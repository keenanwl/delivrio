package schema

import (
	"time"

	"delivrio.io/print-client/ent/schema/pulid_prefix"
	"delivrio.io/shared-utils/pulid"
	"entgo.io/ent"
	"entgo.io/ent/schema/field"
)

// LogError holds the schema definition for the LogError entity.
type LogError struct {
	ent.Schema
}

// Fields of the LogError.
func (LogError) Fields() []ent.Field {
	return []ent.Field{
		field.String("error"),
		field.Time("created_at").
			Default(time.Now).
			Immutable(),
	}
}

// Edges of the LogError.
func (LogError) Edges() []ent.Edge {
	return nil
}

func (LogError) Mixin() []ent.Mixin {
	return []ent.Mixin{
		pulid.MixinWithPrefix(pulid_prefix.TypeToPrefix("log_error")),
	}
}
