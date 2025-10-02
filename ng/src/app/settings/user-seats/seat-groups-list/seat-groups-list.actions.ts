import {FetchSeatGroupsQuery} from "./seat-groups-list.generated";

export namespace SeatGroupsListActions {
	export class FetchSeatGroupsList {
		static readonly type = '[SeatGroupsList] fetch SeatGroupsList';
	}
	export class SetSeatGroupsList {
		static readonly type = '[SeatGroupsList] set SeatGroupsList';
		constructor(public payload: FetchSeatGroups[]) {}
	}
	export type FetchSeatGroups = NonNullable<NonNullable<FetchSeatGroupsQuery['seatGroups']['edges']>[0]>['node'];
}
