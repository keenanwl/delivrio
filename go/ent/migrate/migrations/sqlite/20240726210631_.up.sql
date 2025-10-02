-- Add column "storage_type" to table: "document_files"
ALTER TABLE `document_files` ADD COLUMN `storage_type` text NOT NULL;
-- Add column "storage_path" to table: "document_files"
ALTER TABLE `document_files` ADD COLUMN `storage_path` text NULL;
