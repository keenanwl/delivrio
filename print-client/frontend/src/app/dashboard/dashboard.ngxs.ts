import {Injectable} from "@angular/core";
import {Action, Selector, State, StateContext} from "@ngxs/store";
import {ActivePrintJobs, JobPendingCancel, RecentScans} from "../../../wailsjs/go/main/App";
import {main} from "../../../wailsjs/go/models";
import PrintJobDisplay = main.PrintJobDisplay;
import {DashboardActions} from "./dashboard.actions";
import SetActivePrintJobs = DashboardActions.SetActivePrintJobs;
import {AppActions} from "../app.actions";
import ShowGlobalSnackbar = AppActions.ShowGlobalSnackbar;
import RecentScan = main.RecentScan;

export interface DashboardModel {
	jobs: PrintJobDisplay[];
	recentJobs: PrintJobDisplay[];
	recentScans: RecentScan[];
}

const defaultState: DashboardModel = {
	jobs: [],
	recentJobs: [],
	recentScans: [],
};

@Injectable()
@State<DashboardModel>({
	name: 'dashboard',
	defaults: defaultState,
})
export class DashboardState {

	constructor() {
	}

	@Selector()
	static get(state: DashboardModel) {
		return state;
	}

	@Action(DashboardActions.Refresh)
	Refresh(ctx: StateContext<DashboardModel>, action: DashboardActions.Refresh) {
		return ActivePrintJobs().then((jobs) => {
			ctx.dispatch(new SetActivePrintJobs({current: jobs.current_jobs, recent: jobs.recent_jobs}));
		});
	}

	@Action(DashboardActions.SetActivePrintJobs)
	SetActivePrintJobs(ctx: StateContext<DashboardModel>, action: DashboardActions.SetActivePrintJobs) {
		ctx.patchState({jobs: action.payload.current, recentJobs: action.payload.recent});
	}

	@Action(DashboardActions.SetPrintJobPendingCancel)
	SetPrintJobPendingCancel(ctx: StateContext<DashboardModel>, action: DashboardActions.SetPrintJobPendingCancel) {
		return JobPendingCancel(action.payload.id, action.payload.messages).then((r) => {
			if (r.length === 0) {
				ctx.dispatch(new ShowGlobalSnackbar("Cancelled successfully"));
			} else {
				ctx.dispatch(new ShowGlobalSnackbar("Error: " + r));
			}
			ctx.dispatch(new DashboardActions.Refresh());
		})
	}
	@Action(DashboardActions.FetchRecentScans)
	FetchRecentScans(ctx: StateContext<DashboardModel>, action: DashboardActions.FetchRecentScans) {
		return RecentScans().then((r) => {
			ctx.dispatch(new DashboardActions.SetRecentScans(r));
		}, (e) => {
			ctx.dispatch(new ShowGlobalSnackbar("Error: " + e));
		})
	}

	@Action(DashboardActions.SetRecentScans)
	SetRecentScans(ctx: StateContext<DashboardModel>, action: DashboardActions.SetRecentScans) {
		ctx.patchState({recentScans: action.payload});
	}

}
