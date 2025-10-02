package ratelookup

import (
	"context"
	"delivrio.io/go/ent"
	"delivrio.io/go/ent/carrierservice"
	"delivrio.io/go/ent/deliveryoption"
	"delivrio.io/go/ent/hypothesistest"
	"delivrio.io/go/ent/hypothesistestdeliveryoptionrequest"
	"delivrio.io/go/ent/predicate"
	"delivrio.io/go/schema/mixins"
	"delivrio.io/go/viewer"
	"delivrio.io/shared-utils/pulid"
	"encoding/json"
	"entgo.io/ent/dialect/sql"
	"fmt"
	"hash/fnv"
	"log"
	"math/rand"
	"time"
)

func filterByHypothesisTestAndCompanyField(
	ctx context.Context,
	now time.Time,
	connect *ent.Connection,
	reqOrder interface{},
	reqToAddress interface{},
	hasCompanyField bool,
) ([]*ent.DeliveryOption, *pulid.ID, error) {
	db := ent.FromContext(ctx)
	view := viewer.FromContext(ctx)

	ht, err := connect.QueryHypothesisTest().
		Where(hypothesistest.Active(true)).
		QueryHypothesisTestDeliveryOption().
		WithDeliveryOptionGroupOne().
		WithDeliveryOptionGroupTwo().
		All(ctx)
	if err != nil {
		return nil, nil, err
	}

	if len(ht) > 1 {
		return nil, nil, fmt.Errorf("expected only one active hypothesis test for this connection")
	}

	var trackingID *pulid.ID
	excludeIDs := make([]pulid.ID, 0)

	if len(ht) == 1 {

		isControl := true

		orderJson, err := json.Marshal(reqOrder)
		if err != nil {
			return nil, nil, err
		}
		orderHash, err := fastHashIt(orderJson)
		if err != nil {
			return nil, nil, err
		}
		adrJson, err := json.Marshal(reqToAddress)
		if err != nil {
			return nil, nil, err
		}
		adrHash, err := fastHashIt(adrJson)
		if err != nil {
			return nil, nil, err
		}

		currentRequest, err := db.HypothesisTestDeliveryOptionRequest.Query().
			Where(hypothesistestdeliveryoptionrequest.ShippingAddressHashEQ(adrHash)).
			// Since the hash might have collisions, we allow for multiple records
			Order(hypothesistestdeliveryoptionrequest.ByLastRequestedAt(sql.OrderDesc())).
			First(ctx)
		if !ent.IsNotFound(err) && err != nil {
			return nil, nil, err
		}

		isByOrder := ht[0].ByOrder
		if !ent.IsNotFound(err) && currentRequest != nil {
			isControl = currentRequest.IsControlGroup
		} else if isByOrder {
			// It's a number between o and n-1
			isControl = rand.Intn(2) == 1
		} else if ht[0].ByIntervalRotation {
			interval := ht[0].RotationIntervalHours
			isControl = (now.Hour()/interval)%2 != 0
			isOddWeekDay := now.Weekday() % 2
			if isOddWeekDay != 0 {
				isControl = !isControl
			}
		}

		groupOneIDs := make([]pulid.ID, 0)
		groupTwoIDs := make([]pulid.ID, 0)
		for _, do := range ht[0].Edges.DeliveryOptionGroupOne {
			groupOneIDs = append(groupOneIDs, do.ID)
		}
		for _, do := range ht[0].Edges.DeliveryOptionGroupTwo {
			groupTwoIDs = append(groupTwoIDs, do.ID)
		}

		if isControl {
			excludeIDs = groupTwoIDs
		} else {
			excludeIDs = groupOneIDs
		}

		if currentRequest != nil {
			err = currentRequest.Update().
				SetRequestCount(currentRequest.RequestCount + 1).
				Exec(ctx)
			if err != nil {
				log.Println("err increasing request count", err)
			}
			trackingID = &currentRequest.ID
		} else {
			tracking, err := db.HypothesisTestDeliveryOptionRequest.Create().
				SetShippingAddressHash(adrHash).
				SetOrderHash(orderHash).
				SetIsControlGroup(isControl).
				SetHypothesisTestDeliveryOption(ht[0]).
				SetRequestCount(1).
				SetTenantID(view.TenantID()).
				Save(ctx)
			if err != nil {
				return nil, nil, err
			}
			trackingID = &tracking.ID
		}

	}

	predDeliveryOption := make([]predicate.DeliveryOption, 0)
	predDeliveryOption = append(
		predDeliveryOption,
		deliveryoption.IDNotIn(excludeIDs...),
		deliveryoption.HideDeliveryOptionEQ(false),
		deliveryoption.HasCarrierServiceWith(carrierservice.Return(false)),
	)

	if !hasCompanyField {
		predDeliveryOption = append(predDeliveryOption, deliveryoption.HideIfCompanyEmpty(false))
	}

	do, err := connect.QueryDeliveryOption().
		Where(
			deliveryoption.And(predDeliveryOption...),
		).
		Order(deliveryoption.BySortOrder(sql.OrderAsc())).
		All(mixins.ExcludeArchived(ctx))

	return do, trackingID, err
}

func fastHashIt(byt []byte) (string, error) {
	hash := fnv.New32a()
	_, err := hash.Write(byt)
	if err != nil {
		return "", err
	}
	return fmt.Sprintf("%v", hash.Sum32()), nil
}
