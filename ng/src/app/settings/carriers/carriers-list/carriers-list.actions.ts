import {CarrierBrandsQuery, ListCarriersQuery} from "./carriers-list.generated";

export namespace CarriersListActions {
	export class FetchCarriersList {
		static readonly type = '[CarriersList] fetch CarriersList';
	}
	export class SetCarriersList {
		static readonly type = '[CarriersList] set CarriersList';
		constructor(public payload: SelectCarriersListQueryResponse[]) {}
	}
	export class FetchCarrierBrands {
		static readonly type = '[CarrierList] fetch Carrier brands';
	}
	export class SetCarrierBrands {
		static readonly type = '[CarrierList] set Carrier brands';
		constructor(public payload: SelectCarrierBrandsQueryResponse[]) {}
	}
	export class CreateNewAgreement {
		static readonly type = '[CarrierList] create new agreement';
		constructor(public payload: {name: string, carrierBrandID: string}) {}
	}
	export class Clear {
		static readonly type = '[CarrierList] clear';
	}
	export type SelectCarriersListQueryResponse = NonNullable<NonNullable<NonNullable<ListCarriersQuery['carriers']['edges']>[0]>['node']>;
	export type SelectCarrierBrandsQueryResponse = NonNullable<NonNullable<NonNullable<CarrierBrandsQuery['carrierBrands']['edges']>[0]>['node']>;
}
