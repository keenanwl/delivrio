package seed

import (
	"context"
	"delivrio.io/go/ent"
	"delivrio.io/go/ent/carrieradditionalserviceusps"
	"delivrio.io/go/ent/carrierbrand"
	"delivrio.io/go/ent/carrierserviceusps"
	"delivrio.io/go/ent/packaginguspsprocessingcategory"
)

func USPSServices(ctx context.Context) {
	tx := ent.TxFromContext(ctx)

	uspsBrand := tx.CarrierBrand.Query().
		Where(carrierbrand.InternalIDEQ(carrierbrand.InternalIDUSPS)).
		OnlyX(ctx)

	cs := tx.CarrierService.Create().
		SetLabel("Parcel Select").
		SetInternalID("parcel_select").
		SetCarrierBrand(uspsBrand).
		SaveX(ctx)
	tx.CarrierServiceUSPS.Create().
		SetCarrierService(cs).
		SetAPIKey(carrierserviceusps.APIKeyPARCEL_SELECT).
		ExecX(ctx)

	cs = tx.CarrierService.Create().
		SetLabel("Parcel Select Lightweight").
		SetInternalID("parcel_select_lightweight").
		SetCarrierBrand(uspsBrand).
		SaveX(ctx)
	tx.CarrierServiceUSPS.Create().
		SetCarrierService(cs).
		SetAPIKey(carrierserviceusps.APIKeyPARCEL_SELECT_LIGHTWEIGHT).
		ExecX(ctx)

	cs = tx.CarrierService.Create().
		SetLabel("USPS Connect Local").
		SetInternalID("usps_connect_local").
		SetCarrierBrand(uspsBrand).
		SaveX(ctx)
	tx.CarrierServiceUSPS.Create().
		SetCarrierService(cs).
		SetAPIKey(carrierserviceusps.APIKeyUSPS_CONNECT_LOCAL).
		ExecX(ctx)

	cs = tx.CarrierService.Create().
		SetLabel("USPS Connect Regional").
		SetInternalID("usps_connect_regional").
		SetCarrierBrand(uspsBrand).
		SaveX(ctx)
	tx.CarrierServiceUSPS.Create().
		SetCarrierService(cs).
		SetAPIKey(carrierserviceusps.APIKeyUSPS_CONNECT_REGIONAL).
		ExecX(ctx)

	cs = tx.CarrierService.Create().
		SetLabel("USPS Connect Mail").
		SetInternalID("usps_connect_mail").
		SetCarrierBrand(uspsBrand).
		SaveX(ctx)
	tx.CarrierServiceUSPS.Create().
		SetCarrierService(cs).
		SetAPIKey(carrierserviceusps.APIKeyUSPS_CONNECT_MAIL).
		ExecX(ctx)

	cs = tx.CarrierService.Create().
		SetLabel("USPS Ground Advantage").
		SetInternalID("usps_ground_advantage").
		SetCarrierBrand(uspsBrand).
		SaveX(ctx)
	tx.CarrierServiceUSPS.Create().
		SetCarrierService(cs).
		SetAPIKey(carrierserviceusps.APIKeyUSPS_GROUND_ADVANTAGE).
		ExecX(ctx)

	cs = tx.CarrierService.Create().
		SetLabel("Priority Mail Express").
		SetInternalID("priority_mail_express").
		SetCarrierBrand(uspsBrand).
		SaveX(ctx)
	tx.CarrierServiceUSPS.Create().
		SetCarrierService(cs).
		SetAPIKey(carrierserviceusps.APIKeyPRIORITY_MAIL_EXPRESS).
		ExecX(ctx)

	cs = tx.CarrierService.Create().
		SetLabel("Priority Mail").
		SetInternalID("priority_mail").
		SetCarrierBrand(uspsBrand).
		SaveX(ctx)
	tx.CarrierServiceUSPS.Create().
		SetCarrierService(cs).
		SetAPIKey(carrierserviceusps.APIKeyPRIORITY_MAIL).
		ExecX(ctx)

	cs = tx.CarrierService.Create().
		SetLabel("First-Class Package Service").
		SetInternalID("first-class_package_service").
		SetCarrierBrand(uspsBrand).
		SaveX(ctx)
	tx.CarrierServiceUSPS.Create().
		SetCarrierService(cs).
		SetAPIKey(carrierserviceusps.APIKeyFIRSTCLASSPACKAGESERVICE).
		ExecX(ctx)

	cs = tx.CarrierService.Create().
		SetLabel("Library Mail").
		SetInternalID("library_mail").
		SetCarrierBrand(uspsBrand).
		SaveX(ctx)
	tx.CarrierServiceUSPS.Create().
		SetCarrierService(cs).
		SetAPIKey(carrierserviceusps.APIKeyLIBRARY_MAIL).
		ExecX(ctx)

	cs = tx.CarrierService.Create().
		SetLabel("Media Mail").
		SetInternalID("media_mail").
		SetCarrierBrand(uspsBrand).
		SaveX(ctx)
	tx.CarrierServiceUSPS.Create().
		SetCarrierService(cs).
		SetAPIKey(carrierserviceusps.APIKeyMEDIA_MAIL).
		ExecX(ctx)

	cs = tx.CarrierService.Create().
		SetLabel("Bound Printed Matter").
		SetInternalID("bound_printed_matter").
		SetCarrierBrand(uspsBrand).
		SaveX(ctx)
	tx.CarrierServiceUSPS.Create().
		SetCarrierService(cs).
		SetAPIKey(carrierserviceusps.APIKeyBOUND_PRINTED_MATTER).
		ExecX(ctx)

	cs = tx.CarrierService.Create().
		SetLabel("Domestic Matter for the Blind").
		SetInternalID("domestic_matter_for_the_blind").
		SetCarrierBrand(uspsBrand).
		SaveX(ctx)
	tx.CarrierServiceUSPS.Create().
		SetCarrierService(cs).
		SetAPIKey(carrierserviceusps.APIKeyDOMESTIC_MATTER_FOR_THE_BLIND).
		ExecX(ctx)

	// RETURNS
	cs = tx.CarrierService.Create().
		SetLabel("First-Class Package Return Service").
		SetInternalID("first-class_package_return_service").
		SetCarrierBrand(uspsBrand).
		SetReturn(true).
		SaveX(ctx)
	tx.CarrierServiceUSPS.Create().
		SetCarrierService(cs).
		SetAPIKey(carrierserviceusps.APIKeyFIRSTCLASSPACKAGERETURNSERVICE).
		ExecX(ctx)

	cs = tx.CarrierService.Create().
		SetLabel("Ground Return Service").
		SetInternalID("ground_return_service").
		SetCarrierBrand(uspsBrand).
		SetReturn(true).
		SaveX(ctx)
	tx.CarrierServiceUSPS.Create().
		SetCarrierService(cs).
		SetAPIKey(carrierserviceusps.APIKeyGROUND_RETURN_SERVICE).
		ExecX(ctx)

	cs = tx.CarrierService.Create().
		SetLabel("Priority Mail Express Return Service").
		SetInternalID("priority_mail_express_return_service").
		SetCarrierBrand(uspsBrand).
		SetReturn(true).
		SaveX(ctx)
	tx.CarrierServiceUSPS.Create().
		SetCarrierService(cs).
		SetAPIKey(carrierserviceusps.APIKeyPRIORITY_MAIL_EXPRESS_RETURN_SERVICE).
		ExecX(ctx)

	cs = tx.CarrierService.Create().
		SetLabel("Priority Mail Return Service").
		SetInternalID("priority_mail_return_service").
		SetCarrierBrand(uspsBrand).
		SetReturn(true).
		SaveX(ctx)
	tx.CarrierServiceUSPS.Create().
		SetCarrierService(cs).
		SetAPIKey(carrierserviceusps.APIKeyPRIORITY_MAIL_RETURN_SERVICE).
		ExecX(ctx)

	cs = tx.CarrierService.Create().
		SetLabel("USPS Ground Advantage Return Service").
		SetInternalID("usps_ground_advantage_return_service").
		SetCarrierBrand(uspsBrand).
		SetReturn(true).
		SaveX(ctx)
	tx.CarrierServiceUSPS.Create().
		SetCarrierService(cs).
		SetAPIKey(carrierserviceusps.APIKeyUSPS_GROUND_ADVANTAGE_RETURN_SERVICE).
		ExecX(ctx)

}

