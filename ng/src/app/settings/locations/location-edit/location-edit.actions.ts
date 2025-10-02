import {FetchLocationQuery, LocationSearchCountriesQuery} from "./location-edit.generated";

export namespace LocationEditActions {
	export class FetchLocationEdit {
		static readonly type = '[LocationEdit] fetch LocationEdit';
	}
	export class SetLocationEdit {
		static readonly type = '[LocationEdit] set location edit';
		constructor(public payload: LocationEditResponse) {}
	}
	export class SetLocationID {
		static readonly type = '[LocationEdit] set location ID';
		constructor(public payload: string) {}
	}
	export class SetLocationTags {
		static readonly type = '[LocationEdit] set location tags';
		constructor(public payload: LocationTagResponse[]) {}
	}
	export class SetSelectedLocationTags {
		static readonly type = '[LocationEdit] set selected location tags';
		constructor(public payload: LocationTagResponse[]) {}
	}
	export class Save {
		static readonly type = '[LocationEdit] save';
	}
	export class Clear {
		static readonly type = '[LocationEdit] clear';
	}
	export class SearchCountry {
		static readonly type = '[LocationEdit] search country';
		constructor(public payload: string) {}
	}
	export class SetCountrySearch {
		static readonly type = '[LocationEdit] set country search';
		constructor(public payload: CountriesResponse[]) {}
	}
	export class ChangeCountry {
		static readonly type = '[LocationEdit] change country';
		constructor(public payload: CountriesResponse) {}
	}
	export type LocationEditResponse = NonNullable<FetchLocationQuery['location']>;
	export type LocationTagResponse = NonNullable<NonNullable<LocationEditResponse>['locationTags']>[0];

	export type CountriesResponse = NonNullable<NonNullable<NonNullable<NonNullable<LocationSearchCountriesQuery['countries']>['edges']>[0]>['node']>;
}
