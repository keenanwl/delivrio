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

type BusinessHoursPeriod struct {
	ent.Schema
}

func (BusinessHoursPeriod) Annotations() []schema.Annotation {
	return []schema.Annotation{
		entgql.RelayConnection(),
		entgql.QueryField(),
	}
}

func (BusinessHoursPeriod) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("parcel_shop", ParcelShop.Type).
			Required().
			Unique().
			Ref("business_hours_period"),
	}
}

func (BusinessHoursPeriod) Hooks() []ent.Hook {
	return []ent.Hook{
		//
	}
}

func (BusinessHoursPeriod) Mixin() []ent.Mixin {
	return []ent.Mixin{
		BaseMixin{},
		pulid.MixinWithPrefix(pulid_server_prefix.TypeToPrefix("business_hours_period")),
	}
}

func (BusinessHoursPeriod) Fields() []ent.Field {
	return []ent.Field{
		field.Enum("day_of_week").
			Values("MONDAY", "TUESDAY", "WEDNESDAY", "THURSDAY", "FRIDAY", "SATURDAY", "SUNDAY"),
		field.Time("opening"),
		field.Time("closing"),
	}
}
