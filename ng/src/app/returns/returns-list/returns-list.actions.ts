import {FetchReturnDeliveryOptionsQuery, FetchReturnsListQuery, SearchOrdersQuery} from "./returns-list.generated";
import {MutateReturnDeliveryOption, SearchResult} from "../../../generated/graphql";
import {
	ReturnPortalViewResponse
} from "../../settings/return-portal-viewer/return-portal-frame/return-portal-frame.service";
import {ItemReturn} from "../../settings/return-portal-viewer/return-portal-frame/return-portal-frame.ngxs";

export namespace ReturnsListActions {
	export class FetchReturnsList {
		static readonly type = '[ReturnsList] fetch ReturnsList';
	}
	export class SetReturnsList {
		static readonly type = '[ReturnsList] set ReturnsList';
		constructor(public payload: FetchReturns[]) {}
	}
	export class AddNewProduct {
		static readonly type = '[ReturnsList] add new product';
		constructor(public payload: string) {}
	}
	export class SearchOrders {
		static readonly type = '[ReturnsList] search orders';
		constructor(public payload: string) {}
	}
	export class SetSearchOrders {
		static readonly type = '[ReturnsList] set search orders';
		constructor(public payload: SearchOrdersResult[]) {}
	}
	export class SetSelectedOrder {
		static readonly type = '[ReturnsList] set selected order';
		constructor(public payload: SearchResult) {}
	}
	export class SetOrderLines {
		static readonly type = '[ReturnsList] set order lines';
		constructor(public payload: {view: ReturnPortalViewResponse, selected: ItemReturn[]}) {}
	}
	export class Clear {
		static readonly type = '[ReturnsList] clear';
	}
	export class ClearCreateOrder {
		static readonly type = '[ReturnsList] clear create order';
	}
	export class SearchOrderLines {
		static readonly type = '[ReturnsList] search order lines';
	}
	export class IncrementQuantity {
		static readonly type = '[ReturnsList] increment quantity';
		constructor(public payload: {orderLineID: string}) {}
	}
	export class DecrementQuantity {
		static readonly type = '[ReturnsList] decrement quantity';
		constructor(public payload: {orderLineID: string}) {}
	}
	export class SetSelectedItem {
		static readonly type = '[ReturnsList] set selected item';
		constructor(public payload: {orderLineID: string; selected: boolean}) {}
	}
	export class SetSelectedItemReason {
		static readonly type = '[ReturnsList] set selected item reason';
		constructor(public payload: {orderLineID: string; reasonID: string}) {}
	}
	export class MarkReturnColliDeleted {
		static readonly type = '[ReturnsList] mark return colli deleted';
		constructor(public payload: {returnColliID: string}) {}
	}
	export class CreateReturnOrder {
		static readonly type = '[ReturnsList] create return order';
	}
	export class CreateReturnOrderPending {
		static readonly type = '[ReturnsList] create return order pending';
	}
	export class SetNewReturnColliIDs {
		static readonly type = '[ReturnsList] set new return colli IDs';
		constructor(public payload: string[]) {}
	}
	export class SearchReturnDeliveryOptions {
		static readonly type = '[ReturnsList] search return delivery options';
	}
	export class SetSearchReturnDeliveryOptions {
		static readonly type = '[ReturnsList] set search return delivery options';
		constructor(public payload: SearchDeliveryOptionsResult) {}
	}
	export class ChangeDeliveryOption {
		static readonly type = '[ReturnsList] change delivery option';
		constructor(public payload: MutateReturnDeliveryOption) {}
	}
	export type FetchReturns = NonNullable<NonNullable<NonNullable<NonNullable<FetchReturnsListQuery['returnCollis']>['edges']>[0]>['node']>;
	export type SearchOrdersResult = NonNullable<NonNullable<SearchOrdersQuery['search']>[0]>;
	export type SearchDeliveryOptionsResult = NonNullable<FetchReturnDeliveryOptionsQuery['returnDeliveryOptions']>;
}
