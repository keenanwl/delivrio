package common

import (
	"context"
	"delivrio.io/go/ent"
	"delivrio.io/go/ent/enttest"
	"delivrio.io/go/seed"
	"delivrio.io/go/viewer"
	"delivrio.io/shared-utils/pulid"
	"github.com/stretchr/testify/assert"
	"testing"

	_ "github.com/mattn/go-sqlite3"
)

func TestNetOrderLinePrice(t *testing.T) {
	cli := enttest.Open(t, "sqlite3", "file:ent?mode=memory&cache=shared&_fk=1")
	defer cli.Close()
	ctx := ent.NewContext(viewer.NewContext(context.Background(), viewer.UserViewer{Role: viewer.Background, Context: pulid.MustNew("CH")}), cli.Debug())
	cli = cli.Debug()

	seed.DemoData(ctx)

	txCtx, tx, _ := cli.OpenTx(ctx)
	seed.ExtraOrderLines(txCtx)
	assert.NoError(t, tx.Commit())

	ol := cli.OrderLine.Query().
		AllX(ctx)

	price := NetOrderLinePrice(ol[0])
	assert.Equal(t, 801.0, price)

	price = NetOrderLinePrice(ol[1])
	assert.Equal(t, 801.0, price)

	price = NetOrderLinePrice(ol[2])
	assert.Equal(t, 801.0, price)

	price = NetOrderLinePrice(ol[3])
	assert.Equal(t, 991.0, price)

	price = NetOrderLinePrice(ol[4])
	assert.Equal(t, 50.0, price)
}

func TestColliWeightGram(t *testing.T) {
	cli := enttest.Open(t, "sqlite3", "file:ent?mode=memory&cache=shared&_fk=1")
	defer cli.Close()
	ctx := ent.NewContext(viewer.NewContext(context.Background(), viewer.UserViewer{Role: viewer.Background, Context: pulid.MustNew("CH")}), cli)

	seed.DemoData(ctx)
	txCtx, tx, _ := cli.OpenTx(ctx)
	seed.ExtraOrderLines(txCtx)
	tx.Commit()
	ol := cli.OrderLine.Query().
		AllX(ctx)

	weight, err := ColliWeightGram(ctx, ol)
	assert.NoError(t, err)
	assert.Equal(t, 3008, weight)

	weightKG, err := ColliWeightKG(ctx, ol)
	assert.NoError(t, err)
	assert.Equal(t, 3.008, weightKG)

}
