package mergeutils

import (
	"bytes"
	"context"
	"delivrio.io/go/appconfig"
	"delivrio.io/go/ent"
	"delivrio.io/go/ent/carrier"
	"delivrio.io/go/ent/carrierbrand"
	"delivrio.io/go/ent/colli"
	"delivrio.io/go/ent/deliveryoption"
	"delivrio.io/go/ent/document"
	"delivrio.io/go/ent/documentfile"
	"delivrio.io/go/ent/shipment"
	"delivrio.io/go/ent/shipmentpallet"
	"delivrio.io/go/ent/shipmentparcel"
	"delivrio.io/go/schema/fieldjson"
	"delivrio.io/go/utils"
	"delivrio.io/go/viewer"
	"delivrio.io/shared-utils/pulid"
	"encoding/base64"
	"fmt"
	"github.com/boombuler/barcode"
	"github.com/boombuler/barcode/code128"
	"gocloud.dev/blob"
	"html/template"
	"image"
	"image/png"
	"path"
	"regexp"
	"sort"
	"strconv"
	"strings"
	"time"
)

func EmptyPackingSlip() *PackingSlip {
	return &PackingSlip{
		BasicDocVariables: BasicDocVariables{
			DocCreatedDate: "",
			DocCreatedTime: "",
		},
		CustomerAddress: CustomerAddress{
			CustomerFirstName: "",
			CustomerLastName:  "",
			CustomerCompany:   "",
			CustomerEmail:     "",
			CustomerAddress1:  "",
			CustomerAddress2:  "",
			CustomerZip:       "",
			CustomerCity:      "",
			CustomerState:     "",
			CustomerCountry:   "",
		},
		SenderAddress: SenderAddress{
			SenderFirstName: "",
			SenderLastName:  "",
			SenderCompany:   "",
			SenderEmail:     "",
			SenderAddress1:  "",
			SenderAddress2:  "",
			SenderZip:       "",
			SenderCity:      "",
			SenderState:     "",
			SenderCountry:   "",
		},
		DropPointAddress: DropPointAddress{
			DropPointCompany:  "",
			DropPointAddress1: "",
			DropPointAddress2: "",
			DropPointZip:      "",
			DropPointCity:     "",
			DropPointState:    "",
			DropPointCountry:  "",
		},
		DeliveryOptionName:    "",
		DeliveryOptionCarrier: "",
		OrderCommentExternal:  "",
		OrderCommentInternal:  "",
		OrderPublicID:         "",
		OrderLines:            make([]OrderLine, 0),
		OrderNoteAttributes:   make(fieldjson.NoteAttributes),
		DELIVRIOBarcodeImgTag: "",
		DELIVRIOBarcodeImgSrc: "",
		DELIVRIOBarcode:       "",
	}
}

