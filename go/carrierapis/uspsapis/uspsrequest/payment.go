package uspsrequest

type Payment struct {
	Roles []Role `json:"roles"`
}

type Role struct {
	RoleName      string  `json:"roleName"`
	Crid          string  `json:"CRID"`
	Mid           string  `json:"MID"`
	ManifestMID   string  `json:"manifestMID"`
	AccountType   string  `json:"accountType"`
	AccountNumber string  `json:"accountNumber"`
	PermitNumber  *string `json:"permitNumber,omitempty"`
	PermitZIP     *string `json:"permitZIP,omitempty"`
}
