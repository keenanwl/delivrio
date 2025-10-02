-- Modify "document_files" table
ALTER TABLE "document_files" ADD COLUMN "storage_type" character varying NOT NULL, ADD COLUMN "storage_path" character varying NULL;
