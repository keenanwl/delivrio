package schema

import (
	"context"
	"delivrio.io/go/schema/mixins"
	"fmt"

	"delivrio.io/go/ent/privacy"
	"delivrio.io/go/viewer"
	"delivrio.io/shared-utils/pulid"
	"entgo.io/contrib/entgql"
	"entgo.io/ent"
	"entgo.io/ent/entql"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
	"entgo.io/ent/schema/index"
	"entgo.io/ent/schema/mixin"
)

// BaseMixin for all schemas in the graph.
type BaseMixin struct {
	mixin.Schema
}

func (BaseMixin) Fields() []ent.Field {
	return []ent.Field{}
}

// Policy defines the privacy policy of the BaseMixin.
func (BaseMixin) Policy() ent.Policy {
	return privacy.Policy{
		Mutation: privacy.MutationPolicy{
			mixins.DenyIfNoViewer(),
		},
		Query: privacy.QueryPolicy{
			mixins.DenyIfNoViewer(),
		},
	}
}

type ChangeHistoryMixin struct {
	mixin.Schema
	entity string
}

func ChangeHistoryEntityMixin(entity string) ChangeHistoryMixin {
	return ChangeHistoryMixin{entity: entity}
}

func (c ChangeHistoryMixin) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("change_history", ChangeHistory.Type).
			Ref(fmt.Sprintf("%v_history", c.entity)).
			Unique().
			Required().
			Annotations(entgql.Skip(entgql.SkipMutationCreateInput | entgql.SkipMutationUpdateInput)),
	}
}

// TenantMixin for embedding the tenant info in different schemas.
type TenantMixin struct {
	mixin.Schema
}

func (TenantMixin) Fields() []ent.Field {
	return []ent.Field{
		field.String("tenant_id").
			GoType(pulid.ID("")),
	}
}

// Edges for all schemas that embed TenantMixin.
func (TenantMixin) Edges() []ent.Edge {
	return []ent.Edge{
		edge.To("tenant", Tenant.Type).
			Unique().
			Required().
			Field("tenant_id").
			Annotations(entgql.Skip(entgql.SkipMutationCreateInput | entgql.SkipMutationUpdateInput)),
	}
}

func (TenantMixin) Indexes() []ent.Index {
	return []ent.Index{
		index.Fields("tenant_id"),
	}
}

func FilterTenantRule() privacy.QueryMutationRule {
	// TenantsFilter is an interface to wrap WhereHasTenantWith()
	// predicate.
	type TenantsFilter interface {
		WhereTenantID(p entql.StringP)
	}
	return privacy.FilterFunc(func(ctx context.Context, f privacy.Filter) error {
		view := viewer.FromContext(ctx)
		tid := view.TenantID()

		// Let BackgroundForTenant fallthrough so tenant is included in queries
		if view.Background() {
			// TODO: test this thoroughly..
			return privacy.Skip
		}

		if len(tid) == 0 {
			return privacy.Denyf("missing tenant information in viewer")
		}
		tf, ok := f.(TenantsFilter)
		if !ok {
			return privacy.Denyf("unexpected filter type %T", f)
		}
		// Make sure that a tenant reads only entities that have an edge to it.
		tf.WhereTenantID(entql.StringEQ(string(tid)))
		// Skip to the next privacy rule (equivalent to return nil).
		return privacy.Skip
	})
}

// Policy for all schemas that embed TenantMixin.
func (TenantMixin) Policy() ent.Policy {
	return privacy.Policy{
		Query: privacy.QueryPolicy{
			//rules.AllowIfAdmin(),
			// Filter out entities that are not connected to the tenant.
			// If the viewer is admin, this policy rule is skipped above.
			FilterTenantRule(),
		},
		Mutation: privacy.MutationPolicy{
			//rules.AllowIfAdmin(),
			// Filter out entities that are not connected to the tenant.
			// If the viewer is admin, this policy rule is skipped above.
			FilterTenantRule(),
		},
	}
}
