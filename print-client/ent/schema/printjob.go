package schema

import (
	"delivrio.io/print-client/ent/schema/pulid_prefix"
	"delivrio.io/shared-utils/pulid"
	"entgo.io/contrib/entgql"
	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
	"time"
)

// PrintJob holds the schema definition for the PrintJob entity.
type PrintJob struct {
	ent.Schema
}

// Fields of the PrintJob.
func (PrintJob) Fields() []ent.Field {
	return []ent.Field{
		field.Enum("status").
			Values("pending", "pending_success", "success", "pending_cancel", "canceled").
			Comment("Cancel needs to be acknowledged by server before moving to cancelled"),
		field.Enum("file_extension").
			Values("pdf", "zpl", "txt", "png").
			Immutable(),
		field.Bool("use_shell").
			Default(false).
			Immutable(),
		field.String("base64_print_data").
			Immutable(),
		field.Strings("messages").
			Optional(),
		field.Time("created_at").
			Default(time.Now).
			Annotations(
				entgql.Skip(entgql.SkipMutationCreateInput | entgql.SkipMutationUpdateInput),
			).
			Immutable(),
	}
}

// Edges of the PrintJob.
func (PrintJob) Edges() []ent.Edge {
	return []ent.Edge{
		edge.To("local_device", LocalDevice.Type).
			Unique().
			Required(),
	}
}

func (PrintJob) Mixin() []ent.Mixin {
	return []ent.Mixin{
		pulid.MixinWithPrefix(pulid_prefix.TypeToPrefix("print_job")),
	}
}
