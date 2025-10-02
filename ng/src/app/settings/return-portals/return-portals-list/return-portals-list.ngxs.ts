import {Injectable} from "@angular/core";
import {Action, Selector, State, StateContext, Store} from "@ngxs/store";
import SetReturnPortalsList = ReturnPortalsListActions.SetReturnPortalsList;
import {ReturnPortalsListActions} from "./return-portals-list.actions";
import {CreateReturnPortalGQL, FetchReturnPortalsGQL} from "./return-portals-list.generated";
import ReturnPortalsListResponse = ReturnPortalsListActions.ReturnPortalsListResponse;
import {toNotNullArray} from "../../../functions/not-null-array";
import {AppActions} from "../../../app.actions";
import AppChangeRoute = AppActions.AppChangeRoute;
import {Paths} from "../../../app-routing.module";
import ShowGlobalSnackbar = AppActions.ShowGlobalSnackbar;
import SetConnections = ReturnPortalsListActions.SetConnections;
import ConnectionsResponse = ReturnPortalsListActions.ConnectionsResponse;

export interface ReturnPortalsListModel {
	returnPortalsList: ReturnPortalsListResponse[];
	connections: ConnectionsResponse[];
	loading: boolean;
}

const defaultState: ReturnPortalsListModel = {
	returnPortalsList: [],
	connections: [],
	loading: false,
};

@Injectable()
@State<ReturnPortalsListModel>({
	name: 'returnPortalsList',
	defaults: defaultState,
})
export class ReturnPortalsListState {

	constructor(
		private list: FetchReturnPortalsGQL,
		private store: Store,
		private create: CreateReturnPortalGQL,
	) {
	}

	@Selector()
	static get(state: ReturnPortalsListModel) {
		return state;
	}

	@Action(ReturnPortalsListActions.FetchReturnPortalsList)
	FetchMyReturnPortalsList(ctx: StateContext<ReturnPortalsListModel>, action: ReturnPortalsListActions.FetchReturnPortalsList) {
		ctx.patchState({loading: true});
		return this.list.fetch({})
			.subscribe({
				next: (r) => {
					ctx.patchState({loading: false});
					const portals = toNotNullArray(r.data.returnPortals.edges?.map((rp) => rp?.node));
					ctx.dispatch(new SetReturnPortalsList(portals));

					const connections = toNotNullArray(r.data.connections.edges?.map((rp) => rp?.node));
					ctx.dispatch(new SetConnections(connections));
				},
				error: () => {
					ctx.patchState({loading: false});
				},
			});
	}

	@Action(ReturnPortalsListActions.SetReturnPortalsList)
	SetMyReturnPortalsList(ctx: StateContext<ReturnPortalsListModel>, action: ReturnPortalsListActions.SetReturnPortalsList) {
		ctx.patchState({returnPortalsList: action.payload});
	}

	@Action(ReturnPortalsListActions.SetConnections)
	SetConnections(ctx: StateContext<ReturnPortalsListModel>, action: ReturnPortalsListActions.SetConnections) {
		ctx.patchState({connections: action.payload});
	}

	@Action(ReturnPortalsListActions.Create)
	Create(ctx: StateContext<ReturnPortalsListModel>, action: ReturnPortalsListActions.Create) {
		return this.create.mutate({name: action.payload.name, connection: action.payload.connection})
			.subscribe((rp) => {
				if (!!rp.errors) {
					ctx.dispatch(new ShowGlobalSnackbar("Error creating new return portal"));
				} else {
					const id = rp.data?.createReturnPortal;
					ctx.dispatch(new AppChangeRoute({
						path: Paths.SETTINGS_RETURN_PORTALS_EDIT,
						queryParams: {id: id}}
					));
				}
			});
	}

	@Action(ReturnPortalsListActions.Clear)
	Clear(ctx: StateContext<ReturnPortalsListModel>, action: ReturnPortalsListActions.Clear) {
		ctx.setState(defaultState);
	}

}
