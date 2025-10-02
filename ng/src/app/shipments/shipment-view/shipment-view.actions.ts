import {FetchShipmentQuery} from "./shipment-view.generated";

export namespace ShipmentViewActions {
	export class FetchShipment {
		static readonly type = '[ShipmentView] fetch shipment';
	}
	export class SetShipment {
		static readonly type = '[ShipmentView] set shipment';
		constructor(public payload: FetchShipmentResponse) {}
	}
	export class SetShipmentID {
		static readonly type = '[ShipmentView] set shipment ID';
		constructor(public payload: string) {}
	}
	export class DebugUpdateLabelIDs {
		static readonly type = '[ShipmentView] debug update label ids';
		constructor(public payload: {parcelID: string; itemID: string}) {}
	}
	export class CancelShipment {
		static readonly type = '[ShipmentView] cancel shipment';
	}
	export class CancelCancelSync {
		static readonly type = '[ShipmentView] cancel cancel sync';
		constructor(public payload: string) {}
	}
	export class CancelFulfillmentSync {
		static readonly type = '[ShipmentView] cancel fulfillment sync';
		constructor(public payload: string) {}
	}
	export class Clear {
		static readonly type = '[ShipmentView] clear';
	}
	export type FetchShipmentResponse = NonNullable<FetchShipmentQuery['shipment']>;
	export type ShipmentParcelResponse = NonNullable<NonNullable<FetchShipmentQuery['shipment']>['shipmentParcel']>[0];
	/*export type TimelineResponse = NonNullable<FetchOrderViewQuery['orderTimeline']>[0];*/
}
