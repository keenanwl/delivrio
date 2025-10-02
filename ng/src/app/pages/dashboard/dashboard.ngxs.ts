import {Injectable} from "@angular/core";
import {Action, Selector, State, StateContext, Store} from "@ngxs/store";
import {DashboardActions} from "./dashboard.actions";
import TileResponse = DashboardActions.TileResponse;
import {FetchDashboardTilesGQL} from "./dashboard.generated";
import SetTiles = DashboardActions.SetTiles;
import SetHypothesisTest = DashboardActions.SetHypothesisTest;
import HypothesisTestResponse = DashboardActions.HypothesisTestResponse;
import ProductUpdateCount = DashboardActions.ProductUpdateCount;
import {RateRequests} from "../../../generated/graphql";

export interface DashboardModel {
	tiles: TileResponse[];
	hypothesisTests: HypothesisTestResponse[];
	productUpdateCounts: ProductUpdateCount[];
	rateRequest: RateRequests;
}

const defaultState: DashboardModel = {
	tiles: [],
	hypothesisTests: [],
	productUpdateCounts: [0, 0, 0, 0, 0, 0, 0],
	rateRequest: {
		requests: [],
		requestsError: []
	},
};

@Injectable()
@State<DashboardModel>({
	name: 'dashboard',
	defaults: defaultState,
})
export class DashboardState {

	constructor(
		private fetch: FetchDashboardTilesGQL,
	) {
	}

	@Selector()
	static get(state: DashboardModel) {
		return state;
	}

	@Action(DashboardActions.FetchTiles)
	FetchLoggedInUser(ctx: StateContext<DashboardModel>, action: DashboardActions.FetchTiles) {
		return this.fetch.fetch()
			.subscribe((r) => {
				const tiles = r.data.dashboardTiles;
				if (!!tiles) {
					ctx.dispatch(new SetTiles(tiles));
				}
				const tests = r.data.hypothesisTestResultsDashboard;
				if (!!tests) {
					ctx.dispatch(new SetHypothesisTest(tests));
				}

				const productUpdates = r.data.trailingProductUpdates;
				if (!!tests) {
					ctx.dispatch(new DashboardActions.SetTrailingProductUpdateCounts(productUpdates));
				}
				const rateRequests = r.data.rateRequests;
				if (!!rateRequests) {
					ctx.dispatch(new DashboardActions.SetRateRequests(rateRequests));
				}
			});
	}

	@Action(DashboardActions.SetTiles)
	SetTiles(ctx: StateContext<DashboardModel>, action: DashboardActions.SetTiles) {
		ctx.patchState({tiles: action.payload})
	}

	@Action(DashboardActions.SetHypothesisTest)
	SetHypothesisTest(ctx: StateContext<DashboardModel>, action: DashboardActions.SetHypothesisTest) {
		ctx.patchState({hypothesisTests: action.payload})
	}

	@Action(DashboardActions.SetTrailingProductUpdateCounts)
	SetTrailingProductUpdateCounts(ctx: StateContext<DashboardModel>, action: DashboardActions.SetTrailingProductUpdateCounts) {
		ctx.patchState({productUpdateCounts: action.payload})
	}

}
