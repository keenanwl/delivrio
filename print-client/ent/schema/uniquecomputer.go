package schema

import (
	"delivrio.io/print-client/ent/schema/pulid_prefix"
	"delivrio.io/shared-utils/pulid"
	"entgo.io/ent"
)

// TLDR: single row, creates a unique ID for this machine
type UniqueComputer struct {
	ent.Schema
}

// Fields of the UniqueComputer.
func (UniqueComputer) Fields() []ent.Field {
	return nil
}

// Edges of the UniqueComputer.
func (UniqueComputer) Edges() []ent.Edge {
	return nil
}

func (UniqueComputer) Mixin() []ent.Mixin {
	return []ent.Mixin{
		pulid.MixinWithPrefix(pulid_prefix.TypeToPrefix("unique_computer")),
	}
}
