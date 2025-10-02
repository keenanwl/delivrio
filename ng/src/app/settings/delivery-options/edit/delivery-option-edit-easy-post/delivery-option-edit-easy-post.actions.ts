import {
	FetchAdditionalServiceEasyPostQuery,
	FetchDeliveryOptionEditEasyPostQuery
} from "./delivery-option-edit-easy-post.generated";

export namespace DeliveryOptionEditEasyPostActions {
	export class Fetch {
		static readonly type = '[deliveryOptionEditEasyPost] fetch delivery options edit Post Nord';
	}
	export class SetID {
		static readonly type = '[deliveryOptionEditEasyPost] set ID';
		constructor(public payload: string) {}
	}
	export class SetServices {
		static readonly type = '[deliveryOptionEditEasyPost] set services';
		constructor(public payload: ServicesResponse[]) {}
	}
	export class SetEditEasyPost {
		static readonly type = '[deliveryOptionEditEasyPost] set edit USPS';
		constructor(public payload: DeliveryOptionEditEasyPostResponse) {}
	}
	export class Save {
		static readonly type = '[deliveryOptionEditEasyPost] save';
	}
	export class Clear {
		static readonly type = '[deliveryOptionEditEasyPost] clear';
	}
	export class FetchAdditionalServices {
		static readonly type = '[deliveryOptionEditEasyPost] fetch additional services';
		constructor(public payload: string) {}
	}
	export class SetAdditionalServices {
		static readonly type = '[deliveryOptionEditEasyPost] set additional services';
		constructor(public payload: AdditionalServiceResponse[]) {}
	}
	export class ToggleAdditionalService {
		static readonly type = '[deliveryOptionEditEasyPost] toggle additional services';
		constructor(public payload: {id: string, isAdd: boolean}) {}
	}
	export class SetSelectedAdditionalService {
		static readonly type = '[deliveryOptionEditEasyPost] set selected additional services';
		constructor(public payload: string[]) {}
	}
	export type DeliveryOptionEditEasyPostResponse = NonNullable<NonNullable<FetchDeliveryOptionEditEasyPostQuery['deliveryOptionEasyPost']>['deliveryOption']>;
	export type ServicesResponse = NonNullable<NonNullable<NonNullable<NonNullable<FetchDeliveryOptionEditEasyPostQuery['carrierServices']>['edges']>[0]>['node']>;
	export type AdditionalServiceResponse = NonNullable<NonNullable<NonNullable<NonNullable<FetchAdditionalServiceEasyPostQuery['carrierAdditionalServiceEasyPosts']>['edges']>[0]>['node']>;
}