func PackingSlipMergeData(ctx context.Context, colliID pulid.ID) (*PackingSlip, error) {
	cli := ent.FromContext(ctx)

	col, err := cli.Colli.Query().
		WithOrder().
		WithDeliveryOption().
		Where(colli.ID(colliID)).
		Only(ctx)
	if err != nil {
		return nil, err
	}

	customerAddress, err := col.QueryRecipient().
		Only(ctx)
	if err != nil {
		return nil, err
	}

	cAdr, err := addressToCustomerAddress(ctx, customerAddress)
	if err != nil {
		return nil, err
	}

	senderAddress, err := col.QuerySender().
		Only(ctx)
	if err != nil {
		return nil, err
	}

	sAdr, err := addressToSenderAddress(ctx, senderAddress)
	if err != nil {
		return nil, err
	}

	lines, err := col.QueryOrderLines().
		All(ctx)
	if err != nil {
		return nil, err
	}

	mergeOrderLines, err := orderLinesToMerge(ctx, lines)
	if err != nil {
		return nil, err
	}

	deliveryOptionName := ""
	deliveryOptionCarrierName := ""
	deliveryOption, err := col.QueryDeliveryOption().
		WithCarrier().
		Only(ctx)
	if err != nil && !ent.IsNotFound(err) {
		return nil, err
	} else if err == nil {
		doCB, err := deliveryOption.QueryCarrier().
			QueryCarrierBrand().
			Only(ctx)
		if err != nil {
			return nil, err
		}
		deliveryOptionName = deliveryOption.Name
		deliveryOptionCarrierName = doCB.Label
	}

	source, err := code128.Encode(strconv.FormatInt(col.InternalBarcode, 10))
	if err != nil {
		return nil, fmt.Errorf("encode: %w", err)
	}
	scaledCode, err := barcode.Scale(source, 250, 65)
	if err != nil {
		return nil, fmt.Errorf("scale: %w", err)
	}

	b64Barcode, err := imageToBase64(scaledCode)
	if err != nil {
		return nil, err
	}

	return &PackingSlip{
		BasicDocVariables: BasicDocVariables{
			DocCreatedDate: time.Now().Format(time.DateOnly),
			DocCreatedTime: time.Now().Format(time.Kitchen),
		},
		DeliveryOptionName:    deliveryOptionName,
		DeliveryOptionCarrier: deliveryOptionCarrierName,
		CustomerAddress:       *cAdr,
		SenderAddress:         *sAdr,
		OrderPublicID:         col.Edges.Order.OrderPublicID,
		OrderCommentExternal:  col.Edges.Order.CommentExternal,
		OrderCommentInternal:  col.Edges.Order.CommentInternal,
		OrderNoteAttributes:   col.Edges.Order.NoteAttributes,
		OrderLines:            mergeOrderLines,
		DELIVRIOBarcodeImgTag: template.HTML(fmt.Sprintf("<img src='data:image/png;base64,%v' />", b64Barcode)),
		DELIVRIOBarcodeImgSrc: fmt.Sprintf("data:image/png;base64,%v", b64Barcode),
		DELIVRIOBarcode:       strconv.FormatInt(col.InternalBarcode, 10),
	}, nil

}

func imageToBase64(img image.Image) (string, error) {
	buf := new(bytes.Buffer)
	err := png.Encode(buf, img)
	if err != nil {
		return "", err
	}
	return base64.StdEncoding.EncodeToString(buf.Bytes()), nil
}

func ValidateMergeVariables(ctx context.Context, input ent.UpdateDocumentInput) error {
	// Validate input and discard
	switch *input.MergeType {
	case document.MergeTypeOrders:
		_, err := MergeTemplate(*input.HTMLTemplate, NewOrdersListTest())
		if err != nil {
			return err
		}
		_, err = MergeTemplate(*input.HTMLHeader, NewOrdersListTest())
		if err != nil {
			return err
		}
		_, err = MergeTemplate(*input.HTMLFooter, NewOrdersListTest())
		if err != nil {
			return err
		}
	case document.MergeTypePackingSlip:
		_, err := MergeTemplate(*input.HTMLTemplate, NewPackingSlipTest(ctx))
		if err != nil {
			return err
		}
		_, err = MergeTemplate(*input.HTMLHeader, NewPackingSlipTest(ctx))
		if err != nil {
			return err
		}
		_, err = MergeTemplate(*input.HTMLFooter, NewPackingSlipTest(ctx))
		if err != nil {
			return err
		}
	}
	return nil
}

func NewOrdersListTest() OrdersList {
	return OrdersList{
		BasicDocVariables: BasicDocVariables{
			DocCreatedDate: "2022-01-23",
			DocCreatedTime: "10:48",
		},
		CarrierName:   "PostNord",
		RangeFromDate: "2020-01-01 10:00",
		RangeToDate:   "2021-01-01 11:00",
		ShipmentCount: "49",
		Orders:        testOrderRow(),
	}
}

func NewPackingSlipTest(ctx context.Context) PackingSlip {
	return PackingSlip{
		BasicDocVariables: BasicDocVariables{
			DocCreatedDate: "2022-01-23",
			DocCreatedTime: "10:48",
		},
		CustomerAddress:       testCustomerAddress(),
		SenderAddress:         testSenderAddress(),
		DeliveryOptionName:    "Pakkeshop",
		DeliveryOptionCarrier: "PostNord",
		OrderPublicID:         "#000099999",
		OrderCommentExternal:  "Comment external..",
		OrderCommentInternal:  "Comment internal..",
		OrderNoteAttributes: map[string]string{
			"date-of-pickup": "09/08/2024",
			"time-of-pickup": "13:00",
		},
		OrderLines:            testOrderLines(ctx),
		DELIVRIOBarcodeImgTag: "",
		DELIVRIOBarcodeImgSrc: "",
		DELIVRIOBarcode:       "1000001",
	}
}

