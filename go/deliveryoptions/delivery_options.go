package deliveryoptions

import (
	"context"
	"delivrio.io/go/ent"
	"delivrio.io/go/ent/carrieradditionalservicegls"
	"delivrio.io/go/ent/carrieradditionalservicepostnord"
	"delivrio.io/go/ent/carrierbrand"
	"delivrio.io/go/ent/carrierservice"
	"delivrio.io/go/ent/deliveryoption"
	"delivrio.io/go/ent/productvariant"
	"delivrio.io/go/ent/returncolli"
	"delivrio.io/go/seed"
	"delivrio.io/shared-utils/pulid"
	"fmt"
	"math"
)

func FromReturnColliID(ctx context.Context, returnColliID pulid.ID) ([]*DeliveryOptionBrandName, error) {
	cli := ent.FromContext(ctx)

	orderLines, err := cli.ReturnColli.Query().
		Where(returncolli.ID(returnColliID)).
		QueryReturnOrderLine().
		QueryOrderLine().
		WithProductVariant().
		All(ctx)
	if err != nil {
		return nil, err
	}

	productLines := make([]*DeliveryOptionProductLineInput, 0)
	for _, ol := range orderLines {
		productLines = append(productLines, &DeliveryOptionProductLineInput{
			ProductVariantID: ol.ProductVariantID,
			Units:            ol.Units,
			UnitPrice:        ol.UnitPrice,
		})
	}

	connectionID, err := cli.ReturnColli.Query().
		Where(returncolli.ID(returnColliID)).
		QueryOrder().
		QueryConnection().
		OnlyID(ctx)
	if err != nil {
		return nil, err
	}

	customerAddress, err := cli.ReturnColli.Query().
		Where(returncolli.ID(returnColliID)).
		QueryRecipient().
		WithCountry().
		Only(ctx)
	if err != nil {
		return nil, err
	}

	s := DeliveryOptionSeedInput{
		ConnectionID: connectionID,
		Country:      customerAddress.Edges.Country.ID,
		Zip:          customerAddress.Zip,
		ProductLines: productLines,
	}

	return ByOrder(ctx, s, true)
}

func expandProductMetadata(ctx context.Context, productLines []*DeliveryOptionProductLineInput) ([]*ConstraintProductWeight, error) {
	cli := ent.FromContext(ctx)

	output := make([]*ConstraintProductWeight, 0)
	for _, pl := range productLines {
		pv, err := cli.ProductVariant.Query().
			WithProduct().
			Where(productvariant.ID(pl.ProductVariantID)).
			Only(ctx)
		if err != nil {
			return nil, err
		}

		tags, err := pv.Edges.Product.QueryProductTags().
			IDs(ctx)
		if err != nil {
			return nil, err
		}

		output = append(output, &ConstraintProductWeight{
			WeightG:       *pv.WeightG,
			ProductTagIDs: tags,
			SKU:           pv.EanNumber,
			UnitPrice:     pl.UnitPrice,
			Units:         pl.Units,
		})
	}
	return output, nil
}

