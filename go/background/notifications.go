package background

import (
	"context"
	"delivrio.io/go/ent"
	"delivrio.io/go/ent/address"
	"delivrio.io/go/ent/changehistory"
	"delivrio.io/go/ent/colli"
	"delivrio.io/go/ent/connection"
	"delivrio.io/go/ent/connectionbrand"
	"delivrio.io/go/ent/emailtemplate"
	"delivrio.io/go/ent/notification"
	"delivrio.io/go/ent/order"
	"delivrio.io/go/ent/shipment"
	"delivrio.io/go/ent/shipmentparcel"
	"delivrio.io/go/ent/systemevents"
	"delivrio.io/go/mergeutils"
	"delivrio.io/go/schema/hooks/history"
	"delivrio.io/go/shopify/ordersync"
	"delivrio.io/go/utils"
	"delivrio.io/go/viewer"
	"fmt"
	"log"
	"time"
)

const maxNotificationProcessed = 10

func handleNotifications(ctx context.Context) {
	jobContext.Mu.Lock()
	defer jobContext.Mu.Unlock()
	cli := ent.FromContext(ctx)

	allTenants, err := getSystemTenants(ctx)
	if err != nil {
		log.Println("notifications: new system event: ", err)
		return
	}

	for _, t := range allTenants {
		// We want a new System_event ID
		nextCtx := viewer.MergeViewerTenantID(viewer.NewBackgroundContext(ctx), t.ID)
		systemEvent, err := newSystemEvent(nextCtx, systemevents.EventTypeSendNotifications)
		if err != nil {
			log.Printf("notifications: new system event: %s", err)
			continue
		}

		allErrors := make([]error, 0)

		orderConfirmCount, allErr := handleOrderConfirmationEmails(nextCtx)
		if len(allErr) > 0 {
			allErrors = append(allErrors, allErr...)
		}
		labelPackedCount, allErr := handleLabelPackedEmails(nextCtx)
		if len(allErr) > 0 {
			allErrors = append(allErrors, allErr...)
		}
		returnLabelCount, allErr := handleReturnStatusConfirmationLabelEmails(nextCtx)
		if len(allErr) > 0 {
			allErrors = append(allErrors, allErr...)
		}
		returnQRCount, allErr := handleReturnStatusConfirmationQRCodeEmails(nextCtx)
		if len(allErr) > 0 {
			allErrors = append(allErrors, allErr...)
		}
		returnReceivedCount, allErr := handleReturnStatusReceivedEmails(nextCtx)
		if len(allErr) > 0 {
			allErrors = append(allErrors, allErr...)
		}
		returnAcceptedCount, allErr := handleReturnStatusAcceptedEmails(nextCtx)
		if len(allErr) > 0 {
			allErrors = append(allErrors, allErr...)
		}
		shopifyOrderUpdateCount, allErr := handleShopifyOrderUpdate(nextCtx)
		if len(allErr) > 0 {
			allErrors = append(allErrors, allErr...)
		}

		total := orderConfirmCount +
			labelPackedCount +
			returnLabelCount +
			returnQRCount +
			returnReceivedCount +
			returnAcceptedCount +
			shopifyOrderUpdateCount

		status := systemevents.StatusSuccess
		if len(allErrors) > 0 {
			status = systemevents.StatusFail
		}

		log.Println(systemEvent)

		err = cli.SystemEvents.Update().
			Where(systemevents.ID(systemEvent.ID)).
			SetStatus(status).
			SetDescription(fmt.Sprintf("Notifications for %v events successfully fired", total)).
			SetData(utils.JoinErrors(allErrors, ", ")).
			Exec(nextCtx)
		if err != nil {
			log.Println("updating notifications system event failed: ", err)
		}
	}

}

func getSystemTenants(ctx context.Context) ([]*ent.Tenant, error) {
	cli := ent.FromContext(ctx)

	return cli.Tenant.Query().
		All(ctx)
}

func newSystemEvent(ctx context.Context, systemEventType systemevents.EventType) (*ent.SystemEvents, error) {
	cli := ent.FromContext(ctx)
	view := viewer.FromContext(ctx)

	current, err := cli.SystemEvents.Query().
		Where(systemevents.ID(view.ContextID())).
		Only(ctx)
	if err != nil && !ent.IsNotFound(err) {
		return nil, err
	} else if !ent.IsNotFound(err) {
		return current, nil
	}

	return cli.SystemEvents.Create().
		SetStatus(systemevents.StatusRunning).
		SetDescription("Running...").
		SetData("").
		SetUpdatedAt(time.Now()).
		SetEventType(systemEventType).
		SetTenantID(view.TenantID()).
		Save(ctx)
}

