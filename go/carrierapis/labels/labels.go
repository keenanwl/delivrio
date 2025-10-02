package labels

import (
	"context"
	"delivrio.io/go/carrierapis/bringapis"
	"delivrio.io/go/carrierapis/daoapis"
	"delivrio.io/go/carrierapis/dfapis"
	"delivrio.io/go/carrierapis/easypostapis"
	"delivrio.io/go/carrierapis/glsapis"
	"delivrio.io/go/carrierapis/postnordapis"
	"delivrio.io/go/carrierapis/uspsapis"
	"delivrio.io/go/ent"
	"delivrio.io/go/ent/carrierbrand"
	"delivrio.io/go/ent/colli"
	"delivrio.io/go/ent/printjob"
	"delivrio.io/go/ent/shipment"
	"delivrio.io/go/ent/shipmentparcel"
	"delivrio.io/go/mergeutils"
	"delivrio.io/go/utils"
	"delivrio.io/go/viewer"
	"delivrio.io/shared-utils/pulid"
	"encoding/base64"
	"fmt"
	"go.opentelemetry.io/otel"
	"go.opentelemetry.io/otel/attribute"
)

var (
	tracer = otel.Tracer("labels")
)

// We only support DF by Consolidation right now
type CollisByDeliveryOption struct {
	bring    map[pulid.ID][]*ent.Colli
	dao      map[pulid.ID][]*ent.Colli
	easyPost map[pulid.ID][]*ent.Colli
	postNord map[pulid.ID][]*ent.Colli
	gls      map[pulid.ID][]*ent.Colli
	usps     map[pulid.ID][]*ent.Colli
}

func SortCollis(ctx context.Context, colliIDs []pulid.ID) (*CollisByDeliveryOption, error) {
	ctx, span := tracer.Start(ctx, "SortCollis")
	defer span.End()

	span.SetAttributes(
		attribute.Int("colliCount", len(colliIDs)),
	)

	cli := ent.FromContext(ctx)
	packages, err := cli.Colli.Query().
		WithDeliveryOption().
		WithRecipient().
		// TODO: remember override?
		WithSender().
		WithParcelShop().
		WithOrder().
		Where(colli.IDIn(colliIDs...)).
		All(ctx)
	if err != nil {
		return nil, err
	}

	output := &CollisByDeliveryOption{
		bring:    make(map[pulid.ID][]*ent.Colli),
		dao:      make(map[pulid.ID][]*ent.Colli),
		easyPost: make(map[pulid.ID][]*ent.Colli),
		postNord: make(map[pulid.ID][]*ent.Colli),
		gls:      make(map[pulid.ID][]*ent.Colli),
		usps:     make(map[pulid.ID][]*ent.Colli),
	}

	for pi, p := range packages {

		items, err := p.QueryOrderLines().
			Count(ctx)
		if err != nil {
			return nil, err
		}

		// TODO: convert check to hook
		if items <= 0 {
			return nil,
				fmt.Errorf("Empty shipments are not allowed. Add at least one product to order %s parcel %v.",
					p.Edges.Order.ID,
					pi+1,
				)
		}

		deliveryOption, err := p.QueryDeliveryOption().
			Only(ctx)
		if ent.IsNotFound(err) {
			return nil, fmt.Errorf("An order must be assigned a delivery option before a shipment can be created.")
		} else if err != nil {
			return nil, err
		}

		brand, err := deliveryOption.
			QueryCarrier().
			QueryCarrierBrand().
			Only(ctx)
		if err != nil {
			return nil, err
		}

		do := p.Edges.DeliveryOption

		switch brand.InternalID {

		case carrierbrand.InternalIDBring:
			if _, ok := output.bring[do.ID]; !ok {
				output.bring[do.ID] = make([]*ent.Colli, 0)
			}
			output.bring[do.ID] = append(output.bring[do.ID], p)
		case carrierbrand.InternalIDDAO:
			if _, ok := output.dao[do.ID]; !ok {
				output.dao[do.ID] = make([]*ent.Colli, 0)
			}
			output.dao[do.ID] = append(output.dao[do.ID], p)
		case carrierbrand.InternalIDEasyPost:
			if _, ok := output.dao[do.ID]; !ok {
				output.easyPost[do.ID] = make([]*ent.Colli, 0)
			}
			output.easyPost[do.ID] = append(output.easyPost[do.ID], p)
		case carrierbrand.InternalIDGLS:
			if _, ok := output.gls[do.ID]; !ok {
				output.gls[do.ID] = make([]*ent.Colli, 0)
			}
			output.gls[do.ID] = append(output.gls[do.ID], p)
		case carrierbrand.InternalIDPostNord:
			if _, ok := output.postNord[do.ID]; !ok {
				output.postNord[do.ID] = make([]*ent.Colli, 0)
			}
			output.postNord[do.ID] = append(output.postNord[do.ID], p)
		case carrierbrand.InternalIDUSPS:
			if _, ok := output.gls[do.ID]; !ok {
				output.usps[do.ID] = make([]*ent.Colli, 0)
			}
			output.usps[do.ID] = append(output.gls[do.ID], p)

		default:
			return nil, fmt.Errorf("%s not implemented", brand.Label)
		}

	}
	return output, nil
}

