import {
	Action,
	NgxsOnInit,
	Selector, State,
	StateContext
} from '@ngxs/store';
import {Router, RoutesRecognized} from '@angular/router';
import {filter, pairwise} from 'rxjs/operators';
import {Injectable} from "@angular/core";
import {AppActions} from "./app.actions";

import {IsRegistered, WorkstationName} from "../../wailsjs/go/main/App";

export type Breakpoint = 'mobile' | 'tablet' | 'desktop';

export interface AppModel {
	isRegistered: boolean;
	route: string;
	workstationName: string;
}

const defaultState: AppModel = {
	isRegistered: false,
	route: "",
	workstationName: "",
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

	@Action(AppActions.AppGoBack)
	AppGoBack(ctx: StateContext<AppModel>, action: AppActions.AppGoBack) {

		ctx.dispatch([new AppActions.AppChangeRoute({path: this.lastUrl, queryParams: {}})]);

		// Move this to where changeRoute() is listened for
		this.route.navigate([this.lastUrl], {replaceUrl: true});

	}

	@Action(AppActions.Logout)
	Logout(ctx: StateContext<AppModel>, action: AppActions.Logout) {

		ctx.patchState({
			...defaultState
		});

		//ctx.dispatch(new AppActions.AppChangeRoute({path: Paths.LOGIN, queryParams: {}}))

	}

	@Action(AppActions.FetchIsRegistered)
	FetchIsRegistered(ctx: StateContext<AppModel>, action: AppActions.FetchIsRegistered) {
		IsRegistered().then((r) => {
			ctx.dispatch(new AppActions.SetIsRegistered(r.length === 0));
		});
	}

	@Action(AppActions.SetIsRegistered)
	SetIsRegistered(ctx: StateContext<AppModel>, action: AppActions.SetIsRegistered) {
		ctx.patchState({isRegistered: action.payload});
		ctx.dispatch(new AppActions.AppChangeRoute({path: action.payload ? "/dashboard" : "", queryParams: {}}))
	}

	@Action(AppActions.SetWorkstationName)
	SetWorkstationName(ctx: StateContext<AppModel>, action: AppActions.SetWorkstationName) {
		ctx.patchState({workstationName: action.payload});
	}

	@Action(AppActions.FetchWorkstationName)
	FetchWorkstationName(ctx: StateContext<AppModel>, action: AppActions.FetchWorkstationName) {
		WorkstationName().then((r) => {
			ctx.dispatch(new AppActions.SetWorkstationName(r));
		});
	}

}
