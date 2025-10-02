import {Injectable} from "@angular/core";
import {Action, Selector, State, StateContext, Store} from "@ngxs/store";
import {APITokensActions} from "./api-tokens.actions";
import {CreateApiTokenGQL, DeleteTokenGQL, FetchMyApiTokensGQL} from "./api-tokens.generated";
import {toNotNullArray} from "../../functions/not-null-array";
import SetAPITokens = APITokensActions.SetAPITokens;
import FetchAPITokens = APITokensActions.FetchAPITokens;
import FetchAPITokensResponse = APITokensActions.FetchAPITokensResponse;
import SetNewToken = APITokensActions.SetNewToken;
import {AppActions} from "../../app.actions";
import ShowGlobalSnackbar = AppActions.ShowGlobalSnackbar;

export interface APITokensModel {
	newToken: string;
	confirmDeleteID: string;
	myAPITokens: FetchAPITokensResponse[];
	loading: boolean;
}

const defaultState: APITokensModel = {
	newToken: '',
	confirmDeleteID: '',
	myAPITokens: [],
	loading: false,
};

@Injectable()
@State<APITokensModel>({
	name: 'apiTokens',
	defaults: defaultState,
})
export class APITokensState {

	constructor(
		private list: FetchMyApiTokensGQL,
		private create: CreateApiTokenGQL,
		private del: DeleteTokenGQL,
		private store: Store,
	) {
	}

	@Selector()
	static get(state: APITokensModel) {
		return state;
	}

	@Action(APITokensActions.FetchAPITokens)
	FetchAPITokens(ctx: StateContext<APITokensModel>, action: APITokensActions.FetchAPITokens) {
		ctx.patchState({loading: true});
		return this.list.fetch({})
			.subscribe({next: (r) => {
				ctx.patchState({loading: false});
				const tokens = toNotNullArray(r.data.myAPITokens.map((c) => c));
				if (!!tokens) {
					ctx.dispatch(new SetAPITokens(tokens));
				}
			}});
	}

	@Action(APITokensActions.SetAPITokens)
	SetAPITokens(ctx: StateContext<APITokensModel>, action: APITokensActions.SetAPITokens) {
		ctx.patchState({myAPITokens: action.payload})
	}

	@Action(APITokensActions.SetNewToken)
	SetNewToken(ctx: StateContext<APITokensModel>, action: APITokensActions.SetNewToken) {
		ctx.patchState({newToken: action.payload.token, loading: false})
	}

	@Action(APITokensActions.ClearDialogs)
	ClearNewToken(ctx: StateContext<APITokensModel>, action: APITokensActions.ClearDialogs) {
		ctx.patchState({newToken: '', confirmDeleteID: '', loading: false});
	}

	@Action(APITokensActions.SetConfirmDeleteToken)
	SetConfirmDeleteToken(ctx: StateContext<APITokensModel>, action: APITokensActions.SetConfirmDeleteToken) {
		ctx.patchState({confirmDeleteID: action.payload});
	}

	@Action(APITokensActions.DeleteToken)
	DeleteToken(ctx: StateContext<APITokensModel>, action: APITokensActions.DeleteToken) {
		ctx.patchState({loading: true});
		return this.del.mutate({ID: action.payload})
			.subscribe((resp) => {
				if (!!resp.errors && resp.errors?.length > 0) {
					ctx.dispatch(new ShowGlobalSnackbar(`An error occurred deleting the token: ` + resp.errors?.join(', ')));
				} else {
					ctx.dispatch(new ShowGlobalSnackbar(`Token deleted`));
				}
				const tokens = toNotNullArray(resp.data?.deleteAPIToken.map((c) => c));
				if (!!tokens) {
					ctx.dispatch(new SetAPITokens(tokens));
				}
			});
	}

	@Action(APITokensActions.CreateAPIToken)
	CreateAPIToken(ctx: StateContext<APITokensModel>, action: APITokensActions.CreateAPIToken) {
		ctx.patchState({loading: true});
		return this.create.mutate({
			name: action.payload.name,
		}).subscribe((r) => {
			const resp = r.data?.createAPIToken;
			if (!!resp) {
				ctx.dispatch([
					new SetNewToken({token: resp.token}),
					new FetchAPITokens(),
				]);
			}
		});
	}

}
