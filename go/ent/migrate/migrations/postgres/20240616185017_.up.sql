-- Create "access_rights" table
CREATE TABLE "access_rights" ("id" character varying NOT NULL, "label" character varying NOT NULL, "internal_id" character varying NOT NULL, PRIMARY KEY ("id"));
-- Create index "access_rights_internal_id_key" to table: "access_rights"
CREATE UNIQUE INDEX "access_rights_internal_id_key" ON "access_rights" ("internal_id");
-- Create index "access_rights_label_key" to table: "access_rights"
CREATE UNIQUE INDEX "access_rights_label_key" ON "access_rights" ("label");
-- Create "address_globals" table
CREATE TABLE "address_globals" ("id" character varying NOT NULL, "uniqueness_id" character varying NULL, "company" character varying NULL, "address_one" character varying NOT NULL, "address_two" character varying NULL, "city" character varying NOT NULL, "state" character varying NULL, "zip" character varying NOT NULL, "latitude" double precision NOT NULL DEFAULT 0, "longitude" double precision NOT NULL DEFAULT 0, "address_global_country" character varying NOT NULL, "parcel_shop_address" character varying NULL, "parcel_shop_bring_address_delivery" character varying NULL, "parcel_shop_post_nord_address_delivery" character varying NULL, PRIMARY KEY ("id"));
-- Create index "address_globals_parcel_shop_address_key" to table: "address_globals"
CREATE UNIQUE INDEX "address_globals_parcel_shop_address_key" ON "address_globals" ("parcel_shop_address");
-- Create index "address_globals_parcel_shop_bring_address_delivery_key" to table: "address_globals"
CREATE UNIQUE INDEX "address_globals_parcel_shop_bring_address_delivery_key" ON "address_globals" ("parcel_shop_bring_address_delivery");
-- Create index "address_globals_parcel_shop_post_nord_address_delivery_key" to table: "address_globals"
CREATE UNIQUE INDEX "address_globals_parcel_shop_post_nord_address_delivery_key" ON "address_globals" ("parcel_shop_post_nord_address_delivery");
-- Create index "address_globals_uniqueness_id_key" to table: "address_globals"
CREATE UNIQUE INDEX "address_globals_uniqueness_id_key" ON "address_globals" ("uniqueness_id");
-- Create "addresses" table
CREATE TABLE "addresses" ("id" character varying NOT NULL, "uniqueness_id" character varying NULL, "first_name" character varying NOT NULL, "last_name" character varying NOT NULL, "email" character varying NOT NULL, "phone_number" character varying NOT NULL, "phone_number_2" character varying NULL, "vat_number" character varying NULL, "company" character varying NULL, "address_one" character varying NOT NULL, "address_two" character varying NOT NULL, "city" character varying NOT NULL, "state" character varying NULL, "zip" character varying NOT NULL, "tenant_id" character varying NOT NULL, "address_country" character varying NOT NULL, "consolidation_recipient" character varying NULL, "consolidation_sender" character varying NULL, PRIMARY KEY ("id"));
-- Create index "address_tenant_id" to table: "addresses"
CREATE INDEX "address_tenant_id" ON "addresses" ("tenant_id");
-- Create index "addresses_consolidation_recipient_key" to table: "addresses"
CREATE UNIQUE INDEX "addresses_consolidation_recipient_key" ON "addresses" ("consolidation_recipient");
-- Create index "addresses_consolidation_sender_key" to table: "addresses"
CREATE UNIQUE INDEX "addresses_consolidation_sender_key" ON "addresses" ("consolidation_sender");
-- Create index "addresses_uniqueness_id_key" to table: "addresses"
CREATE UNIQUE INDEX "addresses_uniqueness_id_key" ON "addresses" ("uniqueness_id");
-- Create "api_tokens" table
CREATE TABLE "api_tokens" ("id" character varying NOT NULL, "name" character varying NOT NULL, "hashed_token" character varying NOT NULL, "created_at" timestamptz NULL, "last_used" timestamptz NULL, "tenant_id" character varying NOT NULL, "user_api_token" character varying NOT NULL, PRIMARY KEY ("id"));
-- Create index "api_tokens_hashed_token_key" to table: "api_tokens"
CREATE UNIQUE INDEX "api_tokens_hashed_token_key" ON "api_tokens" ("hashed_token");
-- Create index "apitoken_tenant_id" to table: "api_tokens"
CREATE INDEX "apitoken_tenant_id" ON "api_tokens" ("tenant_id");
-- Create "business_hours_periods" table
CREATE TABLE "business_hours_periods" ("id" character varying NOT NULL, "day_of_week" character varying NOT NULL, "opening" timestamptz NOT NULL, "closing" timestamptz NOT NULL, "parcel_shop_business_hours_period" character varying NOT NULL, PRIMARY KEY ("id"));
-- Create "carrier_additional_service_brings" table
CREATE TABLE "carrier_additional_service_brings" ("id" character varying NOT NULL, "label" character varying NOT NULL, "api_code_booking" character varying NOT NULL, "carrier_service_bring_carrier_additional_service_bring" character varying NULL, PRIMARY KEY ("id"));
-- Create "carrier_additional_service_da_os" table
CREATE TABLE "carrier_additional_service_da_os" ("id" character varying NOT NULL, "label" character varying NOT NULL, "api_code" character varying NOT NULL, PRIMARY KEY ("id"));
-- Create "carrier_additional_service_dfs" table
CREATE TABLE "carrier_additional_service_dfs" ("id" character varying NOT NULL, "label" character varying NOT NULL, "api_code" character varying NOT NULL, PRIMARY KEY ("id"));
-- Create "carrier_additional_service_ds_vs" table
CREATE TABLE "carrier_additional_service_ds_vs" ("id" character varying NOT NULL, "label" character varying NOT NULL, "api_code" character varying NOT NULL, PRIMARY KEY ("id"));
-- Create "carrier_additional_service_easy_posts" table
CREATE TABLE "carrier_additional_service_easy_posts" ("id" character varying NOT NULL, "label" character varying NOT NULL, "api_key" character varying NOT NULL, "api_value" character varying NOT NULL, PRIMARY KEY ("id"));
-- Create "carrier_additional_service_gl_ss" table
CREATE TABLE "carrier_additional_service_gl_ss" ("id" character varying NOT NULL, "label" character varying NOT NULL, "mandatory" boolean NOT NULL, "all_countries_consignor" boolean NOT NULL DEFAULT false, "all_countries_consignee" boolean NOT NULL DEFAULT false, "internal_id" character varying NOT NULL, "carrier_service_gls_carrier_additional_service_gls" character varying NULL, PRIMARY KEY ("id"));
-- Create "carrier_additional_service_gls_countries_consignee" table
CREATE TABLE "carrier_additional_service_gls_countries_consignee" ("carrier_additional_service_gls_id" character varying NOT NULL, "country_id" character varying NOT NULL, PRIMARY KEY ("carrier_additional_service_gls_id", "country_id"));
-- Create "carrier_additional_service_gls_countries_consignor" table
CREATE TABLE "carrier_additional_service_gls_countries_consignor" ("carrier_additional_service_gls_id" character varying NOT NULL, "country_id" character varying NOT NULL, PRIMARY KEY ("carrier_additional_service_gls_id", "country_id"));
-- Create "carrier_additional_service_post_nord_countries_consignee" table
CREATE TABLE "carrier_additional_service_post_nord_countries_consignee" ("carrier_additional_service_post_nord_id" character varying NOT NULL, "country_id" character varying NOT NULL, PRIMARY KEY ("carrier_additional_service_post_nord_id", "country_id"));
-- Create "carrier_additional_service_post_nord_countries_consignor" table
CREATE TABLE "carrier_additional_service_post_nord_countries_consignor" ("carrier_additional_service_post_nord_id" character varying NOT NULL, "country_id" character varying NOT NULL, PRIMARY KEY ("carrier_additional_service_post_nord_id", "country_id"));
-- Create "carrier_additional_service_post_nords" table
CREATE TABLE "carrier_additional_service_post_nords" ("id" character varying NOT NULL, "label" character varying NOT NULL, "mandatory" boolean NOT NULL, "all_countries_consignor" boolean NOT NULL DEFAULT false, "all_countries_consignee" boolean NOT NULL DEFAULT false, "internal_id" character varying NOT NULL, "api_code" character varying NOT NULL, "carrier_service_post_nord_carrier_add_serv_post_nord" character varying NULL, PRIMARY KEY ("id"));
-- Create index "carrieradditionalservicepostnord_internal_id_carrier_service_po" to table: "carrier_additional_service_post_nords"
CREATE UNIQUE INDEX "carrieradditionalservicepostnord_internal_id_carrier_service_po" ON "carrier_additional_service_post_nords" ("internal_id", "carrier_service_post_nord_carrier_add_serv_post_nord");
-- Create "carrier_additional_service_usp_ss" table
CREATE TABLE "carrier_additional_service_usp_ss" ("id" character varying NOT NULL, "label" character varying NOT NULL, "commonly_used" boolean NOT NULL DEFAULT false, "internal_id" character varying NOT NULL, "api_code" character varying NOT NULL, "carrier_service_usps_carrier_additional_service_usps" character varying NULL, PRIMARY KEY ("id"));
-- Create "carrier_brands" table
CREATE TABLE "carrier_brands" ("id" character varying NOT NULL, "label" character varying NOT NULL, "label_short" character varying NOT NULL, "internal_id" character varying NOT NULL, "logo_url" character varying NULL, "text_color" character varying NULL DEFAULT '#FFFFFF', "background_color" character varying NULL DEFAULT '#000000', PRIMARY KEY ("id"));
-- Create index "carrierbrand_internal_id" to table: "carrier_brands"
CREATE UNIQUE INDEX "carrierbrand_internal_id" ON "carrier_brands" ("internal_id");
-- Create "carrier_brings" table
CREATE TABLE "carrier_brings" ("id" character varying NOT NULL, "api_key" character varying NULL, "customer_number" character varying NULL, "test" boolean NOT NULL DEFAULT true, "carrier_carrier_bring" character varying NOT NULL, "tenant_id" character varying NOT NULL, PRIMARY KEY ("id"));
-- Create index "carrier_brings_carrier_carrier_bring_key" to table: "carrier_brings"
CREATE UNIQUE INDEX "carrier_brings_carrier_carrier_bring_key" ON "carrier_brings" ("carrier_carrier_bring");
-- Create index "carrierbring_tenant_id" to table: "carrier_brings"
CREATE INDEX "carrierbring_tenant_id" ON "carrier_brings" ("tenant_id");
-- Create "carrier_da_os" table
CREATE TABLE "carrier_da_os" ("id" character varying NOT NULL, "customer_id" character varying NULL, "api_key" character varying NULL, "test" boolean NOT NULL DEFAULT true, "carrier_carrier_dao" character varying NOT NULL, "tenant_id" character varying NOT NULL, PRIMARY KEY ("id"));
-- Create index "carrier_da_os_carrier_carrier_dao_key" to table: "carrier_da_os"
CREATE UNIQUE INDEX "carrier_da_os_carrier_carrier_dao_key" ON "carrier_da_os" ("carrier_carrier_dao");
-- Create index "carrierdao_tenant_id" to table: "carrier_da_os"
CREATE INDEX "carrierdao_tenant_id" ON "carrier_da_os" ("tenant_id");
-- Create "carrier_dfs" table
CREATE TABLE "carrier_dfs" ("id" character varying NOT NULL, "customer_id" character varying NOT NULL, "agreement_number" character varying NOT NULL, "who_pays" character varying NOT NULL DEFAULT 'Prepaid', "test" boolean NOT NULL DEFAULT true, "carrier_carrier_df" character varying NOT NULL, "tenant_id" character varying NOT NULL, PRIMARY KEY ("id"));
-- Create index "carrier_dfs_carrier_carrier_df_key" to table: "carrier_dfs"
CREATE UNIQUE INDEX "carrier_dfs_carrier_carrier_df_key" ON "carrier_dfs" ("carrier_carrier_df");
-- Create index "carrierdf_tenant_id" to table: "carrier_dfs"
CREATE INDEX "carrierdf_tenant_id" ON "carrier_dfs" ("tenant_id");
-- Create "carrier_ds_vs" table
CREATE TABLE "carrier_ds_vs" ("id" character varying NOT NULL, "carrier_carrier_dsv" character varying NOT NULL, "tenant_id" character varying NOT NULL, PRIMARY KEY ("id"));
-- Create index "carrier_ds_vs_carrier_carrier_dsv_key" to table: "carrier_ds_vs"
CREATE UNIQUE INDEX "carrier_ds_vs_carrier_carrier_dsv_key" ON "carrier_ds_vs" ("carrier_carrier_dsv");
-- Create index "carrierdsv_tenant_id" to table: "carrier_ds_vs"
CREATE INDEX "carrierdsv_tenant_id" ON "carrier_ds_vs" ("tenant_id");
-- Create "carrier_easy_posts" table
CREATE TABLE "carrier_easy_posts" ("id" character varying NOT NULL, "api_key" character varying NOT NULL, "test" boolean NOT NULL DEFAULT true, "carrier_accounts" jsonb NOT NULL, "carrier_carrier_easy_post" character varying NOT NULL, "tenant_id" character varying NOT NULL, PRIMARY KEY ("id"));
-- Create index "carrier_easy_posts_carrier_carrier_easy_post_key" to table: "carrier_easy_posts"
CREATE UNIQUE INDEX "carrier_easy_posts_carrier_carrier_easy_post_key" ON "carrier_easy_posts" ("carrier_carrier_easy_post");
-- Create index "carriereasypost_tenant_id" to table: "carrier_easy_posts"
CREATE INDEX "carriereasypost_tenant_id" ON "carrier_easy_posts" ("tenant_id");
-- Create "carrier_gl_ss" table
CREATE TABLE "carrier_gl_ss" ("id" character varying NOT NULL, "contact_id" character varying NULL, "gls_username" character varying NULL, "gls_password" character varying NULL, "customer_id" character varying NULL, "gls_country_code" character varying NULL, "sync_shipment_cancellation" boolean NULL DEFAULT false, "print_error_on_label" boolean NULL DEFAULT false, "carrier_carrier_gls" character varying NOT NULL, "tenant_id" character varying NOT NULL, PRIMARY KEY ("id"));
-- Create index "carrier_gl_ss_carrier_carrier_gls_key" to table: "carrier_gl_ss"
CREATE UNIQUE INDEX "carrier_gl_ss_carrier_carrier_gls_key" ON "carrier_gl_ss" ("carrier_carrier_gls");
-- Create index "carriergls_tenant_id" to table: "carrier_gl_ss"
CREATE INDEX "carriergls_tenant_id" ON "carrier_gl_ss" ("tenant_id");
-- Create "carrier_post_nords" table
CREATE TABLE "carrier_post_nords" ("id" character varying NOT NULL, "customer_number" character varying NOT NULL DEFAULT '', "carrier_carrier_post_nord" character varying NOT NULL, "tenant_id" character varying NOT NULL, PRIMARY KEY ("id"));
-- Create index "carrier_post_nords_carrier_carrier_post_nord_key" to table: "carrier_post_nords"
CREATE UNIQUE INDEX "carrier_post_nords_carrier_carrier_post_nord_key" ON "carrier_post_nords" ("carrier_carrier_post_nord");
-- Create index "carrierpostnord_tenant_id" to table: "carrier_post_nords"
CREATE INDEX "carrierpostnord_tenant_id" ON "carrier_post_nords" ("tenant_id");
-- Create "carrier_service_brings" table
CREATE TABLE "carrier_service_brings" ("id" character varying NOT NULL, "api_service_code" character varying NOT NULL, "api_request" character varying NOT NULL, "carrier_service_carrier_service_bring" character varying NOT NULL, PRIMARY KEY ("id"));
-- Create index "carrier_service_brings_carrier_service_carrier_service_bring_ke" to table: "carrier_service_brings"
CREATE UNIQUE INDEX "carrier_service_brings_carrier_service_carrier_service_bring_ke" ON "carrier_service_brings" ("carrier_service_carrier_service_bring");
-- Create "carrier_service_da_os" table
CREATE TABLE "carrier_service_da_os" ("id" character varying NOT NULL, "carrier_service_carrier_service_dao" character varying NOT NULL, PRIMARY KEY ("id"));
-- Create index "carrier_service_da_os_carrier_service_carrier_service_dao_key" to table: "carrier_service_da_os"
CREATE UNIQUE INDEX "carrier_service_da_os_carrier_service_carrier_service_dao_key" ON "carrier_service_da_os" ("carrier_service_carrier_service_dao");
-- Create "carrier_service_dao_carrier_additional_service_dao" table
CREATE TABLE "carrier_service_dao_carrier_additional_service_dao" ("carrier_service_dao_id" character varying NOT NULL, "carrier_additional_service_dao_id" character varying NOT NULL, PRIMARY KEY ("carrier_service_dao_id", "carrier_additional_service_dao_id"));
-- Create "carrier_service_df_carrier_additional_service_df" table
CREATE TABLE "carrier_service_df_carrier_additional_service_df" ("carrier_service_df_id" character varying NOT NULL, "carrier_additional_service_df_id" character varying NOT NULL, PRIMARY KEY ("carrier_service_df_id", "carrier_additional_service_df_id"));
-- Create "carrier_service_dfs" table
CREATE TABLE "carrier_service_dfs" ("id" character varying NOT NULL, "carrier_service_carrier_service_df" character varying NOT NULL, PRIMARY KEY ("id"));
-- Create index "carrier_service_dfs_carrier_service_carrier_service_df_key" to table: "carrier_service_dfs"
CREATE UNIQUE INDEX "carrier_service_dfs_carrier_service_carrier_service_df_key" ON "carrier_service_dfs" ("carrier_service_carrier_service_df");
-- Create "carrier_service_ds_vs" table
CREATE TABLE "carrier_service_ds_vs" ("id" character varying NOT NULL, "carrier_service_carrier_service_dsv" character varying NOT NULL, PRIMARY KEY ("id"));
-- Create index "carrier_service_ds_vs_carrier_service_carrier_service_dsv_key" to table: "carrier_service_ds_vs"
CREATE UNIQUE INDEX "carrier_service_ds_vs_carrier_service_carrier_service_dsv_key" ON "carrier_service_ds_vs" ("carrier_service_carrier_service_dsv");
-- Create "carrier_service_dsv_carrier_additional_service_dsv" table
CREATE TABLE "carrier_service_dsv_carrier_additional_service_dsv" ("carrier_service_dsv_id" character varying NOT NULL, "carrier_additional_service_dsv_id" character varying NOT NULL, PRIMARY KEY ("carrier_service_dsv_id", "carrier_additional_service_dsv_id"));
-- Create "carrier_service_easy_post_carrier_add_serv_easy_post" table
CREATE TABLE "carrier_service_easy_post_carrier_add_serv_easy_post" ("carrier_service_easy_post_id" character varying NOT NULL, "carrier_additional_service_easy_post_id" character varying NOT NULL, PRIMARY KEY ("carrier_service_easy_post_id", "carrier_additional_service_easy_post_id"));
-- Create "carrier_service_easy_posts" table
CREATE TABLE "carrier_service_easy_posts" ("id" character varying NOT NULL, "api_key" character varying NOT NULL, "carrier_service_carrier_serv_easy_post" character varying NOT NULL, PRIMARY KEY ("id"));
-- Create index "carrier_service_easy_posts_carrier_service_carrier_serv_easy_po" to table: "carrier_service_easy_posts"
CREATE UNIQUE INDEX "carrier_service_easy_posts_carrier_service_carrier_serv_easy_po" ON "carrier_service_easy_posts" ("carrier_service_carrier_serv_easy_post");
-- Create "carrier_service_gl_ss" table
CREATE TABLE "carrier_service_gl_ss" ("id" character varying NOT NULL, "api_key" character varying NULL, "api_value" character varying NOT NULL, "carrier_service_carrier_service_gls" character varying NOT NULL, PRIMARY KEY ("id"));
-- Create index "carrier_service_gl_ss_api_key_key" to table: "carrier_service_gl_ss"
CREATE UNIQUE INDEX "carrier_service_gl_ss_api_key_key" ON "carrier_service_gl_ss" ("api_key");
-- Create index "carrier_service_gl_ss_carrier_service_carrier_service_gls_key" to table: "carrier_service_gl_ss"
CREATE UNIQUE INDEX "carrier_service_gl_ss_carrier_service_carrier_service_gls_key" ON "carrier_service_gl_ss" ("carrier_service_carrier_service_gls");
-- Create "carrier_service_post_nords" table
CREATE TABLE "carrier_service_post_nords" ("id" character varying NOT NULL, "label" character varying NOT NULL, "internal_id" character varying NOT NULL, "api_code" character varying NOT NULL, "carrier_service_carrier_service_post_nord" character varying NOT NULL, PRIMARY KEY ("id"));
-- Create index "carrier_service_post_nords_carrier_service_carrier_service_post" to table: "carrier_service_post_nords"
CREATE UNIQUE INDEX "carrier_service_post_nords_carrier_service_carrier_service_post" ON "carrier_service_post_nords" ("carrier_service_carrier_service_post_nord");
-- Create index "carrier_service_post_nords_internal_id_key" to table: "carrier_service_post_nords"
CREATE UNIQUE INDEX "carrier_service_post_nords_internal_id_key" ON "carrier_service_post_nords" ("internal_id");
-- Create "carrier_service_usp_ss" table
CREATE TABLE "carrier_service_usp_ss" ("id" character varying NOT NULL, "api_key" character varying NOT NULL, "carrier_service_carrier_service_usps" character varying NOT NULL, PRIMARY KEY ("id"));
-- Create index "carrier_service_usp_ss_carrier_service_carrier_service_usps_key" to table: "carrier_service_usp_ss"
CREATE UNIQUE INDEX "carrier_service_usp_ss_carrier_service_carrier_service_usps_key" ON "carrier_service_usp_ss" ("carrier_service_carrier_service_usps");
-- Create "carrier_services" table
CREATE TABLE "carrier_services" ("id" character varying NOT NULL, "label" character varying NOT NULL, "internal_id" character varying NOT NULL, "return" boolean NOT NULL DEFAULT false, "consolidation" boolean NOT NULL DEFAULT false, "delivery_point_optional" boolean NOT NULL DEFAULT false, "delivery_point_required" boolean NOT NULL DEFAULT false, "carrier_brand_carrier_service" character varying NOT NULL, PRIMARY KEY ("id"));
-- Create index "carrier_services_internal_id_key" to table: "carrier_services"
CREATE UNIQUE INDEX "carrier_services_internal_id_key" ON "carrier_services" ("internal_id");
-- Create "carrier_usp_ss" table
CREATE TABLE "carrier_usp_ss" ("id" character varying NOT NULL, "is_test_api" boolean NOT NULL DEFAULT false, "consumer_key" character varying NULL, "consumer_secret" character varying NULL, "mid" character varying NULL, "manifest_mid" character varying NULL, "crid" character varying NULL, "eps_account_number" character varying NULL, "carrier_carrier_usps" character varying NOT NULL, "tenant_id" character varying NOT NULL, PRIMARY KEY ("id"));
-- Create index "carrier_usp_ss_carrier_carrier_usps_key" to table: "carrier_usp_ss"
CREATE UNIQUE INDEX "carrier_usp_ss_carrier_carrier_usps_key" ON "carrier_usp_ss" ("carrier_carrier_usps");
-- Create index "carrierusps_tenant_id" to table: "carrier_usp_ss"
CREATE INDEX "carrierusps_tenant_id" ON "carrier_usp_ss" ("tenant_id");
-- Create "carriers" table
CREATE TABLE "carriers" ("id" character varying NOT NULL, "name" character varying NOT NULL, "sync_cancelation" boolean NOT NULL DEFAULT false, "tenant_id" character varying NOT NULL, "carrier_carrier_brand" character varying NOT NULL, PRIMARY KEY ("id"));
-- Create index "carrier_tenant_id" to table: "carriers"
CREATE INDEX "carrier_tenant_id" ON "carriers" ("tenant_id");
-- Create "change_histories" table
CREATE TABLE "change_histories" ("id" character varying NOT NULL, "created_at" timestamptz NOT NULL, "origin" character varying NOT NULL DEFAULT 'unknown', "tenant_id" character varying NOT NULL, "change_history_user" character varying NULL, PRIMARY KEY ("id"));
-- Create index "changehistory_tenant_id" to table: "change_histories"
CREATE INDEX "changehistory_tenant_id" ON "change_histories" ("tenant_id");
-- Create "colli_cancelled_shipment_parcel" table
CREATE TABLE "colli_cancelled_shipment_parcel" ("colli_id" character varying NOT NULL, "shipment_parcel_id" character varying NOT NULL, PRIMARY KEY ("colli_id", "shipment_parcel_id"));
-- Create "collis" table
CREATE TABLE "collis" ("id" character varying NOT NULL, "internal_barcode" bigint NULL, "status" character varying NOT NULL, "slip_print_status" character varying NOT NULL DEFAULT 'pending', "created_at" timestamptz NOT NULL, "email_packing_slip_printed_at" timestamptz NULL, "email_label_printed_at" timestamptz NULL, "tenant_id" character varying NOT NULL, "colli_recipient" character varying NOT NULL, "colli_sender" character varying NOT NULL, "colli_parcel_shop" character varying NULL, "colli_click_collect_location" character varying NULL, "colli_delivery_option" character varying NULL, "colli_packaging" character varying NULL, "order_colli" character varying NOT NULL, PRIMARY KEY ("id"));
-- Create index "colli_internal_barcode_tenant_id" to table: "collis"
CREATE UNIQUE INDEX "colli_internal_barcode_tenant_id" ON "collis" ("internal_barcode", "tenant_id");
-- Create index "colli_tenant_id" to table: "collis"
CREATE INDEX "colli_tenant_id" ON "collis" ("tenant_id");
-- Create "connect_option_carriers" table
CREATE TABLE "connect_option_carriers" ("id" character varying NOT NULL, "name" character varying NOT NULL, PRIMARY KEY ("id"));
-- Create index "connect_option_carriers_name_key" to table: "connect_option_carriers"
CREATE UNIQUE INDEX "connect_option_carriers_name_key" ON "connect_option_carriers" ("name");
-- Create "connect_option_platforms" table
CREATE TABLE "connect_option_platforms" ("id" character varying NOT NULL, "name" character varying NOT NULL, PRIMARY KEY ("id"));
-- Create index "connect_option_platforms_name_key" to table: "connect_option_platforms"
CREATE UNIQUE INDEX "connect_option_platforms_name_key" ON "connect_option_platforms" ("name");
-- Create "connection_brands" table
CREATE TABLE "connection_brands" ("id" character varying NOT NULL, "label" character varying NOT NULL, "internal_id" character varying NOT NULL DEFAULT 'shopify', "logo_url" character varying NULL, PRIMARY KEY ("id"));
-- Create index "connection_brands_label_key" to table: "connection_brands"
CREATE UNIQUE INDEX "connection_brands_label_key" ON "connection_brands" ("label");
-- Create "connection_lookups" table
CREATE TABLE "connection_lookups" ("id" character varying NOT NULL, "payload" character varying NOT NULL, "options_output_count" bigint NOT NULL, "error" character varying NULL, "created_at" timestamptz NOT NULL, "tenant_id" character varying NOT NULL, "connection_lookup_connections" character varying NULL, PRIMARY KEY ("id"));
-- Create index "connectionlookup_tenant_id" to table: "connection_lookups"
CREATE INDEX "connectionlookup_tenant_id" ON "connection_lookups" ("tenant_id");
-- Create "connection_shopifies" table
CREATE TABLE "connection_shopifies" ("id" character varying NOT NULL, "rate_integration" boolean NOT NULL DEFAULT false, "store_url" character varying NULL, "api_key" character varying NULL, "lookup_key" character varying NULL, "connection_connection_shopify" character varying NOT NULL, "tenant_id" character varying NOT NULL, PRIMARY KEY ("id"));
-- Create index "connection_shopifies_connection_connection_shopify_key" to table: "connection_shopifies"
CREATE UNIQUE INDEX "connection_shopifies_connection_connection_shopify_key" ON "connection_shopifies" ("connection_connection_shopify");
-- Create index "connection_shopifies_store_url_key" to table: "connection_shopifies"
CREATE UNIQUE INDEX "connection_shopifies_store_url_key" ON "connection_shopifies" ("store_url");
-- Create index "connectionshopify_tenant_id" to table: "connection_shopifies"
CREATE INDEX "connectionshopify_tenant_id" ON "connection_shopifies" ("tenant_id");
-- Create "connections" table
CREATE TABLE "connections" ("id" character varying NOT NULL, "name" character varying NOT NULL, "sync_orders" boolean NOT NULL DEFAULT false, "sync_products" boolean NOT NULL DEFAULT false, "fulfill_automatically" boolean NOT NULL DEFAULT false, "dispatch_automatically" boolean NOT NULL DEFAULT false, "convert_currency" boolean NOT NULL DEFAULT false, "tenant_id" character varying NOT NULL, "connection_connection_brand" character varying NOT NULL, "connection_sender_location" character varying NOT NULL, "connection_pickup_location" character varying NOT NULL, "connection_return_location" character varying NOT NULL, "connection_seller_location" character varying NOT NULL, "connection_currency" character varying NOT NULL, "connection_packing_slip_template" character varying NULL, "return_portal_connection" character varying NULL, PRIMARY KEY ("id"));
-- Create index "connection_tenant_id" to table: "connections"
CREATE INDEX "connection_tenant_id" ON "connections" ("tenant_id");
-- Create index "connections_return_portal_connection_key" to table: "connections"
CREATE UNIQUE INDEX "connections_return_portal_connection_key" ON "connections" ("return_portal_connection");
-- Create "consolidations" table
CREATE TABLE "consolidations" ("id" character varying NOT NULL, "public_id" character varying NOT NULL, "description" character varying NULL, "status" character varying NOT NULL DEFAULT 'Pending', "created_at" timestamptz NULL, "tenant_id" character varying NOT NULL, "consolidation_delivery_option" character varying NULL, "shipment_consolidation" character varying NULL, PRIMARY KEY ("id"));
-- Create index "consolidation_tenant_id" to table: "consolidations"
CREATE INDEX "consolidation_tenant_id" ON "consolidations" ("tenant_id");
-- Create index "consolidations_shipment_consolidation_key" to table: "consolidations"
CREATE UNIQUE INDEX "consolidations_shipment_consolidation_key" ON "consolidations" ("shipment_consolidation");
-- Create "contacts" table
CREATE TABLE "contacts" ("id" character varying NOT NULL, "name" character varying NOT NULL, "surname" character varying NOT NULL, "email" character varying NOT NULL, "phone_number" character varying NOT NULL, "tenant_id" character varying NOT NULL, PRIMARY KEY ("id"));
-- Create index "contact_tenant_id" to table: "contacts"
CREATE INDEX "contact_tenant_id" ON "contacts" ("tenant_id");
-- Create "countries" table
CREATE TABLE "countries" ("id" character varying NOT NULL, "label" character varying NOT NULL, "alpha_2" character varying NOT NULL, "alpha_3" character varying NOT NULL, "code" character varying NOT NULL, "region" character varying NOT NULL, PRIMARY KEY ("id"));
-- Create index "countries_alpha_2_key" to table: "countries"
CREATE UNIQUE INDEX "countries_alpha_2_key" ON "countries" ("alpha_2");
-- Create index "countries_alpha_3_key" to table: "countries"
CREATE UNIQUE INDEX "countries_alpha_3_key" ON "countries" ("alpha_3");
-- Create index "countries_code_key" to table: "countries"
CREATE UNIQUE INDEX "countries_code_key" ON "countries" ("code");
-- Create index "countries_label_key" to table: "countries"
CREATE UNIQUE INDEX "countries_label_key" ON "countries" ("label");
-- Create "country_delivery_rule" table
CREATE TABLE "country_delivery_rule" ("country_id" character varying NOT NULL, "delivery_rule_id" character varying NOT NULL, PRIMARY KEY ("country_id", "delivery_rule_id"));
-- Create "country_harmonized_codes" table
CREATE TABLE "country_harmonized_codes" ("id" character varying NOT NULL, "code" character varying NOT NULL, "tenant_id" character varying NOT NULL, "country_harmonized_code_country" character varying NOT NULL, "inventory_item_country_harmonized_code" character varying NOT NULL, PRIMARY KEY ("id"));
-- Create index "countryharmonizedcode_tenant_id" to table: "country_harmonized_codes"
CREATE INDEX "countryharmonizedcode_tenant_id" ON "country_harmonized_codes" ("tenant_id");
-- Create "currencies" table
CREATE TABLE "currencies" ("id" character varying NOT NULL, "display" character varying NOT NULL, "currency_code" character varying NOT NULL DEFAULT 'DKK', PRIMARY KEY ("id"));
-- Create index "currencies_display_key" to table: "currencies"
CREATE UNIQUE INDEX "currencies_display_key" ON "currencies" ("display");
-- Create index "currency_currency_code" to table: "currencies"
CREATE UNIQUE INDEX "currency_currency_code" ON "currencies" ("currency_code");
-- Create "delivery_option_bring_carrier_additional_service_bring" table
CREATE TABLE "delivery_option_bring_carrier_additional_service_bring" ("delivery_option_bring_id" character varying NOT NULL, "carrier_additional_service_bring_id" character varying NOT NULL, PRIMARY KEY ("delivery_option_bring_id", "carrier_additional_service_bring_id"));
-- Create "delivery_option_brings" table
CREATE TABLE "delivery_option_brings" ("id" character varying NOT NULL, "electronic_customs" boolean NOT NULL DEFAULT false, "delivery_option_delivery_option_bring" character varying NOT NULL, "tenant_id" character varying NOT NULL, PRIMARY KEY ("id"));
-- Create index "delivery_option_brings_delivery_option_delivery_option_bring_ke" to table: "delivery_option_brings"
CREATE UNIQUE INDEX "delivery_option_brings_delivery_option_delivery_option_bring_ke" ON "delivery_option_brings" ("delivery_option_delivery_option_bring");
-- Create index "deliveryoptionbring_tenant_id" to table: "delivery_option_brings"
CREATE INDEX "deliveryoptionbring_tenant_id" ON "delivery_option_brings" ("tenant_id");
-- Create "delivery_option_click_collect_location" table
CREATE TABLE "delivery_option_click_collect_location" ("delivery_option_id" character varying NOT NULL, "location_id" character varying NOT NULL, PRIMARY KEY ("delivery_option_id", "location_id"));
-- Create "delivery_option_da_os" table
CREATE TABLE "delivery_option_da_os" ("id" character varying NOT NULL, "delivery_option_delivery_option_dao" character varying NOT NULL, "tenant_id" character varying NOT NULL, PRIMARY KEY ("id"));
-- Create index "delivery_option_da_os_delivery_option_delivery_option_dao_key" to table: "delivery_option_da_os"
CREATE UNIQUE INDEX "delivery_option_da_os_delivery_option_delivery_option_dao_key" ON "delivery_option_da_os" ("delivery_option_delivery_option_dao");
-- Create index "deliveryoptiondao_tenant_id" to table: "delivery_option_da_os"
CREATE INDEX "deliveryoptiondao_tenant_id" ON "delivery_option_da_os" ("tenant_id");
-- Create "delivery_option_dao_carrier_additional_service_dao" table
CREATE TABLE "delivery_option_dao_carrier_additional_service_dao" ("delivery_option_dao_id" character varying NOT NULL, "carrier_additional_service_dao_id" character varying NOT NULL, PRIMARY KEY ("delivery_option_dao_id", "carrier_additional_service_dao_id"));
-- Create "delivery_option_df_carrier_additional_service_df" table
CREATE TABLE "delivery_option_df_carrier_additional_service_df" ("delivery_option_df_id" character varying NOT NULL, "carrier_additional_service_df_id" character varying NOT NULL, PRIMARY KEY ("delivery_option_df_id", "carrier_additional_service_df_id"));
-- Create "delivery_option_dfs" table
CREATE TABLE "delivery_option_dfs" ("id" character varying NOT NULL, "delivery_option_delivery_option_df" character varying NOT NULL, "tenant_id" character varying NOT NULL, PRIMARY KEY ("id"));
-- Create index "delivery_option_dfs_delivery_option_delivery_option_df_key" to table: "delivery_option_dfs"
CREATE UNIQUE INDEX "delivery_option_dfs_delivery_option_delivery_option_df_key" ON "delivery_option_dfs" ("delivery_option_delivery_option_df");
-- Create index "deliveryoptiondf_tenant_id" to table: "delivery_option_dfs"
CREATE INDEX "deliveryoptiondf_tenant_id" ON "delivery_option_dfs" ("tenant_id");
-- Create "delivery_option_ds_vs" table
CREATE TABLE "delivery_option_ds_vs" ("id" character varying NOT NULL, "delivery_option_delivery_option_dsv" character varying NOT NULL, "tenant_id" character varying NOT NULL, PRIMARY KEY ("id"));
-- Create index "delivery_option_ds_vs_delivery_option_delivery_option_dsv_key" to table: "delivery_option_ds_vs"
CREATE UNIQUE INDEX "delivery_option_ds_vs_delivery_option_delivery_option_dsv_key" ON "delivery_option_ds_vs" ("delivery_option_delivery_option_dsv");
-- Create index "deliveryoptiondsv_tenant_id" to table: "delivery_option_ds_vs"
CREATE INDEX "deliveryoptiondsv_tenant_id" ON "delivery_option_ds_vs" ("tenant_id");
-- Create "delivery_option_dsv_carrier_additional_service_dsv" table
CREATE TABLE "delivery_option_dsv_carrier_additional_service_dsv" ("delivery_option_dsv_id" character varying NOT NULL, "carrier_additional_service_dsv_id" character varying NOT NULL, PRIMARY KEY ("delivery_option_dsv_id", "carrier_additional_service_dsv_id"));
-- Create "delivery_option_easy_post_carrier_add_serv_easy_post" table
CREATE TABLE "delivery_option_easy_post_carrier_add_serv_easy_post" ("delivery_option_easy_post_id" character varying NOT NULL, "carrier_additional_service_easy_post_id" character varying NOT NULL, PRIMARY KEY ("delivery_option_easy_post_id", "carrier_additional_service_easy_post_id"));
-- Create "delivery_option_easy_posts" table
CREATE TABLE "delivery_option_easy_posts" ("id" character varying NOT NULL, "delivery_option_delivery_option_easy_post" character varying NOT NULL, "tenant_id" character varying NOT NULL, PRIMARY KEY ("id"));
-- Create index "delivery_option_easy_posts_delivery_option_delivery_option_easy" to table: "delivery_option_easy_posts"
CREATE UNIQUE INDEX "delivery_option_easy_posts_delivery_option_delivery_option_easy" ON "delivery_option_easy_posts" ("delivery_option_delivery_option_easy_post");
-- Create index "deliveryoptioneasypost_tenant_id" to table: "delivery_option_easy_posts"
CREATE INDEX "deliveryoptioneasypost_tenant_id" ON "delivery_option_easy_posts" ("tenant_id");
-- Create "delivery_option_gl_ss" table
CREATE TABLE "delivery_option_gl_ss" ("id" character varying NOT NULL, "delivery_option_delivery_option_gls" character varying NOT NULL, "tenant_id" character varying NOT NULL, PRIMARY KEY ("id"));
-- Create index "delivery_option_gl_ss_delivery_option_delivery_option_gls_key" to table: "delivery_option_gl_ss"
CREATE UNIQUE INDEX "delivery_option_gl_ss_delivery_option_delivery_option_gls_key" ON "delivery_option_gl_ss" ("delivery_option_delivery_option_gls");
-- Create index "deliveryoptiongls_tenant_id" to table: "delivery_option_gl_ss"
CREATE INDEX "deliveryoptiongls_tenant_id" ON "delivery_option_gl_ss" ("tenant_id");
-- Create "delivery_option_gls_carrier_additional_service_gls" table
CREATE TABLE "delivery_option_gls_carrier_additional_service_gls" ("delivery_option_gls_id" character varying NOT NULL, "carrier_additional_service_gls_id" character varying NOT NULL, PRIMARY KEY ("delivery_option_gls_id", "carrier_additional_service_gls_id"));
-- Create "delivery_option_post_nord_carrier_add_serv_post_nord" table
CREATE TABLE "delivery_option_post_nord_carrier_add_serv_post_nord" ("delivery_option_post_nord_id" character varying NOT NULL, "carrier_additional_service_post_nord_id" character varying NOT NULL, PRIMARY KEY ("delivery_option_post_nord_id", "carrier_additional_service_post_nord_id"));
-- Create "delivery_option_post_nords" table
CREATE TABLE "delivery_option_post_nords" ("id" character varying NOT NULL, "format_zpl" boolean NOT NULL DEFAULT true, "delivery_option_delivery_option_post_nord" character varying NOT NULL, "tenant_id" character varying NOT NULL, PRIMARY KEY ("id"));
-- Create index "delivery_option_post_nords_delivery_option_delivery_option_post" to table: "delivery_option_post_nords"
CREATE UNIQUE INDEX "delivery_option_post_nords_delivery_option_delivery_option_post" ON "delivery_option_post_nords" ("delivery_option_delivery_option_post_nord");
-- Create index "deliveryoptionpostnord_tenant_id" to table: "delivery_option_post_nords"
CREATE INDEX "deliveryoptionpostnord_tenant_id" ON "delivery_option_post_nords" ("tenant_id");
-- Create "delivery_option_usp_ss" table
CREATE TABLE "delivery_option_usp_ss" ("id" character varying NOT NULL, "format_zpl" boolean NOT NULL DEFAULT true, "delivery_option_delivery_option_usps" character varying NOT NULL, "tenant_id" character varying NOT NULL, PRIMARY KEY ("id"));
-- Create index "delivery_option_usp_ss_delivery_option_delivery_option_usps_key" to table: "delivery_option_usp_ss"
CREATE UNIQUE INDEX "delivery_option_usp_ss_delivery_option_delivery_option_usps_key" ON "delivery_option_usp_ss" ("delivery_option_delivery_option_usps");
-- Create index "deliveryoptionusps_tenant_id" to table: "delivery_option_usp_ss"
CREATE INDEX "deliveryoptionusps_tenant_id" ON "delivery_option_usp_ss" ("tenant_id");
-- Create "delivery_option_usps_carrier_additional_service_usps" table
CREATE TABLE "delivery_option_usps_carrier_additional_service_usps" ("delivery_option_usps_id" character varying NOT NULL, "carrier_additional_service_usps_id" character varying NOT NULL, PRIMARY KEY ("delivery_option_usps_id", "carrier_additional_service_usps_id"));
-- Create "delivery_options" table
CREATE TABLE "delivery_options" ("id" character varying NOT NULL, "archived_at" timestamptz NULL, "name" character varying NOT NULL, "sort_order" bigint NOT NULL, "click_option_display_count" bigint NULL DEFAULT 3, "description" character varying NULL, "click_collect" boolean NULL DEFAULT false, "override_sender_address" boolean NULL DEFAULT false, "override_return_address" boolean NULL DEFAULT false, "hide_delivery_option" boolean NULL DEFAULT false, "delivery_estimate_from" bigint NULL, "delivery_estimate_to" bigint NULL, "webshipper_integration" boolean NOT NULL DEFAULT false, "webshipper_id" bigint NULL DEFAULT 1, "shipmondo_integration" boolean NOT NULL DEFAULT false, "shipmondo_delivery_option" character varying NULL, "customs_enabled" boolean NOT NULL DEFAULT false, "customs_signer" character varying NULL, "connection_delivery_option" character varying NOT NULL, "connection_default_delivery_option" character varying NULL, "tenant_id" character varying NOT NULL, "delivery_option_carrier" character varying NOT NULL, "delivery_option_carrier_service" character varying NOT NULL, "delivery_option_email_click_collect_at_store" character varying NULL, "delivery_option_default_packaging" character varying NULL, PRIMARY KEY ("id"));
-- Create index "delivery_options_connection_default_delivery_option_key" to table: "delivery_options"
CREATE UNIQUE INDEX "delivery_options_connection_default_delivery_option_key" ON "delivery_options" ("connection_default_delivery_option");
-- Create index "deliveryoption_tenant_id" to table: "delivery_options"
CREATE INDEX "deliveryoption_tenant_id" ON "delivery_options" ("tenant_id");
-- Create "delivery_rule_constraint_groups" table
CREATE TABLE "delivery_rule_constraint_groups" ("id" character varying NOT NULL, "constraint_logic" character varying NOT NULL DEFAULT 'and', "delivery_rule_delivery_rule_constraint_group" character varying NOT NULL, "tenant_id" character varying NOT NULL, PRIMARY KEY ("id"));
-- Create index "deliveryruleconstraintgroup_tenant_id" to table: "delivery_rule_constraint_groups"
CREATE INDEX "deliveryruleconstraintgroup_tenant_id" ON "delivery_rule_constraint_groups" ("tenant_id");
-- Create "delivery_rule_constraints" table
CREATE TABLE "delivery_rule_constraints" ("id" character varying NOT NULL, "property_type" character varying NOT NULL, "comparison" character varying NOT NULL, "selected_value" jsonb NOT NULL, "tenant_id" character varying NOT NULL, "delivery_rule_constraint_group_delivery_rule_constraints" character varying NOT NULL, PRIMARY KEY ("id"));
-- Create index "deliveryruleconstraint_tenant_id" to table: "delivery_rule_constraints"
CREATE INDEX "deliveryruleconstraint_tenant_id" ON "delivery_rule_constraints" ("tenant_id");
-- Create "delivery_rules" table
CREATE TABLE "delivery_rules" ("id" character varying NOT NULL, "name" character varying NOT NULL, "price" double precision NOT NULL DEFAULT 20, "delivery_option_delivery_rule" character varying NULL, "tenant_id" character varying NOT NULL, "delivery_rule_currency" character varying NULL, PRIMARY KEY ("id"));
-- Create index "deliveryrule_tenant_id" to table: "delivery_rules"
CREATE INDEX "deliveryrule_tenant_id" ON "delivery_rules" ("tenant_id");
-- Create "document_files" table
CREATE TABLE "document_files" ("id" character varying NOT NULL, "created_at" timestamptz NOT NULL, "doc_type" character varying NOT NULL, "data_pdf_base64" character varying NULL, "data_zpl_base64" character varying NULL, "colli_document_file" character varying NULL, "tenant_id" character varying NOT NULL, PRIMARY KEY ("id"));
-- Create index "documentfile_tenant_id" to table: "document_files"
CREATE INDEX "documentfile_tenant_id" ON "document_files" ("tenant_id");
-- Create "documents" table
CREATE TABLE "documents" ("id" character varying NOT NULL, "name" character varying NOT NULL, "html_template" character varying NULL, "html_header" character varying NULL, "html_footer" character varying NULL, "last_base64_pdf" character varying NULL, "merge_type" character varying NOT NULL DEFAULT 'Orders', "paper_size" character varying NOT NULL DEFAULT 'A4', "start_at" timestamptz NOT NULL, "end_at" timestamptz NOT NULL, "created_at" timestamptz NOT NULL, "tenant_id" character varying NOT NULL, "document_carrier_brand" character varying NULL, PRIMARY KEY ("id"));
-- Create index "document_tenant_id" to table: "documents"
CREATE INDEX "document_tenant_id" ON "documents" ("tenant_id");
-- Create "email_templates" table
CREATE TABLE "email_templates" ("id" character varying NOT NULL, "name" character varying NOT NULL, "subject" character varying NOT NULL DEFAULT '', "html_template" character varying NOT NULL DEFAULT '', "merge_type" character varying NOT NULL DEFAULT 'return_colli_label', "created_at" timestamptz NULL, "updated_at" timestamptz NOT NULL, "tenant_id" character varying NOT NULL, PRIMARY KEY ("id"));
-- Create index "emailtemplate_tenant_id" to table: "email_templates"
CREATE INDEX "emailtemplate_tenant_id" ON "email_templates" ("tenant_id");
-- Create "hypothesis_test_delivery_option_delivery_option_group_one" table
CREATE TABLE "hypothesis_test_delivery_option_delivery_option_group_one" ("hypothesis_test_delivery_option_id" character varying NOT NULL, "delivery_option_id" character varying NOT NULL, PRIMARY KEY ("hypothesis_test_delivery_option_id", "delivery_option_id"));
-- Create "hypothesis_test_delivery_option_delivery_option_group_two" table
CREATE TABLE "hypothesis_test_delivery_option_delivery_option_group_two" ("hypothesis_test_delivery_option_id" character varying NOT NULL, "delivery_option_id" character varying NOT NULL, PRIMARY KEY ("hypothesis_test_delivery_option_id", "delivery_option_id"));
-- Create "hypothesis_test_delivery_option_lookups" table
CREATE TABLE "hypothesis_test_delivery_option_lookups" ("id" character varying NOT NULL, "tenant_id" character varying NOT NULL, "hypothesis_test_delivery_option_lookup_delivery_option" character varying NOT NULL, "hypothesis_test_delivery_option_request_hypothesis_test_deliver" character varying NOT NULL, PRIMARY KEY ("id"));
-- Create index "hypothesistestdeliveryoptionlookup_hypothesis_test_delivery_opt" to table: "hypothesis_test_delivery_option_lookups"
CREATE UNIQUE INDEX "hypothesistestdeliveryoptionlookup_hypothesis_test_delivery_opt" ON "hypothesis_test_delivery_option_lookups" ("hypothesis_test_delivery_option_lookup_delivery_option", "hypothesis_test_delivery_option_request_hypothesis_test_deliver");
-- Create index "hypothesistestdeliveryoptionlookup_tenant_id" to table: "hypothesis_test_delivery_option_lookups"
CREATE INDEX "hypothesistestdeliveryoptionlookup_tenant_id" ON "hypothesis_test_delivery_option_lookups" ("tenant_id");
-- Create "hypothesis_test_delivery_option_requests" table
CREATE TABLE "hypothesis_test_delivery_option_requests" ("id" character varying NOT NULL, "order_hash" character varying NOT NULL, "shipping_address_hash" character varying NOT NULL, "is_control_group" boolean NOT NULL, "request_count" bigint NOT NULL, "created_at" timestamptz NOT NULL, "last_requested_at" timestamptz NOT NULL, "tenant_id" character varying NOT NULL, "hypothesis_test_delivery_option_request_hypothesis_test_deliver" character varying NOT NULL, PRIMARY KEY ("id"));
-- Create index "hypothesistestdeliveryoptionrequest_tenant_id" to table: "hypothesis_test_delivery_option_requests"
CREATE INDEX "hypothesistestdeliveryoptionrequest_tenant_id" ON "hypothesis_test_delivery_option_requests" ("tenant_id");
-- Create "hypothesis_test_delivery_options" table
CREATE TABLE "hypothesis_test_delivery_options" ("id" character varying NOT NULL, "randomize_within_group_sort" boolean NOT NULL DEFAULT false, "by_interval_rotation" boolean NOT NULL DEFAULT false, "rotation_interval_hours" bigint NOT NULL DEFAULT 6, "by_order" boolean NOT NULL DEFAULT false, "hypothesis_test_hypothesis_test_delivery_option" character varying NOT NULL, "tenant_id" character varying NOT NULL, PRIMARY KEY ("id"));
-- Create index "hypothesis_test_delivery_options_hypothesis_test_hypothesis_tes" to table: "hypothesis_test_delivery_options"
CREATE UNIQUE INDEX "hypothesis_test_delivery_options_hypothesis_test_hypothesis_tes" ON "hypothesis_test_delivery_options" ("hypothesis_test_hypothesis_test_delivery_option");
-- Create index "hypothesistestdeliveryoption_tenant_id" to table: "hypothesis_test_delivery_options"
CREATE INDEX "hypothesistestdeliveryoption_tenant_id" ON "hypothesis_test_delivery_options" ("tenant_id");
-- Create "hypothesis_tests" table
CREATE TABLE "hypothesis_tests" ("id" character varying NOT NULL, "name" character varying NOT NULL, "active" boolean NOT NULL DEFAULT false, "tenant_id" character varying NOT NULL, "hypothesis_test_connection" character varying NOT NULL, PRIMARY KEY ("id"));
-- Create index "hypothesistest_tenant_id" to table: "hypothesis_tests"
CREATE INDEX "hypothesistest_tenant_id" ON "hypothesis_tests" ("tenant_id");
-- Create "inventory_items" table
CREATE TABLE "inventory_items" ("id" character varying NOT NULL, "external_id" character varying NULL, "code" character varying NULL, "sku" character varying NULL, "tenant_id" character varying NOT NULL, "inventory_item_country_of_origin" character varying NULL, "product_variant_inventory_item" character varying NOT NULL, PRIMARY KEY ("id"));
-- Create index "inventory_items_product_variant_inventory_item_key" to table: "inventory_items"
CREATE UNIQUE INDEX "inventory_items_product_variant_inventory_item_key" ON "inventory_items" ("product_variant_inventory_item");
-- Create index "inventoryitem_tenant_id" to table: "inventory_items"
CREATE INDEX "inventoryitem_tenant_id" ON "inventory_items" ("tenant_id");
-- Create "languages" table
CREATE TABLE "languages" ("id" character varying NOT NULL, "label" character varying NOT NULL, "internal_id" character varying NOT NULL DEFAULT 'EN', PRIMARY KEY ("id"));
-- Create index "languages_label_key" to table: "languages"
CREATE UNIQUE INDEX "languages_label_key" ON "languages" ("label");
-- Create "location_location_tags" table
CREATE TABLE "location_location_tags" ("location_id" character varying NOT NULL, "location_tag_id" character varying NOT NULL, PRIMARY KEY ("location_id", "location_tag_id"));
-- Create "location_tags" table
CREATE TABLE "location_tags" ("id" character varying NOT NULL, "label" character varying NOT NULL, "internal_id" character varying NOT NULL, PRIMARY KEY ("id"));
-- Create index "location_tags_internal_id_key" to table: "location_tags"
CREATE UNIQUE INDEX "location_tags_internal_id_key" ON "location_tags" ("internal_id");
-- Create index "location_tags_label_key" to table: "location_tags"
CREATE UNIQUE INDEX "location_tags_label_key" ON "location_tags" ("label");
-- Create "locations" table
CREATE TABLE "locations" ("id" character varying NOT NULL, "name" character varying NOT NULL, "tenant_id" character varying NOT NULL, "location_address" character varying NOT NULL, PRIMARY KEY ("id"));
-- Create index "location_name_tenant_id" to table: "locations"
CREATE UNIQUE INDEX "location_name_tenant_id" ON "locations" ("name", "tenant_id");
-- Create index "location_tenant_id" to table: "locations"
CREATE INDEX "location_tenant_id" ON "locations" ("tenant_id");
-- Create "notifications" table
CREATE TABLE "notifications" ("id" character varying NOT NULL, "name" character varying NOT NULL, "active" boolean NOT NULL DEFAULT true, "tenant_id" character varying NOT NULL, "notification_connection" character varying NOT NULL, "notification_email_template" character varying NOT NULL, PRIMARY KEY ("id"));
-- Create index "notification_tenant_id" to table: "notifications"
CREATE INDEX "notification_tenant_id" ON "notifications" ("tenant_id");
-- Create "order_histories" table
CREATE TABLE "order_histories" ("id" character varying NOT NULL, "description" character varying NOT NULL, "type" character varying NOT NULL, "change_history_order_history" character varying NOT NULL, "order_order_history" character varying NOT NULL, "tenant_id" character varying NOT NULL, PRIMARY KEY ("id"));
-- Create index "orderhistory_tenant_id" to table: "order_histories"
CREATE INDEX "orderhistory_tenant_id" ON "order_histories" ("tenant_id");
-- Create "order_lines" table
CREATE TABLE "order_lines" ("id" character varying NOT NULL, "unit_price" double precision NOT NULL, "discount_allocation_amount" double precision NOT NULL, "external_id" character varying NULL, "units" bigint NOT NULL, "created_at" timestamptz NULL, "updated_at" timestamptz NOT NULL, "colli_id" character varying NOT NULL, "tenant_id" character varying NOT NULL, "product_variant_id" character varying NOT NULL, "order_line_currency" character varying NOT NULL, PRIMARY KEY ("id"));
-- Create index "orderline_tenant_id" to table: "order_lines"
CREATE INDEX "orderline_tenant_id" ON "order_lines" ("tenant_id");
-- Create "order_senders" table
CREATE TABLE "order_senders" ("id" character varying NOT NULL, "uniqueness_id" character varying NULL, "first_name" character varying NOT NULL, "last_name" character varying NOT NULL, "email" character varying NOT NULL, "phone_number" character varying NOT NULL, "vat_number" character varying NOT NULL, "tenant_id" character varying NOT NULL, PRIMARY KEY ("id"));
-- Create index "order_senders_uniqueness_id_key" to table: "order_senders"
CREATE UNIQUE INDEX "order_senders_uniqueness_id_key" ON "order_senders" ("uniqueness_id");
-- Create index "ordersender_tenant_id" to table: "order_senders"
CREATE INDEX "ordersender_tenant_id" ON "order_senders" ("tenant_id");
-- Create "orders" table
CREATE TABLE "orders" ("id" character varying NOT NULL, "order_public_id" character varying NOT NULL, "external_id" character varying NULL, "comment_internal" character varying NULL, "comment_external" character varying NULL, "created_at" timestamptz NOT NULL, "email_sync_confirmation_at" timestamptz NULL, "status" character varying NOT NULL, "connection_orders" character varying NOT NULL, "consolidation_orders" character varying NULL, "hypothesis_test_delivery_option_request_order" character varying NULL, "tenant_id" character varying NOT NULL, "pallet_orders" character varying NULL, PRIMARY KEY ("id"));
-- Create index "order_created_at" to table: "orders"
CREATE INDEX "order_created_at" ON "orders" ("created_at");
-- Create index "order_external_id_tenant_id" to table: "orders"
CREATE UNIQUE INDEX "order_external_id_tenant_id" ON "orders" ("external_id", "tenant_id");
-- Create index "order_order_public_id_tenant_id" to table: "orders"
CREATE UNIQUE INDEX "order_order_public_id_tenant_id" ON "orders" ("order_public_id", "tenant_id");
-- Create index "order_tenant_id" to table: "orders"
CREATE INDEX "order_tenant_id" ON "orders" ("tenant_id");
-- Create index "orders_hypothesis_test_delivery_option_request_order_key" to table: "orders"
CREATE UNIQUE INDEX "orders_hypothesis_test_delivery_option_request_order_key" ON "orders" ("hypothesis_test_delivery_option_request_order");
-- Create "otk_requests" table
CREATE TABLE "otk_requests" ("id" character varying NOT NULL, "otk" character varying NOT NULL, "tenant_id" character varying NOT NULL, "user_otk_requests" character varying NULL, PRIMARY KEY ("id"));
-- Create index "otkrequests_tenant_id" to table: "otk_requests"
CREATE INDEX "otkrequests_tenant_id" ON "otk_requests" ("tenant_id");
-- Create "packaging_dfs" table
CREATE TABLE "packaging_dfs" ("id" character varying NOT NULL, "api_type" character varying NOT NULL, "max_weight" double precision NULL, "min_weight" double precision NULL, "stackable" boolean NOT NULL DEFAULT false, "packaging_packaging_df" character varying NOT NULL, "tenant_id" character varying NOT NULL, PRIMARY KEY ("id"));
-- Create index "packaging_dfs_packaging_packaging_df_key" to table: "packaging_dfs"
CREATE UNIQUE INDEX "packaging_dfs_packaging_packaging_df_key" ON "packaging_dfs" ("packaging_packaging_df");
-- Create index "packagingdf_tenant_id" to table: "packaging_dfs"
CREATE INDEX "packagingdf_tenant_id" ON "packaging_dfs" ("tenant_id");
-- Create "packaging_usp_ss" table
CREATE TABLE "packaging_usp_ss" ("id" character varying NOT NULL, "packaging_packaging_usps" character varying NOT NULL, "tenant_id" character varying NOT NULL, "packaging_usps_packaging_usps_rate_indicator" character varying NOT NULL, "packaging_usps_packaging_usps_processing_category" character varying NOT NULL, PRIMARY KEY ("id"));
-- Create index "packaging_usp_ss_packaging_packaging_usps_key" to table: "packaging_usp_ss"
CREATE UNIQUE INDEX "packaging_usp_ss_packaging_packaging_usps_key" ON "packaging_usp_ss" ("packaging_packaging_usps");
-- Create index "packagingusps_tenant_id" to table: "packaging_usp_ss"
CREATE INDEX "packagingusps_tenant_id" ON "packaging_usp_ss" ("tenant_id");
-- Create "packaging_usps_processing_categories" table
CREATE TABLE "packaging_usps_processing_categories" ("id" character varying NOT NULL, "name" character varying NOT NULL, "processing_category" character varying NOT NULL, PRIMARY KEY ("id"));
-- Create "packaging_usps_rate_indicators" table
CREATE TABLE "packaging_usps_rate_indicators" ("id" character varying NOT NULL, "code" character varying NOT NULL, "name" character varying NOT NULL, PRIMARY KEY ("id"));
-- Create "packagings" table
CREATE TABLE "packagings" ("id" character varying NOT NULL, "archived_at" timestamptz NULL, "name" character varying NOT NULL, "height_cm" bigint NOT NULL, "width_cm" bigint NOT NULL, "length_cm" bigint NOT NULL, "tenant_id" character varying NOT NULL, "packaging_carrier_brand" character varying NULL, PRIMARY KEY ("id"));
-- Create index "packaging_tenant_id" to table: "packagings"
CREATE INDEX "packaging_tenant_id" ON "packagings" ("tenant_id");
-- Create "pallet_cancelled_shipment_pallet" table
CREATE TABLE "pallet_cancelled_shipment_pallet" ("pallet_id" character varying NOT NULL, "shipment_pallet_id" character varying NOT NULL, PRIMARY KEY ("pallet_id", "shipment_pallet_id"));
-- Create "pallets" table
CREATE TABLE "pallets" ("id" character varying NOT NULL, "public_id" character varying NOT NULL, "description" character varying NOT NULL, "consolidation_pallets" character varying NOT NULL, "tenant_id" character varying NOT NULL, "pallet_packaging" character varying NULL, PRIMARY KEY ("id"));
-- Create index "pallet_tenant_id" to table: "pallets"
CREATE INDEX "pallet_tenant_id" ON "pallets" ("tenant_id");
-- Create "parcel_shop_brings" table
CREATE TABLE "parcel_shop_brings" ("id" character varying NOT NULL, "point_type" character varying NOT NULL, "bring_id" character varying NOT NULL, "parcel_shop_parcel_shop_bring" character varying NOT NULL, PRIMARY KEY ("id"));
-- Create index "parcel_shop_brings_bring_id_key" to table: "parcel_shop_brings"
CREATE UNIQUE INDEX "parcel_shop_brings_bring_id_key" ON "parcel_shop_brings" ("bring_id");
-- Create index "parcel_shop_brings_parcel_shop_parcel_shop_bring_key" to table: "parcel_shop_brings"
CREATE UNIQUE INDEX "parcel_shop_brings_parcel_shop_parcel_shop_bring_key" ON "parcel_shop_brings" ("parcel_shop_parcel_shop_bring");
-- Create "parcel_shop_da_os" table
CREATE TABLE "parcel_shop_da_os" ("id" character varying NOT NULL, "shop_id" character varying NOT NULL, "parcel_shop_parcel_shop_dao" character varying NOT NULL, PRIMARY KEY ("id"));
-- Create index "parcel_shop_da_os_parcel_shop_parcel_shop_dao_key" to table: "parcel_shop_da_os"
CREATE UNIQUE INDEX "parcel_shop_da_os_parcel_shop_parcel_shop_dao_key" ON "parcel_shop_da_os" ("parcel_shop_parcel_shop_dao");
-- Create "parcel_shop_gl_ss" table
CREATE TABLE "parcel_shop_gl_ss" ("id" character varying NOT NULL, "gls_parcel_shop_id" character varying NOT NULL, "partner_id" character varying NULL, "type" character varying NULL, "parcel_shop_parcel_shop_gls" character varying NOT NULL, PRIMARY KEY ("id"));
-- Create index "parcel_shop_gl_ss_gls_parcel_shop_id_key" to table: "parcel_shop_gl_ss"
CREATE UNIQUE INDEX "parcel_shop_gl_ss_gls_parcel_shop_id_key" ON "parcel_shop_gl_ss" ("gls_parcel_shop_id");
-- Create index "parcel_shop_gl_ss_parcel_shop_parcel_shop_gls_key" to table: "parcel_shop_gl_ss"
CREATE UNIQUE INDEX "parcel_shop_gl_ss_parcel_shop_parcel_shop_gls_key" ON "parcel_shop_gl_ss" ("parcel_shop_parcel_shop_gls");
-- Create "parcel_shop_post_nords" table
CREATE TABLE "parcel_shop_post_nords" ("id" character varying NOT NULL, "service_point_id" character varying NOT NULL, "pudoid" character varying NOT NULL, "type_id" character varying NOT NULL, "parcel_shop_parcel_shop_post_nord" character varying NOT NULL, PRIMARY KEY ("id"));
-- Create index "parcel_shop_post_nords_parcel_shop_parcel_shop_post_nord_key" to table: "parcel_shop_post_nords"
CREATE UNIQUE INDEX "parcel_shop_post_nords_parcel_shop_parcel_shop_post_nord_key" ON "parcel_shop_post_nords" ("parcel_shop_parcel_shop_post_nord");
-- Create index "parcel_shop_post_nords_pudoid_key" to table: "parcel_shop_post_nords"
CREATE UNIQUE INDEX "parcel_shop_post_nords_pudoid_key" ON "parcel_shop_post_nords" ("pudoid");
-- Create "parcel_shops" table
CREATE TABLE "parcel_shops" ("id" character varying NOT NULL, "name" character varying NOT NULL, "last_updated" timestamptz NOT NULL, "parcel_shop_carrier_brand" character varying NOT NULL, PRIMARY KEY ("id"));
-- Create "plan_histories" table
CREATE TABLE "plan_histories" ("id" character varying NOT NULL, "created_at" timestamptz NOT NULL, "change_history_plan_history" character varying NOT NULL, "plan_plan_history_plan" character varying NOT NULL, "tenant_id" character varying NOT NULL, "user_plan_history_user" character varying NOT NULL, PRIMARY KEY ("id"));
-- Create index "planhistory_tenant_id" to table: "plan_histories"
CREATE INDEX "planhistory_tenant_id" ON "plan_histories" ("tenant_id");
-- Create "plans" table
CREATE TABLE "plans" ("id" character varying NOT NULL, "label" character varying NOT NULL, "rank" bigint NOT NULL, "price_dkk" bigint NOT NULL, "created_at" timestamptz NOT NULL, PRIMARY KEY ("id"));
-- Create index "plans_label_key" to table: "plans"
CREATE UNIQUE INDEX "plans_label_key" ON "plans" ("label");
-- Create "print_jobs" table
CREATE TABLE "print_jobs" ("id" character varying NOT NULL, "status" character varying NOT NULL, "file_extension" character varying NOT NULL, "document_type" character varying NOT NULL, "base64_print_data" character varying NOT NULL, "created_at" timestamptz NOT NULL, "tenant_id" character varying NOT NULL, "print_job_printer" character varying NOT NULL, "print_job_colli" character varying NULL, "print_job_shipment_parcel" character varying NULL, PRIMARY KEY ("id"));
-- Create index "printjob_tenant_id" to table: "print_jobs"
CREATE INDEX "printjob_tenant_id" ON "print_jobs" ("tenant_id");
-- Create "printers" table
CREATE TABLE "printers" ("id" character varying NOT NULL, "device_id" character varying NOT NULL, "name" character varying NOT NULL, "label_zpl" boolean NOT NULL DEFAULT false, "label_pdf" boolean NOT NULL DEFAULT false, "document" boolean NOT NULL DEFAULT false, "rotate_180" boolean NOT NULL DEFAULT false, "print_size" character varying NOT NULL DEFAULT 'A4', "created_at" timestamptz NOT NULL, "last_ping" timestamptz NOT NULL, "tenant_id" character varying NOT NULL, "workstation_printer" character varying NOT NULL, PRIMARY KEY ("id"));
-- Create index "printer_tenant_id" to table: "printers"
CREATE INDEX "printer_tenant_id" ON "printers" ("tenant_id");
-- Create index "printers_device_id_key" to table: "printers"
CREATE UNIQUE INDEX "printers_device_id_key" ON "printers" ("device_id");
-- Create "product_image_product_variant" table
CREATE TABLE "product_image_product_variant" ("product_image_id" character varying NOT NULL, "product_variant_id" character varying NOT NULL, PRIMARY KEY ("product_image_id", "product_variant_id"));
-- Create "product_images" table
CREATE TABLE "product_images" ("id" character varying NOT NULL, "external_id" character varying NULL, "url" character varying NOT NULL, "tenant_id" character varying NOT NULL, "product_image_product" character varying NOT NULL, PRIMARY KEY ("id"));
-- Create index "productimage_external_id_tenant_id" to table: "product_images"
CREATE UNIQUE INDEX "productimage_external_id_tenant_id" ON "product_images" ("external_id", "tenant_id");
-- Create index "productimage_tenant_id" to table: "product_images"
CREATE INDEX "productimage_tenant_id" ON "product_images" ("tenant_id");
-- Create "product_tag_products" table
CREATE TABLE "product_tag_products" ("product_tag_id" character varying NOT NULL, "product_id" character varying NOT NULL, PRIMARY KEY ("product_tag_id", "product_id"));
-- Create "product_tags" table
CREATE TABLE "product_tags" ("id" character varying NOT NULL, "name" character varying NOT NULL, "created_at" timestamptz NULL, "tenant_id" character varying NOT NULL, PRIMARY KEY ("id"));
-- Create index "producttag_name_tenant_id" to table: "product_tags"
CREATE UNIQUE INDEX "producttag_name_tenant_id" ON "product_tags" ("name", "tenant_id");
-- Create index "producttag_tenant_id" to table: "product_tags"
CREATE INDEX "producttag_tenant_id" ON "product_tags" ("tenant_id");
-- Create "product_variants" table
CREATE TABLE "product_variants" ("id" character varying NOT NULL, "archived" boolean NOT NULL DEFAULT false, "external_id" character varying NULL, "description" character varying NULL, "ean_number" character varying NULL, "weight_g" bigint NULL DEFAULT 0, "dimension_length" bigint NULL, "dimension_width" bigint NULL, "dimension_height" bigint NULL, "created_at" timestamptz NULL, "updated_at" timestamptz NOT NULL, "product_product_variant" character varying NOT NULL, "tenant_id" character varying NOT NULL, PRIMARY KEY ("id"));
-- Create index "productvariant_external_id_tenant_id" to table: "product_variants"
CREATE UNIQUE INDEX "productvariant_external_id_tenant_id" ON "product_variants" ("external_id", "tenant_id");
-- Create index "productvariant_tenant_id" to table: "product_variants"
CREATE INDEX "productvariant_tenant_id" ON "product_variants" ("tenant_id");
-- Create "products" table
CREATE TABLE "products" ("id" character varying NOT NULL, "external_id" character varying NULL, "title" character varying NOT NULL, "body_html" character varying NULL, "status" character varying NOT NULL DEFAULT 'active', "created_at" timestamptz NULL, "updated_at" timestamptz NOT NULL, "tenant_id" character varying NOT NULL, PRIMARY KEY ("id"));
-- Create index "product_external_id_tenant_id" to table: "products"
CREATE UNIQUE INDEX "product_external_id_tenant_id" ON "products" ("external_id", "tenant_id");
-- Create index "product_tenant_id" to table: "products"
CREATE INDEX "product_tenant_id" ON "products" ("tenant_id");
-- Create "return_colli_histories" table
CREATE TABLE "return_colli_histories" ("id" character varying NOT NULL, "description" character varying NOT NULL, "type" character varying NOT NULL, "change_history_return_colli_history" character varying NOT NULL, "return_colli_return_colli_history" character varying NOT NULL, "tenant_id" character varying NOT NULL, PRIMARY KEY ("id"));
-- Create index "returncollihistory_tenant_id" to table: "return_colli_histories"
CREATE INDEX "returncollihistory_tenant_id" ON "return_colli_histories" ("tenant_id");
-- Create "return_collis" table
CREATE TABLE "return_collis" ("id" character varying NOT NULL, "expected_at" timestamptz NULL, "label_pdf" character varying NULL, "label_png" character varying NULL, "qr_code_png" character varying NULL, "comment" character varying NULL, "created_at" timestamptz NOT NULL, "status" character varying NOT NULL DEFAULT 'Opened', "email_received" timestamptz NULL, "email_accepted" timestamptz NULL, "email_confirmation_label" timestamptz NULL, "email_confirmation_qr_code" timestamptz NULL, "order_return_colli" character varying NOT NULL, "tenant_id" character varying NOT NULL, "return_colli_recipient" character varying NOT NULL, "return_colli_sender" character varying NOT NULL, "return_colli_delivery_option" character varying NULL, "return_colli_return_portal" character varying NOT NULL, "return_colli_packaging" character varying NULL, PRIMARY KEY ("id"));
-- Create index "returncolli_tenant_id" to table: "return_collis"
CREATE INDEX "returncolli_tenant_id" ON "return_collis" ("tenant_id");
-- Create "return_order_lines" table
CREATE TABLE "return_order_lines" ("id" character varying NOT NULL, "units" bigint NOT NULL, "return_colli_return_order_line" character varying NOT NULL, "tenant_id" character varying NOT NULL, "return_order_line_order_line" character varying NOT NULL, "return_order_line_return_portal_claim" character varying NOT NULL, PRIMARY KEY ("id"));
-- Create index "returnorderline_tenant_id" to table: "return_order_lines"
CREATE INDEX "returnorderline_tenant_id" ON "return_order_lines" ("tenant_id");
-- Create "return_portal_claims" table
CREATE TABLE "return_portal_claims" ("id" character varying NOT NULL, "name" character varying NOT NULL, "description" character varying NOT NULL, "restockable" boolean NOT NULL, "archived" boolean NOT NULL, "return_portal_return_portal_claim" character varying NOT NULL, "tenant_id" character varying NOT NULL, "return_portal_claim_return_location" character varying NULL, PRIMARY KEY ("id"));
-- Create index "returnportalclaim_tenant_id" to table: "return_portal_claims"
CREATE INDEX "returnportalclaim_tenant_id" ON "return_portal_claims" ("tenant_id");
-- Create "return_portal_delivery_options" table
CREATE TABLE "return_portal_delivery_options" ("return_portal_id" character varying NOT NULL, "delivery_option_id" character varying NOT NULL, PRIMARY KEY ("return_portal_id", "delivery_option_id"));
-- Create "return_portal_return_location" table
CREATE TABLE "return_portal_return_location" ("return_portal_id" character varying NOT NULL, "location_id" character varying NOT NULL, PRIMARY KEY ("return_portal_id", "location_id"));
-- Create "return_portals" table
CREATE TABLE "return_portals" ("id" character varying NOT NULL, "name" character varying NOT NULL, "return_open_hours" bigint NOT NULL DEFAULT 720, "automatically_accept" boolean NOT NULL DEFAULT false, "tenant_id" character varying NOT NULL, "return_portal_email_confirmation_label" character varying NULL, "return_portal_email_confirmation_qr_code" character varying NULL, "return_portal_email_received" character varying NULL, "return_portal_email_accepted" character varying NULL, PRIMARY KEY ("id"));
-- Create index "returnportal_tenant_id" to table: "return_portals"
CREATE INDEX "returnportal_tenant_id" ON "return_portals" ("tenant_id");
-- Create "seat_group_access_rights" table
CREATE TABLE "seat_group_access_rights" ("id" character varying NOT NULL, "level" character varying NOT NULL DEFAULT 'none', "tenant_id" character varying NOT NULL, "access_right_id" character varying NOT NULL, "seat_group_id" character varying NOT NULL, PRIMARY KEY ("id"));
-- Create index "seatgroupaccessright_seat_group_id_access_right_id" to table: "seat_group_access_rights"
CREATE UNIQUE INDEX "seatgroupaccessright_seat_group_id_access_right_id" ON "seat_group_access_rights" ("seat_group_id", "access_right_id");
-- Create index "seatgroupaccessright_tenant_id" to table: "seat_group_access_rights"
CREATE INDEX "seatgroupaccessright_tenant_id" ON "seat_group_access_rights" ("tenant_id");
-- Create "seat_groups" table
CREATE TABLE "seat_groups" ("id" character varying NOT NULL, "name" character varying NOT NULL, "created_at" timestamptz NULL, "tenant_id" character varying NOT NULL, PRIMARY KEY ("id"));
-- Create index "seatgroup_tenant_id" to table: "seat_groups"
CREATE INDEX "seatgroup_tenant_id" ON "seat_groups" ("tenant_id");
-- Create "shipment_brings" table
CREATE TABLE "shipment_brings" ("id" character varying NOT NULL, "consignment_number" character varying NOT NULL, "shipment_shipment_bring" character varying NOT NULL, "tenant_id" character varying NOT NULL, PRIMARY KEY ("id"));
-- Create index "shipment_brings_shipment_shipment_bring_key" to table: "shipment_brings"
CREATE UNIQUE INDEX "shipment_brings_shipment_shipment_bring_key" ON "shipment_brings" ("shipment_shipment_bring");
-- Create index "shipmentbring_tenant_id" to table: "shipment_brings"
CREATE INDEX "shipmentbring_tenant_id" ON "shipment_brings" ("tenant_id");
-- Create "shipment_da_os" table
CREATE TABLE "shipment_da_os" ("id" character varying NOT NULL, "barcode_id" character varying NOT NULL, "shipment_shipment_dao" character varying NOT NULL, "tenant_id" character varying NOT NULL, PRIMARY KEY ("id"));
-- Create index "shipment_da_os_shipment_shipment_dao_key" to table: "shipment_da_os"
CREATE UNIQUE INDEX "shipment_da_os_shipment_shipment_dao_key" ON "shipment_da_os" ("shipment_shipment_dao");
-- Create index "shipmentdao_tenant_id" to table: "shipment_da_os"
CREATE INDEX "shipmentdao_tenant_id" ON "shipment_da_os" ("tenant_id");
-- Create "shipment_dfs" table
CREATE TABLE "shipment_dfs" ("id" character varying NOT NULL, "shipment_shipment_df" character varying NOT NULL, "tenant_id" character varying NOT NULL, PRIMARY KEY ("id"));
-- Create index "shipment_dfs_shipment_shipment_df_key" to table: "shipment_dfs"
CREATE UNIQUE INDEX "shipment_dfs_shipment_shipment_df_key" ON "shipment_dfs" ("shipment_shipment_df");
-- Create index "shipmentdf_tenant_id" to table: "shipment_dfs"
CREATE INDEX "shipmentdf_tenant_id" ON "shipment_dfs" ("tenant_id");
-- Create "shipment_ds_vs" table
CREATE TABLE "shipment_ds_vs" ("id" character varying NOT NULL, "barcode_id" character varying NOT NULL, "shipment_shipment_dsv" character varying NOT NULL, "tenant_id" character varying NOT NULL, PRIMARY KEY ("id"));
-- Create index "shipment_ds_vs_shipment_shipment_dsv_key" to table: "shipment_ds_vs"
CREATE UNIQUE INDEX "shipment_ds_vs_shipment_shipment_dsv_key" ON "shipment_ds_vs" ("shipment_shipment_dsv");
-- Create index "shipmentdsv_tenant_id" to table: "shipment_ds_vs"
CREATE INDEX "shipmentdsv_tenant_id" ON "shipment_ds_vs" ("tenant_id");
-- Create "shipment_easy_posts" table
CREATE TABLE "shipment_easy_posts" ("id" character varying NOT NULL, "tracking_number" character varying NULL, "ep_shipment_id" character varying NULL, "rate" double precision NULL, "est_delivery_date" timestamptz NULL, "shipment_shipment_easy_post" character varying NOT NULL, "tenant_id" character varying NOT NULL, PRIMARY KEY ("id"));
-- Create index "shipment_easy_posts_shipment_shipment_easy_post_key" to table: "shipment_easy_posts"
CREATE UNIQUE INDEX "shipment_easy_posts_shipment_shipment_easy_post_key" ON "shipment_easy_posts" ("shipment_shipment_easy_post");
-- Create index "shipmenteasypost_tenant_id" to table: "shipment_easy_posts"
CREATE INDEX "shipmenteasypost_tenant_id" ON "shipment_easy_posts" ("tenant_id");
-- Create "shipment_gl_ss" table
CREATE TABLE "shipment_gl_ss" ("id" character varying NOT NULL, "consignment_id" character varying NOT NULL, "shipment_shipment_gls" character varying NOT NULL, "tenant_id" character varying NOT NULL, PRIMARY KEY ("id"));
-- Create index "shipment_gl_ss_shipment_shipment_gls_key" to table: "shipment_gl_ss"
CREATE UNIQUE INDEX "shipment_gl_ss_shipment_shipment_gls_key" ON "shipment_gl_ss" ("shipment_shipment_gls");
-- Create index "shipmentgls_tenant_id" to table: "shipment_gl_ss"
CREATE INDEX "shipmentgls_tenant_id" ON "shipment_gl_ss" ("tenant_id");
-- Create "shipment_histories" table
CREATE TABLE "shipment_histories" ("id" character varying NOT NULL, "type" character varying NOT NULL, "change_history_shipment_history" character varying NOT NULL, "shipment_shipment_history" character varying NOT NULL, "tenant_id" character varying NOT NULL, PRIMARY KEY ("id"));
-- Create index "shipmenthistory_tenant_id" to table: "shipment_histories"
CREATE INDEX "shipmenthistory_tenant_id" ON "shipment_histories" ("tenant_id");
-- Create "shipment_old_consolidation" table
CREATE TABLE "shipment_old_consolidation" ("shipment_id" character varying NOT NULL, "consolidation_id" character varying NOT NULL, PRIMARY KEY ("shipment_id", "consolidation_id"));
-- Create "shipment_pallets" table
CREATE TABLE "shipment_pallets" ("id" character varying NOT NULL, "barcode" character varying NOT NULL, "colli_number" character varying NOT NULL, "carrier_id" character varying NOT NULL, "label_pdf" character varying NULL, "label_zpl" character varying NULL, "status" character varying NOT NULL DEFAULT 'pending', "pallet_shipment_pallet" character varying NULL, "shipment_shipment_pallet" character varying NOT NULL, "tenant_id" character varying NOT NULL, PRIMARY KEY ("id"));
-- Create index "shipment_pallets_pallet_shipment_pallet_key" to table: "shipment_pallets"
CREATE UNIQUE INDEX "shipment_pallets_pallet_shipment_pallet_key" ON "shipment_pallets" ("pallet_shipment_pallet");
-- Create index "shipmentpallet_tenant_id" to table: "shipment_pallets"
CREATE INDEX "shipmentpallet_tenant_id" ON "shipment_pallets" ("tenant_id");
-- Create "shipment_parcels" table
CREATE TABLE "shipment_parcels" ("id" character varying NOT NULL, "label_pdf" character varying NULL, "label_zpl" character varying NULL, "item_id" character varying NULL, "status" character varying NOT NULL DEFAULT 'pending', "cc_pickup_signature_urls" jsonb NULL, "expected_at" timestamptz NULL, "fulfillment_synced_at" timestamptz NULL, "cancel_synced_at" timestamptz NULL, "colli_shipment_parcel" character varying NULL, "shipment_shipment_parcel" character varying NOT NULL, "tenant_id" character varying NOT NULL, "shipment_parcel_packaging" character varying NULL, PRIMARY KEY ("id"));
-- Create index "shipment_parcels_colli_shipment_parcel_key" to table: "shipment_parcels"
CREATE UNIQUE INDEX "shipment_parcels_colli_shipment_parcel_key" ON "shipment_parcels" ("colli_shipment_parcel");
-- Create index "shipmentparcel_tenant_id" to table: "shipment_parcels"
CREATE INDEX "shipmentparcel_tenant_id" ON "shipment_parcels" ("tenant_id");
-- Create "shipment_post_nords" table
CREATE TABLE "shipment_post_nords" ("id" character varying NOT NULL, "booking_id" character varying NOT NULL, "item_id" character varying NOT NULL, "shipment_reference_no" character varying NOT NULL, "shipment_shipment_post_nord" character varying NOT NULL, "tenant_id" character varying NOT NULL, PRIMARY KEY ("id"));
-- Create index "shipment_post_nords_shipment_shipment_post_nord_key" to table: "shipment_post_nords"
CREATE UNIQUE INDEX "shipment_post_nords_shipment_shipment_post_nord_key" ON "shipment_post_nords" ("shipment_shipment_post_nord");
-- Create index "shipmentpostnord_tenant_id" to table: "shipment_post_nords"
CREATE INDEX "shipmentpostnord_tenant_id" ON "shipment_post_nords" ("tenant_id");
-- Create "shipment_usp_ss" table
CREATE TABLE "shipment_usp_ss" ("id" character varying NOT NULL, "tracking_number" character varying NULL, "postage" double precision NULL, "scheduled_delivery_date" timestamptz NULL, "shipment_shipment_usps" character varying NOT NULL, "tenant_id" character varying NOT NULL, PRIMARY KEY ("id"));
-- Create index "shipment_usp_ss_shipment_shipment_usps_key" to table: "shipment_usp_ss"
CREATE UNIQUE INDEX "shipment_usp_ss_shipment_shipment_usps_key" ON "shipment_usp_ss" ("shipment_shipment_usps");
-- Create index "shipmentusps_tenant_id" to table: "shipment_usp_ss"
CREATE INDEX "shipmentusps_tenant_id" ON "shipment_usp_ss" ("tenant_id");
-- Create "shipments" table
CREATE TABLE "shipments" ("id" character varying NOT NULL, "shipment_public_id" character varying NOT NULL, "created_at" timestamptz NOT NULL, "status" character varying NOT NULL, "tenant_id" character varying NOT NULL, "shipment_carrier" character varying NOT NULL, PRIMARY KEY ("id"));
-- Create index "shipment_tenant_id" to table: "shipments"
CREATE INDEX "shipment_tenant_id" ON "shipments" ("tenant_id");
-- Create "signup_options" table
CREATE TABLE "signup_options" ("id" character varying NOT NULL, "better_delivery_options" boolean NOT NULL, "improve_pick_pack" boolean NOT NULL, "shipping_label" boolean NOT NULL, "custom_docs" boolean NOT NULL, "reduced_costs" boolean NOT NULL, "easy_returns" boolean NOT NULL, "click_collect" boolean NOT NULL, "num_shipments" bigint NOT NULL, "user_signup_options" character varying NOT NULL, PRIMARY KEY ("id"));
-- Create index "signup_options_user_signup_options_key" to table: "signup_options"
CREATE UNIQUE INDEX "signup_options_user_signup_options_key" ON "signup_options" ("user_signup_options");
-- Create "system_events" table
CREATE TABLE "system_events" ("id" character varying NOT NULL, "event_type" character varying NOT NULL, "event_type_id" character varying NULL, "status" character varying NOT NULL, "description" character varying NOT NULL, "data" character varying NULL, "updated_at" timestamptz NOT NULL, "created_at" timestamptz NULL, "tenant_id" character varying NOT NULL, PRIMARY KEY ("id"));
-- Create index "systemevents_tenant_id" to table: "system_events"
CREATE INDEX "systemevents_tenant_id" ON "system_events" ("tenant_id");
-- Create "tenant_connect_option_carriers" table
CREATE TABLE "tenant_connect_option_carriers" ("tenant_id" character varying NOT NULL, "connect_option_carrier_id" character varying NOT NULL, PRIMARY KEY ("tenant_id", "connect_option_carrier_id"));
-- Create "tenant_connect_option_platforms" table
CREATE TABLE "tenant_connect_option_platforms" ("tenant_id" character varying NOT NULL, "connect_option_platform_id" character varying NOT NULL, PRIMARY KEY ("tenant_id", "connect_option_platform_id"));
-- Create "tenants" table
CREATE TABLE "tenants" ("id" character varying NOT NULL, "name" character varying NOT NULL, "vat_number" character varying NULL, "invoice_reference" character varying NULL, "plan_tenant" character varying NOT NULL, "tenant_company_address" character varying NULL, "tenant_default_language" character varying NOT NULL, "tenant_billing_contact" character varying NULL, "tenant_admin_contact" character varying NULL, PRIMARY KEY ("id"));
-- Create index "tenants_name_key" to table: "tenants"
CREATE UNIQUE INDEX "tenants_name_key" ON "tenants" ("name");
-- Create "user_seats" table
CREATE TABLE "user_seats" ("id" character varying NOT NULL, "name" character varying NULL, "surname" character varying NULL, "email" character varying NOT NULL, "password" character varying NOT NULL, "hash" character varying NOT NULL, "created_at" timestamptz NOT NULL, "tenant_id" character varying NOT NULL, PRIMARY KEY ("id"));
-- Create index "user_seats_email_key" to table: "user_seats"
CREATE UNIQUE INDEX "user_seats_email_key" ON "user_seats" ("email");
-- Create index "userseat_tenant_id" to table: "user_seats"
CREATE INDEX "userseat_tenant_id" ON "user_seats" ("tenant_id");
-- Create "users" table
CREATE TABLE "users" ("id" character varying NOT NULL, "name" character varying NULL, "surname" character varying NULL, "phone_number" character varying NULL, "email" character varying NOT NULL, "password" character varying NULL, "hash" character varying NOT NULL, "is_account_owner" boolean NOT NULL DEFAULT false, "is_global_admin" boolean NOT NULL DEFAULT false, "marketing_consent" boolean NULL DEFAULT true, "created_at" timestamptz NULL, "archived_at" timestamptz NULL, "pickup_day" character varying NOT NULL DEFAULT 'Today', "pickup_day_last_changed" timestamptz NULL, "seat_group_user" character varying NULL, "tenant_id" character varying NOT NULL, "user_language" character varying NULL, PRIMARY KEY ("id"));
-- Create index "user_tenant_id" to table: "users"
CREATE INDEX "user_tenant_id" ON "users" ("tenant_id");
-- Create index "users_email_key" to table: "users"
CREATE UNIQUE INDEX "users_email_key" ON "users" ("email");
-- Create "workspace_recent_scans" table
CREATE TABLE "workspace_recent_scans" ("id" character varying NOT NULL, "created_at" timestamptz NOT NULL, "tenant_id" character varying NOT NULL, "workspace_recent_scan_shipment_parcel" character varying NULL, "workspace_recent_scan_user" character varying NOT NULL, PRIMARY KEY ("id"));
-- Create index "workspacerecentscan_tenant_id" to table: "workspace_recent_scans"
CREATE INDEX "workspacerecentscan_tenant_id" ON "workspace_recent_scans" ("tenant_id");
-- Create "workstations" table
CREATE TABLE "workstations" ("id" character varying NOT NULL, "archived_at" timestamptz NULL, "name" character varying NOT NULL, "device_type" character varying NOT NULL DEFAULT 'label_station', "registration_code" character varying NOT NULL, "workstation_id" character varying NOT NULL, "created_at" timestamptz NOT NULL, "last_ping" timestamptz NULL, "status" character varying NOT NULL DEFAULT 'pending', "user_selected_workstation" character varying NULL, "tenant_id" character varying NOT NULL, "workstation_user" character varying NULL, PRIMARY KEY ("id"));
-- Create index "workstation_tenant_id" to table: "workstations"
CREATE INDEX "workstation_tenant_id" ON "workstations" ("tenant_id");
-- Create index "workstations_user_selected_workstation_key" to table: "workstations"
CREATE UNIQUE INDEX "workstations_user_selected_workstation_key" ON "workstations" ("user_selected_workstation");
-- Modify "address_globals" table
ALTER TABLE "address_globals" ADD CONSTRAINT "address_globals_countries_country" FOREIGN KEY ("address_global_country") REFERENCES "countries" ("id") ON UPDATE NO ACTION ON DELETE NO ACTION, ADD CONSTRAINT "address_globals_parcel_shop_brings_address_delivery" FOREIGN KEY ("parcel_shop_bring_address_delivery") REFERENCES "parcel_shop_brings" ("id") ON UPDATE NO ACTION ON DELETE SET NULL, ADD CONSTRAINT "address_globals_parcel_shop_post_nords_address_delivery" FOREIGN KEY ("parcel_shop_post_nord_address_delivery") REFERENCES "parcel_shop_post_nords" ("id") ON UPDATE NO ACTION ON DELETE SET NULL, ADD CONSTRAINT "address_globals_parcel_shops_address" FOREIGN KEY ("parcel_shop_address") REFERENCES "parcel_shops" ("id") ON UPDATE NO ACTION ON DELETE SET NULL;
-- Modify "addresses" table
ALTER TABLE "addresses" ADD CONSTRAINT "addresses_consolidations_recipient" FOREIGN KEY ("consolidation_recipient") REFERENCES "consolidations" ("id") ON UPDATE NO ACTION ON DELETE SET NULL, ADD CONSTRAINT "addresses_consolidations_sender" FOREIGN KEY ("consolidation_sender") REFERENCES "consolidations" ("id") ON UPDATE NO ACTION ON DELETE SET NULL, ADD CONSTRAINT "addresses_countries_country" FOREIGN KEY ("address_country") REFERENCES "countries" ("id") ON UPDATE NO ACTION ON DELETE NO ACTION, ADD CONSTRAINT "addresses_tenants_tenant" FOREIGN KEY ("tenant_id") REFERENCES "tenants" ("id") ON UPDATE NO ACTION ON DELETE NO ACTION;
-- Modify "api_tokens" table
ALTER TABLE "api_tokens" ADD CONSTRAINT "api_tokens_tenants_tenant" FOREIGN KEY ("tenant_id") REFERENCES "tenants" ("id") ON UPDATE NO ACTION ON DELETE NO ACTION, ADD CONSTRAINT "api_tokens_users_api_token" FOREIGN KEY ("user_api_token") REFERENCES "users" ("id") ON UPDATE NO ACTION ON DELETE NO ACTION;
-- Modify "business_hours_periods" table
ALTER TABLE "business_hours_periods" ADD CONSTRAINT "business_hours_periods_parcel_shops_business_hours_period" FOREIGN KEY ("parcel_shop_business_hours_period") REFERENCES "parcel_shops" ("id") ON UPDATE NO ACTION ON DELETE NO ACTION;
-- Modify "carrier_additional_service_brings" table
ALTER TABLE "carrier_additional_service_brings" ADD CONSTRAINT "carrier_additional_service_brings_carrier_service_brings_carrie" FOREIGN KEY ("carrier_service_bring_carrier_additional_service_bring") REFERENCES "carrier_service_brings" ("id") ON UPDATE NO ACTION ON DELETE SET NULL;
-- Modify "carrier_additional_service_gl_ss" table
ALTER TABLE "carrier_additional_service_gl_ss" ADD CONSTRAINT "carrier_additional_service_gl_ss_carrier_service_gl_ss_carrier_" FOREIGN KEY ("carrier_service_gls_carrier_additional_service_gls") REFERENCES "carrier_service_gl_ss" ("id") ON UPDATE NO ACTION ON DELETE SET NULL;
-- Modify "carrier_additional_service_gls_countries_consignee" table
ALTER TABLE "carrier_additional_service_gls_countries_consignee" ADD CONSTRAINT "carrier_additional_service_gls_countries_consignee_carrier_addi" FOREIGN KEY ("carrier_additional_service_gls_id") REFERENCES "carrier_additional_service_gl_ss" ("id") ON UPDATE NO ACTION ON DELETE CASCADE, ADD CONSTRAINT "carrier_additional_service_gls_countries_consignee_country_id" FOREIGN KEY ("country_id") REFERENCES "countries" ("id") ON UPDATE NO ACTION ON DELETE CASCADE;
-- Modify "carrier_additional_service_gls_countries_consignor" table
ALTER TABLE "carrier_additional_service_gls_countries_consignor" ADD CONSTRAINT "carrier_additional_service_gls_countries_consignor_carrier_addi" FOREIGN KEY ("carrier_additional_service_gls_id") REFERENCES "carrier_additional_service_gl_ss" ("id") ON UPDATE NO ACTION ON DELETE CASCADE, ADD CONSTRAINT "carrier_additional_service_gls_countries_consignor_country_id" FOREIGN KEY ("country_id") REFERENCES "countries" ("id") ON UPDATE NO ACTION ON DELETE CASCADE;
-- Modify "carrier_additional_service_post_nord_countries_consignee" table
ALTER TABLE "carrier_additional_service_post_nord_countries_consignee" ADD CONSTRAINT "carrier_additional_service_post_nord_countries_consignee_carrie" FOREIGN KEY ("carrier_additional_service_post_nord_id") REFERENCES "carrier_additional_service_post_nords" ("id") ON UPDATE NO ACTION ON DELETE CASCADE, ADD CONSTRAINT "carrier_additional_service_post_nord_countries_consignee_countr" FOREIGN KEY ("country_id") REFERENCES "countries" ("id") ON UPDATE NO ACTION ON DELETE CASCADE;
-- Modify "carrier_additional_service_post_nord_countries_consignor" table
ALTER TABLE "carrier_additional_service_post_nord_countries_consignor" ADD CONSTRAINT "carrier_additional_service_post_nord_countries_consignor_carrie" FOREIGN KEY ("carrier_additional_service_post_nord_id") REFERENCES "carrier_additional_service_post_nords" ("id") ON UPDATE NO ACTION ON DELETE CASCADE, ADD CONSTRAINT "carrier_additional_service_post_nord_countries_consignor_countr" FOREIGN KEY ("country_id") REFERENCES "countries" ("id") ON UPDATE NO ACTION ON DELETE CASCADE;
-- Modify "carrier_additional_service_post_nords" table
ALTER TABLE "carrier_additional_service_post_nords" ADD CONSTRAINT "carrier_additional_service_post_nords_carrier_service_post_nord" FOREIGN KEY ("carrier_service_post_nord_carrier_add_serv_post_nord") REFERENCES "carrier_service_post_nords" ("id") ON UPDATE NO ACTION ON DELETE SET NULL;
-- Modify "carrier_additional_service_usp_ss" table
ALTER TABLE "carrier_additional_service_usp_ss" ADD CONSTRAINT "carrier_additional_service_usp_ss_carrier_service_usp_ss_carrie" FOREIGN KEY ("carrier_service_usps_carrier_additional_service_usps") REFERENCES "carrier_service_usp_ss" ("id") ON UPDATE NO ACTION ON DELETE SET NULL;
-- Modify "carrier_brings" table
ALTER TABLE "carrier_brings" ADD CONSTRAINT "carrier_brings_carriers_carrier_bring" FOREIGN KEY ("carrier_carrier_bring") REFERENCES "carriers" ("id") ON UPDATE NO ACTION ON DELETE NO ACTION, ADD CONSTRAINT "carrier_brings_tenants_tenant" FOREIGN KEY ("tenant_id") REFERENCES "tenants" ("id") ON UPDATE NO ACTION ON DELETE NO ACTION;
-- Modify "carrier_da_os" table
ALTER TABLE "carrier_da_os" ADD CONSTRAINT "carrier_da_os_carriers_carrier_dao" FOREIGN KEY ("carrier_carrier_dao") REFERENCES "carriers" ("id") ON UPDATE NO ACTION ON DELETE NO ACTION, ADD CONSTRAINT "carrier_da_os_tenants_tenant" FOREIGN KEY ("tenant_id") REFERENCES "tenants" ("id") ON UPDATE NO ACTION ON DELETE NO ACTION;
-- Modify "carrier_dfs" table
ALTER TABLE "carrier_dfs" ADD CONSTRAINT "carrier_dfs_carriers_carrier_df" FOREIGN KEY ("carrier_carrier_df") REFERENCES "carriers" ("id") ON UPDATE NO ACTION ON DELETE NO ACTION, ADD CONSTRAINT "carrier_dfs_tenants_tenant" FOREIGN KEY ("tenant_id") REFERENCES "tenants" ("id") ON UPDATE NO ACTION ON DELETE NO ACTION;
-- Modify "carrier_ds_vs" table
ALTER TABLE "carrier_ds_vs" ADD CONSTRAINT "carrier_ds_vs_carriers_carrier_dsv" FOREIGN KEY ("carrier_carrier_dsv") REFERENCES "carriers" ("id") ON UPDATE NO ACTION ON DELETE NO ACTION, ADD CONSTRAINT "carrier_ds_vs_tenants_tenant" FOREIGN KEY ("tenant_id") REFERENCES "tenants" ("id") ON UPDATE NO ACTION ON DELETE NO ACTION;
-- Modify "carrier_easy_posts" table
ALTER TABLE "carrier_easy_posts" ADD CONSTRAINT "carrier_easy_posts_carriers_carrier_easy_post" FOREIGN KEY ("carrier_carrier_easy_post") REFERENCES "carriers" ("id") ON UPDATE NO ACTION ON DELETE NO ACTION, ADD CONSTRAINT "carrier_easy_posts_tenants_tenant" FOREIGN KEY ("tenant_id") REFERENCES "tenants" ("id") ON UPDATE NO ACTION ON DELETE NO ACTION;
-- Modify "carrier_gl_ss" table
ALTER TABLE "carrier_gl_ss" ADD CONSTRAINT "carrier_gl_ss_carriers_carrier_gls" FOREIGN KEY ("carrier_carrier_gls") REFERENCES "carriers" ("id") ON UPDATE NO ACTION ON DELETE NO ACTION, ADD CONSTRAINT "carrier_gl_ss_tenants_tenant" FOREIGN KEY ("tenant_id") REFERENCES "tenants" ("id") ON UPDATE NO ACTION ON DELETE NO ACTION;
-- Modify "carrier_post_nords" table
ALTER TABLE "carrier_post_nords" ADD CONSTRAINT "carrier_post_nords_carriers_carrier_post_nord" FOREIGN KEY ("carrier_carrier_post_nord") REFERENCES "carriers" ("id") ON UPDATE NO ACTION ON DELETE NO ACTION, ADD CONSTRAINT "carrier_post_nords_tenants_tenant" FOREIGN KEY ("tenant_id") REFERENCES "tenants" ("id") ON UPDATE NO ACTION ON DELETE NO ACTION;
-- Modify "carrier_service_brings" table
ALTER TABLE "carrier_service_brings" ADD CONSTRAINT "carrier_service_brings_carrier_services_carrier_service_bring" FOREIGN KEY ("carrier_service_carrier_service_bring") REFERENCES "carrier_services" ("id") ON UPDATE NO ACTION ON DELETE NO ACTION;
-- Modify "carrier_service_da_os" table
ALTER TABLE "carrier_service_da_os" ADD CONSTRAINT "carrier_service_da_os_carrier_services_carrier_service_dao" FOREIGN KEY ("carrier_service_carrier_service_dao") REFERENCES "carrier_services" ("id") ON UPDATE NO ACTION ON DELETE NO ACTION;
-- Modify "carrier_service_dao_carrier_additional_service_dao" table
ALTER TABLE "carrier_service_dao_carrier_additional_service_dao" ADD CONSTRAINT "carrier_service_dao_carrier_additional_service_dao_carrier_addi" FOREIGN KEY ("carrier_additional_service_dao_id") REFERENCES "carrier_additional_service_da_os" ("id") ON UPDATE NO ACTION ON DELETE CASCADE, ADD CONSTRAINT "carrier_service_dao_carrier_additional_service_dao_carrier_serv" FOREIGN KEY ("carrier_service_dao_id") REFERENCES "carrier_service_da_os" ("id") ON UPDATE NO ACTION ON DELETE CASCADE;
-- Modify "carrier_service_df_carrier_additional_service_df" table
ALTER TABLE "carrier_service_df_carrier_additional_service_df" ADD CONSTRAINT "carrier_service_df_carrier_additional_service_df_carrier_additi" FOREIGN KEY ("carrier_additional_service_df_id") REFERENCES "carrier_additional_service_dfs" ("id") ON UPDATE NO ACTION ON DELETE CASCADE, ADD CONSTRAINT "carrier_service_df_carrier_additional_service_df_carrier_servic" FOREIGN KEY ("carrier_service_df_id") REFERENCES "carrier_service_dfs" ("id") ON UPDATE NO ACTION ON DELETE CASCADE;
-- Modify "carrier_service_dfs" table
ALTER TABLE "carrier_service_dfs" ADD CONSTRAINT "carrier_service_dfs_carrier_services_carrier_service_df" FOREIGN KEY ("carrier_service_carrier_service_df") REFERENCES "carrier_services" ("id") ON UPDATE NO ACTION ON DELETE NO ACTION;
-- Modify "carrier_service_ds_vs" table
ALTER TABLE "carrier_service_ds_vs" ADD CONSTRAINT "carrier_service_ds_vs_carrier_services_carrier_service_dsv" FOREIGN KEY ("carrier_service_carrier_service_dsv") REFERENCES "carrier_services" ("id") ON UPDATE NO ACTION ON DELETE NO ACTION;
-- Modify "carrier_service_dsv_carrier_additional_service_dsv" table
ALTER TABLE "carrier_service_dsv_carrier_additional_service_dsv" ADD CONSTRAINT "carrier_service_dsv_carrier_additional_service_dsv_carrier_addi" FOREIGN KEY ("carrier_additional_service_dsv_id") REFERENCES "carrier_additional_service_ds_vs" ("id") ON UPDATE NO ACTION ON DELETE CASCADE, ADD CONSTRAINT "carrier_service_dsv_carrier_additional_service_dsv_carrier_serv" FOREIGN KEY ("carrier_service_dsv_id") REFERENCES "carrier_service_ds_vs" ("id") ON UPDATE NO ACTION ON DELETE CASCADE;
-- Modify "carrier_service_easy_post_carrier_add_serv_easy_post" table
ALTER TABLE "carrier_service_easy_post_carrier_add_serv_easy_post" ADD CONSTRAINT "carrier_service_easy_post_carrier_add_serv_easy_post_carrier_ad" FOREIGN KEY ("carrier_additional_service_easy_post_id") REFERENCES "carrier_additional_service_easy_posts" ("id") ON UPDATE NO ACTION ON DELETE CASCADE, ADD CONSTRAINT "carrier_service_easy_post_carrier_add_serv_easy_post_carrier_se" FOREIGN KEY ("carrier_service_easy_post_id") REFERENCES "carrier_service_easy_posts" ("id") ON UPDATE NO ACTION ON DELETE CASCADE;
-- Modify "carrier_service_easy_posts" table
ALTER TABLE "carrier_service_easy_posts" ADD CONSTRAINT "carrier_service_easy_posts_carrier_services_carrier_serv_easy_p" FOREIGN KEY ("carrier_service_carrier_serv_easy_post") REFERENCES "carrier_services" ("id") ON UPDATE NO ACTION ON DELETE NO ACTION;
-- Modify "carrier_service_gl_ss" table
ALTER TABLE "carrier_service_gl_ss" ADD CONSTRAINT "carrier_service_gl_ss_carrier_services_carrier_service_gls" FOREIGN KEY ("carrier_service_carrier_service_gls") REFERENCES "carrier_services" ("id") ON UPDATE NO ACTION ON DELETE NO ACTION;
-- Modify "carrier_service_post_nords" table
ALTER TABLE "carrier_service_post_nords" ADD CONSTRAINT "carrier_service_post_nords_carrier_services_carrier_service_pos" FOREIGN KEY ("carrier_service_carrier_service_post_nord") REFERENCES "carrier_services" ("id") ON UPDATE NO ACTION ON DELETE NO ACTION;
-- Modify "carrier_service_usp_ss" table
ALTER TABLE "carrier_service_usp_ss" ADD CONSTRAINT "carrier_service_usp_ss_carrier_services_carrier_service_usps" FOREIGN KEY ("carrier_service_carrier_service_usps") REFERENCES "carrier_services" ("id") ON UPDATE NO ACTION ON DELETE NO ACTION;
-- Modify "carrier_services" table
ALTER TABLE "carrier_services" ADD CONSTRAINT "carrier_services_carrier_brands_carrier_service" FOREIGN KEY ("carrier_brand_carrier_service") REFERENCES "carrier_brands" ("id") ON UPDATE NO ACTION ON DELETE NO ACTION;
-- Modify "carrier_usp_ss" table
ALTER TABLE "carrier_usp_ss" ADD CONSTRAINT "carrier_usp_ss_carriers_carrier_usps" FOREIGN KEY ("carrier_carrier_usps") REFERENCES "carriers" ("id") ON UPDATE NO ACTION ON DELETE NO ACTION, ADD CONSTRAINT "carrier_usp_ss_tenants_tenant" FOREIGN KEY ("tenant_id") REFERENCES "tenants" ("id") ON UPDATE NO ACTION ON DELETE NO ACTION;
-- Modify "carriers" table
ALTER TABLE "carriers" ADD CONSTRAINT "carriers_carrier_brands_carrier_brand" FOREIGN KEY ("carrier_carrier_brand") REFERENCES "carrier_brands" ("id") ON UPDATE NO ACTION ON DELETE NO ACTION, ADD CONSTRAINT "carriers_tenants_tenant" FOREIGN KEY ("tenant_id") REFERENCES "tenants" ("id") ON UPDATE NO ACTION ON DELETE NO ACTION;
-- Modify "change_histories" table
ALTER TABLE "change_histories" ADD CONSTRAINT "change_histories_tenants_tenant" FOREIGN KEY ("tenant_id") REFERENCES "tenants" ("id") ON UPDATE NO ACTION ON DELETE NO ACTION, ADD CONSTRAINT "change_histories_users_user" FOREIGN KEY ("change_history_user") REFERENCES "users" ("id") ON UPDATE NO ACTION ON DELETE SET NULL;
-- Modify "colli_cancelled_shipment_parcel" table
ALTER TABLE "colli_cancelled_shipment_parcel" ADD CONSTRAINT "colli_cancelled_shipment_parcel_colli_id" FOREIGN KEY ("colli_id") REFERENCES "collis" ("id") ON UPDATE NO ACTION ON DELETE CASCADE, ADD CONSTRAINT "colli_cancelled_shipment_parcel_shipment_parcel_id" FOREIGN KEY ("shipment_parcel_id") REFERENCES "shipment_parcels" ("id") ON UPDATE NO ACTION ON DELETE CASCADE;
-- Modify "collis" table
ALTER TABLE "collis" ADD CONSTRAINT "collis_addresses_recipient" FOREIGN KEY ("colli_recipient") REFERENCES "addresses" ("id") ON UPDATE NO ACTION ON DELETE NO ACTION, ADD CONSTRAINT "collis_addresses_sender" FOREIGN KEY ("colli_sender") REFERENCES "addresses" ("id") ON UPDATE NO ACTION ON DELETE NO ACTION, ADD CONSTRAINT "collis_delivery_options_delivery_option" FOREIGN KEY ("colli_delivery_option") REFERENCES "delivery_options" ("id") ON UPDATE NO ACTION ON DELETE SET NULL, ADD CONSTRAINT "collis_locations_click_collect_location" FOREIGN KEY ("colli_click_collect_location") REFERENCES "locations" ("id") ON UPDATE NO ACTION ON DELETE SET NULL, ADD CONSTRAINT "collis_orders_colli" FOREIGN KEY ("order_colli") REFERENCES "orders" ("id") ON UPDATE NO ACTION ON DELETE NO ACTION, ADD CONSTRAINT "collis_packagings_packaging" FOREIGN KEY ("colli_packaging") REFERENCES "packagings" ("id") ON UPDATE NO ACTION ON DELETE SET NULL, ADD CONSTRAINT "collis_parcel_shops_parcel_shop" FOREIGN KEY ("colli_parcel_shop") REFERENCES "parcel_shops" ("id") ON UPDATE NO ACTION ON DELETE SET NULL, ADD CONSTRAINT "collis_tenants_tenant" FOREIGN KEY ("tenant_id") REFERENCES "tenants" ("id") ON UPDATE NO ACTION ON DELETE NO ACTION;
-- Modify "connection_lookups" table
ALTER TABLE "connection_lookups" ADD CONSTRAINT "connection_lookups_connections_connections" FOREIGN KEY ("connection_lookup_connections") REFERENCES "connections" ("id") ON UPDATE NO ACTION ON DELETE SET NULL, ADD CONSTRAINT "connection_lookups_tenants_tenant" FOREIGN KEY ("tenant_id") REFERENCES "tenants" ("id") ON UPDATE NO ACTION ON DELETE NO ACTION;
-- Modify "connection_shopifies" table
ALTER TABLE "connection_shopifies" ADD CONSTRAINT "connection_shopifies_connections_connection_shopify" FOREIGN KEY ("connection_connection_shopify") REFERENCES "connections" ("id") ON UPDATE NO ACTION ON DELETE NO ACTION, ADD CONSTRAINT "connection_shopifies_tenants_tenant" FOREIGN KEY ("tenant_id") REFERENCES "tenants" ("id") ON UPDATE NO ACTION ON DELETE NO ACTION;
-- Modify "connections" table
ALTER TABLE "connections" ADD CONSTRAINT "connections_connection_brands_connection_brand" FOREIGN KEY ("connection_connection_brand") REFERENCES "connection_brands" ("id") ON UPDATE NO ACTION ON DELETE NO ACTION, ADD CONSTRAINT "connections_currencies_currency" FOREIGN KEY ("connection_currency") REFERENCES "currencies" ("id") ON UPDATE NO ACTION ON DELETE NO ACTION, ADD CONSTRAINT "connections_documents_packing_slip_template" FOREIGN KEY ("connection_packing_slip_template") REFERENCES "documents" ("id") ON UPDATE NO ACTION ON DELETE SET NULL, ADD CONSTRAINT "connections_locations_pickup_location" FOREIGN KEY ("connection_pickup_location") REFERENCES "locations" ("id") ON UPDATE NO ACTION ON DELETE NO ACTION, ADD CONSTRAINT "connections_locations_return_location" FOREIGN KEY ("connection_return_location") REFERENCES "locations" ("id") ON UPDATE NO ACTION ON DELETE NO ACTION, ADD CONSTRAINT "connections_locations_seller_location" FOREIGN KEY ("connection_seller_location") REFERENCES "locations" ("id") ON UPDATE NO ACTION ON DELETE NO ACTION, ADD CONSTRAINT "connections_locations_sender_location" FOREIGN KEY ("connection_sender_location") REFERENCES "locations" ("id") ON UPDATE NO ACTION ON DELETE NO ACTION, ADD CONSTRAINT "connections_return_portals_connection" FOREIGN KEY ("return_portal_connection") REFERENCES "return_portals" ("id") ON UPDATE NO ACTION ON DELETE SET NULL, ADD CONSTRAINT "connections_tenants_tenant" FOREIGN KEY ("tenant_id") REFERENCES "tenants" ("id") ON UPDATE NO ACTION ON DELETE NO ACTION;
-- Modify "consolidations" table
ALTER TABLE "consolidations" ADD CONSTRAINT "consolidations_delivery_options_delivery_option" FOREIGN KEY ("consolidation_delivery_option") REFERENCES "delivery_options" ("id") ON UPDATE NO ACTION ON DELETE SET NULL, ADD CONSTRAINT "consolidations_shipments_consolidation" FOREIGN KEY ("shipment_consolidation") REFERENCES "shipments" ("id") ON UPDATE NO ACTION ON DELETE SET NULL, ADD CONSTRAINT "consolidations_tenants_tenant" FOREIGN KEY ("tenant_id") REFERENCES "tenants" ("id") ON UPDATE NO ACTION ON DELETE NO ACTION;
-- Modify "contacts" table
ALTER TABLE "contacts" ADD CONSTRAINT "contacts_tenants_tenant" FOREIGN KEY ("tenant_id") REFERENCES "tenants" ("id") ON UPDATE NO ACTION ON DELETE NO ACTION;
-- Modify "country_delivery_rule" table
ALTER TABLE "country_delivery_rule" ADD CONSTRAINT "country_delivery_rule_country_id" FOREIGN KEY ("country_id") REFERENCES "countries" ("id") ON UPDATE NO ACTION ON DELETE CASCADE, ADD CONSTRAINT "country_delivery_rule_delivery_rule_id" FOREIGN KEY ("delivery_rule_id") REFERENCES "delivery_rules" ("id") ON UPDATE NO ACTION ON DELETE CASCADE;
-- Modify "country_harmonized_codes" table
ALTER TABLE "country_harmonized_codes" ADD CONSTRAINT "country_harmonized_codes_countries_country" FOREIGN KEY ("country_harmonized_code_country") REFERENCES "countries" ("id") ON UPDATE NO ACTION ON DELETE NO ACTION, ADD CONSTRAINT "country_harmonized_codes_inventory_items_country_harmonized_cod" FOREIGN KEY ("inventory_item_country_harmonized_code") REFERENCES "inventory_items" ("id") ON UPDATE NO ACTION ON DELETE NO ACTION, ADD CONSTRAINT "country_harmonized_codes_tenants_tenant" FOREIGN KEY ("tenant_id") REFERENCES "tenants" ("id") ON UPDATE NO ACTION ON DELETE NO ACTION;
-- Modify "delivery_option_bring_carrier_additional_service_bring" table
ALTER TABLE "delivery_option_bring_carrier_additional_service_bring" ADD CONSTRAINT "delivery_option_bring_carrier_additional_service_bring_carrier_" FOREIGN KEY ("carrier_additional_service_bring_id") REFERENCES "carrier_additional_service_brings" ("id") ON UPDATE NO ACTION ON DELETE CASCADE, ADD CONSTRAINT "delivery_option_bring_carrier_additional_service_bring_delivery" FOREIGN KEY ("delivery_option_bring_id") REFERENCES "delivery_option_brings" ("id") ON UPDATE NO ACTION ON DELETE CASCADE;
-- Modify "delivery_option_brings" table
ALTER TABLE "delivery_option_brings" ADD CONSTRAINT "delivery_option_brings_delivery_options_delivery_option_bring" FOREIGN KEY ("delivery_option_delivery_option_bring") REFERENCES "delivery_options" ("id") ON UPDATE NO ACTION ON DELETE NO ACTION, ADD CONSTRAINT "delivery_option_brings_tenants_tenant" FOREIGN KEY ("tenant_id") REFERENCES "tenants" ("id") ON UPDATE NO ACTION ON DELETE NO ACTION;
-- Modify "delivery_option_click_collect_location" table
ALTER TABLE "delivery_option_click_collect_location" ADD CONSTRAINT "delivery_option_click_collect_location_delivery_option_id" FOREIGN KEY ("delivery_option_id") REFERENCES "delivery_options" ("id") ON UPDATE NO ACTION ON DELETE CASCADE, ADD CONSTRAINT "delivery_option_click_collect_location_location_id" FOREIGN KEY ("location_id") REFERENCES "locations" ("id") ON UPDATE NO ACTION ON DELETE CASCADE;
-- Modify "delivery_option_da_os" table
ALTER TABLE "delivery_option_da_os" ADD CONSTRAINT "delivery_option_da_os_delivery_options_delivery_option_dao" FOREIGN KEY ("delivery_option_delivery_option_dao") REFERENCES "delivery_options" ("id") ON UPDATE NO ACTION ON DELETE NO ACTION, ADD CONSTRAINT "delivery_option_da_os_tenants_tenant" FOREIGN KEY ("tenant_id") REFERENCES "tenants" ("id") ON UPDATE NO ACTION ON DELETE NO ACTION;
-- Modify "delivery_option_dao_carrier_additional_service_dao" table
ALTER TABLE "delivery_option_dao_carrier_additional_service_dao" ADD CONSTRAINT "delivery_option_dao_carrier_additional_service_dao_carrier_addi" FOREIGN KEY ("carrier_additional_service_dao_id") REFERENCES "carrier_additional_service_da_os" ("id") ON UPDATE NO ACTION ON DELETE CASCADE, ADD CONSTRAINT "delivery_option_dao_carrier_additional_service_dao_delivery_opt" FOREIGN KEY ("delivery_option_dao_id") REFERENCES "delivery_option_da_os" ("id") ON UPDATE NO ACTION ON DELETE CASCADE;
-- Modify "delivery_option_df_carrier_additional_service_df" table
ALTER TABLE "delivery_option_df_carrier_additional_service_df" ADD CONSTRAINT "delivery_option_df_carrier_additional_service_df_carrier_additi" FOREIGN KEY ("carrier_additional_service_df_id") REFERENCES "carrier_additional_service_dfs" ("id") ON UPDATE NO ACTION ON DELETE CASCADE, ADD CONSTRAINT "delivery_option_df_carrier_additional_service_df_delivery_optio" FOREIGN KEY ("delivery_option_df_id") REFERENCES "delivery_option_dfs" ("id") ON UPDATE NO ACTION ON DELETE CASCADE;
-- Modify "delivery_option_dfs" table
ALTER TABLE "delivery_option_dfs" ADD CONSTRAINT "delivery_option_dfs_delivery_options_delivery_option_df" FOREIGN KEY ("delivery_option_delivery_option_df") REFERENCES "delivery_options" ("id") ON UPDATE NO ACTION ON DELETE NO ACTION, ADD CONSTRAINT "delivery_option_dfs_tenants_tenant" FOREIGN KEY ("tenant_id") REFERENCES "tenants" ("id") ON UPDATE NO ACTION ON DELETE NO ACTION;
-- Modify "delivery_option_ds_vs" table
ALTER TABLE "delivery_option_ds_vs" ADD CONSTRAINT "delivery_option_ds_vs_delivery_options_delivery_option_dsv" FOREIGN KEY ("delivery_option_delivery_option_dsv") REFERENCES "delivery_options" ("id") ON UPDATE NO ACTION ON DELETE NO ACTION, ADD CONSTRAINT "delivery_option_ds_vs_tenants_tenant" FOREIGN KEY ("tenant_id") REFERENCES "tenants" ("id") ON UPDATE NO ACTION ON DELETE NO ACTION;
-- Modify "delivery_option_dsv_carrier_additional_service_dsv" table
ALTER TABLE "delivery_option_dsv_carrier_additional_service_dsv" ADD CONSTRAINT "delivery_option_dsv_carrier_additional_service_dsv_carrier_addi" FOREIGN KEY ("carrier_additional_service_dsv_id") REFERENCES "carrier_additional_service_ds_vs" ("id") ON UPDATE NO ACTION ON DELETE CASCADE, ADD CONSTRAINT "delivery_option_dsv_carrier_additional_service_dsv_delivery_opt" FOREIGN KEY ("delivery_option_dsv_id") REFERENCES "delivery_option_ds_vs" ("id") ON UPDATE NO ACTION ON DELETE CASCADE;
-- Modify "delivery_option_easy_post_carrier_add_serv_easy_post" table
ALTER TABLE "delivery_option_easy_post_carrier_add_serv_easy_post" ADD CONSTRAINT "delivery_option_easy_post_carrier_add_serv_easy_post_carrier_ad" FOREIGN KEY ("carrier_additional_service_easy_post_id") REFERENCES "carrier_additional_service_easy_posts" ("id") ON UPDATE NO ACTION ON DELETE CASCADE, ADD CONSTRAINT "delivery_option_easy_post_carrier_add_serv_easy_post_delivery_o" FOREIGN KEY ("delivery_option_easy_post_id") REFERENCES "delivery_option_easy_posts" ("id") ON UPDATE NO ACTION ON DELETE CASCADE;
-- Modify "delivery_option_easy_posts" table
ALTER TABLE "delivery_option_easy_posts" ADD CONSTRAINT "delivery_option_easy_posts_delivery_options_delivery_option_eas" FOREIGN KEY ("delivery_option_delivery_option_easy_post") REFERENCES "delivery_options" ("id") ON UPDATE NO ACTION ON DELETE NO ACTION, ADD CONSTRAINT "delivery_option_easy_posts_tenants_tenant" FOREIGN KEY ("tenant_id") REFERENCES "tenants" ("id") ON UPDATE NO ACTION ON DELETE NO ACTION;
-- Modify "delivery_option_gl_ss" table
ALTER TABLE "delivery_option_gl_ss" ADD CONSTRAINT "delivery_option_gl_ss_delivery_options_delivery_option_gls" FOREIGN KEY ("delivery_option_delivery_option_gls") REFERENCES "delivery_options" ("id") ON UPDATE NO ACTION ON DELETE NO ACTION, ADD CONSTRAINT "delivery_option_gl_ss_tenants_tenant" FOREIGN KEY ("tenant_id") REFERENCES "tenants" ("id") ON UPDATE NO ACTION ON DELETE NO ACTION;
-- Modify "delivery_option_gls_carrier_additional_service_gls" table
ALTER TABLE "delivery_option_gls_carrier_additional_service_gls" ADD CONSTRAINT "delivery_option_gls_carrier_additional_service_gls_carrier_addi" FOREIGN KEY ("carrier_additional_service_gls_id") REFERENCES "carrier_additional_service_gl_ss" ("id") ON UPDATE NO ACTION ON DELETE CASCADE, ADD CONSTRAINT "delivery_option_gls_carrier_additional_service_gls_delivery_opt" FOREIGN KEY ("delivery_option_gls_id") REFERENCES "delivery_option_gl_ss" ("id") ON UPDATE NO ACTION ON DELETE CASCADE;
-- Modify "delivery_option_post_nord_carrier_add_serv_post_nord" table
ALTER TABLE "delivery_option_post_nord_carrier_add_serv_post_nord" ADD CONSTRAINT "delivery_option_post_nord_carrier_add_serv_post_nord_carrier_ad" FOREIGN KEY ("carrier_additional_service_post_nord_id") REFERENCES "carrier_additional_service_post_nords" ("id") ON UPDATE NO ACTION ON DELETE CASCADE, ADD CONSTRAINT "delivery_option_post_nord_carrier_add_serv_post_nord_delivery_o" FOREIGN KEY ("delivery_option_post_nord_id") REFERENCES "delivery_option_post_nords" ("id") ON UPDATE NO ACTION ON DELETE CASCADE;
-- Modify "delivery_option_post_nords" table
ALTER TABLE "delivery_option_post_nords" ADD CONSTRAINT "delivery_option_post_nords_delivery_options_delivery_option_pos" FOREIGN KEY ("delivery_option_delivery_option_post_nord") REFERENCES "delivery_options" ("id") ON UPDATE NO ACTION ON DELETE NO ACTION, ADD CONSTRAINT "delivery_option_post_nords_tenants_tenant" FOREIGN KEY ("tenant_id") REFERENCES "tenants" ("id") ON UPDATE NO ACTION ON DELETE NO ACTION;
-- Modify "delivery_option_usp_ss" table
ALTER TABLE "delivery_option_usp_ss" ADD CONSTRAINT "delivery_option_usp_ss_delivery_options_delivery_option_usps" FOREIGN KEY ("delivery_option_delivery_option_usps") REFERENCES "delivery_options" ("id") ON UPDATE NO ACTION ON DELETE NO ACTION, ADD CONSTRAINT "delivery_option_usp_ss_tenants_tenant" FOREIGN KEY ("tenant_id") REFERENCES "tenants" ("id") ON UPDATE NO ACTION ON DELETE NO ACTION;
-- Modify "delivery_option_usps_carrier_additional_service_usps" table
ALTER TABLE "delivery_option_usps_carrier_additional_service_usps" ADD CONSTRAINT "delivery_option_usps_carrier_additional_service_usps_carrier_ad" FOREIGN KEY ("carrier_additional_service_usps_id") REFERENCES "carrier_additional_service_usp_ss" ("id") ON UPDATE NO ACTION ON DELETE CASCADE, ADD CONSTRAINT "delivery_option_usps_carrier_additional_service_usps_delivery_o" FOREIGN KEY ("delivery_option_usps_id") REFERENCES "delivery_option_usp_ss" ("id") ON UPDATE NO ACTION ON DELETE CASCADE;
-- Modify "delivery_options" table
ALTER TABLE "delivery_options" ADD CONSTRAINT "delivery_options_carrier_services_carrier_service" FOREIGN KEY ("delivery_option_carrier_service") REFERENCES "carrier_services" ("id") ON UPDATE NO ACTION ON DELETE NO ACTION, ADD CONSTRAINT "delivery_options_carriers_carrier" FOREIGN KEY ("delivery_option_carrier") REFERENCES "carriers" ("id") ON UPDATE NO ACTION ON DELETE NO ACTION, ADD CONSTRAINT "delivery_options_connections_default_delivery_option" FOREIGN KEY ("connection_default_delivery_option") REFERENCES "connections" ("id") ON UPDATE NO ACTION ON DELETE SET NULL, ADD CONSTRAINT "delivery_options_connections_delivery_option" FOREIGN KEY ("connection_delivery_option") REFERENCES "connections" ("id") ON UPDATE NO ACTION ON DELETE NO ACTION, ADD CONSTRAINT "delivery_options_email_templates_email_click_collect_at_store" FOREIGN KEY ("delivery_option_email_click_collect_at_store") REFERENCES "email_templates" ("id") ON UPDATE NO ACTION ON DELETE SET NULL, ADD CONSTRAINT "delivery_options_packagings_default_packaging" FOREIGN KEY ("delivery_option_default_packaging") REFERENCES "packagings" ("id") ON UPDATE NO ACTION ON DELETE SET NULL, ADD CONSTRAINT "delivery_options_tenants_tenant" FOREIGN KEY ("tenant_id") REFERENCES "tenants" ("id") ON UPDATE NO ACTION ON DELETE NO ACTION;
-- Modify "delivery_rule_constraint_groups" table
ALTER TABLE "delivery_rule_constraint_groups" ADD CONSTRAINT "delivery_rule_constraint_groups_delivery_rules_delivery_rule_co" FOREIGN KEY ("delivery_rule_delivery_rule_constraint_group") REFERENCES "delivery_rules" ("id") ON UPDATE NO ACTION ON DELETE NO ACTION, ADD CONSTRAINT "delivery_rule_constraint_groups_tenants_tenant" FOREIGN KEY ("tenant_id") REFERENCES "tenants" ("id") ON UPDATE NO ACTION ON DELETE NO ACTION;
-- Modify "delivery_rule_constraints" table
ALTER TABLE "delivery_rule_constraints" ADD CONSTRAINT "delivery_rule_constraints_delivery_rule_constraint_groups_deliv" FOREIGN KEY ("delivery_rule_constraint_group_delivery_rule_constraints") REFERENCES "delivery_rule_constraint_groups" ("id") ON UPDATE NO ACTION ON DELETE NO ACTION, ADD CONSTRAINT "delivery_rule_constraints_tenants_tenant" FOREIGN KEY ("tenant_id") REFERENCES "tenants" ("id") ON UPDATE NO ACTION ON DELETE NO ACTION;
-- Modify "delivery_rules" table
ALTER TABLE "delivery_rules" ADD CONSTRAINT "delivery_rules_currencies_currency" FOREIGN KEY ("delivery_rule_currency") REFERENCES "currencies" ("id") ON UPDATE NO ACTION ON DELETE SET NULL, ADD CONSTRAINT "delivery_rules_delivery_options_delivery_rule" FOREIGN KEY ("delivery_option_delivery_rule") REFERENCES "delivery_options" ("id") ON UPDATE NO ACTION ON DELETE SET NULL, ADD CONSTRAINT "delivery_rules_tenants_tenant" FOREIGN KEY ("tenant_id") REFERENCES "tenants" ("id") ON UPDATE NO ACTION ON DELETE NO ACTION;
-- Modify "document_files" table
ALTER TABLE "document_files" ADD CONSTRAINT "document_files_collis_document_file" FOREIGN KEY ("colli_document_file") REFERENCES "collis" ("id") ON UPDATE NO ACTION ON DELETE SET NULL, ADD CONSTRAINT "document_files_tenants_tenant" FOREIGN KEY ("tenant_id") REFERENCES "tenants" ("id") ON UPDATE NO ACTION ON DELETE NO ACTION;
-- Modify "documents" table
ALTER TABLE "documents" ADD CONSTRAINT "documents_carrier_brands_carrier_brand" FOREIGN KEY ("document_carrier_brand") REFERENCES "carrier_brands" ("id") ON UPDATE NO ACTION ON DELETE SET NULL, ADD CONSTRAINT "documents_tenants_tenant" FOREIGN KEY ("tenant_id") REFERENCES "tenants" ("id") ON UPDATE NO ACTION ON DELETE NO ACTION;
-- Modify "email_templates" table
ALTER TABLE "email_templates" ADD CONSTRAINT "email_templates_tenants_tenant" FOREIGN KEY ("tenant_id") REFERENCES "tenants" ("id") ON UPDATE NO ACTION ON DELETE NO ACTION;
-- Modify "hypothesis_test_delivery_option_delivery_option_group_one" table
ALTER TABLE "hypothesis_test_delivery_option_delivery_option_group_one" ADD CONSTRAINT "hypothesis_test_delivery_option_delivery_option_group_one_deliv" FOREIGN KEY ("delivery_option_id") REFERENCES "delivery_options" ("id") ON UPDATE NO ACTION ON DELETE CASCADE, ADD CONSTRAINT "hypothesis_test_delivery_option_delivery_option_group_one_hypot" FOREIGN KEY ("hypothesis_test_delivery_option_id") REFERENCES "hypothesis_test_delivery_options" ("id") ON UPDATE NO ACTION ON DELETE CASCADE;
-- Modify "hypothesis_test_delivery_option_delivery_option_group_two" table
ALTER TABLE "hypothesis_test_delivery_option_delivery_option_group_two" ADD CONSTRAINT "hypothesis_test_delivery_option_delivery_option_group_two_deliv" FOREIGN KEY ("delivery_option_id") REFERENCES "delivery_options" ("id") ON UPDATE NO ACTION ON DELETE CASCADE, ADD CONSTRAINT "hypothesis_test_delivery_option_delivery_option_group_two_hypot" FOREIGN KEY ("hypothesis_test_delivery_option_id") REFERENCES "hypothesis_test_delivery_options" ("id") ON UPDATE NO ACTION ON DELETE CASCADE;
-- Modify "hypothesis_test_delivery_option_lookups" table
ALTER TABLE "hypothesis_test_delivery_option_lookups" ADD CONSTRAINT "hypothesis_test_delivery_option_lookups_delivery_options_delive" FOREIGN KEY ("hypothesis_test_delivery_option_lookup_delivery_option") REFERENCES "delivery_options" ("id") ON UPDATE NO ACTION ON DELETE NO ACTION, ADD CONSTRAINT "hypothesis_test_delivery_option_lookups_hypothesis_test_deliver" FOREIGN KEY ("hypothesis_test_delivery_option_request_hypothesis_test_deliver") REFERENCES "hypothesis_test_delivery_option_requests" ("id") ON UPDATE NO ACTION ON DELETE NO ACTION, ADD CONSTRAINT "hypothesis_test_delivery_option_lookups_tenants_tenant" FOREIGN KEY ("tenant_id") REFERENCES "tenants" ("id") ON UPDATE NO ACTION ON DELETE NO ACTION;
-- Modify "hypothesis_test_delivery_option_requests" table
ALTER TABLE "hypothesis_test_delivery_option_requests" ADD CONSTRAINT "hypothesis_test_delivery_option_requests_hypothesis_test_delive" FOREIGN KEY ("hypothesis_test_delivery_option_request_hypothesis_test_deliver") REFERENCES "hypothesis_test_delivery_options" ("id") ON UPDATE NO ACTION ON DELETE NO ACTION, ADD CONSTRAINT "hypothesis_test_delivery_option_requests_tenants_tenant" FOREIGN KEY ("tenant_id") REFERENCES "tenants" ("id") ON UPDATE NO ACTION ON DELETE NO ACTION;
-- Modify "hypothesis_test_delivery_options" table
ALTER TABLE "hypothesis_test_delivery_options" ADD CONSTRAINT "hypothesis_test_delivery_options_hypothesis_tests_hypothesis_te" FOREIGN KEY ("hypothesis_test_hypothesis_test_delivery_option") REFERENCES "hypothesis_tests" ("id") ON UPDATE NO ACTION ON DELETE NO ACTION, ADD CONSTRAINT "hypothesis_test_delivery_options_tenants_tenant" FOREIGN KEY ("tenant_id") REFERENCES "tenants" ("id") ON UPDATE NO ACTION ON DELETE NO ACTION;
-- Modify "hypothesis_tests" table
ALTER TABLE "hypothesis_tests" ADD CONSTRAINT "hypothesis_tests_connections_connection" FOREIGN KEY ("hypothesis_test_connection") REFERENCES "connections" ("id") ON UPDATE NO ACTION ON DELETE NO ACTION, ADD CONSTRAINT "hypothesis_tests_tenants_tenant" FOREIGN KEY ("tenant_id") REFERENCES "tenants" ("id") ON UPDATE NO ACTION ON DELETE NO ACTION;
-- Modify "inventory_items" table
ALTER TABLE "inventory_items" ADD CONSTRAINT "inventory_items_countries_country_of_origin" FOREIGN KEY ("inventory_item_country_of_origin") REFERENCES "countries" ("id") ON UPDATE NO ACTION ON DELETE SET NULL, ADD CONSTRAINT "inventory_items_product_variants_inventory_item" FOREIGN KEY ("product_variant_inventory_item") REFERENCES "product_variants" ("id") ON UPDATE NO ACTION ON DELETE NO ACTION, ADD CONSTRAINT "inventory_items_tenants_tenant" FOREIGN KEY ("tenant_id") REFERENCES "tenants" ("id") ON UPDATE NO ACTION ON DELETE NO ACTION;
-- Modify "location_location_tags" table
ALTER TABLE "location_location_tags" ADD CONSTRAINT "location_location_tags_location_id" FOREIGN KEY ("location_id") REFERENCES "locations" ("id") ON UPDATE NO ACTION ON DELETE CASCADE, ADD CONSTRAINT "location_location_tags_location_tag_id" FOREIGN KEY ("location_tag_id") REFERENCES "location_tags" ("id") ON UPDATE NO ACTION ON DELETE CASCADE;
-- Modify "locations" table
ALTER TABLE "locations" ADD CONSTRAINT "locations_addresses_address" FOREIGN KEY ("location_address") REFERENCES "addresses" ("id") ON UPDATE NO ACTION ON DELETE NO ACTION, ADD CONSTRAINT "locations_tenants_tenant" FOREIGN KEY ("tenant_id") REFERENCES "tenants" ("id") ON UPDATE NO ACTION ON DELETE NO ACTION;
-- Modify "notifications" table
ALTER TABLE "notifications" ADD CONSTRAINT "notifications_connections_connection" FOREIGN KEY ("notification_connection") REFERENCES "connections" ("id") ON UPDATE NO ACTION ON DELETE NO ACTION, ADD CONSTRAINT "notifications_email_templates_email_template" FOREIGN KEY ("notification_email_template") REFERENCES "email_templates" ("id") ON UPDATE NO ACTION ON DELETE NO ACTION, ADD CONSTRAINT "notifications_tenants_tenant" FOREIGN KEY ("tenant_id") REFERENCES "tenants" ("id") ON UPDATE NO ACTION ON DELETE NO ACTION;
-- Modify "order_histories" table
ALTER TABLE "order_histories" ADD CONSTRAINT "order_histories_change_histories_order_history" FOREIGN KEY ("change_history_order_history") REFERENCES "change_histories" ("id") ON UPDATE NO ACTION ON DELETE NO ACTION, ADD CONSTRAINT "order_histories_orders_order_history" FOREIGN KEY ("order_order_history") REFERENCES "orders" ("id") ON UPDATE NO ACTION ON DELETE NO ACTION, ADD CONSTRAINT "order_histories_tenants_tenant" FOREIGN KEY ("tenant_id") REFERENCES "tenants" ("id") ON UPDATE NO ACTION ON DELETE NO ACTION;
-- Modify "order_lines" table
ALTER TABLE "order_lines" ADD CONSTRAINT "order_lines_collis_order_lines" FOREIGN KEY ("colli_id") REFERENCES "collis" ("id") ON UPDATE NO ACTION ON DELETE NO ACTION, ADD CONSTRAINT "order_lines_currencies_currency" FOREIGN KEY ("order_line_currency") REFERENCES "currencies" ("id") ON UPDATE NO ACTION ON DELETE NO ACTION, ADD CONSTRAINT "order_lines_product_variants_product_variant" FOREIGN KEY ("product_variant_id") REFERENCES "product_variants" ("id") ON UPDATE NO ACTION ON DELETE NO ACTION, ADD CONSTRAINT "order_lines_tenants_tenant" FOREIGN KEY ("tenant_id") REFERENCES "tenants" ("id") ON UPDATE NO ACTION ON DELETE NO ACTION;
-- Modify "order_senders" table
ALTER TABLE "order_senders" ADD CONSTRAINT "order_senders_tenants_tenant" FOREIGN KEY ("tenant_id") REFERENCES "tenants" ("id") ON UPDATE NO ACTION ON DELETE NO ACTION;
-- Modify "orders" table
ALTER TABLE "orders" ADD CONSTRAINT "orders_connections_orders" FOREIGN KEY ("connection_orders") REFERENCES "connections" ("id") ON UPDATE NO ACTION ON DELETE NO ACTION, ADD CONSTRAINT "orders_consolidations_orders" FOREIGN KEY ("consolidation_orders") REFERENCES "consolidations" ("id") ON UPDATE NO ACTION ON DELETE SET NULL, ADD CONSTRAINT "orders_hypothesis_test_delivery_option_requests_order" FOREIGN KEY ("hypothesis_test_delivery_option_request_order") REFERENCES "hypothesis_test_delivery_option_requests" ("id") ON UPDATE NO ACTION ON DELETE SET NULL, ADD CONSTRAINT "orders_pallets_orders" FOREIGN KEY ("pallet_orders") REFERENCES "pallets" ("id") ON UPDATE NO ACTION ON DELETE SET NULL, ADD CONSTRAINT "orders_tenants_tenant" FOREIGN KEY ("tenant_id") REFERENCES "tenants" ("id") ON UPDATE NO ACTION ON DELETE NO ACTION;
-- Modify "otk_requests" table
ALTER TABLE "otk_requests" ADD CONSTRAINT "otk_requests_tenants_tenant" FOREIGN KEY ("tenant_id") REFERENCES "tenants" ("id") ON UPDATE NO ACTION ON DELETE NO ACTION, ADD CONSTRAINT "otk_requests_users_otk_requests" FOREIGN KEY ("user_otk_requests") REFERENCES "users" ("id") ON UPDATE NO ACTION ON DELETE SET NULL;
-- Modify "packaging_dfs" table
ALTER TABLE "packaging_dfs" ADD CONSTRAINT "packaging_dfs_packagings_packaging_df" FOREIGN KEY ("packaging_packaging_df") REFERENCES "packagings" ("id") ON UPDATE NO ACTION ON DELETE NO ACTION, ADD CONSTRAINT "packaging_dfs_tenants_tenant" FOREIGN KEY ("tenant_id") REFERENCES "tenants" ("id") ON UPDATE NO ACTION ON DELETE NO ACTION;
-- Modify "packaging_usp_ss" table
ALTER TABLE "packaging_usp_ss" ADD CONSTRAINT "packaging_usp_ss_packaging_usps_processing_categories_packaging" FOREIGN KEY ("packaging_usps_packaging_usps_processing_category") REFERENCES "packaging_usps_processing_categories" ("id") ON UPDATE NO ACTION ON DELETE NO ACTION, ADD CONSTRAINT "packaging_usp_ss_packaging_usps_rate_indicators_packaging_usps_" FOREIGN KEY ("packaging_usps_packaging_usps_rate_indicator") REFERENCES "packaging_usps_rate_indicators" ("id") ON UPDATE NO ACTION ON DELETE NO ACTION, ADD CONSTRAINT "packaging_usp_ss_packagings_packaging_usps" FOREIGN KEY ("packaging_packaging_usps") REFERENCES "packagings" ("id") ON UPDATE NO ACTION ON DELETE NO ACTION, ADD CONSTRAINT "packaging_usp_ss_tenants_tenant" FOREIGN KEY ("tenant_id") REFERENCES "tenants" ("id") ON UPDATE NO ACTION ON DELETE NO ACTION;
-- Modify "packagings" table
ALTER TABLE "packagings" ADD CONSTRAINT "packagings_carrier_brands_carrier_brand" FOREIGN KEY ("packaging_carrier_brand") REFERENCES "carrier_brands" ("id") ON UPDATE NO ACTION ON DELETE SET NULL, ADD CONSTRAINT "packagings_tenants_tenant" FOREIGN KEY ("tenant_id") REFERENCES "tenants" ("id") ON UPDATE NO ACTION ON DELETE NO ACTION;
-- Modify "pallet_cancelled_shipment_pallet" table
ALTER TABLE "pallet_cancelled_shipment_pallet" ADD CONSTRAINT "pallet_cancelled_shipment_pallet_pallet_id" FOREIGN KEY ("pallet_id") REFERENCES "pallets" ("id") ON UPDATE NO ACTION ON DELETE CASCADE, ADD CONSTRAINT "pallet_cancelled_shipment_pallet_shipment_pallet_id" FOREIGN KEY ("shipment_pallet_id") REFERENCES "shipment_pallets" ("id") ON UPDATE NO ACTION ON DELETE CASCADE;
-- Modify "pallets" table
ALTER TABLE "pallets" ADD CONSTRAINT "pallets_consolidations_pallets" FOREIGN KEY ("consolidation_pallets") REFERENCES "consolidations" ("id") ON UPDATE NO ACTION ON DELETE NO ACTION, ADD CONSTRAINT "pallets_packagings_packaging" FOREIGN KEY ("pallet_packaging") REFERENCES "packagings" ("id") ON UPDATE NO ACTION ON DELETE SET NULL, ADD CONSTRAINT "pallets_tenants_tenant" FOREIGN KEY ("tenant_id") REFERENCES "tenants" ("id") ON UPDATE NO ACTION ON DELETE NO ACTION;
-- Modify "parcel_shop_brings" table
ALTER TABLE "parcel_shop_brings" ADD CONSTRAINT "parcel_shop_brings_parcel_shops_parcel_shop_bring" FOREIGN KEY ("parcel_shop_parcel_shop_bring") REFERENCES "parcel_shops" ("id") ON UPDATE NO ACTION ON DELETE NO ACTION;
-- Modify "parcel_shop_da_os" table
ALTER TABLE "parcel_shop_da_os" ADD CONSTRAINT "parcel_shop_da_os_parcel_shops_parcel_shop_dao" FOREIGN KEY ("parcel_shop_parcel_shop_dao") REFERENCES "parcel_shops" ("id") ON UPDATE NO ACTION ON DELETE NO ACTION;
-- Modify "parcel_shop_gl_ss" table
ALTER TABLE "parcel_shop_gl_ss" ADD CONSTRAINT "parcel_shop_gl_ss_parcel_shops_parcel_shop_gls" FOREIGN KEY ("parcel_shop_parcel_shop_gls") REFERENCES "parcel_shops" ("id") ON UPDATE NO ACTION ON DELETE NO ACTION;
-- Modify "parcel_shop_post_nords" table
ALTER TABLE "parcel_shop_post_nords" ADD CONSTRAINT "parcel_shop_post_nords_parcel_shops_parcel_shop_post_nord" FOREIGN KEY ("parcel_shop_parcel_shop_post_nord") REFERENCES "parcel_shops" ("id") ON UPDATE NO ACTION ON DELETE NO ACTION;
-- Modify "parcel_shops" table
ALTER TABLE "parcel_shops" ADD CONSTRAINT "parcel_shops_carrier_brands_carrier_brand" FOREIGN KEY ("parcel_shop_carrier_brand") REFERENCES "carrier_brands" ("id") ON UPDATE NO ACTION ON DELETE NO ACTION;
-- Modify "plan_histories" table
ALTER TABLE "plan_histories" ADD CONSTRAINT "plan_histories_change_histories_plan_history" FOREIGN KEY ("change_history_plan_history") REFERENCES "change_histories" ("id") ON UPDATE NO ACTION ON DELETE NO ACTION, ADD CONSTRAINT "plan_histories_plans_plan_history_plan" FOREIGN KEY ("plan_plan_history_plan") REFERENCES "plans" ("id") ON UPDATE NO ACTION ON DELETE NO ACTION, ADD CONSTRAINT "plan_histories_tenants_tenant" FOREIGN KEY ("tenant_id") REFERENCES "tenants" ("id") ON UPDATE NO ACTION ON DELETE NO ACTION, ADD CONSTRAINT "plan_histories_users_plan_history_user" FOREIGN KEY ("user_plan_history_user") REFERENCES "users" ("id") ON UPDATE NO ACTION ON DELETE NO ACTION;
-- Modify "print_jobs" table
ALTER TABLE "print_jobs" ADD CONSTRAINT "print_jobs_collis_colli" FOREIGN KEY ("print_job_colli") REFERENCES "collis" ("id") ON UPDATE NO ACTION ON DELETE SET NULL, ADD CONSTRAINT "print_jobs_printers_printer" FOREIGN KEY ("print_job_printer") REFERENCES "printers" ("id") ON UPDATE NO ACTION ON DELETE NO ACTION, ADD CONSTRAINT "print_jobs_shipment_parcels_shipment_parcel" FOREIGN KEY ("print_job_shipment_parcel") REFERENCES "shipment_parcels" ("id") ON UPDATE NO ACTION ON DELETE SET NULL, ADD CONSTRAINT "print_jobs_tenants_tenant" FOREIGN KEY ("tenant_id") REFERENCES "tenants" ("id") ON UPDATE NO ACTION ON DELETE NO ACTION;
-- Modify "printers" table
ALTER TABLE "printers" ADD CONSTRAINT "printers_tenants_tenant" FOREIGN KEY ("tenant_id") REFERENCES "tenants" ("id") ON UPDATE NO ACTION ON DELETE NO ACTION, ADD CONSTRAINT "printers_workstations_printer" FOREIGN KEY ("workstation_printer") REFERENCES "workstations" ("id") ON UPDATE NO ACTION ON DELETE NO ACTION;
-- Modify "product_image_product_variant" table
ALTER TABLE "product_image_product_variant" ADD CONSTRAINT "product_image_product_variant_product_image_id" FOREIGN KEY ("product_image_id") REFERENCES "product_images" ("id") ON UPDATE NO ACTION ON DELETE CASCADE, ADD CONSTRAINT "product_image_product_variant_product_variant_id" FOREIGN KEY ("product_variant_id") REFERENCES "product_variants" ("id") ON UPDATE NO ACTION ON DELETE CASCADE;
-- Modify "product_images" table
ALTER TABLE "product_images" ADD CONSTRAINT "product_images_products_product" FOREIGN KEY ("product_image_product") REFERENCES "products" ("id") ON UPDATE NO ACTION ON DELETE NO ACTION, ADD CONSTRAINT "product_images_tenants_tenant" FOREIGN KEY ("tenant_id") REFERENCES "tenants" ("id") ON UPDATE NO ACTION ON DELETE NO ACTION;
-- Modify "product_tag_products" table
ALTER TABLE "product_tag_products" ADD CONSTRAINT "product_tag_products_product_id" FOREIGN KEY ("product_id") REFERENCES "products" ("id") ON UPDATE NO ACTION ON DELETE CASCADE, ADD CONSTRAINT "product_tag_products_product_tag_id" FOREIGN KEY ("product_tag_id") REFERENCES "product_tags" ("id") ON UPDATE NO ACTION ON DELETE CASCADE;
-- Modify "product_tags" table
ALTER TABLE "product_tags" ADD CONSTRAINT "product_tags_tenants_tenant" FOREIGN KEY ("tenant_id") REFERENCES "tenants" ("id") ON UPDATE NO ACTION ON DELETE NO ACTION;
-- Modify "product_variants" table
ALTER TABLE "product_variants" ADD CONSTRAINT "product_variants_products_product_variant" FOREIGN KEY ("product_product_variant") REFERENCES "products" ("id") ON UPDATE NO ACTION ON DELETE NO ACTION, ADD CONSTRAINT "product_variants_tenants_tenant" FOREIGN KEY ("tenant_id") REFERENCES "tenants" ("id") ON UPDATE NO ACTION ON DELETE NO ACTION;
-- Modify "products" table
ALTER TABLE "products" ADD CONSTRAINT "products_tenants_tenant" FOREIGN KEY ("tenant_id") REFERENCES "tenants" ("id") ON UPDATE NO ACTION ON DELETE NO ACTION;
-- Modify "return_colli_histories" table
ALTER TABLE "return_colli_histories" ADD CONSTRAINT "return_colli_histories_change_histories_return_colli_history" FOREIGN KEY ("change_history_return_colli_history") REFERENCES "change_histories" ("id") ON UPDATE NO ACTION ON DELETE NO ACTION, ADD CONSTRAINT "return_colli_histories_return_collis_return_colli_history" FOREIGN KEY ("return_colli_return_colli_history") REFERENCES "return_collis" ("id") ON UPDATE NO ACTION ON DELETE NO ACTION, ADD CONSTRAINT "return_colli_histories_tenants_tenant" FOREIGN KEY ("tenant_id") REFERENCES "tenants" ("id") ON UPDATE NO ACTION ON DELETE NO ACTION;
-- Modify "return_collis" table
ALTER TABLE "return_collis" ADD CONSTRAINT "return_collis_addresses_recipient" FOREIGN KEY ("return_colli_recipient") REFERENCES "addresses" ("id") ON UPDATE NO ACTION ON DELETE NO ACTION, ADD CONSTRAINT "return_collis_addresses_sender" FOREIGN KEY ("return_colli_sender") REFERENCES "addresses" ("id") ON UPDATE NO ACTION ON DELETE NO ACTION, ADD CONSTRAINT "return_collis_delivery_options_delivery_option" FOREIGN KEY ("return_colli_delivery_option") REFERENCES "delivery_options" ("id") ON UPDATE NO ACTION ON DELETE SET NULL, ADD CONSTRAINT "return_collis_orders_return_colli" FOREIGN KEY ("order_return_colli") REFERENCES "orders" ("id") ON UPDATE NO ACTION ON DELETE NO ACTION, ADD CONSTRAINT "return_collis_packagings_packaging" FOREIGN KEY ("return_colli_packaging") REFERENCES "packagings" ("id") ON UPDATE NO ACTION ON DELETE SET NULL, ADD CONSTRAINT "return_collis_return_portals_return_portal" FOREIGN KEY ("return_colli_return_portal") REFERENCES "return_portals" ("id") ON UPDATE NO ACTION ON DELETE NO ACTION, ADD CONSTRAINT "return_collis_tenants_tenant" FOREIGN KEY ("tenant_id") REFERENCES "tenants" ("id") ON UPDATE NO ACTION ON DELETE NO ACTION;
-- Modify "return_order_lines" table
ALTER TABLE "return_order_lines" ADD CONSTRAINT "return_order_lines_order_lines_order_line" FOREIGN KEY ("return_order_line_order_line") REFERENCES "order_lines" ("id") ON UPDATE NO ACTION ON DELETE NO ACTION, ADD CONSTRAINT "return_order_lines_return_collis_return_order_line" FOREIGN KEY ("return_colli_return_order_line") REFERENCES "return_collis" ("id") ON UPDATE NO ACTION ON DELETE NO ACTION, ADD CONSTRAINT "return_order_lines_return_portal_claims_return_portal_claim" FOREIGN KEY ("return_order_line_return_portal_claim") REFERENCES "return_portal_claims" ("id") ON UPDATE NO ACTION ON DELETE NO ACTION, ADD CONSTRAINT "return_order_lines_tenants_tenant" FOREIGN KEY ("tenant_id") REFERENCES "tenants" ("id") ON UPDATE NO ACTION ON DELETE NO ACTION;
-- Modify "return_portal_claims" table
ALTER TABLE "return_portal_claims" ADD CONSTRAINT "return_portal_claims_locations_return_location" FOREIGN KEY ("return_portal_claim_return_location") REFERENCES "locations" ("id") ON UPDATE NO ACTION ON DELETE SET NULL, ADD CONSTRAINT "return_portal_claims_return_portals_return_portal_claim" FOREIGN KEY ("return_portal_return_portal_claim") REFERENCES "return_portals" ("id") ON UPDATE NO ACTION ON DELETE NO ACTION, ADD CONSTRAINT "return_portal_claims_tenants_tenant" FOREIGN KEY ("tenant_id") REFERENCES "tenants" ("id") ON UPDATE NO ACTION ON DELETE NO ACTION;
-- Modify "return_portal_delivery_options" table
ALTER TABLE "return_portal_delivery_options" ADD CONSTRAINT "return_portal_delivery_options_delivery_option_id" FOREIGN KEY ("delivery_option_id") REFERENCES "delivery_options" ("id") ON UPDATE NO ACTION ON DELETE CASCADE, ADD CONSTRAINT "return_portal_delivery_options_return_portal_id" FOREIGN KEY ("return_portal_id") REFERENCES "return_portals" ("id") ON UPDATE NO ACTION ON DELETE CASCADE;
-- Modify "return_portal_return_location" table
ALTER TABLE "return_portal_return_location" ADD CONSTRAINT "return_portal_return_location_location_id" FOREIGN KEY ("location_id") REFERENCES "locations" ("id") ON UPDATE NO ACTION ON DELETE CASCADE, ADD CONSTRAINT "return_portal_return_location_return_portal_id" FOREIGN KEY ("return_portal_id") REFERENCES "return_portals" ("id") ON UPDATE NO ACTION ON DELETE CASCADE;
-- Modify "return_portals" table
ALTER TABLE "return_portals" ADD CONSTRAINT "return_portals_email_templates_email_accepted" FOREIGN KEY ("return_portal_email_accepted") REFERENCES "email_templates" ("id") ON UPDATE NO ACTION ON DELETE SET NULL, ADD CONSTRAINT "return_portals_email_templates_email_confirmation_label" FOREIGN KEY ("return_portal_email_confirmation_label") REFERENCES "email_templates" ("id") ON UPDATE NO ACTION ON DELETE SET NULL, ADD CONSTRAINT "return_portals_email_templates_email_confirmation_qr_code" FOREIGN KEY ("return_portal_email_confirmation_qr_code") REFERENCES "email_templates" ("id") ON UPDATE NO ACTION ON DELETE SET NULL, ADD CONSTRAINT "return_portals_email_templates_email_received" FOREIGN KEY ("return_portal_email_received") REFERENCES "email_templates" ("id") ON UPDATE NO ACTION ON DELETE SET NULL, ADD CONSTRAINT "return_portals_tenants_tenant" FOREIGN KEY ("tenant_id") REFERENCES "tenants" ("id") ON UPDATE NO ACTION ON DELETE NO ACTION;
-- Modify "seat_group_access_rights" table
ALTER TABLE "seat_group_access_rights" ADD CONSTRAINT "seat_group_access_rights_access_rights_access_right" FOREIGN KEY ("access_right_id") REFERENCES "access_rights" ("id") ON UPDATE NO ACTION ON DELETE NO ACTION, ADD CONSTRAINT "seat_group_access_rights_seat_groups_seat_group" FOREIGN KEY ("seat_group_id") REFERENCES "seat_groups" ("id") ON UPDATE NO ACTION ON DELETE NO ACTION, ADD CONSTRAINT "seat_group_access_rights_tenants_tenant" FOREIGN KEY ("tenant_id") REFERENCES "tenants" ("id") ON UPDATE NO ACTION ON DELETE NO ACTION;
-- Modify "seat_groups" table
ALTER TABLE "seat_groups" ADD CONSTRAINT "seat_groups_tenants_tenant" FOREIGN KEY ("tenant_id") REFERENCES "tenants" ("id") ON UPDATE NO ACTION ON DELETE NO ACTION;
-- Modify "shipment_brings" table
ALTER TABLE "shipment_brings" ADD CONSTRAINT "shipment_brings_shipments_shipment_bring" FOREIGN KEY ("shipment_shipment_bring") REFERENCES "shipments" ("id") ON UPDATE NO ACTION ON DELETE NO ACTION, ADD CONSTRAINT "shipment_brings_tenants_tenant" FOREIGN KEY ("tenant_id") REFERENCES "tenants" ("id") ON UPDATE NO ACTION ON DELETE NO ACTION;
-- Modify "shipment_da_os" table
ALTER TABLE "shipment_da_os" ADD CONSTRAINT "shipment_da_os_shipments_shipment_dao" FOREIGN KEY ("shipment_shipment_dao") REFERENCES "shipments" ("id") ON UPDATE NO ACTION ON DELETE NO ACTION, ADD CONSTRAINT "shipment_da_os_tenants_tenant" FOREIGN KEY ("tenant_id") REFERENCES "tenants" ("id") ON UPDATE NO ACTION ON DELETE NO ACTION;
-- Modify "shipment_dfs" table
ALTER TABLE "shipment_dfs" ADD CONSTRAINT "shipment_dfs_shipments_shipment_df" FOREIGN KEY ("shipment_shipment_df") REFERENCES "shipments" ("id") ON UPDATE NO ACTION ON DELETE NO ACTION, ADD CONSTRAINT "shipment_dfs_tenants_tenant" FOREIGN KEY ("tenant_id") REFERENCES "tenants" ("id") ON UPDATE NO ACTION ON DELETE NO ACTION;
-- Modify "shipment_ds_vs" table
ALTER TABLE "shipment_ds_vs" ADD CONSTRAINT "shipment_ds_vs_shipments_shipment_dsv" FOREIGN KEY ("shipment_shipment_dsv") REFERENCES "shipments" ("id") ON UPDATE NO ACTION ON DELETE NO ACTION, ADD CONSTRAINT "shipment_ds_vs_tenants_tenant" FOREIGN KEY ("tenant_id") REFERENCES "tenants" ("id") ON UPDATE NO ACTION ON DELETE NO ACTION;
-- Modify "shipment_easy_posts" table
ALTER TABLE "shipment_easy_posts" ADD CONSTRAINT "shipment_easy_posts_shipments_shipment_easy_post" FOREIGN KEY ("shipment_shipment_easy_post") REFERENCES "shipments" ("id") ON UPDATE NO ACTION ON DELETE NO ACTION, ADD CONSTRAINT "shipment_easy_posts_tenants_tenant" FOREIGN KEY ("tenant_id") REFERENCES "tenants" ("id") ON UPDATE NO ACTION ON DELETE NO ACTION;
-- Modify "shipment_gl_ss" table
ALTER TABLE "shipment_gl_ss" ADD CONSTRAINT "shipment_gl_ss_shipments_shipment_gls" FOREIGN KEY ("shipment_shipment_gls") REFERENCES "shipments" ("id") ON UPDATE NO ACTION ON DELETE NO ACTION, ADD CONSTRAINT "shipment_gl_ss_tenants_tenant" FOREIGN KEY ("tenant_id") REFERENCES "tenants" ("id") ON UPDATE NO ACTION ON DELETE NO ACTION;
-- Modify "shipment_histories" table
ALTER TABLE "shipment_histories" ADD CONSTRAINT "shipment_histories_change_histories_shipment_history" FOREIGN KEY ("change_history_shipment_history") REFERENCES "change_histories" ("id") ON UPDATE NO ACTION ON DELETE NO ACTION, ADD CONSTRAINT "shipment_histories_shipments_shipment_history" FOREIGN KEY ("shipment_shipment_history") REFERENCES "shipments" ("id") ON UPDATE NO ACTION ON DELETE NO ACTION, ADD CONSTRAINT "shipment_histories_tenants_tenant" FOREIGN KEY ("tenant_id") REFERENCES "tenants" ("id") ON UPDATE NO ACTION ON DELETE NO ACTION;
-- Modify "shipment_old_consolidation" table
ALTER TABLE "shipment_old_consolidation" ADD CONSTRAINT "shipment_old_consolidation_consolidation_id" FOREIGN KEY ("consolidation_id") REFERENCES "consolidations" ("id") ON UPDATE NO ACTION ON DELETE CASCADE, ADD CONSTRAINT "shipment_old_consolidation_shipment_id" FOREIGN KEY ("shipment_id") REFERENCES "shipments" ("id") ON UPDATE NO ACTION ON DELETE CASCADE;
-- Modify "shipment_pallets" table
ALTER TABLE "shipment_pallets" ADD CONSTRAINT "shipment_pallets_pallets_shipment_pallet" FOREIGN KEY ("pallet_shipment_pallet") REFERENCES "pallets" ("id") ON UPDATE NO ACTION ON DELETE SET NULL, ADD CONSTRAINT "shipment_pallets_shipments_shipment_pallet" FOREIGN KEY ("shipment_shipment_pallet") REFERENCES "shipments" ("id") ON UPDATE NO ACTION ON DELETE NO ACTION, ADD CONSTRAINT "shipment_pallets_tenants_tenant" FOREIGN KEY ("tenant_id") REFERENCES "tenants" ("id") ON UPDATE NO ACTION ON DELETE NO ACTION;
-- Modify "shipment_parcels" table
ALTER TABLE "shipment_parcels" ADD CONSTRAINT "shipment_parcels_collis_shipment_parcel" FOREIGN KEY ("colli_shipment_parcel") REFERENCES "collis" ("id") ON UPDATE NO ACTION ON DELETE SET NULL, ADD CONSTRAINT "shipment_parcels_packagings_packaging" FOREIGN KEY ("shipment_parcel_packaging") REFERENCES "packagings" ("id") ON UPDATE NO ACTION ON DELETE SET NULL, ADD CONSTRAINT "shipment_parcels_shipments_shipment_parcel" FOREIGN KEY ("shipment_shipment_parcel") REFERENCES "shipments" ("id") ON UPDATE NO ACTION ON DELETE NO ACTION, ADD CONSTRAINT "shipment_parcels_tenants_tenant" FOREIGN KEY ("tenant_id") REFERENCES "tenants" ("id") ON UPDATE NO ACTION ON DELETE NO ACTION;
-- Modify "shipment_post_nords" table
ALTER TABLE "shipment_post_nords" ADD CONSTRAINT "shipment_post_nords_shipments_shipment_post_nord" FOREIGN KEY ("shipment_shipment_post_nord") REFERENCES "shipments" ("id") ON UPDATE NO ACTION ON DELETE NO ACTION, ADD CONSTRAINT "shipment_post_nords_tenants_tenant" FOREIGN KEY ("tenant_id") REFERENCES "tenants" ("id") ON UPDATE NO ACTION ON DELETE NO ACTION;
-- Modify "shipment_usp_ss" table
ALTER TABLE "shipment_usp_ss" ADD CONSTRAINT "shipment_usp_ss_shipments_shipment_usps" FOREIGN KEY ("shipment_shipment_usps") REFERENCES "shipments" ("id") ON UPDATE NO ACTION ON DELETE NO ACTION, ADD CONSTRAINT "shipment_usp_ss_tenants_tenant" FOREIGN KEY ("tenant_id") REFERENCES "tenants" ("id") ON UPDATE NO ACTION ON DELETE NO ACTION;
-- Modify "shipments" table
ALTER TABLE "shipments" ADD CONSTRAINT "shipments_carriers_carrier" FOREIGN KEY ("shipment_carrier") REFERENCES "carriers" ("id") ON UPDATE NO ACTION ON DELETE NO ACTION, ADD CONSTRAINT "shipments_tenants_tenant" FOREIGN KEY ("tenant_id") REFERENCES "tenants" ("id") ON UPDATE NO ACTION ON DELETE NO ACTION;
-- Modify "signup_options" table
ALTER TABLE "signup_options" ADD CONSTRAINT "signup_options_users_signup_options" FOREIGN KEY ("user_signup_options") REFERENCES "users" ("id") ON UPDATE NO ACTION ON DELETE NO ACTION;
-- Modify "system_events" table
ALTER TABLE "system_events" ADD CONSTRAINT "system_events_tenants_tenant" FOREIGN KEY ("tenant_id") REFERENCES "tenants" ("id") ON UPDATE NO ACTION ON DELETE NO ACTION;
-- Modify "tenant_connect_option_carriers" table
ALTER TABLE "tenant_connect_option_carriers" ADD CONSTRAINT "tenant_connect_option_carriers_connect_option_carrier_id" FOREIGN KEY ("connect_option_carrier_id") REFERENCES "connect_option_carriers" ("id") ON UPDATE NO ACTION ON DELETE CASCADE, ADD CONSTRAINT "tenant_connect_option_carriers_tenant_id" FOREIGN KEY ("tenant_id") REFERENCES "tenants" ("id") ON UPDATE NO ACTION ON DELETE CASCADE;
-- Modify "tenant_connect_option_platforms" table
ALTER TABLE "tenant_connect_option_platforms" ADD CONSTRAINT "tenant_connect_option_platforms_connect_option_platform_id" FOREIGN KEY ("connect_option_platform_id") REFERENCES "connect_option_platforms" ("id") ON UPDATE NO ACTION ON DELETE CASCADE, ADD CONSTRAINT "tenant_connect_option_platforms_tenant_id" FOREIGN KEY ("tenant_id") REFERENCES "tenants" ("id") ON UPDATE NO ACTION ON DELETE CASCADE;
-- Modify "tenants" table
ALTER TABLE "tenants" ADD CONSTRAINT "tenants_addresses_company_address" FOREIGN KEY ("tenant_company_address") REFERENCES "addresses" ("id") ON UPDATE NO ACTION ON DELETE SET NULL, ADD CONSTRAINT "tenants_contacts_admin_contact" FOREIGN KEY ("tenant_admin_contact") REFERENCES "contacts" ("id") ON UPDATE NO ACTION ON DELETE SET NULL, ADD CONSTRAINT "tenants_contacts_billing_contact" FOREIGN KEY ("tenant_billing_contact") REFERENCES "contacts" ("id") ON UPDATE NO ACTION ON DELETE SET NULL, ADD CONSTRAINT "tenants_languages_default_language" FOREIGN KEY ("tenant_default_language") REFERENCES "languages" ("id") ON UPDATE NO ACTION ON DELETE NO ACTION, ADD CONSTRAINT "tenants_plans_tenant" FOREIGN KEY ("plan_tenant") REFERENCES "plans" ("id") ON UPDATE NO ACTION ON DELETE NO ACTION;
-- Modify "user_seats" table
ALTER TABLE "user_seats" ADD CONSTRAINT "user_seats_tenants_tenant" FOREIGN KEY ("tenant_id") REFERENCES "tenants" ("id") ON UPDATE NO ACTION ON DELETE NO ACTION;
-- Modify "users" table
ALTER TABLE "users" ADD CONSTRAINT "users_languages_language" FOREIGN KEY ("user_language") REFERENCES "languages" ("id") ON UPDATE NO ACTION ON DELETE SET NULL, ADD CONSTRAINT "users_seat_groups_user" FOREIGN KEY ("seat_group_user") REFERENCES "seat_groups" ("id") ON UPDATE NO ACTION ON DELETE SET NULL, ADD CONSTRAINT "users_tenants_tenant" FOREIGN KEY ("tenant_id") REFERENCES "tenants" ("id") ON UPDATE NO ACTION ON DELETE NO ACTION;
-- Modify "workspace_recent_scans" table
ALTER TABLE "workspace_recent_scans" ADD CONSTRAINT "workspace_recent_scans_shipment_parcels_shipment_parcel" FOREIGN KEY ("workspace_recent_scan_shipment_parcel") REFERENCES "shipment_parcels" ("id") ON UPDATE NO ACTION ON DELETE SET NULL, ADD CONSTRAINT "workspace_recent_scans_tenants_tenant" FOREIGN KEY ("tenant_id") REFERENCES "tenants" ("id") ON UPDATE NO ACTION ON DELETE NO ACTION, ADD CONSTRAINT "workspace_recent_scans_users_user" FOREIGN KEY ("workspace_recent_scan_user") REFERENCES "users" ("id") ON UPDATE NO ACTION ON DELETE NO ACTION;
-- Modify "workstations" table
ALTER TABLE "workstations" ADD CONSTRAINT "workstations_tenants_tenant" FOREIGN KEY ("tenant_id") REFERENCES "tenants" ("id") ON UPDATE NO ACTION ON DELETE NO ACTION, ADD CONSTRAINT "workstations_users_selected_workstation" FOREIGN KEY ("user_selected_workstation") REFERENCES "users" ("id") ON UPDATE NO ACTION ON DELETE SET NULL, ADD CONSTRAINT "workstations_users_user" FOREIGN KEY ("workstation_user") REFERENCES "users" ("id") ON UPDATE NO ACTION ON DELETE SET NULL;
