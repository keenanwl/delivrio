package mixins

import (
	"context"

	"delivrio.io/go/ent/privacy"
	"delivrio.io/go/viewer"
	"entgo.io/ent/entql"
)

type (
	// Filter is the interface that wraps the Where function
	// for filtering nodes in queries and mutations.
	Filter interface {
		// Where applies a filter on the executed query/mutation.
		Where(entql.P)
	}

	// The FilterFunc type is an adapter that allows the use of ordinary
	// functions as where for query and mutation types.
	FilterFunc func(context.Context, Filter) error
)

// DenyIfNoViewer is a rule that returns Deny decision if the viewer is
// missing in the context.
func DenyIfNoViewer() privacy.QueryMutationRule {
	return privacy.ContextQueryMutationRule(func(ctx context.Context) error {
		view := viewer.FromContext(ctx)
		if view == nil {
			panic("view is nil")
			return privacy.Denyf("viewer-context is missing")
		}
		// Skip to the next privacy rule (equivalent to returning nil).
		return privacy.Skip
	})
}

// AllowIfAdmin is a rule that returns Allow decision if the viewer is admin.
func AllowIfAdmin() privacy.QueryMutationRule {
	return privacy.ContextQueryMutationRule(func(ctx context.Context) error {
		view := viewer.FromContext(ctx)
		if view.Admin() {
			return privacy.Allow
		}
		// Skip to the next privacy rule (equivalent to returning nil).
		return privacy.Skip
	})
}

/*// FilterTenantRule is a query/mutation rule that where out entities that are not in the tenant.
func FilterTenantRule() privacy.QueryMutationRule {
	// TenantsFilter is an interface to wrap WhereHasTenantWith()
	// predicate that is used by both `Group` and `User` schemas.
	type TenantsFilter interface {
		WhereHasTenantWith(...predicate.Tenant)
	}
	return privacy.FilterFunc(func(ctx context.Context, f privacy.Filter) error {
		view := viewer.FromContext(ctx)
		if view.TenantID().String() == "" {
			return privacy.Denyf("missing tenant information in viewer")
		}
		tf, ok := f.(TenantsFilter)
		if !ok {
			return privacy.Denyf("unexpected filter type %T", f)
		}
		// Make sure that a tenant reads only entities that has an edge to it.
		tf.WhereHasTenantWith(tenant.ID(view.TenantID()))
		// Skip to the next privacy rule (equivalent to return nil).
		return privacy.Skip
	})
}*/

/*func OrderAllowIfOwner() privacy.MutationRule {
	return privacy.OrderMutationRuleFunc(func(ctx context.Context, m *ent.OrderMutation) error {
		tid, exists := m.TenantID()
		if !exists {
			return privacy.Denyf("missing tenant information in mutation")
		}

		view := viewer.FromContext(ctx)
		if view.TenantID() != tid {
			return privacy.Denyf("mismatch tenant-ids for users/order %d != %d", tid, view.TenantID())
		}

		return privacy.Skip
	})
}*/
