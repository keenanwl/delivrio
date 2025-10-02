-- Create "local_devices" table
CREATE TABLE `local_devices` (`id` text NOT NULL, `name` text NOT NULL, `system_name` text NOT NULL, `vendor_id` integer NULL, `product_id` integer NULL, `address` text NULL, `active` bool NOT NULL, `archived` bool NOT NULL DEFAULT false, `category` text NOT NULL, PRIMARY KEY (`id`));
-- Create index "local_devices_system_name_key" to table: "local_devices"
CREATE UNIQUE INDEX `local_devices_system_name_key` ON `local_devices` (`system_name`);
-- Create "log_errors" table
CREATE TABLE `log_errors` (`id` text NOT NULL, `error` text NOT NULL, `created_at` datetime NOT NULL, PRIMARY KEY (`id`));
-- Create "print_jobs" table
CREATE TABLE `print_jobs` (`id` text NOT NULL, `status` text NOT NULL, `file_extension` text NOT NULL, `base64_print_data` text NOT NULL, `messages` json NULL, `created_at` datetime NOT NULL, `print_job_local_device` text NOT NULL, PRIMARY KEY (`id`), CONSTRAINT `print_jobs_local_devices_local_device` FOREIGN KEY (`print_job_local_device`) REFERENCES `local_devices` (`id`) ON DELETE NO ACTION);
-- Create "recent_scans" table
CREATE TABLE `recent_scans` (`id` integer NOT NULL PRIMARY KEY AUTOINCREMENT, `scan_value` text NOT NULL, `response` text NOT NULL, `scan_type` text NOT NULL DEFAULT 'label_request', `created_at` datetime NOT NULL);
-- Create "remote_connections" table
CREATE TABLE `remote_connections` (`id` text NOT NULL, `remote_url` text NOT NULL, `registration_token` text NOT NULL, `workstation_name` text NOT NULL, `last_ping` datetime NULL, PRIMARY KEY (`id`));
-- Create index "remote_connections_remote_url_key" to table: "remote_connections"
CREATE UNIQUE INDEX `remote_connections_remote_url_key` ON `remote_connections` (`remote_url`);
-- Create "unique_computers" table
CREATE TABLE `unique_computers` (`id` text NOT NULL, PRIMARY KEY (`id`));
