import {OrderFilter, OrderPagination, WhereChipOption} from "./orders.ngxs";
import {FetchOrdersQuery} from "./orders.generated";
import {CreateOrderInput, OrderDirection} from "../../generated/graphql";
import {SortDirection} from "@angular/material/sort";

export namespace OrdersActions {

	export class ResetWhereTop {
		static readonly type = '[Orders] reset where top';
	}
	export class SelectWhereTop {
		static readonly type = '[Orders] select where top';
		constructor(public payload: WhereChipOption) {}
	}
	export class SelectWhere {
		static readonly type = '[Orders] select where';
		constructor(public payload: WhereChipOption) {}
	}
	export class WhereChipClicked {
		static readonly type = '[Orders] where chip clicked';
		constructor(public payload: string) {}
	}
	export class WhereChipRemove {
		static readonly type = '[Orders] where chip remove';
		constructor(public payload: string) {}
	}
	export class WhereChipRemoveAll {
		static readonly type = '[Orders] where chip remove all';
	}
	export class FetchOrders {
		static readonly type = '[Orders] fetch orders';
		constructor(public payload: "previous" | "next" = "next") {}
	}
	export class FetchOptionsOrderStatus {
		static readonly type = '[Orders] fetch options order status';
	}
	export class FetchOptionsCountry {
		static readonly type = '[Orders] fetch options country';
	}
	export type FetchOrdersQueryResponse = NonNullable<NonNullable<NonNullable<FetchOrdersQuery['orders']['edges']>[0]>['node']>;
	export type OrderLineResponse = NonNullable<NonNullable<FetchOrdersQueryResponse['colli']>[0]>;
	export type ConnectionsResponse = NonNullable<NonNullable<NonNullable<FetchOrdersQuery['connections']['edges']>[0]>['node']>;
	export type LocationResponse = NonNullable<NonNullable<NonNullable<FetchOrdersQuery['locations']['edges']>[0]>['node']>;

	export class SetOrders {
		static readonly type = '[Orders] set orders';
		constructor(public payload: FetchOrdersQueryResponse[]) {}
	}
	export class SetConnections {
		static readonly type = '[Orders] set connections';
		constructor(public payload: ConnectionsResponse[]) {}
	}
	export class OrderRowsToggleRows {
		static readonly type = '[Orders] order rows toggle rows';
		constructor(public payload: FetchOrdersQueryResponse[]) {}
	}
	export class OrderRowsToggleAll {
		static readonly type = '[Orders] order rows toggle all';
	}
	export class AddOrderFilter {
		static readonly type = '[Orders] add order filter';
		constructor(public payload: OrderFilter) {}
	}
	export class SetWhereOptions {
		static readonly type = '[Orders] set where options';
		constructor(public payload: WhereChipOption[]) {}
	}
	export class LastPage {
		static readonly type = '[Orders] last page';
	}
	export class NextPage {
		static readonly type = '[Orders] next page';
	}
	export class PreviousPage {
		static readonly type = '[Orders] previous page';
	}
	export class FirstPage {
		static readonly type = '[Orders] first page';
	}
	export class SetPagination {
		static readonly type = '[Orders] set pagination';
		constructor(public payload: OrderPagination) {}
	}
	export class ResetPagination {
		static readonly type = '[Orders] reset pagination';
	}
	export class ResetState {
		static readonly type = '[Orders] reset state';
	}
	export class CreateNewOrder {
		static readonly type = '[Orders] create new order';
		constructor(public payload: {input: CreateOrderInput}) {}
	}
	export class SetSenderLocations {
		static readonly type = '[Orders] set sender locations';
		constructor(public payload: LocationResponse[]) {}
	}
	export class ShowHideColumn {
		static readonly type = '[Orders] show hide column';
		constructor(public payload: string[]) {}
	}
	export class ChangeSortBy {
		static readonly type = '[Orders] change sort by';
		constructor(public payload: OrderDirection) {}
	}
	export class BulkUpdatePackaging {
		static readonly type = '[Orders] bulk update packaging';
		constructor(public payload: string | null) {}
	}
	export class BulkFetchPackingSlips {
		static readonly type = '[Orders] bulk fetch packing slips';
	}
	export class SetPackingSlips {
		static readonly type = '[Orders] set packing slips';
		constructor(public payload: {packingSlips: string[]; allPackingSlips: string}) {}
	}
	export class CreatePackingSlipPrintJobs {
		static readonly type = '[Orders] create packing slip print jobs';
	}
	export class LocalFilterWhere {
		static readonly type = '[Orders] local filter where';
		constructor(public payload: string) {}
	}
}
