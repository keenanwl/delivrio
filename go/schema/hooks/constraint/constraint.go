package constraint

import (
	"context"
	"fmt"
	"strings"
	"time"

	ent2 "delivrio.io/go/ent"
	"delivrio.io/go/ent/deliveryruleconstraint"
	"delivrio.io/go/ent/hook"
	"delivrio.io/go/ent/producttag"
	"delivrio.io/go/schema/fieldjson"
	"delivrio.io/go/schema/hooks"
	"delivrio.io/shared-utils/pulid"
	"entgo.io/ent"
)

func TrimDeliveryRuleConstraint() ent.Hook {
	hk := func(next ent.Mutator) ent.Mutator {
		return hook.DeliveryRuleConstraintFunc(func(ctx context.Context, m *ent2.DeliveryRuleConstraintMutation) (ent.Value, error) {

			selectedValue, exists := m.SelectedValue()
			if exists {
				nextSelectedValue := selectedValue
				nextSelectedValue.Values = trimSpaces(nextSelectedValue.Values)
				nextSelectedValue.IDs = trimSpaces(nextSelectedValue.IDs)
				nextSelectedValue.DayOfWeek = trimSpaces(nextSelectedValue.DayOfWeek)
				nextSelectedValue.TimeOfDay = trimSpaces(nextSelectedValue.TimeOfDay)
				m.SetSelectedValue(nextSelectedValue)
			}
			return next.Mutate(ctx, m)
		})
	}
	return hook.On(hk, ent.OpCreate|ent.OpUpdate|ent.OpUpdateOne)
}
func trimSpaces(values []string) []string {
	for i := range values {
		values[i] = strings.TrimSpace(values[i])
	}
	return values
}

func SaveDeliveryOptionConstraintCreate() ent.Hook {
	hk := func(next ent.Mutator) ent.Mutator {
		return hook.DeliveryRuleConstraintFunc(func(ctx context.Context, m *ent2.DeliveryRuleConstraintMutation) (ent.Value, error) {
			errs := hooks.NewValidationError()

			constraintPropertyType, constraintTypeSet := m.PropertyType()
			if !constraintTypeSet {
				return nil, fmt.Errorf("constraint type must be set")
			}

			comparison, comparisonSet := m.Comparison()
			if !comparisonSet {
				errs.SetError(deliveryruleconstraint.FieldComparison, "missing comparison field value")
			}

			constraintValue, constraintValueSet := m.SelectedValue()
			if !constraintValueSet {
				errs.SetError(deliveryruleconstraint.FieldComparison, "missing option value field value")
			}

			err := validConstraint(ctx, m.Client(), constraintPropertyType, comparison, constraintValue)
			if err != nil {
				errs.SetError(deliveryruleconstraint.FieldSelectedValue, err.Error())
			}

			if len(errs.InvalidFields(ctx)) > 0 {
				return nil, errs
			}

			return next.Mutate(ctx, m)
		})
	}
	return hook.On(hk, ent.OpCreate)
}

func SaveDeliveryOptionConstraintUpdate() ent.Hook {
	hk := func(next ent.Mutator) ent.Mutator {
		return hook.DeliveryRuleConstraintFunc(func(ctx context.Context, m *ent2.DeliveryRuleConstraintMutation) (ent.Value, error) {
			errs := hooks.NewValidationError()
			var err error

			constraintPropertyType, constraintTypeSet := m.PropertyType()
			if !constraintTypeSet {
				return nil, fmt.Errorf("constraint type must be set")
			}

			comparison, comparisonSet := m.Comparison()
			if !comparisonSet {
				comparison, err = m.OldComparison(ctx)
				if err != nil {
					errs.SetError(deliveryruleconstraint.FieldComparison, err.Error())
				}
			}

			constraintValue, constraintValueSet := m.SelectedValue()
			if !constraintValueSet {
				constraintValue, err = m.OldSelectedValue(ctx)
				if err != nil {
					errs.SetError(deliveryruleconstraint.FieldSelectedValue, err.Error())
				}
			}

			err = validConstraint(ctx, m.Client(), constraintPropertyType, comparison, constraintValue)
			if err != nil {
				errs.SetError(deliveryruleconstraint.FieldSelectedValue, err.Error())
			}

			if len(errs.InvalidFields(ctx)) > 0 {
				return nil, errs
			}

			return next.Mutate(ctx, m)
		})
	}
	return hook.On(hk, ent.OpUpdate|ent.OpUpdateOne)
}

