import {Injectable} from "@angular/core";
import {Action, Selector, State, StateContext, Store} from "@ngxs/store";
import {Apollo} from "apollo-angular";
import {ConnectionsListActions} from "./connections-list.actions";
import SelectConnectionsListQueryResponse = ConnectionsListActions.ConnectionResponse;
import SetConnectionsList = ConnectionsListActions.SetConnectionsList;
import {AppService} from "../../../app.service";
import {AppState} from "../../../app.ngxs";
import {ListConnectionsGQL} from "./connections-list.generated";
import {toNotNullArray} from "../../../functions/not-null-array";

export interface ConnectionsListModel {
	connectionsList: SelectConnectionsListQueryResponse[];
	loading: boolean;
}

const defaultState: ConnectionsListModel = {
	connectionsList: [],
	loading: false,
};

@Injectable()
@State<ConnectionsListModel>({
	name: 'connectionsList',
	defaults: defaultState,
})
export class ConnectionsListState {

	constructor(
		private list: ListConnectionsGQL,
		private store: Store,
	) {
	}

	@Selector()
	static get(state: ConnectionsListModel) {
		return state;
	}

	@Action(ConnectionsListActions.FetchConnectionsList)
	FetchMyConnectionsList(ctx: StateContext<ConnectionsListModel>, action: ConnectionsListActions.FetchConnectionsList) {
		ctx.patchState({loading: true});
		return this.list.fetch({})
			.subscribe({next: (r) => {
					ctx.patchState({loading: false});
				const conns = toNotNullArray(r.data.connections.edges?.map((n) => n?.node));
				ctx.dispatch(new SetConnectionsList(conns));
			}});
	}

	@Action(ConnectionsListActions.SetConnectionsList)
	SetMyConnectionsList(ctx: StateContext<ConnectionsListModel>, action: ConnectionsListActions.SetConnectionsList) {
		ctx.patchState({connectionsList: action.payload});
	}

	@Action(ConnectionsListActions.Clear)
	Clear(ctx: StateContext<ConnectionsListModel>, action: ConnectionsListActions.Clear) {
		ctx.setState(defaultState);
	}

}
