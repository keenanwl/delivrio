import {FetchDeliveryOptionEditDfQuery} from "./delivery-option-edit-df.generated";

export namespace DeliveryOptionEditDFActions {
	export class Fetch {
		static readonly type = '[deliveryOptionEditDF] fetch delivery options edit DF';
	}
	export class SetID {
		static readonly type = '[deliveryOptionEditDF] set ID';
		constructor(public payload: string) {}
	}
	export class SetServices {
		static readonly type = '[deliveryOptionEditDF] set services';
		constructor(public payload: ServicesResponse[]) {}
	}
	export class SetEditDF {
		static readonly type = '[deliveryOptionEditDF] set edit DF';
		constructor(public payload: DeliveryOptionEditDFResponse) {}
	}
	export class Save {
		static readonly type = '[deliveryOptionEditDF] save';
	}
	export class Clear {
		static readonly type = '[deliveryOptionEditDF] clear';
	}
	export type DeliveryOptionEditDFResponse = NonNullable<NonNullable<FetchDeliveryOptionEditDfQuery['deliveryOptionDF']>['deliveryOption']>;
	export type ServicesResponse = NonNullable<NonNullable<NonNullable<NonNullable<FetchDeliveryOptionEditDfQuery['carrierServices']>['edges']>[0]>['node']>;
}
