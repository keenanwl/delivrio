import {FetchUserSeatsQuery} from "./user-seats-list.generated";

export namespace UserSeatsListActions {
	export class FetchUserSeatsList {
		static readonly type = '[UserSeatsList] fetch UserSeatsList';
	}
	export class SetUserSeatsList {
		static readonly type = '[UserSeatsList] set UserSeatsList';
		constructor(public payload: FecthUserSeats[]) {}
	}
	export class CreateNewGLSDeliveryOption {
		static readonly type = '[UserSeatsList] create new GLS delivery option';
		constructor(public payload: {name: string; agreementId: string;}) {}
	}
	export class FetchCarrierAgreements {
		static readonly type = '[UserSeatsList] fetch carrier agreements';
	}
	export class SetCarrierAgreements {
		static readonly type = '[UserSeatsList] set carrier agreements';
	}
	export type FecthUserSeats = NonNullable<NonNullable<FetchUserSeatsQuery['users']['edges']>[0]>['node'];
}
