package seed

import (
	"context"
	"delivrio.io/go/ent"
	"delivrio.io/go/ent/carrierbrand"
)

func DFServices(ctx context.Context) {
	tx := ent.TxFromContext(ctx)

	brand := tx.CarrierBrand.Query().
		Where(carrierbrand.InternalIDEQ(carrierbrand.InternalIDDF)).
		OnlyX(ctx)

	cs1 := tx.CarrierService.Create().
		SetLabel("Pick-up - Cargo").
		SetInternalID("DF_PICK_UP_CARGO").
		SetConsolidation(true).
		SetCarrierBrand(brand).
		SaveX(ctx)
	tx.CarrierServiceDF.Create().
		SetCarrierService(cs1).
		ExecX(ctx)

	cs1 = tx.CarrierService.Create().
		SetLabel("Pick-up - Pallet").
		SetInternalID("DF_PICK_UP_PALLET").
		SetConsolidation(true).
		SetCarrierBrand(brand).
		SaveX(ctx)
	tx.CarrierServiceDF.Create().
		SetCarrierService(cs1).
		ExecX(ctx)

	cs1 = tx.CarrierService.Create().
		SetLabel("Pick-up - Vehicle Parcel").
		SetInternalID("DF_PICK_UP_VEHICLE_PARCEL").
		SetConsolidation(true).
		SetCarrierBrand(brand).
		SaveX(ctx)
	tx.CarrierServiceDF.Create().
		SetCarrierService(cs1).
		ExecX(ctx)

	cs1 = tx.CarrierService.Create().
		SetLabel("Return - Cargo").
		SetReturn(true).
		SetInternalID("DF_RETURN_CARGO").
		SetConsolidation(true).
		SetCarrierBrand(brand).
		SaveX(ctx)
	tx.CarrierServiceDF.Create().
		SetCarrierService(cs1).
		ExecX(ctx)

	cs1 = tx.CarrierService.Create().
		SetLabel("Return - Pallet").
		SetInternalID("DF_RETURN_PALLET").
		SetReturn(true).
		SetConsolidation(true).
		SetCarrierBrand(brand).
		SaveX(ctx)
	tx.CarrierServiceDF.Create().
		SetCarrierService(cs1).
		ExecX(ctx)

	cs1 = tx.CarrierService.Create().
		SetLabel("Return - Vehicle Parcel").
		SetInternalID("DF_RETURN_VEHICLE_PARCEL").
		SetReturn(true).
		SetConsolidation(true).
		SetCarrierBrand(brand).
		SaveX(ctx)
	tx.CarrierServiceDF.Create().
		SetCarrierService(cs1).
		ExecX(ctx)

	cs1 = tx.CarrierService.Create().
		SetLabel("Relocation - Cargo").
		SetInternalID("DF_RELOCATION_CARGO").
		SetConsolidation(true).
		SetCarrierBrand(brand).
		SaveX(ctx)
	tx.CarrierServiceDF.Create().
		SetCarrierService(cs1).
		ExecX(ctx)

	cs1 = tx.CarrierService.Create().
		SetLabel("Relocation - Pallet").
		SetInternalID("DF_RELOCATION_PALLET").
		SetConsolidation(true).
		SetCarrierBrand(brand).
		SaveX(ctx)
	tx.CarrierServiceDF.Create().
		SetCarrierService(cs1).
		ExecX(ctx)

	cs1 = tx.CarrierService.Create().
		SetLabel("Relocation - Vehicle Parcel").
		SetInternalID("DF_RELOCATION_VEHICLE_PARCEL").
		SetConsolidation(true).
		SetCarrierBrand(brand).
		SaveX(ctx)
	tx.CarrierServiceDF.Create().
		SetCarrierService(cs1).
		ExecX(ctx)
}
