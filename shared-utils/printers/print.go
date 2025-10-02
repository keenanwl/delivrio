package printers

import (
	"bytes"
	"context"

	"delivrio.io/shared-utils/pulid"
)

type Client struct {
	ctx context.Context
}

type Printer struct {
	ID      pulid.ID `json:"id"`
	Name    string   `json:"name"`
	Active  bool     `json:"active"`
	Network bool     `json:"network"`
}

func NewClient(ctx context.Context) *Client {
	return &Client{ctx: ctx}
}

func IsValidPDF(data []byte) bool {
	// Check if the file starts with the PDF signature
	if len(data) < 5 || !bytes.Equal(data[:5], []byte("%PDF-")) {
		return false
	}

	// These may be too specific
	/*// Check if the file ends with %%EOF
	if len(data) < 5 || !bytes.Equal(bytes.TrimSpace(data[len(data)-5:]), []byte("%%EOF")) {
		return false
	}

	// Check for the presence of key PDF objects
	requiredObjects := []string{"/Pages", "/Type", "/Contents"}
	for _, obj := range requiredObjects {
		if !bytes.Contains(data, []byte(obj)) {
			return false
		}
	}*/

	return true
}
