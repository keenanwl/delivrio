package deliveryoptions

import (
	"delivrio.io/go/ent/deliveryruleconstraint"
	"delivrio.io/go/ent/deliveryruleconstraintgroup"
	"delivrio.io/go/schema/fieldjson"
	"delivrio.io/shared-utils/pulid"
	"fmt"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"testing"
	"time"
)

func Test_tagsToMap(t *testing.T) {
	r := tagsToMap([]string{
		"1234",
		"abc",
		"ABC",
	})
	expected := map[pulid.ID]bool{
		"1234": true,
		"abc":  true,
		"ABC":  true,
	}
	require.Equal(t, expected, r)
}

func Test_groupEvaluation(t *testing.T) {
	done, matches := groupEvaluation(
		deliveryruleconstraintgroup.ConstraintLogicAnd,
		true,
	)
	require.False(t, done, "expect AND logic where condition true to continue to next loop iteration")
	require.True(t, matches, "expect true input to also come out again")

	done, matches = groupEvaluation(
		deliveryruleconstraintgroup.ConstraintLogicAnd,
		false,
	)
	require.Truef(t, done, "expect AND logic where condition false to break loop -> done")
	require.False(t, matches, "expect false input to also come out again")

	done, matches = groupEvaluation(
		deliveryruleconstraintgroup.ConstraintLogicOr,
		true,
	)
	require.Truef(t, done, "expect OR logic where condition true to break loop -> done")
	require.Truef(t, matches, "expect true input to also come out again")

	done, matches = groupEvaluation(
		deliveryruleconstraintgroup.ConstraintLogicOr,
		false,
	)
	require.Falsef(t, done, "expect OR logic where condition false to continue to next loop iteration")
	require.Falsef(t, matches, "expect false input to also come out again")
}

