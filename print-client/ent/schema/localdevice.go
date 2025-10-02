package schema

import (
	"delivrio.io/print-client/ent/schema/pulid_prefix"
	"delivrio.io/shared-utils/pulid"
	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
)

// LocalDevice holds the schema definition for the LocalDevice entity.
type LocalDevice struct {
	ent.Schema
}

// Fields of the LocalDevice.
func (LocalDevice) Fields() []ent.Field {
	return []ent.Field{
		field.String("name"),
		field.String("system_name").
			Unique(),
		field.Int("vendor_id").
			Optional(),
		field.Int("product_id").
			Optional(),
		field.String("address").
			Optional(),
		field.Bool("active").
			Comment("Toggles which devices are available to remote"),
		field.Bool("archived").
			Comment("Toggles whether to show printers on the local list (meant for hiding network printers)").
			Default(false),
		field.Enum("category").
			Values("scanner", "printer"),
	}
}

// Edges of the LocalDevice.
func (LocalDevice) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("print_job", PrintJob.Type).
			Ref("local_device"),
	}
}

func (LocalDevice) Mixin() []ent.Mixin {
	return []ent.Mixin{
		pulid.MixinWithPrefix(pulid_prefix.TypeToPrefix("local_device")),
	}
}
