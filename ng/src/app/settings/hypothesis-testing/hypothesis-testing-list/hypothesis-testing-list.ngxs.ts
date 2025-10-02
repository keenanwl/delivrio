import {Injectable} from "@angular/core";
import {Action, Selector, State, StateContext} from "@ngxs/store";
import SetHypothesisTestingList = HypothesisTestingListActions.SetHypothesisTestingList;
import {toNotNullArray} from "../../../functions/not-null-array";
import {HypothesisTestingListActions} from "./hypothesis-testing-list.actions";
import HypothesisTestingResponse = HypothesisTestingListActions.HypothesisTestingResponse;
import {
	CreateHypothesisTestDeliveryOptionGQL,
	FetchHypothesisTestingListGQL
} from "./hypothesis-testing-list.generated";
import SetConnections = HypothesisTestingListActions.SetConnections;
import ConnectionsResponse = HypothesisTestingListActions.ConnectionsResponse;
import {AppActions} from "../../../app.actions";
import AppChangeRoute = AppActions.AppChangeRoute;
import {Paths} from "../../../app-routing.module";
import ShowGlobalSnackbar = AppActions.ShowGlobalSnackbar;

export interface HypothesisTestingListModel {
	hypothesisTestingList: HypothesisTestingResponse[];
	connections: ConnectionsResponse[];
	loading: boolean;
}

const defaultState: HypothesisTestingListModel = {
	hypothesisTestingList: [],
	connections: [],
	loading: false,
};

@Injectable()
@State<HypothesisTestingListModel>({
	name: 'hypothesisTestingList',
	defaults: defaultState,
})
export class HypothesisTestingListState {

	constructor(
		private list: FetchHypothesisTestingListGQL,
		private create: CreateHypothesisTestDeliveryOptionGQL,
	) {}

	@Selector()
	static get(state: HypothesisTestingListModel) {
		return state;
	}

	@Action(HypothesisTestingListActions.FetchHypothesisTestingList)
	FetchMyHypothesisTestingList(ctx: StateContext<HypothesisTestingListModel>, action: HypothesisTestingListActions.FetchHypothesisTestingList) {
		ctx.patchState({loading: true});
		return this.list.fetch({})
			.subscribe({
				next: (r) => {
					ctx.patchState({loading: false});
					const hypothesisTesting = toNotNullArray(r.data.hypothesisTests.edges?.map((l) => l?.node));
					ctx.dispatch(new SetHypothesisTestingList(hypothesisTesting));

					const connections = toNotNullArray(r.data.connections.edges?.map((l) => l?.node));
					ctx.dispatch(new SetConnections(connections));
				}, error: () => {
					ctx.patchState({loading: false});
				}});
	}

	@Action(HypothesisTestingListActions.SetHypothesisTestingList)
	SetMyHypothesisTestingList(ctx: StateContext<HypothesisTestingListModel>, action: HypothesisTestingListActions.SetHypothesisTestingList) {
		ctx.patchState({hypothesisTestingList: action.payload});
	}

	@Action(HypothesisTestingListActions.SetConnections)
	SetConnections(ctx: StateContext<HypothesisTestingListModel>, action: HypothesisTestingListActions.SetConnections) {
		ctx.patchState({connections: action.payload});
	}

	@Action(HypothesisTestingListActions.CreateNewTest)
	CreateNewTest(ctx: StateContext<HypothesisTestingListModel>, action: HypothesisTestingListActions.CreateNewTest) {
		return this.create.mutate({name: action.payload.name, connectionID: action.payload.connectionID})
			.subscribe((r) => {
				if (!r.errors) {
					ctx.dispatch(new AppChangeRoute({
						path: Paths.SETTINGS_HYPOTHESIS_TESTING_EDIT,
						queryParams: {id: r.data?.createHypothesisTestDeliveryOption}})
					);
				} else {
					ctx.dispatch(new ShowGlobalSnackbar("An error occurred."));
				}
			});
	}

	@Action(HypothesisTestingListActions.Reset)
	Reset(ctx: StateContext<HypothesisTestingListModel>, action: HypothesisTestingListActions.Reset) {
		ctx.setState(defaultState);
	}

}