func TestValuesMatches(t *testing.T) {
	tests := []struct {
		logicType        deliveryruleconstraint.Comparison
		constraintValues []string
		checkValues      []string
		expected         bool
	}{
		{
			logicType:        deliveryruleconstraint.ComparisonNotEquals,
			constraintValues: []string{"A", "B"},
			checkValues:      []string{"a", "b"},
			expected:         false,
		},
		{
			logicType:        deliveryruleconstraint.ComparisonNotEquals,
			constraintValues: []string{"5960", "5970", "5985"},
			checkValues:      []string{"5960"},
			expected:         false,
		},
		{
			logicType:        deliveryruleconstraint.ComparisonNotEquals,
			constraintValues: []string{"5960", "5970", "5985"},
			checkValues:      []string{"59601"},
			expected:         true,
		},
		{
			logicType:        deliveryruleconstraint.ComparisonEquals,
			constraintValues: []string{"A", "B"},
			checkValues:      []string{"a", "b"},
			expected:         true,
		},
		{
			logicType:        deliveryruleconstraint.ComparisonEquals,
			constraintValues: []string{"A", "B"},
			checkValues:      []string{"A", "B"},
			expected:         true,
		},
		{
			logicType:        deliveryruleconstraint.ComparisonEquals,
			constraintValues: []string{"A", "B"},
			checkValues:      []string{"A", "A"},
			expected:         true,
		},
		{
			logicType:        deliveryruleconstraint.ComparisonEquals,
			constraintValues: []string{"A", "B"},
			checkValues:      []string{"A", "C"},
			expected:         true,
		},
		{
			logicType:        deliveryruleconstraint.ComparisonEquals,
			constraintValues: []string{"B", "A"},
			checkValues:      []string{"A", "C"},
			expected:         true,
		},
		{
			logicType:        deliveryruleconstraint.ComparisonEquals,
			constraintValues: []string{"A", "B"},
			checkValues:      []string{"C", "D"},
			expected:         false,
		},
		{
			logicType:        deliveryruleconstraint.ComparisonPrefix,
			constraintValues: []string{"aB", "bc"},
			checkValues:      []string{"abc", "cd"},
			expected:         true,
		},
		{
			logicType:        deliveryruleconstraint.ComparisonPrefix,
			constraintValues: []string{"ab", "bc"},
			checkValues:      []string{"abc", "cd"},
			expected:         true,
		},
		{
			logicType:        deliveryruleconstraint.ComparisonPrefix,
			constraintValues: []string{"a1", "cd"},
			checkValues:      []string{"abc", "cd"},
			expected:         true,
		},
		{
			logicType:        deliveryruleconstraint.ComparisonPrefix,
			constraintValues: []string{"a1", "CD"},
			checkValues:      []string{"abc", "cd"},
			expected:         true,
		},
		{
			logicType:        deliveryruleconstraint.ComparisonPrefix,
			constraintValues: []string{"a1", "ABCD"},
			checkValues:      []string{"abc", "cd"},
			expected:         false,
		},
		{
			logicType:        deliveryruleconstraint.ComparisonPrefix,
			constraintValues: []string{"a1", "CD"},
			checkValues:      []string{"abc", "abcd"},
			expected:         false,
		},
		{
			logicType:        deliveryruleconstraint.ComparisonContains,
			constraintValues: []string{"abc", "def"},
			checkValues:      []string{"123", "abcdef"},
			expected:         true,
		},
		{
			logicType:        deliveryruleconstraint.ComparisonContains,
			constraintValues: []string{"abc", "cde"},
			checkValues:      []string{"123", "abcdef"},
			expected:         true,
		},
		{
			logicType:        deliveryruleconstraint.ComparisonContains,
			constraintValues: []string{"abc", "def"},
			checkValues:      []string{"123", "456"},
			expected:         false,
		},
		{
			logicType:        deliveryruleconstraint.ComparisonSuffix,
			constraintValues: []string{"aBC", "BC"},
			checkValues:      []string{"abc", "cd"},
			expected:         true,
		},
		{
			logicType:        deliveryruleconstraint.ComparisonSuffix,
			constraintValues: []string{"ab", "12"},
			checkValues:      []string{"abc", "cd"},
			expected:         false,
		},
		{
			logicType:        deliveryruleconstraint.ComparisonSuffix,
			constraintValues: []string{"a1", "cd"},
			checkValues:      []string{"abc", "cd"},
			expected:         true,
		},
		{
			logicType:        deliveryruleconstraint.ComparisonSuffix,
			constraintValues: []string{"a1", "CD"},
			checkValues:      []string{"abc", "cd"},
			expected:         true,
		},
		{
			logicType:        deliveryruleconstraint.ComparisonSuffix,
			constraintValues: []string{"a1", "ABCD"},
			checkValues:      []string{"abc", "AB"},
			expected:         false,
		},
		{
			logicType:        deliveryruleconstraint.ComparisonSuffix,
			constraintValues: []string{"a1", "CD"},
			checkValues:      []string{"abc", "abcd"},
			expected:         true,
		},
	}

	for _, tt := range tests {
		result := valuesMatches(tt.logicType, tt.constraintValues, tt.checkValues)
		assert.Equal(t, tt.expected, result, "valuesMatches failed for logicType %v, %v %v", tt.logicType, tt.constraintValues, tt.checkValues)
	}
}

