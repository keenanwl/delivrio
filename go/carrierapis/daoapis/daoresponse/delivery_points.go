package daoresponse

type ResponseDeliveryPoints struct {
	BaseResponse
	Results DeliveryPointResults `json:"resultat"`
	Antal   string               `json:"antal"`
}

type DeliveryPointResults struct {
	ParcelShops   []ParcelShops `json:"pakkeshops"`
	StartingPoint Udgangspunkt  `json:"udgangspunkt"`
}

type ParcelShops struct {
	ShopId        string `json:"shopId"`
	Type          string `json:"type"`
	Navn          string `json:"navn"`
	Adresse       string `json:"adresse"`
	Postnr        string `json:"postnr"`
	Bynavn        string `json:"bynavn"`
	Udsortering   string `json:"udsortering"`
	Latitude      string `json:"latitude"`
	Longitude     string `json:"longitude"`
	Afstand       string `json:"afstand"`
	Aabningstider struct {
		Man string `json:"man"`
		Tir string `json:"tir"`
		Ons string `json:"ons"`
		Tor string `json:"tor"`
		Fre string `json:"fre"`
		Lor string `json:"lor"`
		Son string `json:"son"`
	} `json:"aabningstider"`
	AfstandDirekte  string `json:"afstand_direkte"`
	AfstandMinutter string `json:"afstand_minutter"`
}
type Udgangspunkt struct {
	Latitude           float64 `json:"latitude"`
	Longtide           float64 `json:"longtide"`
	PositionFromPostal bool    `json:"position_from_postal"`
}