func ByOrder(ctx context.Context, orderInfo DeliveryOptionSeedInput, returns bool) ([]*DeliveryOptionBrandName, error) {
	cli := ent.FromContext(ctx)
	deliveryOptions, err := cli.DeliveryOption.Query().
		WithCarrier().
		Where(deliveryoption.And(
			deliveryoption.HideDeliveryOptionEQ(false),
			deliveryoption.HasCarrierServiceWith(
				carrierservice.ReturnEQ(returns),
			)),
		).All(ctx)
	if err != nil {
		return nil, err
	}

	extendedOrderLineInfo, err := expandProductMetadata(ctx, orderInfo.ProductLines)
	if err != nil {
		return nil, err
	}

	output := make([]*DeliveryOptionBrandName, 0)
	for _, do := range deliveryOptions {

		matches, price, err := DeliveryOptionMatches(ctx, do, orderInfo.Zip, extendedOrderLineInfo, orderInfo.Country)
		servicePointOptional, servicePointRequired, err := ServicePointConfig(ctx, do)
		if err != nil {
			return nil, err
		}

		if matches {
			output = append(
				output,
				&DeliveryOptionBrandName{
					DeliveryOptionID:      do.ID,
					Name:                  do.Name,
					Description:           do.Description,
					Status:                DeliveryOptionBrandNameStatusAvailable,
					Price:                 fmt.Sprintf("%v", math.Ceil(price.Price)),
					Currency:              price.Currency,
					Warning:               nil,
					RequiresDeliveryPoint: servicePointRequired,
					DeliveryPoint:         servicePointOptional,
					ClickAndCollect:       do.ClickCollect,
				},
			)
		} else if !returns {
			warning := "Delivery option not matched"
			output = append(
				output,
				&DeliveryOptionBrandName{
					DeliveryOptionID:      do.ID,
					Name:                  do.Name,
					Description:           do.Description,
					Status:                DeliveryOptionBrandNameStatusNotAvailable,
					Price:                 "-",
					Currency:              nil,
					Warning:               &warning,
					RequiresDeliveryPoint: servicePointRequired,
					DeliveryPoint:         servicePointOptional,
					ClickAndCollect:       do.ClickCollect,
				},
			)
		}
	}

	return output, nil
}

func ServicePointConfig(ctx context.Context, do *ent.DeliveryOption) (bool, bool, error) {
	servicePointOptional := false
	servicePointRequired := false

	cb, err := do.QueryCarrier().
		QueryCarrierBrand().
		Only(ctx)
	if err != nil {
		return false, false, err
	}
	switch cb.InternalID {
	case carrierbrand.InternalIDPostNord:
		servicePointService, err := do.QueryCarrierService().
			QueryCarrierServicePostNord().
			QueryCarrierAddServPostNord().
			Where(carrieradditionalservicepostnord.And(
				carrieradditionalservicepostnord.InternalIDEQ(seed.PostNordDeliveryPointServiceInternalID),
			)).Only(ctx)
		if err != nil && !ent.IsNotFound(err) {
			return false, false, err
		} else if ent.IsNotFound(err) {
			servicePointOptional = false
		} else {
			servicePointOptional = true
			servicePointRequired = servicePointService.Mandatory
		}
		break
	case carrierbrand.InternalIDGLS:
		servicePointService, err := do.QueryCarrierService().
			QueryCarrierServiceGLS().
			QueryCarrierAdditionalServiceGLS().
			Where(
				carrieradditionalservicegls.InternalIDEQ(seed.GLSDeliveryPointServiceInternalID),
			).Only(ctx)
		if err != nil && !ent.IsNotFound(err) {
			return false, false, err
			// Logic below might be wrong
		} else if ent.IsNotFound(err) {
			servicePointOptional = false
		} else {
			servicePointOptional = true
			servicePointRequired = servicePointService.Mandatory
		}
		break
	case carrierbrand.InternalIDBring:
		// Should always be defined
		carrierService, err := do.CarrierService(ctx)
		if err != nil {
			return false, false, err
		}
		servicePointRequired = carrierService.DeliveryPointRequired
		servicePointOptional = carrierService.DeliveryPointOptional
		break
	case carrierbrand.InternalIDDAO:
		// Should always be defined
		carrierService, err := do.CarrierService(ctx)
		if err != nil {
			return false, false, err
		}
		servicePointRequired = carrierService.DeliveryPointRequired
		servicePointOptional = carrierService.DeliveryPointOptional
		break
	case carrierbrand.InternalIDUSPS:
		// No USPS service points supported yet
		return false, false, nil
	case carrierbrand.InternalIDEasyPost:
		// No EP service points supported yet
		return false, false, nil
	case carrierbrand.InternalIDDF:
		// No DF service points supported yet
		return false, false, nil
	default:
		return false, false, fmt.Errorf("deliveryoptions: unsupported carrier brand %v", cb.InternalID)
	}
	return servicePointOptional, servicePointRequired, nil
}
