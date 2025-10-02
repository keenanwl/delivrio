package schema

import (
	pulid_server_prefix "delivrio.io/go/schema/pulid-server-prefix"
	"delivrio.io/shared-utils/pulid"
	"entgo.io/contrib/entgql"
	"entgo.io/ent"
	"entgo.io/ent/schema"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
	"time"
)

type DocumentFile struct {
	ent.Schema
}

func (DocumentFile) Annotations() []schema.Annotation {
	return []schema.Annotation{
		entgql.RelayConnection(),
		entgql.QueryField(),
	}
}

func (DocumentFile) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("colli", Colli.Type).
			Ref("document_file").
			Unique(),
		edge.From("shipment_parcel", ShipmentParcel.Type).
			Ref("document_file").
			Unique(),
	}
}

func (DocumentFile) Hooks() []ent.Hook {
	return []ent.Hook{
		//
	}
}

func (DocumentFile) Mixin() []ent.Mixin {
	return []ent.Mixin{
		BaseMixin{},
		TenantMixin{},
		pulid.MixinWithPrefix(pulid_server_prefix.TypeToPrefix("document_file")),
	}
}

func (DocumentFile) Fields() []ent.Field {
	return []ent.Field{
		field.Time("created_at").
			Default(time.Now).
			Annotations(
				entgql.Skip(entgql.SkipMutationCreateInput | entgql.SkipMutationUpdateInput),
			).
			Immutable(),
		field.Enum("storage_type").
			Values("database", "bucket").
			Immutable(),
		field.String("storage_path").
			Optional(),
		field.String("storage_path_zpl").
			Default("").
			Optional(),
		field.Enum("doc_type").
			Values("carrier_label", "packing_slip").
			Immutable(),
		field.String("data_pdf_base64").
			Optional().
			Immutable(),
		field.String("data_zpl_base64").
			Optional().
			Immutable(),
	}
}
