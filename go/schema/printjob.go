package schema

import (
	"delivrio.io/go/schema/hooks/history"
	"delivrio.io/go/schema/pulid-server-prefix"
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
			Values("pending", "at_printer", "success", "canceled"),
		field.Enum("file_extension").
			Values("pdf", "zpl", "txt", "png").
			Immutable(),
		field.Enum("document_type").
			Values("parcel_label", "unknown", "packing_list").
			Immutable(),
		field.Strings("printer_messages").
			Optional(),
		field.String("base64_print_data").
			Immutable(),
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
		edge.To("printer", Printer.Type).
			Unique().
			Required(),
		edge.To("colli", Colli.Type).
			Unique(),
		edge.To("shipment_parcel", ShipmentParcel.Type).
			Unique(),
	}
}

func (PrintJob) Mixin() []ent.Mixin {
	return []ent.Mixin{
		BaseMixin{},
		TenantMixin{},
		pulid.MixinWithPrefix(pulid_server_prefix.TypeToPrefix("print_job")),
	}
}

func (PrintJob) Hooks() []ent.Hook {
	return []ent.Hook{
		history.PrintJobCreate(),
	}
}