// Only DF for now
func RequestAndSaveConsolidation(ctx context.Context, prebook bool, c *ent.Consolidation) (*SavedShipment, error) {
	ctx, span := tracer.Start(ctx, "RequestAndSaveConsolidation")
	defer span.End()

	output, err := dfapis.FetchLabels(ctx, prebook, c)
	if err != nil {
		return nil, err
	}

	if output.IsNoop {
		return shipmentFromConsolidation(ctx, c)
	}

	// TODO: potentially saves a lot of labels in a single TX
	// Not good for scale
	created, err := dfapis.SaveLabel(ctx, output)
	if err != nil {
		return nil, err
	}

	allLabels, err := utils.JoinPDFs(created.Labels...)
	if err != nil {
		return nil, err
	}

	// A bit convoluted to match the structure of the others
	return &SavedShipment{
		ShipmentID: created.Shipment,
		LabelsPDF:  created.Labels,
		AllLabels:  allLabels,
	}, nil
}

func shipmentFromConsolidation(ctx context.Context, c *ent.Consolidation) (*SavedShipment, error) {
	ship, err := c.QueryShipment().
		Only(ctx)
	if err != nil && !ent.IsNotFound(err) {
		return nil, err
	} else if ent.IsNotFound(err) {
		return nil, fmt.Errorf("df: this consolidation does not have a shipment to retrieve")
	}

	allParcels, err := ship.QueryShipmentParcel().
		All(ctx)
	if err != nil {
		return nil, err
	}

	allPallets, err := ship.QueryShipmentPallet().
		All(ctx)
	if err != nil {
		return nil, err
	}

	labelPages := make([]string, 0)
	for _, p := range allParcels {
		labelPDF, err := mergeutils.FetchShipmentDocumentPDF(ctx, p.ID)
		if err != nil {
			return nil, err
		}
		labelPages = append(labelPages, labelPDF)
	}

	for _, p := range allPallets {
		labelPages = append(labelPages, p.LabelPdf)
	}

	joinedLabels, err := utils.JoinPDFs(labelPages...)
	if err != nil {
		return nil, err
	}

	return &SavedShipment{
		ShipmentID: ship.ID,
		LabelsPDF:  labelPages,
		AllLabels:  joinedLabels,
	}, nil
}

type SavedShipment struct {
	ShipmentID       pulid.ID `json:"shipmentID"`
	ShipmentParcelID pulid.ID `json:"shipment_parcel_id"`
	LabelsPDF        []string `json:"labelsPDF"`
	AllLabels        string   `json:"allLabels"`
}

