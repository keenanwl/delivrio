package schema

import (
	"delivrio.io/go/schema/pulid-server-prefix"
	"delivrio.io/shared-utils/pulid"
	"entgo.io/contrib/entgql"
	"entgo.io/ent"
	"entgo.io/ent/schema"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
)

// Country holds the schema definition for the Country entity.
type Country struct {
	ent.Schema
}

func (Country) Annotations() []schema.Annotation {
	return []schema.Annotation{
		entgql.RelayConnection(),
		entgql.QueryField(),
	}
}

// Fields of the Country.
func (Country) Fields() []ent.Field {
	return []ent.Field{
		field.String("label").Unique(),
		field.String("alpha_2").Unique().MaxLen(2),
		field.String("alpha_3").Unique().MaxLen(3),
		field.String("code").Unique(),
		field.Enum("region").
			Values("Asia", "Europe", "Oceania", "Americas", "Africa"),
	}
}

// Edges of the Country.
func (Country) Edges() []ent.Edge {
	return []ent.Edge{
		edge.To("delivery_rule", DeliveryRule.Type),
		edge.From("address", Address.Type).
			Ref("country"),
		edge.From("address_global", AddressGlobal.Type).
			Ref("country"),
		edge.From("carrier_additional_service_post_nord_consignee", CarrierAdditionalServicePostNord.Type).
			Ref("countries_consignee"),
		edge.From("carrier_additional_service_post_nord_consignor", CarrierAdditionalServicePostNord.Type).
			Ref("countries_consignor"),
		edge.From("carrier_additional_service_gls_consignee", CarrierAdditionalServiceGLS.Type).
			Ref("countries_consignee"),
		edge.From("carrier_additional_service_gls_consignor", CarrierAdditionalServiceGLS.Type).
			Ref("countries_consignor"),
		edge.From("country_harmonized_code", CountryHarmonizedCode.Type).
			Ref("country"),
		edge.From("inventory_item", InventoryItem.Type).
			Ref("country_of_origin"),
	}
}

func (Country) Mixin() []ent.Mixin {
	return []ent.Mixin{
		BaseMixin{},
		pulid.MixinWithPrefix(pulid_server_prefix.TypeToPrefix("country")),
	}
}
