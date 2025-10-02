package deliverypoints

import (
	"delivrio.io/go/ent"
	"time"
)

var UpdateInterval = time.Now().Add(time.Hour * -12)

type DropPointLookupAddress struct {
	Address1 string
	Zip      string
	Country  *ent.Country
}
