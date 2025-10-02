package schema

import (
	"context"
	ent2 "delivrio.io/go/ent"
	"delivrio.io/go/ent/hook"
	"delivrio.io/go/ent/user"
	"delivrio.io/go/schema/pulid-server-prefix"
	"delivrio.io/shared-utils/pulid"
	"entgo.io/contrib/entgql"
	"entgo.io/ent"
	"entgo.io/ent/schema"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
)

// Contact holds the schema definition for the Contact entity.
type Contact struct {
	ent.Schema
}

func (Contact) Annotations() []schema.Annotation {
	return []schema.Annotation{
		entgql.RelayConnection(),
		entgql.QueryField(),
		entgql.Mutations(
			entgql.MutationCreate(),
			entgql.MutationUpdate(),
		),
	}
}

type ErrContactFields struct {
	InvalidFields map[string]string
}

func (ErrContactFields) Error() string {
	return "one or more contact fields is invalid"
}
func (e ErrContactFields) String() string {
	return e.Error()
}
func NewErrContactFields() ErrContactFields {
	return ErrContactFields{
		InvalidFields: map[string]string{},
	}
}

func (Contact) Hooks() []ent.Hook {

	return []ent.Hook{
		hook.On(
			func(next ent.Mutator) ent.Mutator {
				return hook.ContactFunc(func(ctx context.Context, m *ent2.ContactMutation) (ent.Value, error) {
					err := NewErrContactFields()
					if num, ok := m.Name(); ok && len(num) > 15 {
						err.InvalidFields[user.FieldName] = "name must be less than 15"
					}
					if num, ok := m.Surname(); ok && len(num) > 15 {
						err.InvalidFields[user.FieldSurname] = "surname must be less than 15"
					}

					if len(err.InvalidFields) == 0 {
						return next.Mutate(ctx, m)
					}
					return nil, err
				})
			},
			ent.OpCreate|ent.OpUpdate|ent.OpUpdateOne,
		),
	}

}

// Fields of the Contact.
func (Contact) Fields() []ent.Field {
	return []ent.Field{
		field.String("name"),
		field.String("surname"),
		field.String("email"),
		field.String("phone_number"),
	}
}

// Edges of the Contact.
func (Contact) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("billing_contact", Tenant.Type).
			Ref("billing_contact"),
		edge.From("admin_contact", Tenant.Type).
			Ref("admin_contact"),
	}
}

func (Contact) Mixin() []ent.Mixin {
	return []ent.Mixin{
		BaseMixin{},
		TenantMixin{},
		pulid.MixinWithPrefix(pulid_server_prefix.TypeToPrefix("contact")),
	}
}
