package utils

import (
	"context"
	"delivrio.io/go/appconfig"
	"delivrio.io/go/ent"
	"delivrio.io/go/ent/documentfile"
	"delivrio.io/go/viewer"
	"delivrio.io/shared-utils/pulid"
	"fmt"
	"gocloud.dev/blob"
	"regexp"
	"strings"
)

func DeleteDocument(ctx context.Context, doc *ent.DocumentFile) error {
	cli := ent.FromContext(ctx)

	switch conf.BlobStorage.Type {
	case appconfig.BlobStorageTypeBucketLocal:
		fallthrough
	case appconfig.BlobStorageTypeBucketS3:
		if len(doc.StoragePath) > 0 {
			err := cli.Bucket.Delete(ctx, doc.StoragePath)
			if err != nil {
				return err
			}
		}

		if len(doc.StoragePathZpl) > 0 {
			err := cli.Bucket.Delete(ctx, doc.StoragePathZpl)
			if err != nil {
				return err
			}
		}
	}

	_, err := cli.DocumentFile.Delete().
		Where(documentfile.ID(doc.ID)).
		Exec(ctx)

	return err
}

// Function to sanitize the path components
func sanitizePathComponent(component string) string {
	// Replace all invalid characters with an underscore
	re := regexp.MustCompile(`[^a-zA-Z0-9._/-]`)
	return re.ReplaceAllString(component, "_")
}

func CreateShipmentDocument(ctx context.Context, shipmentParcel *ent.ShipmentParcel, base64PDF *string, base64ZPL *string) (*ent.DocumentFile, error) {
	cli := ent.FromContext(ctx)
	view := viewer.FromContext(ctx)

	create := cli.DocumentFile.Create().
		SetDocType(documentfile.DocTypeCarrierLabel).
		SetTenantID(view.TenantID())

	if shipmentParcel != nil {
		create = create.SetShipmentParcel(shipmentParcel)
	}

	switch conf.BlobStorage.Type {
	case appconfig.BlobStorageTypeBucketLocal:
		fallthrough
	case appconfig.BlobStorageTypeBucketS3:

		create = create.SetStorageType(documentfile.StorageTypeBucket)
		if base64PDF != nil && *base64PDF != "" {
			saveKeyID := pulid.MustNew("")
			saveKey := fmt.Sprintf("%s/%s.pdf", sanitizePathComponent(conf.ServerID), saveKeyID.String())
			err := cli.Bucket.Upload(
				ctx,
				saveKey,
				strings.NewReader(*base64PDF),
				&blob.WriterOptions{ContentType: "application/pdf"},
			)
			if err != nil {
				return nil, err
			}
			// TODO: fix path to zpl
			create = create.SetStoragePath(saveKey)
		}

		if base64ZPL != nil && *base64ZPL != "" {
			saveKeyID := pulid.MustNew("")
			saveKey := fmt.Sprintf("%s/%s.zpl", sanitizePathComponent(conf.ServerID), saveKeyID.String())
			err := cli.Bucket.Upload(
				ctx,
				saveKey,
				strings.NewReader(*base64ZPL),
				&blob.WriterOptions{ContentType: "text/plain"},
			)
			if err != nil {
				return nil, err
			}
			create = create.SetStoragePathZpl(saveKey)
		}

	default:
		create = create.SetNillableDataPdfBase64(base64PDF).
			SetNillableDataZplBase64(base64ZPL).
			SetStorageType(documentfile.StorageTypeDatabase)
	}

	return create.Save(ctx)
}