func FetchOrdersList(ctx context.Context, carrierBrand *pulid.ID, start time.Time, end time.Time) (*OrdersList, error) {
	cli := ent.FromContext(ctx)

	brandFilter := shipmentparcel.Or()
	if carrierBrand != nil {
		brandFilter = shipmentparcel.HasColliWith(colli.HasDeliveryOptionWith(deliveryoption.HasCarrierWith(carrier.HasCarrierBrandWith(carrierbrand.ID(*carrierBrand)))))
	}

	par, err := cli.ShipmentParcel.Query().
		WithShipment().
		WithColli().
		Where(
			brandFilter,
			shipmentparcel.StatusIn(shipmentparcel.StatusPrinted),
			shipmentparcel.HasShipmentWith(shipment.And(
				shipment.CreatedAtLTE(end),
				shipment.CreatedAtGTE(start),
			)),
		).All(ctx)
	if err != nil {
		return nil, err
	}

	pal, err := cli.ShipmentPallet.Query().
		WithShipment().
		Where(
			shipmentpallet.StatusIn(shipmentpallet.StatusPrinted),
			shipmentpallet.HasShipmentWith(shipment.And(
				shipment.CreatedAtLTE(end),
				shipment.CreatedAtGTE(start),
			)),
		).All(ctx)
	if err != nil {
		return nil, err
	}

	cbOutput := ""
	if carrierBrand != nil {
		cb, err := cli.CarrierBrand.Query().
			Where(carrierbrand.ID(*carrierBrand)).
			Only(ctx)
		if err != nil {
			return nil, err
		}
		cbOutput = cb.Label
	}

	output := &OrdersList{
		CarrierName: cbOutput,
		BasicDocVariables: BasicDocVariables{
			DocCreatedDate: time.Now().Format(time.DateOnly),
			DocCreatedTime: time.Now().Format(time.Kitchen),
		},
		RangeFromDate: fmt.Sprintf("%v %v", start.Format(time.DateOnly), start.Format(time.TimeOnly)),
		RangeToDate:   fmt.Sprintf("%v %v", end.Format(time.DateOnly), end.Format(time.TimeOnly)),
		ShipmentCount: "",
		Orders:        make([]OrderRow, 0),
	}

	c := 0
	for c < 50 {
		c++
		output.Orders = append(output.Orders, demoLine(c))
	}

	for _, p := range pal {

		output.Orders = append(output.Orders, OrderRow{
			RowCount: "",
			OrderID:  "",
			ShipmentCreatedDate: fmt.Sprintf(
				"%v %v",
				p.Edges.Shipment.CreatedAt.Format(time.DateOnly),
				p.Edges.Shipment.CreatedAt.Format(time.TimeOnly),
			),
			ShipmentID:         p.Edges.Shipment.ShipmentPublicID,
			ShipmentTrackingID: p.Barcode,
			created:            p.Edges.Shipment.CreatedAt,
		})
	}

	for _, p := range par {
		ord, err := p.Edges.Colli.QueryOrder().
			Only(ctx)
		if err != nil {
			return nil, err
		}

		output.Orders = append(output.Orders, OrderRow{
			RowCount: "",
			OrderID:  ord.OrderPublicID,
			ShipmentCreatedDate: fmt.Sprintf(
				"%v %v",
				p.Edges.Shipment.CreatedAt.Format(time.DateOnly),
				p.Edges.Shipment.CreatedAt.Format(time.TimeOnly),
			),
			ShipmentID:         p.Edges.Shipment.ShipmentPublicID,
			ShipmentTrackingID: p.ItemID,
			created:            p.Edges.Shipment.CreatedAt,
		})
	}

	sort.Slice(output.Orders, func(i, j int) bool {
		return output.Orders[i].created.Before(output.Orders[j].created)
	})

	for i, _ := range output.Orders {
		output.Orders[i].RowCount = fmt.Sprintf("%v", i+1)
	}

	output.ShipmentCount = fmt.Sprintf("%v", len(output.Orders))

	return output, nil
}

func demoLine(count int) OrderRow {
	return OrderRow{
		RowCount: "",
		OrderID:  fmt.Sprintf("#10000000%v", count),
		ShipmentCreatedDate: fmt.Sprintf(
			"%v %v",
			time.Now().Format(time.DateOnly),
			time.Now().Format(time.TimeOnly),
		),
		ShipmentID:         fmt.Sprintf("shipment-id-%v", count),
		ShipmentTrackingID: fmt.Sprintf("999999999999999999%v", count),
		created:            time.Now(),
	}
}

