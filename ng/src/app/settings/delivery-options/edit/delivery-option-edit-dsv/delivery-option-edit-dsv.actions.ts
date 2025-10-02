import {FetchDeliveryOptionEditDsvQuery} from "./delivery-option-edit-dsv.generated";

export namespace DeliveryOptionEditDSVActions {
	export class Fetch {
		static readonly type = '[deliveryOptionEditDSV] fetch delivery options edit Post Nord';
	}
	export class SetID {
		static readonly type = '[deliveryOptionEditDSV] set ID';
		constructor(public payload: string) {}
	}
	export class SetServices {
		static readonly type = '[deliveryOptionEditDSV] set services';
		constructor(public payload: ServicesResponse[]) {}
	}
	export class SetEditDSV {
		static readonly type = '[deliveryOptionEditDSV] set edit USPS';
		constructor(public payload: DeliveryOptionEditDSVResponse) {}
	}
	export class Save {
		static readonly type = '[deliveryOptionEditDSV] save';
	}
	export class Clear {
		static readonly type = '[deliveryOptionEditDSV] clear';
	}
	export type DeliveryOptionEditDSVResponse = NonNullable<NonNullable<FetchDeliveryOptionEditDsvQuery['deliveryOptionDSV']>['deliveryOption']>;
	export type ServicesResponse = NonNullable<NonNullable<NonNullable<NonNullable<FetchDeliveryOptionEditDsvQuery['carrierServices']>['edges']>[0]>['node']>;
}
