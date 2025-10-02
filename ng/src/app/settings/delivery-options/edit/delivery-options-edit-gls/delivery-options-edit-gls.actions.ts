import {
	FetchDeliveryOptionsGlsEditQuery,
	UpdateDeliveryOptionGlsMutationVariables
} from "./delivery-options-edit-gls.generated";

export namespace DeliveryOptionsGLSEditActions {
	export class FetchDeliveryOptionsGLSEdit {
		static readonly type = '[deliveryOptionsGLS] fetch delivery options GLS edit';
	}
	export class SetDeliveryOptionsGLSEdit {
		static readonly type = '[deliveryOptionsGLS] set delivery options GLS edit';
		constructor(public payload: SelectDeliveryOptionsGLSEditQueryResponse) {}
	}
	export class SaveForm {
		static readonly type = '[deliveryOptionsGLS] save form';
		constructor(public payload: UpdateDeliveryOptionGlsMutationVariables) {}
	}
	export class SetSelectedOption {
		static readonly type = '[deliveryOptionsGLS] set selected option';
		constructor(public payload: string) {}
	}
	export class Clear {
		static readonly type = '[deliveryOptionsGLS] clear';
	}
	export class SetServices {
		static readonly type = '[deliveryOptionsGLS] set carrier services';
		constructor(public payload: GLSServicesResponse[]) {}
	}

	export type SelectDeliveryOptionsGLSEditQueryResponse = NonNullable<NonNullable<FetchDeliveryOptionsGlsEditQuery['deliveryOptionGLS']>['deliveryOption']>;
	export type GLSServicesResponse = NonNullable<NonNullable<NonNullable<NonNullable<FetchDeliveryOptionsGlsEditQuery['carrierServices']>['edges']>[0]>['node']>;

}
