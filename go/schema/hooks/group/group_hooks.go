package grouphooks

/* Deprecated entity
func RecordSubscriptionChange() ent.Hook {
	hk := func(next ent.Mutator) ent.Mutator {
		return hook.GroupFunc(func(ctx context.Context, m *ent2.GroupMutation) (ent.Value, error) {
			view := viewer.FromContext(ctx)
			ids, err := m.IDs(ctx)
			if err != nil {
				return nil, err
			}

			if len(ids) == 0 {
				return nil, errors.New("could not determine which group to update")
			}

			tx, err := m.Client().Tx(ctx)
			if err != nil {
				return nil, err
			}

			for _, id := range ids {
				g, err := tx.Group.Query().WithPlan().Where(group.ID(id)).Only(ctx)
				if err != nil {
					return nil, fmt.Errorf("getting group: %w", err)
				}

				nextPlanID, exists := m.PlanID()
				if exists && nextPlanID != g.Edges.Plan.ID {

					err = tx.ChangeHistory.Create().
						SetTenantID(view.TenantID()).
						SetID(view.ContextID()).
						OnConflict(sql.ConflictColumns("id")).
						DoNothing().
						Exec(ctx)
					if err != nil {
						return nil, utils.Rollback(tx, err)
					}

					_, err = tx.PlanHistory.Create().
						SetChangedByID(view.MyId()).
						SetChangedFromID(g.Edges.Plan.ID).
						SetTenantID(view.TenantID()).
						SetChangeHistoryID(view.ContextID()).
						Save(ctx)
					if err != nil {
						return nil, utils.Rollback(tx, fmt.Errorf("creating plan history record: %w", err))
					}
				}
			}

			err = tx.Commit()
			if err != nil {
				return nil, utils.Rollback(tx, err)
			}

			return next.Mutate(ctx, m)
		})
	}
	return hook.On(hk, ent.OpUpdate|ent.OpUpdateOne)
}*/
