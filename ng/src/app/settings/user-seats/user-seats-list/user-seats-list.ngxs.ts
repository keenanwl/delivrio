import {Injectable} from "@angular/core";
import {Action, Selector, State, StateContext} from "@ngxs/store";
import {UserSeatsListActions} from "./user-seats-list.actions";
import SetUserSeatsList = UserSeatsListActions.SetUserSeatsList;
import FetchUserSeats = UserSeatsListActions.FecthUserSeats;
import {FetchUserSeatsGQL} from "./user-seats-list.generated";
import {toNotNullArray} from "../../../functions/not-null-array";

export interface UserSeatsListModel {
	list: FetchUserSeats[];
	loading: boolean;
}

const defaultState: UserSeatsListModel = {
	list: [],
	loading: false,
};

@Injectable()
@State<UserSeatsListModel>({
	name: 'userSeatsList',
	defaults: defaultState,
})
export class UserSeatsListState {

	constructor(
		private list: FetchUserSeatsGQL,
	) {
	}

	@Selector()
	static get(state: UserSeatsListModel) {
		return state;
	}

	@Action(UserSeatsListActions.FetchUserSeatsList)
	FetchMyUserSeatsList(ctx: StateContext<UserSeatsListModel>, action: UserSeatsListActions.FetchUserSeatsList) {
		ctx.patchState({loading: true});
		return this.list.fetch({}, {fetchPolicy: 'no-cache'})
			.subscribe({next: (r) => {
				ctx.patchState({loading: false});
				const list = toNotNullArray(r.data.users.edges?.map((n) => n?.node));
				ctx.dispatch(new SetUserSeatsList(list));
			}});
	}

	@Action(UserSeatsListActions.SetUserSeatsList)
	SetMyUserSeatsList(ctx: StateContext<UserSeatsListModel>, action: UserSeatsListActions.SetUserSeatsList) {
		ctx.patchState({
			list: action.payload,
		});
	}

}
