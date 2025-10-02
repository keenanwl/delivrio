package postnordresponse

type ErrorComposite struct {
	Message        string         `json:"message"`
	CompositeFault CompositeFault `json:"compositeFault"`
}

type CompositeFault struct {
	Faults []Fault `json:"faults"`
}

type Fault struct {
	ExplanationText string `json:"explanationText"`
}