func TestOrderPostalCodeMatchesNumeric(t *testing.T) {
	tests := []struct {
		logicType       deliveryruleconstraint.Comparison
		postalCode      string
		constraintRange []int64
		expected        bool
	}{
		{
			logicType:       deliveryruleconstraint.ComparisonBetween,
			postalCode:      "12345",
			constraintRange: []int64{10000, 20000},
			expected:        true,
		},
		{
			logicType:       deliveryruleconstraint.ComparisonOutside,
			postalCode:      "12345",
			constraintRange: []int64{10000, 20000},
			expected:        false,
		},
		{
			logicType:       deliveryruleconstraint.ComparisonBetween,
			postalCode:      "12345",
			constraintRange: []int64{20000, 30000},
			expected:        false,
		},
		{
			logicType:       deliveryruleconstraint.ComparisonOutside,
			postalCode:      "12345",
			constraintRange: []int64{20000, 30000},
			expected:        true,
		},
		{
			logicType:       deliveryruleconstraint.ComparisonOutside,
			postalCode:      "ABCDEFG",
			constraintRange: []int64{20000, 30000},
			expected:        false,
		},
		{
			logicType:       deliveryruleconstraint.ComparisonEquals,
			postalCode:      "ABCDEFG",
			constraintRange: []int64{20000, 30000},
			expected:        false,
		},
		{
			logicType:       deliveryruleconstraint.ComparisonNotEquals,
			postalCode:      "ABCDEFG",
			constraintRange: []int64{20000, 30000},
			expected:        false,
		},
		{
			logicType:       deliveryruleconstraint.ComparisonEquals,
			postalCode:      "20000",
			constraintRange: []int64{20000, 30000},
			expected:        true,
		},
		{
			logicType:       deliveryruleconstraint.ComparisonEquals,
			postalCode:      "200001",
			constraintRange: []int64{20000, 30000},
			expected:        false,
		},
		{
			logicType:       deliveryruleconstraint.ComparisonNotEquals,
			postalCode:      "20000",
			constraintRange: []int64{20000, 30000},
			expected:        false,
		},
		{
			logicType:       deliveryruleconstraint.ComparisonNotEquals,
			postalCode:      "200001",
			constraintRange: []int64{20000, 30000},
			expected:        true,
		},
	}

	for _, tt := range tests {
		result := orderPostalCodeMatchesNumeric(tt.logicType, tt.constraintRange, tt.postalCode)
		assert.Equal(t, tt.expected, result, "orderPostalCodeMatchesNumeric failed for logicType %v", tt.logicType)
	}
}

func TestCountItemsMatches(t *testing.T) {
	testCases := []struct {
		logicType    deliveryruleconstraint.Comparison
		count        int64
		products     []*ConstraintProductWeight
		expectedBool bool
	}{
		{
			deliveryruleconstraint.ComparisonEquals,
			9,
			[]*ConstraintProductWeight{{Units: 2}, {Units: 3}, {Units: 4}},
			true,
		},
		{
			deliveryruleconstraint.ComparisonNotEquals,
			9,
			[]*ConstraintProductWeight{{Units: 2}, {Units: 3}, {Units: 4}},
			false,
		},
		{
			deliveryruleconstraint.ComparisonGreaterThan,
			6,
			[]*ConstraintProductWeight{{Units: 2}, {Units: 3}, {Units: 4}},
			true,
		},
		{
			deliveryruleconstraint.ComparisonGreaterThan,
			6,
			[]*ConstraintProductWeight{{Units: 1}, {Units: 2}, {Units: 3}},
			false,
		},
		{
			deliveryruleconstraint.ComparisonLessThan,
			10,
			[]*ConstraintProductWeight{{Units: 2}, {Units: 3}, {Units: 4}},
			true,
		},
		{
			deliveryruleconstraint.ComparisonLessThan,
			10,
			[]*ConstraintProductWeight{{Units: 2}, {Units: 3}, {Units: 5}},
			false,
		},
	}

	for _, tc := range testCases {
		result := countItemsMatches(tc.logicType, tc.count, tc.products)
		if result != tc.expectedBool {
			t.Errorf("Expected %v, got %v for countItemsMatches(%v, %v, %v)", tc.expectedBool, result, tc.logicType, tc.count, tc.products)
		}
	}
}

func TestAllProductsTagMatch(t *testing.T) {
	testCases := []struct {
		tags            map[pulid.ID]bool
		products        []*ConstraintProductWeight
		expectedMatches bool
	}{
		{
			map[pulid.ID]bool{"1": true},
			[]*ConstraintProductWeight{{ProductTagIDs: []pulid.ID{"1"}}},
			true,
		},
		{
			map[pulid.ID]bool{"1": true, "2": true},
			[]*ConstraintProductWeight{{ProductTagIDs: []pulid.ID{"1"}}, {ProductTagIDs: []pulid.ID{"2"}}},
			true,
		},
		{
			map[pulid.ID]bool{"1": true},
			[]*ConstraintProductWeight{{ProductTagIDs: []pulid.ID{"1"}}, {ProductTagIDs: []pulid.ID{"2"}}},
			false,
		},
	}

	for _, tc := range testCases {
		result := allProductsTagMatch(tc.tags, tc.products)
		if result != tc.expectedMatches {
			t.Errorf("Expected %v, got %v for allProductsTagMatch(%v, %v)", tc.expectedMatches, result, tc.tags, tc.products)
		}
	}
}

