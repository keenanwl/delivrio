package accessrights_test

import (
	"context"
	"delivrio.io/go/ent"
	"delivrio.io/go/ent/accessright"
	"delivrio.io/go/ent/user"
	"delivrio.io/go/seed"
	"delivrio.io/go/viewer"
	"delivrio.io/shared-utils/pulid"
	"testing"

	"delivrio.io/go/ent/enttest"
	"delivrio.io/go/ent/privacy"
	_ "delivrio.io/go/ent/runtime"
	"delivrio.io/go/ent/seatgroupaccessright"
	"delivrio.io/go/schema/testingutils"
	"github.com/stretchr/testify/require"

	_ "github.com/mattn/go-sqlite3"
)

type Setup struct {
	ctx       context.Context
	bgCtx     context.Context
	tx        *ent.Tx
	client    *ent.Client
	sgar      *ent.SeatGroupAccessRight
	cust      *ent.User
	seatGroup *ent.SeatGroup
}

func initSetup(t *testing.T) *Setup {
	client := enttest.Open(t, "sqlite3", "file:ent?mode=memory&_fk=1")
	ctx := ent.NewContext(context.Background(), client)

	tx, _ := client.Tx(ctx)
	ctx = ent.NewTxContext(ctx, tx)
	seed.Base(ctx)
	bgCtx := testingutils.DefaultBackgroundViewer()
	bgCtx = ent.NewTxContext(bgCtx, tx)

	group := tx.SeatGroup.Create().
		SetName("Test group").
		SetTenantID(seed.GetTenantID()).
		SaveX(bgCtx)

	customer := seed.SeedCustomerUser(bgCtx)

	ctx = viewer.NewContext(ctx, viewer.UserViewer{
		Role:    viewer.CustomerTenant,
		Context: pulid.MustNew("CH"),
		Tenant:  seed.GetTenantID(),
		MyID:    customer.ID,
	})

	accessRight := tx.AccessRight.Query().
		Where(accessright.InternalID("orders")).
		OnlyX(ctx)

	ar := tx.SeatGroupAccessRight.Create().
		SetAccessRight(accessRight).
		SetSeatGroup(group).
		SetTenantID(seed.GetTenantID()).
		SetLevel(seatgroupaccessright.LevelNone).
		SaveX(bgCtx)

	return &Setup{
		ctx:       ctx,
		bgCtx:     bgCtx,
		tx:        tx,
		client:    client,
		sgar:      ar,
		cust:      customer,
		seatGroup: group,
	}
}

func Test_NoGroupButStillMayQuery(t *testing.T) {
	r := initSetup(t)
	defer r.client.Close()

	r.tx.User.Update().
		ClearSeatGroup().
		SetIsAccountOwner(false).
		Where(user.ID(r.cust.ID)).
		ExecX(r.bgCtx)

	_, err := r.tx.Order.Query().All(r.ctx)
	require.NoError(t, err)

	r.tx.User.Update().
		SetSeatGroup(r.seatGroup).
		SetIsAccountOwner(false).
		Where(user.ID(r.cust.ID)).
		ExecX(r.bgCtx)

	_, err = r.tx.Order.Query().All(r.ctx)
	require.ErrorIs(t, err, privacy.Deny)

}

func Test_AccountOwnerAlwaysAllowedQuery(t *testing.T) {
	r := initSetup(t)
	defer r.client.Close()

	r.tx.User.Update().
		SetSeatGroup(r.seatGroup).
		SetIsAccountOwner(true).
		Where(user.ID(r.cust.ID)).
		ExecX(r.bgCtx)

	_, err := r.tx.Order.Query().All(r.ctx)
	require.NoError(t, err)

	r.tx.User.Update().
		SetSeatGroup(r.seatGroup).
		SetIsAccountOwner(false).
		Where(user.ID(r.cust.ID)).
		ExecX(r.bgCtx)

	_, err = r.tx.Order.Query().All(r.ctx)
	require.ErrorIs(t, err, privacy.Deny)

}

func Test_BackgroundViewerMayQuery(t *testing.T) {
	r := initSetup(t)
	defer r.client.Close()

	r.tx.User.Update().
		SetSeatGroup(r.seatGroup).
		SetIsAccountOwner(false).
		Where(user.ID(r.cust.ID)).
		ExecX(r.bgCtx)

	_, err := r.tx.Order.Query().All(r.bgCtx)
	require.NoError(t, err)

	_, err = r.tx.Order.Query().All(r.ctx)
	require.ErrorIs(t, err, privacy.Deny)

}

func Test_AccessRightsMayQuery(t *testing.T) {
	r := initSetup(t)
	defer r.client.Close()

	r.tx.User.Update().
		SetSeatGroup(r.seatGroup).
		SetIsAccountOwner(false).
		Where(user.ID(r.cust.ID)).
		ExecX(r.bgCtx)

	_, err := r.tx.Order.Query().All(r.ctx)
	require.ErrorIs(t, err, privacy.Deny)

	r.tx.SeatGroupAccessRight.Update().
		SetLevel(seatgroupaccessright.LevelRead).
		ExecX(r.ctx)

	_, err = r.tx.Order.Query().All(r.ctx)
	require.NoError(t, err)

	r.tx.SeatGroupAccessRight.Update().
		SetLevel(seatgroupaccessright.LevelWrite).
		ExecX(r.ctx)

	_, err = r.tx.Order.Query().All(r.ctx)
	require.NoError(t, err)

}

func Test_BackgroundViewerMayMutate(t *testing.T) {
	r := initSetup(t)
	defer r.client.Close()

	view := viewer.NewContext(context.Background(), viewer.UserViewer{
		Role: viewer.Background,
	})

	r.tx.User.Update().
		SetSeatGroup(r.seatGroup).
		SetIsAccountOwner(false).
		Where(user.ID(r.cust.ID)).
		ExecX(view)

	_, err := r.tx.Order.Create().Save(view)
	require.NotErrorIs(t, err, privacy.Deny)

	_, err = r.tx.Order.Create().Save(r.ctx)
	require.ErrorIs(t, err, privacy.Deny)

}

func Test_AccountOwnerAlwaysAllowed(t *testing.T) {
	r := initSetup(t)
	defer r.client.Close()

	r.tx.User.Update().
		SetSeatGroup(r.seatGroup).
		SetIsAccountOwner(true).
		Where(user.ID(r.cust.ID)).
		ExecX(r.bgCtx)

	_, err := r.tx.Order.Create().Save(r.ctx)
	require.NotErrorIs(t, err, privacy.Deny)

	r.tx.User.Update().
		SetSeatGroup(r.seatGroup).
		SetIsAccountOwner(false).
		Where(user.ID(r.cust.ID)).
		ExecX(r.bgCtx)

	_, err = r.tx.Order.Create().Save(r.ctx)
	require.ErrorIs(t, err, privacy.Deny)

}

func Test_NoGroupButStillAllowedAccess(t *testing.T) {
	r := initSetup(t)
	defer r.client.Close()

	r.tx.User.Update().
		ClearSeatGroup().
		SetIsAccountOwner(false).
		Where(user.ID(r.cust.ID)).
		ExecX(r.bgCtx)

	_, err := r.tx.Order.Create().Save(r.ctx)
	require.NotErrorIs(t, err, privacy.Deny)

	r.tx.User.Update().
		SetSeatGroup(r.seatGroup).
		SetIsAccountOwner(false).
		Where(user.ID(r.cust.ID)).
		ExecX(r.bgCtx)

	_, err = r.tx.Order.Create().Save(r.ctx)
	require.ErrorIs(t, err, privacy.Deny)

}
