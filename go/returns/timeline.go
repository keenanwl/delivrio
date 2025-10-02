package returns

import (
	"context"
	"delivrio.io/go/ent"
	"delivrio.io/go/ent/changehistory"
	"delivrio.io/go/ent/returncolli"
	"delivrio.io/go/ent/returncollihistory"
	"delivrio.io/shared-utils/pulid"
)

func ReturnColliTimeline(ctx context.Context, cli *ent.Client, returnColliID pulid.ID) ([]*ent.ChangeHistory, error) {
	changeHistory, err := cli.ChangeHistory.Query().
		Where(changehistory.HasReturnColliHistoryWith(
			returncollihistory.HasReturnColliWith(
				returncolli.ID(returnColliID)))).
		Limit(3).
		Order(ent.Desc(changehistory.FieldCreatedAt)).
		All(ctx)
	if err != nil {
		return nil, err
	}

	return changeHistory, nil
}
