import {FetchDeliveryOptionEditDaoQuery} from "./delivery-option-edit-dao.generated";

export namespace DeliveryOptionEditDAOActions {
	export class Fetch {
		static readonly type = '[deliveryOptionEditDAO] fetch delivery options edit Post Nord';
	}
	export class SetID {
		static readonly type = '[deliveryOptionEditDAO] set ID';
		constructor(public payload: string) {}
	}
	export class SetServices {
		static readonly type = '[deliveryOptionEditDAO] set services';
		constructor(public payload: ServicesResponse[]) {}
	}
	export class SetEditDAO {
		static readonly type = '[deliveryOptionEditDAO] set edit USPS';
		constructor(public payload: DeliveryOptionEditDAOResponse) {}
	}
	export class Save {
		static readonly type = '[deliveryOptionEditDAO] save';
	}
	export class Clear {
		static readonly type = '[deliveryOptionEditDAO] clear';
	}
	export type DeliveryOptionEditDAOResponse = NonNullable<NonNullable<FetchDeliveryOptionEditDaoQuery['deliveryOptionDAO']>['deliveryOption']>;
	export type ServicesResponse = NonNullable<NonNullable<NonNullable<NonNullable<FetchDeliveryOptionEditDaoQuery['carrierServices']>['edges']>[0]>['node']>;
}
