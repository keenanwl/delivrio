-- Modify "document_files" table
ALTER TABLE "document_files" ADD COLUMN "storage_path_zpl" character varying NULL DEFAULT '', ADD COLUMN "shipment_parcel_document_file" character varying NULL, ADD CONSTRAINT "document_files_shipment_parcels_document_file" FOREIGN KEY ("shipment_parcel_document_file") REFERENCES "shipment_parcels" ("id") ON UPDATE NO ACTION ON DELETE SET NULL;
-- Create index "document_files_shipment_parcel_document_file_key" to table: "document_files"
CREATE UNIQUE INDEX "document_files_shipment_parcel_document_file_key" ON "document_files" ("shipment_parcel_document_file");
