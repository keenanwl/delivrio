package mergeutils

import (
	"bytes"
	"context"
	"delivrio.io/go/ent"
	"delivrio.io/go/ent/document"
	"fmt"
	"google.golang.org/api/idtoken"
	"google.golang.org/api/option"
	"io"
	"mime/multipart"
	"net/http"
	"net/url"
)

type PDFOptions struct {
	// Inches due to PDF converter
	Width        *float64
	Height       *float64
	MarginTop    *float64
	MarginBottom *float64
	MarginRight  *float64
	MarginLeft   *float64
	BodyInput    *[]byte
	HeaderInput  *[]byte
	FooterInput  *[]byte
	// Default works here
	TurboMode bool
}

func PDFPageSize(v document.PaperSize) func(*PDFOptions) {
	switch v {
	case document.PaperSizeA4:
		h := 11.69
		w := 8.27
		return func(o *PDFOptions) {
			o.Height = &h
			o.Width = &w
		}
	case document.PaperSizeFour_x_six:
		h := 6.0
		w := 4.0
		return func(o *PDFOptions) {
			o.Height = &h
			o.Width = &w
		}
	}

	h := 11.0
	w := 8.5
	return func(o *PDFOptions) {
		o.Height = &h
		o.Width = &w
	}
}

func PDFWidth(width float64) func(*PDFOptions) {
	return func(o *PDFOptions) {
		o.Width = &width
	}
}

func PDFHeight(height float64) func(*PDFOptions) {
	return func(o *PDFOptions) {
		o.Height = &height
	}
}

func PDFMarginTop(marginTop float64) func(*PDFOptions) {
	return func(o *PDFOptions) {
		o.MarginTop = &marginTop
	}
}

func PDFMarginBottom(marginBottom float64) func(*PDFOptions) {
	return func(o *PDFOptions) {
		o.MarginBottom = &marginBottom
	}
}

func PDFMarginRight(marginRight float64) func(*PDFOptions) {
	return func(o *PDFOptions) {
		o.MarginRight = &marginRight
	}
}

func PDFMarginLeft(marginLeft float64) func(*PDFOptions) {
	return func(o *PDFOptions) {
		o.MarginLeft = &marginLeft
	}
}

func PDFTurboMode(turbo bool) func(*PDFOptions) {
	return func(o *PDFOptions) {
		o.TurboMode = turbo
	}
}

func PDFBodyInput(bodyInput []byte) func(*PDFOptions) {
	return func(o *PDFOptions) {
		o.BodyInput = &bodyInput
	}
}

func PDFHeaderInput(headerInput []byte) func(*PDFOptions) {
	return func(o *PDFOptions) {
		o.HeaderInput = &headerInput
	}
}

func PDFFooterInput(footerInput []byte) func(*PDFOptions) {
	return func(o *PDFOptions) {
		o.FooterInput = &footerInput
	}
}

func Doc2PDFOptions(ctx context.Context, doc *ent.Document, mergeVars interface{}) ([]func(*PDFOptions), error) {
	var err error

	opts := make([]func(*PDFOptions), 0)
	opts = append(opts,
		PDFPageSize(doc.PaperSize),
	)

	tmpBody, err := MergeTemplate(doc.HTMLTemplate, mergeVars)
	if err != nil {
		return nil, err
	}
	opts = append(opts, PDFBodyInput(tmpBody.Bytes()))

	if len(doc.HTMLHeader) > 0 {
		tmpHeader, err := MergeTemplate(doc.HTMLHeader, mergeVars)
		if err != nil {
			return nil, err
		}
		opts = append(opts, PDFHeaderInput(tmpHeader.Bytes()))
	}

	if len(doc.HTMLFooter) > 0 {
		tmpFooter, err := MergeTemplate(doc.HTMLFooter, mergeVars)
		if err != nil {
			return nil, err
		}
		opts = append(opts, PDFFooterInput(tmpFooter.Bytes()))
	}

	return opts, nil
}

func HTML2PDF(ctx context.Context, options ...func(*PDFOptions)) ([]byte, error) {
	if conf.HTML2PDF.Gotenberg == nil {
		return nil, fmt.Errorf("html2pdf: missing gotenberg config")
	}

	opt := &PDFOptions{}
	for _, o := range options {
		o(opt)
	}

	// Gets used as Audience
	baseURL := conf.HTML2PDF.Gotenberg.Base

	req, err := multiPartRequest(baseURL, opt)
	if err != nil {
		return nil, err
	}
	/*
		r, _ := httputil.DumpRequest(req, true)
		fmt.Println(string(r))*/

	keyFile := conf.HTML2PDF.Gotenberg.CloudRunAuthJSON

	body, err := gcpAuthRequest(ctx, req, []byte(*keyFile), baseURL)
	if err != nil {
		return nil, err
	}

	return body, nil
}