func TestProductHasAtLeastOneTagMatch(t *testing.T) {
	testCases := []struct {
		tags            map[pulid.ID]bool
		products        []*ConstraintProductWeight
		expectedMatches bool
	}{
		{
			map[pulid.ID]bool{"1": true},
			[]*ConstraintProductWeight{{ProductTagIDs: []pulid.ID{"1"}}},
			true,
		},
		{
			map[pulid.ID]bool{"1": true},
			[]*ConstraintProductWeight{{ProductTagIDs: []pulid.ID{"2"}}},
			false,
		},
		{
			map[pulid.ID]bool{"1": true, "2": true},
			[]*ConstraintProductWeight{{ProductTagIDs: []pulid.ID{"1"}}, {ProductTagIDs: []pulid.ID{"2"}}},
			true,
		},
		{
			map[pulid.ID]bool{"1": true},
			[]*ConstraintProductWeight{{ProductTagIDs: []pulid.ID{"2"}}, {ProductTagIDs: []pulid.ID{"3"}}},
			false,
		},
	}

	for _, tc := range testCases {
		result := productHasAtLeastOneTagMatch(tc.tags, tc.products)
		if result != tc.expectedMatches {
			t.Errorf("Expected %v, got %v for productHasAtLeastOneTagMatch(%v, %v)", tc.expectedMatches, result, tc.tags, tc.products)
		}
	}
}

func TestOrderLinePriceSumMatches(t *testing.T) {
	testCases := []struct {
		logicType deliveryruleconstraint.Comparison
		value     int64
		products  []*ConstraintProductWeight
		expected  bool
	}{
		{
			deliveryruleconstraint.ComparisonEquals,
			122,
			[]*ConstraintProductWeight{{UnitPrice: 20.5, Units: 2}, {UnitPrice: 40.5, Units: 2}},
			true,
		},
		{
			deliveryruleconstraint.ComparisonEquals,
			120,
			[]*ConstraintProductWeight{{UnitPrice: 20.5, Units: 3}, {UnitPrice: 40.0, Units: 2}},
			false,
		},
		{
			deliveryruleconstraint.ComparisonNotEquals,
			120,
			[]*ConstraintProductWeight{{UnitPrice: 20.5, Units: 3}, {UnitPrice: 40.0, Units: 2}},
			true,
		},
		{
			deliveryruleconstraint.ComparisonGreaterThan,
			100,
			[]*ConstraintProductWeight{{UnitPrice: 20.0, Units: 3}, {UnitPrice: 40.0, Units: 2}},
			true,
		},
		{
			deliveryruleconstraint.ComparisonGreaterThan,
			100,
			[]*ConstraintProductWeight{},
			false,
		},
		{
			deliveryruleconstraint.ComparisonGreaterThan,
			100,
			[]*ConstraintProductWeight{{UnitPrice: 0.00001, Units: 300000}},
			false,
		},
		{
			deliveryruleconstraint.ComparisonLessThan,
			150,
			[]*ConstraintProductWeight{{UnitPrice: 20.0, Units: 3}, {UnitPrice: 40.0, Units: 2}},
			true,
		},
		{
			deliveryruleconstraint.ComparisonLessThan,
			150,
			[]*ConstraintProductWeight{{UnitPrice: 20.0, Units: 5}, {UnitPrice: 40.0, Units: 2}},
			false,
		},
	}

	for _, tc := range testCases {
		result := orderLinePriceSumMatches(tc.logicType, tc.value, tc.products)
		if result != tc.expected {
			t.Errorf("Expected %v, got %v for orderLinePriceSumMatches(%v, %v, %v)", tc.expected, result, tc.logicType, tc.value, tc.products)
		}
	}
}