func RequestAndSave(ctx context.Context, sorted *CollisByDeliveryOption) ([]SavedShipment, error) {

	ctx, span := tracer.Start(ctx, "RequestAndSave")
	defer span.End()

	span.SetAttributes(
		attribute.Int("bringCount", len(sorted.bring)),
		attribute.Int("daoCount", len(sorted.dao)),
		attribute.Int("daoCount", len(sorted.easyPost)),
		attribute.Int("postNordCount", len(sorted.postNord)),
		attribute.Int("GLSCount", len(sorted.gls)),
		attribute.Int("USPSCount", len(sorted.usps)),
	)

	cli := ent.FromContext(ctx)
	allCreateShipments := make([]SavedShipment, 0)

	for _, g := range sorted.bring {
		fetchedResponses, err := bringapis.FetchLabels(ctx, g)
		if err != nil {
			return nil, err
		}

		for _, r := range fetchedResponses {

			ship, err := bringapis.SaveLabelData(ctx, r)
			if err != nil {
				return nil, err
			}

			allCreateShipments = append(allCreateShipments, SavedShipment{
				ShipmentID: ship.Shipment,
				LabelsPDF:  ship.Labels,
				AllLabels:  r.ResponseB64PDF,
			})

		}
	}

	for _, g := range sorted.dao {
		fetchedResponses, err := daoapis.FetchLabels(ctx, g)
		if err != nil {
			return nil, err
		}

		for _, r := range fetchedResponses {

			ship, err := daoapis.SaveLabelData(ctx, r)
			if err != nil {
				return nil, err
			}

			allCreateShipments = append(allCreateShipments, SavedShipment{
				ShipmentID: ship.Shipment,
				LabelsPDF:  ship.Labels,
				AllLabels:  r.ResponseB64PDF,
			})

		}
	}

	for doID, c := range sorted.easyPost {
		fetchedResponses, err := easypostapis.FetchLabels(ctx, doID, c)
		if err != nil {
			return nil, err
		}

		for _, r := range fetchedResponses {

			ctxTX, tx, err := cli.OpenTx(ctx)
			if err != nil {
				return nil, err
			}
			defer tx.Rollback()

			ship, err := easypostapis.SaveLabelData(ctxTX, r)
			if err != nil {
				return nil, err
			}

			err = tx.Commit()
			if err != nil {
				return nil, err
			}

			allCreateShipments = append(allCreateShipments, SavedShipment{
				ShipmentID: ship.Shipment,
				LabelsPDF:  ship.Labels,
				AllLabels:  r.Responseb64PDF,
			})

		}
	}

	for doID, pn := range sorted.postNord {
		responseData, err := postnordapis.FetchLabels(ctx, doID, pn)
		if err != nil {
			return nil, err
		}

		createShipments, err := postnordapis.SaveFlattenedResponse(ctx, responseData)
		if err != nil {
			return nil, err
		}

		for _, s := range createShipments {
			labelDownload, err := utils.JoinPDFs(s.Labels...)
			if err != nil {
				return nil, err
			}
			allCreateShipments = append(allCreateShipments, SavedShipment{
				ShipmentID: s.Shipment,
				LabelsPDF:  s.Labels,
				AllLabels:  labelDownload,
			})
		}
	}

	for doID, g := range sorted.gls {
		fetchedResponses, err := glsapis.FetchLabels(ctx, doID, g)
		if err != nil {
			return nil, err
		}

		for _, r := range fetchedResponses {

			ctxTX, tx, err := cli.OpenTx(ctx)
			if err != nil {
				return nil, err
			}
			defer tx.Rollback()

			ship, err := glsapis.SaveLabelData(ctxTX, r.ShipmentConfig, r)
			if err != nil {
				return nil, err
			}

			err = tx.Commit()
			if err != nil {
				return nil, err
			}

			for _, s := range ship {
				allCreateShipments = append(allCreateShipments, SavedShipment{
					ShipmentID: s.Shipment,
					LabelsPDF:  s.Labels,
					AllLabels:  r.Response.PDF,
				})
			}
		}
	}

	for doID, u := range sorted.usps {
		fetchedResponses, err := uspsapis.FetchLabels(ctx, doID, u)
		if err != nil {
			return nil, err
		}

		for _, r := range fetchedResponses {

			ctxTX, tx, err := cli.OpenTx(ctx)
			if err != nil {
				return nil, err
			}
			defer tx.Rollback()

			ship, err := uspsapis.SaveLabelData(ctxTX, r)
			if err != nil {
				return nil, err
			}

			err = tx.Commit()
			if err != nil {
				return nil, fmt.Errorf("requestAndSave: %w", err)
			}

			allCreateShipments = append(allCreateShipments, SavedShipment{
				ShipmentID: ship.Shipment,
				LabelsPDF:  ship.Labels,
				AllLabels:  r.Responseb64PDF,
			})

		}
	}

	return allCreateShipments, nil

}

