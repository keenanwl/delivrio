import {ListLocationsQuery} from "./locations-list.generated";

export namespace LocationsListActions {
	export class FetchLocationsList {
		static readonly type = '[LocationsList] fetch LocationsList';
	}
	export class SetLocationsList {
		static readonly type = '[LocationsList] set LocationsList';
		constructor(public payload: LocationsResponse[]) {}
	}
	export class CreateNewLocation {
		static readonly type = '[LocationsList] create new location';
	}
	export class Reset {
		static readonly type = '[LocationsList] reset';
	}
	export type LocationsResponse = NonNullable<NonNullable<NonNullable<NonNullable<ListLocationsQuery['locations']>['edges']>[0]>['node']>;
}