func TestProductWeightSumMatches(t *testing.T) {
	testCases := []struct {
		logicType deliveryruleconstraint.Comparison
		value     int64
		products  []*ConstraintProductWeight
		expected  bool
	}{
		{
			deliveryruleconstraint.ComparisonEquals,
			10,
			[]*ConstraintProductWeight{{WeightG: 5}, {WeightG: 5}},
			true,
		},
		{
			deliveryruleconstraint.ComparisonNotEquals,
			10,
			[]*ConstraintProductWeight{{WeightG: 5}, {WeightG: 5}},
			false,
		},
		{
			deliveryruleconstraint.ComparisonGreaterThan,
			10,
			[]*ConstraintProductWeight{{WeightG: 6}, {WeightG: 5}},
			true,
		},
		{
			deliveryruleconstraint.ComparisonGreaterThan,
			10,
			[]*ConstraintProductWeight{{WeightG: 5}, {WeightG: 5}},
			false,
		},
		{
			deliveryruleconstraint.ComparisonLessThan,
			10,
			[]*ConstraintProductWeight{{WeightG: 5}, {WeightG: 5}},
			false,
		},
	}

	for _, tc := range testCases {
		result := productWeightSumMatches(tc.logicType, tc.value, tc.products)
		if result != tc.expected {
			t.Errorf("Expected %v, got %v for productWeightSumMatches(%v, %v, %v)", tc.expected, result, tc.logicType, tc.value, tc.products)
		}
	}
}

func TestTimeOfDayBetween(t *testing.T) {
	testCases := []struct {
		now       time.Time
		timeOfDay []string
		expected  bool
		err       error
	}{
		{
			time.Date(2023, 1, 1, 12, 0, 0, 0, time.UTC),
			[]string{"11:00", "13:00"},
			true,
			nil,
		},
		{
			time.Date(2023, 1, 1, 13, 0, 0, 0, time.UTC),
			[]string{"11:00", "13:00"},
			false,
			nil,
		},
		{
			time.Date(2023, 1, 1, 11, 0, 0, 0, time.UTC),
			[]string{"11:00", "13:00"},
			true,
			nil,
		},
		{
			time.Date(2023, 1, 1, 12, 59, 59, 0, time.UTC),
			[]string{"11:00", "13:00"},
			true,
			nil,
		},
		{
			time.Date(2023, 1, 1, 14, 0, 0, 0, time.UTC),
			[]string{"11:00", "13:00"},
			false,
			nil,
		},
		{
			time.Date(2023, 1, 1, 12, 0, 0, 0, time.UTC),
			[]string{"13:00", "11:00"},
			false,
			nil,
		},
		{
			time.Date(2023, 1, 1, 12, 0, 0, 0, time.UTC),
			[]string{"11:00"},
			false,
			fmt.Errorf("expected exactly 2 times to calculate \"between\""),
		},
	}

	for _, tc := range testCases {
		result, err := timeOfDayBetween(tc.now, tc.timeOfDay)
		assert.Equal(t, tc.err, err)
		assert.Equal(t, tc.expected, result, "%v is between %v", tc.now.Format(time.TimeOnly), tc.timeOfDay)
	}
}

func intptr(i int) *int {
	return &i
}

