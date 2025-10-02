import {FetchPackagingListQuery} from "./packaging-list.generated";
import {CreatePackagingDfInput, CreatePackagingInput, CreatePackagingUspsInput} from "../../../../generated/graphql";

export namespace PackagingListActions {
	export class FetchPackagingList {
		static readonly type = '[PackagingList] fetch PackagingList';
	}
	export class SetPackagingList {
		static readonly type = '[PackagingList] set PackagingList';
		constructor(public payload: PackagingResponse[]) {}
	}
	export class SetCarrierBrands {
		static readonly type = '[PackagingList] set carrier brands';
		constructor(public payload: CarrierBrandResponse[]) {}
	}
	export class SetRateIndicators {
		static readonly type = '[PackagingList] set rate indicators';
		constructor(public payload: USPSRateIndicatorResponse[]) {}
	}
	export class SetProcessingCategories {
		static readonly type = '[PackagingList] set processing categories';
		constructor(public payload: USPSProcessingCategoryResponse[]) {}
	}
	export class Archive {
		static readonly type = '[PackagingList] archive';
		constructor(public payload: string) {}
	}
	export class Clear {
		static readonly type = '[PackagingList] clear';
	}
	export class Create {
		static readonly type = '[PackagingList] create';
		constructor(public payload: {
			packaging: CreatePackagingInput,
			uspsPackaging?: CreatePackagingUspsInput,
			dfPackaging?: CreatePackagingDfInput,
		}) {}
	}
	export type PackagingResponse = NonNullable<NonNullable<FetchPackagingListQuery['packagingFiltered']>[0]>;
	export type CarrierBrandResponse = NonNullable<NonNullable<NonNullable<NonNullable<FetchPackagingListQuery['carrierBrands']>['edges']>[0]>['node']>;
	export type USPSRateIndicatorResponse = NonNullable<NonNullable<NonNullable<NonNullable<FetchPackagingListQuery['packagingUSPSRateIndicators']>['edges']>[0]>['node']>;
	export type USPSProcessingCategoryResponse = NonNullable<NonNullable<NonNullable<NonNullable<FetchPackagingListQuery['packagingUSPSProcessingCategories']>['edges']>[0]>['node']>;
}
