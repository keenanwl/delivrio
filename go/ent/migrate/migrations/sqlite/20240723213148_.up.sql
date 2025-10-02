-- Add column "printer_messages" to table: "print_jobs"
ALTER TABLE `print_jobs` ADD COLUMN `printer_messages` json NULL;
