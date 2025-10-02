package schema

import (
	"entgo.io/contrib/entgql"
	"entgo.io/ent"
	"entgo.io/ent/schema/field"
	"time"
)

// RecentScan holds the schema definition for the RecentScan entity.
type RecentScan struct {
	ent.Schema
}

// Fields of the RecentScan.
func (RecentScan) Fields() []ent.Field {
	return []ent.Field{
		field.String("scan_value"),
		field.String("response"),
		field.Enum("scan_type").
			Default("label_request").
			Values("label_request"),
		field.Time("created_at").
			Default(time.Now).
			Annotations(
				entgql.Skip(entgql.SkipMutationCreateInput | entgql.SkipMutationUpdateInput),
			).
			Immutable(),
	}
}

// Edges of the RecentScan.
func (RecentScan) Edges() []ent.Edge {
	return nil
}
