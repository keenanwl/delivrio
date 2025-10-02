-- Disable the enforcement of foreign-keys constraints
PRAGMA foreign_keys = off;
-- Create "new_print_jobs" table
CREATE TABLE `new_print_jobs` (`id` text NOT NULL, `status` text NOT NULL, `file_extension` text NOT NULL, `use_shell` bool NOT NULL DEFAULT (false), `base64_print_data` text NOT NULL, `messages` json NULL, `created_at` datetime NOT NULL, `print_job_local_device` text NOT NULL, PRIMARY KEY (`id`), CONSTRAINT `print_jobs_local_devices_local_device` FOREIGN KEY (`print_job_local_device`) REFERENCES `local_devices` (`id`) ON DELETE NO ACTION);
-- Copy rows from old table "print_jobs" to new temporary table "new_print_jobs"
INSERT INTO `new_print_jobs` (`id`, `status`, `file_extension`, `base64_print_data`, `messages`, `created_at`, `print_job_local_device`) SELECT `id`, `status`, `file_extension`, `base64_print_data`, `messages`, `created_at`, `print_job_local_device` FROM `print_jobs`;
-- Drop "print_jobs" table after copying rows
DROP TABLE `print_jobs`;
-- Rename temporary table "new_print_jobs" to "print_jobs"
ALTER TABLE `new_print_jobs` RENAME TO `print_jobs`;
-- Enable back the enforcement of foreign-keys constraints
PRAGMA foreign_keys = on;
