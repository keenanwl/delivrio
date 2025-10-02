-- Disable the enforcement of foreign-keys constraints
PRAGMA foreign_keys = off;
-- Create "new_workstations" table
CREATE TABLE `new_workstations` (`id` text NOT NULL, `archived_at` datetime NULL, `name` text NOT NULL, `device_type` text NOT NULL DEFAULT ('label_station'), `registration_code` text NOT NULL, `workstation_id` text NOT NULL, `created_at` datetime NOT NULL, `last_ping` datetime NULL, `status` text NOT NULL DEFAULT ('pending'), `auto_print_receiver` bool NOT NULL DEFAULT (false), `user_selected_workstation` text NULL, `tenant_id` text NOT NULL, `workstation_user` text NULL, PRIMARY KEY (`id`), CONSTRAINT `workstations_users_selected_workstation` FOREIGN KEY (`user_selected_workstation`) REFERENCES `users` (`id`) ON DELETE SET NULL, CONSTRAINT `workstations_tenants_tenant` FOREIGN KEY (`tenant_id`) REFERENCES `tenants` (`id`) ON DELETE NO ACTION, CONSTRAINT `workstations_users_user` FOREIGN KEY (`workstation_user`) REFERENCES `users` (`id`) ON DELETE SET NULL);
-- Copy rows from old table "workstations" to new temporary table "new_workstations"
INSERT INTO `new_workstations` (`id`, `archived_at`, `name`, `device_type`, `registration_code`, `workstation_id`, `created_at`, `last_ping`, `status`, `user_selected_workstation`, `tenant_id`, `workstation_user`) SELECT `id`, `archived_at`, `name`, `device_type`, `registration_code`, `workstation_id`, `created_at`, `last_ping`, `status`, `user_selected_workstation`, `tenant_id`, `workstation_user` FROM `workstations`;
-- Drop "workstations" table after copying rows
DROP TABLE `workstations`;
-- Rename temporary table "new_workstations" to "workstations"
ALTER TABLE `new_workstations` RENAME TO `workstations`;
-- Create index "workstations_user_selected_workstation_key" to table: "workstations"
CREATE UNIQUE INDEX `workstations_user_selected_workstation_key` ON `workstations` (`user_selected_workstation`);
-- Create index "workstation_tenant_id" to table: "workstations"
CREATE INDEX `workstation_tenant_id` ON `workstations` (`tenant_id`);
-- Enable back the enforcement of foreign-keys constraints
PRAGMA foreign_keys = on;
