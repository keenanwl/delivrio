import {Injectable} from "@angular/core";
import {Action, Selector, State, StateContext} from "@ngxs/store";
import {GraphQLError} from "graphql";
/*import FetchReturnResponse = ReturnEditActions.FetchReturnResponse;
import FetchReturnTagsResponse = ReturnEditActions.FetchReturnTagsResponse;*/
import {produce} from "immer";
import {ReturnEditActions} from "./return-edit.actions";

export interface ReturnModel {
	returnForm: {
		//model: FetchReturnResponse | undefined;
		dirty: boolean;
		status: string;
		errors: readonly GraphQLError[];
	},
	returnID: string;
}

const defaultState: ReturnModel = {
	returnForm: {
		//model: undefined,
		dirty: false,
		status: '',
		errors: [],
	},
	returnID: '',
};

@Injectable()
@State<ReturnModel>({
	name: 'returnEdit',
	defaults: defaultState,
})
export class ReturnEditState {

	constructor(
		/*private fetchReturn: FetchReturnGQL,
		private createReturn: CreateReturnGQL,
		private updateReturn: UpdateReturnGQL,
		private createVariant: CreateVariantGQL,*/
	) {}

	@Selector()
	static get(state: ReturnModel) {
		return state;
	}

	@Action(ReturnEditActions.FetchReturn)
	FetchMyReturn(ctx: StateContext<ReturnModel>, action: ReturnEditActions.FetchReturn) {
		/*const state = ctx.getState();
		return this.fetchReturn.fetch({id: state.returnID}, {fetchPolicy: "no-cache", errorPolicy: "all"})
			.subscribe({next: (r) => {
					const ret = r.data.return;
					if (!!ret) {
						ctx.dispatch(new SetReturn(ret));
					}
					const groups = r.data.returnTags.edges?.map(g => g?.node);
					if (!!groups) {
						ctx.dispatch(new SetReturnTags(groups));
					}
				}, error: (e) => {

				}});*/
	}

	@Action(ReturnEditActions.SetReturnID)
	SetReturnID(ctx: StateContext<ReturnModel>, action: ReturnEditActions.SetReturnID) {
		ctx.patchState({
			returnID: action.payload,
		})
	}

	@Action(ReturnEditActions.SetReturn)
	SetMyReturn(ctx: StateContext<ReturnModel>, action: ReturnEditActions.SetReturn) {
/*		const state = ctx.getState();
		const next = Object.assign({}, state.returnForm, {
			model: Object.assign({}, action.payload)
		});
		ctx.patchState({
			returnForm: next,
		})*/
	}

}
