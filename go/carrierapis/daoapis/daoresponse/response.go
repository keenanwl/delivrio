package daoresponse

import "fmt"

type ResponseStatus string

var (
	StatusOK     ResponseStatus = "OK"
	StatusFailed ResponseStatus = "FEJL"
)

func (l ResponseStatus) String() string {
	return string(l)
}

type BaseResponse struct {
	Status       ResponseStatus `json:"status"`
	ErrorCode    interface{}    `json:"fejlkode"`
	ErrorMessage string         `json:"fejltekst"`
}

func (b *BaseResponse) ErrorCodeParsed() string {
	switch v := b.ErrorCode.(type) {
	case int:
		return fmt.Sprintf("%v", v)
	case string:
		return fmt.Sprintf("%v", v)
	}
	return ""
}

type ShipmentCreateResponse struct {
	BaseResponse
	Result Resultat `json:"resultat"`
}

type Resultat struct {
	Barcode     string `json:"stregkode"`
	LabelTekst1 string `json:"labelTekst1"`
	LabelTekst2 string `json:"labelTekst2"`
	LabelTekst3 string `json:"labelTekst3"`
	Udsortering string `json:"udsortering"`
	Eta         string `json:"ETA"`
}
