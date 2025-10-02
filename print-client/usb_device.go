package main

import "delivrio.io/shared-utils/pulid"

type USBDevice struct {
	ID        pulid.ID `json:"id"`
	VendorID  string   `json:"vendor_id"`
	ProductID string   `json:"product_id"`
	Name      string   `json:"name"`
	Active    bool     `json:"active"`
}

/*func (d USBDevice) WorkstationName() string {
	return d.name
}
*/
