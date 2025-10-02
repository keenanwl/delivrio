package seed

import (
	"context"
	"delivrio.io/go/ent"
	"delivrio.io/go/ent/carrierbrand"
)

func BringServices(ctx context.Context) {
	tx := ent.TxFromContext(ctx)

	cb := tx.CarrierBrand.Query().
		Where(carrierbrand.InternalIDEQ(carrierbrand.InternalIDBring)).
		OnlyX(ctx)

	cs := tx.CarrierService.Create().
		SetCarrierBrand(cb).
		SetLabel("PickUp Parcel").
		SetInternalID("BRING_BRING_PICKUP_PARCEL").
		SetDeliveryPointRequired(true).
		SetDeliveryPointRequired(false).
		SaveX(ctx)
	tx.CarrierServiceBring.Create().
		SetCarrierService(cs).
		SetAPIServiceCode("0340").
		SetAPIRequest("PICKUP_PARCEL").
		ExecX(ctx)

	cs = tx.CarrierService.Create().
		SetCarrierBrand(cb).
		SetLabel("PickUp Parcel Bulk").
		SetDeliveryPointRequired(true).
		SetDeliveryPointRequired(false).
		SetInternalID("BRING_PICKUP_PARCEL_BULK").
		SaveX(ctx)
	tx.CarrierServiceBring.Create().
		SetCarrierService(cs).
		SetAPIServiceCode("0342").
		SetAPIRequest("PICKUP_PARCEL_BULK").
		ExecX(ctx)

	cs = tx.CarrierService.Create().
		SetCarrierBrand(cb).
		SetLabel("Home Delivery Parcel").
		SetInternalID("BRING_HOME_DELIVERY_PARCEL").
		SaveX(ctx)
	tx.CarrierServiceBring.Create().
		SetCarrierService(cs).
		SetAPIServiceCode("0349").
		SetAPIRequest("HOME_DELIVERY_PARCEL").
		ExecX(ctx)

	cs = tx.CarrierService.Create().
		SetCarrierBrand(cb).
		SetLabel("Business Parcel").
		SetInternalID("BRING_BUSINESS_PARCEL").
		SaveX(ctx)
	tx.CarrierServiceBring.Create().
		SetCarrierService(cs).
		SetAPIServiceCode("0330").
		SetAPIRequest("BUSINESS_PARCEL").
		ExecX(ctx)

	cs = tx.CarrierService.Create().
		SetCarrierBrand(cb).
		SetLabel("Business Parcel Bulk").
		SetInternalID("BRING_BUSINESS_PARCEL_BULK").
		SaveX(ctx)
	tx.CarrierServiceBring.Create().
		SetCarrierService(cs).
		SetAPIServiceCode("0332").
		SetAPIRequest("BUSINESS_PARCEL_BULK").
		ExecX(ctx)

	cs = tx.CarrierService.Create().
		SetCarrierBrand(cb).
		SetLabel("Express Nordic 09:00 Bulk").
		SetInternalID("BRING_EXPRESS_NORDIC_0900_BULK").
		SaveX(ctx)
	tx.CarrierServiceBring.Create().
		SetCarrierService(cs).
		SetAPIServiceCode("0334").
		SetAPIRequest("EXPRESS_NORDIC_0900_BULK").
		ExecX(ctx)

	cs = tx.CarrierService.Create().
		SetCarrierBrand(cb).
		SetLabel("Business Pallet").
		SetInternalID("BRING_BUSINESS_PALLET").
		SaveX(ctx)
	tx.CarrierServiceBring.Create().
		SetCarrierService(cs).
		SetAPIServiceCode("0336").
		SetAPIRequest("BUSINESS_PALLET").
		ExecX(ctx)

	cs = tx.CarrierService.Create().
		SetCarrierBrand(cb).
		SetLabel("Business Pallet (1/2 pallet)").
		SetInternalID("BRING_BUSINESS_HALFPALLET").
		SaveX(ctx)
	tx.CarrierServiceBring.Create().
		SetCarrierService(cs).
		SetAPIServiceCode("0336").
		SetAPIRequest("BUSINESS_HALFPALLET").
		ExecX(ctx)

	cs = tx.CarrierService.Create().
		SetCarrierBrand(cb).
		SetLabel("Business Pallet (1/4 pallet)").
		SetInternalID("BRING_BUSINESS_QUARTERPALLET").
		SaveX(ctx)
	tx.CarrierServiceBring.Create().
		SetCarrierService(cs).
		SetAPIServiceCode("0336").
		SetAPIRequest("BUSINESS_QUARTERPALLET").
		ExecX(ctx)

	cs = tx.CarrierService.Create().
		SetReturn(true).
		SetCarrierBrand(cb).
		SetLabel("PickUp Parcel Return").
		SetInternalID("BRING_PICKUP_PARCEL_RETURN").
		SaveX(ctx)
	tx.CarrierServiceBring.Create().
		SetCarrierService(cs).
		SetAPIServiceCode("0341").
		SetAPIRequest("PICKUP_PARCEL_RETURN").
		ExecX(ctx)

	cs = tx.CarrierService.Create().
		SetReturn(true).
		SetCarrierBrand(cb).
		SetLabel("PickUp Parcel Return Bulk").
		SetInternalID("BRING_PICKUP_PARCEL_RETURN_BULK").
		SaveX(ctx)
	tx.CarrierServiceBring.Create().
		SetCarrierService(cs).
		SetAPIServiceCode("0343").
		SetAPIRequest("PICKUP_PARCEL_RETURN_BULK").
		ExecX(ctx)

	cs = tx.CarrierService.Create().
		SetReturn(true).
		SetCarrierBrand(cb).
		SetLabel("Business Parcel Return").
		SetInternalID("BRING_BUSINESS_PARCEL_RETURN").
		SaveX(ctx)
	tx.CarrierServiceBring.Create().
		SetCarrierService(cs).
		SetAPIServiceCode("0331").
		SetAPIRequest("BUSINESS_PARCEL_RETURN").
		ExecX(ctx)

	cs = tx.CarrierService.Create().
		SetReturn(true).
		SetCarrierBrand(cb).
		SetLabel("Business Parcel Return Bulk").
		SetInternalID("BRING_BUSINESS_PARCEL_RETURN_BULK").
		SaveX(ctx)
	tx.CarrierServiceBring.Create().
		SetCarrierService(cs).
		SetAPIServiceCode("0333").
		SetAPIRequest("BUSINESS_PARCEL_RETURN_BULK").
		ExecX(ctx)

	cs = tx.CarrierService.Create().
		SetCarrierBrand(cb).
		SetLabel("Home Delivery Single Indoor").
		SetInternalID("BRING_SINGLE_INDOOR").
		SaveX(ctx)
	tx.CarrierServiceBring.Create().
		SetCarrierService(cs).
		SetAPIServiceCode("3150").
		SetAPIRequest("SINGLE_INDOOR").
		ExecX(ctx)

	cs = tx.CarrierService.Create().
		SetCarrierBrand(cb).
		SetLabel("Home Delivery Indoor").
		SetInternalID("BRING_DOUBLE_INDOOR").
		SaveX(ctx)
	tx.CarrierServiceBring.Create().
		SetCarrierService(cs).
		SetAPIServiceCode("2870").
		SetAPIRequest("DOUBLE_INDOOR").
		ExecX(ctx)

	cs = tx.CarrierService.Create().
		SetCarrierBrand(cb).
		SetLabel("Home Delivery Curbside").
		SetInternalID("BRING_CURBSIDE").
		SaveX(ctx)
	tx.CarrierServiceBring.Create().
		SetCarrierService(cs).
		SetAPIServiceCode("3123").
		SetAPIRequest("CURBSIDE").
		ExecX(ctx)

	cs = tx.CarrierService.Create().
		SetCarrierBrand(cb).
		SetLabel("Home Delivery Curbside Evening").
		SetInternalID("BRING_CURBSIDE_EVENING").
		SaveX(ctx)
	tx.CarrierServiceBring.Create().
		SetCarrierService(cs).
		SetAPIServiceCode("3457").
		SetAPIRequest("CURBSIDE_EVENING").
		ExecX(ctx)

	cs = tx.CarrierService.Create().
		SetCarrierBrand(cb).
		SetLabel("Urban Home Delivery").
		SetInternalID("BRING_3332").
		SaveX(ctx)
	tx.CarrierServiceBring.Create().
		SetCarrierService(cs).
		SetAPIServiceCode("3332").
		SetAPIRequest("3332").
		ExecX(ctx)

	cs = tx.CarrierService.Create().
		SetReturn(true).
		SetCarrierBrand(cb).
		SetLabel("Home Delivery Return").
		SetInternalID("BRING_HOME_DELIVERY_RETURN").
		SaveX(ctx)
	tx.CarrierServiceBring.Create().
		SetCarrierService(cs).
		SetAPIServiceCode("HOME_DELIVERY_RETURN").
		SetAPIRequest("2778").
		ExecX(ctx)

	cs = tx.CarrierService.Create().
		SetReturn(true).
		SetCarrierBrand(cb).
		SetLabel("Indoor Return").
		SetInternalID("BRING_RETURN_INDOOR").
		SaveX(ctx)
	tx.CarrierServiceBring.Create().
		SetCarrierService(cs).
		SetAPIServiceCode("RETURN_INDOOR").
		SetAPIRequest("3578").
		ExecX(ctx)

	cs = tx.CarrierService.Create().
		SetReturn(true).
		SetCarrierBrand(cb).
		SetLabel("Curbside Return").
		SetInternalID("BRING_RETURN_CURBSIDE").
		SaveX(ctx)
	tx.CarrierServiceBring.Create().
		SetCarrierService(cs).
		SetAPIServiceCode("RETURN_CURBSIDE").
		SetAPIRequest("3577").
		ExecX(ctx)

}
