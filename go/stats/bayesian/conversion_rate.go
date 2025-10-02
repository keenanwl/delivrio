package bayesian

import (
	"math"
)

type BinaryVariant struct {
	participants int
	conversions  int
}

type BinaryTest struct {
	variants []BinaryVariant
}

func NewBinaryTest() *BinaryTest {
	return &BinaryTest{
		variants: make([]BinaryVariant, 0, 4),
	}
}

func (bt *BinaryTest) Add(participants, conversions int) {
	if conversions > participants {
		panic("Conversions cannot exceed participants")
	}
	if len(bt.variants) >= 4 {
		panic("Cannot add more than 4 variants")
	}

	bt.variants = append(bt.variants, BinaryVariant{participants: participants, conversions: conversions})
}

func (bt *BinaryTest) Probabilities() []float64 {
	numVariants := len(bt.variants)

	switch numVariants {
	case 0:
		return []float64{}
	case 1:
		return []float64{1.0}
	case 2:
		a := &bt.variants[1]
		b := &bt.variants[0]

		prob := probBBWinsOverA(1+a.conversions, 1+a.participants-a.conversions, 1+b.conversions, 1+b.participants-b.conversions)
		return []float64{prob, 1.0 - prob}
	case 3:
		probs := make([]float64, 0, 3)
		total := 0.0
		for i := 0; i < 2; i++ {
			a := &bt.variants[(i+2)%3]
			b := &bt.variants[(i+1)%3]
			c := &bt.variants[i]

			prob := probCBWinsOverAB(1+a.conversions, 1+a.participants-a.conversions, 1+b.conversions, 1+b.participants-b.conversions, 1+c.conversions, 1+c.participants-c.conversions)

			probs = append(probs, prob)
			total += prob
		}
		probs = append(probs, 1.0-total)
		return probs
	default:
		probs := make([]float64, 0, 4)
		total := 0.0
		for i := 0; i < 3; i++ {
			a := &bt.variants[(i+3)%4]
			b := &bt.variants[(i+2)%4]
			c := &bt.variants[(i+1)%4]
			d := &bt.variants[i]

			prob := probDBWinsOverABC(1+a.conversions, 1+a.participants-a.conversions, 1+b.conversions, 1+b.participants-b.conversions, 1+c.conversions, 1+c.participants-c.conversions, 1+d.conversions, 1+d.participants-d.conversions)

			probs = append(probs, prob)
			total += prob
		}
		probs = append(probs, 1.0-total)
		return probs
	}
}

func probBBWinsOverA(alphaA, betaA, alphaB, betaB int) float64 {
	total := 0.0
	logbetaAABa := logbeta(float64(alphaA), float64(betaA))
	betaBa := float64(betaB + betaA)

	for i := int(0); i < alphaB; i++ {
		total += math.Exp(logbeta(float64(alphaA+i), betaBa) - math.Log(float64(betaB+i)) - logbeta(float64(1+i), float64(betaB)) - logbetaAABa)
	}

	return total
}

func probCBWinsOverAB(alphaA, betaA, alphaB, betaB, alphaC, betaC int) float64 {
	total := 0.0
	logbetaAAAC := logbeta(float64(alphaA), float64(betaA+betaC))
	logbetaABBC := logbeta(float64(alphaA+alphaB), float64(betaA+betaB+betaC))
	betaABBC := float64(betaB + betaA + betaC)

	for i := int(0); i < alphaC; i++ {
		total += math.Exp(logbeta(float64(alphaA+i), float64(betaA+betaC)) + logbeta(float64(alphaA+alphaB+i), betaABBC) - logbeta(float64(betaC+i), float64(betaC)) - logbetaAAAC - logbetaABBC)
	}

	return total
}

func probDBWinsOverABC(alphaA, betaA, alphaB, betaB, alphaC, betaC, alphaD, betaD int) float64 {
	total := 0.0
	logbetaAAAD := logbeta(float64(alphaA), float64(betaA+betaD))
	logbetaABBD := logbeta(float64(alphaA+alphaB), float64(betaA+betaB+betaD))
	logbetaACCD := logbeta(float64(alphaA+alphaB+alphaC), float64(betaA+betaB+betaC+betaD))
	betaABCD := float64(betaC + betaB + betaA + betaD)

	for i := int(0); i < alphaD; i++ {
		total += math.Exp(logbeta(float64(alphaA+i), float64(betaA+betaD)) + logbeta(float64(alphaA+alphaB+i), float64(betaA+betaB+betaD)) + logbeta(float64(alphaA+alphaB+alphaC+i), betaABCD) - logbeta(float64(betaD+i), float64(betaD)) - logbetaAAAD - logbetaABBD - logbetaACCD)
	}

	return total
}

func logbeta(a, b float64) float64 {
	lgammaA, _ := math.Lgamma(a)
	lgammaB, _ := math.Lgamma(b)
	lgammaAB, _ := math.Lgamma(a + b)
	return lgammaA + lgammaB - lgammaAB
}
