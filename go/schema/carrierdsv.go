package schema

import (
	pulid_server_prefix "delivrio.io/go/schema/pulid-server-prefix"
	"delivrio.io/shared-utils/pulid"
	"entgo.io/contrib/entgql"
	"entgo.io/ent"
	"entgo.io/ent/schema"
	"entgo.io/ent/schema/edge"
)

type CarrierDSV struct {
	ent.Schema
}

func (CarrierDSV) Annotations() []schema.Annotation {
	return []schema.Annotation{
		entgql.RelayConnection(),
		entgql.QueryField(),
		entgql.Mutations(
			entgql.MutationCreate(),
			entgql.MutationUpdate(),
		),
	}
}

func (CarrierDSV) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("carrier", Carrier.Type).
			Ref("carrier_dsv").
			Unique().
			Required(),
	}
}

func (CarrierDSV) Hooks() []ent.Hook {
	return []ent.Hook{
		//
	}
}

func (CarrierDSV) Mixin() []ent.Mixin {
	return []ent.Mixin{
		BaseMixin{},
		TenantMixin{},
		pulid.MixinWithPrefix(pulid_server_prefix.TypeToPrefix("carrier_dsv")),
	}
}

func (CarrierDSV) Fields() []ent.Field {
	return []ent.Field{
		//
	}
}
