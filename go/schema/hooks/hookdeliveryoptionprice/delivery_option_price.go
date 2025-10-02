package hookdeliveryoptionprice

/*
func CreateUpdatePrice() ent.Hook {
	hk := func(next ent.Mutator) ent.Mutator {
		return hook.DeliveryOptionPriceFunc(func(ctx context.Context, m *ent2.DeliveryOptionPriceMutation) (ent.Value, error) {
			errs := hooks.NewValidationError()

			price, priceSet := m.Price()
			if priceSet && price < 0 {
				errs.SetError(deliveryoptionprice.FieldPrice, "price may not be negative")
				return nil, errs
			}

			return next.Mutate(ctx, m)
		})
	}
	return hook.On(hk, ent.OpCreate|ent.OpUpdate|ent.OpUpdateOne)
}

func CreateUpdatePriceMargin() ent.Hook {
	hk := func(next ent.Mutator) ent.Mutator {
		return hook.DeliveryOptionPriceMarginFunc(func(ctx context.Context, m *ent2.DeliveryOptionPriceMarginMutation) (ent.Value, error) {
			errs := hooks.NewValidationError()

			value, _ := m.MarginValue()
			marginType, _ := m.MarginType()

			if value < 0 {
				errs.SetError(deliveryoptionpricemargin.FieldMarginValue, "value must be at least 0")
				return nil, errs
			}

			if marginType == deliveryoptionpricemargin.MarginTypeRounded && value > 9 {
				errs.SetError(deliveryoptionpricemargin.FieldMarginValue, "value of nearest type must be between 0 and 9")
				return nil, errs
			}

			return next.Mutate(ctx, m)
		})
	}
	return hook.On(hk, ent.OpCreate|ent.OpUpdate|ent.OpUpdateOne)
}
*/
