package bringapis

import (
	"delivrio.io/go/carrierapis/bringapis/bringrequest"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestCombineCustomsLines(t *testing.T) {
	base := bringrequest.EDICustomsDeclaration{
		Quantity:             10,
		GoodsDescription:     "Test Product",
		CustomsArticleNumber: "1234",
		ItemNetWeightInKg:    1.5,
		TarriffLineAmount:    20.50,
		Currency:             "USD",
		CountryOfOrigin:      "US",
	}

	adl := []bringrequest.EDICustomsDeclaration{
		{
			Quantity:             5,
			GoodsDescription:     "Test Product should not display",
			CustomsArticleNumber: "not shown",
			ItemNetWeightInKg:    1.0,
			TarriffLineAmount:    10.00,
			Currency:             "USD",
			CountryOfOrigin:      "USA<__Wrong",
		},
		{
			Quantity:             7,
			GoodsDescription:     "Test Product not displayed since it should be the same anyways",
			CustomsArticleNumber: "not shown2",
			ItemNetWeightInKg:    0.5,
			TarriffLineAmount:    5.00,
			Currency:             "USD",
			CountryOfOrigin:      "USA<-Wrong",
		},
	}

	expected := bringrequest.EDICustomsDeclaration{
		Quantity:             22,
		GoodsDescription:     "Test Product",
		CustomsArticleNumber: "1234",
		ItemNetWeightInKg:    3.0,
		TarriffLineAmount:    35.50,
		Currency:             "USD",
		CountryOfOrigin:      "US",
	}

	result := combineCustomsLines(base, adl...)

	assert.Equal(t, expected, result)
}
