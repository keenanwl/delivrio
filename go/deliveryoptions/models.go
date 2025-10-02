package deliveryoptions

import (
	"fmt"
	"io"
	"strconv"

	"delivrio.io/go/ent"
	"delivrio.io/shared-utils/pulid"
)

type DeliveryOptionSeedInput struct {
	// For connection info
	ConnectionID pulid.ID                          `json:"connectionID"`
	Country      pulid.ID                          `json:"country"`
	Zip          string                            `json:"zip"`
	ProductLines []*DeliveryOptionProductLineInput `json:"productLines"`
}

type DeliveryOptionProductLineInput struct {
	ProductVariantID pulid.ID `json:"productVariantID"`
	Units            int      `json:"units"`
	UnitPrice        float64  `json:"unitPrice"`
}

type DeliveryOptionBrandName struct {
	DeliveryOptionID      pulid.ID                      `json:"deliveryOptionID"`
	Name                  string                        `json:"name"`
	Description           string                        `json:"description"`
	Status                DeliveryOptionBrandNameStatus `json:"status"`
	Price                 string                        `json:"price"`
	Currency              *ent.Currency                 `json:"currency"`
	Warning               *string                       `json:"warning"`
	RequiresDeliveryPoint bool                          `json:"requiresDeliveryPoint"`
	DeliveryPoint         bool                          `json:"deliveryPoint"`
	ClickAndCollect       bool                          `json:"clickAndCollect"`
}

type DeliveryOptionBrandNameStatus string

const (
	DeliveryOptionBrandNameStatusSelected     DeliveryOptionBrandNameStatus = "SELECTED"
	DeliveryOptionBrandNameStatusAvailable    DeliveryOptionBrandNameStatus = "AVAILABLE"
	DeliveryOptionBrandNameStatusNotAvailable DeliveryOptionBrandNameStatus = "NOT_AVAILABLE"
)

var AllDeliveryOptionBrandNameStatus = []DeliveryOptionBrandNameStatus{
	DeliveryOptionBrandNameStatusSelected,
	DeliveryOptionBrandNameStatusAvailable,
	DeliveryOptionBrandNameStatusNotAvailable,
}

func (e DeliveryOptionBrandNameStatus) IsValid() bool {
	switch e {
	case DeliveryOptionBrandNameStatusSelected, DeliveryOptionBrandNameStatusAvailable, DeliveryOptionBrandNameStatusNotAvailable:
		return true
	}
	return false
}

func (e DeliveryOptionBrandNameStatus) String() string {
	return string(e)
}

// TODO: fix reciever inconsistency
func (e *DeliveryOptionBrandNameStatus) UnmarshalGQL(v interface{}) error {
	str, ok := v.(string)
	if !ok {
		return fmt.Errorf("enums must be strings")
	}

	*e = DeliveryOptionBrandNameStatus(str)
	if !e.IsValid() {
		return fmt.Errorf("%s is not a valid DeliveryOptionBrandNameStatus", str)
	}
	return nil
}

func (e DeliveryOptionBrandNameStatus) MarshalGQL(w io.Writer) {
	fmt.Fprint(w, strconv.Quote(e.String()))
}
