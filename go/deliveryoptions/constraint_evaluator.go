package deliveryoptions

import (
	"context"
	"delivrio.io/go/ent/deliveryrule"
	"entgo.io/ent/dialect/sql"
	"errors"
	"fmt"
	"strconv"
	"strings"
	"time"

	"delivrio.io/go/ent"
	"delivrio.io/go/ent/deliveryruleconstraint"
	"delivrio.io/go/ent/deliveryruleconstraintgroup"
	"delivrio.io/go/schema/fieldjson"
	"delivrio.io/go/schema/hooks/constraint"
	"delivrio.io/shared-utils/pulid"
)

type ConstraintProductWeight struct {
	WeightG       int
	ProductTagIDs []pulid.ID
	SKU           *string
	UnitPrice     float64
	Units         int
}

type PricePair struct {
	Price    float64
	Currency *ent.Currency
}

func DeliveryOptionMatches(
	ctx context.Context,
	option *ent.DeliveryOption,
	zip string,
	products []*ConstraintProductWeight,
	country pulid.ID,
) (bool, *PricePair, error) {

	rules, err := option.QueryDeliveryRule().
		// TODO: Add test; very important that we show the cheapest option available
		// and then break, so duplicate DO's are not returned
		Order(deliveryrule.ByPrice(sql.OrderAsc())).
		WithCurrency().
		All(ctx)
	if err != nil {
		return false, nil, err
	}

	for _, rule := range rules {
		ruleMatches, err := deliveryRuleMatches(
			ctx,
			rule,
			zip,
			products,
			country,
		)
		if err != nil {
			return false, nil, err
		}
		if ruleMatches {
			return true, &PricePair{
				Price:    rule.Price,
				Currency: rule.Edges.Currency,
			}, nil
		}
	}

	return false, nil, nil
}

func deliveryRuleMatches(
	ctx context.Context,
	rule *ent.DeliveryRule,
	zip string,
	products []*ConstraintProductWeight,
	country pulid.ID,
) (bool, error) {
	constraintGroups, err := rule.QueryDeliveryRuleConstraintGroup().
		All(ctx)
	if err != nil {
		return false, err
	}

	for _, g := range constraintGroups {
		constraintsMatch, err := GroupMatches(
			ctx,
			g,
			zip,
			products,
			country,
		)
		if err != nil {
			return false, nil
		}
		if constraintsMatch {
			return true, nil
		}
	}

	return false, nil
}

func GroupMatches(
	ctx context.Context,
	group *ent.DeliveryRuleConstraintGroup,
	orderPostalCode string,
	products []*ConstraintProductWeight,
	country pulid.ID,
) (bool, error) {

	ruleCountries, err := group.QueryDeliveryRule().
		QueryCountry().
		All(ctx)
	if err != nil {
		return false, err
	}

	countryMatches := false
	for _, c := range ruleCountries {
		if c.ID == country {
			countryMatches = true
			break
		}
	}

	// Allow a potential match below if there are no
	// countries assigned to this rule
	if !countryMatches && len(ruleCountries) > 0 {
		return false, nil
	}

	constraints, err := group.QueryDeliveryRuleConstraints().
		All(ctx)
	if err != nil {
		return false, err
	}

	done := false
	matches := false

	for _, c := range constraints {

		switch c.PropertyType {
		case deliveryruleconstraint.PropertyTypeDayOfWeek:
			// TODO: check store TZ
			done, matches = groupEvaluation(
				group.ConstraintLogic,
				dayOfWeekContains(time.Now(), c.SelectedValue.DayOfWeek),
			)
			break
		case deliveryruleconstraint.PropertyTypeTimeOfDay:
			// TODO: check store TZ
			timeOfDayMatches, err := timeOfDayBetween(time.Now(), c.SelectedValue.TimeOfDay)
			if err != nil {
				// TODO ignore this error
				return false, err
			}
			done, matches = groupEvaluation(
				group.ConstraintLogic,
				timeOfDayMatches,
			)
			break
		case deliveryruleconstraint.PropertyTypeTotalWeight:
			weightSumMatches := productWeightSumMatches(c.Comparison, c.SelectedValue.Numeric, products)
			done, matches = groupEvaluation(
				group.ConstraintLogic,
				weightSumMatches,
			)
			break
		case deliveryruleconstraint.PropertyTypeCartTotal:
			priceSumMatches := orderLinePriceSumMatches(c.Comparison, c.SelectedValue.Numeric, products)
			fmt.Println(priceSumMatches)
			done, matches = groupEvaluation(
				group.ConstraintLogic,
				priceSumMatches,
			)
			break
		case deliveryruleconstraint.PropertyTypeProductTag:
			done, matches = groupEvaluation(
				group.ConstraintLogic,
				productHasAtLeastOneTagMatch(tagsToMap(c.SelectedValue.IDs), products),
			)
			break
		case deliveryruleconstraint.PropertyTypeAllProductsTagged:
			done, matches = groupEvaluation(
				group.ConstraintLogic,
				allProductsTagMatch(tagsToMap(c.SelectedValue.IDs), products),
			)
			break
		case deliveryruleconstraint.PropertyTypeSku:
			allSKU := make([]string, 0)
			for _, p := range products {
				if p.SKU != nil {
					allSKU = append(allSKU, *p.SKU)
				}
			}
			done, matches = groupEvaluation(
				group.ConstraintLogic,
				valuesMatches(c.Comparison, c.SelectedValue.Values, allSKU),
			)
			break
		case deliveryruleconstraint.PropertyTypeOrderLines:
			done, matches = groupEvaluation(
				group.ConstraintLogic,
				countItemsMatches(c.Comparison, c.SelectedValue.Numeric, products),
			)
			break
		case deliveryruleconstraint.PropertyTypePostalCodeNumeric:
			done, matches = groupEvaluation(
				group.ConstraintLogic,
				orderPostalCodeMatchesNumeric(c.Comparison, c.SelectedValue.NumericRange, orderPostalCode),
			)
			break
		case deliveryruleconstraint.PropertyTypePostalCodeString:
			done, matches = groupEvaluation(
				group.ConstraintLogic,
				valuesMatches(c.Comparison, trimSpaces(c.SelectedValue.Values), []string{orderPostalCode}),
			)
			break
		default:
			return false, fmt.Errorf("constraint not implemented: %w", err)
		}

		if done {
			break
		}

	}

	return matches, nil

}

