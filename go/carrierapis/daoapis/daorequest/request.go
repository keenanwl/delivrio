package daorequest

type Authentication struct {
	CustomerID string // "kundeid"
	Code       string // "kode"
}

type LabelType string

var (
	LabelTypeLabelless LabelType = "labelless" // Default
	LabelTypeWithLabel LabelType = "withlabel"
)

func (l LabelType) String() string {
	return string(l)
}

type LabelFormat string

func (l LabelFormat) String() string {
	return string(l)
}

var (
	LabelFormatJSON LabelFormat = "JSON" // Default
	LabelFormatXML  LabelFormat = "XML"
	LabelFormatCSV  LabelFormat = "CSV"
)

type LabelPaperSize string

func (l LabelPaperSize) String() string {
	return string(l)
}

var (
	LabelPaperSize100x150    LabelPaperSize = "100x150" // Default
	LabelPaperSize150x100    LabelPaperSize = "150x100"
	LabelPaperSizeA4Foldable LabelPaperSize = "A4Foldable"
	LabelPaperSize100x1501   LabelPaperSize = "100x1501"
	LabelPaperSize100x150p   LabelPaperSize = "100x150p"
)
