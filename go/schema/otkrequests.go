package schema

import (
	"delivrio.io/go/schema/pulid-server-prefix"
	"delivrio.io/shared-utils/pulid"
	"entgo.io/ent"
	"entgo.io/ent/schema"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
)

// OTKRequests holds the schema definition for the OTKRequests entity.
type OTKRequests struct {
	ent.Schema
}

// Annotations of the User.
func (OTKRequests) Annotations() []schema.Annotation {
	return nil
}

// Fields of the OTKRequests.
func (OTKRequests) Fields() []ent.Field {
	return []ent.Field{
		field.String("otk"),
	}
}

// Edges of the OTKRequests.
func (OTKRequests) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("users", User.Type).
			Ref("otk_requests").
			Unique(),
	}
}

func (OTKRequests) Mixin() []ent.Mixin {
	return []ent.Mixin{
		BaseMixin{},
		TenantMixin{},
		pulid.MixinWithPrefix(pulid_server_prefix.TypeToPrefix("otk_requests")),
	}
}