func TestDayOfWeekContains(t *testing.T) {
	testCases := []struct {
		name       string
		now        time.Time
		daysOfWeek []string
		want       bool
	}{
		{
			name:       "Today is Monday",
			now:        time.Date(2023, 1, 2, 0, 0, 0, 0, time.UTC),
			daysOfWeek: []string{fieldjson.Monday.String()},
			want:       true,
		},
		{
			name:       "Today is Tuesday",
			now:        time.Date(2023, 1, 3, 0, 0, 0, 0, time.UTC),
			daysOfWeek: []string{fieldjson.Tuesday.String()},
			want:       true,
		},
		{
			name:       "Today is Wednesday",
			now:        time.Date(2023, 1, 4, 0, 0, 0, 0, time.UTC),
			daysOfWeek: []string{fieldjson.Wednesday.String()},
			want:       true,
		},
		{
			name:       "Today is Thursday",
			now:        time.Date(2023, 1, 5, 0, 0, 0, 0, time.UTC),
			daysOfWeek: []string{fieldjson.Thursday.String()},
			want:       true,
		},
		{
			name:       "Today is Friday",
			now:        time.Date(2023, 1, 6, 0, 0, 0, 0, time.UTC),
			daysOfWeek: []string{fieldjson.Friday.String()},
			want:       true,
		},
		{
			name:       "Today is Saturday",
			now:        time.Date(2023, 1, 7, 0, 0, 0, 0, time.UTC),
			daysOfWeek: []string{fieldjson.Saturday.String()},
			want:       true,
		},
		{
			name:       "Today is Sunday",
			now:        time.Date(2023, 1, 8, 0, 0, 0, 0, time.UTC),
			daysOfWeek: []string{fieldjson.Sunday.String()},
			want:       true,
		},
		{
			name:       "Today is not specified in daysOfWeek",
			now:        time.Date(2023, 1, 9, 0, 0, 0, 0, time.UTC),
			daysOfWeek: []string{fieldjson.Thursday.String(), fieldjson.Tuesday.String(), fieldjson.Wednesday.String()},
			want:       false,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			got := dayOfWeekContains(tc.now, tc.daysOfWeek)
			if got != tc.want {
				t.Errorf("got %t, want %t", got, tc.want)
			}
		})
	}
}

func TestTagsToMap(t *testing.T) {
	testCases := []struct {
		name     string
		tags     []string
		expected map[pulid.ID]bool
	}{
		{
			name:     "Empty input",
			tags:     []string{},
			expected: map[pulid.ID]bool{},
		},
		{
			name: "Unique tags",
			tags: []string{"tag1", "tag2", "tag3"},
			expected: map[pulid.ID]bool{
				pulid.ID("tag1"): true,
				pulid.ID("tag2"): true,
				pulid.ID("tag3"): true,
			},
		},
		{
			name: "Duplicate tags",
			tags: []string{"tag1", "tag2", "tag1"},
			expected: map[pulid.ID]bool{
				pulid.ID("tag1"): true,
				pulid.ID("tag2"): true,
			},
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			got := tagsToMap(tc.tags)
			assert.Equal(t, got, tc.expected)
		})
	}
}

func TestGroupEvaluation(t *testing.T) {
	testCases := []struct {
		name             string
		constraintLogic  deliveryruleconstraintgroup.ConstraintLogic
		conditionMatches bool
		expectedDone     bool
		expectedMatches  bool
	}{
		{
			name:             "ConstraintLogicAnd with false condition",
			constraintLogic:  deliveryruleconstraintgroup.ConstraintLogicAnd,
			conditionMatches: false,
			expectedDone:     true,
			expectedMatches:  false,
		},
		{
			name:             "ConstraintLogicAnd with true condition",
			constraintLogic:  deliveryruleconstraintgroup.ConstraintLogicAnd,
			conditionMatches: true,
			expectedDone:     false,
			expectedMatches:  true,
		},
		{
			name:             "ConstraintLogicOr with true condition",
			constraintLogic:  deliveryruleconstraintgroup.ConstraintLogicOr,
			conditionMatches: true,
			expectedDone:     true,
			expectedMatches:  true,
		},
		{
			name:             "ConstraintLogicOr with true condition",
			constraintLogic:  deliveryruleconstraintgroup.ConstraintLogicOr,
			conditionMatches: false,
			expectedDone:     false,
			expectedMatches:  false,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			done, matches := groupEvaluation(tc.constraintLogic, tc.conditionMatches)
			if done != tc.expectedDone || matches != tc.expectedMatches {
				t.Errorf("got (%v, %v), want (%v, %v)", done, matches, tc.expectedDone, tc.expectedMatches)
			}
		})
	}
}