func validConstraint(
	ctx context.Context,
	client *ent2.Client,
	constraintType deliveryruleconstraint.PropertyType,
	comparison deliveryruleconstraint.Comparison,
	constraintValue *fieldjson.DeliveryRuleConstraintSelectedValue,
) error {
	switch constraintType {
	case deliveryruleconstraint.PropertyTypeAllProductsTagged:
	case deliveryruleconstraint.PropertyTypeProductTag:
		err := validProductTags(ctx, client, comparison, constraintValue)
		if err != nil {
			return err
		}
		break
	case deliveryruleconstraint.PropertyTypeSku:
		err := validSku(comparison, constraintValue)
		if err != nil {
			return err
		}
		break
	case deliveryruleconstraint.PropertyTypeOrderLines:
		err := validOrderLines(comparison, constraintValue)
		if err != nil {
			return err
		}
		break
	case deliveryruleconstraint.PropertyTypePostalCodeNumeric:
		err := validPostalCodesNumeric(comparison, constraintValue)
		if err != nil {
			return err
		}
		break
	case deliveryruleconstraint.PropertyTypePostalCodeString:
		err := validPostalCodesString(comparison, constraintValue)
		if err != nil {
			return err
		}
		break
	case deliveryruleconstraint.PropertyTypeTimeOfDay:
		err := validTimeOfDay(comparison, constraintValue)
		if err != nil {
			return err
		}
		break
	case deliveryruleconstraint.PropertyTypeDayOfWeek:
		err := validDayOfWeek(comparison, constraintValue)
		if err != nil {
			return err
		}
		break
	case deliveryruleconstraint.PropertyTypeCartTotal:
		err := validCartTotal(comparison, constraintValue)
		if err != nil {
			return err
		}
		break
	case deliveryruleconstraint.PropertyTypeTotalWeight:
		err := validTotalWeight(comparison, constraintValue)
		if err != nil {
			return err
		}
		break
	default:
		return fmt.Errorf("unknown constraint type")

	}
	return nil
}

const ExpectedTimeFormat = `15:04`

func validDayOfWeek(
	comparison deliveryruleconstraint.Comparison,
	value *fieldjson.DeliveryRuleConstraintSelectedValue,
) error {
	switch comparison {
	case deliveryruleconstraint.ComparisonEquals:
	case deliveryruleconstraint.ComparisonNotEquals:
		break
	default:
		return fmt.Errorf("unsupported comparison %v", comparison)
	}

	if len(value.DayOfWeek) == 0 {
		return fmt.Errorf("at least 1 day must be selected")
	}

	for _, d := range value.DayOfWeek {

		switch fieldjson.Weekday(d) {
		case fieldjson.Monday:
		case fieldjson.Tuesday:
		case fieldjson.Wednesday:
		case fieldjson.Thursday:
		case fieldjson.Friday:
		case fieldjson.Saturday:
		case fieldjson.Sunday:
			break
		default:
			return fmt.Errorf("unknown day of week option: %v", d)
		}

	}

	return nil
}

func validTimeOfDay(
	comparison deliveryruleconstraint.Comparison,
	value *fieldjson.DeliveryRuleConstraintSelectedValue,
) error {
	switch comparison {
	case deliveryruleconstraint.ComparisonBetween:
		if len(value.TimeOfDay) != 2 {
			return fmt.Errorf("between comparison requires exactly 2 times")
		}
		break
	case deliveryruleconstraint.ComparisonOutside:
		if len(value.TimeOfDay) != 2 {
			return fmt.Errorf("outside comparison requires exactly 2 times")
		}
		break
	default:
		return fmt.Errorf("unsupported comparison %v", comparison)
	}

	timeFrom, err := time.Parse(ExpectedTimeFormat, value.TimeOfDay[0])
	if err != nil {
		return err
	}

	timeTo, err := time.Parse(ExpectedTimeFormat, value.TimeOfDay[1])
	if err != nil {
		return err
	}

	if timeFrom.After(timeTo) || timeFrom.Equal(timeTo) {
		return fmt.Errorf("start time must be before end time")
	}

	return nil
}

