import {Injectable} from "@angular/core";
import {Action, Selector, State, StateContext} from "@ngxs/store";
import {ReturnViewActions} from "./return-view.actions";
import {FetchReturnCollisViewGQL, UpdateReturnColliStatusGQL} from "./return-view.generated";
import {TimelineViewerFragment} from "../../shared/timeline-viewer/timeline-viewer.generated";
import {toNotNullArray} from "../../functions/not-null-array";
import {ReturnColliStatus} from "../../../generated/graphql";
import ReturnColliResponse = ReturnViewActions.ReturnColliResponse;
import SetReturnView = ReturnViewActions.SetReturnView;
import SetOrderPublicID = ReturnViewActions.SetOrderPublicID;
import SetTimeline = ReturnViewActions.SetTimeline;
import {AppActions} from "../../app.actions";
import ShowGlobalSnackbar = AppActions.ShowGlobalSnackbar;

export interface ReturnViewModel {
	orderPublicID: string;
	orderID: string;
	view: ReturnColliResponse[];
	showDeleted: boolean;
}

const defaultState: ReturnViewModel = {
	orderID: '',
	orderPublicID: '',
	view: [],
	showDeleted: false,
};

@Injectable()
@State<ReturnViewModel>({
	name: 'returnView',
	defaults: defaultState,
})
export class ReturnViewState {

	constructor(
		private view: FetchReturnCollisViewGQL,
		private changeStatus: UpdateReturnColliStatusGQL,
	) {}

	@Selector()
	static get(state: ReturnViewModel) {
		return state;
	}

	@Action(ReturnViewActions.FetchReturnView)
	FetchReturnView(ctx: StateContext<ReturnViewModel>, action: ReturnViewActions.FetchReturnView) {
		const state = ctx.getState();
		return this.view.fetch({orderID: state.orderID}, {fetchPolicy: "no-cache", errorPolicy: "all"})
			.subscribe({next: (r) => {
					ctx.dispatch([
						new SetReturnView(toNotNullArray(r.data.returnColli.collis)),
						new SetOrderPublicID(r.data.returnColli.order.orderPublicID),
					]);
				}});
	}

	@Action(ReturnViewActions.SetOrderID)
	SetOrderID(ctx: StateContext<ReturnViewModel>, action: ReturnViewActions.SetOrderID) {
		ctx.patchState({orderID: action.payload});
	}

	@Action(ReturnViewActions.SetReturnView)
	SetReturnView(ctx: StateContext<ReturnViewModel>, action: ReturnViewActions.SetReturnView) {
		ctx.patchState({view: action.payload});
	}

	@Action(ReturnViewActions.ToggleShowDeleted)
	ToggleShowDeleted(ctx: StateContext<ReturnViewModel>, action: ReturnViewActions.ToggleShowDeleted) {
		ctx.patchState({showDeleted: !ctx.getState().showDeleted});
	}

	@Action(ReturnViewActions.SetOrderPublicID)
	SetOrderPublicID(ctx: StateContext<ReturnViewModel>, action: ReturnViewActions.SetOrderPublicID) {
		ctx.patchState({orderPublicID: action.payload});
	}

	@Action(ReturnViewActions.MarkAccepted)
	MarkAccepted(ctx: StateContext<ReturnViewModel>, action: ReturnViewActions.MarkAccepted) {
		return this.changeStatus.mutate({returnColliID: action.payload, status: ReturnColliStatus.Accepted}, {errorPolicy: "all"})
			.subscribe((r) => {
				if (!r.errors) {
					ctx.dispatch([
						new SetReturnView(toNotNullArray(r.data?.updateReturnColliStatus.collis)),
						new ShowGlobalSnackbar("Return colli marked as accepted"),
					]);

				} else {
					ctx.dispatch([
						new ShowGlobalSnackbar("An error occurred: " + r.errors[0].message),
					]);
				}
			});
	}

	@Action(ReturnViewActions.MarkDeclined)
	MarkDeclined(ctx: StateContext<ReturnViewModel>, action: ReturnViewActions.MarkDeclined) {
		return this.changeStatus.mutate({returnColliID: action.payload, status: ReturnColliStatus.Declined}, {errorPolicy: "all"})
			.subscribe((r) => {
				if (!r.errors) {
					ctx.dispatch([
						new SetReturnView(toNotNullArray(r.data?.updateReturnColliStatus.collis)),
						new ShowGlobalSnackbar("Return colli marked as declined"),
					]);
				} else {
					ctx.dispatch([
						new ShowGlobalSnackbar("An error occurred: " + r.errors[0].message),
					]);
				}
			});
	}

}
