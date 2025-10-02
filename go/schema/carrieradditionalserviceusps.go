package schema

import (
	pulid_server_prefix "delivrio.io/go/schema/pulid-server-prefix"
	"delivrio.io/shared-utils/pulid"
	"entgo.io/contrib/entgql"
	"entgo.io/ent"
	"entgo.io/ent/schema"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
)

type CarrierAdditionalServiceUSPS struct {
	ent.Schema
}

func (CarrierAdditionalServiceUSPS) Annotations() []schema.Annotation {
	return []schema.Annotation{
		entgql.RelayConnection(),
		entgql.QueryField(),
	}
}

func (CarrierAdditionalServiceUSPS) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("carrier_service_usps", CarrierServiceUSPS.Type).
			Ref("carrier_additional_service_usps").
			Unique(),
		edge.From("delivery_option_usps", DeliveryOptionUSPS.Type).
			Ref("carrier_additional_service_usps"),
	}
}

func (CarrierAdditionalServiceUSPS) Hooks() []ent.Hook {
	return []ent.Hook{
		//
	}
}

func (CarrierAdditionalServiceUSPS) Mixin() []ent.Mixin {
	return []ent.Mixin{
		BaseMixin{},
		pulid.MixinWithPrefix(pulid_server_prefix.TypeToPrefix("carrier_additional_service_usps")),
	}
}

func (CarrierAdditionalServiceUSPS) Fields() []ent.Field {
	return []ent.Field{
		field.String("label"),
		field.Bool("commonly_used").
			Default(false).
			Comment("For filtering away rarely used options in the UI"),
		field.Enum("internal_id").
			Values("global_direct_entry",
				"usps_label_delivery_service",
				"tracking_plus_6_months",
				"tracking_plus_1_year",
				"tracking_plus_3_years",
				"tracking_plus_5_years",
				"tracking_plus_7_years",
				"tracking_plus_10_years",
				"tracking_plus_signature_3_years",
				"tracking_plus_signature_5_years",
				"tracking_plus_signature_7_years",
				"tracking_plus_signature_10_years",
				"hazardous_materials_air_eligible_ethanol",
				"hazardous_materials_class_1_toy_propellant",
				"hazardous_materials_class_3_flammable_and_combustible_liquids",
				"hazardous_materials_class_7_radioactive_materials",
				"hazardous_materials_class_8_air_eligible_corrosive_materials",
				"hazardous_materials_class_8_nonspillable_wet_batteries",
				"hazardous_materials_class_9_lithium_battery_marked_ground_only",
				"hazardous_materials_class_9_lithium_battery_returns",
				"hazardous_materials_class_9_marked_lithium_battery",
				"hazardous_materials_class_9_dry_ice",
				"hazardous_materials_class_9_unmarked_lithium_batteries",
				"hazardous_materials_class_9_magnetized_materials",
				"hazardous_materials_division_4_1_mailable_flammable_solids_and_safety_matches",
				"hazardous_materials_division_5_2_organic_peroxides",
				"hazardous_materials_division_6_1_toxic_materials",
				"hazardous_materials_division_6_2_biological_materials",
				"hazardous_materials_excepted_quantity_provision",
				"hazardous_materials_ground_only_hazardous_materials",
				"hazardous_materials_air_eligible_id8000_consumer_commodity",
				"hazardous_materials_lighters",
				"hazardous_materials_limited_quantity_ground",
				"hazardous_materials_small_quantity_provision_markings_required",
				"hazardous_materials",
				"certified_mail",
				"certified_mail_restricted_delivery",
				"certified_mail_adult_signature_required",
				"certified_mail_adult_signature_restricted_delivery",
				"usps_tracking_electronic",
				"signature_confirmation",
				"adult_signature_required",
				"adult_signature_restricted_delivery",
				"signature_confirmation_restricted_delivery",
				"priority_mail_express_insurance",
				"insurance",
				"insurance_restricted_delivery",
				"return_receipt",
				"return_receipt_electronic",
				"signature_requested_priority_mail_express_only",
				"po_to_addressee_priority_mail_express_only",
				"sunday_delivery",
			).
			Comment(""),
		field.String("api_code").
			Comment("ServiceID to be included in XML payload"),
	}
}
