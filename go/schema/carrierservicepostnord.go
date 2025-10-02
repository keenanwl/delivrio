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

type CarrierServicePostNord struct {
	ent.Schema
}

func (CarrierServicePostNord) Annotations() []schema.Annotation {
	return []schema.Annotation{
		entgql.QueryField(),
	}
}

func (CarrierServicePostNord) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("carrier_service", CarrierService.Type).
			Ref("carrier_service_post_nord").
			Unique().
			Required(),
		edge.To("carrier_add_serv_post_nord", CarrierAdditionalServicePostNord.Type),
	}
}

func (CarrierServicePostNord) Hooks() []ent.Hook {
	return []ent.Hook{
		//
	}
}

func (CarrierServicePostNord) Mixin() []ent.Mixin {
	return []ent.Mixin{
		BaseMixin{},
		pulid.MixinWithPrefix(pulid_server_prefix.TypeToPrefix("carrier_service_post_nord")),
	}
}

func (CarrierServicePostNord) Fields() []ent.Field {
	return []ent.Field{
		field.String("label"),
		field.String("internal_id").Unique(),
		field.String("api_code"),
	}
}
