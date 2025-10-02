package ordersync

import (
	"context"
	"delivrio.io/go/ent"
	"delivrio.io/go/ent/changehistory"
	"delivrio.io/go/ent/colli"
	"delivrio.io/go/ent/connection"
	"delivrio.io/go/ent/order"
	"delivrio.io/go/ent/shipment"
	"delivrio.io/go/ent/shipmentparcel"
	"delivrio.io/go/ent/systemevents"
	"delivrio.io/go/schema/hooks/history"
	"delivrio.io/go/utils"
	"delivrio.io/shared-utils/pulid"
	"fmt"
	"log"
	"strconv"
	"time"
)

// Sync orders which have been cancelled in Shopify

const apiDate = "2024-04"

func ProcessShopifyOrderCancelledSync(ctx context.Context, connectShopify *ent.ConnectionShopify, lastSyncEvent time.Time, evt pulid.ID) {
	db := ent.FromContext(ctx)

	fetchURLBase := fmt.Sprintf(
		// When updating a fulfillment, we are updating the update_at value...
		"%s/admin/api/%s/orders.json?status=cancelled&fields=id,name&updated_at_min=%s",
		connectShopify.StoreURL,
		apiDate,
		// UTC() is a hack because the Shopify API isn't working as documented???
		lastSyncEvent.UTC().Format(time.RFC3339),
	)

	orderCount, err := pageCancelledOrders(ctx, fetchURLBase, connectShopify)
	if err != nil {
		setCancelled(
			ctx,
			fmt.Sprintf("Failed fetching orders, after syncing: %v orders", orderCount),
			fmt.Sprintf("%v -> %v", fetchURLBase, err.Error()),
			evt,
		)
	}

	err = db.SystemEvents.Update().
		SetStatus(systemevents.StatusSuccess).
		SetDescription(fmt.Sprintf("Successfully synced %v cancelled orders", orderCount)).
		Where(systemevents.ID(evt)).
		Exec(ctx)
	if err != nil {
		log.Println(err.Error())
	}
}

func setCancelled(ctx context.Context, description string, data string, systemEventID pulid.ID) {
	db := ent.FromContext(ctx)
	err := db.SystemEvents.Update().
		SetStatus(systemevents.StatusFail).
		SetDescription(description).
		SetData(data).
		Where(systemevents.ID(systemEventID)).
		Exec(ctx)
	if err != nil {
		log.Println("could not save cancelled error to DB: %w", err)
	}
}

func pageCancelledOrders(ctx context.Context, fetchURLbase string, connectShopify *ent.ConnectionShopify) (int, error) {
	var sinceID uint64 = 0
	fetchURL := fmt.Sprintf(fetchURLbase+"&since_id=%v", sinceID)
	orders, err := fetchShopifyOrdersPage(fetchURL, connectShopify.APIKey)
	if err != nil {
		return 0, fmt.Errorf("sync: shopify: cancelled: %w", err)
	}

	connect, err := connectShopify.QueryConnection().
		Only(ctx)
	if err != nil {
		return 0, fmt.Errorf("sync: shopify: cancelled: %w", err)
	}

	orderCount := len(orders.Orders)

	for len(orders.Orders) > 0 {
		select {
		case <-ctx.Done():
			return orderCount, fmt.Errorf("sync: shopify: cancelled: context was cancelled before completed")
		default:
			for _, o := range orders.Orders {
				if o.ID > sinceID {
					sinceID = o.ID
				}

				err = setOrderCancelled(ctx, connect.ID, o.ID)
				if err != nil {
					return orderCount, fmt.Errorf("error: sync: shopify: cancelled: set cancelled: %w", err)
				}
			}
		}

		fetchURL = fmt.Sprintf(fetchURLbase+"&since_id=%v", sinceID)
		orders, err = fetchShopifyOrdersPage(fetchURL, connectShopify.APIKey)
		if err != nil {
			return orderCount, fmt.Errorf("sync: shopify: cancelled: %w", err)
		}

		orderCount += len(orders.Orders)
	}

	return orderCount, nil
}

func setOrderCancelled(ctx context.Context, connectionID pulid.ID, shopifyOrderID uint64) error {
	cli := ent.FromContext(ctx)
	ord, err := cli.Order.Query().
		Where(
			order.And(
				order.HasConnectionWith(connection.ID(connectionID)),
				order.ExternalID(strconv.FormatUint(shopifyOrderID, 10)),
			),
		).Only(ctx)
	if err != nil {
		return err
	}

	orderHasActiveShipment, err := ord.QueryColli().
		QueryShipmentParcel().
		Where(shipmentparcel.HasShipmentWith(shipment.StatusNEQ(shipment.StatusDeleted))).
		Exist(ctx)
	if err != nil {
		return err
	}

	if orderHasActiveShipment {
		return nil
	}

	tx, err := cli.Tx(ctx)
	if err != nil {
		return utils.Rollback(tx, err)
	}
	defer tx.Rollback()

	// The hook should cancel the order itself after all collis are cancelled
	err = tx.Colli.Update().
		SetStatus(colli.StatusCancelled).
		Where(colli.HasOrderWith(order.ID(ord.ID))).
		Exec(history.FromContext(ctx).SetDescription("Sync cancel from Shopify").SetOrigin(changehistory.OriginBackground).Ctx())
	if err != nil {
		return utils.Rollback(tx, err)
	}

	return tx.Commit()

}
