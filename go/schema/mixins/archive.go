package mixins

import (
	"context"
	"delivrio.io/go/ent/intercept"
	"entgo.io/ent"
	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/schema/field"
	"entgo.io/ent/schema/mixin"
)

type ArchiveMixin struct {
	// We embed the `mixin.Schema` to avoid
	// implementing the rest of the methods.
	mixin.Schema
}

type archiveKey struct{}

func (ArchiveMixin) Fields() []ent.Field {
	return []ent.Field{
		field.Time("archived_at").
			Optional(),
	}
}

// ExcludeArchived Allow archived to be generally seen so that FK required lookups don't fail
// Add this to the CTX of lists which should exclude the archived entity
func ExcludeArchived(parent context.Context) context.Context {
	return context.WithValue(parent, archiveKey{}, true)
}

func (a ArchiveMixin) Interceptors() []ent.Interceptor {
	return []ent.Interceptor{
		intercept.TraverseFunc(func(ctx context.Context, q intercept.Query) error {
			// Skip soft-delete, means include soft-deleted entities.
			if excludeArchived, _ := ctx.Value(archiveKey{}).(bool); excludeArchived {
				a.P(q)
				return nil
			}
			return nil
		}),
	}
}

// P adds a storage-level predicate to the queries and mutations.
func (a ArchiveMixin) P(w interface{ WhereP(...func(*sql.Selector)) }) {
	w.WhereP(
		sql.FieldIsNull(a.Fields()[0].Descriptor().Name),
	)
}
