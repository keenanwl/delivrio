package schema

import (
	"context"
	"time"

	ent2 "delivrio.io/go/ent"
	"delivrio.io/go/ent/hook"
	"delivrio.io/go/ent/privacy"
	"delivrio.io/go/schema/pulid-server-prefix"
	"delivrio.io/go/utils"
	"delivrio.io/shared-utils/pulid"
	"entgo.io/contrib/entgql"
	"entgo.io/ent"
	"entgo.io/ent/schema"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
)

// User holds the schema definition for the User entity.
type User struct {
	ent.Schema
}

func (User) Annotations() []schema.Annotation {
	return []schema.Annotation{
		entgql.RelayConnection(),
		entgql.QueryField(),
		entgql.Mutations(
			entgql.MutationCreate(),
			entgql.MutationUpdate(),
		),
	}
}

type ErrUserFields struct {
	InvalidFields map[string]string
}

func (ErrUserFields) Error() string {
	return "one or more user fields is invalid"
}
func (e ErrUserFields) String() string {
	return e.Error()
}
func NewErrUserFields() ErrUserFields {
	return ErrUserFields{
		InvalidFields: map[string]string{},
	}
}

func (User) Hooks() []ent.Hook {

	return []ent.Hook{
		hook.On(
			func(next ent.Mutator) ent.Mutator {
				return hook.UserFunc(func(ctx context.Context, m *ent2.UserMutation) (ent.Value, error) {
					err := NewErrUserFields()
					/*					if num, ok := m.CompanyName(); ok && len(num) > 30 {
											err.InvalidFields[user.FieldCompanyName] = "company name must be less than 30"
										}
										if num, ok := m.VatNumber(); ok && len(num) > 15 {
											err.InvalidFields[user.FieldVatNumber] = "vat number must be less than 15"
										}*/

					if len(err.InvalidFields) == 0 {
						return next.Mutate(ctx, m)
					}
					return nil, err
				})
			},
			ent.OpCreate|ent.OpUpdate|ent.OpUpdateOne,
		),
		hook.On(
			func(next ent.Mutator) ent.Mutator {
				return hook.UserFunc(func(ctx context.Context, m *ent2.UserMutation) (ent.Value, error) {
					password, exists := m.Password()
					if exists {
						m.SetHash(utils.HashPasswordX(password))
						m.SetPassword("X")
					}

					return next.Mutate(ctx, m)
				})
			},
			ent.OpCreate|ent.OpUpdate|ent.OpUpdateOne,
		),
	}

}

// Fields of the User.
func (User) Fields() []ent.Field {
	return []ent.Field{
		field.String("name").Optional(),
		field.String("surname").Optional(),
		field.String("phone_number").Optional(),
		field.String("email").Unique().NotEmpty(),
		field.String("password").
			Sensitive().
			Optional().
			Comment("Field is just for front end convenience. Password gets stored as hash."),
		field.String("hash").Sensitive().
			Annotations(entgql.Skip(entgql.SkipMutationCreateInput | entgql.SkipMutationUpdateInput)),
		field.Bool("is_account_owner").
			Default(false).
			Annotations(entgql.Skip(entgql.SkipMutationCreateInput | entgql.SkipMutationUpdateInput)),
		field.Bool("is_global_admin").
			Default(false).
			Annotations(entgql.Skip(entgql.SkipMutationCreateInput | entgql.SkipMutationUpdateInput)),
		field.Bool("marketing_consent").
			Default(true).
			Optional(),
		field.Time("created_at").
			Default(time.Now).
			Optional().
			Immutable(),
		field.Time("archived_at").
			Optional().
			Nillable(),
		field.Enum("pickup_day").
			Comment("When fulfilling, the next carrier pickup date for the package can be selected").
			Default("Today").
			Values("Today", "Tomorrow", "In_2_Days", "In_3_Days", "In_4_Days", "In_5_Days"),
		field.Time("pickup_day_last_changed").
			Comment("So we can ask the user to confirm their pickup day after X hours").
			Optional(),
	}
}

// Edges of the User.
func (User) Edges() []ent.Edge {
	return []ent.Edge{
		edge.To("otk_requests", OTKRequests.Type),
		edge.To("signup_options", SignupOptions.Type).
			Unique(),
		edge.To("language", Language.Type).
			Unique(),
		edge.From("change_history", ChangeHistory.Type).
			Ref("user"),
		edge.To("plan_history_user", PlanHistory.Type),
		edge.To("api_token", APIToken.Type),
		edge.To("selected_workstation", Workstation.Type).
			Unique(),
		edge.From("seat_group", SeatGroup.Type).
			Ref("user").
			Unique(),
		edge.From("workspace_recent_scan", WorkspaceRecentScan.Type).
			Ref("user"),
	}
}

func (User) Policy() ent.Policy {
	return privacy.Policy{
		Mutation: privacy.MutationPolicy{
			// Deny if not set otherwise.
			//privacy.AlwaysDenyRule(),
		},
		Query: privacy.QueryPolicy{
			// Allow any viewer to read anything.
			//privacy.AlwaysAllowRule(),
		},
	}
}

func (User) Mixin() []ent.Mixin {
	return []ent.Mixin{
		BaseMixin{},
		TenantMixin{},
		pulid.MixinWithPrefix(pulid_server_prefix.TypeToPrefix("user")),
	}
}
