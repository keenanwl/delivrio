package pulid_server_prefix

import (
	"context"
	"delivrio.io/go/ent/colli"
	"delivrio.io/go/ent/hypothesistestdeliveryoptionlookup"
	"delivrio.io/go/ent/hypothesistestdeliveryoptionrequest"
	"fmt"

	"delivrio.io/go/ent/accessright"
	"delivrio.io/go/ent/address"
	"delivrio.io/go/ent/addressglobal"
	"delivrio.io/go/ent/apitoken"
	"delivrio.io/go/ent/carrier"
	"delivrio.io/go/ent/carrieradditionalservicepostnord"
	"delivrio.io/go/ent/carrierbrand"
	"delivrio.io/go/ent/carriergls"
	"delivrio.io/go/ent/carrierservicepostnord"
	"delivrio.io/go/ent/changehistory"
	"delivrio.io/go/ent/connection"
	"delivrio.io/go/ent/connectionbrand"
	"delivrio.io/go/ent/connectionshopify"
	"delivrio.io/go/ent/connectoptioncarrier"
	"delivrio.io/go/ent/connectoptionplatform"
	"delivrio.io/go/ent/contact"
	"delivrio.io/go/ent/country"
	"delivrio.io/go/ent/currency"
	"delivrio.io/go/ent/deliveryoptiongls"
	"delivrio.io/go/ent/deliveryrule"
	"delivrio.io/go/ent/deliveryruleconstraint"
	"delivrio.io/go/ent/deliveryruleconstraintgroup"
	"delivrio.io/go/ent/language"
	"delivrio.io/go/ent/location"
	"delivrio.io/go/ent/locationtag"
	"delivrio.io/go/ent/order"
	"delivrio.io/go/ent/orderhistory"
	"delivrio.io/go/ent/orderline"
	"delivrio.io/go/ent/ordersender"
	"delivrio.io/go/ent/otkrequests"
	"delivrio.io/go/ent/parcelshop"
	"delivrio.io/go/ent/parcelshoppostnord"
	"delivrio.io/go/ent/plan"
	"delivrio.io/go/ent/planhistory"
	"delivrio.io/go/ent/printer"
	"delivrio.io/go/ent/printjob"
	"delivrio.io/go/ent/product"
	"delivrio.io/go/ent/producttag"
	"delivrio.io/go/ent/productvariant"
	"delivrio.io/go/ent/seatgroup"
	"delivrio.io/go/ent/seatgroupaccessright"
	"delivrio.io/go/ent/shipment"
	"delivrio.io/go/ent/shipmenthistory"
	"delivrio.io/go/ent/shipmentparcel"
	"delivrio.io/go/ent/shipmentpostnord"
	"delivrio.io/go/ent/signupoptions"
	"delivrio.io/go/ent/systemevents"
	"delivrio.io/go/ent/tenant"
	"delivrio.io/go/ent/user"
	"delivrio.io/go/ent/userseat"
	"delivrio.io/go/ent/workstation"
	"delivrio.io/shared-utils/pulid"
)

type LabelTable struct {
	L string
	T string
}

