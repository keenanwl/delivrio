-- Modify "delivery_options" table
ALTER TABLE "delivery_options" ADD COLUMN "hide_if_company_empty" boolean NOT NULL DEFAULT false;