func USPSRateIndicators(ctx context.Context) {
	tx := ent.TxFromContext(ctx)

	tx.PackagingUSPSRateIndicator.Create().
		SetName("Dimensional nonrectangular price").
		SetCode("DN").
		SaveX(ctx)

	tx.PackagingUSPSRateIndicator.Create().
		SetName("Dimensional rectangular price").
		SetCode("DR").
		SaveX(ctx)

	tx.PackagingUSPSRateIndicator.Create().
		SetName("Legal Flat Rate Envelope").
		SetCode("FA").
		SaveX(ctx)

	tx.PackagingUSPSRateIndicator.Create().
		SetName("Medium Flat Rate Box").
		SetCode("FB").
		SaveX(ctx)

	tx.PackagingUSPSRateIndicator.Create().
		SetName("Flat Rate Envelope").
		SetCode("FE").
		SaveX(ctx)

	tx.PackagingUSPSRateIndicator.Create().
		SetName("Padded Flat Rate Envelope").
		SetCode("FP").
		SaveX(ctx)

	tx.PackagingUSPSRateIndicator.Create().
		SetName("Small Flat Rate Box").
		SetCode("FS").
		SaveX(ctx)

	tx.PackagingUSPSRateIndicator.Create().
		SetName("Oversized price").
		SetCode("OS").
		SaveX(ctx)

	tx.PackagingUSPSRateIndicator.Create().
		SetName("Large Flat Rate Box").
		SetCode("PL").
		SaveX(ctx)

	tx.PackagingUSPSRateIndicator.Create().
		SetName("Single–piece price").
		SetCode("SP").
		SaveX(ctx)

	tx.PackagingUSPSProcessingCategory.Create().
		SetName("Machinable").
		SetProcessingCategory(packaginguspsprocessingcategory.ProcessingCategoryMACHINABLE).
		SaveX(ctx)

	tx.PackagingUSPSProcessingCategory.Create().
		SetName("Nonmachinable").
		SetProcessingCategory(packaginguspsprocessingcategory.ProcessingCategoryNON_MACHINABLE).
		SaveX(ctx)

	tx.PackagingUSPSProcessingCategory.Create().
		SetName("Letters").
		SetProcessingCategory(packaginguspsprocessingcategory.ProcessingCategoryLETTERS).
		SaveX(ctx)

	tx.PackagingUSPSProcessingCategory.Create().
		SetName("Flats").
		SetProcessingCategory(packaginguspsprocessingcategory.ProcessingCategoryFLATS).
		SaveX(ctx)

	tx.PackagingUSPSProcessingCategory.Create().
		SetName("Irregular").
		SetProcessingCategory(packaginguspsprocessingcategory.ProcessingCategoryIRREGULAR).
		SaveX(ctx)

}

