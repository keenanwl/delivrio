import {Action, NgxsOnInit, Selector, State, StateContext} from '@ngxs/store';
import {Router, RoutesRecognized} from '@angular/router';
import {filter, pairwise} from 'rxjs/operators';
import {Injectable} from "@angular/core";
import {AppActions} from './app.actions';
import {CookieService} from "ngx-cookie-service";
import {Paths} from "./app-routing.module";
import {
	FetchSelectableWorkstationsGQL,
	FetchSelectedWorkstationGQL,
	FetchUserGQL,
	SaveSelectedWorkstationGQL,
	SearchGQL
} from "./app.generated";
import {SearchResult, UserPickupDay} from "../generated/graphql";
import {LoginActions} from "./login/login.actions";
import {toNotNullArray} from "./functions/not-null-array";
import SetLoggedInUser = AppActions.SetLoggedInUser;
import SetSearchResults = AppActions.SetSearchResults;
import ClearAllLoginData = LoginActions.ClearAllLoginData;
import Workstation = AppActions.Workstation;
import ShowGlobalSnackbar = AppActions.ShowGlobalSnackbar;
import PrintJob = AppActions.PrintJob;
import BuildInfoResponse = AppActions.BuildInfoResponse;
import SetBuildInfo = AppActions.SetBuildInfo;
import {timer} from "rxjs";
import FetchLoggedInUserQueryResponse = AppActions.FetchLoggedInUserQueryResponse;

export type Breakpoint = 'mobile' | 'tablet' | 'desktop';

export interface AppModel {
	route: string;
	breakpoint: Breakpoint;
	my_ids: AppActions.MyIds;
	my_info: FetchLoggedInUserQueryResponse;
	isAdmin: boolean;
	loading: boolean;
	searchResults: SearchResult[];
	selectableWorkstations: Workstation[];
	selectedWorkstation: Workstation | undefined | null;
	selectedPickupDay: UserPickupDay;
	selectedWorkstationJobs: PrintJob[];
	printJobLimitExceeded: boolean;
	buildInfo?: BuildInfoResponse;
}

const defaultState: AppModel = {
	route: '/login',
	breakpoint: "mobile",
	my_ids: {
		'my_pulid': '',
		'my_tenant_pulid': '',
	},
	my_info: {
		email: '',
		id: '',
		tenant: {
			id: '',
			name: ''
		}
	},
	isAdmin: false,
	loading: false,
	searchResults: [],
	selectableWorkstations: [],
	selectedPickupDay: UserPickupDay.Today,
	selectedWorkstation: null,
	selectedWorkstationJobs: [],
	printJobLimitExceeded: false,
	buildInfo: undefined,
};

@Injectable()
@State<AppModel>({
	name: 'appModel',
	defaults: defaultState,
})
export class AppState implements NgxsOnInit {

	private lastUrl = '/';

	constructor(
		public route: Router,
		private fetchUser: FetchUserGQL,
		private search: SearchGQL,
		private selectedWS: FetchSelectedWorkstationGQL,
		private ws: FetchSelectableWorkstationsGQL,
		private saveWS: SaveSelectedWorkstationGQL,
		private cookieService: CookieService,
	) {}

	ngxsOnInit(ctx: StateContext<AppModel>) {

		this.route.events
			.pipe(
				filter((e: any) => e instanceof RoutesRecognized),
				pairwise()
			).subscribe((e: any) => {
				this.lastUrl = e[0].urlAfterRedirects;
			});

	}

	@Selector()
	static get(state: AppModel) {
		return state;
	}

	@Action(AppActions.FetchLoggedInUser)
	FetchLoggedInUser(ctx: StateContext<AppModel>, action: AppActions.FetchLoggedInUser) {
		return this.fetchUser.fetch()
			.subscribe((r) => {
				const user = r.data.user;
				if (!!user) {
					ctx.dispatch(new SetLoggedInUser(user, {my_pulid: user.id, my_tenant_pulid: user.tenant.id}));
				}
				ctx.dispatch(new SetBuildInfo(r.data.buildInfo))
			});
	}

	@Action(AppActions.SetLoggedInUser)
	SetLoggedInUser(ctx: StateContext<AppModel>, action: AppActions.SetLoggedInUser) {
		ctx.patchState({
			my_info: action.payload,
			my_ids: action.myIDs,
		});
	}

	@Action(AppActions.FetchAppConfig)
	FetchAppConfig(ctx: StateContext<AppModel>, action: AppActions.FetchAppConfig) {
		ctx.patchState({loading: true});
	}

	@Action(AppActions.AppChangeRoute)
	AppChangeRoute(ctx: StateContext<AppModel>, action: AppActions.AppChangeRoute) {
		ctx.patchState({
			route: action.payload.path
		});
	}

	@Action(AppActions.AppSetRoute)
	AppSetRoute(ctx: StateContext<AppModel>, action: AppActions.AppSetRoute) {
		ctx.patchState({
			route: action.payload
		});
	}