func FetchShipmentDocumentPDF(ctx context.Context, shipmentParcelID pulid.ID) (string, error) {
	cli := ent.FromContext(ctx)

	df, err := cli.ShipmentParcel.Query().
		Where(shipmentparcel.IDEQ(shipmentParcelID)).
		QueryDocumentFile().
		Only(ctx)
	if err != nil {
		return "", nil
	}

	switch df.StorageType {
	case documentfile.StorageTypeBucket:
		out, err := PDFDocFromFile(ctx, []*ent.DocumentFile{df})
		if err != nil {
			return "", err
		} else if len(out) == 0 {
			return "", fmt.Errorf("no document found")
		}
		return out[0], nil
	}

	return "", nil
}

// PDFDocFromFile abstracts handling fetching PDF from the bucket
func PDFDocFromFile(ctx context.Context, docs []*ent.DocumentFile) ([]string, error) {

	output := make([]string, 0)

	for _, doc := range docs {
		if doc.StorageType == documentfile.StorageTypeDatabase {
			output = append(output, doc.DataPdfBase64)
			continue
		}

		cli := ent.FromContext(ctx)
		bucketDoc, err := cli.Bucket.ReadAll(ctx, doc.StoragePath)
		if err != nil {
			return nil, err
		}

		output = append(output, string(bucketDoc))
	}

	return output, nil
}

func QueryOrFetchPackingSlipZPL(ctx context.Context, c *ent.Colli) ([]string, error) {
	pdfPackingSlip, err := QueryOrFetchPackingSlip(ctx, c)
	if err != nil {
		return nil, err
	}

	allZpl := make([]string, 0)
	for _, p := range pdfPackingSlip {
		zpl, err := utils.Base64PDFToZPL(p)
		if err != nil {
			return nil, err
		}
		allZpl = append(allZpl, zpl)
	}

	return allZpl, nil
}

func QueryOrFetchPackingSlipPng(ctx context.Context, c *ent.Colli) ([]string, error) {
	pdfPackingSlip, err := QueryOrFetchPackingSlip(ctx, c)
	if err != nil {
		return nil, err
	}

	allPngBase64 := make([]string, 0)
	for _, p := range pdfPackingSlip {
		pngImg, err := utils.Base64PDFToPNG(p)
		if err != nil {
			return nil, err
		}
		var buf bytes.Buffer

		// Encode the image to PNG format and write to buffer
		err = png.Encode(&buf, *pngImg)
		if err != nil {
			return nil, err
		}

		// Convert buffer to byte slice
		imgBytes := buf.Bytes()

		allPngBase64 = append(allPngBase64, base64.StdEncoding.EncodeToString(imgBytes))
	}

	return allPngBase64, nil
}

// QueryOrFetchPackingSlip looks at local cache or generates PDF from scratch
// May not be run in TX due to external network calls
// Creates it's own TX for replacing existing cache
func QueryOrFetchPackingSlip(ctx context.Context, c *ent.Colli) ([]string, error) {
	cli := ent.FromContext(ctx)

	pdfDocs, err := c.QueryDocumentFile().
		Where(documentfile.DocTypeEQ(documentfile.DocTypePackingSlip)).
		All(ctx)
	if err != nil && !ent.IsNotFound(err) {
		return nil, err
	} else if len(pdfDocs) > 0 {
		docOutput, err := PDFDocFromFile(ctx, pdfDocs)
		if err != nil {
			return nil, err
		}
		return docOutput, nil
	}

	ord, err := c.QueryOrder().
		Only(ctx)
	if err != nil {
		return nil, err
	}

	// Probably the same for every colli, but may be different
	doc, err := ord.QueryConnection().
		QueryPackingSlipTemplate().
		Only(ctx)
	if err != nil {
		return nil, err
	}

	mergeVars, err := PackingSlipMergeData(ctx, c.ID)
	if err != nil {
		return nil, err
	}

	opts, err := Doc2PDFOptions(ctx, doc, mergeVars)
	if err != nil {
		return nil, err
	}
	opts = append(opts,
		PDFMarginBottom(1),
		PDFTurboMode(true),
	)

	output, err := HTML2PDF(
		ctx,
		opts...,
	)
	if err != nil {
		return nil, err
	}

	b64Docs, err := utils.SplitPDFPagesToB64(output)
	if err != nil {
		return nil, err
	}

	if conf.BlobStorage.Type != appconfig.BlobStorageTypeDatabase {
		err = ReplacePackingSlipBlob(ctx, c.ID, b64Docs)
		if err != nil {
			return nil, err
		}
	} else {
		txCtx, tx, err := cli.OpenTx(ctx)
		if err != nil {
			return nil, err
		}
		defer tx.Rollback()

		err = ReplacePackingSlip(txCtx, c.ID, b64Docs)
		if err != nil {
			return nil, utils.Rollback(tx, err)
		}

		err = tx.Commit()
		if err != nil {
			return nil, utils.Rollback(tx, err)
		}
	}

	return b64Docs, nil
}

