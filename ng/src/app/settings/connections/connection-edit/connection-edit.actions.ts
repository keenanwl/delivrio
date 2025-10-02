import {
	ConnectionBrandsQuery, ConnectionDeliveryOptionFragment,
	CreateConnectionShopifyMutationVariables, FetchConnectionQuery, UpdateConnectionShopifyMutationVariables
} from "./connection-edit.generated";
import {SelectedLocations} from "./locations-selector/locations-selector.component";

export namespace ConnectionEditActions {
	export class FetchConnectionEdit {
		static readonly type = '[ConnectionEdit] fetch ConnectionEdit';
	}
	export class Clear {
		static readonly type = '[ConnectionEdit] clear';
	}
	export class SetConnectionEdit {
		static readonly type = '[ConnectionEdit] set connection edit';
		constructor(public payload: SelectConnectionsEditQueryResponse) {}
	}
	export class SetConnectionID {
		static readonly type = '[ConnectionEdit] set connection ID';
		constructor(public payload: string) {}
	}
	export class FetchConnectionBrands {
		static readonly type = '[ConnectionEdit] fetch connection brands';
	}
	export class SetConnectionBrands {
		static readonly type = '[ConnectionEdit] set connection brands';
		constructor(public payload: SelectConnectionBrandsQueryResponse[]) {}
	}
	export class SaveForm {
		static readonly type = '[ConnectionEdit] save form create';
		constructor(public payload: CreateConnectionShopifyMutationVariables) {}
	}
	export class SaveFormUpdate {
		static readonly type = '[ConnectionEdit] save form update';
		constructor(public payload: UpdateConnectionShopifyMutationVariables) {}
	}
	export class SetLocations {
		static readonly type = '[ConnectionEdit] set locations';
		constructor(public payload: LocationsResponse[]) {}
	}
	export class UpdateLocations {
		static readonly type = '[ConnectionEdit] update locations';
		constructor(public payload: SelectedLocations) {}
	}
	export class SetDeliveryOptions {
		static readonly type = '[ConnectionEdit] set delivery options';
		constructor(public payload: ConnectionDeliveryOptionFragment[]) {}
	}
	export class SetDocs {
		static readonly type = '[ConnectionEdit] set docs';
		constructor(public payload: DocsResponse[]) {}
	}
	export class SetCurrencies {
		static readonly type = '[ConnectionEdit] set currencies';
		constructor(public payload: CurrencyResponse[]) {}
	}
	export class AddFilterTag {
		static readonly type = '[ConnectionEdit] add filter tag';
		constructor(public payload: string) {}
	}
	export class RemoveFilterTag {
		static readonly type = '[ConnectionEdit] remove filter tag';
		constructor(public payload: string) {}
	}
	export type SelectConnectionsEditQueryResponse = NonNullable<FetchConnectionQuery['connection']>;
	export type SelectConnectionBrandsQueryResponse = NonNullable<NonNullable<ConnectionBrandsQuery['connectionBrands']['edges']>[0]>;
	export type LocationsResponse = NonNullable<NonNullable<NonNullable<FetchConnectionQuery['locations']['edges']>[0]>['node']>;
	export type DocsResponse = NonNullable<NonNullable<NonNullable<FetchConnectionQuery['documents']['edges']>[0]>['node']>;
	export type CurrencyResponse = NonNullable<NonNullable<NonNullable<FetchConnectionQuery['currencies']['edges']>[0]>['node']>;
}