func trimSpaces(values []string) []string {
	for i := range values {
		values[i] = strings.TrimSpace(values[i])
	}
	return values
}

func tagsToMap(tags []string) map[pulid.ID]bool {
	out := make(map[pulid.ID]bool)
	for _, t := range tags {
		out[pulid.ID(t)] = true
	}
	return out
}

func groupEvaluation(
	constraintLogic deliveryruleconstraintgroup.ConstraintLogic,
	conditionMatches bool,
) (done bool, matches bool) {

	if constraintLogic == deliveryruleconstraintgroup.ConstraintLogicAnd && !conditionMatches {
		return true, false
	} else if constraintLogic == deliveryruleconstraintgroup.ConstraintLogicAnd && conditionMatches {
		return false, true
	} else if constraintLogic == deliveryruleconstraintgroup.ConstraintLogicOr && conditionMatches {
		return true, true
	}

	return false, false

}

func orderPostalCodeMatchesNumeric(
	logicType deliveryruleconstraint.Comparison,
	constraintRange []int64,
	postalCode string,
) bool {

	postalCodeNumeric, err := strconv.ParseInt(postalCode, 10, 64)
	if err != nil {
		return false
	}

	if logicType == deliveryruleconstraint.ComparisonEquals {
		for _, cr := range constraintRange {
			if cr == postalCodeNumeric {
				return true
			}
		}
		return false
	} else if logicType == deliveryruleconstraint.ComparisonNotEquals {
		for _, cr := range constraintRange {
			if cr == postalCodeNumeric {
				return false
			}
		}
		return true
	}

	if len(constraintRange) != 2 {
		return false
	}

	inBetween := false
	if constraintRange[0] <= postalCodeNumeric && postalCodeNumeric <= constraintRange[1] {
		inBetween = true
	}

	if logicType == deliveryruleconstraint.ComparisonBetween {
		return inBetween
	} else if logicType == deliveryruleconstraint.ComparisonOutside {
		return !inBetween
	}

	return false

}

func countItemsMatches(
	logicType deliveryruleconstraint.Comparison,
	count int64,
	products []*ConstraintProductWeight,
) bool {

	var totalItems int64 = 0
	for _, p := range products {
		totalItems += int64(p.Units)
	}

	switch logicType {
	case deliveryruleconstraint.ComparisonEquals:
		return totalItems == count
	case deliveryruleconstraint.ComparisonNotEquals:
		return totalItems != count
	case deliveryruleconstraint.ComparisonGreaterThan:
		return totalItems > count
	case deliveryruleconstraint.ComparisonLessThan:
		return totalItems < count
	}

	return false

}

func notInList(
	constraintValues []string,
	checkValues []string,
) bool {
	for _, p := range checkValues {
		for _, check := range constraintValues {
			if strings.EqualFold(p, check) {
				return false
			}
		}
	}
	return true
}

