package seed

import (
	"context"
	"delivrio.io/go/ent"
	"delivrio.io/go/ent/carrierbrand"
)

type DAOService string

var (
	DAOServiceHOME       DAOService = "DAO_HOME"
	DAOServiceSHOP       DAOService = "DAO_SHOP"
	DAOServiceSHOPReturn DAOService = "DAO_SHOP_RETURN"
)

func (s DAOService) String() string {
	return string(s)
}

func DAOServices(ctx context.Context) {
	tx := ent.TxFromContext(ctx)

	cb := tx.CarrierBrand.Query().
		Where(carrierbrand.InternalIDEQ(carrierbrand.InternalIDDAO)).
		OnlyX(ctx)

	csHome := tx.CarrierService.Create().
		SetCarrierBrand(cb).
		SetLabel("daoHOME").
		SetInternalID(DAOServiceHOME.String()).
		SaveX(ctx)
	csHomeDao := tx.CarrierServiceDAO.Create().
		SetCarrierService(csHome).
		SaveX(ctx)

	csShop := tx.CarrierService.Create().
		SetCarrierBrand(cb).
		SetLabel("daoSHOP").
		SetDeliveryPointRequired(true).
		SetDeliveryPointOptional(false).
		SetInternalID(DAOServiceSHOP.String()).
		SaveX(ctx)
	csShopDao := tx.CarrierServiceDAO.Create().
		SetCarrierService(csShop).
		SaveX(ctx)

	csReturn := tx.CarrierService.Create().
		SetCarrierBrand(cb).
		SetReturn(true).
		SetLabel("daoSHOP Return").
		SetInternalID(DAOServiceSHOPReturn.String()).
		SaveX(ctx)
	tx.CarrierServiceDAO.Create().
		SetCarrierService(csReturn).
		ExecX(ctx)

	tx.CarrierAdditionalServiceDAO.Create().
		SetLabel("Priority").
		SetAPICode("priority").
		AddCarrierServiceDAO(csShopDao).
		ExecX(ctx)

	tx.CarrierAdditionalServiceDAO.Create().
		SetLabel("Same day").
		SetAPICode("sameday").
		AddCarrierServiceDAO(csHomeDao).
		ExecX(ctx)

	tx.CarrierAdditionalServiceDAO.Create().
		SetLabel("Automatically select pickup point").
		SetAPICode("auto-pickup-point").
		AddCarrierServiceDAO(csShopDao).
		ExecX(ctx)
}
