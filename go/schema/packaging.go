package schema

import (
	"delivrio.io/go/schema/mixins"
	pulid_server_prefix "delivrio.io/go/schema/pulid-server-prefix"
	"delivrio.io/shared-utils/pulid"
	"entgo.io/contrib/entgql"
	"entgo.io/ent"
	"entgo.io/ent/schema"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
)

type Packaging struct {
	ent.Schema
}

func (Packaging) Annotations() []schema.Annotation {
	return []schema.Annotation{
		entgql.RelayConnection(),
		entgql.QueryField(),
		entgql.Mutations(
			entgql.MutationCreate(),
			entgql.MutationUpdate(),
		),
	}
}

func (Packaging) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("shipment_parcel", ShipmentParcel.Type).
			Ref("packaging"),
		edge.From("pallet", Pallet.Type).
			Ref("packaging"),
		edge.From("colli", Colli.Type).
			Ref("packaging"),
		edge.From("return_colli", ReturnColli.Type).
			Ref("packaging"),
		edge.To("packaging_df", PackagingDF.Type).
			Unique(),
		edge.To("packaging_usps", PackagingUSPS.Type).
			Unique(),
		edge.To("carrier_brand", CarrierBrand.Type).
			Unique(),
		edge.From("delivery_option", DeliveryOption.Type).
			Ref("default_packaging"),
	}
}

func (Packaging) Hooks() []ent.Hook {
	return []ent.Hook{
		//
	}
}

func (Packaging) Mixin() []ent.Mixin {
	return []ent.Mixin{
		BaseMixin{},
		TenantMixin{},
		pulid.MixinWithPrefix(pulid_server_prefix.TypeToPrefix("packaging")),
		mixins.ArchiveMixin{},
	}
}

func (Packaging) Fields() []ent.Field {
	return []ent.Field{
		field.String("name"),
		field.Int("height_cm"),
		field.Int("width_cm"),
		field.Int("length_cm"),
	}
}