// Function to sanitize the path components
func sanitizePathComponent(component string) string {
	// Replace all invalid characters with an underscore
	re := regexp.MustCompile(`[^a-zA-Z0-9._/-]`)
	return re.ReplaceAllString(component, "_")
}

// Skipped the tx for now since there are a couple network requests mixed in
func ReplacePackingSlipBlob(ctx context.Context, colliID pulid.ID, b64PDF []string) error {
	cli := ent.FromContext(ctx)
	view := viewer.FromContext(ctx)

	allKeys := make([]string, 0)
	for _, doc := range b64PDF {
		saveKeyID := pulid.MustNew("")
		saveKey := fmt.Sprintf("%s.pdf", saveKeyID.String())
		saveKey = path.Join(sanitizePathComponent(conf.ServerID), saveKey)
		err := cli.Bucket.Upload(
			ctx,
			saveKey,
			strings.NewReader(doc),
			&blob.WriterOptions{ContentType: "application/pdf"},
		)
		if err != nil {
			return err
		}
		allKeys = append(allKeys, saveKey)
	}

	allNewDocs := make([]pulid.ID, 0)
	for _, key := range allKeys {
		next, err := cli.DocumentFile.Create().
			SetTenantID(view.TenantID()).
			SetDocType(documentfile.DocTypePackingSlip).
			SetStorageType(documentfile.StorageTypeBucket).
			SetStoragePath(key).
			SetColliID(colliID).
			Save(ctx)
		if err != nil {
			return err
		}
		allNewDocs = append(allNewDocs, next.ID)
	}

	oldDocs, err := cli.DocumentFile.Query().
		Where(documentfile.And(
			documentfile.IDNotIn(allNewDocs...),
			documentfile.HasColliWith(colli.ID(colliID))),
			// Remove all packing slips as they should differ only by format
			documentfile.DocTypeIn(documentfile.DocTypePackingSlip),
		).All(ctx)
	if err != nil {
		return fmt.Errorf("replace packing slip: fetching existing: %w", err)
	}

	for _, file := range oldDocs {
		err = cli.Bucket.Delete(ctx, file.StoragePath)
		if err != nil {
			return fmt.Errorf("replace packing slip: deleting existing blob: %w", err)
		}

		_, err = cli.DocumentFile.Delete().
			Where(documentfile.ID(file.ID)).
			Exec(ctx)
		if err != nil {
			return fmt.Errorf("replace packing slip: deleting existing: %w", err)
		}
	}

	return nil
}

func ReplacePackingSlip(ctx context.Context, colliID pulid.ID, b64PDFDocs []string) error {
	tx := ent.TxFromContext(ctx)
	view := viewer.FromContext(ctx)

	_, err := tx.DocumentFile.Delete().
		Where(documentfile.And(
			documentfile.HasColliWith(colli.ID(colliID))),
			// Remove all packing slips as they should differ only by format
			documentfile.DocTypeIn(documentfile.DocTypePackingSlip),
		).Exec(ctx)
	if err != nil {
		return fmt.Errorf("replace packing slip: fetching existing: %w", err)
	}

	for _, doc := range b64PDFDocs {
		err = tx.DocumentFile.Create().
			SetTenantID(view.TenantID()).
			SetDocType(documentfile.DocTypePackingSlip).
			SetDataPdfBase64(doc).
			SetStorageType(documentfile.StorageTypeDatabase).
			SetColliID(colliID).
			Exec(ctx)
		if err != nil {
			return err
		}
	}

	return nil
}
