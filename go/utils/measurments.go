package utils

import (
	"context"
	"delivrio.io/go/ent"
	"math"
)

func CmToInches(cm int) int64 {
	return int64(CmToInchesFloat(cm))
}

func CmToInchesFloat(cm int) float64 {
	return math.Ceil(float64(cm) / cmPerInch)
}

const cmPerInch = 2.54
const gramsPerPound = 453.592

func PoundsFromOrderLines(ctx context.Context, variants []*ent.OrderLine) (float64, error) {
	totalGrams, err := OuncesFromOrderLines(ctx, variants)
	if err != nil {
		return 0.0, err
	}

	return (math.Ceil(totalGrams*100) / 100) / gramsPerPound, nil
}

const gramsPerOunce = 28.3495

func OuncesFromOrderLines(ctx context.Context, orderLines []*ent.OrderLine) (float64, error) {
	var totalGrams float64 = 0
	for _, ol := range orderLines {
		variant, err := ol.QueryProductVariant().
			Only(ctx)
		if err != nil {
			return 0.0, err
		}

		if variant.WeightG != nil {
			totalGrams += float64(*variant.WeightG) * float64(ol.Units)
		}
	}

	return (math.Ceil(totalGrams*100) / 100) / gramsPerOunce, nil
}

func GramsToOunces(grams int) int {
	return int(float64(grams) / gramsPerOunce)
}
