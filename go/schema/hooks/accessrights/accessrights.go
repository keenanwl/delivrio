package accessrights

import (
	"context"
	"delivrio.io/go/ent"
	"delivrio.io/go/ent/accessright"
	"delivrio.io/go/ent/privacy"
	"delivrio.io/go/ent/seatgroup"
	"delivrio.io/go/ent/seatgroupaccessright"
	"delivrio.io/go/ent/user"
	"delivrio.io/go/viewer"
)

func CheckAccessRightsMutation(accessRightInternalID string) privacy.MutationRule {
	return privacy.ContextQueryMutationRule(func(ctx context.Context) error {
		view := viewer.FromContext(ctx)
		tx := ent.TxFromContext(ctx)

		if view.Background() ||
			view.Admin() ||
			view.CurrentRole() == viewer.BackgroundForTenant {
			return privacy.Skip
		}

		u, err := tx.User.Query().
			Where(user.ID(view.MyId())).
			Only(ctx)
		if err != nil {
			if err != nil {
				return privacy.Denyf("access rights mutate: viewer user not found: %v", err)
			}
		}

		if u.IsAccountOwner {
			return privacy.Skip
		}

		right, err := tx.SeatGroupAccessRight.Query().
			Where(
				seatgroupaccessright.And(
					seatgroupaccessright.HasSeatGroupWith(seatgroup.HasUserWith(user.ID(view.MyId()))),
					seatgroupaccessright.HasAccessRightWith(accessright.InternalID(accessRightInternalID)),
				),
			).Only(ctx)

		if err != nil && !ent.IsNotFound(err) {
			return err
		} else if ent.IsNotFound(err) {
			// Users are not required to have a group
			return privacy.Skip
		}

		if right.Level == seatgroupaccessright.LevelWrite {
			return privacy.Skip
		}

		return privacy.Denyf("you have not been granted access to modify these resources")

	})
}

func CheckAccessRightsQuery(accessRightInternalID string) privacy.QueryRule {
	return privacy.ContextQueryMutationRule(func(ctx context.Context) error {
		view := viewer.FromContext(ctx)
		client := ent.FromContext(ctx)
		tx := ent.TxFromContext(ctx)

		if view.Background() ||
			view.Admin() ||
			view.AnonymousOverride() ||
			view.CurrentRole() == viewer.BackgroundForTenant {
			return privacy.Skip
		}

		// Had hoped for more abstraction client vs tx
		userQuery := client.User.Query()
		if tx != nil {
			userQuery = tx.User.Query()
		}

		u, err := userQuery.
			Where(user.ID(view.MyId())).
			Only(ctx)
		if err != nil {
			return privacy.Denyf("access rights query: viewer user not found: %v", err)
		}

		if u.IsAccountOwner {
			return privacy.Skip
		}

		accessRightsQuery := client.SeatGroupAccessRight.Query()
		if tx != nil {
			accessRightsQuery = tx.SeatGroupAccessRight.Query()
		}

		right, err := accessRightsQuery.
			Where(
				seatgroupaccessright.And(
					seatgroupaccessright.HasSeatGroupWith(seatgroup.HasUserWith(user.ID(view.MyId()))),
					seatgroupaccessright.HasAccessRightWith(accessright.InternalID(accessRightInternalID)),
				),
			).Only(ctx)
		if err != nil && !ent.IsNotFound(err) {
			return err
		} else if ent.IsNotFound(err) {
			// Users are not required to have a group
			return privacy.Skip
		}

		if right.Level == seatgroupaccessright.LevelWrite || right.Level == seatgroupaccessright.LevelRead {
			return privacy.Skip
		}

		return privacy.Denyf("you have not been granted access to view these resources")
	})
}
