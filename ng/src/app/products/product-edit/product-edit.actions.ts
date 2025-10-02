import {
	CreateProductMutationVariables,
	FetchProductQuery, MustInventoryItemMutation, ProductsSearchCountriesQuery
} from "./product-edit.generated";
import {UpdateInventoryItemInput, UpdateProductInput, UpdateProductVariantIdInput} from "../../../generated/graphql";

export namespace ProductActions {
	export class FetchProduct {
		static readonly type = '[ProductEdit] fetch Product';
	}
	export class SetProductID {
		static readonly type = '[ProductEdit] set product ID';
		constructor(public payload: string) {}
	}
	export class SetProduct {
		static readonly type = '[ProductEdit] set product';
		constructor(public payload: FetchProductResponse) {}
	}
	export class SetVariants {
		static readonly type = '[ProductEdit] set variants';
		constructor(public payload: VariantResponse[]) {}
	}
	export class SetProductTags {
		static readonly type = '[ProductEdit] set product tags';
		constructor(public payload: FetchProductTagsResponse[]) {}
	}
	export class AddTag {
		static readonly type = '[ProductEdit] add tag';
		constructor(public payload: NonNullable<FetchProductResponse['productTags']>[0]) {}
	}
	export class RemoveTag {
		static readonly type = '[ProductEdit] remove tag';
		constructor(public payload: string) {}
	}
	export class SaveFormNew {
		static readonly type = '[ProductEdit] save form new';
		constructor(public payload: CreateProductMutationVariables) {}
	}
	export class SaveFormUpdate {
		static readonly type = '[ProductEdit] save form update';
		constructor(public payload: {input: UpdateProductInput, variants: UpdateProductVariantIdInput[]}) {}
	}
	export class CreateVariant {
		static readonly type = '[ProductEdit] create variant';
	}
	export class UploadImage {
		static readonly type = '[ProductEdit] upload image';
		constructor(public payload: string) {}
	}
	export class SetProductImageInfo {
		static readonly type = '[ProductEdit] set product image info';
		constructor(public payload: ProductImageResponse) {}
	}
	export class NextImage {
		static readonly type = '[ProductEdit] next image';
	}
	export class PreviousImage {
		static readonly type = '[ProductEdit] previous image';
	}
	export class AddImageVariant {
		static readonly type = '[ProductEdit] add image variant';
		constructor(public payload: {variantID: string}) {}
	}
	export class RemoveImageVariant {
		static readonly type = '[ProductEdit] remove image variant';
		constructor(public payload: {variantID: string}) {}
	}
	export class DeleteImage {
		static readonly type = '[ProductEdit] delete image';
		constructor(public payload: {imageID: string}) {}
	}
	export class ArchiveVariant {
		static readonly type = '[ProductEdit] archive variant';
		constructor(public payload: string) {}
	}
	export class MustInventoryItem {
		static readonly type = '[ProductEdit] must inventory item';
		constructor(public payload: string) {}
	}
	export class SetInventoryItem {
		static readonly type = '[ProductEdit] set inventory item';
		constructor(public payload?: MustInventoryItemResponse) {}
	}
	export class SetCountries {
		static readonly type = '[ProductEdit] set countries';
		constructor(public payload: CountriesResponse[]) {}
	}
	export class SearchCountries {
		static readonly type = '[ProductEdit] search countries';
		constructor(public payload: string) {}
	}
	export class ChangeCountry {
		static readonly type = '[ProductEdit] change country';
		constructor(public payload: CountriesResponse) {}
	}
	export class SaveInventoryItem {
		static readonly type = '[ProductEdit] save inventory item';
		constructor(public payload: {id: string; input: UpdateInventoryItemInput}) {}
	}
	export class CloseInventoryForm {
		static readonly type = '[ProductEdit] close inventory form';
	}
	export class ResetInventoryForm {
		static readonly type = '[ProductEdit] reset inventory form';
	}
	export class Reset {
		static readonly type = '[ProductEdit] reset';
	}
	export type FetchProductResponse = NonNullable<FetchProductQuery['product']>;
	export type VariantResponse = NonNullable<NonNullable<NonNullable<NonNullable<FetchProductQuery['productVariants']>['edges']>[0]>['node']>;
	export type FetchProductTagsResponse = NonNullable<NonNullable<NonNullable<FetchProductQuery['productTags']>['edges']>[0]>['node'];
	export type ProductImageResponse = NonNullable<NonNullable<FetchProductQuery['productForImage']>['productImage']>;
	export type MustInventoryItemResponse = NonNullable<MustInventoryItemMutation['mustInventory']>;
	export type CountriesResponse = NonNullable<NonNullable<NonNullable<NonNullable<ProductsSearchCountriesQuery['countries']>['edges']>[0]>['node']>;
}
