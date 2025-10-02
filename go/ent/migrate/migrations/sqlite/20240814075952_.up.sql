-- Disable the enforcement of foreign-keys constraints
PRAGMA foreign_keys = off;
-- Create "new_shipment_parcels" table
CREATE TABLE `new_shipment_parcels` (`id` text NOT NULL, `item_id` text NULL, `status` text NOT NULL DEFAULT ('pending'), `cc_pickup_signature_urls` json NULL, `expected_at` datetime NULL, `fulfillment_synced_at` datetime NULL, `cancel_synced_at` datetime NULL, `colli_shipment_parcel` text NULL, `shipment_shipment_parcel` text NOT NULL, `tenant_id` text NOT NULL, `shipment_parcel_packaging` text NULL, PRIMARY KEY (`id`), CONSTRAINT `shipment_parcels_collis_shipment_parcel` FOREIGN KEY (`colli_shipment_parcel`) REFERENCES `collis` (`id`) ON DELETE SET NULL, CONSTRAINT `shipment_parcels_shipments_shipment_parcel` FOREIGN KEY (`shipment_shipment_parcel`) REFERENCES `shipments` (`id`) ON DELETE NO ACTION, CONSTRAINT `shipment_parcels_tenants_tenant` FOREIGN KEY (`tenant_id`) REFERENCES `tenants` (`id`) ON DELETE NO ACTION, CONSTRAINT `shipment_parcels_packagings_packaging` FOREIGN KEY (`shipment_parcel_packaging`) REFERENCES `packagings` (`id`) ON DELETE SET NULL);
-- Copy rows from old table "shipment_parcels" to new temporary table "new_shipment_parcels"
INSERT INTO `new_shipment_parcels` (`id`, `item_id`, `status`, `cc_pickup_signature_urls`, `expected_at`, `fulfillment_synced_at`, `cancel_synced_at`, `colli_shipment_parcel`, `shipment_shipment_parcel`, `tenant_id`, `shipment_parcel_packaging`) SELECT `id`, `item_id`, `status`, `cc_pickup_signature_urls`, `expected_at`, `fulfillment_synced_at`, `cancel_synced_at`, `colli_shipment_parcel`, `shipment_shipment_parcel`, `tenant_id`, `shipment_parcel_packaging` FROM `shipment_parcels`;
-- Drop "shipment_parcels" table after copying rows
DROP TABLE `shipment_parcels`;
-- Rename temporary table "new_shipment_parcels" to "shipment_parcels"
ALTER TABLE `new_shipment_parcels` RENAME TO `shipment_parcels`;
-- Create index "shipment_parcels_colli_shipment_parcel_key" to table: "shipment_parcels"
CREATE UNIQUE INDEX `shipment_parcels_colli_shipment_parcel_key` ON `shipment_parcels` (`colli_shipment_parcel`);
-- Create index "shipmentparcel_tenant_id" to table: "shipment_parcels"
CREATE INDEX `shipmentparcel_tenant_id` ON `shipment_parcels` (`tenant_id`);
-- Enable back the enforcement of foreign-keys constraints
PRAGMA foreign_keys = on;
