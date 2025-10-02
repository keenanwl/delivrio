import {
	AvailableAdditionalServicesPostNordQuery,
	FetchDeliveryOptionEditPostNordQuery
} from "./delivery-option-edit-post-nord.generated";
import {SelectedEmailTemplates} from "../delivery-option-email-templates/delivery-option-email-templates.component";

export namespace DeliveryOptionEditPostNordActions {
	export class Fetch {
		static readonly type = '[deliveryOptionEditPostNord] fetch delivery options edit Post Nord';
	}
	export class SetID {
		static readonly type = '[deliveryOptionEditPostNord] set ID';
		constructor(public payload: string) {}
	}
	export class SetServices {
		static readonly type = '[deliveryOptionEditPostNord] set services';
		constructor(public payload: ServicesResponse[]) {}
	}
	export class SetEditPostNord {
		static readonly type = '[deliveryOptionEditPostNord] set edit PostNord';
		constructor(public payload: SelectDeliveryOptionEditPostNordQueryResponse) {}
	}
	export class SetEnabledAdditionalServices {
		static readonly type = '[deliveryOptionEditPostNord] set enabled additional services';
		constructor(public payload: AddedAdditionalServiceResponse[]) {}
	}
	export class SetAdditionalServiceEnabled {
		static readonly type = '[deliveryOptionEditPostNord] set additional services is enabled';
		constructor(public payload: {internalID: string, checked: boolean}) {}
	}
	export class FetchAvailableAdditionalServices {
		static readonly type = '[deliveryOptionEditPostNord] fetch available additional services';
	}
	export class SetAvailableAdditionalServices {
		static readonly type = '[deliveryOptionEditPostNord] set available additional services';
		constructor(public payload: AvailableAdditionalServiceResponse[]) {}
	}
	export class SetShowUnavailableAdditionalServices {
		static readonly type = '[deliveryOptionEditPostNord] set show unavailable additional services';
		constructor(public payload: boolean) {}
	}
	export class SetLocations {
		static readonly type = '[deliveryOptionEditPostNord] set locations';
		constructor(public payload: LocationResponse[]) {}
	}
	export class SetSelectedLocations {
		static readonly type = '[deliveryOptionEditPostNord] set selected locations';
		constructor(public payload: LocationResponse[]) {}
	}
	export class AddLocation {
		static readonly type = '[deliveryOptionEditPostNord] add locations';
		constructor(public payload: LocationResponse) {}
	}
	export class SetEmailTemplates {
		static readonly type = '[deliveryOptionEditPostNord] set email templates';
		constructor(public payload: EmailTemplateResponse[]) {}
	}
	export class SetSelectedEmailTemplates {
		static readonly type = '[deliveryOptionEditPostNord] set selected email templates';
		constructor(public payload: SelectedEmailTemplates) {}
	}
	export class RemoveLocation {
		static readonly type = '[deliveryOptionEditPostNord] remove locations';
		constructor(public payload: string) {}
	}
	export class Save {
		static readonly type = '[deliveryOptionEditPostNord] save';
	}
	export class Clear {
		static readonly type = '[deliveryOptionEditPostNord] clear';
	}
	export type SelectDeliveryOptionEditPostNordQueryResponse = NonNullable<NonNullable<FetchDeliveryOptionEditPostNordQuery['deliveryOptionPostNord']>['deliveryOption']>;
	export type ServicesResponse = NonNullable<NonNullable<NonNullable<NonNullable<FetchDeliveryOptionEditPostNordQuery['carrierServices']>['edges']>[0]>['node']>;
	export type AddedAdditionalServiceResponse = NonNullable<NonNullable<NonNullable<NonNullable<FetchDeliveryOptionEditPostNordQuery['carrierAdditionalServicePostNords']>['edges']>[0]>['node']>;
	export type AvailableAdditionalServiceResponse = NonNullable<AvailableAdditionalServicesPostNordQuery>['availableAdditionalServicesPostNord'][0];
	export type LocationResponse = NonNullable<NonNullable<NonNullable<NonNullable<FetchDeliveryOptionEditPostNordQuery['locations']>['edges']>[0]>['node']>;
	export type EmailTemplateResponse = NonNullable<NonNullable<NonNullable<NonNullable<FetchDeliveryOptionEditPostNordQuery['emailTemplates']>['edges']>[0]>['node']>;
}
