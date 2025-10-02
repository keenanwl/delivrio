import {ConsolidationSearchCountriesQuery, FetchConsolidationQuery} from "./consolidation-edit.generated";
import {AddressInfoFragment} from "../../orders/order-edit/order-edit.generated";
import {GraphQLError} from "graphql/index";

export namespace ConsolidationEditActions {
	export class FetchConsolidationEdit {
		static readonly type = '[ConsolidationEdit] fetch ConsolidationEdit';
	}
	export class SetConsolidationEdit {
		static readonly type = '[ConsolidationEdit] set ConsolidationEdit';
		constructor(public payload: ConsolidationResponse) {}
	}
	export class SetConsolidationID {
		static readonly type = '[ConsolidationEdit] set ConsolidationEdit ID';
		constructor(public payload: string) {}
	}
	export class SetOrders {
		static readonly type = '[ConsolidationEdit] set orders';
		constructor(public payload: OrderResponse[]) {}
	}
	export class SetPallets {
		static readonly type = '[ConsolidationEdit] set pallets';
		constructor(public payload: PalletResponse[]) {}
	}
	export class EditPallet {
		static readonly type = '[ConsolidationEdit] edit pallet';
		constructor(public payload: PalletResponse | undefined) {}
	}
	export class UpdatePallet {
		static readonly type = '[ConsolidationEdit] update pallet';
		constructor(public payload: {publicID: string; description: string; packagingID: string;}) {}
	}
	export class SearchOrders {
		static readonly type = '[ConsolidationEdit] search orders';
		constructor(public payload: string) {}
	}
	export class SetSearchOrders {
		static readonly type = '[ConsolidationEdit] set search orders';
		constructor(public payload: OrderResponse[]) {}
	}
	export class MoveOrder {
		static readonly type = '[ConsolidationEdit] add order to pallet';
		constructor(public payload: {item: OrderResponse; destination: ListContainer}) {}
	}
	export class RemoveOrder {
		static readonly type = '[ConsolidationEdit] remove order';
		constructor(public payload: string) {}
	}
	export class SetDeliveryOptions {
		static readonly type = '[ConsolidationEdit] set delivery options';
		constructor(public payload: DeliveryOptionItem[]) {}
	}
	export class SetPackagings {
		static readonly type = '[ConsolidationEdit] set packagings';
		constructor(public payload: PackagingItem[]) {}
	}
	export class SetRecipient {
		static readonly type = '[ConsolidationEdit] set recipient';
		constructor(public payload: AddressInfoFragment) {}
	}
	export class SetSender {
		static readonly type = '[ConsolidationEdit] set sender';
		constructor(public payload: AddressInfoFragment) {}
	}
	export class SearchCountries {
		static readonly type = '[ConsolidationEdit] search countries';
		constructor(public payload: string) {}
	}
	export class SetSearchCountries {
		static readonly type = '[ConsolidationEdit] set search countries';
		constructor(public payload: CountryResult[]) {}
	}
	export class CreateShipment {
		static readonly type = '[ConsolidationEdit] create shipment';
		constructor(public payload: {prebook: boolean}) {}
	}
	export class Clear {
		static readonly type = '[ConsolidationEdit] clear';
	}
	export class Save {
		static readonly type = '[ConsolidationEdit] create';
	}

	export class SetShipmentInfo {
		static readonly type = '[ConsolidationEdit] set shipment info';
		constructor(public payload: ConsolidationShipment) {}
	}
	export class SetLabels {
		static readonly type = '[ConsolidationEdit] set labels';
		constructor(public payload: {labelsPDF: string[], allLabels: string}) {}
	}
	export class SetShipmentErrors {
		static readonly type = '[ConsolidationEdit] set shipment errors';
		constructor(public payload: {errors: readonly GraphQLError[]}) {}
	}
	export class IncrementLabelViewerOffset {
		static readonly type = '[ConsolidationEdit] increment label viewer offset';
	}
	export class DecrementLabelViewerOffset {
		static readonly type = '[ConsolidationEdit] decrement label viewer offset';
	}

	export enum ListContainerType {
		SEARCH,
		ORDERS,
		PALLET,
	}

	export type ListContainer = {
		type: ListContainerType;
		index: number;
	}

	export type ConsolidationResponse = NonNullable<FetchConsolidationQuery['consolidation']>;
	export type PalletResponse = NonNullable<NonNullable<FetchConsolidationQuery['consolidation']>['pallets']>[0];
	export type OrderResponse = NonNullable<NonNullable<NonNullable<FetchConsolidationQuery['consolidation']>['orders']>[0]>;
	export type DeliveryOptionItem = NonNullable<NonNullable<NonNullable<NonNullable<FetchConsolidationQuery['deliveryOptions']>['edges']>[0]>['node']>;
	export type PackagingItem = NonNullable<NonNullable<NonNullable<NonNullable<FetchConsolidationQuery['packagings']>['edges']>[0]>['node']>;
	export type CountryResult = NonNullable<NonNullable<NonNullable<NonNullable<ConsolidationSearchCountriesQuery['countries']>['edges']>[0]>['node']>;
	export type ConsolidationShipment = NonNullable<FetchConsolidationQuery['consolidationShipments']>;
}
