import {FetchOrderViewQuery} from "./order-view.generated";
import {UpdateOrderInput} from "../../../generated/graphql";
import {GraphQLError} from "graphql";

export namespace OrderViewActions {
	export class FetchOrder {
		static readonly type = '[OrderView] fetch Order';
	}
	export class SetOrderID {
		static readonly type = '[OrderView] set order ID';
		constructor(public payload: string) {}
	}
	export class SetOrder {
		static readonly type = '[OrderView] set order';
		constructor(public payload: FetchOrderResponse) {}
	}
	export class SetTimeline {
		static readonly type = '[OrderView] set timeline';
		constructor(public payload: TimelineResponse[]) {}
	}
	export class SaveFormNew {
		static readonly type = '[OrderView] save form new';
	}
	export class SaveFormUpdate {
		static readonly type = '[OrderView] save form update';
	}
	export class SaveOrder {
		static readonly type = '[OrderView] save order';
		constructor(public payload: UpdateOrderInput) {}
	}
	export class SaveOrderSuccess {
		static readonly type = '[OrderView] save order success';
	}
	export class ResetState {
		static readonly type = '[OrderView] reset state';
	}
	export class DuplicatePackage {
		static readonly type = '[OrderView] duplicate package';
		constructor(public payload: {fromColliID: string}) {}
	}
	export class AddPackage {
		static readonly type = '[OrderView] add package';
		constructor(public payload: ColliResponse) {}
	}
	export class SetPackages {
		static readonly type = '[OrderView] set packages';
		constructor(public payload: ColliResponse[]) {}
	}
	export class SetIsDragging {
		static readonly type = '[OrderView] set is dragging';
		constructor(public payload: boolean) {}
	}
	export class FireMoveOrderLine {
		static readonly type = '[OrderView] fire move order line';
		constructor(public payload: {orderLineID: string, colliID: string}) {}
	}
	export class DeletePackage {
		static readonly type = '[OrderView] delete package';
		constructor(public payload: {colliID: string}) {}
	}
	export class SetConnections {
		static readonly type = '[OrderView] set connections';
		constructor(public payload: ConnectionResponse[]) {}
	}
	export class CreateShipments {
		static readonly type = '[OrderView] create shipments';
		constructor(public payload: {parcelIDs: string[]}) {}
	}
	export class SetShipmentErrors {
		static readonly type = '[OrderView] set shipment errors';
		constructor(public payload: {errors: readonly GraphQLError[]}) {}
	}
	export class SetLabels {
		static readonly type = '[OrderView] set labels';
		constructor(public payload: {labels: string[], allLabels: string}) {}
	}
	export class IncrementLabelViewerOffset {
		static readonly type = '[OrderView] increment label viewer offset';
	}
	export class DecrementLabelViewerOffset {
		static readonly type = '[OrderView] decrement label viewer offset';
	}
	export class Clear {
		static readonly type = '[OrderView] clear';
	}
	export class SetShipmentStatuses {
		static readonly type = '[OrderView] set shipment statuses';
		constructor(public payload: ShipmentStatusesResponse) {}
	}
	export class FetchShipmentLabels {
		static readonly type = '[OrderView] fetch shipment labels';
		constructor(public payload: string[]) {}
	}
	export class CancelShipment {
		static readonly type = '[OrderView] cancel selected shipment';
	}
	export class ClearDialogs {
		static readonly type = '[OrderView] clear dialogs';
	}
	export class FetchPackingSlips {
		static readonly type = '[OrderView] fetch packing slips';
		constructor(public payload: {parcelIDs: string[]}) {}
	}
	export class SetPackingSlips {
		static readonly type = '[OrderView] set packing slips';
		constructor(public payload: {packingSlips: string[]; allPackingSlips: string}) {}
	}
	export class CreatePackingSlipPrintJobs {
		static readonly type = '[OrderView] create packing slip print jobs';
		constructor(public payload: {parcelIDs: string[]}) {}
	}
	export class CreateLabelPrintJobs {
		static readonly type = '[OrderView] create label print jobs';
		constructor(public payload: {parcelIDs: string[]}) {}
	}
	export class PackingSlipsClearCache {
		static readonly type = '[OrderView] packing slips clear cache';
		constructor(public payload: {orderIDs: string[]}) {}
	}
	export type FetchOrderResponse = NonNullable<FetchOrderViewQuery['order']>;
	export type ColliResponse = NonNullable<NonNullable<NonNullable<FetchOrderViewQuery['order']>['colli']>[0]>;
	export type TimelineResponse = NonNullable<FetchOrderViewQuery['orderTimeline']>[0];
	export type ConnectionResponse = NonNullable<NonNullable<NonNullable<NonNullable<FetchOrderViewQuery['connections']>['edges']>[0]>['node']>;
	export type ShipmentStatusesResponse = NonNullable<FetchOrderViewQuery['orderShipments']>;
	export type ShipmentStatusResponse = NonNullable<NonNullable<NonNullable<FetchOrderViewQuery['orderShipments']>['shipmentStatuses']>[0]>;
}
