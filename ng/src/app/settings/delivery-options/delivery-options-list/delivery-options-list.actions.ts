import {FetchCarrierAgreementsAndConnectionsQuery, ListDeliveryOptionsQuery} from "./delivery-options-list.generated";

export namespace DeliveryOptionsListActions {
	export class FetchDeliveryOptionsList {
		static readonly type = '[deliveryOptionsList] fetch deliveryOptionsList';
	}
	export class SetDeliveryOptionsList {
		static readonly type = '[deliveryOptionsList] set deliveryOptionsList';
		constructor(public payload: SelectDeliveryOptionsListQueryResponse[]) {}
	}
	export class CreateNewDeliveryOption {
		static readonly type = '[deliveryOptionsList] create new GLS delivery option';
		constructor(public payload: {name: string; agreementId: string; connectionID: string;}) {}
	}
	export class FetchCarrierAgreements {
		static readonly type = '[deliveryOptionsList] fetch carrier agreements';
	}
	export class SetCarrierAgreements {
		static readonly type = '[deliveryOptionsList] set carrier agreements';
		constructor(public payload: FetchCarrierAgreementsQueryResponse[]) {}
	}
	export class SetConnections {
		static readonly type = '[deliveryOptionsList] set connections';
		constructor(public payload: FetchConnectionsResponse[]) {}
	}
	export class UpdateSortOrder {
		static readonly type = '[deliveryOptionsList] update sort order';
		constructor(public payload: {deliveryOptionID: string; nextIndex: number;}) {}
	}
	export class Archive {
		static readonly type = '[deliveryOptionsList] archive';
		constructor(public payload: {deliveryOptionID: string}) {}
	}
	export class ToggleShowArchive {
		static readonly type = '[deliveryOptionsList] toggle show archive';
	}
	export class Clear {
		static readonly type = '[deliveryOptionsList] clear';
	}
	export type SelectDeliveryOptionsListQueryResponse = NonNullable<NonNullable<ListDeliveryOptionsQuery['deliveryOptionsFiltered']>[0]>;
	export type FetchCarrierAgreementsQueryResponse = NonNullable<NonNullable<NonNullable<FetchCarrierAgreementsAndConnectionsQuery['carriers']['edges']>[0]>['node']>;
	export type FetchConnectionsResponse = NonNullable<NonNullable<NonNullable<FetchCarrierAgreementsAndConnectionsQuery['connections']['edges']>[0]>['node']>;
}
