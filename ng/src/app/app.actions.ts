import {Breakpoint} from "./app.ngxs";
import {
	FetchSelectableWorkstationsQuery,
	FetchSelectedWorkstationQuery,
	FetchUserQuery,
	SearchQuery
} from "./app.generated";
import {UserPickupDay} from "../generated/graphql";

export interface GQLFieldError {
	message: string;
	path: Array<string | number>;
}

export namespace AppActions {

	export class FetchAppConfig {
		static readonly type = '[Login] fetch app config';
	}
	export class AppChangeRoute {
		static readonly type = '[App] change route';
		constructor(public payload: {path: string; queryParams: {[key: string]: any}}) {}
	}
	// Differs from above by only changing URL
	// in store. Then calls the .back() to maintain
	// browser navigation. Does not get listened to.
	export class AppSetRoute {
		static readonly type = '[App] set route';
		constructor(public payload: string) {}
	}
	export class AppGoBack {
		static readonly type = '[App] go back';
	}
	export type MyIds = {
		my_pulid: string;
		my_tenant_pulid: string;
	}
	export class UpdateBreakpoint {
		static readonly type = '[App] update breakpoint';
		constructor(public payload: Breakpoint) {}
	}
	export class Logout {
		static readonly type = '[App] logout';
	}
	export class ShowGlobalSnackbar {
		static readonly type = '[App] show global snackbar';
		constructor(public payload: string) {}
	}

	export class FetchLoggedInUser {
		static readonly type = '[App] fetch logged in user';
	}
	export type FetchLoggedInUserQueryResponse = NonNullable<FetchUserQuery['user']>;
	export class SetLoggedInUser {
		static readonly type = '[App] set logged in user';
		constructor(public payload: FetchLoggedInUserQueryResponse, public myIDs: MyIds) {}
	}
	export class Search {
		static readonly type = '[App] search';
		constructor(public payload: string) {}
	}
	export class SetSearchResults {
		static readonly type = '[App] set search results';
		constructor(public payload: SearchResult[]) {}
	}
	export class ClearSearchResults {
		static readonly type = '[App] clear search results';
	}
	export class FetchSelectedWorkstation {
		static readonly type = '[App] fetch selected workstation';
	}
	export class SetSelectedWorkstation {
		static readonly type = '[App] set selected workstation';
		constructor(public payload: {workstation: Workstation | undefined | null; printJobs: PrintJob[]; limitExceeded: boolean;}) {}
	}
	export class FetchSelectableWorkstations {
		static readonly type = '[App] fetch selectable workstations';
	}
	export class SetSelectableWorkstations {
		static readonly type = '[App] set selectable workstations';
		constructor(public payload: Workstation[]) {}
	}
	export class SetSelectedPickup {
		static readonly type = '[App] set selected pickup';
		constructor(public payload: UserPickupDay) {}
	}
	export class SaveSelectedWorkstation {
		static readonly type = '[App] save selectable workstation';
		constructor(public payload: {workstationID: string; pickupDay: UserPickupDay}) {}
	}
	export class SetBuildInfo {
		static readonly type = '[App] set build info';
		constructor(public payload: BuildInfoResponse) {}
	}

	export type SearchResult = NonNullable<NonNullable<SearchQuery['search']>[0]>;
	export type Workstation = NonNullable<NonNullable<NonNullable<NonNullable<FetchSelectableWorkstationsQuery['workstations']>['edges']>[0]>['node']>;
	export type PrintJob = NonNullable<NonNullable<NonNullable<FetchSelectedWorkstationQuery['selectedWorkstation']>['jobs']>[0]>;
	export type BuildInfoResponse = NonNullable<FetchUserQuery['buildInfo']>;

}
