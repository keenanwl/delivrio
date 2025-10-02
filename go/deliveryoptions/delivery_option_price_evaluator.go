package deliveryoptions

const validTimeFormat = `Mon, 02 Jan 2006`

type RateSheetWarningTimeLimited struct {
	Err string
}

func (e *RateSheetWarningTimeLimited) Error() string {
	return e.Err
}

// Deprecated until we can make this better
/*func CalculatePrice(
	ctx context.Context,
	now time.Time,
	price *ent.DeliveryOptionPrice,
	c pulid.ID,
	totalWeightG *int,
) (float64, *ent.Currency, error) {

	tx := ent.TxFromContext(ctx)

	if !price.CarrierRates {
		currency, err := price.Currency(ctx)
		return price.Price, currency, err
	}

	rs, err := price.QueryRateService().WithRateGroup().Only(ctx)
	if err != nil {
		return 0, nil, err
	}

	var warning error

	if totalWeightG == nil {
		return 0.0, nil, errors.New("weight required for all products to calculate rates dynamically")
	}

	if rs.Edges.RateGroup.ValidityLimited {
		if rs.Edges.RateGroup.ValidFrom == nil || rs.Edges.RateGroup.ValidUntil == nil {
			warning = &RateSheetWarningTimeLimited{Err: "rate sheet time limited, but valid range not set"}
		} else if now.Before(*rs.Edges.RateGroup.ValidFrom) || now.After(*rs.Edges.RateGroup.ValidUntil) {
			warning = &RateSheetWarningTimeLimited{
				Err: fmt.Sprintf("rate sheet is only valid from %s until %s",
					rs.Edges.RateGroup.ValidFrom.Format(validTimeFormat),
					rs.Edges.RateGroup.ValidUntil.Format(validTimeFormat),
				),
			}
		}
	}

	totalProductWeightKG := float64(*totalWeightG) / 1000.0

	dynamicRate, err := tx.Rate.Query().
		Where(rate.MaxWeightGTE(totalProductWeightKG)).
		Where(
			rate.HasRateZoneWith(
				ratezone.HasCountryWith(country.ID(c)),
			),
			rate.HasRateServiceWith(rateservice.ID(rs.ID)),
		).
		// TODO: test ASC/DESC logic
		Order(ent.Asc(rate.FieldMaxWeight)).
		First(ctx)
	if err != nil {
		return 0, nil, err
	}

	rateCurrency, err := rs.Edges.RateGroup.Currency(ctx)
	if err != nil {
		return 0, nil, err
	}

	// TODO: sort order?
	margins, err := price.DeliveryOptionPriceMargin(ctx)
	if err != nil {
		return 0, nil, err
	}

	priceWithMargins, err := AddMargins(dynamicRate.Price, margins)
	if err != nil {
		return 0, nil, err
	}

	return priceWithMargins, rateCurrency, warning

}*/
/*
func AddMargins(basePrice float64, margins []*ent.DeliveryOptionPriceMargin) (float64, error) {
	basePlusMargins := basePrice
	var err error

	for _, m := range margins {
		switch m.MarginType {
		case deliveryoptionpricemargin.MarginTypeRounded:
			basePlusMargins, err = nearest(basePlusMargins, m.MarginValue)
			if err != nil {
				return 0, err
			}
			break
		case deliveryoptionpricemargin.MarginTypeRelative:
			basePlusMargins = (basePlusMargins * m.MarginValue / 100) + basePlusMargins
			break
		case deliveryoptionpricemargin.MarginTypeAbsolute:
			basePlusMargins += m.MarginValue
			break
		}
	}

	return basePlusMargins, nil
}

func nearest(basePrice float64, toDigit float64) (float64, error) {
	synthetic := fmt.Sprintf("%d", int(basePrice))
	syntheticNum, err := strconv.ParseFloat(synthetic[:len(synthetic)-1]+fmt.Sprintf("%d", int(toDigit)), 64)
	if err != nil {
		return 0, err
	}

	diff := math.Abs((math.Abs(syntheticNum-basePrice) - 10))
	high := basePrice + diff
	low := syntheticNum

	if high-basePrice > basePrice-low {
		return low, nil
	}
	return high, nil
}*/