func USPSAdditionalServices(ctx context.Context) {
	// Extra Service Code requested. If creating a HAZMAT label, users must include Extra Service Code 857.
	//Print and Deliver labels will always generate a 4x6 label with a receipt on the page.
	// \n * 365 - Global Direct Entry\n *
	// 415 - USPS Label Delivery Service\n *
	// 480 - Tracking Plus 6 Months\n *
	// 481 - Tracking Plus 1 Year\n *
	// 482 - Tracking Plus 3 Years\n *
	// 483 - Tracking Plus 5 Years\n *
	// 484 - Tracking Plus 7 Years\n *
	// 485 - Tracking Plus 10 Years\n *
	// 486 - Tracking Plus Signature 3 Years\n *
	// 487 - Tracking Plus Signature 5 Years\n *
	// 488 - Tracking Plus Signature 7 Years\n *
	// 489 - Tracking Plus Signature 10 Years\n *
	// 810 - Hazardous Materials - Air Eligible Ethanol\n *
	//811 - Hazardous Materials - Class 1 – Toy Propellant/Safety Fuse Package\n *
	//812 - Hazardous Materials - Class 3 - Flammable and Combustible Liquids\n *
	//813 - Hazardous Materials - Class 7 – Radioactive Materials\n *
	//814 - Hazardous Materials - Class 8 – Air Eligible Corrosive Materials\n *
	//815 - Hazardous Materials - Class 8 – Nonspillable Wet Batteries\n *
	//816 - Hazardous Materials - Class 9 - Lithium Battery Marked Ground Only  \n *
	//817 - Hazardous Materials - Class 9 - Lithium Battery Returns\n *
	//818 - Hazardous Materials - Class 9 - Marked Lithium Batteries\n *
	//819 - Hazardous Materials - Class 9 – Dry Ice\n *
	//820 - Hazardous Materials - Class 9 – Unmarked Lithium Batteries\n *
	//821 - Hazardous Materials - Class 9 – Magnetized Materials\n *
	//822 - Hazardous Materials - Division 4.1 – Mailable Flammable Solids and Safety Matches \n *
	//823 - Hazardous Materials - Division 5.1 – Oxidizers \n *
	//824 - Hazardous Materials - Division 5.2 – Organic Peroxides \n *
	//825 - Hazardous Materials - Division 6.1 – Toxic Materials \n *
	//826 - Hazardous Materials - Division 6.2 Biological Materials\n *
	//827 - Hazardous Materials - Excepted Quantity Provision \n *
	//828 - Hazardous Materials - Ground Only Hazardous Materials\n *
	//829 - Hazardous Materials - Air Eligible ID8000 Consumer Commodity\n *
	//830 - Hazardous Materials - Lighters \n *
	//831 - Hazardous Materials - Limited Quantity Ground \n *
	//832 - Hazardous Materials - Small Quantity Provision (Markings Required)\n *
	//857 - Hazardous Materials\n *
	//910 - Certified Mail\n *
	//911 - Certified Mail Restricted Delivery\n *
	//912 - Certified Mail Adult Signature Required\n *
	//913 - Certified Mail Adult Signature Restricted Delivery\n *
	//920 - USPS Tracking Electronic\n *
	//921 - Signature Confirmation\n *
	//922 - Adult Signature Required\n *
	//923 - Adult Signature Restricted Delivery\n *
	//924 - Signature Confirmation Restricted Delivery\n *
	//925 - Priority Mail Express Insurance \n *
	//930 - Insurance\n *
	//934 - Insurance Restricted Delivery\n *
	//955 - Return Receipt\n *
	//957 - Return Receipt Electronic\n *
	//981 - Signature Requested (PRIORITY_MAIL_EXPRESS only)\n *
	//986 - PO to Addressee (PRIORITY_MAIL_EXPRESS only)\n *
	//991 - Sunday Delivery
	tx := ent.TxFromContext(ctx)

	tx.CarrierAdditionalServiceUSPS.Create().
		SetLabel("Global Direct Entry").
		SetInternalID(carrieradditionalserviceusps.InternalIDGlobalDirectEntry).
		SetAPICode("365").
		ExecX(ctx)
	tx.CarrierAdditionalServiceUSPS.Create().
		SetLabel("USPS Label Delivery Service").
		SetInternalID(carrieradditionalserviceusps.InternalIDUSPSLabelDeliveryService).
		SetAPICode("415").
		SetCommonlyUsed(true).
		ExecX(ctx)
	tx.CarrierAdditionalServiceUSPS.Create().
		SetLabel("Tracking Plus 6 Months").
		SetInternalID(carrieradditionalserviceusps.InternalIDTrackingPlus6Months).
		SetAPICode("480").
		SetCommonlyUsed(true).
		ExecX(ctx)
	tx.CarrierAdditionalServiceUSPS.Create().
		SetLabel("Tracking Plus 1 Year").
		SetInternalID(carrieradditionalserviceusps.InternalIDTrackingPlus1Year).
		SetAPICode("481").
		ExecX(ctx)
	tx.CarrierAdditionalServiceUSPS.Create().
		SetLabel("Tracking Plus 3 Years").
		SetInternalID(carrieradditionalserviceusps.InternalIDTrackingPlus3Years).
		SetAPICode("482").
		ExecX(ctx)
	tx.CarrierAdditionalServiceUSPS.Create().
		SetLabel("Tracking Plus 5 Years").
		SetInternalID(carrieradditionalserviceusps.InternalIDTrackingPlus5Years).
		SetAPICode("483").
		ExecX(ctx)
	tx.CarrierAdditionalServiceUSPS.Create().
		SetLabel("Tracking Plus 7 Years").
		SetInternalID(carrieradditionalserviceusps.InternalIDTrackingPlus7Years).
		SetAPICode("484").
		ExecX(ctx)
	tx.CarrierAdditionalServiceUSPS.Create().
		SetLabel("Tracking Plus 10 Years").
		SetInternalID(carrieradditionalserviceusps.InternalIDTrackingPlus10Years).
		SetAPICode("485").
		ExecX(ctx)

	tx.CarrierAdditionalServiceUSPS.Create().
		SetLabel("Tracking Plus Signature 3 Years").
		SetInternalID(carrieradditionalserviceusps.InternalIDTrackingPlusSignature3Years).
		SetAPICode("486").
		SetCommonlyUsed(true).
		ExecX(ctx)
	tx.CarrierAdditionalServiceUSPS.Create().
		SetLabel("Tracking Plus Signature 5 Years").
		SetInternalID(carrieradditionalserviceusps.InternalIDTrackingPlusSignature5Years).
		SetAPICode("487").
		ExecX(ctx)
	tx.CarrierAdditionalServiceUSPS.Create().
		SetLabel("Tracking Plus Signature 7 Years").
		SetInternalID(carrieradditionalserviceusps.InternalIDTrackingPlusSignature7Years).
		SetAPICode("488").
		ExecX(ctx)
	tx.CarrierAdditionalServiceUSPS.Create().
		SetLabel("Tracking Plus Signature 10 Years").
		SetInternalID(carrieradditionalserviceusps.InternalIDTrackingPlusSignature10Years).
		SetAPICode("489").
		ExecX(ctx)

	tx.CarrierAdditionalServiceUSPS.Create().
		SetLabel("Hazardous Materials - Air Eligible Ethanol").
		SetInternalID(carrieradditionalserviceusps.InternalIDHazardousMaterialsAirEligibleEthanol).
		SetAPICode("810").
		ExecX(ctx)

	tx.CarrierAdditionalServiceUSPS.Create().
		SetLabel("Hazardous Materials - Class 1 – Toy Propellant").
		SetInternalID(carrieradditionalserviceusps.InternalIDHazardousMaterialsClass1ToyPropellant).
		SetAPICode("811").
		ExecX(ctx)

	tx.CarrierAdditionalServiceUSPS.Create().
		SetLabel("Hazardous Materials - Class 3 - Flammable and Combustible Liquids").
		SetInternalID(carrieradditionalserviceusps.InternalIDHazardousMaterialsClass3FlammableAndCombustibleLiquids).
		SetAPICode("812").
		ExecX(ctx)

	tx.CarrierAdditionalServiceUSPS.Create().
		SetLabel("Hazardous Materials - Class 7 – Radioactive Materials").
		SetInternalID(carrieradditionalserviceusps.InternalIDHazardousMaterialsClass7RadioactiveMaterials).
		SetAPICode("813").
		ExecX(ctx)
	tx.CarrierAdditionalServiceUSPS.Create().
		SetLabel("Hazardous Materials - Class 8 – Air Eligible Corrosive Materials").
		SetInternalID(carrieradditionalserviceusps.InternalIDHazardousMaterialsClass8AirEligibleCorrosiveMaterials).
		SetAPICode("814").
		ExecX(ctx)
	tx.CarrierAdditionalServiceUSPS.Create().
		SetLabel("Hazardous Materials - Class 8 – Nonspillable Wet Batteries").
		SetInternalID(carrieradditionalserviceusps.InternalIDHazardousMaterialsClass8NonspillableWetBatteries).
		SetAPICode("815").
		ExecX(ctx)
	tx.CarrierAdditionalServiceUSPS.Create().
		SetLabel("Hazardous Materials - Class 9 - Lithium Battery Marked Ground Only").
		SetInternalID(carrieradditionalserviceusps.InternalIDHazardousMaterialsClass9LithiumBatteryMarkedGroundOnly).
		SetAPICode("816").
		SetCommonlyUsed(true).
		ExecX(ctx)
	tx.CarrierAdditionalServiceUSPS.Create().
		SetLabel("Hazardous Materials - Class 9 - Lithium Battery Returns").
		SetInternalID(carrieradditionalserviceusps.InternalIDHazardousMaterialsClass9LithiumBatteryReturns).
		SetAPICode("817").
		SetCommonlyUsed(true).
		ExecX(ctx)
	tx.CarrierAdditionalServiceUSPS.Create().
		SetLabel("Hazardous Materials - Class 9 - Marked Lithium Battery").
		SetInternalID(carrieradditionalserviceusps.InternalIDHazardousMaterialsClass9MarkedLithiumBattery).
		SetAPICode("818").
		ExecX(ctx)
	tx.CarrierAdditionalServiceUSPS.Create().
		SetLabel("Hazardous Materials - Class 9 - Dry Ice").
		SetInternalID(carrieradditionalserviceusps.InternalIDHazardousMaterialsClass9DryIce).
		SetAPICode("819").
		ExecX(ctx)
	tx.CarrierAdditionalServiceUSPS.Create().
		SetLabel("Hazardous Materials - Class 9 - Unmarked Lithium Batteries").
		SetInternalID(carrieradditionalserviceusps.InternalIDHazardousMaterialsClass9UnmarkedLithiumBatteries).
		SetAPICode("820").
		ExecX(ctx)
	tx.CarrierAdditionalServiceUSPS.Create().
		SetLabel("Hazardous Materials - Class 9 - Magnetized Materials").
		SetInternalID(carrieradditionalserviceusps.InternalIDHazardousMaterialsClass9MagnetizedMaterials).
		SetAPICode("821").
		ExecX(ctx)

	tx.CarrierAdditionalServiceUSPS.Create().
		SetLabel("Hazardous Materials - Division 4.1 – Mailable Flammable Solids and Safety Matches").
		SetInternalID(carrieradditionalserviceusps.InternalIDHazardousMaterialsDivision41MailableFlammableSolidsAndSafetyMatches).
		SetAPICode("822").
		ExecX(ctx)
	tx.CarrierAdditionalServiceUSPS.Create().
		SetLabel("Hazardous Materials - Division 4.1 – Mailable Flammable Solids and Safety Matches").
		SetInternalID(carrieradditionalserviceusps.InternalIDHazardousMaterialsDivision41MailableFlammableSolidsAndSafetyMatches).
		SetAPICode("823").
		ExecX(ctx)
	tx.CarrierAdditionalServiceUSPS.Create().
		SetLabel("Hazardous Materials - Division 5.2 – Organic Peroxides").
		SetInternalID(carrieradditionalserviceusps.InternalIDHazardousMaterialsDivision52OrganicPeroxides).
		SetAPICode("824").
		ExecX(ctx)
	tx.CarrierAdditionalServiceUSPS.Create().
		SetLabel("Hazardous Materials - Division 6.1 – Toxic Materials").
		SetInternalID(carrieradditionalserviceusps.InternalIDHazardousMaterialsDivision61ToxicMaterials).
		SetAPICode("825").
		ExecX(ctx)
	tx.CarrierAdditionalServiceUSPS.Create().
		SetLabel("Hazardous Materials - Division 6.2 Biological Materials").
		SetInternalID(carrieradditionalserviceusps.InternalIDHazardousMaterialsDivision62BiologicalMaterials).
		SetAPICode("826").
		ExecX(ctx)
	tx.CarrierAdditionalServiceUSPS.Create().
		SetLabel("Hazardous Materials - Excepted Quantity Provision").
		SetInternalID(carrieradditionalserviceusps.InternalIDHazardousMaterialsExceptedQuantityProvision).
		SetAPICode("827").
		ExecX(ctx)
	tx.CarrierAdditionalServiceUSPS.Create().
		SetLabel("Hazardous Materials - Ground Only Hazardous Materials").
		SetInternalID(carrieradditionalserviceusps.InternalIDHazardousMaterialsGroundOnlyHazardousMaterials).
		SetAPICode("828").
		ExecX(ctx)
	tx.CarrierAdditionalServiceUSPS.Create().
		SetLabel("Hazardous Materials - Air Eligible ID8000 Consumer Commodity").
		SetInternalID(carrieradditionalserviceusps.InternalIDHazardousMaterialsAirEligibleId8000ConsumerCommodity).
		SetAPICode("829").
		ExecX(ctx)
	tx.CarrierAdditionalServiceUSPS.Create().
		SetLabel("Hazardous Materials - Lighters").
		SetInternalID(carrieradditionalserviceusps.InternalIDHazardousMaterialsLighters).
		SetAPICode("830").
		ExecX(ctx)
	tx.CarrierAdditionalServiceUSPS.Create().
		SetLabel("Hazardous Materials - Limited Quantity Ground").
		SetInternalID(carrieradditionalserviceusps.InternalIDHazardousMaterialsLimitedQuantityGround).
		SetAPICode("831").
		ExecX(ctx)
	tx.CarrierAdditionalServiceUSPS.Create().
		SetLabel("Hazardous Materials - Small Quantity Provision (Markings Required)").
		SetInternalID(carrieradditionalserviceusps.InternalIDHazardousMaterialsSmallQuantityProvisionMarkingsRequired).
		SetAPICode("832").
		ExecX(ctx)
	tx.CarrierAdditionalServiceUSPS.Create().
		SetLabel("Hazardous Materials").
		SetInternalID(carrieradditionalserviceusps.InternalIDHazardousMaterials).
		SetAPICode("857").
		ExecX(ctx)
	tx.CarrierAdditionalServiceUSPS.Create().
		SetLabel("Certified Mail").
		SetInternalID(carrieradditionalserviceusps.InternalIDCertifiedMail).
		SetAPICode("910").
		SetCommonlyUsed(true).
		ExecX(ctx)
	tx.CarrierAdditionalServiceUSPS.Create().
		SetLabel("Certified Mail Restricted Delivery").
		SetInternalID(carrieradditionalserviceusps.InternalIDCertifiedMailRestrictedDelivery).
		SetAPICode("911").
		ExecX(ctx)

	tx.CarrierAdditionalServiceUSPS.Create().
		SetLabel("Certified Mail Adult Signature Required").
		SetInternalID(carrieradditionalserviceusps.InternalIDCertifiedMailAdultSignatureRequired).
		SetAPICode("912").
		ExecX(ctx)
	tx.CarrierAdditionalServiceUSPS.Create().
		SetLabel("Certified Mail Adult Signature Restricted Delivery").
		SetInternalID(carrieradditionalserviceusps.InternalIDCertifiedMailAdultSignatureRestrictedDelivery).
		SetAPICode("913").
		ExecX(ctx)
	tx.CarrierAdditionalServiceUSPS.Create().
		SetLabel("USPS Tracking Electronic").
		SetInternalID(carrieradditionalserviceusps.InternalIDUSPSTrackingElectronic).
		SetAPICode("920").
		SetCommonlyUsed(true).
		ExecX(ctx)
	tx.CarrierAdditionalServiceUSPS.Create().
		SetLabel("Signature Confirmation").
		SetInternalID(carrieradditionalserviceusps.InternalIDSignatureConfirmation).
		SetAPICode("921").
		SetCommonlyUsed(true).
		ExecX(ctx)
	tx.CarrierAdditionalServiceUSPS.Create().
		SetLabel("Adult Signature Required").
		SetInternalID(carrieradditionalserviceusps.InternalIDAdultSignatureRequired).
		SetAPICode("922").
		ExecX(ctx)
	tx.CarrierAdditionalServiceUSPS.Create().
		SetLabel("Adult Signature Restricted Delivery").
		SetInternalID(carrieradditionalserviceusps.InternalIDAdultSignatureRestrictedDelivery).
		SetAPICode("923").
		ExecX(ctx)
	tx.CarrierAdditionalServiceUSPS.Create().
		SetLabel("Signature Confirmation Restricted Delivery").
		SetInternalID(carrieradditionalserviceusps.InternalIDSignatureConfirmationRestrictedDelivery).
		SetAPICode("924").
		ExecX(ctx)
	tx.CarrierAdditionalServiceUSPS.Create().
		SetLabel("Priority Mail Express Insurance").
		SetInternalID(carrieradditionalserviceusps.InternalIDPriorityMailExpressInsurance).
		SetAPICode("925").
		SetCommonlyUsed(true).
		ExecX(ctx)
	tx.CarrierAdditionalServiceUSPS.Create().
		SetLabel("Insurance").
		SetInternalID(carrieradditionalserviceusps.InternalIDInsurance).
		SetAPICode("930").
		SetCommonlyUsed(true).
		ExecX(ctx)
	tx.CarrierAdditionalServiceUSPS.Create().
		SetLabel("Insurance Restricted Delivery").
		SetInternalID(carrieradditionalserviceusps.InternalIDInsuranceRestrictedDelivery).
		SetAPICode("934").
		ExecX(ctx)
	tx.CarrierAdditionalServiceUSPS.Create().
		SetLabel("Return Receipt").
		SetInternalID(carrieradditionalserviceusps.InternalIDReturnReceipt).
		SetAPICode("955").
		SetCommonlyUsed(true).
		ExecX(ctx)
	tx.CarrierAdditionalServiceUSPS.Create().
		SetLabel("Return Receipt Electronic").
		SetInternalID(carrieradditionalserviceusps.InternalIDReturnReceiptElectronic).
		SetAPICode("957").
		ExecX(ctx)
	tx.CarrierAdditionalServiceUSPS.Create().
		SetLabel("Signature Requested (PRIORITY_MAIL_EXPRESS only)").
		SetInternalID(carrieradditionalserviceusps.InternalIDSignatureRequestedPriorityMailExpressOnly).
		SetAPICode("981").
		SetCommonlyUsed(true).
		ExecX(ctx)

	tx.CarrierAdditionalServiceUSPS.Create().
		SetLabel("PO to Addressee (PRIORITY_MAIL_EXPRESS only)").
		SetInternalID(carrieradditionalserviceusps.InternalIDPoToAddresseePriorityMailExpressOnly).
		SetAPICode("986").
		ExecX(ctx)
	tx.CarrierAdditionalServiceUSPS.Create().
		SetLabel("Sunday Delivery").
		SetInternalID(carrieradditionalserviceusps.InternalIDSundayDelivery).
		SetAPICode("991").
		SetCommonlyUsed(true).
		ExecX(ctx)

}