	@Action(AppActions.UpdateBreakpoint)
	UpdateBreakpoint(ctx: StateContext<AppModel>, action: AppActions.UpdateBreakpoint) {
		ctx.patchState({
			breakpoint: action.payload
		});
	}

	@Action(AppActions.AppGoBack)
	AppGoBack(ctx: StateContext<AppModel>, action: AppActions.AppGoBack) {

		ctx.dispatch([new AppActions.AppChangeRoute({path: this.lastUrl, queryParams: {}})]);

		// Move this to where changeRoute() is listened for
		this.route.navigate([this.lastUrl], {replaceUrl: true});

	}

	@Action(AppActions.Logout)
	Logout(ctx: StateContext<AppModel>, action: AppActions.Logout) {

		this.cookieService.deleteAll(undefined, undefined, true);

		ctx.patchState({
			...defaultState
		});

		ctx.dispatch([
			new ClearAllLoginData(),
			new AppActions.AppChangeRoute({path: Paths.LOGIN, queryParams: {}}),
		]);

		timer(500).subscribe(() => {
			// Force refresh of state data for login.
			window.location.reload();
		});

	}

	@Action(AppActions.Search)
	Search(ctx: StateContext<AppModel>, action: AppActions.Search) {
		return this.search.fetch({term: action.payload})
			.subscribe((s) => {
				ctx.dispatch(new SetSearchResults(s.data.search))
			});
	}

	@Action(AppActions.SetSearchResults)
	SetSearchResults(ctx: StateContext<AppModel>, action: AppActions.SetSearchResults) {
		ctx.patchState({searchResults: action.payload});
	}

	@Action(AppActions.ClearSearchResults)
	ClearSearchResults(ctx: StateContext<AppModel>, action: AppActions.ClearSearchResults) {
		ctx.patchState({searchResults: []});
	}

	@Action(AppActions.FetchSelectedWorkstation)
	FetchSelectedWorkstation(ctx: StateContext<AppModel>, action: AppActions.FetchSelectedWorkstation) {
		return this.selectedWS.fetch().subscribe((resp) => {
			ctx.dispatch(new AppActions.SetSelectedWorkstation({
				workstation: resp.data.selectedWorkstation?.workstation,
				printJobs: resp.data.selectedWorkstation?.jobs || [],
				limitExceeded: resp.data.selectedWorkstation?.limitExceeded || false,
			}));
			const user = resp.data.user;
			if (!!user) {
				ctx.dispatch(new AppActions.SetSelectedPickup(user.pickupDay));
			}
		})
	}

	@Action(AppActions.SetSelectedWorkstation)
	SetSelectedWorkstation(ctx: StateContext<AppModel>, action: AppActions.SetSelectedWorkstation) {
		ctx.patchState({
			selectedWorkstation: action.payload.workstation,
			selectedWorkstationJobs: action.payload.printJobs,
			printJobLimitExceeded: action.payload.limitExceeded,
		})
	}

	@Action(AppActions.FetchSelectableWorkstations)
	FetchSelectableWorkstations(ctx: StateContext<AppModel>, action: AppActions.FetchSelectableWorkstations) {
		return this.ws.fetch().subscribe((resp) => {
			const ws = toNotNullArray(resp.data.workstations.edges?.map((n) => n?.node));
			if (!!ws) {
				ctx.dispatch(new AppActions.SetSelectableWorkstations(ws));
			}
		})
	}

	@Action(AppActions.SetSelectableWorkstations)
	SetSelectableWorkstations(ctx: StateContext<AppModel>, action: AppActions.SetSelectableWorkstations) {
		ctx.patchState({
			selectableWorkstations: action.payload,
		})
	}

	@Action(AppActions.SetSelectedPickup)
	SetSelectedPickup(ctx: StateContext<AppModel>, action: AppActions.SetSelectedPickup) {
		ctx.patchState({selectedPickupDay: action.payload});
	}

	@Action(AppActions.SaveSelectedWorkstation)
	SaveSelectableWorkstation(ctx: StateContext<AppModel>, action: AppActions.SaveSelectedWorkstation) {
		return this.saveWS.mutate({id: action.payload.workstationID, pickupDay: action.payload.pickupDay}, {errorPolicy: "all"})
			.subscribe((resp) => {
				if (!resp.errors) {
					ctx.dispatch(new ShowGlobalSnackbar("Saved successfully"));
				} else {
					ctx.dispatch(new ShowGlobalSnackbar("Error: " + JSON.stringify(resp.errors)));
				}
				ctx.dispatch(new AppActions.FetchSelectedWorkstation())
			});
	}

	@Action(AppActions.SetBuildInfo)
	SetBuildInfo(ctx: StateContext<AppModel>, action: AppActions.SetBuildInfo) {
		ctx.patchState({buildInfo: action.payload});
	}

}
