-- Disable the enforcement of foreign-keys constraints
PRAGMA foreign_keys = off;
-- Create "new_printers" table
CREATE TABLE `new_printers` (`id` text NOT NULL, `device_id` text NOT NULL, `name` text NOT NULL, `label_zpl` bool NOT NULL DEFAULT (false), `label_pdf` bool NOT NULL DEFAULT (false), `label_png` bool NOT NULL DEFAULT (false), `document` bool NOT NULL DEFAULT (false), `rotate_180` bool NOT NULL DEFAULT (false), `use_shell` bool NOT NULL DEFAULT (false), `print_size` text NOT NULL DEFAULT ('A4'), `created_at` datetime NOT NULL, `last_ping` datetime NOT NULL, `tenant_id` text NOT NULL, `workstation_printer` text NOT NULL, PRIMARY KEY (`id`), CONSTRAINT `printers_tenants_tenant` FOREIGN KEY (`tenant_id`) REFERENCES `tenants` (`id`) ON DELETE NO ACTION, CONSTRAINT `printers_workstations_printer` FOREIGN KEY (`workstation_printer`) REFERENCES `workstations` (`id`) ON DELETE NO ACTION);
-- Copy rows from old table "printers" to new temporary table "new_printers"
INSERT INTO `new_printers` (`id`, `device_id`, `name`, `label_zpl`, `label_pdf`, `label_png`, `document`, `rotate_180`, `print_size`, `created_at`, `last_ping`, `tenant_id`, `workstation_printer`) SELECT `id`, `device_id`, `name`, `label_zpl`, `label_pdf`, `label_png`, `document`, `rotate_180`, `print_size`, `created_at`, `last_ping`, `tenant_id`, `workstation_printer` FROM `printers`;
-- Drop "printers" table after copying rows
DROP TABLE `printers`;
-- Rename temporary table "new_printers" to "printers"
ALTER TABLE `new_printers` RENAME TO `printers`;
-- Create index "printers_device_id_key" to table: "printers"
CREATE UNIQUE INDEX `printers_device_id_key` ON `printers` (`device_id`);
-- Create index "printer_tenant_id" to table: "printers"
CREATE INDEX `printer_tenant_id` ON `printers` (`tenant_id`);
-- Enable back the enforcement of foreign-keys constraints
PRAGMA foreign_keys = on;