// MustShipmentFromColliID ensures a shipment is retrieved or created for the given colliID
// expects all existing shipments to be PDF for now.
func MustShipmentFromColliID(ctx context.Context, colliID pulid.ID) ([]SavedShipment, error) {
	cli := ent.FromContext(ctx)
	existingShip, err := cli.Colli.Query().
		Where(colli.ID(colliID)).
		QueryShipmentParcel().
		WithShipment().
		Where(
			shipmentparcel.HasShipmentWith(
				shipment.StatusNEQ(shipment.StatusDeleted),
			),
		).
		Only(ctx)
	if err != nil && !ent.IsNotFound(err) {
		return nil, fmt.Errorf("must shipment: %w", err)
	} else if ent.IsNotFound(err) {
		sorted, err := SortCollis(ctx, []pulid.ID{colliID})
		if err != nil {
			return nil, err
		}
		// Tx is created here to prevent blocking while requesting labels
		// Should only be 1 shipment given the context
		createdShipments, err := RequestAndSave(ctx, sorted)
		if err != nil {
			return nil, fmt.Errorf("must shipment: create new label: %w", err)
		}

		// Kind of a hack to get the parcel ID
		shipParcel, err := cli.Colli.Query().
			Where(colli.ID(colliID)).
			QueryShipmentParcel().
			Where(
				shipmentparcel.HasShipmentWith(
					shipment.StatusNEQ(shipment.StatusDeleted),
				),
			).OnlyID(ctx)
		if err != nil {
			return nil, fmt.Errorf("must shipment: query after create: %w", err)
		}

		if len(createdShipments) == 1 {
			createdShipments[0].ShipmentParcelID = shipParcel
		} else {
			return nil, fmt.Errorf("must shipment: expected only 1 colli/shipment")
		}

		return createdShipments, nil
	}

	labelPDF, err := mergeutils.FetchShipmentDocumentPDF(ctx, existingShip.ID)
	if err != nil {
		return nil, err
	}

	// Everything should be PDF labels for now
	if len(labelPDF) > 0 {
		return []SavedShipment{
			{
				ShipmentID:       existingShip.Edges.Shipment.ID,
				ShipmentParcelID: existingShip.ID,
				LabelsPDF: []string{
					labelPDF,
				},
			},
		}, nil
	}

	return nil, fmt.Errorf("expected existing shipment to have PDF")
}

func CreatePrintJobs(ctx context.Context, printers []*ent.Printer, shipments []SavedShipment) error {

	if len(printers) == 0 {
		return fmt.Errorf("could not create print jobs: 0 printers selected")
	}

	printerConfigured := false

	for _, prin := range printers {
		if prin.LabelPdf {
			for _, s := range shipments {
				for _, l := range s.LabelsPDF {
					err := createPrintjob(
						ctx,
						l,
						prin,
						printjob.FileExtensionPdf,
						s.ShipmentParcelID,
					)
					if err != nil {
						return err
					}
					printerConfigured = true
				}
			}
		}

		if prin.LabelZpl {
			for _, s := range shipments {
				for _, l := range s.LabelsPDF {
					zplLabel, err := utils.Base64PDFToZPL(l)
					if err != nil {
						return err
					}

					err = createPrintjob(
						ctx,
						base64.StdEncoding.EncodeToString([]byte(zplLabel)),
						prin,
						printjob.FileExtensionZpl,
						s.ShipmentParcelID,
					)
					if err != nil {
						return err
					}
					printerConfigured = true
				}
			}
		}
	}

	if !printerConfigured {
		return fmt.Errorf("could not create print jobs: printer not configured")
	}

	return nil
}

func createPrintjob(
	ctx context.Context,
	dataBase64 string,
	printer *ent.Printer,
	ext printjob.FileExtension,
	shipmentParcelID pulid.ID,
) error {
	cli := ent.FromContext(ctx)
	view := viewer.FromContext(ctx)
	return cli.PrintJob.Create().
		SetStatus(printjob.StatusPending).
		SetDocumentType(printjob.DocumentTypeParcelLabel).
		SetShipmentParcelID(shipmentParcelID).
		SetFileExtension(ext).
		SetBase64PrintData(dataBase64).
		SetPrinter(printer).
		SetTenantID(view.TenantID()).
		Exec(ctx)
}
