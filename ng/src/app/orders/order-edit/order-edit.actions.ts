import {
	FetchAvailableClickCollectLocationsQuery,
	FetchColliQuery,
	FetchDeliveryOptionsQuery,
	OrdersSearchCountriesQuery,
	SearchProductsQuery
} from "./order-edit.generated";
import {PriceEdited, RowIndex, UnitsEdited} from "../../shared/order-lines/order-lines.component";
import {SelectPackagingActions} from "../../shared/select-packaging/select-packaging.actions";

export namespace OrderEditActions {

	import PackagingResponse = SelectPackagingActions.PackagingResponse;

	export class FetchPackage {
		static readonly type = '[Order] fetch Order';
	}
	export class SetColliID {
		static readonly type = '[Order] set colli ID';
		constructor(public payload: string) {}
	}
	export class SetOrder {
		static readonly type = '[Order] set order';
		constructor(public payload: FetchOrderResponse) {}
	}
	export class SetOrderID {
		static readonly type = '[Order] set order ID';
		constructor(public payload: string) {}
	}
	export class SearchProducts {
		static readonly type = '[Order] search products';
		constructor(public payload: string) {}
	}
	export class SetProducts {
		static readonly type = '[Order] set products';
		constructor(public payload: SearchProductsResponse[]) {}
	}
	export class AddProduct {
		static readonly type = '[Order] add product';
		constructor(public payload: SearchProductsResponse) {}
	}
	export class ProductIncreaseUnits {
		static readonly type = '[Order] product increase units';
		constructor(public payload: string) {}
	}
	export class ProductDecreaseUnits {
		static readonly type = '[Order] product decrease units';
		constructor(public payload: string) {}
	}
	export class RemoveOrderLine {
		static readonly type = '[Order] delete order line';
		constructor(public payload: RowIndex) {}
	}
	export class SaveFormNew {
		static readonly type = '[Order] save form new';
	}
	export class SaveFormUpdate {
		static readonly type = '[Order] save form update';
	}
	export class ResetState {
		static readonly type = '[Order] reset state';
	}
	export class FetchDeliveryOptions {
		static readonly type = '[Order] fetch delivery options';
	}
	export class SetDeliveryOptions {
		static readonly type = '[Order] set delivery options';
		constructor(public payload: FetchDeliveryOptionsResponse[]) {}
	}
	export class SelectDeliveryOption {
		static readonly type = '[Order] select delivery option';
		constructor(public payload: {id: string; clickCollect: boolean}) {}
	}
	export class SetConnectionList {
		static readonly type = '[Order] set connection list';
		constructor(public payload: ConnectionListResponse[]) {}
	}
	export class RowEditedPrice {
		static readonly type = '[Order] row edited price';
		constructor(public payload: PriceEdited) {}
	}
	export class RowEditedUnits {
		static readonly type = '[Order] row edited units';
		constructor(public payload: UnitsEdited) {}
	}
	export class StopLoading {
		static readonly type = '[Order] stop loading';
	}
	export class SearchCountry {
		static readonly type = '[Order] search country';
		constructor(public payload: string) {}
	}
	export class SetCountrySearch {
		static readonly type = '[Order] set country search';
		constructor(public payload: CountriesResponse[]) {}
	}
	export class ChangeCountrySender {
		static readonly type = '[Order] change country sender';
		constructor(public payload: CountriesResponse) {}
	}
	export class ChangeCountryRecipient {
		static readonly type = '[Order] change country recipient';
		constructor(public payload: CountriesResponse) {}
	}
	export class SetDeliveryPoint {
		static readonly type = '[Order] set delivery point';
		constructor(public payload: DeliveryPointResponse) {}
	}
	export class SearchDeliveryPoints {
		static readonly type = '[Order] search delivery points';
	}
	export class SetDeliveryPointsSearch {
		static readonly type = '[Order] set delivery points search';
		constructor(public payload: DeliveryPointResponse[]) {}
	}
	export class SetClickCollectLocation {
		static readonly type = '[Order] set click collect location';
		constructor(public payload: ClickCollectResponse) {}
	}
	export class FetchAvailableClickCollectLocations {
		static readonly type = '[Order] fetch click collect locations';
	}
	export class SetSelectedPackaging {
		static readonly type = '[Order] set selected packaging';
		constructor(public payload: PackagingResponse | null) {}
	}
	export class SetAvailableClickCollectLocations {
		static readonly type = '[Order] set click collect locations';
		constructor(public payload: AvailableClickCollectResponse[]) {}
	}

	export type FetchOrderResponse = NonNullable<FetchColliQuery['colli']>;
	export type FetchDeliveryOptionsResponse = NonNullable<NonNullable<FetchDeliveryOptionsQuery['deliveryOptionsList']>[0]>;
	export type OrderLineResponse = NonNullable<NonNullable<FetchColliQuery['colli']>['orderLines']>[0];
	export type SearchProductsResponse = NonNullable<NonNullable<NonNullable<SearchProductsQuery['productVariants']>['edges']>[0]>['node'];

	export type ConnectionListResponse = NonNullable<NonNullable<NonNullable<FetchColliQuery['connections']>['edges']>[0]>['node'];

	export type CountriesResponse = NonNullable<NonNullable<NonNullable<NonNullable<OrdersSearchCountriesQuery['countries']>['edges']>[0]>['node']>;
	export type DeliveryPointResponse = NonNullable<FetchColliQuery['deliveryPoint']>;

	export type ClickCollectResponse = NonNullable<FetchColliQuery['clickCollectLocation']>;
	export type AvailableClickCollectResponse = NonNullable<NonNullable<NonNullable<NonNullable<FetchAvailableClickCollectLocationsQuery['locations']>['edges']>[0]>['node']>;
}
