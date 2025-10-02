import {
	CreateSeatGroupMutationVariables,
	FetchSeatGroupQuery,
	ReplaceSeatGroupMutationVariables
} from "./seat-group-edit.generated";

export namespace SeatGroupActions {
	export class FetchSeatGroup {
		static readonly type = '[SeatGroup] fetch SeatGroup';
	}
	export class SetSeatGroupID {
		static readonly type = '[SeatGroup] set user seat ID';
		constructor(public payload: string) {}
	}
	export class SetSeatGroup {
		static readonly type = '[SeatGroup] set user seat';
		constructor(public payload: FetchSeatGroupResponse) {}
	}
	export class SetAccessRights {
		static readonly type = '[SeatGroup] set access rights';
		constructor(public payload: FetchAccessRightsResponse[]) {}
	}
	export class SaveFormNew {
		static readonly type = '[SeatGroup] save form new';
		constructor(public payload: CreateSeatGroupMutationVariables) {}
	}
	export class SaveFormEdit {
		static readonly type = '[SeatGroup] save form edit';
		constructor(public payload: ReplaceSeatGroupMutationVariables) {}
	}
	export class Clear {
		static readonly type = '[SeatGroup] clear';
	}
	export type FetchSeatGroupResponse = NonNullable<FetchSeatGroupQuery['seatGroup']>;
	export type FetchAccessRightsResponse = NonNullable<NonNullable<NonNullable<FetchSeatGroupQuery['accessRights']>['edges']>[0]>["node"];
}