func multiPartRequest(targetURL string, opt *PDFOptions) (*http.Request, error) {

	// Create a new multipart form
	var buf bytes.Buffer
	mw := multipart.NewWriter(&buf)

	if opt.BodyInput != nil && len(*opt.BodyInput) > 0 {
		err := writeFile(mw, "index.html", *opt.BodyInput)
		if err != nil {
			return nil, fmt.Errorf("writeFile: %w", err)
		}
	}

	if opt.HeaderInput != nil && len(*opt.HeaderInput) > 0 {
		err := writeFile(mw, "header.html", *opt.HeaderInput)
		if err != nil {
			return nil, fmt.Errorf("writeFile: %w", err)
		}
	}

	if opt.FooterInput != nil && len(*opt.FooterInput) > 0 {
		err := writeFile(mw, "footer.html", *opt.FooterInput)
		if err != nil {
			return nil, fmt.Errorf("writeFile: %w", err)
		}
	}

	err := mw.WriteField("skipNetworkIdleEvent", fmt.Sprintf("%t", opt.TurboMode))
	if err != nil {
		return nil, fmt.Errorf("writeField: %w", err)
	}

	if opt.Width != nil {
		err = mw.WriteField("paperWidth", fmt.Sprintf("%v", *opt.Width))
		if err != nil {
			return nil, fmt.Errorf("writeField: %w", err)
		}
	}

	if opt.Height != nil {
		err = mw.WriteField("paperHeight", fmt.Sprintf("%v", *opt.Height))
		if err != nil {
			return nil, fmt.Errorf("writeField: %w", err)
		}
	}

	if opt.MarginTop != nil {
		err = mw.WriteField("marginTop", fmt.Sprintf("%v", *opt.MarginTop))
		if err != nil {
			return nil, fmt.Errorf("writeField: %w", err)
		}
	}
	if opt.MarginBottom != nil {
		err = mw.WriteField("marginBottom", fmt.Sprintf("%v", *opt.MarginBottom))
		if err != nil {
			return nil, fmt.Errorf("writeField: %w", err)
		}
	}
	if opt.MarginLeft != nil {
		err = mw.WriteField("marginLeft", fmt.Sprintf("%v", *opt.MarginLeft))
		if err != nil {
			return nil, fmt.Errorf("writeField: %w", err)
		}
	}
	if opt.MarginRight != nil {
		err = mw.WriteField("marginRight", fmt.Sprintf("%v", *opt.MarginRight))
		if err != nil {
			return nil, fmt.Errorf("writeField: %w", err)
		}
	}

	// Should not be deferred. This writes the closing boundary.
	err = mw.Close()
	if err != nil {
		return nil, fmt.Errorf("mw.Close: %w", err)
	}

	p, err := url.JoinPath(targetURL, `/forms/chromium/convert/html`)
	if err != nil {
		return nil, fmt.Errorf("url.JoinPath: %w", err)
	}
	req, err := http.NewRequest("POST", p, &buf)
	if err != nil {
		return nil, fmt.Errorf("req: %w", err)
	}
	req.Header.Set("Content-Type", mw.FormDataContentType())

	return req, nil
}

func writeFile(mw *multipart.Writer, inputName string, input []byte) error {

	// Add the HTML file to the form
	fw, err := mw.CreateFormFile("files", inputName)
	if err != nil {
		return fmt.Errorf("mw.CreateFormFile: %w", err)
	}
	_, err = fw.Write(input)
	if err != nil {
		return fmt.Errorf("fw.Write: %w", err)
	}

	return nil
}

func gcpAuthRequest(ctx context.Context, req *http.Request, keyfile []byte, audience string) ([]byte, error) {

	// client is a http.Client that automatically adds an "Authorization" header
	// to any requests made.
	client, err := idtoken.NewClient(ctx, audience, option.WithCredentialsJSON(keyfile))
	if err != nil {
		return nil, fmt.Errorf("idtoken.NewClient: %w", err)
	}

	resp, err := client.Do(req)
	if err != nil {
		return nil, fmt.Errorf("client.Get: %w", err)
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, fmt.Errorf("io: %w", err)
	}

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("gcp request: status: %v: %v", resp.StatusCode, string(body))
	}

	return body, nil
}
