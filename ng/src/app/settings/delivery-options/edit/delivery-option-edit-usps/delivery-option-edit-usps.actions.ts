import {BaseDeliveryOptionFragment} from "../edit-common.generated";
import {SelectedEmailTemplates} from "../delivery-option-email-templates/delivery-option-email-templates.component";
import {
	FetchDeliveryOptionEditUspsQuery, UspsAdditionalServicesFragment
} from "./delivery-option-edit-usps.generated";

export namespace DeliveryOptionEditUSPSActions {
	export class Fetch {
		static readonly type = '[deliveryOptionEditUSPS] fetch delivery options edit Post Nord';
	}
	export class SetID {
		static readonly type = '[deliveryOptionEditUSPS] set ID';
		constructor(public payload: string) {}
	}
	export class SetServices {
		static readonly type = '[deliveryOptionEditUSPS] set services';
		constructor(public payload: ServicesResponse[]) {}
	}
	export class SetEditUSPS {
		static readonly type = '[deliveryOptionEditUSPS] set edit USPS';
		constructor(public payload: BaseDeliveryOptionFragment) {}
	}
	export class SetEnabledAdditionalServices {
		static readonly type = '[deliveryOptionEditUSPS] set enabled additional services';
		constructor(public payload: UspsAdditionalServicesFragment[]) {}
	}
	export class SetAdditionalServiceEnabled {
		static readonly type = '[deliveryOptionEditUSPS] set additional services is enabled';
		constructor(public payload: UspsAdditionalServicesFragment) {}
	}
	export class SetAdditionalServiceDisabled {
		static readonly type = '[deliveryOptionEditUSPS] set additional services is disabled';
		constructor(public payload: UspsAdditionalServicesFragment) {}
	}
	export class FetchAvailableAdditionalServices {
		static readonly type = '[deliveryOptionEditUSPS] fetch available additional services';
	}
	export class SetAvailableAdditionalServices {
		static readonly type = '[deliveryOptionEditUSPS] set available additional services';
		constructor(public payload: UspsAdditionalServicesFragment[]) {}
	}
	export class SetShowUnavailableAdditionalServices {
		static readonly type = '[deliveryOptionEditUSPS] set show unavailable additional services';
		constructor(public payload: boolean) {}
	}
	export class SetLocations {
		static readonly type = '[deliveryOptionEditUSPS] set locations';
		constructor(public payload: LocationResponse[]) {}
	}
	export class SetSelectedLocations {
		static readonly type = '[deliveryOptionEditUSPS] set selected locations';
		constructor(public payload: LocationResponse[]) {}
	}
	export class AddLocation {
		static readonly type = '[deliveryOptionEditUSPS] add locations';
		constructor(public payload: LocationResponse) {}
	}
	export class SetEmailTemplates {
		static readonly type = '[deliveryOptionEditUSPS] set email templates';
		constructor(public payload: EmailTemplateResponse[]) {}
	}
	export class SetSelectedEmailTemplates {
		static readonly type = '[deliveryOptionEditUSPS] set selected email templates';
		constructor(public payload: SelectedEmailTemplates) {}
	}
	export class RemoveLocation {
		static readonly type = '[deliveryOptionEditUSPS] remove locations';
		constructor(public payload: string) {}
	}
	export class Save {
		static readonly type = '[deliveryOptionEditUSPS] save';
	}
	export class Clear {
		static readonly type = '[deliveryOptionEditUSPS] clear';
	}
	export type ServicesResponse = NonNullable<NonNullable<NonNullable<NonNullable<FetchDeliveryOptionEditUspsQuery['carrierServices']>['edges']>[0]>['node']>;
	export type LocationResponse = NonNullable<NonNullable<NonNullable<NonNullable<FetchDeliveryOptionEditUspsQuery['locations']>['edges']>[0]>['node']>;
	export type EmailTemplateResponse = NonNullable<NonNullable<NonNullable<NonNullable<FetchDeliveryOptionEditUspsQuery['emailTemplates']>['edges']>[0]>['node']>;
}