func validPostalCodesNumeric(
	comparison deliveryruleconstraint.Comparison,
	value *fieldjson.DeliveryRuleConstraintSelectedValue,
) error {
	switch comparison {
	case deliveryruleconstraint.ComparisonEquals:
	case deliveryruleconstraint.ComparisonNotEquals:
		break
	case deliveryruleconstraint.ComparisonBetween:
		if len(value.NumericRange) != 2 {
			return fmt.Errorf("between comparison requires exactly 2 zip codes")
		}
		break
	case deliveryruleconstraint.ComparisonOutside:
		if len(value.NumericRange) != 2 {
			return fmt.Errorf("outside comparison requires exactly 2 zip codes")
		}
		break
	default:
		return fmt.Errorf("unsupported comparison %v", comparison)
	}

	if len(value.NumericRange) == 0 {
		return fmt.Errorf("at least 1 zip code must be included")
	}

	return nil
}

func validPostalCodesString(
	comparison deliveryruleconstraint.Comparison,
	value *fieldjson.DeliveryRuleConstraintSelectedValue,
) error {
	switch comparison {
	case deliveryruleconstraint.ComparisonEquals:
	case deliveryruleconstraint.ComparisonNotEquals:
	case deliveryruleconstraint.ComparisonContains:
	case deliveryruleconstraint.ComparisonPrefix:
	case deliveryruleconstraint.ComparisonSuffix:
		break
	default:
		return fmt.Errorf("unsupported comparison %v", comparison)
	}

	if len(value.Values) == 0 {
		return fmt.Errorf("at least 1 postal code must be included")
	}

	return nil
}

func validCartTotal(
	comparison deliveryruleconstraint.Comparison,
	value *fieldjson.DeliveryRuleConstraintSelectedValue,
) error {

	switch comparison {
	case deliveryruleconstraint.ComparisonEquals:
	case deliveryruleconstraint.ComparisonNotEquals:
	case deliveryruleconstraint.ComparisonLessThan:
	case deliveryruleconstraint.ComparisonGreaterThan:
		break
	default:
		return fmt.Errorf("unsupported comparison %v", comparison)
	}

	return nil
}
func validTotalWeight(
	comparison deliveryruleconstraint.Comparison,
	value *fieldjson.DeliveryRuleConstraintSelectedValue,
) error {

	switch comparison {
	case deliveryruleconstraint.ComparisonEquals:
	case deliveryruleconstraint.ComparisonNotEquals:
	case deliveryruleconstraint.ComparisonLessThan:
	case deliveryruleconstraint.ComparisonGreaterThan:
		break
	default:
		return fmt.Errorf("unsupported comparison %v", comparison)
	}

	return nil
}
func validOrderLines(
	comparison deliveryruleconstraint.Comparison,
	value *fieldjson.DeliveryRuleConstraintSelectedValue,
) error {

	switch comparison {
	case deliveryruleconstraint.ComparisonEquals:
	case deliveryruleconstraint.ComparisonNotEquals:
	case deliveryruleconstraint.ComparisonLessThan:
	case deliveryruleconstraint.ComparisonGreaterThan:
		break
	default:
		return fmt.Errorf("unsupported comparison %v", comparison)
	}

	return nil
}

func validSku(
	comparison deliveryruleconstraint.Comparison,
	value *fieldjson.DeliveryRuleConstraintSelectedValue,
) error {

	switch comparison {
	case deliveryruleconstraint.ComparisonEquals:
	case deliveryruleconstraint.ComparisonNotEquals:
	case deliveryruleconstraint.ComparisonContains:
	case deliveryruleconstraint.ComparisonPrefix:
	case deliveryruleconstraint.ComparisonSuffix:
		break
	default:
		return fmt.Errorf("unsupported comparison %v", comparison)
	}

	if len(value.Values) == 0 {
		return fmt.Errorf("SKU value empty")
	}

	return nil

}

func validProductTags(
	ctx context.Context,
	client *ent2.Client,
	comparison deliveryruleconstraint.Comparison,
	value *fieldjson.DeliveryRuleConstraintSelectedValue,
) error {

	switch comparison {
	case deliveryruleconstraint.ComparisonEquals:
	case deliveryruleconstraint.ComparisonNotEquals:
		break
	default:
		return fmt.Errorf("unsupported comparison %v", comparison)
	}

	if len(value.IDs) == 0 {
		return fmt.Errorf("unsupported comparison %v", comparison)
	}

	for i, id := range value.IDs {
		exists, err := client.ProductTag.Query().Where(producttag.ID(pulid.ID(id))).Exist(ctx)
		if err != nil {
			return err
		}

		if !exists {
			return fmt.Errorf("product tag #%v does not exist", i)
		}
	}

	return nil

}