// Constraint value matches are all independent as the single product use case
// is better handled with other constraints (?)
// So 1 match = true
func valuesMatches(
	logicType deliveryruleconstraint.Comparison,
	constraintValues []string,
	checkValues []string,
) bool {
	fmt.Printf("Check values: %v\n", checkValues)
	if logicType == deliveryruleconstraint.ComparisonNotEquals {
		return notInList(constraintValues, checkValues)
	}

	for _, p := range checkValues {
		for _, check := range constraintValues {
			switch logicType {
			case deliveryruleconstraint.ComparisonEquals:
				fmt.Printf("Check values: %v == %v\n", p, check)
				if strings.EqualFold(p, check) {
					return true
				}
			case deliveryruleconstraint.ComparisonPrefix:
				if strings.HasPrefix(strings.ToLower(p), strings.ToLower(check)) {
					return true
				}
			case deliveryruleconstraint.ComparisonContains:
				if strings.Contains(strings.ToLower(p), strings.ToLower(check)) {
					return true
				}
			case deliveryruleconstraint.ComparisonSuffix:
				if strings.HasSuffix(strings.ToLower(p), strings.ToLower(check)) {
					return true
				}
			}
		}
	}

	return false
}

// pulid.ID is sortable, so we can
// speed up the matching later
func allProductsTagMatch(
	tags map[pulid.ID]bool,
	products []*ConstraintProductWeight,
) bool {

	for _, p := range products {
		found := false
		for _, t := range p.ProductTagIDs {
			if tags[t] {
				found = true
				continue
			}
		}
		if !found {
			return false
		}
	}

	return true
}

func productHasAtLeastOneTagMatch(
	tags map[pulid.ID]bool,
	products []*ConstraintProductWeight,
) bool {

	for _, p := range products {
		for _, t := range p.ProductTagIDs {
			if tags[t] {
				return true
			}
		}
	}

	return false
}

func orderLinePriceSumMatches(
	logicType deliveryruleconstraint.Comparison,
	value int64,
	products []*ConstraintProductWeight,
) bool {
	var totalPrice float64 = 0
	for _, l := range products {
		totalPrice += l.UnitPrice * float64(l.Units)
	}

	fmt.Println("PRICE SUM", totalPrice, products)

	switch logicType {
	case deliveryruleconstraint.ComparisonEquals:
		return totalPrice == float64(value)
	case deliveryruleconstraint.ComparisonNotEquals:
		return totalPrice != float64(value)
	case deliveryruleconstraint.ComparisonGreaterThan:
		return totalPrice > float64(value)
	case deliveryruleconstraint.ComparisonLessThan:
		return totalPrice < float64(value)
	}

	return false

}

func productWeightSumMatches(
	logicType deliveryruleconstraint.Comparison,
	value int64,
	products []*ConstraintProductWeight,
) bool {
	totalWeight := 0
	for _, v := range products {
		totalWeight += v.WeightG
	}

	switch logicType {
	case deliveryruleconstraint.ComparisonEquals:
		return int64(totalWeight) == value
	case deliveryruleconstraint.ComparisonNotEquals:
		return int64(totalWeight) != value
	case deliveryruleconstraint.ComparisonGreaterThan:
		return int64(totalWeight) > value
	case deliveryruleconstraint.ComparisonLessThan:
		return int64(totalWeight) < value
	}

	return false

}

func timeOfDayBetween(now time.Time, timeOfDay []string) (bool, error) {

	if len(timeOfDay) != 2 {
		return false, errors.New("expected exactly 2 times to calculate \"between\"")
	}

	start, err := time.ParseInLocation(constraint.ExpectedTimeFormat, timeOfDay[0], now.Location())
	if err != nil {
		return false, err
	}

	start = time.Date(
		now.Year(),
		now.Month(),
		now.Day(),
		start.Hour(),
		start.Minute(),
		start.Second(),
		start.Nanosecond(),
		now.Location(),
	)

	end, err := time.Parse(constraint.ExpectedTimeFormat, timeOfDay[1])
	if err != nil {
		return false, err
	}

	end = time.Date(
		now.Year(),
		now.Month(),
		now.Day(),
		end.Hour(),
		end.Minute(),
		end.Second(),
		end.Nanosecond(),
		now.Location(),
	)

	if now.Equal(start) {
		return true, nil
	}

	if now.After(start) && now.Before(end) {
		return true, nil
	}

	return false, nil
}

func dayOfWeekContains(now time.Time, daysOfWeek []string) bool {
	weekday := now.Weekday()

	dowMap := make(map[time.Weekday]bool)
	for _, d := range daysOfWeek {
		switch fieldjson.Weekday(d) {
		case fieldjson.Monday:
			dowMap[time.Monday] = true
			break
		case fieldjson.Tuesday:
			dowMap[time.Tuesday] = true
			break
		case fieldjson.Wednesday:
			dowMap[time.Wednesday] = true
			break
		case fieldjson.Thursday:
			dowMap[time.Thursday] = true
			break
		case fieldjson.Friday:
			dowMap[time.Friday] = true
			break
		case fieldjson.Saturday:
			dowMap[time.Saturday] = true
			break
		case fieldjson.Sunday:
			dowMap[time.Sunday] = true
			break
		}

	}

	fmt.Println(dowMap, weekday)

	if dowMap[weekday] {
		return true
	}

	return false

}
