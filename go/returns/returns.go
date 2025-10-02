package returns

import (
	"context"
	"delivrio.io/go/ent/colli"
	"delivrio.io/go/ent/returncolli"
	"golang.org/x/exp/maps"

	"delivrio.io/go/ent"
	"delivrio.io/go/ent/order"
	"delivrio.io/go/ent/orderline"
	"delivrio.io/go/viewer"
	"delivrio.io/shared-utils/pulid"
)

type ReturnLineItem struct {
	Units   int
	ClaimID pulid.ID
}

func CreateReturn(ctx context.Context, orderID pulid.ID, portalID pulid.ID, items map[pulid.ID]ReturnLineItem, comment string) ([]pulid.ID, error) {
	tx := ent.TxFromContext(ctx)
	view := viewer.FromContext(ctx)

	linesToReturn, err := tx.OrderLine.Query().
		WithColli().
		Where(
			orderline.And(
				orderline.IDIn(maps.Keys(items)...),
				orderline.HasColliWith(colli.HasOrderWith(order.ID(orderID))),
			),
		).All(ctx)
	if err != nil {
		return nil, err
	}

	outputReturnColliIDs := make([]pulid.ID, 0)
	returnCollis := make(map[pulid.ID]*ent.ReturnColli, 0)
	for _, l := range linesToReturn {

		if _, ok := returnCollis[l.Edges.Colli.ID]; !ok {

			recip, err := l.Edges.Colli.QueryRecipient().
				WithCountry().
				Only(ctx)
			if err != nil {
				return nil, err
			}

			recipCountry, err := recip.QueryCountry().
				Only(ctx)
			if err != nil {
				return nil, err
			}

			recipAdr, err := recip.CloneEntity(tx).
				SetCountry(recipCountry).
				SetTenantID(view.TenantID()).
				Save(ctx)
			if err != nil {
				return nil, err
			}

			connReturnLocation, err := l.Edges.Colli.QueryOrder().
				QueryConnection().
				QueryReturnLocation().
				WithAddress().
				Only(ctx)
			if err != nil {
				return nil, err
			}

			senderCountry, err := connReturnLocation.Edges.Address.Country(ctx)
			if err != nil {
				return nil, err
			}

			senderAdr, err := connReturnLocation.Edges.Address.
				CloneEntity(tx).
				SetTenantID(view.TenantID()).
				SetCountry(senderCountry).
				Save(ctx)
			if err != nil {
				return nil, err
			}

			col, err := tx.ReturnColli.Create().
				SetStatus(returncolli.StatusOpened).
				// Addresses are reversed on the way back
				SetRecipient(senderAdr).
				SetSender(recipAdr).
				SetOrderID(orderID).
				SetReturnPortalID(portalID).
				SetTenantID(view.TenantID()).
				SetComment(comment).
				Save(ctx)
			if err != nil {
				return nil, err
			}
			returnCollis[l.Edges.Colli.ID] = col
			outputReturnColliIDs = append(outputReturnColliIDs, col.ID)
		}
		c := returnCollis[l.Edges.Colli.ID]

		err = tx.ReturnOrderLine.Create().
			SetReturnColli(c).
			SetUnits(items[l.ID].Units).
			SetReturnPortalClaimID(items[l.ID].ClaimID).
			SetTenantID(view.TenantID()).
			SetOrderLine(l).
			Exec(ctx)
		if err != nil {
			return nil, err
		}

	}

	return outputReturnColliIDs, nil
}