// prefixMap maps PULID prefixes to table names.
var prefixMap = map[pulid.ID]LabelTable{
	"AD": {T: address.Table, L: address.Label},
	"AT": {T: apitoken.Table, L: apitoken.Label},
	"AS": {T: carrieradditionalservicepostnord.Table, L: carrieradditionalservicepostnord.Label},
	"AG": {T: addressglobal.Table, L: addressglobal.Label},
	"C1": {T: carrierservicepostnord.Table, L: carrierservicepostnord.Label},
	"AR": {T: accessright.Table, L: accessright.Label},
	"CA": {T: carrier.Table, L: carrier.Label},
	"CG": {T: carriergls.Table, L: carriergls.Label},
	"CD": {T: carrierbrand.Table, L: carrierbrand.Label},
	"CE": {T: "carrier_service", L: "carrier_service"},
	"CB": {T: connectionbrand.Table, L: connectionbrand.Label},
	"CC": {T: connectoptioncarrier.Table, L: connectoptioncarrier.Label},
	"CL": {T: colli.Table, L: colli.Label},
	"CO": {T: country.Table, L: country.Label},
	"CH": {T: changehistory.Table, L: changehistory.Label},
	"CT": {T: connection.Table, L: connection.Label},
	"CS": {T: connectionshopify.Table, L: connectionshopify.Label},
	"CN": {T: contact.Table, L: contact.Label},
	"CP": {T: connectoptionplatform.Table, L: connectoptionplatform.Label},
	"CK": {T: "carrier_service_gls", L: "carrier_service_gls"},
	"C2": {T: "carrier_service_usps", L: "carrier_service_usps"},
	"CQ": {T: "carrier_additional_service_gls", L: "carrier_additional_service_gls"},
	"C3": {T: "carrier_additional_service_usps", L: "carrier_additional_service_usps"},
	"PN": {T: "carrier_post_nord", L: "carrier_post_nord"},
	"CM": {T: "carrier_usps", L: "carrier_usps"},
	"CU": {T: currency.Table, L: currency.Label},
	"DG": {T: deliveryoptiongls.Table, L: deliveryoptiongls.Label},
	"DU": {T: "delivery_option_usps", L: "delivery_option_usps"},
	"DR": {T: deliveryrule.Table, L: deliveryrule.Label},
	"DC": {T: deliveryruleconstraint.Table, L: deliveryruleconstraint.Label},
	"DT": {T: "delivery_option", L: "delivery_option"},
	"DN": {T: "delivery_option_post_nord", L: "delivery_option_post_nord"},
	"DP": {T: "delivery_option_price", L: "delivery_option_price"},
	"DL": {T: "delivery_rule_price", L: "delivery_rule_price"},
	"DM": {T: "delivery_option_price_margin", L: "delivery_option_price_margin"},
	"ET": {T: "email_template", L: "email_template"},
	"R1": {T: "return_portal", L: "return_portal"},
	"R2": {T: "return_portal_claim", L: "return_portal_claim"},
	"RC": {T: "return_colli", L: "return_colli"},
	"RL": {T: "return_order_line", L: "return_order_line"},
	"RG": {T: deliveryruleconstraintgroup.Table, L: deliveryruleconstraintgroup.Label},
	"LA": {T: language.Table, L: language.Label},
	"LO": {T: location.Table, L: location.Label},
	"LT": {T: locationtag.Table, L: locationtag.Label},
	"OR": {T: order.Table, L: order.Label},
	"OH": {T: orderhistory.Table, L: orderhistory.Label},
	"OI": {T: orderline.Table, L: orderline.Label},
	"O3": {T: ordersender.Table, L: ordersender.Label},
	"OT": {T: otkrequests.Table, L: otkrequests.Label},
	"SA": {T: seatgroupaccessright.Table, L: seatgroupaccessright.Label},
	"SE": {T: systemevents.Table, L: systemevents.Label},
	"SO": {T: signupoptions.Table, L: signupoptions.Label},
	"SG": {T: seatgroup.Table, L: seatgroup.Label},
	"SH": {T: shipment.Table, L: shipment.Label},
	"SP": {T: shipmentpostnord.Table, L: shipmentpostnord.Label},
	"S1": {T: shipmentparcel.Table, L: shipmentparcel.Label},
	"SY": {T: shipmenthistory.Table, L: shipmenthistory.Label},
	"SW": {T: "shipment_gls", L: "shipment_gls"},
	"TE": {T: tenant.Table, L: tenant.Label},
	"PL": {T: plan.Table, L: plan.Label},
	"PI": {T: "product_image", L: "product_image"},
	"PD": {T: product.Table, L: product.Label},
	"PT": {T: producttag.Table, L: producttag.Label},
	"PR": {T: printer.Table, L: printer.Label},
	"PV": {T: productvariant.Table, L: productvariant.Label},
	"HT": {T: "hypothesis_test", L: "hypothesis_test"},
	"NO": {T: "notification", L: "notification"},
	"HD": {T: "hypothesis_test_delivery_option", L: "hypothesis_test_delivery_option"},
	"HU": {T: hypothesistestdeliveryoptionlookup.Table, L: hypothesistestdeliveryoptionlookup.Label},
	"HR": {T: hypothesistestdeliveryoptionrequest.Table, L: "hypothesis_test_delivery_option_request"},
	"PH": {T: planhistory.Table, L: planhistory.Label},
	"PG": {T: "parcel_shop_gls", L: "parcel_shop_gls"},
	"PJ": {T: printjob.Table, L: printjob.Label},
	"RH": {T: "return_colli_history", L: "return_colli_history"},
	"P1": {T: parcelshop.Table, L: parcelshop.Label},
	"SU": {T: "shipment_usps", L: "shipment_usps"},
	"P2": {T: parcelshoppostnord.Table, L: parcelshoppostnord.Label},
	"PK": {T: "packaging", L: "packaging"},
	"P3": {T: "packaging_usps", L: "packaging_usps"},
	"P4": {T: "packaging_usps_rate_indicator", L: "packaging_usps_rate_indicator"},
	"P5": {T: "packaging_usps_processing_category", L: "packaging_usps_processing_category"},
	"US": {T: user.Table, L: user.Label},
	"U1": {T: userseat.Table, L: userseat.Label},
	"WR": {T: "workspace_recent_scan", L: "workspace_recent_scan"},
	"WS": {T: workstation.Table, L: workstation.Label},
	"PB": {T: "parcel_shop_bring", L: "parcel_shop_bring"},
	"DB": {T: "delivery_option_bring", L: "delivery_option_bring"},
	"SB": {T: "shipment_bring", L: "shipment_bring"},
	"BA": {T: "carrier_additional_service_bring", L: "carrier_additional_service_bring"},
	"BS": {T: "carrier_service_bring", L: "carrier_service_bring"},
	"BC": {T: "carrier_bring", L: "carrier_bring"},

	"D1": {T: "parcel_shop_dao", L: "parcel_shop_dao"},
	"D2": {T: "delivery_option_dao", L: "delivery_option_dao"},
	"D3": {T: "shipment_dao", L: "shipment_dao"},
	"D4": {T: "carrier_additional_service_dao", L: "carrier_additional_service_dao"},
	"D5": {T: "carrier_service_dao", L: "carrier_service_dao"},
	"D6": {T: "carrier_dao", L: "carrier_dao"},

	"V2": {T: "delivery_option_dsv", L: "delivery_option_dsv"},
	"V3": {T: "shipment_dsv", L: "shipment_dsv"},
	"V4": {T: "carrier_additional_service_dsv", L: "carrier_additional_service_dsv"},
	"V5": {T: "carrier_service_dsv", L: "carrier_service_dsv"},
	"V6": {T: "carrier_dsv", L: "carrier_dsv"},

	"F1": {T: "delivery_option_df", L: "delivery_option_df"},
	"F2": {T: "shipment_df", L: "shipment_df"},
	"F3": {T: "carrier_additional_service_df", L: "carrier_additional_service_df"},
	"F4": {T: "carrier_service_df", L: "carrier_service_df"},
	"F5": {T: "carrier_df", L: "carrier_df"},
	"F6": {T: "packaging_df", L: "packaging_df"},

	"E1": {T: "delivery_option_easy_post", L: "delivery_option_easy_post"},
	"E2": {T: "shipment_easy_post", L: "shipment_easy_post"},
	"E3": {T: "carrier_additional_service_easy_post", L: "carrier_additional_service_easy_post"},
	"E4": {T: "carrier_service_easy_post", L: "carrier_service_easy_post"},
	"E5": {T: "carrier_easy_post", L: "carrier_easy_post"},
	"E6": {T: "packaging_easy_post", L: "packaging_easy_post"},

	"PA": {T: "pallet", L: "pallet"},
	"1C": {T: "consolidation", L: "consolidation"},

	"P6": {T: "shipment_pallet", L: "shipment_pallet"},
	"D7": {T: "document", L: "document"},

	"DF": {T: "document_file", L: "document_file"},

	"HS": {T: "country_harmonized_code", L: "country_harmonized_code"},
	"II": {T: "inventory_item", L: "inventory_item"},
	"BH": {T: "business_hours_period", L: "business_hours_period"},

	"LC": {T: "connection_lookup", L: "connection_lookup"},
}

// IDToType maps a pulid.ID to the underlying table.
func IDToType(ctx context.Context, id pulid.ID) (string, error) {
	if len(id) < 2 {
		return "", fmt.Errorf("IDToType: id too short")
	}
	prefix := id[:2]
	if val, ok := prefixMap[prefix]; ok {
		return val.T, nil
	}

	return "", fmt.Errorf("IDToType: could not map prefix '%s' to a type", prefix)
}

func TypeToPrefix(label string) string {

	for p, t := range prefixMap {
		if t.L == label {
			return string(p)
		}
	}

	panic(fmt.Sprintf("label not found: %s", label))
}
