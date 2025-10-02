import {FetchDeliveryOptionEditBringQuery} from "./delivery-option-edit-bring.generated";

export namespace DeliveryOptionEditBringActions {
	export class Fetch {
		static readonly type = '[deliveryOptionEditBring] fetch delivery options edit Post Nord';
	}
	export class SetID {
		static readonly type = '[deliveryOptionEditBring] set ID';
		constructor(public payload: string) {}
	}
	export class SetServices {
		static readonly type = '[deliveryOptionEditBring] set services';
		constructor(public payload: ServicesResponse[]) {}
	}
	export class SetEditBring {
		static readonly type = '[deliveryOptionEditBring] set edit USPS';
		constructor(public payload: DeliveryOptionEditBringResponse) {}
	}
	export class Save {
		static readonly type = '[deliveryOptionEditBring] save';
	}
	export class Clear {
		static readonly type = '[deliveryOptionEditBring] clear';
	}
	export type DeliveryOptionEditBringResponse = NonNullable<NonNullable<FetchDeliveryOptionEditBringQuery['deliveryOptionBring']>['deliveryOption']>;
	export type ServicesResponse = NonNullable<NonNullable<NonNullable<NonNullable<FetchDeliveryOptionEditBringQuery['carrierServices']>['edges']>[0]>['node']>;
}
