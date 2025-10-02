package i18n

import (
	"delivrio.io/go/ent/language"
	"log"
)

type i18nKey int

const (
	ReturnsOrderNotFound i18nKey = iota
	ReturnsAuthenticationNotFound
	ReturnsMissingCollis
)

var englishBase = map[i18nKey]string{
	ReturnsOrderNotFound:          "Order could not be found. Please try again.",
	ReturnsAuthenticationNotFound: "Order ID and Email required.",
	ReturnsMissingCollis:          "There were no packages associated with this order.",
}

var translationsDanish = map[i18nKey]string{
	ReturnsOrderNotFound:          "Order could not be found. Please try again.",
	ReturnsAuthenticationNotFound: "Order ID and Email required.",
	ReturnsMissingCollis:          "There were no packages associated with this order.",
}

func Value(lang language.InternalID, key i18nKey) string {

	translations := englishBase
	switch lang {
	case language.InternalIDDA:
		translations = translationsDanish
	default:
		log.Printf("language %v not found", lang)
	}

	if v, ok := translations[key]; ok {
		return v
	} else {
		log.Printf("translations %v->%v not found", lang, key)
	}

	return "Translation not found"
}
