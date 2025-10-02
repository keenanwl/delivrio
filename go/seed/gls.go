package seed

import (
	"context"
	"delivrio.io/go/ent"
	"delivrio.io/go/ent/carrierbrand"
	"delivrio.io/go/ent/carrierservicegls"
	"delivrio.io/go/ent/country"
)

const GLSDeliveryPointServiceInternalID = "service_point"

func GLSServices(ctx context.Context) {
	tx := ent.TxFromContext(ctx)

	glsBrand := tx.CarrierBrand.Query().
		Where(carrierbrand.InternalIDEQ(carrierbrand.InternalIDGLS)).
		OnlyX(ctx)

	cs1 := tx.CarrierService.Create().
		SetLabel("BusinessParcel").
		SetInternalID("GLS_BUSINESS_PARCEL").
		SetCarrierBrand(glsBrand).
		SaveX(ctx)
	tx.CarrierServiceGLS.Create().
		SetCarrierService(cs1).
		SetNillableAPIKey(nil).
		SetAPIValue(carrierservicegls.APIValueNone).
		ExecX(ctx)

	cs2 := tx.CarrierService.Create().
		SetLabel("EuroBusinessParcel").
		SetInternalID("GLS_EURO_BUSINESS_PARCEL").
		SetCarrierBrand(glsBrand).
		SaveX(ctx)
	tx.CarrierServiceGLS.Create().
		SetCarrierService(cs2).
		SetNillableAPIKey(nil).
		SetAPIValue(carrierservicegls.APIValueNone).
		ExecX(ctx)

	cs3 := tx.CarrierService.Create().
		SetLabel("DirectShopService").
		SetInternalID("GLS_DIRECT_SHOP_SERVICE").
		SetCarrierBrand(glsBrand).
		SaveX(ctx)
	tx.CarrierServiceGLS.Create().
		SetCarrierService(cs3).
		SetAPIKey("DirectShop").
		SetAPIValue(carrierservicegls.APIValueY).
		ExecX(ctx)

	cs4 := tx.CarrierService.Create().
		SetLabel("ShopDeliveryService").
		SetInternalID("GLS_SHOP_DELIVERY_SERVICE").
		SetCarrierBrand(glsBrand).
		SaveX(ctx)
	shopDeliveryService := tx.CarrierServiceGLS.Create().
		SetCarrierService(cs4).
		SetAPIKey("ShopDelivery").
		SetAPIValue(carrierservicegls.APIValueNumericString).
		SaveX(ctx)

	cs5 := tx.CarrierService.Create().
		SetLabel("Express12Service").
		SetInternalID("GLS_EXPRESS_12_SERVICE").
		SetCarrierBrand(glsBrand).
		SaveX(ctx)
	tx.CarrierServiceGLS.Create().
		SetCarrierService(cs5).
		SetAPIKey("Express12").
		SetAPIValue(carrierservicegls.APIValueY).
		ExecX(ctx)

	cs6 := tx.CarrierService.Create().
		SetLabel("Express10Service").
		SetInternalID("GLS_EXPRESS_10_SERVICE").
		SetCarrierBrand(glsBrand).
		SaveX(ctx)
	tx.CarrierServiceGLS.Create().
		SetCarrierService(cs6).
		SetAPIKey("Express10").
		SetAPIValue(carrierservicegls.APIValueY).
		ExecX(ctx)

	cs7 := tx.CarrierService.Create().
		SetLabel("ShopReturnService").
		SetReturn(true).
		SetInternalID("GLS_SHOP_RETURN_SERVICE").
		SetCarrierBrand(glsBrand).
		SaveX(ctx)
	tx.CarrierServiceGLS.Create().
		SetCarrierService(cs7).
		SetAPIKey("ShopReturn").
		SetAPIValue(carrierservicegls.APIValueY).
		ExecX(ctx)

	cs8 := tx.CarrierService.Create().
		SetLabel("PrivateDeliveryService").
		SetInternalID("GLS_PRIVATE_DELIVERY_SERVICE").
		SetCarrierBrand(glsBrand).
		SaveX(ctx)
	tx.CarrierServiceGLS.Create().
		SetCarrierService(cs8).
		SetAPIKey("PrivateDelivery").
		SetAPIValue(carrierservicegls.APIValueY).
		ExecX(ctx)

	cs9 := tx.CarrierService.Create().
		SetLabel("FlexDeliveryService").
		SetInternalID("GLS_FLEX_DELIVERY_SERVICE").
		SetCarrierBrand(glsBrand).
		SaveX(ctx)
	tx.CarrierServiceGLS.Create().
		SetCarrierService(cs9).
		SetAPIKey("FlexDelivery").
		SetAPIValue(carrierservicegls.APIValueY).
		ExecX(ctx)

	//______________________

	tx.CarrierAdditionalServiceGLS.Create().
		SetLabel("Delivery point").
		SetInternalID(GLSDeliveryPointServiceInternalID).
		SetMandatory(true).
		SetCarrierServiceGLS(shopDeliveryService).
		AddCountriesConsignee(tx.Country.Query().Where(
			// Revisit this list...
			country.Alpha2("DK"),
			country.Alpha2("SE"),
			country.Alpha2("FR"),
			country.Alpha2("DE"),
			country.Alpha2("IT"),
			country.Alpha2("ES"),
			country.Alpha2("PL"),
			country.Alpha2("NL"),
			country.Alpha2("AT"),
			country.Alpha2("BE"),
			country.Alpha2("LU"),
		).AllX(ctx)...,
		).SaveX(ctx)
}
