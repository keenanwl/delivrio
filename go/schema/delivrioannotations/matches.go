package delivrioannotations

import "entgo.io/ent/schema"

type (
	// Annotation annotates fields and edges with metadata for templates.
	Annotation struct {
		// Check ensures this property is checked in the matching method
		Check bool `json:"Check,omitempty"`
		Clone bool `json:"Clone,omitempty"`
		// Skip is required since bool defaults to false
		Skip      bool `json:"Skip,omitempty"`
		SkipClone bool `json:"SkipClone,omitempty"`
	}
)

func (a Annotation) Merge(other schema.Annotation) schema.Annotation {
	var ant Annotation
	switch other := other.(type) {
	case Annotation:
		ant = other
	case *Annotation:
		if other != nil {
			ant = *other
		}
	default:
		return a
	}

	a.Check = a.Check || ant.Check
	a.Clone = a.Clone || ant.Clone
	a.Skip = a.Skip || ant.Skip
	a.SkipClone = a.SkipClone || ant.SkipClone

	return a
}

func (a Annotation) Name() string {
	return "DelivrioAnnotations"
}

func Check() Annotation {
	return Annotation{Check: true}
}
func Clone() Annotation {
	return Annotation{Clone: true}
}
func Skip() Annotation {
	return Annotation{Skip: true}
}
func SkipClone() Annotation {
	return Annotation{SkipClone: true}
}
