-- Disable the enforcement of foreign-keys constraints
PRAGMA foreign_keys = off;
-- Create "new_document_files" table
CREATE TABLE `new_document_files` (`id` text NOT NULL, `created_at` datetime NOT NULL, `storage_type` text NOT NULL, `storage_path` text NULL, `storage_path_zpl` text NULL DEFAULT (''), `doc_type` text NOT NULL, `data_pdf_base64` text NULL, `data_zpl_base64` text NULL, `colli_document_file` text NULL, `tenant_id` text NOT NULL, `shipment_parcel_document_file` text NULL, PRIMARY KEY (`id`), CONSTRAINT `document_files_collis_document_file` FOREIGN KEY (`colli_document_file`) REFERENCES `collis` (`id`) ON DELETE SET NULL, CONSTRAINT `document_files_tenants_tenant` FOREIGN KEY (`tenant_id`) REFERENCES `tenants` (`id`) ON DELETE NO ACTION, CONSTRAINT `document_files_shipment_parcels_document_file` FOREIGN KEY (`shipment_parcel_document_file`) REFERENCES `shipment_parcels` (`id`) ON DELETE SET NULL);
-- Copy rows from old table "document_files" to new temporary table "new_document_files"
INSERT INTO `new_document_files` (`id`, `created_at`, `storage_type`, `storage_path`, `doc_type`, `data_pdf_base64`, `data_zpl_base64`, `colli_document_file`, `tenant_id`) SELECT `id`, `created_at`, `storage_type`, `storage_path`, `doc_type`, `data_pdf_base64`, `data_zpl_base64`, `colli_document_file`, `tenant_id` FROM `document_files`;
-- Drop "document_files" table after copying rows
DROP TABLE `document_files`;
-- Rename temporary table "new_document_files" to "document_files"
ALTER TABLE `new_document_files` RENAME TO `document_files`;
-- Create index "document_files_shipment_parcel_document_file_key" to table: "document_files"
CREATE UNIQUE INDEX `document_files_shipment_parcel_document_file_key` ON `document_files` (`shipment_parcel_document_file`);
-- Create index "documentfile_tenant_id" to table: "document_files"
CREATE INDEX `documentfile_tenant_id` ON `document_files` (`tenant_id`);
-- Enable back the enforcement of foreign-keys constraints
PRAGMA foreign_keys = on;
