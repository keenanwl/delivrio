import {
	CreateUserSeatMutationVariables,
	FetchUserSeatQuery,
	UpdateUserSeatMutationVariables
} from "./user-seat-edit.generated";

export namespace UserSeatsActions {
	export class FetchUserSeat {
		static readonly type = '[UserSeats] fetch UserSeat';
	}
	export class SetUserSeatID {
		static readonly type = '[UserSeats] set user seat ID';
		constructor(public payload: string) {}
	}
	export class SetUserSeat {
		static readonly type = '[UserSeats] set user seat';
		constructor(public payload: FetchUserSeatResponse) {}
	}
	export class SetSeatGroups {
		static readonly type = '[UserSeats] set seat groups';
		constructor(public payload: FetchSeatGroupsResponse[]) {}
	}
	export class SaveFormNew {
		static readonly type = '[UserSeats] save form new';
		constructor(public payload: CreateUserSeatMutationVariables) {}
	}
	export class SaveFormUpdate {
		static readonly type = '[UserSeats] save form update';
		constructor(public payload: UpdateUserSeatMutationVariables) {}
	}
	export class UpdatePassword {
		static readonly type = '[UserSeats] update password';
		constructor(public payload: {userID: string; password: string}) {}
	}
	export class Clear {
		static readonly type = '[UserSeats] clear';
	}
	export type FetchUserSeatResponse = NonNullable<FetchUserSeatQuery['user']>;
	export type FetchSeatGroupsResponse = NonNullable<NonNullable<NonNullable<FetchUserSeatQuery['seatGroups']>['edges']>[0]>['node'];
}