func handleShopifyOrderUpdate(ctx context.Context) (int, []error) {
	cli := ent.FromContext(ctx)
	count := 0
	allErrors := make([]error, 0)

	parcelsToUpdate, err := cli.ShipmentParcel.Query().Where(
		shipmentparcel.And(
			shipmentparcel.HasShipmentWith(shipment.StatusNotIn(shipment.StatusDeleted)),
			shipmentparcel.FulfillmentSyncedAtIsNil(),
			shipmentparcel.StatusIn(shipmentparcel.StatusPrinted, shipmentparcel.StatusPending),
			shipmentparcel.HasColliWith(colli.HasOrderWith(
				order.And(
					// Replace with string len > 0?
					order.ExternalIDNotNil(),
					order.HasConnectionWith(connection.HasConnectionBrandWith(connectionbrand.InternalIDEQ(connectionbrand.InternalIDShopify))),
					order.StatusIn(order.StatusDispatched, order.StatusPartially_dispatched),
				)),
			))).
		Order(shipmentparcel.ByID()).
		Limit(20).
		All(ctx)
	if err != nil {
		return 0, append(allErrors, fmt.Errorf("fetching shopify orders update: %w", err))
	}

	for _, p := range parcelsToUpdate {

		ord, err := p.QueryColli().
			QueryOrder().
			Only(ctx)
		if err != nil {
			allErrors = append(allErrors, fmt.Errorf("sync: fulfillment: query order: %w", err))
			continue
		}

		conn, err := ord.
			QueryConnection().
			QueryConnectionShopify().
			Only(ctx)
		if err != nil {
			allErrors = append(allErrors, fmt.Errorf("sync: fulfillment: query connection: %w", err))
			continue
		}

		syncErrors := ordersync.SyncOrderFulfilled(ctx, conn, ord.ExternalID, p)
		if len(syncErrors) > 0 {
			allErrors = append(allErrors, fmt.Errorf("sync: fulfillment: all errors"))
			allErrors = append(allErrors, syncErrors...)
			continue
		}

		count++
	}

	return count, allErrors

}

func handleLabelPackedEmails(ctx context.Context) (int, []error) {
	cli := ent.FromContext(ctx)
	count := 0
	allErrors := make([]error, 0)

	colliesToNotify, err := cli.Colli.Query().
		Where(colli.And(
			colli.EmailLabelPrintedAtIsNil(),
			colli.HasRecipientWith(address.EmailNEQ("")),
			colli.StatusEQ(colli.StatusDispatched),
			colli.HasOrderWith(
				order.HasConnectionWith(
					connection.HasNotificationsWith(
						notification.And(
							notification.HasEmailTemplateWith(
								emailtemplate.MergeTypeEQ(emailtemplate.MergeTypeOrderPicked),
							),
							notification.Active(true),
						),
					),
				),
			))).
		Order(colli.ByID()).
		Limit(maxNotificationProcessed).
		All(ctx)
	if err != nil {
		return count, append(allErrors, fmt.Errorf("error handle label email: %w", err))
	}

	for _, c := range colliesToNotify {
		// Required for the change history
		tenantCtx := viewer.MergeViewerTenantID(ctx, c.TenantID)

		err := mergeutils.SendColliPackedConfirmation(ctx, c)
		if err != nil {
			log.Println("error handle label email: ", err)
			continue
		} else {
			err := c.Update().
				SetEmailLabelPrintedAt(time.Now()).
				Exec(history.NewConfig(tenantCtx).
					SetOrigin(changehistory.OriginBackground).
					SetDescription("Pick & pack email sent to customer").
					Ctx())
			if err != nil {
				allErrors = append(allErrors, fmt.Errorf("error handle label email: %w", err))
				continue
			}
		}
	}

	return count, allErrors
}

func handleOrderConfirmationEmails(ctx context.Context) (int, []error) {
	cli := ent.FromContext(ctx)
	count := 0
	allErrors := make([]error, 0)

	ordersToNotify, err := cli.Order.Query().
		Where(order.And(
			order.EmailSyncConfirmationAtIsNil(),
			order.HasConnectionWith(
				connection.HasNotificationsWith(
					notification.And(
						notification.HasEmailTemplateWith(
							emailtemplate.MergeTypeEQ(emailtemplate.MergeTypeOrderConfirmation),
						),
						notification.Active(true),
					),
				),
			),
		)).
		Order(order.ByID()).
		Limit(maxNotificationProcessed).
		All(ctx)
	if err != nil {
		return count, append(allErrors, fmt.Errorf("error handle order confirmation email: %w", err))
	}

	for _, o := range ordersToNotify {
		// Required for the change history
		tenantCtx := viewer.MergeViewerTenantID(ctx, o.TenantID)

		err := mergeutils.SendOrderConfirmation(ctx, o)
		if err != nil {
			allErrors = append(allErrors, fmt.Errorf("error handle order confirmation email: %w", err))
			continue
		} else {
			err := o.Update().
				SetEmailSyncConfirmationAt(time.Now()).
				Exec(history.NewConfig(tenantCtx).
					SetOrigin(changehistory.OriginBackground).
					SetDescription("Order confirmation email sent to customer").
					Ctx())
			if err != nil {
				allErrors = append(allErrors, fmt.Errorf("error handle order confirmation email: %w", err))
				continue
			}
		}
	}

	return count, allErrors
}
