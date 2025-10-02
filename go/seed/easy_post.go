package seed

import (
	"context"
	"delivrio.io/go/ent"
	"delivrio.io/go/ent/carrierbrand"
	"delivrio.io/go/ent/carrierserviceeasypost"
)

func EasyPostServices(ctx context.Context) {
	tx := ent.TxFromContext(ctx)

	cb := tx.CarrierBrand.Query().
		Where(carrierbrand.InternalIDEQ(carrierbrand.InternalIDEasyPost)).
		OnlyX(ctx)

	csHome := tx.CarrierService.Create().
		SetCarrierBrand(cb).
		SetLabel("USPS - First").
		SetInternalID(carrierserviceeasypost.APIKeyFirst.String()).
		SaveX(ctx)
	tx.CarrierServiceEasyPost.Create().
		SetCarrierService(csHome).
		SetAPIKey(carrierserviceeasypost.APIKeyFirst).
		SaveX(ctx)

	csHome = tx.CarrierService.Create().
		SetCarrierBrand(cb).
		SetLabel("USPS - Priority").
		SetInternalID(carrierserviceeasypost.APIKeyPriority.String()).
		SaveX(ctx)
	tx.CarrierServiceEasyPost.Create().
		SetCarrierService(csHome).
		SetAPIKey(carrierserviceeasypost.APIKeyPriority).
		SaveX(ctx)

	csHome = tx.CarrierService.Create().
		SetCarrierBrand(cb).
		SetLabel("USPS - Express").
		SetInternalID(carrierserviceeasypost.APIKeyExpress.String()).
		SaveX(ctx)
	tx.CarrierServiceEasyPost.Create().
		SetCarrierService(csHome).
		SetAPIKey(carrierserviceeasypost.APIKeyExpress).
		SaveX(ctx)

	csHome = tx.CarrierService.Create().
		SetCarrierBrand(cb).
		SetLabel("USPS - GroundAdvantage").
		SetInternalID(carrierserviceeasypost.APIKeyGroundAdvantage.String()).
		SaveX(ctx)
	tx.CarrierServiceEasyPost.Create().
		SetCarrierService(csHome).
		SetAPIKey(carrierserviceeasypost.APIKeyGroundAdvantage).
		SaveX(ctx)

	csHome = tx.CarrierService.Create().
		SetCarrierBrand(cb).
		SetLabel("USPS - LibraryMail").
		SetInternalID(carrierserviceeasypost.APIKeyLibraryMail.String()).
		SaveX(ctx)
	tx.CarrierServiceEasyPost.Create().
		SetCarrierService(csHome).
		SetAPIKey(carrierserviceeasypost.APIKeyLibraryMail).
		SaveX(ctx)

	csHome = tx.CarrierService.Create().
		SetCarrierBrand(cb).
		SetLabel("USPS - MediaMail").
		SetInternalID(carrierserviceeasypost.APIKeyMediaMail.String()).
		SaveX(ctx)
	tx.CarrierServiceEasyPost.Create().
		SetCarrierService(csHome).
		SetAPIKey(carrierserviceeasypost.APIKeyMediaMail).
		SaveX(ctx)

	csHome = tx.CarrierService.Create().
		SetCarrierBrand(cb).
		SetLabel("USPS - FirstClassMailInternational").
		SetInternalID(carrierserviceeasypost.APIKeyFirstClassMailInternational.String()).
		SaveX(ctx)
	tx.CarrierServiceEasyPost.Create().
		SetCarrierService(csHome).
		SetAPIKey(carrierserviceeasypost.APIKeyFirstClassMailInternational).
		SaveX(ctx)

	csHome = tx.CarrierService.Create().
		SetCarrierBrand(cb).
		SetLabel("USPS - FirstClassPackageInternationalService").
		SetInternalID(carrierserviceeasypost.APIKeyFirstClassPackageInternationalService.String()).
		SaveX(ctx)
	tx.CarrierServiceEasyPost.Create().
		SetCarrierService(csHome).
		SetAPIKey(carrierserviceeasypost.APIKeyFirstClassPackageInternationalService).
		SaveX(ctx)

	csHome = tx.CarrierService.Create().
		SetCarrierBrand(cb).
		SetLabel("USPS - PriorityMailInternational").
		SetInternalID(carrierserviceeasypost.APIKeyPriorityMailInternational.String()).
		SaveX(ctx)
	tx.CarrierServiceEasyPost.Create().
		SetCarrierService(csHome).
		SetAPIKey(carrierserviceeasypost.APIKeyPriorityMailInternational).
		SaveX(ctx)

	csHome = tx.CarrierService.Create().
		SetCarrierBrand(cb).
		SetLabel("USPS - ExpressMailInternational").
		SetInternalID(carrierserviceeasypost.APIKeyExpressMailInternational.String()).
		SaveX(ctx)
	tx.CarrierServiceEasyPost.Create().
		SetCarrierService(csHome).
		SetAPIKey(carrierserviceeasypost.APIKeyExpressMailInternational).
		SaveX(ctx)

	tx.CarrierAdditionalServiceEasyPost.Create().
		AddCarrierServiceEasyPost(tx.CarrierServiceEasyPost.Query().AllX(ctx)...).
		SetAPIKey("delivery_confirmation").
		SetAPIValue("SIGNATURE").
		SetLabel("Delivery Confirmation - Signature").
		SaveX(ctx)
	tx.CarrierAdditionalServiceEasyPost.Create().
		AddCarrierServiceEasyPost(tx.CarrierServiceEasyPost.Query().AllX(ctx)...).
		SetAPIKey("delivery_confirmation").
		SetAPIValue("ADULT_SIGNATURE").
		SetLabel("Delivery Confirmation - Adult Signature").
		SaveX(ctx)
	tx.CarrierAdditionalServiceEasyPost.Create().
		AddCarrierServiceEasyPost(tx.CarrierServiceEasyPost.Query().AllX(ctx)...).
		SetAPIKey("delivery_confirmation").
		SetAPIValue("SIGNATURE_RESTRICTED").
		SetLabel("Delivery Confirmation - Signature Restricted").
		SaveX(ctx)
	tx.CarrierAdditionalServiceEasyPost.Create().
		AddCarrierServiceEasyPost(tx.CarrierServiceEasyPost.Query().AllX(ctx)...).
		SetAPIKey("delivery_confirmation").
		SetAPIValue("ADULT_SIGNATURE_RESTRICTED").
		SetLabel("Delivery Confirmation - Adult Signature Restricted").
		ExecX(ctx)

	tx.CarrierAdditionalServiceEasyPost.Create().
		AddCarrierServiceEasyPost(tx.CarrierServiceEasyPost.Query().AllX(ctx)...).
		SetAPIKey("address_validation_level").
		SetAPIValue("0").
		SetLabel("Disable Address Validation").
		ExecX(ctx)

	tx.CarrierAdditionalServiceEasyPost.Create().
		AddCarrierServiceEasyPost(tx.CarrierServiceEasyPost.Query().AllX(ctx)...).
		SetAPIKey("endorsement").
		SetAPIValue("ADDRESS_SERVICE_REQUESTED").
		SetLabel("Address Service Requested").
		SaveX(ctx)
	tx.CarrierAdditionalServiceEasyPost.Create().
		AddCarrierServiceEasyPost(tx.CarrierServiceEasyPost.Query().AllX(ctx)...).
		SetAPIKey("endorsement").
		SetAPIValue("FORWARDING_SERVICE_REQUESTED").
		SetLabel("Forwarding Service Requested").
		SaveX(ctx)
	tx.CarrierAdditionalServiceEasyPost.Create().
		AddCarrierServiceEasyPost(tx.CarrierServiceEasyPost.Query().AllX(ctx)...).
		SetAPIKey("endorsement").
		SetAPIValue("CHANGE_SERVICE_REQUESTED").
		SetLabel("Change Service Requested").
		SaveX(ctx)
	tx.CarrierAdditionalServiceEasyPost.Create().
		AddCarrierServiceEasyPost(tx.CarrierServiceEasyPost.Query().AllX(ctx)...).
		SetAPIKey("endorsement").
		SetAPIValue("RETURN_SERVICE_REQUESTED").
		SetLabel("Return Service Requested").
		SaveX(ctx)
	tx.CarrierAdditionalServiceEasyPost.Create().
		AddCarrierServiceEasyPost(tx.CarrierServiceEasyPost.Query().AllX(ctx)...).
		SetAPIKey("endorsement").
		SetAPIValue("LEAVE_IF_NO_RESPONSE").
		SetLabel("Leave if no Response").
		SaveX(ctx)

}
