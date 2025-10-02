-- Create "api_tokens" table
CREATE TABLE `api_tokens` (`id` text NOT NULL, `name` text NOT NULL, `hashed_token` text NOT NULL, `created_at` datetime NULL, `last_used` datetime NULL, `tenant_id` text NOT NULL, `user_api_token` text NOT NULL, PRIMARY KEY (`id`), CONSTRAINT `api_tokens_tenants_tenant` FOREIGN KEY (`tenant_id`) REFERENCES `tenants` (`id`) ON DELETE NO ACTION, CONSTRAINT `api_tokens_users_api_token` FOREIGN KEY (`user_api_token`) REFERENCES `users` (`id`) ON DELETE NO ACTION);
-- Create index "api_tokens_hashed_token_key" to table: "api_tokens"
CREATE UNIQUE INDEX `api_tokens_hashed_token_key` ON `api_tokens` (`hashed_token`);
-- Create index "apitoken_tenant_id" to table: "api_tokens"
CREATE INDEX `apitoken_tenant_id` ON `api_tokens` (`tenant_id`);
-- Create "access_rights" table
CREATE TABLE `access_rights` (`id` text NOT NULL, `label` text NOT NULL, `internal_id` text NOT NULL, PRIMARY KEY (`id`));
-- Create index "access_rights_label_key" to table: "access_rights"
CREATE UNIQUE INDEX `access_rights_label_key` ON `access_rights` (`label`);
-- Create index "access_rights_internal_id_key" to table: "access_rights"
CREATE UNIQUE INDEX `access_rights_internal_id_key` ON `access_rights` (`internal_id`);
-- Create "addresses" table
CREATE TABLE `addresses` (`id` text NOT NULL, `uniqueness_id` text NULL, `first_name` text NOT NULL, `last_name` text NOT NULL, `email` text NOT NULL, `phone_number` text NOT NULL, `phone_number_2` text NULL, `vat_number` text NULL, `company` text NULL, `address_one` text NOT NULL, `address_two` text NOT NULL, `city` text NOT NULL, `state` text NULL, `zip` text NOT NULL, `tenant_id` text NOT NULL, `address_country` text NOT NULL, `consolidation_recipient` text NULL, `consolidation_sender` text NULL, PRIMARY KEY (`id`), CONSTRAINT `addresses_tenants_tenant` FOREIGN KEY (`tenant_id`) REFERENCES `tenants` (`id`) ON DELETE NO ACTION, CONSTRAINT `addresses_countries_country` FOREIGN KEY (`address_country`) REFERENCES `countries` (`id`) ON DELETE NO ACTION, CONSTRAINT `addresses_consolidations_recipient` FOREIGN KEY (`consolidation_recipient`) REFERENCES `consolidations` (`id`) ON DELETE SET NULL, CONSTRAINT `addresses_consolidations_sender` FOREIGN KEY (`consolidation_sender`) REFERENCES `consolidations` (`id`) ON DELETE SET NULL);
-- Create index "addresses_uniqueness_id_key" to table: "addresses"
CREATE UNIQUE INDEX `addresses_uniqueness_id_key` ON `addresses` (`uniqueness_id`);
-- Create index "addresses_consolidation_recipient_key" to table: "addresses"
CREATE UNIQUE INDEX `addresses_consolidation_recipient_key` ON `addresses` (`consolidation_recipient`);
-- Create index "addresses_consolidation_sender_key" to table: "addresses"
CREATE UNIQUE INDEX `addresses_consolidation_sender_key` ON `addresses` (`consolidation_sender`);
-- Create index "address_tenant_id" to table: "addresses"
CREATE INDEX `address_tenant_id` ON `addresses` (`tenant_id`);
-- Create "address_globals" table
CREATE TABLE `address_globals` (`id` text NOT NULL, `uniqueness_id` text NULL, `company` text NULL, `address_one` text NOT NULL, `address_two` text NULL, `city` text NOT NULL, `state` text NULL, `zip` text NOT NULL, `latitude` real NOT NULL DEFAULT (0), `longitude` real NOT NULL DEFAULT (0), `address_global_country` text NOT NULL, `parcel_shop_address` text NULL, `parcel_shop_bring_address_delivery` text NULL, `parcel_shop_post_nord_address_delivery` text NULL, PRIMARY KEY (`id`), CONSTRAINT `address_globals_countries_country` FOREIGN KEY (`address_global_country`) REFERENCES `countries` (`id`) ON DELETE NO ACTION, CONSTRAINT `address_globals_parcel_shops_address` FOREIGN KEY (`parcel_shop_address`) REFERENCES `parcel_shops` (`id`) ON DELETE SET NULL, CONSTRAINT `address_globals_parcel_shop_brings_address_delivery` FOREIGN KEY (`parcel_shop_bring_address_delivery`) REFERENCES `parcel_shop_brings` (`id`) ON DELETE SET NULL, CONSTRAINT `address_globals_parcel_shop_post_nords_address_delivery` FOREIGN KEY (`parcel_shop_post_nord_address_delivery`) REFERENCES `parcel_shop_post_nords` (`id`) ON DELETE SET NULL);
-- Create index "address_globals_uniqueness_id_key" to table: "address_globals"
CREATE UNIQUE INDEX `address_globals_uniqueness_id_key` ON `address_globals` (`uniqueness_id`);
-- Create index "address_globals_parcel_shop_address_key" to table: "address_globals"
CREATE UNIQUE INDEX `address_globals_parcel_shop_address_key` ON `address_globals` (`parcel_shop_address`);
-- Create index "address_globals_parcel_shop_bring_address_delivery_key" to table: "address_globals"
CREATE UNIQUE INDEX `address_globals_parcel_shop_bring_address_delivery_key` ON `address_globals` (`parcel_shop_bring_address_delivery`);
-- Create index "address_globals_parcel_shop_post_nord_address_delivery_key" to table: "address_globals"
CREATE UNIQUE INDEX `address_globals_parcel_shop_post_nord_address_delivery_key` ON `address_globals` (`parcel_shop_post_nord_address_delivery`);
-- Create "business_hours_periods" table
CREATE TABLE `business_hours_periods` (`id` text NOT NULL, `day_of_week` text NOT NULL, `opening` datetime NOT NULL, `closing` datetime NOT NULL, `parcel_shop_business_hours_period` text NOT NULL, PRIMARY KEY (`id`), CONSTRAINT `business_hours_periods_parcel_shops_business_hours_period` FOREIGN KEY (`parcel_shop_business_hours_period`) REFERENCES `parcel_shops` (`id`) ON DELETE NO ACTION);
-- Create "carriers" table
CREATE TABLE `carriers` (`id` text NOT NULL, `name` text NOT NULL, `sync_cancelation` bool NOT NULL DEFAULT (false), `tenant_id` text NOT NULL, `carrier_carrier_brand` text NOT NULL, PRIMARY KEY (`id`), CONSTRAINT `carriers_tenants_tenant` FOREIGN KEY (`tenant_id`) REFERENCES `tenants` (`id`) ON DELETE NO ACTION, CONSTRAINT `carriers_carrier_brands_carrier_brand` FOREIGN KEY (`carrier_carrier_brand`) REFERENCES `carrier_brands` (`id`) ON DELETE NO ACTION);
-- Create index "carrier_tenant_id" to table: "carriers"
CREATE INDEX `carrier_tenant_id` ON `carriers` (`tenant_id`);
-- Create "carrier_additional_service_brings" table
CREATE TABLE `carrier_additional_service_brings` (`id` text NOT NULL, `label` text NOT NULL, `api_code_booking` text NOT NULL, `carrier_service_bring_carrier_additional_service_bring` text NULL, PRIMARY KEY (`id`), CONSTRAINT `carrier_additional_service_brings_carrier_service_brings_carrier_additional_service_bring` FOREIGN KEY (`carrier_service_bring_carrier_additional_service_bring`) REFERENCES `carrier_service_brings` (`id`) ON DELETE SET NULL);
-- Create "carrier_additional_service_da_os" table
CREATE TABLE `carrier_additional_service_da_os` (`id` text NOT NULL, `label` text NOT NULL, `api_code` text NOT NULL, PRIMARY KEY (`id`));
-- Create "carrier_additional_service_dfs" table
CREATE TABLE `carrier_additional_service_dfs` (`id` text NOT NULL, `label` text NOT NULL, `api_code` text NOT NULL, PRIMARY KEY (`id`));
-- Create "carrier_additional_service_ds_vs" table
CREATE TABLE `carrier_additional_service_ds_vs` (`id` text NOT NULL, `label` text NOT NULL, `api_code` text NOT NULL, PRIMARY KEY (`id`));
-- Create "carrier_additional_service_easy_posts" table
CREATE TABLE `carrier_additional_service_easy_posts` (`id` text NOT NULL, `label` text NOT NULL, `api_key` text NOT NULL, `api_value` text NOT NULL, PRIMARY KEY (`id`));
-- Create "carrier_additional_service_gl_ss" table
CREATE TABLE `carrier_additional_service_gl_ss` (`id` text NOT NULL, `label` text NOT NULL, `mandatory` bool NOT NULL, `all_countries_consignor` bool NOT NULL DEFAULT (false), `all_countries_consignee` bool NOT NULL DEFAULT (false), `internal_id` text NOT NULL, `carrier_service_gls_carrier_additional_service_gls` text NULL, PRIMARY KEY (`id`), CONSTRAINT `carrier_additional_service_gl_ss_carrier_service_gl_ss_carrier_additional_service_gls` FOREIGN KEY (`carrier_service_gls_carrier_additional_service_gls`) REFERENCES `carrier_service_gl_ss` (`id`) ON DELETE SET NULL);
-- Create "carrier_additional_service_post_nords" table
CREATE TABLE `carrier_additional_service_post_nords` (`id` text NOT NULL, `label` text NOT NULL, `mandatory` bool NOT NULL, `all_countries_consignor` bool NOT NULL DEFAULT (false), `all_countries_consignee` bool NOT NULL DEFAULT (false), `internal_id` text NOT NULL, `api_code` text NOT NULL, `carrier_service_post_nord_carrier_add_serv_post_nord` text NULL, PRIMARY KEY (`id`), CONSTRAINT `carrier_additional_service_post_nords_carrier_service_post_nords_carrier_add_serv_post_nord` FOREIGN KEY (`carrier_service_post_nord_carrier_add_serv_post_nord`) REFERENCES `carrier_service_post_nords` (`id`) ON DELETE SET NULL);
-- Create index "carrieradditionalservicepostnord_internal_id_carrier_service_post_nord_carrier_add_serv_post_nord" to table: "carrier_additional_service_post_nords"
CREATE UNIQUE INDEX `carrieradditionalservicepostnord_internal_id_carrier_service_post_nord_carrier_add_serv_post_nord` ON `carrier_additional_service_post_nords` (`internal_id`, `carrier_service_post_nord_carrier_add_serv_post_nord`);
-- Create "carrier_additional_service_usp_ss" table
CREATE TABLE `carrier_additional_service_usp_ss` (`id` text NOT NULL, `label` text NOT NULL, `commonly_used` bool NOT NULL DEFAULT (false), `internal_id` text NOT NULL, `api_code` text NOT NULL, `carrier_service_usps_carrier_additional_service_usps` text NULL, PRIMARY KEY (`id`), CONSTRAINT `carrier_additional_service_usp_ss_carrier_service_usp_ss_carrier_additional_service_usps` FOREIGN KEY (`carrier_service_usps_carrier_additional_service_usps`) REFERENCES `carrier_service_usp_ss` (`id`) ON DELETE SET NULL);
-- Create "carrier_brands" table
CREATE TABLE `carrier_brands` (`id` text NOT NULL, `label` text NOT NULL, `label_short` text NOT NULL, `internal_id` text NOT NULL, `logo_url` text NULL, `text_color` text NULL DEFAULT ('#FFFFFF'), `background_color` text NULL DEFAULT ('#000000'), PRIMARY KEY (`id`));
-- Create index "carrierbrand_internal_id" to table: "carrier_brands"
CREATE UNIQUE INDEX `carrierbrand_internal_id` ON `carrier_brands` (`internal_id`);
-- Create "carrier_brings" table
CREATE TABLE `carrier_brings` (`id` text NOT NULL, `api_key` text NULL, `customer_number` text NULL, `test` bool NOT NULL DEFAULT (true), `carrier_carrier_bring` text NOT NULL, `tenant_id` text NOT NULL, PRIMARY KEY (`id`), CONSTRAINT `carrier_brings_carriers_carrier_bring` FOREIGN KEY (`carrier_carrier_bring`) REFERENCES `carriers` (`id`) ON DELETE NO ACTION, CONSTRAINT `carrier_brings_tenants_tenant` FOREIGN KEY (`tenant_id`) REFERENCES `tenants` (`id`) ON DELETE NO ACTION);
-- Create index "carrier_brings_carrier_carrier_bring_key" to table: "carrier_brings"
CREATE UNIQUE INDEX `carrier_brings_carrier_carrier_bring_key` ON `carrier_brings` (`carrier_carrier_bring`);
-- Create index "carrierbring_tenant_id" to table: "carrier_brings"
CREATE INDEX `carrierbring_tenant_id` ON `carrier_brings` (`tenant_id`);
-- Create "carrier_da_os" table
CREATE TABLE `carrier_da_os` (`id` text NOT NULL, `customer_id` text NULL, `api_key` text NULL, `test` bool NOT NULL DEFAULT (true), `carrier_carrier_dao` text NOT NULL, `tenant_id` text NOT NULL, PRIMARY KEY (`id`), CONSTRAINT `carrier_da_os_carriers_carrier_dao` FOREIGN KEY (`carrier_carrier_dao`) REFERENCES `carriers` (`id`) ON DELETE NO ACTION, CONSTRAINT `carrier_da_os_tenants_tenant` FOREIGN KEY (`tenant_id`) REFERENCES `tenants` (`id`) ON DELETE NO ACTION);
-- Create index "carrier_da_os_carrier_carrier_dao_key" to table: "carrier_da_os"
CREATE UNIQUE INDEX `carrier_da_os_carrier_carrier_dao_key` ON `carrier_da_os` (`carrier_carrier_dao`);
-- Create index "carrierdao_tenant_id" to table: "carrier_da_os"
CREATE INDEX `carrierdao_tenant_id` ON `carrier_da_os` (`tenant_id`);
-- Create "carrier_dfs" table
CREATE TABLE `carrier_dfs` (`id` text NOT NULL, `customer_id` text NOT NULL, `agreement_number` text NOT NULL, `who_pays` text NOT NULL DEFAULT ('Prepaid'), `test` bool NOT NULL DEFAULT (true), `carrier_carrier_df` text NOT NULL, `tenant_id` text NOT NULL, PRIMARY KEY (`id`), CONSTRAINT `carrier_dfs_carriers_carrier_df` FOREIGN KEY (`carrier_carrier_df`) REFERENCES `carriers` (`id`) ON DELETE NO ACTION, CONSTRAINT `carrier_dfs_tenants_tenant` FOREIGN KEY (`tenant_id`) REFERENCES `tenants` (`id`) ON DELETE NO ACTION);
-- Create index "carrier_dfs_carrier_carrier_df_key" to table: "carrier_dfs"
CREATE UNIQUE INDEX `carrier_dfs_carrier_carrier_df_key` ON `carrier_dfs` (`carrier_carrier_df`);
-- Create index "carrierdf_tenant_id" to table: "carrier_dfs"
CREATE INDEX `carrierdf_tenant_id` ON `carrier_dfs` (`tenant_id`);
-- Create "carrier_ds_vs" table
CREATE TABLE `carrier_ds_vs` (`id` text NOT NULL, `carrier_carrier_dsv` text NOT NULL, `tenant_id` text NOT NULL, PRIMARY KEY (`id`), CONSTRAINT `carrier_ds_vs_carriers_carrier_dsv` FOREIGN KEY (`carrier_carrier_dsv`) REFERENCES `carriers` (`id`) ON DELETE NO ACTION, CONSTRAINT `carrier_ds_vs_tenants_tenant` FOREIGN KEY (`tenant_id`) REFERENCES `tenants` (`id`) ON DELETE NO ACTION);
-- Create index "carrier_ds_vs_carrier_carrier_dsv_key" to table: "carrier_ds_vs"
CREATE UNIQUE INDEX `carrier_ds_vs_carrier_carrier_dsv_key` ON `carrier_ds_vs` (`carrier_carrier_dsv`);
-- Create index "carrierdsv_tenant_id" to table: "carrier_ds_vs"
CREATE INDEX `carrierdsv_tenant_id` ON `carrier_ds_vs` (`tenant_id`);
-- Create "carrier_easy_posts" table
CREATE TABLE `carrier_easy_posts` (`id` text NOT NULL, `api_key` text NOT NULL, `test` bool NOT NULL DEFAULT (true), `carrier_accounts` json NOT NULL, `carrier_carrier_easy_post` text NOT NULL, `tenant_id` text NOT NULL, PRIMARY KEY (`id`), CONSTRAINT `carrier_easy_posts_carriers_carrier_easy_post` FOREIGN KEY (`carrier_carrier_easy_post`) REFERENCES `carriers` (`id`) ON DELETE NO ACTION, CONSTRAINT `carrier_easy_posts_tenants_tenant` FOREIGN KEY (`tenant_id`) REFERENCES `tenants` (`id`) ON DELETE NO ACTION);
-- Create index "carrier_easy_posts_carrier_carrier_easy_post_key" to table: "carrier_easy_posts"
CREATE UNIQUE INDEX `carrier_easy_posts_carrier_carrier_easy_post_key` ON `carrier_easy_posts` (`carrier_carrier_easy_post`);
-- Create index "carriereasypost_tenant_id" to table: "carrier_easy_posts"
CREATE INDEX `carriereasypost_tenant_id` ON `carrier_easy_posts` (`tenant_id`);
-- Create "carrier_gl_ss" table
CREATE TABLE `carrier_gl_ss` (`id` text NOT NULL, `contact_id` text NULL, `gls_username` text NULL, `gls_password` text NULL, `customer_id` text NULL, `gls_country_code` text NULL, `sync_shipment_cancellation` bool NULL DEFAULT (false), `print_error_on_label` bool NULL DEFAULT (false), `carrier_carrier_gls` text NOT NULL, `tenant_id` text NOT NULL, PRIMARY KEY (`id`), CONSTRAINT `carrier_gl_ss_carriers_carrier_gls` FOREIGN KEY (`carrier_carrier_gls`) REFERENCES `carriers` (`id`) ON DELETE NO ACTION, CONSTRAINT `carrier_gl_ss_tenants_tenant` FOREIGN KEY (`tenant_id`) REFERENCES `tenants` (`id`) ON DELETE NO ACTION);
-- Create index "carrier_gl_ss_carrier_carrier_gls_key" to table: "carrier_gl_ss"
CREATE UNIQUE INDEX `carrier_gl_ss_carrier_carrier_gls_key` ON `carrier_gl_ss` (`carrier_carrier_gls`);
-- Create index "carriergls_tenant_id" to table: "carrier_gl_ss"
CREATE INDEX `carriergls_tenant_id` ON `carrier_gl_ss` (`tenant_id`);
-- Create "carrier_post_nords" table
CREATE TABLE `carrier_post_nords` (`id` text NOT NULL, `customer_number` text NOT NULL DEFAULT (''), `carrier_carrier_post_nord` text NOT NULL, `tenant_id` text NOT NULL, PRIMARY KEY (`id`), CONSTRAINT `carrier_post_nords_carriers_carrier_post_nord` FOREIGN KEY (`carrier_carrier_post_nord`) REFERENCES `carriers` (`id`) ON DELETE NO ACTION, CONSTRAINT `carrier_post_nords_tenants_tenant` FOREIGN KEY (`tenant_id`) REFERENCES `tenants` (`id`) ON DELETE NO ACTION);
-- Create index "carrier_post_nords_carrier_carrier_post_nord_key" to table: "carrier_post_nords"
CREATE UNIQUE INDEX `carrier_post_nords_carrier_carrier_post_nord_key` ON `carrier_post_nords` (`carrier_carrier_post_nord`);
-- Create index "carrierpostnord_tenant_id" to table: "carrier_post_nords"
CREATE INDEX `carrierpostnord_tenant_id` ON `carrier_post_nords` (`tenant_id`);
-- Create "carrier_services" table
CREATE TABLE `carrier_services` (`id` text NOT NULL, `label` text NOT NULL, `internal_id` text NOT NULL, `return` bool NOT NULL DEFAULT (false), `consolidation` bool NOT NULL DEFAULT (false), `delivery_point_optional` bool NOT NULL DEFAULT (false), `delivery_point_required` bool NOT NULL DEFAULT (false), `carrier_brand_carrier_service` text NOT NULL, PRIMARY KEY (`id`), CONSTRAINT `carrier_services_carrier_brands_carrier_service` FOREIGN KEY (`carrier_brand_carrier_service`) REFERENCES `carrier_brands` (`id`) ON DELETE NO ACTION);
-- Create index "carrier_services_internal_id_key" to table: "carrier_services"
CREATE UNIQUE INDEX `carrier_services_internal_id_key` ON `carrier_services` (`internal_id`);
-- Create "carrier_service_brings" table
CREATE TABLE `carrier_service_brings` (`id` text NOT NULL, `api_service_code` text NOT NULL, `api_request` text NOT NULL, `carrier_service_carrier_service_bring` text NOT NULL, PRIMARY KEY (`id`), CONSTRAINT `carrier_service_brings_carrier_services_carrier_service_bring` FOREIGN KEY (`carrier_service_carrier_service_bring`) REFERENCES `carrier_services` (`id`) ON DELETE NO ACTION);
-- Create index "carrier_service_brings_carrier_service_carrier_service_bring_key" to table: "carrier_service_brings"
CREATE UNIQUE INDEX `carrier_service_brings_carrier_service_carrier_service_bring_key` ON `carrier_service_brings` (`carrier_service_carrier_service_bring`);
-- Create "carrier_service_da_os" table
CREATE TABLE `carrier_service_da_os` (`id` text NOT NULL, `carrier_service_carrier_service_dao` text NOT NULL, PRIMARY KEY (`id`), CONSTRAINT `carrier_service_da_os_carrier_services_carrier_service_dao` FOREIGN KEY (`carrier_service_carrier_service_dao`) REFERENCES `carrier_services` (`id`) ON DELETE NO ACTION);
-- Create index "carrier_service_da_os_carrier_service_carrier_service_dao_key" to table: "carrier_service_da_os"
CREATE UNIQUE INDEX `carrier_service_da_os_carrier_service_carrier_service_dao_key` ON `carrier_service_da_os` (`carrier_service_carrier_service_dao`);
-- Create "carrier_service_dfs" table
CREATE TABLE `carrier_service_dfs` (`id` text NOT NULL, `carrier_service_carrier_service_df` text NOT NULL, PRIMARY KEY (`id`), CONSTRAINT `carrier_service_dfs_carrier_services_carrier_service_df` FOREIGN KEY (`carrier_service_carrier_service_df`) REFERENCES `carrier_services` (`id`) ON DELETE NO ACTION);
-- Create index "carrier_service_dfs_carrier_service_carrier_service_df_key" to table: "carrier_service_dfs"
CREATE UNIQUE INDEX `carrier_service_dfs_carrier_service_carrier_service_df_key` ON `carrier_service_dfs` (`carrier_service_carrier_service_df`);
-- Create "carrier_service_ds_vs" table
CREATE TABLE `carrier_service_ds_vs` (`id` text NOT NULL, `carrier_service_carrier_service_dsv` text NOT NULL, PRIMARY KEY (`id`), CONSTRAINT `carrier_service_ds_vs_carrier_services_carrier_service_dsv` FOREIGN KEY (`carrier_service_carrier_service_dsv`) REFERENCES `carrier_services` (`id`) ON DELETE NO ACTION);
-- Create index "carrier_service_ds_vs_carrier_service_carrier_service_dsv_key" to table: "carrier_service_ds_vs"
CREATE UNIQUE INDEX `carrier_service_ds_vs_carrier_service_carrier_service_dsv_key` ON `carrier_service_ds_vs` (`carrier_service_carrier_service_dsv`);
-- Create "carrier_service_easy_posts" table
CREATE TABLE `carrier_service_easy_posts` (`id` text NOT NULL, `api_key` text NOT NULL, `carrier_service_carrier_serv_easy_post` text NOT NULL, PRIMARY KEY (`id`), CONSTRAINT `carrier_service_easy_posts_carrier_services_carrier_serv_easy_post` FOREIGN KEY (`carrier_service_carrier_serv_easy_post`) REFERENCES `carrier_services` (`id`) ON DELETE NO ACTION);
-- Create index "carrier_service_easy_posts_carrier_service_carrier_serv_easy_post_key" to table: "carrier_service_easy_posts"
CREATE UNIQUE INDEX `carrier_service_easy_posts_carrier_service_carrier_serv_easy_post_key` ON `carrier_service_easy_posts` (`carrier_service_carrier_serv_easy_post`);
-- Create "carrier_service_gl_ss" table
CREATE TABLE `carrier_service_gl_ss` (`id` text NOT NULL, `api_key` text NULL, `api_value` text NOT NULL, `carrier_service_carrier_service_gls` text NOT NULL, PRIMARY KEY (`id`), CONSTRAINT `carrier_service_gl_ss_carrier_services_carrier_service_gls` FOREIGN KEY (`carrier_service_carrier_service_gls`) REFERENCES `carrier_services` (`id`) ON DELETE NO ACTION);
-- Create index "carrier_service_gl_ss_api_key_key" to table: "carrier_service_gl_ss"
CREATE UNIQUE INDEX `carrier_service_gl_ss_api_key_key` ON `carrier_service_gl_ss` (`api_key`);
-- Create index "carrier_service_gl_ss_carrier_service_carrier_service_gls_key" to table: "carrier_service_gl_ss"
CREATE UNIQUE INDEX `carrier_service_gl_ss_carrier_service_carrier_service_gls_key` ON `carrier_service_gl_ss` (`carrier_service_carrier_service_gls`);
-- Create "carrier_service_post_nords" table
CREATE TABLE `carrier_service_post_nords` (`id` text NOT NULL, `label` text NOT NULL, `internal_id` text NOT NULL, `api_code` text NOT NULL, `carrier_service_carrier_service_post_nord` text NOT NULL, PRIMARY KEY (`id`), CONSTRAINT `carrier_service_post_nords_carrier_services_carrier_service_post_nord` FOREIGN KEY (`carrier_service_carrier_service_post_nord`) REFERENCES `carrier_services` (`id`) ON DELETE NO ACTION);
-- Create index "carrier_service_post_nords_internal_id_key" to table: "carrier_service_post_nords"
CREATE UNIQUE INDEX `carrier_service_post_nords_internal_id_key` ON `carrier_service_post_nords` (`internal_id`);
-- Create index "carrier_service_post_nords_carrier_service_carrier_service_post_nord_key" to table: "carrier_service_post_nords"
CREATE UNIQUE INDEX `carrier_service_post_nords_carrier_service_carrier_service_post_nord_key` ON `carrier_service_post_nords` (`carrier_service_carrier_service_post_nord`);
-- Create "carrier_service_usp_ss" table
CREATE TABLE `carrier_service_usp_ss` (`id` text NOT NULL, `api_key` text NOT NULL, `carrier_service_carrier_service_usps` text NOT NULL, PRIMARY KEY (`id`), CONSTRAINT `carrier_service_usp_ss_carrier_services_carrier_service_usps` FOREIGN KEY (`carrier_service_carrier_service_usps`) REFERENCES `carrier_services` (`id`) ON DELETE NO ACTION);
-- Create index "carrier_service_usp_ss_carrier_service_carrier_service_usps_key" to table: "carrier_service_usp_ss"
CREATE UNIQUE INDEX `carrier_service_usp_ss_carrier_service_carrier_service_usps_key` ON `carrier_service_usp_ss` (`carrier_service_carrier_service_usps`);
-- Create "carrier_usp_ss" table
CREATE TABLE `carrier_usp_ss` (`id` text NOT NULL, `is_test_api` bool NOT NULL DEFAULT (false), `consumer_key` text NULL, `consumer_secret` text NULL, `mid` text NULL, `manifest_mid` text NULL, `crid` text NULL, `eps_account_number` text NULL, `carrier_carrier_usps` text NOT NULL, `tenant_id` text NOT NULL, PRIMARY KEY (`id`), CONSTRAINT `carrier_usp_ss_carriers_carrier_usps` FOREIGN KEY (`carrier_carrier_usps`) REFERENCES `carriers` (`id`) ON DELETE NO ACTION, CONSTRAINT `carrier_usp_ss_tenants_tenant` FOREIGN KEY (`tenant_id`) REFERENCES `tenants` (`id`) ON DELETE NO ACTION);
-- Create index "carrier_usp_ss_carrier_carrier_usps_key" to table: "carrier_usp_ss"
CREATE UNIQUE INDEX `carrier_usp_ss_carrier_carrier_usps_key` ON `carrier_usp_ss` (`carrier_carrier_usps`);
-- Create index "carrierusps_tenant_id" to table: "carrier_usp_ss"
CREATE INDEX `carrierusps_tenant_id` ON `carrier_usp_ss` (`tenant_id`);
-- Create "change_histories" table
CREATE TABLE `change_histories` (`id` text NOT NULL, `created_at` datetime NOT NULL, `origin` text NOT NULL DEFAULT ('unknown'), `tenant_id` text NOT NULL, `change_history_user` text NULL, PRIMARY KEY (`id`), CONSTRAINT `change_histories_tenants_tenant` FOREIGN KEY (`tenant_id`) REFERENCES `tenants` (`id`) ON DELETE NO ACTION, CONSTRAINT `change_histories_users_user` FOREIGN KEY (`change_history_user`) REFERENCES `users` (`id`) ON DELETE SET NULL);
-- Create index "changehistory_tenant_id" to table: "change_histories"
CREATE INDEX `changehistory_tenant_id` ON `change_histories` (`tenant_id`);
-- Create "collis" table
CREATE TABLE `collis` (`id` text NOT NULL, `internal_barcode` integer NULL, `status` text NOT NULL, `slip_print_status` text NOT NULL DEFAULT ('pending'), `created_at` datetime NOT NULL, `email_packing_slip_printed_at` datetime NULL, `email_label_printed_at` datetime NULL, `tenant_id` text NOT NULL, `colli_recipient` text NOT NULL, `colli_sender` text NOT NULL, `colli_parcel_shop` text NULL, `colli_click_collect_location` text NULL, `colli_delivery_option` text NULL, `colli_packaging` text NULL, `order_colli` text NOT NULL, PRIMARY KEY (`id`), CONSTRAINT `collis_tenants_tenant` FOREIGN KEY (`tenant_id`) REFERENCES `tenants` (`id`) ON DELETE NO ACTION, CONSTRAINT `collis_addresses_recipient` FOREIGN KEY (`colli_recipient`) REFERENCES `addresses` (`id`) ON DELETE NO ACTION, CONSTRAINT `collis_addresses_sender` FOREIGN KEY (`colli_sender`) REFERENCES `addresses` (`id`) ON DELETE NO ACTION, CONSTRAINT `collis_parcel_shops_parcel_shop` FOREIGN KEY (`colli_parcel_shop`) REFERENCES `parcel_shops` (`id`) ON DELETE SET NULL, CONSTRAINT `collis_locations_click_collect_location` FOREIGN KEY (`colli_click_collect_location`) REFERENCES `locations` (`id`) ON DELETE SET NULL, CONSTRAINT `collis_delivery_options_delivery_option` FOREIGN KEY (`colli_delivery_option`) REFERENCES `delivery_options` (`id`) ON DELETE SET NULL, CONSTRAINT `collis_packagings_packaging` FOREIGN KEY (`colli_packaging`) REFERENCES `packagings` (`id`) ON DELETE SET NULL, CONSTRAINT `collis_orders_colli` FOREIGN KEY (`order_colli`) REFERENCES `orders` (`id`) ON DELETE NO ACTION);
-- Create index "colli_tenant_id" to table: "collis"
CREATE INDEX `colli_tenant_id` ON `collis` (`tenant_id`);
-- Create index "colli_internal_barcode_tenant_id" to table: "collis"
CREATE UNIQUE INDEX `colli_internal_barcode_tenant_id` ON `collis` (`internal_barcode`, `tenant_id`);
-- Create "connect_option_carriers" table
CREATE TABLE `connect_option_carriers` (`id` text NOT NULL, `name` text NOT NULL, PRIMARY KEY (`id`));
-- Create index "connect_option_carriers_name_key" to table: "connect_option_carriers"
CREATE UNIQUE INDEX `connect_option_carriers_name_key` ON `connect_option_carriers` (`name`);
-- Create "connect_option_platforms" table
CREATE TABLE `connect_option_platforms` (`id` text NOT NULL, `name` text NOT NULL, PRIMARY KEY (`id`));
-- Create index "connect_option_platforms_name_key" to table: "connect_option_platforms"
CREATE UNIQUE INDEX `connect_option_platforms_name_key` ON `connect_option_platforms` (`name`);
-- Create "connections" table
CREATE TABLE `connections` (`id` text NOT NULL, `name` text NOT NULL, `sync_orders` bool NOT NULL DEFAULT (false), `sync_products` bool NOT NULL DEFAULT (false), `fulfill_automatically` bool NOT NULL DEFAULT (false), `dispatch_automatically` bool NOT NULL DEFAULT (false), `convert_currency` bool NOT NULL DEFAULT (false), `tenant_id` text NOT NULL, `connection_connection_brand` text NOT NULL, `connection_sender_location` text NOT NULL, `connection_pickup_location` text NOT NULL, `connection_return_location` text NOT NULL, `connection_seller_location` text NOT NULL, `connection_currency` text NOT NULL, `connection_packing_slip_template` text NULL, `return_portal_connection` text NULL, PRIMARY KEY (`id`), CONSTRAINT `connections_tenants_tenant` FOREIGN KEY (`tenant_id`) REFERENCES `tenants` (`id`) ON DELETE NO ACTION, CONSTRAINT `connections_connection_brands_connection_brand` FOREIGN KEY (`connection_connection_brand`) REFERENCES `connection_brands` (`id`) ON DELETE NO ACTION, CONSTRAINT `connections_locations_sender_location` FOREIGN KEY (`connection_sender_location`) REFERENCES `locations` (`id`) ON DELETE NO ACTION, CONSTRAINT `connections_locations_pickup_location` FOREIGN KEY (`connection_pickup_location`) REFERENCES `locations` (`id`) ON DELETE NO ACTION, CONSTRAINT `connections_locations_return_location` FOREIGN KEY (`connection_return_location`) REFERENCES `locations` (`id`) ON DELETE NO ACTION, CONSTRAINT `connections_locations_seller_location` FOREIGN KEY (`connection_seller_location`) REFERENCES `locations` (`id`) ON DELETE NO ACTION, CONSTRAINT `connections_currencies_currency` FOREIGN KEY (`connection_currency`) REFERENCES `currencies` (`id`) ON DELETE NO ACTION, CONSTRAINT `connections_documents_packing_slip_template` FOREIGN KEY (`connection_packing_slip_template`) REFERENCES `documents` (`id`) ON DELETE SET NULL, CONSTRAINT `connections_return_portals_connection` FOREIGN KEY (`return_portal_connection`) REFERENCES `return_portals` (`id`) ON DELETE SET NULL);
-- Create index "connections_return_portal_connection_key" to table: "connections"
CREATE UNIQUE INDEX `connections_return_portal_connection_key` ON `connections` (`return_portal_connection`);
-- Create index "connection_tenant_id" to table: "connections"
CREATE INDEX `connection_tenant_id` ON `connections` (`tenant_id`);
-- Create "connection_brands" table
CREATE TABLE `connection_brands` (`id` text NOT NULL, `label` text NOT NULL, `internal_id` text NOT NULL DEFAULT ('shopify'), `logo_url` text NULL, PRIMARY KEY (`id`));
-- Create index "connection_brands_label_key" to table: "connection_brands"
CREATE UNIQUE INDEX `connection_brands_label_key` ON `connection_brands` (`label`);
-- Create "connection_lookups" table
CREATE TABLE `connection_lookups` (`id` text NOT NULL, `payload` text NOT NULL, `options_output_count` integer NOT NULL, `error` text NULL, `created_at` datetime NOT NULL, `tenant_id` text NOT NULL, `connection_lookup_connections` text NULL, PRIMARY KEY (`id`), CONSTRAINT `connection_lookups_tenants_tenant` FOREIGN KEY (`tenant_id`) REFERENCES `tenants` (`id`) ON DELETE NO ACTION, CONSTRAINT `connection_lookups_connections_connections` FOREIGN KEY (`connection_lookup_connections`) REFERENCES `connections` (`id`) ON DELETE SET NULL);
-- Create index "connectionlookup_tenant_id" to table: "connection_lookups"
CREATE INDEX `connectionlookup_tenant_id` ON `connection_lookups` (`tenant_id`);
-- Create "connection_shopifies" table
CREATE TABLE `connection_shopifies` (`id` text NOT NULL, `rate_integration` bool NOT NULL DEFAULT (false), `store_url` text NULL, `api_key` text NULL, `lookup_key` text NULL, `connection_connection_shopify` text NOT NULL, `tenant_id` text NOT NULL, PRIMARY KEY (`id`), CONSTRAINT `connection_shopifies_connections_connection_shopify` FOREIGN KEY (`connection_connection_shopify`) REFERENCES `connections` (`id`) ON DELETE NO ACTION, CONSTRAINT `connection_shopifies_tenants_tenant` FOREIGN KEY (`tenant_id`) REFERENCES `tenants` (`id`) ON DELETE NO ACTION);
-- Create index "connection_shopifies_store_url_key" to table: "connection_shopifies"
CREATE UNIQUE INDEX `connection_shopifies_store_url_key` ON `connection_shopifies` (`store_url`);
-- Create index "connection_shopifies_connection_connection_shopify_key" to table: "connection_shopifies"
CREATE UNIQUE INDEX `connection_shopifies_connection_connection_shopify_key` ON `connection_shopifies` (`connection_connection_shopify`);
-- Create index "connectionshopify_tenant_id" to table: "connection_shopifies"
CREATE INDEX `connectionshopify_tenant_id` ON `connection_shopifies` (`tenant_id`);
-- Create "consolidations" table
CREATE TABLE `consolidations` (`id` text NOT NULL, `public_id` text NOT NULL, `description` text NULL, `status` text NOT NULL DEFAULT ('Pending'), `created_at` datetime NULL, `tenant_id` text NOT NULL, `consolidation_delivery_option` text NULL, `shipment_consolidation` text NULL, PRIMARY KEY (`id`), CONSTRAINT `consolidations_tenants_tenant` FOREIGN KEY (`tenant_id`) REFERENCES `tenants` (`id`) ON DELETE NO ACTION, CONSTRAINT `consolidations_delivery_options_delivery_option` FOREIGN KEY (`consolidation_delivery_option`) REFERENCES `delivery_options` (`id`) ON DELETE SET NULL, CONSTRAINT `consolidations_shipments_consolidation` FOREIGN KEY (`shipment_consolidation`) REFERENCES `shipments` (`id`) ON DELETE SET NULL);
-- Create index "consolidations_shipment_consolidation_key" to table: "consolidations"
CREATE UNIQUE INDEX `consolidations_shipment_consolidation_key` ON `consolidations` (`shipment_consolidation`);
-- Create index "consolidation_tenant_id" to table: "consolidations"
CREATE INDEX `consolidation_tenant_id` ON `consolidations` (`tenant_id`);
-- Create "contacts" table
CREATE TABLE `contacts` (`id` text NOT NULL, `name` text NOT NULL, `surname` text NOT NULL, `email` text NOT NULL, `phone_number` text NOT NULL, `tenant_id` text NOT NULL, PRIMARY KEY (`id`), CONSTRAINT `contacts_tenants_tenant` FOREIGN KEY (`tenant_id`) REFERENCES `tenants` (`id`) ON DELETE NO ACTION);
-- Create index "contact_tenant_id" to table: "contacts"
CREATE INDEX `contact_tenant_id` ON `contacts` (`tenant_id`);
-- Create "countries" table
CREATE TABLE `countries` (`id` text NOT NULL, `label` text NOT NULL, `alpha_2` text NOT NULL, `alpha_3` text NOT NULL, `code` text NOT NULL, `region` text NOT NULL, PRIMARY KEY (`id`));
-- Create index "countries_label_key" to table: "countries"
CREATE UNIQUE INDEX `countries_label_key` ON `countries` (`label`);
-- Create index "countries_alpha_2_key" to table: "countries"
CREATE UNIQUE INDEX `countries_alpha_2_key` ON `countries` (`alpha_2`);
-- Create index "countries_alpha_3_key" to table: "countries"
CREATE UNIQUE INDEX `countries_alpha_3_key` ON `countries` (`alpha_3`);
-- Create index "countries_code_key" to table: "countries"
CREATE UNIQUE INDEX `countries_code_key` ON `countries` (`code`);
-- Create "country_harmonized_codes" table
CREATE TABLE `country_harmonized_codes` (`id` text NOT NULL, `code` text NOT NULL, `tenant_id` text NOT NULL, `country_harmonized_code_country` text NOT NULL, `inventory_item_country_harmonized_code` text NOT NULL, PRIMARY KEY (`id`), CONSTRAINT `country_harmonized_codes_tenants_tenant` FOREIGN KEY (`tenant_id`) REFERENCES `tenants` (`id`) ON DELETE NO ACTION, CONSTRAINT `country_harmonized_codes_countries_country` FOREIGN KEY (`country_harmonized_code_country`) REFERENCES `countries` (`id`) ON DELETE NO ACTION, CONSTRAINT `country_harmonized_codes_inventory_items_country_harmonized_code` FOREIGN KEY (`inventory_item_country_harmonized_code`) REFERENCES `inventory_items` (`id`) ON DELETE NO ACTION);
-- Create index "countryharmonizedcode_tenant_id" to table: "country_harmonized_codes"
CREATE INDEX `countryharmonizedcode_tenant_id` ON `country_harmonized_codes` (`tenant_id`);
-- Create "currencies" table
CREATE TABLE `currencies` (`id` text NOT NULL, `display` text NOT NULL, `currency_code` text NOT NULL DEFAULT ('DKK'), PRIMARY KEY (`id`));
-- Create index "currencies_display_key" to table: "currencies"
CREATE UNIQUE INDEX `currencies_display_key` ON `currencies` (`display`);
-- Create index "currency_currency_code" to table: "currencies"
CREATE UNIQUE INDEX `currency_currency_code` ON `currencies` (`currency_code`);
-- Create "delivery_options" table
CREATE TABLE `delivery_options` (`id` text NOT NULL, `archived_at` datetime NULL, `name` text NOT NULL, `sort_order` integer NOT NULL, `click_option_display_count` integer NULL DEFAULT (3), `description` text NULL, `click_collect` bool NULL DEFAULT (false), `override_sender_address` bool NULL DEFAULT (false), `override_return_address` bool NULL DEFAULT (false), `hide_delivery_option` bool NULL DEFAULT (false), `delivery_estimate_from` integer NULL, `delivery_estimate_to` integer NULL, `webshipper_integration` bool NOT NULL DEFAULT (false), `webshipper_id` integer NULL DEFAULT (1), `shipmondo_integration` bool NOT NULL DEFAULT (false), `shipmondo_delivery_option` text NULL, `customs_enabled` bool NOT NULL DEFAULT (false), `customs_signer` text NULL, `connection_delivery_option` text NOT NULL, `connection_default_delivery_option` text NULL, `tenant_id` text NOT NULL, `delivery_option_carrier` text NOT NULL, `delivery_option_carrier_service` text NOT NULL, `delivery_option_email_click_collect_at_store` text NULL, `delivery_option_default_packaging` text NULL, PRIMARY KEY (`id`), CONSTRAINT `delivery_options_connections_delivery_option` FOREIGN KEY (`connection_delivery_option`) REFERENCES `connections` (`id`) ON DELETE NO ACTION, CONSTRAINT `delivery_options_connections_default_delivery_option` FOREIGN KEY (`connection_default_delivery_option`) REFERENCES `connections` (`id`) ON DELETE SET NULL, CONSTRAINT `delivery_options_tenants_tenant` FOREIGN KEY (`tenant_id`) REFERENCES `tenants` (`id`) ON DELETE NO ACTION, CONSTRAINT `delivery_options_carriers_carrier` FOREIGN KEY (`delivery_option_carrier`) REFERENCES `carriers` (`id`) ON DELETE NO ACTION, CONSTRAINT `delivery_options_carrier_services_carrier_service` FOREIGN KEY (`delivery_option_carrier_service`) REFERENCES `carrier_services` (`id`) ON DELETE NO ACTION, CONSTRAINT `delivery_options_email_templates_email_click_collect_at_store` FOREIGN KEY (`delivery_option_email_click_collect_at_store`) REFERENCES `email_templates` (`id`) ON DELETE SET NULL, CONSTRAINT `delivery_options_packagings_default_packaging` FOREIGN KEY (`delivery_option_default_packaging`) REFERENCES `packagings` (`id`) ON DELETE SET NULL);
-- Create index "delivery_options_connection_default_delivery_option_key" to table: "delivery_options"
CREATE UNIQUE INDEX `delivery_options_connection_default_delivery_option_key` ON `delivery_options` (`connection_default_delivery_option`);
-- Create index "deliveryoption_tenant_id" to table: "delivery_options"
CREATE INDEX `deliveryoption_tenant_id` ON `delivery_options` (`tenant_id`);
-- Create "delivery_option_brings" table
CREATE TABLE `delivery_option_brings` (`id` text NOT NULL, `electronic_customs` bool NOT NULL DEFAULT (false), `delivery_option_delivery_option_bring` text NOT NULL, `tenant_id` text NOT NULL, PRIMARY KEY (`id`), CONSTRAINT `delivery_option_brings_delivery_options_delivery_option_bring` FOREIGN KEY (`delivery_option_delivery_option_bring`) REFERENCES `delivery_options` (`id`) ON DELETE NO ACTION, CONSTRAINT `delivery_option_brings_tenants_tenant` FOREIGN KEY (`tenant_id`) REFERENCES `tenants` (`id`) ON DELETE NO ACTION);
-- Create index "delivery_option_brings_delivery_option_delivery_option_bring_key" to table: "delivery_option_brings"
CREATE UNIQUE INDEX `delivery_option_brings_delivery_option_delivery_option_bring_key` ON `delivery_option_brings` (`delivery_option_delivery_option_bring`);
-- Create index "deliveryoptionbring_tenant_id" to table: "delivery_option_brings"
CREATE INDEX `deliveryoptionbring_tenant_id` ON `delivery_option_brings` (`tenant_id`);
-- Create "delivery_option_da_os" table
CREATE TABLE `delivery_option_da_os` (`id` text NOT NULL, `delivery_option_delivery_option_dao` text NOT NULL, `tenant_id` text NOT NULL, PRIMARY KEY (`id`), CONSTRAINT `delivery_option_da_os_delivery_options_delivery_option_dao` FOREIGN KEY (`delivery_option_delivery_option_dao`) REFERENCES `delivery_options` (`id`) ON DELETE NO ACTION, CONSTRAINT `delivery_option_da_os_tenants_tenant` FOREIGN KEY (`tenant_id`) REFERENCES `tenants` (`id`) ON DELETE NO ACTION);
-- Create index "delivery_option_da_os_delivery_option_delivery_option_dao_key" to table: "delivery_option_da_os"
CREATE UNIQUE INDEX `delivery_option_da_os_delivery_option_delivery_option_dao_key` ON `delivery_option_da_os` (`delivery_option_delivery_option_dao`);
-- Create index "deliveryoptiondao_tenant_id" to table: "delivery_option_da_os"
CREATE INDEX `deliveryoptiondao_tenant_id` ON `delivery_option_da_os` (`tenant_id`);
-- Create "delivery_option_dfs" table
CREATE TABLE `delivery_option_dfs` (`id` text NOT NULL, `delivery_option_delivery_option_df` text NOT NULL, `tenant_id` text NOT NULL, PRIMARY KEY (`id`), CONSTRAINT `delivery_option_dfs_delivery_options_delivery_option_df` FOREIGN KEY (`delivery_option_delivery_option_df`) REFERENCES `delivery_options` (`id`) ON DELETE NO ACTION, CONSTRAINT `delivery_option_dfs_tenants_tenant` FOREIGN KEY (`tenant_id`) REFERENCES `tenants` (`id`) ON DELETE NO ACTION);
-- Create index "delivery_option_dfs_delivery_option_delivery_option_df_key" to table: "delivery_option_dfs"
CREATE UNIQUE INDEX `delivery_option_dfs_delivery_option_delivery_option_df_key` ON `delivery_option_dfs` (`delivery_option_delivery_option_df`);
-- Create index "deliveryoptiondf_tenant_id" to table: "delivery_option_dfs"
CREATE INDEX `deliveryoptiondf_tenant_id` ON `delivery_option_dfs` (`tenant_id`);
-- Create "delivery_option_ds_vs" table
CREATE TABLE `delivery_option_ds_vs` (`id` text NOT NULL, `delivery_option_delivery_option_dsv` text NOT NULL, `tenant_id` text NOT NULL, PRIMARY KEY (`id`), CONSTRAINT `delivery_option_ds_vs_delivery_options_delivery_option_dsv` FOREIGN KEY (`delivery_option_delivery_option_dsv`) REFERENCES `delivery_options` (`id`) ON DELETE NO ACTION, CONSTRAINT `delivery_option_ds_vs_tenants_tenant` FOREIGN KEY (`tenant_id`) REFERENCES `tenants` (`id`) ON DELETE NO ACTION);
-- Create index "delivery_option_ds_vs_delivery_option_delivery_option_dsv_key" to table: "delivery_option_ds_vs"
CREATE UNIQUE INDEX `delivery_option_ds_vs_delivery_option_delivery_option_dsv_key` ON `delivery_option_ds_vs` (`delivery_option_delivery_option_dsv`);
-- Create index "deliveryoptiondsv_tenant_id" to table: "delivery_option_ds_vs"
CREATE INDEX `deliveryoptiondsv_tenant_id` ON `delivery_option_ds_vs` (`tenant_id`);
-- Create "delivery_option_easy_posts" table
CREATE TABLE `delivery_option_easy_posts` (`id` text NOT NULL, `delivery_option_delivery_option_easy_post` text NOT NULL, `tenant_id` text NOT NULL, PRIMARY KEY (`id`), CONSTRAINT `delivery_option_easy_posts_delivery_options_delivery_option_easy_post` FOREIGN KEY (`delivery_option_delivery_option_easy_post`) REFERENCES `delivery_options` (`id`) ON DELETE NO ACTION, CONSTRAINT `delivery_option_easy_posts_tenants_tenant` FOREIGN KEY (`tenant_id`) REFERENCES `tenants` (`id`) ON DELETE NO ACTION);
-- Create index "delivery_option_easy_posts_delivery_option_delivery_option_easy_post_key" to table: "delivery_option_easy_posts"
CREATE UNIQUE INDEX `delivery_option_easy_posts_delivery_option_delivery_option_easy_post_key` ON `delivery_option_easy_posts` (`delivery_option_delivery_option_easy_post`);
-- Create index "deliveryoptioneasypost_tenant_id" to table: "delivery_option_easy_posts"
CREATE INDEX `deliveryoptioneasypost_tenant_id` ON `delivery_option_easy_posts` (`tenant_id`);
-- Create "delivery_option_gl_ss" table
CREATE TABLE `delivery_option_gl_ss` (`id` text NOT NULL, `delivery_option_delivery_option_gls` text NOT NULL, `tenant_id` text NOT NULL, PRIMARY KEY (`id`), CONSTRAINT `delivery_option_gl_ss_delivery_options_delivery_option_gls` FOREIGN KEY (`delivery_option_delivery_option_gls`) REFERENCES `delivery_options` (`id`) ON DELETE NO ACTION, CONSTRAINT `delivery_option_gl_ss_tenants_tenant` FOREIGN KEY (`tenant_id`) REFERENCES `tenants` (`id`) ON DELETE NO ACTION);
-- Create index "delivery_option_gl_ss_delivery_option_delivery_option_gls_key" to table: "delivery_option_gl_ss"
CREATE UNIQUE INDEX `delivery_option_gl_ss_delivery_option_delivery_option_gls_key` ON `delivery_option_gl_ss` (`delivery_option_delivery_option_gls`);
-- Create index "deliveryoptiongls_tenant_id" to table: "delivery_option_gl_ss"
CREATE INDEX `deliveryoptiongls_tenant_id` ON `delivery_option_gl_ss` (`tenant_id`);
-- Create "delivery_option_post_nords" table
CREATE TABLE `delivery_option_post_nords` (`id` text NOT NULL, `format_zpl` bool NOT NULL DEFAULT (true), `delivery_option_delivery_option_post_nord` text NOT NULL, `tenant_id` text NOT NULL, PRIMARY KEY (`id`), CONSTRAINT `delivery_option_post_nords_delivery_options_delivery_option_post_nord` FOREIGN KEY (`delivery_option_delivery_option_post_nord`) REFERENCES `delivery_options` (`id`) ON DELETE NO ACTION, CONSTRAINT `delivery_option_post_nords_tenants_tenant` FOREIGN KEY (`tenant_id`) REFERENCES `tenants` (`id`) ON DELETE NO ACTION);
-- Create index "delivery_option_post_nords_delivery_option_delivery_option_post_nord_key" to table: "delivery_option_post_nords"
CREATE UNIQUE INDEX `delivery_option_post_nords_delivery_option_delivery_option_post_nord_key` ON `delivery_option_post_nords` (`delivery_option_delivery_option_post_nord`);
-- Create index "deliveryoptionpostnord_tenant_id" to table: "delivery_option_post_nords"
CREATE INDEX `deliveryoptionpostnord_tenant_id` ON `delivery_option_post_nords` (`tenant_id`);
-- Create "delivery_option_usp_ss" table
CREATE TABLE `delivery_option_usp_ss` (`id` text NOT NULL, `format_zpl` bool NOT NULL DEFAULT (true), `delivery_option_delivery_option_usps` text NOT NULL, `tenant_id` text NOT NULL, PRIMARY KEY (`id`), CONSTRAINT `delivery_option_usp_ss_delivery_options_delivery_option_usps` FOREIGN KEY (`delivery_option_delivery_option_usps`) REFERENCES `delivery_options` (`id`) ON DELETE NO ACTION, CONSTRAINT `delivery_option_usp_ss_tenants_tenant` FOREIGN KEY (`tenant_id`) REFERENCES `tenants` (`id`) ON DELETE NO ACTION);
-- Create index "delivery_option_usp_ss_delivery_option_delivery_option_usps_key" to table: "delivery_option_usp_ss"
CREATE UNIQUE INDEX `delivery_option_usp_ss_delivery_option_delivery_option_usps_key` ON `delivery_option_usp_ss` (`delivery_option_delivery_option_usps`);
-- Create index "deliveryoptionusps_tenant_id" to table: "delivery_option_usp_ss"
CREATE INDEX `deliveryoptionusps_tenant_id` ON `delivery_option_usp_ss` (`tenant_id`);
-- Create "delivery_rules" table
CREATE TABLE `delivery_rules` (`id` text NOT NULL, `name` text NOT NULL, `price` real NOT NULL DEFAULT (20), `delivery_option_delivery_rule` text NULL, `tenant_id` text NOT NULL, `delivery_rule_currency` text NULL, PRIMARY KEY (`id`), CONSTRAINT `delivery_rules_delivery_options_delivery_rule` FOREIGN KEY (`delivery_option_delivery_rule`) REFERENCES `delivery_options` (`id`) ON DELETE SET NULL, CONSTRAINT `delivery_rules_tenants_tenant` FOREIGN KEY (`tenant_id`) REFERENCES `tenants` (`id`) ON DELETE NO ACTION, CONSTRAINT `delivery_rules_currencies_currency` FOREIGN KEY (`delivery_rule_currency`) REFERENCES `currencies` (`id`) ON DELETE SET NULL);
-- Create index "deliveryrule_tenant_id" to table: "delivery_rules"
CREATE INDEX `deliveryrule_tenant_id` ON `delivery_rules` (`tenant_id`);
-- Create "delivery_rule_constraints" table
CREATE TABLE `delivery_rule_constraints` (`id` text NOT NULL, `property_type` text NOT NULL, `comparison` text NOT NULL, `selected_value` json NOT NULL, `tenant_id` text NOT NULL, `delivery_rule_constraint_group_delivery_rule_constraints` text NOT NULL, PRIMARY KEY (`id`), CONSTRAINT `delivery_rule_constraints_tenants_tenant` FOREIGN KEY (`tenant_id`) REFERENCES `tenants` (`id`) ON DELETE NO ACTION, CONSTRAINT `delivery_rule_constraints_delivery_rule_constraint_groups_delivery_rule_constraints` FOREIGN KEY (`delivery_rule_constraint_group_delivery_rule_constraints`) REFERENCES `delivery_rule_constraint_groups` (`id`) ON DELETE NO ACTION);
-- Create index "deliveryruleconstraint_tenant_id" to table: "delivery_rule_constraints"
CREATE INDEX `deliveryruleconstraint_tenant_id` ON `delivery_rule_constraints` (`tenant_id`);
-- Create "delivery_rule_constraint_groups" table
CREATE TABLE `delivery_rule_constraint_groups` (`id` text NOT NULL, `constraint_logic` text NOT NULL DEFAULT ('and'), `delivery_rule_delivery_rule_constraint_group` text NOT NULL, `tenant_id` text NOT NULL, PRIMARY KEY (`id`), CONSTRAINT `delivery_rule_constraint_groups_delivery_rules_delivery_rule_constraint_group` FOREIGN KEY (`delivery_rule_delivery_rule_constraint_group`) REFERENCES `delivery_rules` (`id`) ON DELETE NO ACTION, CONSTRAINT `delivery_rule_constraint_groups_tenants_tenant` FOREIGN KEY (`tenant_id`) REFERENCES `tenants` (`id`) ON DELETE NO ACTION);
-- Create index "deliveryruleconstraintgroup_tenant_id" to table: "delivery_rule_constraint_groups"
CREATE INDEX `deliveryruleconstraintgroup_tenant_id` ON `delivery_rule_constraint_groups` (`tenant_id`);
-- Create "documents" table
CREATE TABLE `documents` (`id` text NOT NULL, `name` text NOT NULL, `html_template` text NULL, `html_header` text NULL, `html_footer` text NULL, `last_base64_pdf` text NULL, `merge_type` text NOT NULL DEFAULT ('Orders'), `paper_size` text NOT NULL DEFAULT ('A4'), `start_at` datetime NOT NULL, `end_at` datetime NOT NULL, `created_at` datetime NOT NULL, `tenant_id` text NOT NULL, `document_carrier_brand` text NULL, PRIMARY KEY (`id`), CONSTRAINT `documents_tenants_tenant` FOREIGN KEY (`tenant_id`) REFERENCES `tenants` (`id`) ON DELETE NO ACTION, CONSTRAINT `documents_carrier_brands_carrier_brand` FOREIGN KEY (`document_carrier_brand`) REFERENCES `carrier_brands` (`id`) ON DELETE SET NULL);
-- Create index "document_tenant_id" to table: "documents"
CREATE INDEX `document_tenant_id` ON `documents` (`tenant_id`);
-- Create "document_files" table
CREATE TABLE `document_files` (`id` text NOT NULL, `created_at` datetime NOT NULL, `doc_type` text NOT NULL, `data_pdf_base64` text NULL, `data_zpl_base64` text NULL, `colli_document_file` text NULL, `tenant_id` text NOT NULL, PRIMARY KEY (`id`), CONSTRAINT `document_files_collis_document_file` FOREIGN KEY (`colli_document_file`) REFERENCES `collis` (`id`) ON DELETE SET NULL, CONSTRAINT `document_files_tenants_tenant` FOREIGN KEY (`tenant_id`) REFERENCES `tenants` (`id`) ON DELETE NO ACTION);
-- Create index "documentfile_tenant_id" to table: "document_files"
CREATE INDEX `documentfile_tenant_id` ON `document_files` (`tenant_id`);
-- Create "email_templates" table
CREATE TABLE `email_templates` (`id` text NOT NULL, `name` text NOT NULL, `subject` text NOT NULL DEFAULT (''), `html_template` text NOT NULL DEFAULT (''), `merge_type` text NOT NULL DEFAULT ('return_colli_label'), `created_at` datetime NULL, `updated_at` datetime NOT NULL, `tenant_id` text NOT NULL, PRIMARY KEY (`id`), CONSTRAINT `email_templates_tenants_tenant` FOREIGN KEY (`tenant_id`) REFERENCES `tenants` (`id`) ON DELETE NO ACTION);
-- Create index "emailtemplate_tenant_id" to table: "email_templates"
CREATE INDEX `emailtemplate_tenant_id` ON `email_templates` (`tenant_id`);
-- Create "hypothesis_tests" table
CREATE TABLE `hypothesis_tests` (`id` text NOT NULL, `name` text NOT NULL, `active` bool NOT NULL DEFAULT (false), `tenant_id` text NOT NULL, `hypothesis_test_connection` text NOT NULL, PRIMARY KEY (`id`), CONSTRAINT `hypothesis_tests_tenants_tenant` FOREIGN KEY (`tenant_id`) REFERENCES `tenants` (`id`) ON DELETE NO ACTION, CONSTRAINT `hypothesis_tests_connections_connection` FOREIGN KEY (`hypothesis_test_connection`) REFERENCES `connections` (`id`) ON DELETE NO ACTION);
-- Create index "hypothesistest_tenant_id" to table: "hypothesis_tests"
CREATE INDEX `hypothesistest_tenant_id` ON `hypothesis_tests` (`tenant_id`);
-- Create "hypothesis_test_delivery_options" table
CREATE TABLE `hypothesis_test_delivery_options` (`id` text NOT NULL, `randomize_within_group_sort` bool NOT NULL DEFAULT (false), `by_interval_rotation` bool NOT NULL DEFAULT (false), `rotation_interval_hours` integer NOT NULL DEFAULT (6), `by_order` bool NOT NULL DEFAULT (false), `hypothesis_test_hypothesis_test_delivery_option` text NOT NULL, `tenant_id` text NOT NULL, PRIMARY KEY (`id`), CONSTRAINT `hypothesis_test_delivery_options_hypothesis_tests_hypothesis_test_delivery_option` FOREIGN KEY (`hypothesis_test_hypothesis_test_delivery_option`) REFERENCES `hypothesis_tests` (`id`) ON DELETE NO ACTION, CONSTRAINT `hypothesis_test_delivery_options_tenants_tenant` FOREIGN KEY (`tenant_id`) REFERENCES `tenants` (`id`) ON DELETE NO ACTION);
-- Create index "hypothesis_test_delivery_options_hypothesis_test_hypothesis_test_delivery_option_key" to table: "hypothesis_test_delivery_options"
CREATE UNIQUE INDEX `hypothesis_test_delivery_options_hypothesis_test_hypothesis_test_delivery_option_key` ON `hypothesis_test_delivery_options` (`hypothesis_test_hypothesis_test_delivery_option`);
-- Create index "hypothesistestdeliveryoption_tenant_id" to table: "hypothesis_test_delivery_options"
CREATE INDEX `hypothesistestdeliveryoption_tenant_id` ON `hypothesis_test_delivery_options` (`tenant_id`);
-- Create "hypothesis_test_delivery_option_lookups" table
CREATE TABLE `hypothesis_test_delivery_option_lookups` (`id` text NOT NULL, `tenant_id` text NOT NULL, `hypothesis_test_delivery_option_lookup_delivery_option` text NOT NULL, `hypothesis_test_delivery_option_request_hypothesis_test_delivery_option_lookup` text NOT NULL, PRIMARY KEY (`id`), CONSTRAINT `hypothesis_test_delivery_option_lookups_tenants_tenant` FOREIGN KEY (`tenant_id`) REFERENCES `tenants` (`id`) ON DELETE NO ACTION, CONSTRAINT `hypothesis_test_delivery_option_lookups_delivery_options_delivery_option` FOREIGN KEY (`hypothesis_test_delivery_option_lookup_delivery_option`) REFERENCES `delivery_options` (`id`) ON DELETE NO ACTION, CONSTRAINT `hypothesis_test_delivery_option_lookups_hypothesis_test_delivery_option_requests_hypothesis_test_delivery_option_lookup` FOREIGN KEY (`hypothesis_test_delivery_option_request_hypothesis_test_delivery_option_lookup`) REFERENCES `hypothesis_test_delivery_option_requests` (`id`) ON DELETE NO ACTION);
-- Create index "hypothesistestdeliveryoptionlookup_tenant_id" to table: "hypothesis_test_delivery_option_lookups"
CREATE INDEX `hypothesistestdeliveryoptionlookup_tenant_id` ON `hypothesis_test_delivery_option_lookups` (`tenant_id`);
-- Create index "hypothesistestdeliveryoptionlookup_hypothesis_test_delivery_option_lookup_delivery_option_hypothesis_test_delivery_option_request_hypothesis_test_delivery_option_lookup" to table: "hypothesis_test_delivery_option_lookups"
CREATE UNIQUE INDEX `hypothesistestdeliveryoptionlookup_hypothesis_test_delivery_option_lookup_delivery_option_hypothesis_test_delivery_option_request_hypothesis_test_delivery_option_lookup` ON `hypothesis_test_delivery_option_lookups` (`hypothesis_test_delivery_option_lookup_delivery_option`, `hypothesis_test_delivery_option_request_hypothesis_test_delivery_option_lookup`);
-- Create "hypothesis_test_delivery_option_requests" table
CREATE TABLE `hypothesis_test_delivery_option_requests` (`id` text NOT NULL, `order_hash` text NOT NULL, `shipping_address_hash` text NOT NULL, `is_control_group` bool NOT NULL, `request_count` integer NOT NULL, `created_at` datetime NOT NULL, `last_requested_at` datetime NOT NULL, `tenant_id` text NOT NULL, `hypothesis_test_delivery_option_request_hypothesis_test_delivery_option` text NOT NULL, PRIMARY KEY (`id`), CONSTRAINT `hypothesis_test_delivery_option_requests_tenants_tenant` FOREIGN KEY (`tenant_id`) REFERENCES `tenants` (`id`) ON DELETE NO ACTION, CONSTRAINT `hypothesis_test_delivery_option_requests_hypothesis_test_delivery_options_hypothesis_test_delivery_option` FOREIGN KEY (`hypothesis_test_delivery_option_request_hypothesis_test_delivery_option`) REFERENCES `hypothesis_test_delivery_options` (`id`) ON DELETE NO ACTION);
-- Create index "hypothesistestdeliveryoptionrequest_tenant_id" to table: "hypothesis_test_delivery_option_requests"
CREATE INDEX `hypothesistestdeliveryoptionrequest_tenant_id` ON `hypothesis_test_delivery_option_requests` (`tenant_id`);
-- Create "inventory_items" table
CREATE TABLE `inventory_items` (`id` text NOT NULL, `external_id` text NULL, `code` text NULL, `sku` text NULL, `tenant_id` text NOT NULL, `inventory_item_country_of_origin` text NULL, `product_variant_inventory_item` text NOT NULL, PRIMARY KEY (`id`), CONSTRAINT `inventory_items_tenants_tenant` FOREIGN KEY (`tenant_id`) REFERENCES `tenants` (`id`) ON DELETE NO ACTION, CONSTRAINT `inventory_items_countries_country_of_origin` FOREIGN KEY (`inventory_item_country_of_origin`) REFERENCES `countries` (`id`) ON DELETE SET NULL, CONSTRAINT `inventory_items_product_variants_inventory_item` FOREIGN KEY (`product_variant_inventory_item`) REFERENCES `product_variants` (`id`) ON DELETE NO ACTION);
-- Create index "inventory_items_product_variant_inventory_item_key" to table: "inventory_items"
CREATE UNIQUE INDEX `inventory_items_product_variant_inventory_item_key` ON `inventory_items` (`product_variant_inventory_item`);
-- Create index "inventoryitem_tenant_id" to table: "inventory_items"
CREATE INDEX `inventoryitem_tenant_id` ON `inventory_items` (`tenant_id`);
-- Create "languages" table
CREATE TABLE `languages` (`id` text NOT NULL, `label` text NOT NULL, `internal_id` text NOT NULL DEFAULT ('EN'), PRIMARY KEY (`id`));
-- Create index "languages_label_key" to table: "languages"
CREATE UNIQUE INDEX `languages_label_key` ON `languages` (`label`);
-- Create "locations" table
CREATE TABLE `locations` (`id` text NOT NULL, `name` text NOT NULL, `tenant_id` text NOT NULL, `location_address` text NOT NULL, PRIMARY KEY (`id`), CONSTRAINT `locations_tenants_tenant` FOREIGN KEY (`tenant_id`) REFERENCES `tenants` (`id`) ON DELETE NO ACTION, CONSTRAINT `locations_addresses_address` FOREIGN KEY (`location_address`) REFERENCES `addresses` (`id`) ON DELETE NO ACTION);
-- Create index "location_tenant_id" to table: "locations"
CREATE INDEX `location_tenant_id` ON `locations` (`tenant_id`);
-- Create index "location_name_tenant_id" to table: "locations"
CREATE UNIQUE INDEX `location_name_tenant_id` ON `locations` (`name`, `tenant_id`);
-- Create "location_tags" table
CREATE TABLE `location_tags` (`id` text NOT NULL, `label` text NOT NULL, `internal_id` text NOT NULL, PRIMARY KEY (`id`));
-- Create index "location_tags_label_key" to table: "location_tags"
CREATE UNIQUE INDEX `location_tags_label_key` ON `location_tags` (`label`);
-- Create index "location_tags_internal_id_key" to table: "location_tags"
CREATE UNIQUE INDEX `location_tags_internal_id_key` ON `location_tags` (`internal_id`);
-- Create "notifications" table
CREATE TABLE `notifications` (`id` text NOT NULL, `name` text NOT NULL, `active` bool NOT NULL DEFAULT (true), `tenant_id` text NOT NULL, `notification_connection` text NOT NULL, `notification_email_template` text NOT NULL, PRIMARY KEY (`id`), CONSTRAINT `notifications_tenants_tenant` FOREIGN KEY (`tenant_id`) REFERENCES `tenants` (`id`) ON DELETE NO ACTION, CONSTRAINT `notifications_connections_connection` FOREIGN KEY (`notification_connection`) REFERENCES `connections` (`id`) ON DELETE NO ACTION, CONSTRAINT `notifications_email_templates_email_template` FOREIGN KEY (`notification_email_template`) REFERENCES `email_templates` (`id`) ON DELETE NO ACTION);
-- Create index "notification_tenant_id" to table: "notifications"
CREATE INDEX `notification_tenant_id` ON `notifications` (`tenant_id`);
-- Create "otk_requests" table
CREATE TABLE `otk_requests` (`id` text NOT NULL, `otk` text NOT NULL, `tenant_id` text NOT NULL, `user_otk_requests` text NULL, PRIMARY KEY (`id`), CONSTRAINT `otk_requests_tenants_tenant` FOREIGN KEY (`tenant_id`) REFERENCES `tenants` (`id`) ON DELETE NO ACTION, CONSTRAINT `otk_requests_users_otk_requests` FOREIGN KEY (`user_otk_requests`) REFERENCES `users` (`id`) ON DELETE SET NULL);
-- Create index "otkrequests_tenant_id" to table: "otk_requests"
CREATE INDEX `otkrequests_tenant_id` ON `otk_requests` (`tenant_id`);
-- Create "orders" table
CREATE TABLE `orders` (`id` text NOT NULL, `order_public_id` text NOT NULL, `external_id` text NULL, `comment_internal` text NULL, `comment_external` text NULL, `created_at` datetime NOT NULL, `email_sync_confirmation_at` datetime NULL, `status` text NOT NULL, `connection_orders` text NOT NULL, `consolidation_orders` text NULL, `hypothesis_test_delivery_option_request_order` text NULL, `tenant_id` text NOT NULL, `pallet_orders` text NULL, PRIMARY KEY (`id`), CONSTRAINT `orders_connections_orders` FOREIGN KEY (`connection_orders`) REFERENCES `connections` (`id`) ON DELETE NO ACTION, CONSTRAINT `orders_consolidations_orders` FOREIGN KEY (`consolidation_orders`) REFERENCES `consolidations` (`id`) ON DELETE SET NULL, CONSTRAINT `orders_hypothesis_test_delivery_option_requests_order` FOREIGN KEY (`hypothesis_test_delivery_option_request_order`) REFERENCES `hypothesis_test_delivery_option_requests` (`id`) ON DELETE SET NULL, CONSTRAINT `orders_tenants_tenant` FOREIGN KEY (`tenant_id`) REFERENCES `tenants` (`id`) ON DELETE NO ACTION, CONSTRAINT `orders_pallets_orders` FOREIGN KEY (`pallet_orders`) REFERENCES `pallets` (`id`) ON DELETE SET NULL);
-- Create index "orders_hypothesis_test_delivery_option_request_order_key" to table: "orders"
CREATE UNIQUE INDEX `orders_hypothesis_test_delivery_option_request_order_key` ON `orders` (`hypothesis_test_delivery_option_request_order`);
-- Create index "order_tenant_id" to table: "orders"
CREATE INDEX `order_tenant_id` ON `orders` (`tenant_id`);
-- Create index "order_order_public_id_tenant_id" to table: "orders"
CREATE UNIQUE INDEX `order_order_public_id_tenant_id` ON `orders` (`order_public_id`, `tenant_id`);
-- Create index "order_external_id_tenant_id" to table: "orders"
CREATE UNIQUE INDEX `order_external_id_tenant_id` ON `orders` (`external_id`, `tenant_id`);
-- Create index "order_created_at" to table: "orders"
CREATE INDEX `order_created_at` ON `orders` (`created_at`);
-- Create "order_histories" table
CREATE TABLE `order_histories` (`id` text NOT NULL, `description` text NOT NULL, `type` text NOT NULL, `change_history_order_history` text NOT NULL, `order_order_history` text NOT NULL, `tenant_id` text NOT NULL, PRIMARY KEY (`id`), CONSTRAINT `order_histories_change_histories_order_history` FOREIGN KEY (`change_history_order_history`) REFERENCES `change_histories` (`id`) ON DELETE NO ACTION, CONSTRAINT `order_histories_orders_order_history` FOREIGN KEY (`order_order_history`) REFERENCES `orders` (`id`) ON DELETE NO ACTION, CONSTRAINT `order_histories_tenants_tenant` FOREIGN KEY (`tenant_id`) REFERENCES `tenants` (`id`) ON DELETE NO ACTION);
-- Create index "orderhistory_tenant_id" to table: "order_histories"
CREATE INDEX `orderhistory_tenant_id` ON `order_histories` (`tenant_id`);
-- Create "order_lines" table
CREATE TABLE `order_lines` (`id` text NOT NULL, `unit_price` real NOT NULL, `discount_allocation_amount` real NOT NULL, `external_id` text NULL, `units` integer NOT NULL, `created_at` datetime NULL, `updated_at` datetime NOT NULL, `colli_id` text NOT NULL, `tenant_id` text NOT NULL, `product_variant_id` text NOT NULL, `order_line_currency` text NOT NULL, PRIMARY KEY (`id`), CONSTRAINT `order_lines_collis_order_lines` FOREIGN KEY (`colli_id`) REFERENCES `collis` (`id`) ON DELETE NO ACTION, CONSTRAINT `order_lines_tenants_tenant` FOREIGN KEY (`tenant_id`) REFERENCES `tenants` (`id`) ON DELETE NO ACTION, CONSTRAINT `order_lines_product_variants_product_variant` FOREIGN KEY (`product_variant_id`) REFERENCES `product_variants` (`id`) ON DELETE NO ACTION, CONSTRAINT `order_lines_currencies_currency` FOREIGN KEY (`order_line_currency`) REFERENCES `currencies` (`id`) ON DELETE NO ACTION);
-- Create index "orderline_tenant_id" to table: "order_lines"
CREATE INDEX `orderline_tenant_id` ON `order_lines` (`tenant_id`);
-- Create "order_senders" table
CREATE TABLE `order_senders` (`id` text NOT NULL, `uniqueness_id` text NULL, `first_name` text NOT NULL, `last_name` text NOT NULL, `email` text NOT NULL, `phone_number` text NOT NULL, `vat_number` text NOT NULL, `tenant_id` text NOT NULL, PRIMARY KEY (`id`), CONSTRAINT `order_senders_tenants_tenant` FOREIGN KEY (`tenant_id`) REFERENCES `tenants` (`id`) ON DELETE NO ACTION);
-- Create index "order_senders_uniqueness_id_key" to table: "order_senders"
CREATE UNIQUE INDEX `order_senders_uniqueness_id_key` ON `order_senders` (`uniqueness_id`);
-- Create index "ordersender_tenant_id" to table: "order_senders"
CREATE INDEX `ordersender_tenant_id` ON `order_senders` (`tenant_id`);
-- Create "packagings" table
CREATE TABLE `packagings` (`id` text NOT NULL, `archived_at` datetime NULL, `name` text NOT NULL, `height_cm` integer NOT NULL, `width_cm` integer NOT NULL, `length_cm` integer NOT NULL, `tenant_id` text NOT NULL, `packaging_carrier_brand` text NULL, PRIMARY KEY (`id`), CONSTRAINT `packagings_tenants_tenant` FOREIGN KEY (`tenant_id`) REFERENCES `tenants` (`id`) ON DELETE NO ACTION, CONSTRAINT `packagings_carrier_brands_carrier_brand` FOREIGN KEY (`packaging_carrier_brand`) REFERENCES `carrier_brands` (`id`) ON DELETE SET NULL);
-- Create index "packaging_tenant_id" to table: "packagings"
CREATE INDEX `packaging_tenant_id` ON `packagings` (`tenant_id`);
-- Create "packaging_dfs" table
CREATE TABLE `packaging_dfs` (`id` text NOT NULL, `api_type` text NOT NULL, `max_weight` real NULL, `min_weight` real NULL, `stackable` bool NOT NULL DEFAULT (false), `packaging_packaging_df` text NOT NULL, `tenant_id` text NOT NULL, PRIMARY KEY (`id`), CONSTRAINT `packaging_dfs_packagings_packaging_df` FOREIGN KEY (`packaging_packaging_df`) REFERENCES `packagings` (`id`) ON DELETE NO ACTION, CONSTRAINT `packaging_dfs_tenants_tenant` FOREIGN KEY (`tenant_id`) REFERENCES `tenants` (`id`) ON DELETE NO ACTION);
-- Create index "packaging_dfs_packaging_packaging_df_key" to table: "packaging_dfs"
CREATE UNIQUE INDEX `packaging_dfs_packaging_packaging_df_key` ON `packaging_dfs` (`packaging_packaging_df`);
-- Create index "packagingdf_tenant_id" to table: "packaging_dfs"
CREATE INDEX `packagingdf_tenant_id` ON `packaging_dfs` (`tenant_id`);
-- Create "packaging_usp_ss" table
CREATE TABLE `packaging_usp_ss` (`id` text NOT NULL, `packaging_packaging_usps` text NOT NULL, `tenant_id` text NOT NULL, `packaging_usps_packaging_usps_rate_indicator` text NOT NULL, `packaging_usps_packaging_usps_processing_category` text NOT NULL, PRIMARY KEY (`id`), CONSTRAINT `packaging_usp_ss_packagings_packaging_usps` FOREIGN KEY (`packaging_packaging_usps`) REFERENCES `packagings` (`id`) ON DELETE NO ACTION, CONSTRAINT `packaging_usp_ss_tenants_tenant` FOREIGN KEY (`tenant_id`) REFERENCES `tenants` (`id`) ON DELETE NO ACTION, CONSTRAINT `packaging_usp_ss_packaging_usps_rate_indicators_packaging_usps_rate_indicator` FOREIGN KEY (`packaging_usps_packaging_usps_rate_indicator`) REFERENCES `packaging_usps_rate_indicators` (`id`) ON DELETE NO ACTION, CONSTRAINT `packaging_usp_ss_packaging_usps_processing_categories_packaging_usps_processing_category` FOREIGN KEY (`packaging_usps_packaging_usps_processing_category`) REFERENCES `packaging_usps_processing_categories` (`id`) ON DELETE NO ACTION);
-- Create index "packaging_usp_ss_packaging_packaging_usps_key" to table: "packaging_usp_ss"
CREATE UNIQUE INDEX `packaging_usp_ss_packaging_packaging_usps_key` ON `packaging_usp_ss` (`packaging_packaging_usps`);
-- Create index "packagingusps_tenant_id" to table: "packaging_usp_ss"
CREATE INDEX `packagingusps_tenant_id` ON `packaging_usp_ss` (`tenant_id`);
-- Create "packaging_usps_processing_categories" table
CREATE TABLE `packaging_usps_processing_categories` (`id` text NOT NULL, `name` text NOT NULL, `processing_category` text NOT NULL, PRIMARY KEY (`id`));
-- Create "packaging_usps_rate_indicators" table
CREATE TABLE `packaging_usps_rate_indicators` (`id` text NOT NULL, `code` text NOT NULL, `name` text NOT NULL, PRIMARY KEY (`id`));
-- Create "pallets" table
CREATE TABLE `pallets` (`id` text NOT NULL, `public_id` text NOT NULL, `description` text NOT NULL, `consolidation_pallets` text NOT NULL, `tenant_id` text NOT NULL, `pallet_packaging` text NULL, PRIMARY KEY (`id`), CONSTRAINT `pallets_consolidations_pallets` FOREIGN KEY (`consolidation_pallets`) REFERENCES `consolidations` (`id`) ON DELETE NO ACTION, CONSTRAINT `pallets_tenants_tenant` FOREIGN KEY (`tenant_id`) REFERENCES `tenants` (`id`) ON DELETE NO ACTION, CONSTRAINT `pallets_packagings_packaging` FOREIGN KEY (`pallet_packaging`) REFERENCES `packagings` (`id`) ON DELETE SET NULL);
-- Create index "pallet_tenant_id" to table: "pallets"
CREATE INDEX `pallet_tenant_id` ON `pallets` (`tenant_id`);
-- Create "parcel_shops" table
CREATE TABLE `parcel_shops` (`id` text NOT NULL, `name` text NOT NULL, `last_updated` datetime NOT NULL, `parcel_shop_carrier_brand` text NOT NULL, PRIMARY KEY (`id`), CONSTRAINT `parcel_shops_carrier_brands_carrier_brand` FOREIGN KEY (`parcel_shop_carrier_brand`) REFERENCES `carrier_brands` (`id`) ON DELETE NO ACTION);
-- Create "parcel_shop_brings" table
CREATE TABLE `parcel_shop_brings` (`id` text NOT NULL, `point_type` text NOT NULL, `bring_id` text NOT NULL, `parcel_shop_parcel_shop_bring` text NOT NULL, PRIMARY KEY (`id`), CONSTRAINT `parcel_shop_brings_parcel_shops_parcel_shop_bring` FOREIGN KEY (`parcel_shop_parcel_shop_bring`) REFERENCES `parcel_shops` (`id`) ON DELETE NO ACTION);
-- Create index "parcel_shop_brings_bring_id_key" to table: "parcel_shop_brings"
CREATE UNIQUE INDEX `parcel_shop_brings_bring_id_key` ON `parcel_shop_brings` (`bring_id`);
-- Create index "parcel_shop_brings_parcel_shop_parcel_shop_bring_key" to table: "parcel_shop_brings"
CREATE UNIQUE INDEX `parcel_shop_brings_parcel_shop_parcel_shop_bring_key` ON `parcel_shop_brings` (`parcel_shop_parcel_shop_bring`);
-- Create "parcel_shop_da_os" table
CREATE TABLE `parcel_shop_da_os` (`id` text NOT NULL, `shop_id` text NOT NULL, `parcel_shop_parcel_shop_dao` text NOT NULL, PRIMARY KEY (`id`), CONSTRAINT `parcel_shop_da_os_parcel_shops_parcel_shop_dao` FOREIGN KEY (`parcel_shop_parcel_shop_dao`) REFERENCES `parcel_shops` (`id`) ON DELETE NO ACTION);
-- Create index "parcel_shop_da_os_parcel_shop_parcel_shop_dao_key" to table: "parcel_shop_da_os"
CREATE UNIQUE INDEX `parcel_shop_da_os_parcel_shop_parcel_shop_dao_key` ON `parcel_shop_da_os` (`parcel_shop_parcel_shop_dao`);
-- Create "parcel_shop_gl_ss" table
CREATE TABLE `parcel_shop_gl_ss` (`id` text NOT NULL, `gls_parcel_shop_id` text NOT NULL, `partner_id` text NULL, `type` text NULL, `parcel_shop_parcel_shop_gls` text NOT NULL, PRIMARY KEY (`id`), CONSTRAINT `parcel_shop_gl_ss_parcel_shops_parcel_shop_gls` FOREIGN KEY (`parcel_shop_parcel_shop_gls`) REFERENCES `parcel_shops` (`id`) ON DELETE NO ACTION);
-- Create index "parcel_shop_gl_ss_gls_parcel_shop_id_key" to table: "parcel_shop_gl_ss"
CREATE UNIQUE INDEX `parcel_shop_gl_ss_gls_parcel_shop_id_key` ON `parcel_shop_gl_ss` (`gls_parcel_shop_id`);
-- Create index "parcel_shop_gl_ss_parcel_shop_parcel_shop_gls_key" to table: "parcel_shop_gl_ss"
CREATE UNIQUE INDEX `parcel_shop_gl_ss_parcel_shop_parcel_shop_gls_key` ON `parcel_shop_gl_ss` (`parcel_shop_parcel_shop_gls`);
-- Create "parcel_shop_post_nords" table
CREATE TABLE `parcel_shop_post_nords` (`id` text NOT NULL, `service_point_id` text NOT NULL, `pudoid` text NOT NULL, `type_id` text NOT NULL, `parcel_shop_parcel_shop_post_nord` text NOT NULL, PRIMARY KEY (`id`), CONSTRAINT `parcel_shop_post_nords_parcel_shops_parcel_shop_post_nord` FOREIGN KEY (`parcel_shop_parcel_shop_post_nord`) REFERENCES `parcel_shops` (`id`) ON DELETE NO ACTION);
-- Create index "parcel_shop_post_nords_pudoid_key" to table: "parcel_shop_post_nords"
CREATE UNIQUE INDEX `parcel_shop_post_nords_pudoid_key` ON `parcel_shop_post_nords` (`pudoid`);
-- Create index "parcel_shop_post_nords_parcel_shop_parcel_shop_post_nord_key" to table: "parcel_shop_post_nords"
CREATE UNIQUE INDEX `parcel_shop_post_nords_parcel_shop_parcel_shop_post_nord_key` ON `parcel_shop_post_nords` (`parcel_shop_parcel_shop_post_nord`);
-- Create "plans" table
CREATE TABLE `plans` (`id` text NOT NULL, `label` text NOT NULL, `rank` integer NOT NULL, `price_dkk` integer NOT NULL, `created_at` datetime NOT NULL, PRIMARY KEY (`id`));
-- Create index "plans_label_key" to table: "plans"
CREATE UNIQUE INDEX `plans_label_key` ON `plans` (`label`);
-- Create "plan_histories" table
CREATE TABLE `plan_histories` (`id` text NOT NULL, `created_at` datetime NOT NULL, `change_history_plan_history` text NOT NULL, `plan_plan_history_plan` text NOT NULL, `tenant_id` text NOT NULL, `user_plan_history_user` text NOT NULL, PRIMARY KEY (`id`), CONSTRAINT `plan_histories_change_histories_plan_history` FOREIGN KEY (`change_history_plan_history`) REFERENCES `change_histories` (`id`) ON DELETE NO ACTION, CONSTRAINT `plan_histories_plans_plan_history_plan` FOREIGN KEY (`plan_plan_history_plan`) REFERENCES `plans` (`id`) ON DELETE NO ACTION, CONSTRAINT `plan_histories_tenants_tenant` FOREIGN KEY (`tenant_id`) REFERENCES `tenants` (`id`) ON DELETE NO ACTION, CONSTRAINT `plan_histories_users_plan_history_user` FOREIGN KEY (`user_plan_history_user`) REFERENCES `users` (`id`) ON DELETE NO ACTION);
-- Create index "planhistory_tenant_id" to table: "plan_histories"
CREATE INDEX `planhistory_tenant_id` ON `plan_histories` (`tenant_id`);
-- Create "print_jobs" table
CREATE TABLE `print_jobs` (`id` text NOT NULL, `status` text NOT NULL, `file_extension` text NOT NULL, `document_type` text NOT NULL, `base64_print_data` text NOT NULL, `created_at` datetime NOT NULL, `tenant_id` text NOT NULL, `print_job_printer` text NOT NULL, `print_job_colli` text NULL, `print_job_shipment_parcel` text NULL, PRIMARY KEY (`id`), CONSTRAINT `print_jobs_tenants_tenant` FOREIGN KEY (`tenant_id`) REFERENCES `tenants` (`id`) ON DELETE NO ACTION, CONSTRAINT `print_jobs_printers_printer` FOREIGN KEY (`print_job_printer`) REFERENCES `printers` (`id`) ON DELETE NO ACTION, CONSTRAINT `print_jobs_collis_colli` FOREIGN KEY (`print_job_colli`) REFERENCES `collis` (`id`) ON DELETE SET NULL, CONSTRAINT `print_jobs_shipment_parcels_shipment_parcel` FOREIGN KEY (`print_job_shipment_parcel`) REFERENCES `shipment_parcels` (`id`) ON DELETE SET NULL);
-- Create index "printjob_tenant_id" to table: "print_jobs"
CREATE INDEX `printjob_tenant_id` ON `print_jobs` (`tenant_id`);
-- Create "printers" table
CREATE TABLE `printers` (`id` text NOT NULL, `device_id` text NOT NULL, `name` text NOT NULL, `label_zpl` bool NOT NULL DEFAULT (false), `label_pdf` bool NOT NULL DEFAULT (false), `document` bool NOT NULL DEFAULT (false), `rotate_180` bool NOT NULL DEFAULT (false), `print_size` text NOT NULL DEFAULT ('A4'), `created_at` datetime NOT NULL, `last_ping` datetime NOT NULL, `tenant_id` text NOT NULL, `workstation_printer` text NOT NULL, PRIMARY KEY (`id`), CONSTRAINT `printers_tenants_tenant` FOREIGN KEY (`tenant_id`) REFERENCES `tenants` (`id`) ON DELETE NO ACTION, CONSTRAINT `printers_workstations_printer` FOREIGN KEY (`workstation_printer`) REFERENCES `workstations` (`id`) ON DELETE NO ACTION);
-- Create index "printers_device_id_key" to table: "printers"
CREATE UNIQUE INDEX `printers_device_id_key` ON `printers` (`device_id`);
-- Create index "printer_tenant_id" to table: "printers"
CREATE INDEX `printer_tenant_id` ON `printers` (`tenant_id`);
-- Create "products" table
CREATE TABLE `products` (`id` text NOT NULL, `external_id` text NULL, `title` text NOT NULL, `body_html` text NULL, `status` text NOT NULL DEFAULT ('active'), `created_at` datetime NULL, `updated_at` datetime NOT NULL, `tenant_id` text NOT NULL, PRIMARY KEY (`id`), CONSTRAINT `products_tenants_tenant` FOREIGN KEY (`tenant_id`) REFERENCES `tenants` (`id`) ON DELETE NO ACTION);
-- Create index "product_tenant_id" to table: "products"
CREATE INDEX `product_tenant_id` ON `products` (`tenant_id`);
-- Create index "product_external_id_tenant_id" to table: "products"
CREATE UNIQUE INDEX `product_external_id_tenant_id` ON `products` (`external_id`, `tenant_id`);
-- Create "product_images" table
CREATE TABLE `product_images` (`id` text NOT NULL, `external_id` text NULL, `url` text NOT NULL, `tenant_id` text NOT NULL, `product_image_product` text NOT NULL, PRIMARY KEY (`id`), CONSTRAINT `product_images_tenants_tenant` FOREIGN KEY (`tenant_id`) REFERENCES `tenants` (`id`) ON DELETE NO ACTION, CONSTRAINT `product_images_products_product` FOREIGN KEY (`product_image_product`) REFERENCES `products` (`id`) ON DELETE NO ACTION);
-- Create index "productimage_tenant_id" to table: "product_images"
CREATE INDEX `productimage_tenant_id` ON `product_images` (`tenant_id`);
-- Create index "productimage_external_id_tenant_id" to table: "product_images"
CREATE UNIQUE INDEX `productimage_external_id_tenant_id` ON `product_images` (`external_id`, `tenant_id`);
-- Create "product_tags" table
CREATE TABLE `product_tags` (`id` text NOT NULL, `name` text NOT NULL, `created_at` datetime NULL, `tenant_id` text NOT NULL, PRIMARY KEY (`id`), CONSTRAINT `product_tags_tenants_tenant` FOREIGN KEY (`tenant_id`) REFERENCES `tenants` (`id`) ON DELETE NO ACTION);
-- Create index "producttag_tenant_id" to table: "product_tags"
CREATE INDEX `producttag_tenant_id` ON `product_tags` (`tenant_id`);
-- Create index "producttag_name_tenant_id" to table: "product_tags"
CREATE UNIQUE INDEX `producttag_name_tenant_id` ON `product_tags` (`name`, `tenant_id`);
-- Create "product_variants" table
CREATE TABLE `product_variants` (`id` text NOT NULL, `archived` bool NOT NULL DEFAULT (false), `external_id` text NULL, `description` text NULL, `ean_number` text NULL, `weight_g` integer NULL DEFAULT (0), `dimension_length` integer NULL, `dimension_width` integer NULL, `dimension_height` integer NULL, `created_at` datetime NULL, `updated_at` datetime NOT NULL, `product_product_variant` text NOT NULL, `tenant_id` text NOT NULL, PRIMARY KEY (`id`), CONSTRAINT `product_variants_products_product_variant` FOREIGN KEY (`product_product_variant`) REFERENCES `products` (`id`) ON DELETE NO ACTION, CONSTRAINT `product_variants_tenants_tenant` FOREIGN KEY (`tenant_id`) REFERENCES `tenants` (`id`) ON DELETE NO ACTION);
-- Create index "productvariant_tenant_id" to table: "product_variants"
CREATE INDEX `productvariant_tenant_id` ON `product_variants` (`tenant_id`);
-- Create index "productvariant_external_id_tenant_id" to table: "product_variants"
CREATE UNIQUE INDEX `productvariant_external_id_tenant_id` ON `product_variants` (`external_id`, `tenant_id`);
-- Create "return_collis" table
CREATE TABLE `return_collis` (`id` text NOT NULL, `expected_at` datetime NULL, `label_pdf` text NULL, `label_png` text NULL, `qr_code_png` text NULL, `comment` text NULL, `created_at` datetime NOT NULL, `status` text NOT NULL DEFAULT ('Opened'), `email_received` datetime NULL, `email_accepted` datetime NULL, `email_confirmation_label` datetime NULL, `email_confirmation_qr_code` datetime NULL, `order_return_colli` text NOT NULL, `tenant_id` text NOT NULL, `return_colli_recipient` text NOT NULL, `return_colli_sender` text NOT NULL, `return_colli_delivery_option` text NULL, `return_colli_return_portal` text NOT NULL, `return_colli_packaging` text NULL, PRIMARY KEY (`id`), CONSTRAINT `return_collis_orders_return_colli` FOREIGN KEY (`order_return_colli`) REFERENCES `orders` (`id`) ON DELETE NO ACTION, CONSTRAINT `return_collis_tenants_tenant` FOREIGN KEY (`tenant_id`) REFERENCES `tenants` (`id`) ON DELETE NO ACTION, CONSTRAINT `return_collis_addresses_recipient` FOREIGN KEY (`return_colli_recipient`) REFERENCES `addresses` (`id`) ON DELETE NO ACTION, CONSTRAINT `return_collis_addresses_sender` FOREIGN KEY (`return_colli_sender`) REFERENCES `addresses` (`id`) ON DELETE NO ACTION, CONSTRAINT `return_collis_delivery_options_delivery_option` FOREIGN KEY (`return_colli_delivery_option`) REFERENCES `delivery_options` (`id`) ON DELETE SET NULL, CONSTRAINT `return_collis_return_portals_return_portal` FOREIGN KEY (`return_colli_return_portal`) REFERENCES `return_portals` (`id`) ON DELETE NO ACTION, CONSTRAINT `return_collis_packagings_packaging` FOREIGN KEY (`return_colli_packaging`) REFERENCES `packagings` (`id`) ON DELETE SET NULL);
-- Create index "returncolli_tenant_id" to table: "return_collis"
CREATE INDEX `returncolli_tenant_id` ON `return_collis` (`tenant_id`);
-- Create "return_colli_histories" table
CREATE TABLE `return_colli_histories` (`id` text NOT NULL, `description` text NOT NULL, `type` text NOT NULL, `change_history_return_colli_history` text NOT NULL, `return_colli_return_colli_history` text NOT NULL, `tenant_id` text NOT NULL, PRIMARY KEY (`id`), CONSTRAINT `return_colli_histories_change_histories_return_colli_history` FOREIGN KEY (`change_history_return_colli_history`) REFERENCES `change_histories` (`id`) ON DELETE NO ACTION, CONSTRAINT `return_colli_histories_return_collis_return_colli_history` FOREIGN KEY (`return_colli_return_colli_history`) REFERENCES `return_collis` (`id`) ON DELETE NO ACTION, CONSTRAINT `return_colli_histories_tenants_tenant` FOREIGN KEY (`tenant_id`) REFERENCES `tenants` (`id`) ON DELETE NO ACTION);
-- Create index "returncollihistory_tenant_id" to table: "return_colli_histories"
CREATE INDEX `returncollihistory_tenant_id` ON `return_colli_histories` (`tenant_id`);
-- Create "return_order_lines" table
CREATE TABLE `return_order_lines` (`id` text NOT NULL, `units` integer NOT NULL, `return_colli_return_order_line` text NOT NULL, `tenant_id` text NOT NULL, `return_order_line_order_line` text NOT NULL, `return_order_line_return_portal_claim` text NOT NULL, PRIMARY KEY (`id`), CONSTRAINT `return_order_lines_return_collis_return_order_line` FOREIGN KEY (`return_colli_return_order_line`) REFERENCES `return_collis` (`id`) ON DELETE NO ACTION, CONSTRAINT `return_order_lines_tenants_tenant` FOREIGN KEY (`tenant_id`) REFERENCES `tenants` (`id`) ON DELETE NO ACTION, CONSTRAINT `return_order_lines_order_lines_order_line` FOREIGN KEY (`return_order_line_order_line`) REFERENCES `order_lines` (`id`) ON DELETE NO ACTION, CONSTRAINT `return_order_lines_return_portal_claims_return_portal_claim` FOREIGN KEY (`return_order_line_return_portal_claim`) REFERENCES `return_portal_claims` (`id`) ON DELETE NO ACTION);
-- Create index "returnorderline_tenant_id" to table: "return_order_lines"
CREATE INDEX `returnorderline_tenant_id` ON `return_order_lines` (`tenant_id`);
-- Create "return_portals" table
CREATE TABLE `return_portals` (`id` text NOT NULL, `name` text NOT NULL, `return_open_hours` integer NOT NULL DEFAULT (720), `automatically_accept` bool NOT NULL DEFAULT (false), `tenant_id` text NOT NULL, `return_portal_email_confirmation_label` text NULL, `return_portal_email_confirmation_qr_code` text NULL, `return_portal_email_received` text NULL, `return_portal_email_accepted` text NULL, PRIMARY KEY (`id`), CONSTRAINT `return_portals_tenants_tenant` FOREIGN KEY (`tenant_id`) REFERENCES `tenants` (`id`) ON DELETE NO ACTION, CONSTRAINT `return_portals_email_templates_email_confirmation_label` FOREIGN KEY (`return_portal_email_confirmation_label`) REFERENCES `email_templates` (`id`) ON DELETE SET NULL, CONSTRAINT `return_portals_email_templates_email_confirmation_qr_code` FOREIGN KEY (`return_portal_email_confirmation_qr_code`) REFERENCES `email_templates` (`id`) ON DELETE SET NULL, CONSTRAINT `return_portals_email_templates_email_received` FOREIGN KEY (`return_portal_email_received`) REFERENCES `email_templates` (`id`) ON DELETE SET NULL, CONSTRAINT `return_portals_email_templates_email_accepted` FOREIGN KEY (`return_portal_email_accepted`) REFERENCES `email_templates` (`id`) ON DELETE SET NULL);
-- Create index "returnportal_tenant_id" to table: "return_portals"
CREATE INDEX `returnportal_tenant_id` ON `return_portals` (`tenant_id`);
-- Create "return_portal_claims" table
CREATE TABLE `return_portal_claims` (`id` text NOT NULL, `name` text NOT NULL, `description` text NOT NULL, `restockable` bool NOT NULL, `archived` bool NOT NULL, `return_portal_return_portal_claim` text NOT NULL, `tenant_id` text NOT NULL, `return_portal_claim_return_location` text NULL, PRIMARY KEY (`id`), CONSTRAINT `return_portal_claims_return_portals_return_portal_claim` FOREIGN KEY (`return_portal_return_portal_claim`) REFERENCES `return_portals` (`id`) ON DELETE NO ACTION, CONSTRAINT `return_portal_claims_tenants_tenant` FOREIGN KEY (`tenant_id`) REFERENCES `tenants` (`id`) ON DELETE NO ACTION, CONSTRAINT `return_portal_claims_locations_return_location` FOREIGN KEY (`return_portal_claim_return_location`) REFERENCES `locations` (`id`) ON DELETE SET NULL);
-- Create index "returnportalclaim_tenant_id" to table: "return_portal_claims"
CREATE INDEX `returnportalclaim_tenant_id` ON `return_portal_claims` (`tenant_id`);
-- Create "seat_groups" table
CREATE TABLE `seat_groups` (`id` text NOT NULL, `name` text NOT NULL, `created_at` datetime NULL, `tenant_id` text NOT NULL, PRIMARY KEY (`id`), CONSTRAINT `seat_groups_tenants_tenant` FOREIGN KEY (`tenant_id`) REFERENCES `tenants` (`id`) ON DELETE NO ACTION);
-- Create index "seatgroup_tenant_id" to table: "seat_groups"
CREATE INDEX `seatgroup_tenant_id` ON `seat_groups` (`tenant_id`);
-- Create "seat_group_access_rights" table
CREATE TABLE `seat_group_access_rights` (`id` text NOT NULL, `level` text NOT NULL DEFAULT ('none'), `tenant_id` text NOT NULL, `access_right_id` text NOT NULL, `seat_group_id` text NOT NULL, PRIMARY KEY (`id`), CONSTRAINT `seat_group_access_rights_tenants_tenant` FOREIGN KEY (`tenant_id`) REFERENCES `tenants` (`id`) ON DELETE NO ACTION, CONSTRAINT `seat_group_access_rights_access_rights_access_right` FOREIGN KEY (`access_right_id`) REFERENCES `access_rights` (`id`) ON DELETE NO ACTION, CONSTRAINT `seat_group_access_rights_seat_groups_seat_group` FOREIGN KEY (`seat_group_id`) REFERENCES `seat_groups` (`id`) ON DELETE NO ACTION);
-- Create index "seatgroupaccessright_tenant_id" to table: "seat_group_access_rights"
CREATE INDEX `seatgroupaccessright_tenant_id` ON `seat_group_access_rights` (`tenant_id`);
-- Create index "seatgroupaccessright_seat_group_id_access_right_id" to table: "seat_group_access_rights"
CREATE UNIQUE INDEX `seatgroupaccessright_seat_group_id_access_right_id` ON `seat_group_access_rights` (`seat_group_id`, `access_right_id`);
-- Create "shipments" table
CREATE TABLE `shipments` (`id` text NOT NULL, `shipment_public_id` text NOT NULL, `created_at` datetime NOT NULL, `status` text NOT NULL, `tenant_id` text NOT NULL, `shipment_carrier` text NOT NULL, PRIMARY KEY (`id`), CONSTRAINT `shipments_tenants_tenant` FOREIGN KEY (`tenant_id`) REFERENCES `tenants` (`id`) ON DELETE NO ACTION, CONSTRAINT `shipments_carriers_carrier` FOREIGN KEY (`shipment_carrier`) REFERENCES `carriers` (`id`) ON DELETE NO ACTION);
-- Create index "shipment_tenant_id" to table: "shipments"
CREATE INDEX `shipment_tenant_id` ON `shipments` (`tenant_id`);
-- Create "shipment_brings" table
CREATE TABLE `shipment_brings` (`id` text NOT NULL, `consignment_number` text NOT NULL, `shipment_shipment_bring` text NOT NULL, `tenant_id` text NOT NULL, PRIMARY KEY (`id`), CONSTRAINT `shipment_brings_shipments_shipment_bring` FOREIGN KEY (`shipment_shipment_bring`) REFERENCES `shipments` (`id`) ON DELETE NO ACTION, CONSTRAINT `shipment_brings_tenants_tenant` FOREIGN KEY (`tenant_id`) REFERENCES `tenants` (`id`) ON DELETE NO ACTION);
-- Create index "shipment_brings_shipment_shipment_bring_key" to table: "shipment_brings"
CREATE UNIQUE INDEX `shipment_brings_shipment_shipment_bring_key` ON `shipment_brings` (`shipment_shipment_bring`);
-- Create index "shipmentbring_tenant_id" to table: "shipment_brings"
CREATE INDEX `shipmentbring_tenant_id` ON `shipment_brings` (`tenant_id`);
-- Create "shipment_da_os" table
CREATE TABLE `shipment_da_os` (`id` text NOT NULL, `barcode_id` text NOT NULL, `shipment_shipment_dao` text NOT NULL, `tenant_id` text NOT NULL, PRIMARY KEY (`id`), CONSTRAINT `shipment_da_os_shipments_shipment_dao` FOREIGN KEY (`shipment_shipment_dao`) REFERENCES `shipments` (`id`) ON DELETE NO ACTION, CONSTRAINT `shipment_da_os_tenants_tenant` FOREIGN KEY (`tenant_id`) REFERENCES `tenants` (`id`) ON DELETE NO ACTION);
-- Create index "shipment_da_os_shipment_shipment_dao_key" to table: "shipment_da_os"
CREATE UNIQUE INDEX `shipment_da_os_shipment_shipment_dao_key` ON `shipment_da_os` (`shipment_shipment_dao`);
-- Create index "shipmentdao_tenant_id" to table: "shipment_da_os"
CREATE INDEX `shipmentdao_tenant_id` ON `shipment_da_os` (`tenant_id`);
-- Create "shipment_dfs" table
CREATE TABLE `shipment_dfs` (`id` text NOT NULL, `shipment_shipment_df` text NOT NULL, `tenant_id` text NOT NULL, PRIMARY KEY (`id`), CONSTRAINT `shipment_dfs_shipments_shipment_df` FOREIGN KEY (`shipment_shipment_df`) REFERENCES `shipments` (`id`) ON DELETE NO ACTION, CONSTRAINT `shipment_dfs_tenants_tenant` FOREIGN KEY (`tenant_id`) REFERENCES `tenants` (`id`) ON DELETE NO ACTION);
-- Create index "shipment_dfs_shipment_shipment_df_key" to table: "shipment_dfs"
CREATE UNIQUE INDEX `shipment_dfs_shipment_shipment_df_key` ON `shipment_dfs` (`shipment_shipment_df`);
-- Create index "shipmentdf_tenant_id" to table: "shipment_dfs"
CREATE INDEX `shipmentdf_tenant_id` ON `shipment_dfs` (`tenant_id`);
-- Create "shipment_ds_vs" table
CREATE TABLE `shipment_ds_vs` (`id` text NOT NULL, `barcode_id` text NOT NULL, `shipment_shipment_dsv` text NOT NULL, `tenant_id` text NOT NULL, PRIMARY KEY (`id`), CONSTRAINT `shipment_ds_vs_shipments_shipment_dsv` FOREIGN KEY (`shipment_shipment_dsv`) REFERENCES `shipments` (`id`) ON DELETE NO ACTION, CONSTRAINT `shipment_ds_vs_tenants_tenant` FOREIGN KEY (`tenant_id`) REFERENCES `tenants` (`id`) ON DELETE NO ACTION);
-- Create index "shipment_ds_vs_shipment_shipment_dsv_key" to table: "shipment_ds_vs"
CREATE UNIQUE INDEX `shipment_ds_vs_shipment_shipment_dsv_key` ON `shipment_ds_vs` (`shipment_shipment_dsv`);
-- Create index "shipmentdsv_tenant_id" to table: "shipment_ds_vs"
CREATE INDEX `shipmentdsv_tenant_id` ON `shipment_ds_vs` (`tenant_id`);
-- Create "shipment_easy_posts" table
CREATE TABLE `shipment_easy_posts` (`id` text NOT NULL, `tracking_number` text NULL, `ep_shipment_id` text NULL, `rate` real NULL, `est_delivery_date` datetime NULL, `shipment_shipment_easy_post` text NOT NULL, `tenant_id` text NOT NULL, PRIMARY KEY (`id`), CONSTRAINT `shipment_easy_posts_shipments_shipment_easy_post` FOREIGN KEY (`shipment_shipment_easy_post`) REFERENCES `shipments` (`id`) ON DELETE NO ACTION, CONSTRAINT `shipment_easy_posts_tenants_tenant` FOREIGN KEY (`tenant_id`) REFERENCES `tenants` (`id`) ON DELETE NO ACTION);
-- Create index "shipment_easy_posts_shipment_shipment_easy_post_key" to table: "shipment_easy_posts"
CREATE UNIQUE INDEX `shipment_easy_posts_shipment_shipment_easy_post_key` ON `shipment_easy_posts` (`shipment_shipment_easy_post`);
-- Create index "shipmenteasypost_tenant_id" to table: "shipment_easy_posts"
CREATE INDEX `shipmenteasypost_tenant_id` ON `shipment_easy_posts` (`tenant_id`);
-- Create "shipment_gl_ss" table
CREATE TABLE `shipment_gl_ss` (`id` text NOT NULL, `consignment_id` text NOT NULL, `shipment_shipment_gls` text NOT NULL, `tenant_id` text NOT NULL, PRIMARY KEY (`id`), CONSTRAINT `shipment_gl_ss_shipments_shipment_gls` FOREIGN KEY (`shipment_shipment_gls`) REFERENCES `shipments` (`id`) ON DELETE NO ACTION, CONSTRAINT `shipment_gl_ss_tenants_tenant` FOREIGN KEY (`tenant_id`) REFERENCES `tenants` (`id`) ON DELETE NO ACTION);
-- Create index "shipment_gl_ss_shipment_shipment_gls_key" to table: "shipment_gl_ss"
CREATE UNIQUE INDEX `shipment_gl_ss_shipment_shipment_gls_key` ON `shipment_gl_ss` (`shipment_shipment_gls`);
-- Create index "shipmentgls_tenant_id" to table: "shipment_gl_ss"
CREATE INDEX `shipmentgls_tenant_id` ON `shipment_gl_ss` (`tenant_id`);
-- Create "shipment_histories" table
CREATE TABLE `shipment_histories` (`id` text NOT NULL, `type` text NOT NULL, `change_history_shipment_history` text NOT NULL, `shipment_shipment_history` text NOT NULL, `tenant_id` text NOT NULL, PRIMARY KEY (`id`), CONSTRAINT `shipment_histories_change_histories_shipment_history` FOREIGN KEY (`change_history_shipment_history`) REFERENCES `change_histories` (`id`) ON DELETE NO ACTION, CONSTRAINT `shipment_histories_shipments_shipment_history` FOREIGN KEY (`shipment_shipment_history`) REFERENCES `shipments` (`id`) ON DELETE NO ACTION, CONSTRAINT `shipment_histories_tenants_tenant` FOREIGN KEY (`tenant_id`) REFERENCES `tenants` (`id`) ON DELETE NO ACTION);
-- Create index "shipmenthistory_tenant_id" to table: "shipment_histories"
CREATE INDEX `shipmenthistory_tenant_id` ON `shipment_histories` (`tenant_id`);
-- Create "shipment_pallets" table
CREATE TABLE `shipment_pallets` (`id` text NOT NULL, `barcode` text NOT NULL, `colli_number` text NOT NULL, `carrier_id` text NOT NULL, `label_pdf` text NULL, `label_zpl` text NULL, `status` text NOT NULL DEFAULT ('pending'), `pallet_shipment_pallet` text NULL, `shipment_shipment_pallet` text NOT NULL, `tenant_id` text NOT NULL, PRIMARY KEY (`id`), CONSTRAINT `shipment_pallets_pallets_shipment_pallet` FOREIGN KEY (`pallet_shipment_pallet`) REFERENCES `pallets` (`id`) ON DELETE SET NULL, CONSTRAINT `shipment_pallets_shipments_shipment_pallet` FOREIGN KEY (`shipment_shipment_pallet`) REFERENCES `shipments` (`id`) ON DELETE NO ACTION, CONSTRAINT `shipment_pallets_tenants_tenant` FOREIGN KEY (`tenant_id`) REFERENCES `tenants` (`id`) ON DELETE NO ACTION);
-- Create index "shipment_pallets_pallet_shipment_pallet_key" to table: "shipment_pallets"
CREATE UNIQUE INDEX `shipment_pallets_pallet_shipment_pallet_key` ON `shipment_pallets` (`pallet_shipment_pallet`);
-- Create index "shipmentpallet_tenant_id" to table: "shipment_pallets"
CREATE INDEX `shipmentpallet_tenant_id` ON `shipment_pallets` (`tenant_id`);
-- Create "shipment_parcels" table
CREATE TABLE `shipment_parcels` (`id` text NOT NULL, `label_pdf` text NULL, `label_zpl` text NULL, `item_id` text NULL, `status` text NOT NULL DEFAULT ('pending'), `cc_pickup_signature_urls` json NULL, `expected_at` datetime NULL, `fulfillment_synced_at` datetime NULL, `cancel_synced_at` datetime NULL, `colli_shipment_parcel` text NULL, `shipment_shipment_parcel` text NOT NULL, `tenant_id` text NOT NULL, `shipment_parcel_packaging` text NULL, PRIMARY KEY (`id`), CONSTRAINT `shipment_parcels_collis_shipment_parcel` FOREIGN KEY (`colli_shipment_parcel`) REFERENCES `collis` (`id`) ON DELETE SET NULL, CONSTRAINT `shipment_parcels_shipments_shipment_parcel` FOREIGN KEY (`shipment_shipment_parcel`) REFERENCES `shipments` (`id`) ON DELETE NO ACTION, CONSTRAINT `shipment_parcels_tenants_tenant` FOREIGN KEY (`tenant_id`) REFERENCES `tenants` (`id`) ON DELETE NO ACTION, CONSTRAINT `shipment_parcels_packagings_packaging` FOREIGN KEY (`shipment_parcel_packaging`) REFERENCES `packagings` (`id`) ON DELETE SET NULL);
-- Create index "shipment_parcels_colli_shipment_parcel_key" to table: "shipment_parcels"
CREATE UNIQUE INDEX `shipment_parcels_colli_shipment_parcel_key` ON `shipment_parcels` (`colli_shipment_parcel`);
-- Create index "shipmentparcel_tenant_id" to table: "shipment_parcels"
CREATE INDEX `shipmentparcel_tenant_id` ON `shipment_parcels` (`tenant_id`);
-- Create "shipment_post_nords" table
CREATE TABLE `shipment_post_nords` (`id` text NOT NULL, `booking_id` text NOT NULL, `item_id` text NOT NULL, `shipment_reference_no` text NOT NULL, `shipment_shipment_post_nord` text NOT NULL, `tenant_id` text NOT NULL, PRIMARY KEY (`id`), CONSTRAINT `shipment_post_nords_shipments_shipment_post_nord` FOREIGN KEY (`shipment_shipment_post_nord`) REFERENCES `shipments` (`id`) ON DELETE NO ACTION, CONSTRAINT `shipment_post_nords_tenants_tenant` FOREIGN KEY (`tenant_id`) REFERENCES `tenants` (`id`) ON DELETE NO ACTION);
-- Create index "shipment_post_nords_shipment_shipment_post_nord_key" to table: "shipment_post_nords"
CREATE UNIQUE INDEX `shipment_post_nords_shipment_shipment_post_nord_key` ON `shipment_post_nords` (`shipment_shipment_post_nord`);
-- Create index "shipmentpostnord_tenant_id" to table: "shipment_post_nords"
CREATE INDEX `shipmentpostnord_tenant_id` ON `shipment_post_nords` (`tenant_id`);
-- Create "shipment_usp_ss" table
CREATE TABLE `shipment_usp_ss` (`id` text NOT NULL, `tracking_number` text NULL, `postage` real NULL, `scheduled_delivery_date` datetime NULL, `shipment_shipment_usps` text NOT NULL, `tenant_id` text NOT NULL, PRIMARY KEY (`id`), CONSTRAINT `shipment_usp_ss_shipments_shipment_usps` FOREIGN KEY (`shipment_shipment_usps`) REFERENCES `shipments` (`id`) ON DELETE NO ACTION, CONSTRAINT `shipment_usp_ss_tenants_tenant` FOREIGN KEY (`tenant_id`) REFERENCES `tenants` (`id`) ON DELETE NO ACTION);
-- Create index "shipment_usp_ss_shipment_shipment_usps_key" to table: "shipment_usp_ss"
CREATE UNIQUE INDEX `shipment_usp_ss_shipment_shipment_usps_key` ON `shipment_usp_ss` (`shipment_shipment_usps`);
-- Create index "shipmentusps_tenant_id" to table: "shipment_usp_ss"
CREATE INDEX `shipmentusps_tenant_id` ON `shipment_usp_ss` (`tenant_id`);
-- Create "signup_options" table
CREATE TABLE `signup_options` (`id` text NOT NULL, `better_delivery_options` bool NOT NULL, `improve_pick_pack` bool NOT NULL, `shipping_label` bool NOT NULL, `custom_docs` bool NOT NULL, `reduced_costs` bool NOT NULL, `easy_returns` bool NOT NULL, `click_collect` bool NOT NULL, `num_shipments` integer NOT NULL, `user_signup_options` text NOT NULL, PRIMARY KEY (`id`), CONSTRAINT `signup_options_users_signup_options` FOREIGN KEY (`user_signup_options`) REFERENCES `users` (`id`) ON DELETE NO ACTION);
-- Create index "signup_options_user_signup_options_key" to table: "signup_options"
CREATE UNIQUE INDEX `signup_options_user_signup_options_key` ON `signup_options` (`user_signup_options`);
-- Create "system_events" table
CREATE TABLE `system_events` (`id` text NOT NULL, `event_type` text NOT NULL, `event_type_id` text NULL, `status` text NOT NULL, `description` text NOT NULL, `data` text NULL, `updated_at` datetime NOT NULL, `created_at` datetime NULL, `tenant_id` text NOT NULL, PRIMARY KEY (`id`), CONSTRAINT `system_events_tenants_tenant` FOREIGN KEY (`tenant_id`) REFERENCES `tenants` (`id`) ON DELETE NO ACTION);
-- Create index "systemevents_tenant_id" to table: "system_events"
CREATE INDEX `systemevents_tenant_id` ON `system_events` (`tenant_id`);
-- Create "tenants" table
CREATE TABLE `tenants` (`id` text NOT NULL, `name` text NOT NULL, `vat_number` text NULL, `invoice_reference` text NULL, `plan_tenant` text NOT NULL, `tenant_company_address` text NULL, `tenant_default_language` text NOT NULL, `tenant_billing_contact` text NULL, `tenant_admin_contact` text NULL, PRIMARY KEY (`id`), CONSTRAINT `tenants_plans_tenant` FOREIGN KEY (`plan_tenant`) REFERENCES `plans` (`id`) ON DELETE NO ACTION, CONSTRAINT `tenants_addresses_company_address` FOREIGN KEY (`tenant_company_address`) REFERENCES `addresses` (`id`) ON DELETE SET NULL, CONSTRAINT `tenants_languages_default_language` FOREIGN KEY (`tenant_default_language`) REFERENCES `languages` (`id`) ON DELETE NO ACTION, CONSTRAINT `tenants_contacts_billing_contact` FOREIGN KEY (`tenant_billing_contact`) REFERENCES `contacts` (`id`) ON DELETE SET NULL, CONSTRAINT `tenants_contacts_admin_contact` FOREIGN KEY (`tenant_admin_contact`) REFERENCES `contacts` (`id`) ON DELETE SET NULL);
-- Create index "tenants_name_key" to table: "tenants"
CREATE UNIQUE INDEX `tenants_name_key` ON `tenants` (`name`);
-- Create "users" table
CREATE TABLE `users` (`id` text NOT NULL, `name` text NULL, `surname` text NULL, `phone_number` text NULL, `email` text NOT NULL, `password` text NULL, `hash` text NOT NULL, `is_account_owner` bool NOT NULL DEFAULT (false), `is_global_admin` bool NOT NULL DEFAULT (false), `marketing_consent` bool NULL DEFAULT (true), `created_at` datetime NULL, `archived_at` datetime NULL, `pickup_day` text NOT NULL DEFAULT ('Today'), `pickup_day_last_changed` datetime NULL, `seat_group_user` text NULL, `tenant_id` text NOT NULL, `user_language` text NULL, PRIMARY KEY (`id`), CONSTRAINT `users_seat_groups_user` FOREIGN KEY (`seat_group_user`) REFERENCES `seat_groups` (`id`) ON DELETE SET NULL, CONSTRAINT `users_tenants_tenant` FOREIGN KEY (`tenant_id`) REFERENCES `tenants` (`id`) ON DELETE NO ACTION, CONSTRAINT `users_languages_language` FOREIGN KEY (`user_language`) REFERENCES `languages` (`id`) ON DELETE SET NULL);
-- Create index "users_email_key" to table: "users"
CREATE UNIQUE INDEX `users_email_key` ON `users` (`email`);
-- Create index "user_tenant_id" to table: "users"
CREATE INDEX `user_tenant_id` ON `users` (`tenant_id`);
-- Create "user_seats" table
CREATE TABLE `user_seats` (`id` text NOT NULL, `name` text NULL, `surname` text NULL, `email` text NOT NULL, `password` text NOT NULL, `hash` text NOT NULL, `created_at` datetime NOT NULL, `tenant_id` text NOT NULL, PRIMARY KEY (`id`), CONSTRAINT `user_seats_tenants_tenant` FOREIGN KEY (`tenant_id`) REFERENCES `tenants` (`id`) ON DELETE NO ACTION);
-- Create index "user_seats_email_key" to table: "user_seats"
CREATE UNIQUE INDEX `user_seats_email_key` ON `user_seats` (`email`);
-- Create index "userseat_tenant_id" to table: "user_seats"
CREATE INDEX `userseat_tenant_id` ON `user_seats` (`tenant_id`);
-- Create "workspace_recent_scans" table
CREATE TABLE `workspace_recent_scans` (`id` text NOT NULL, `created_at` datetime NOT NULL, `tenant_id` text NOT NULL, `workspace_recent_scan_shipment_parcel` text NULL, `workspace_recent_scan_user` text NOT NULL, PRIMARY KEY (`id`), CONSTRAINT `workspace_recent_scans_tenants_tenant` FOREIGN KEY (`tenant_id`) REFERENCES `tenants` (`id`) ON DELETE NO ACTION, CONSTRAINT `workspace_recent_scans_shipment_parcels_shipment_parcel` FOREIGN KEY (`workspace_recent_scan_shipment_parcel`) REFERENCES `shipment_parcels` (`id`) ON DELETE SET NULL, CONSTRAINT `workspace_recent_scans_users_user` FOREIGN KEY (`workspace_recent_scan_user`) REFERENCES `users` (`id`) ON DELETE NO ACTION);
-- Create index "workspacerecentscan_tenant_id" to table: "workspace_recent_scans"
CREATE INDEX `workspacerecentscan_tenant_id` ON `workspace_recent_scans` (`tenant_id`);
-- Create "workstations" table
CREATE TABLE `workstations` (`id` text NOT NULL, `archived_at` datetime NULL, `name` text NOT NULL, `device_type` text NOT NULL DEFAULT ('label_station'), `registration_code` text NOT NULL, `workstation_id` text NOT NULL, `created_at` datetime NOT NULL, `last_ping` datetime NULL, `status` text NOT NULL DEFAULT ('pending'), `user_selected_workstation` text NULL, `tenant_id` text NOT NULL, `workstation_user` text NULL, PRIMARY KEY (`id`), CONSTRAINT `workstations_users_selected_workstation` FOREIGN KEY (`user_selected_workstation`) REFERENCES `users` (`id`) ON DELETE SET NULL, CONSTRAINT `workstations_tenants_tenant` FOREIGN KEY (`tenant_id`) REFERENCES `tenants` (`id`) ON DELETE NO ACTION, CONSTRAINT `workstations_users_user` FOREIGN KEY (`workstation_user`) REFERENCES `users` (`id`) ON DELETE SET NULL);
-- Create index "workstations_user_selected_workstation_key" to table: "workstations"
CREATE UNIQUE INDEX `workstations_user_selected_workstation_key` ON `workstations` (`user_selected_workstation`);
-- Create index "workstation_tenant_id" to table: "workstations"
CREATE INDEX `workstation_tenant_id` ON `workstations` (`tenant_id`);
-- Create "carrier_additional_service_gls_countries_consignee" table
CREATE TABLE `carrier_additional_service_gls_countries_consignee` (`carrier_additional_service_gls_id` text NOT NULL, `country_id` text NOT NULL, PRIMARY KEY (`carrier_additional_service_gls_id`, `country_id`), CONSTRAINT `carrier_additional_service_gls_countries_consignee_carrier_additional_service_gls_id` FOREIGN KEY (`carrier_additional_service_gls_id`) REFERENCES `carrier_additional_service_gl_ss` (`id`) ON DELETE CASCADE, CONSTRAINT `carrier_additional_service_gls_countries_consignee_country_id` FOREIGN KEY (`country_id`) REFERENCES `countries` (`id`) ON DELETE CASCADE);
-- Create "carrier_additional_service_gls_countries_consignor" table
CREATE TABLE `carrier_additional_service_gls_countries_consignor` (`carrier_additional_service_gls_id` text NOT NULL, `country_id` text NOT NULL, PRIMARY KEY (`carrier_additional_service_gls_id`, `country_id`), CONSTRAINT `carrier_additional_service_gls_countries_consignor_carrier_additional_service_gls_id` FOREIGN KEY (`carrier_additional_service_gls_id`) REFERENCES `carrier_additional_service_gl_ss` (`id`) ON DELETE CASCADE, CONSTRAINT `carrier_additional_service_gls_countries_consignor_country_id` FOREIGN KEY (`country_id`) REFERENCES `countries` (`id`) ON DELETE CASCADE);
-- Create "carrier_additional_service_post_nord_countries_consignee" table
CREATE TABLE `carrier_additional_service_post_nord_countries_consignee` (`carrier_additional_service_post_nord_id` text NOT NULL, `country_id` text NOT NULL, PRIMARY KEY (`carrier_additional_service_post_nord_id`, `country_id`), CONSTRAINT `carrier_additional_service_post_nord_countries_consignee_carrier_additional_service_post_nord_id` FOREIGN KEY (`carrier_additional_service_post_nord_id`) REFERENCES `carrier_additional_service_post_nords` (`id`) ON DELETE CASCADE, CONSTRAINT `carrier_additional_service_post_nord_countries_consignee_country_id` FOREIGN KEY (`country_id`) REFERENCES `countries` (`id`) ON DELETE CASCADE);
-- Create "carrier_additional_service_post_nord_countries_consignor" table
CREATE TABLE `carrier_additional_service_post_nord_countries_consignor` (`carrier_additional_service_post_nord_id` text NOT NULL, `country_id` text NOT NULL, PRIMARY KEY (`carrier_additional_service_post_nord_id`, `country_id`), CONSTRAINT `carrier_additional_service_post_nord_countries_consignor_carrier_additional_service_post_nord_id` FOREIGN KEY (`carrier_additional_service_post_nord_id`) REFERENCES `carrier_additional_service_post_nords` (`id`) ON DELETE CASCADE, CONSTRAINT `carrier_additional_service_post_nord_countries_consignor_country_id` FOREIGN KEY (`country_id`) REFERENCES `countries` (`id`) ON DELETE CASCADE);
-- Create "carrier_service_dao_carrier_additional_service_dao" table
CREATE TABLE `carrier_service_dao_carrier_additional_service_dao` (`carrier_service_dao_id` text NOT NULL, `carrier_additional_service_dao_id` text NOT NULL, PRIMARY KEY (`carrier_service_dao_id`, `carrier_additional_service_dao_id`), CONSTRAINT `carrier_service_dao_carrier_additional_service_dao_carrier_service_dao_id` FOREIGN KEY (`carrier_service_dao_id`) REFERENCES `carrier_service_da_os` (`id`) ON DELETE CASCADE, CONSTRAINT `carrier_service_dao_carrier_additional_service_dao_carrier_additional_service_dao_id` FOREIGN KEY (`carrier_additional_service_dao_id`) REFERENCES `carrier_additional_service_da_os` (`id`) ON DELETE CASCADE);
-- Create "carrier_service_df_carrier_additional_service_df" table
CREATE TABLE `carrier_service_df_carrier_additional_service_df` (`carrier_service_df_id` text NOT NULL, `carrier_additional_service_df_id` text NOT NULL, PRIMARY KEY (`carrier_service_df_id`, `carrier_additional_service_df_id`), CONSTRAINT `carrier_service_df_carrier_additional_service_df_carrier_service_df_id` FOREIGN KEY (`carrier_service_df_id`) REFERENCES `carrier_service_dfs` (`id`) ON DELETE CASCADE, CONSTRAINT `carrier_service_df_carrier_additional_service_df_carrier_additional_service_df_id` FOREIGN KEY (`carrier_additional_service_df_id`) REFERENCES `carrier_additional_service_dfs` (`id`) ON DELETE CASCADE);
-- Create "carrier_service_dsv_carrier_additional_service_dsv" table
CREATE TABLE `carrier_service_dsv_carrier_additional_service_dsv` (`carrier_service_dsv_id` text NOT NULL, `carrier_additional_service_dsv_id` text NOT NULL, PRIMARY KEY (`carrier_service_dsv_id`, `carrier_additional_service_dsv_id`), CONSTRAINT `carrier_service_dsv_carrier_additional_service_dsv_carrier_service_dsv_id` FOREIGN KEY (`carrier_service_dsv_id`) REFERENCES `carrier_service_ds_vs` (`id`) ON DELETE CASCADE, CONSTRAINT `carrier_service_dsv_carrier_additional_service_dsv_carrier_additional_service_dsv_id` FOREIGN KEY (`carrier_additional_service_dsv_id`) REFERENCES `carrier_additional_service_ds_vs` (`id`) ON DELETE CASCADE);
-- Create "carrier_service_easy_post_carrier_add_serv_easy_post" table
CREATE TABLE `carrier_service_easy_post_carrier_add_serv_easy_post` (`carrier_service_easy_post_id` text NOT NULL, `carrier_additional_service_easy_post_id` text NOT NULL, PRIMARY KEY (`carrier_service_easy_post_id`, `carrier_additional_service_easy_post_id`), CONSTRAINT `carrier_service_easy_post_carrier_add_serv_easy_post_carrier_service_easy_post_id` FOREIGN KEY (`carrier_service_easy_post_id`) REFERENCES `carrier_service_easy_posts` (`id`) ON DELETE CASCADE, CONSTRAINT `carrier_service_easy_post_carrier_add_serv_easy_post_carrier_additional_service_easy_post_id` FOREIGN KEY (`carrier_additional_service_easy_post_id`) REFERENCES `carrier_additional_service_easy_posts` (`id`) ON DELETE CASCADE);
-- Create "colli_cancelled_shipment_parcel" table
CREATE TABLE `colli_cancelled_shipment_parcel` (`colli_id` text NOT NULL, `shipment_parcel_id` text NOT NULL, PRIMARY KEY (`colli_id`, `shipment_parcel_id`), CONSTRAINT `colli_cancelled_shipment_parcel_colli_id` FOREIGN KEY (`colli_id`) REFERENCES `collis` (`id`) ON DELETE CASCADE, CONSTRAINT `colli_cancelled_shipment_parcel_shipment_parcel_id` FOREIGN KEY (`shipment_parcel_id`) REFERENCES `shipment_parcels` (`id`) ON DELETE CASCADE);
-- Create "country_delivery_rule" table
CREATE TABLE `country_delivery_rule` (`country_id` text NOT NULL, `delivery_rule_id` text NOT NULL, PRIMARY KEY (`country_id`, `delivery_rule_id`), CONSTRAINT `country_delivery_rule_country_id` FOREIGN KEY (`country_id`) REFERENCES `countries` (`id`) ON DELETE CASCADE, CONSTRAINT `country_delivery_rule_delivery_rule_id` FOREIGN KEY (`delivery_rule_id`) REFERENCES `delivery_rules` (`id`) ON DELETE CASCADE);
-- Create "delivery_option_click_collect_location" table
CREATE TABLE `delivery_option_click_collect_location` (`delivery_option_id` text NOT NULL, `location_id` text NOT NULL, PRIMARY KEY (`delivery_option_id`, `location_id`), CONSTRAINT `delivery_option_click_collect_location_delivery_option_id` FOREIGN KEY (`delivery_option_id`) REFERENCES `delivery_options` (`id`) ON DELETE CASCADE, CONSTRAINT `delivery_option_click_collect_location_location_id` FOREIGN KEY (`location_id`) REFERENCES `locations` (`id`) ON DELETE CASCADE);
-- Create "delivery_option_bring_carrier_additional_service_bring" table
CREATE TABLE `delivery_option_bring_carrier_additional_service_bring` (`delivery_option_bring_id` text NOT NULL, `carrier_additional_service_bring_id` text NOT NULL, PRIMARY KEY (`delivery_option_bring_id`, `carrier_additional_service_bring_id`), CONSTRAINT `delivery_option_bring_carrier_additional_service_bring_delivery_option_bring_id` FOREIGN KEY (`delivery_option_bring_id`) REFERENCES `delivery_option_brings` (`id`) ON DELETE CASCADE, CONSTRAINT `delivery_option_bring_carrier_additional_service_bring_carrier_additional_service_bring_id` FOREIGN KEY (`carrier_additional_service_bring_id`) REFERENCES `carrier_additional_service_brings` (`id`) ON DELETE CASCADE);
-- Create "delivery_option_dao_carrier_additional_service_dao" table
CREATE TABLE `delivery_option_dao_carrier_additional_service_dao` (`delivery_option_dao_id` text NOT NULL, `carrier_additional_service_dao_id` text NOT NULL, PRIMARY KEY (`delivery_option_dao_id`, `carrier_additional_service_dao_id`), CONSTRAINT `delivery_option_dao_carrier_additional_service_dao_delivery_option_dao_id` FOREIGN KEY (`delivery_option_dao_id`) REFERENCES `delivery_option_da_os` (`id`) ON DELETE CASCADE, CONSTRAINT `delivery_option_dao_carrier_additional_service_dao_carrier_additional_service_dao_id` FOREIGN KEY (`carrier_additional_service_dao_id`) REFERENCES `carrier_additional_service_da_os` (`id`) ON DELETE CASCADE);
-- Create "delivery_option_df_carrier_additional_service_df" table
CREATE TABLE `delivery_option_df_carrier_additional_service_df` (`delivery_option_df_id` text NOT NULL, `carrier_additional_service_df_id` text NOT NULL, PRIMARY KEY (`delivery_option_df_id`, `carrier_additional_service_df_id`), CONSTRAINT `delivery_option_df_carrier_additional_service_df_delivery_option_df_id` FOREIGN KEY (`delivery_option_df_id`) REFERENCES `delivery_option_dfs` (`id`) ON DELETE CASCADE, CONSTRAINT `delivery_option_df_carrier_additional_service_df_carrier_additional_service_df_id` FOREIGN KEY (`carrier_additional_service_df_id`) REFERENCES `carrier_additional_service_dfs` (`id`) ON DELETE CASCADE);
-- Create "delivery_option_dsv_carrier_additional_service_dsv" table
CREATE TABLE `delivery_option_dsv_carrier_additional_service_dsv` (`delivery_option_dsv_id` text NOT NULL, `carrier_additional_service_dsv_id` text NOT NULL, PRIMARY KEY (`delivery_option_dsv_id`, `carrier_additional_service_dsv_id`), CONSTRAINT `delivery_option_dsv_carrier_additional_service_dsv_delivery_option_dsv_id` FOREIGN KEY (`delivery_option_dsv_id`) REFERENCES `delivery_option_ds_vs` (`id`) ON DELETE CASCADE, CONSTRAINT `delivery_option_dsv_carrier_additional_service_dsv_carrier_additional_service_dsv_id` FOREIGN KEY (`carrier_additional_service_dsv_id`) REFERENCES `carrier_additional_service_ds_vs` (`id`) ON DELETE CASCADE);
-- Create "delivery_option_easy_post_carrier_add_serv_easy_post" table
CREATE TABLE `delivery_option_easy_post_carrier_add_serv_easy_post` (`delivery_option_easy_post_id` text NOT NULL, `carrier_additional_service_easy_post_id` text NOT NULL, PRIMARY KEY (`delivery_option_easy_post_id`, `carrier_additional_service_easy_post_id`), CONSTRAINT `delivery_option_easy_post_carrier_add_serv_easy_post_delivery_option_easy_post_id` FOREIGN KEY (`delivery_option_easy_post_id`) REFERENCES `delivery_option_easy_posts` (`id`) ON DELETE CASCADE, CONSTRAINT `delivery_option_easy_post_carrier_add_serv_easy_post_carrier_additional_service_easy_post_id` FOREIGN KEY (`carrier_additional_service_easy_post_id`) REFERENCES `carrier_additional_service_easy_posts` (`id`) ON DELETE CASCADE);
-- Create "delivery_option_gls_carrier_additional_service_gls" table
CREATE TABLE `delivery_option_gls_carrier_additional_service_gls` (`delivery_option_gls_id` text NOT NULL, `carrier_additional_service_gls_id` text NOT NULL, PRIMARY KEY (`delivery_option_gls_id`, `carrier_additional_service_gls_id`), CONSTRAINT `delivery_option_gls_carrier_additional_service_gls_delivery_option_gls_id` FOREIGN KEY (`delivery_option_gls_id`) REFERENCES `delivery_option_gl_ss` (`id`) ON DELETE CASCADE, CONSTRAINT `delivery_option_gls_carrier_additional_service_gls_carrier_additional_service_gls_id` FOREIGN KEY (`carrier_additional_service_gls_id`) REFERENCES `carrier_additional_service_gl_ss` (`id`) ON DELETE CASCADE);
-- Create "delivery_option_post_nord_carrier_add_serv_post_nord" table
CREATE TABLE `delivery_option_post_nord_carrier_add_serv_post_nord` (`delivery_option_post_nord_id` text NOT NULL, `carrier_additional_service_post_nord_id` text NOT NULL, PRIMARY KEY (`delivery_option_post_nord_id`, `carrier_additional_service_post_nord_id`), CONSTRAINT `delivery_option_post_nord_carrier_add_serv_post_nord_delivery_option_post_nord_id` FOREIGN KEY (`delivery_option_post_nord_id`) REFERENCES `delivery_option_post_nords` (`id`) ON DELETE CASCADE, CONSTRAINT `delivery_option_post_nord_carrier_add_serv_post_nord_carrier_additional_service_post_nord_id` FOREIGN KEY (`carrier_additional_service_post_nord_id`) REFERENCES `carrier_additional_service_post_nords` (`id`) ON DELETE CASCADE);
-- Create "delivery_option_usps_carrier_additional_service_usps" table
CREATE TABLE `delivery_option_usps_carrier_additional_service_usps` (`delivery_option_usps_id` text NOT NULL, `carrier_additional_service_usps_id` text NOT NULL, PRIMARY KEY (`delivery_option_usps_id`, `carrier_additional_service_usps_id`), CONSTRAINT `delivery_option_usps_carrier_additional_service_usps_delivery_option_usps_id` FOREIGN KEY (`delivery_option_usps_id`) REFERENCES `delivery_option_usp_ss` (`id`) ON DELETE CASCADE, CONSTRAINT `delivery_option_usps_carrier_additional_service_usps_carrier_additional_service_usps_id` FOREIGN KEY (`carrier_additional_service_usps_id`) REFERENCES `carrier_additional_service_usp_ss` (`id`) ON DELETE CASCADE);
-- Create "hypothesis_test_delivery_option_delivery_option_group_one" table
CREATE TABLE `hypothesis_test_delivery_option_delivery_option_group_one` (`hypothesis_test_delivery_option_id` text NOT NULL, `delivery_option_id` text NOT NULL, PRIMARY KEY (`hypothesis_test_delivery_option_id`, `delivery_option_id`), CONSTRAINT `hypothesis_test_delivery_option_delivery_option_group_one_hypothesis_test_delivery_option_id` FOREIGN KEY (`hypothesis_test_delivery_option_id`) REFERENCES `hypothesis_test_delivery_options` (`id`) ON DELETE CASCADE, CONSTRAINT `hypothesis_test_delivery_option_delivery_option_group_one_delivery_option_id` FOREIGN KEY (`delivery_option_id`) REFERENCES `delivery_options` (`id`) ON DELETE CASCADE);
-- Create "hypothesis_test_delivery_option_delivery_option_group_two" table
CREATE TABLE `hypothesis_test_delivery_option_delivery_option_group_two` (`hypothesis_test_delivery_option_id` text NOT NULL, `delivery_option_id` text NOT NULL, PRIMARY KEY (`hypothesis_test_delivery_option_id`, `delivery_option_id`), CONSTRAINT `hypothesis_test_delivery_option_delivery_option_group_two_hypothesis_test_delivery_option_id` FOREIGN KEY (`hypothesis_test_delivery_option_id`) REFERENCES `hypothesis_test_delivery_options` (`id`) ON DELETE CASCADE, CONSTRAINT `hypothesis_test_delivery_option_delivery_option_group_two_delivery_option_id` FOREIGN KEY (`delivery_option_id`) REFERENCES `delivery_options` (`id`) ON DELETE CASCADE);
-- Create "location_location_tags" table
CREATE TABLE `location_location_tags` (`location_id` text NOT NULL, `location_tag_id` text NOT NULL, PRIMARY KEY (`location_id`, `location_tag_id`), CONSTRAINT `location_location_tags_location_id` FOREIGN KEY (`location_id`) REFERENCES `locations` (`id`) ON DELETE CASCADE, CONSTRAINT `location_location_tags_location_tag_id` FOREIGN KEY (`location_tag_id`) REFERENCES `location_tags` (`id`) ON DELETE CASCADE);
-- Create "pallet_cancelled_shipment_pallet" table
CREATE TABLE `pallet_cancelled_shipment_pallet` (`pallet_id` text NOT NULL, `shipment_pallet_id` text NOT NULL, PRIMARY KEY (`pallet_id`, `shipment_pallet_id`), CONSTRAINT `pallet_cancelled_shipment_pallet_pallet_id` FOREIGN KEY (`pallet_id`) REFERENCES `pallets` (`id`) ON DELETE CASCADE, CONSTRAINT `pallet_cancelled_shipment_pallet_shipment_pallet_id` FOREIGN KEY (`shipment_pallet_id`) REFERENCES `shipment_pallets` (`id`) ON DELETE CASCADE);
-- Create "product_image_product_variant" table
CREATE TABLE `product_image_product_variant` (`product_image_id` text NOT NULL, `product_variant_id` text NOT NULL, PRIMARY KEY (`product_image_id`, `product_variant_id`), CONSTRAINT `product_image_product_variant_product_image_id` FOREIGN KEY (`product_image_id`) REFERENCES `product_images` (`id`) ON DELETE CASCADE, CONSTRAINT `product_image_product_variant_product_variant_id` FOREIGN KEY (`product_variant_id`) REFERENCES `product_variants` (`id`) ON DELETE CASCADE);
-- Create "product_tag_products" table
CREATE TABLE `product_tag_products` (`product_tag_id` text NOT NULL, `product_id` text NOT NULL, PRIMARY KEY (`product_tag_id`, `product_id`), CONSTRAINT `product_tag_products_product_tag_id` FOREIGN KEY (`product_tag_id`) REFERENCES `product_tags` (`id`) ON DELETE CASCADE, CONSTRAINT `product_tag_products_product_id` FOREIGN KEY (`product_id`) REFERENCES `products` (`id`) ON DELETE CASCADE);
-- Create "return_portal_return_location" table
CREATE TABLE `return_portal_return_location` (`return_portal_id` text NOT NULL, `location_id` text NOT NULL, PRIMARY KEY (`return_portal_id`, `location_id`), CONSTRAINT `return_portal_return_location_return_portal_id` FOREIGN KEY (`return_portal_id`) REFERENCES `return_portals` (`id`) ON DELETE CASCADE, CONSTRAINT `return_portal_return_location_location_id` FOREIGN KEY (`location_id`) REFERENCES `locations` (`id`) ON DELETE CASCADE);
-- Create "return_portal_delivery_options" table
CREATE TABLE `return_portal_delivery_options` (`return_portal_id` text NOT NULL, `delivery_option_id` text NOT NULL, PRIMARY KEY (`return_portal_id`, `delivery_option_id`), CONSTRAINT `return_portal_delivery_options_return_portal_id` FOREIGN KEY (`return_portal_id`) REFERENCES `return_portals` (`id`) ON DELETE CASCADE, CONSTRAINT `return_portal_delivery_options_delivery_option_id` FOREIGN KEY (`delivery_option_id`) REFERENCES `delivery_options` (`id`) ON DELETE CASCADE);
-- Create "shipment_old_consolidation" table
CREATE TABLE `shipment_old_consolidation` (`shipment_id` text NOT NULL, `consolidation_id` text NOT NULL, PRIMARY KEY (`shipment_id`, `consolidation_id`), CONSTRAINT `shipment_old_consolidation_shipment_id` FOREIGN KEY (`shipment_id`) REFERENCES `shipments` (`id`) ON DELETE CASCADE, CONSTRAINT `shipment_old_consolidation_consolidation_id` FOREIGN KEY (`consolidation_id`) REFERENCES `consolidations` (`id`) ON DELETE CASCADE);
-- Create "tenant_connect_option_carriers" table
CREATE TABLE `tenant_connect_option_carriers` (`tenant_id` text NOT NULL, `connect_option_carrier_id` text NOT NULL, PRIMARY KEY (`tenant_id`, `connect_option_carrier_id`), CONSTRAINT `tenant_connect_option_carriers_tenant_id` FOREIGN KEY (`tenant_id`) REFERENCES `tenants` (`id`) ON DELETE CASCADE, CONSTRAINT `tenant_connect_option_carriers_connect_option_carrier_id` FOREIGN KEY (`connect_option_carrier_id`) REFERENCES `connect_option_carriers` (`id`) ON DELETE CASCADE);
-- Create "tenant_connect_option_platforms" table
CREATE TABLE `tenant_connect_option_platforms` (`tenant_id` text NOT NULL, `connect_option_platform_id` text NOT NULL, PRIMARY KEY (`tenant_id`, `connect_option_platform_id`), CONSTRAINT `tenant_connect_option_platforms_tenant_id` FOREIGN KEY (`tenant_id`) REFERENCES `tenants` (`id`) ON DELETE CASCADE, CONSTRAINT `tenant_connect_option_platforms_connect_option_platform_id` FOREIGN KEY (`connect_option_platform_id`) REFERENCES `connect_option_platforms` (`id`) ON DELETE CASCADE);
