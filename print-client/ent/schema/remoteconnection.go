package schema

import (
	"delivrio.io/print-client/ent/schema/pulid_prefix"
	"delivrio.io/shared-utils/pulid"
	"entgo.io/ent"
	"entgo.io/ent/schema/field"
)

// RemoteConnection holds the schema definition for the RemoteConnection entity.
type RemoteConnection struct {
	ent.Schema
}

// Fields of the RemoteConnection.
func (RemoteConnection) Fields() []ent.Field {
	return []ent.Field{
		field.String("remote_url").Unique(),
		field.String("registration_token"),
		field.String("workstation_name"),
		field.Time("last_ping").
			Nillable().
			Optional(),
	}
}

// Edges of the RemoteConnection.
func (RemoteConnection) Edges() []ent.Edge {
	return nil
}

func (RemoteConnection) Mixin() []ent.Mixin {
	return []ent.Mixin{
		pulid.MixinWithPrefix(pulid_prefix.TypeToPrefix("remote_connection")),
	}
}
