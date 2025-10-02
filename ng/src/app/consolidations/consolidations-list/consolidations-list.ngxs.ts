import {Injectable} from "@angular/core";
import {Action, Selector, State, StateContext, Store} from "@ngxs/store";
import {toNotNullArray} from "../../functions/not-null-array";
import {ConsolidationsListActions} from "./consolidations-list.actions";
import ConsolidationRecord = ConsolidationsListActions.ConsolidationRecord;
import {AddConsolidationGQL, FetchConsolidationsGQL} from "./consolidations-list.generated";
import SetConsolidationsList = ConsolidationsListActions.SetConsolidationsList;
import {AppActions} from "../../app.actions";
import AppChangeRoute = AppActions.AppChangeRoute;
import {Paths} from "../../app-routing.module";
import ShowGlobalSnackbar = AppActions.ShowGlobalSnackbar;
import { HttpErrorResponse } from "@angular/common/http";

export interface ConsolidationsListModel {
	consolidations: ConsolidationRecord[];
	loading: boolean;
}

const defaultState: ConsolidationsListModel = {
	consolidations: [],
	loading: false,
};

@Injectable()
@State<ConsolidationsListModel>({
	name: 'consolidationsList',
	defaults: defaultState,
})
export class ConsolidationsListState {

	constructor(
		private fetch: FetchConsolidationsGQL,
		private add: AddConsolidationGQL,
	) {}

	@Selector()
	static get(state: ConsolidationsListModel) {
		return state;
	}

	@Action(ConsolidationsListActions.FetchConsolidationsList)
	FetchConsolidationsList(ctx: StateContext<ConsolidationsListModel>, action: ConsolidationsListActions.FetchConsolidationsList) {
		ctx.patchState({loading: true});
		return this.fetch.fetch({}, {errorPolicy: "all"})
			.subscribe((r) => {
				ctx.patchState({loading: false});
				ctx.dispatch([
					new SetConsolidationsList(toNotNullArray(r.data.consolidations.edges?.map((n) => n?.node))),
				]);
			});
	}

	@Action(ConsolidationsListActions.SetConsolidationsList)
	SetHypothesisTest(ctx: StateContext<ConsolidationsListModel>, action: ConsolidationsListActions.SetConsolidationsList) {
		ctx.patchState({consolidations: action.payload})
	}

	@Action(ConsolidationsListActions.AddConsolidation)
	AddConsolidation(ctx: StateContext<ConsolidationsListModel>, action: ConsolidationsListActions.AddConsolidation) {
		ctx.patchState({loading: true});
		return this.add.mutate(action.payload)
			.subscribe({
				next: (r) => {
					ctx.dispatch(new AppChangeRoute({path: Paths.CONSOLIDATIONS, queryParams: {id: r.data?.createConsolidation.id}}));
				},
				error: (e: HttpErrorResponse) => {
					ctx.dispatch(new ShowGlobalSnackbar("Error: " + JSON.stringify(e.message)));
				}
			})
	}

}
