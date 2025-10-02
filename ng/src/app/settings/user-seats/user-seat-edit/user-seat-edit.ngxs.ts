import {Injectable} from "@angular/core";
import {Action, Selector, State, StateContext} from "@ngxs/store";
import {AppActions} from "../../../app.actions";
import AppChangeRoute = AppActions.AppChangeRoute;
import {GraphQLError} from "graphql";
import {UserSeatsActions} from "./user-seat-edit.actions";
import FetchUserSeatResponse = UserSeatsActions.FetchUserSeatResponse;
import SetUserSeats = UserSeatsActions.SetUserSeat;
import {UserSeatsListActions} from "../user-seats-list/user-seats-list.actions";
import SetSeatGroups = UserSeatsActions.SetSeatGroups;
import FetchSeatGroupsResponse = UserSeatsActions.FetchSeatGroupsResponse;
import {Paths} from "../../../app-routing.module";
import {CreateUserSeatGQL, FetchUserSeatGQL, UpdatePasswordGQL, UpdateUserSeatGQL} from "./user-seat-edit.generated";
import ShowGlobalSnackbar = AppActions.ShowGlobalSnackbar;

export interface UserSeatsModel {
	userSeatsForm: {
		model: FetchUserSeatResponse | undefined;
		dirty: boolean;
		status: string;
		errors: readonly GraphQLError[];
	},
	userSeatID: string;
	seatGroups: FetchSeatGroupsResponse[];
}

const defaultState: UserSeatsModel = {
	userSeatsForm: {
		model: undefined,
		dirty: false,
		status: '',
		errors: [],
	},
	userSeatID: '',
	seatGroups: [],
};

@Injectable()
@State<UserSeatsModel>({
	name: 'userSeats',
	defaults: defaultState,
})
export class UserSeatsState {

	constructor(
		private fetchUserSeat: FetchUserSeatGQL,
		private createUserSeat: CreateUserSeatGQL,
		private updateUserSeat: UpdateUserSeatGQL,
		private updatePassword: UpdatePasswordGQL,
	) {}

	@Selector()
	static get(state: UserSeatsModel) {
		return state;
	}

	@Action(UserSeatsActions.FetchUserSeat)
	FetchMyUserSeats(ctx: StateContext<UserSeatsModel>, action: UserSeatsActions.FetchUserSeat) {
		const state = ctx.getState();
		return this.fetchUserSeat.fetch({id: state.userSeatID}, {fetchPolicy: "no-cache", errorPolicy: "all"})
			.subscribe({next: (r) => {
				const user = r.data.user;
				if (!!user) {
					ctx.dispatch(new SetUserSeats(user));
				}
				const groups = r.data.seatGroups.edges?.map(g => g?.node);
				if (!!groups) {
					ctx.dispatch(new SetSeatGroups(groups));
				}
			}});
	}

	@Action(UserSeatsActions.SetUserSeatID)
	SetUserSeatID(ctx: StateContext<UserSeatsModel>, action: UserSeatsActions.SetUserSeatID) {
		ctx.patchState({
			userSeatID: action.payload,
		})
	}

	@Action(UserSeatsActions.SetUserSeat)
	SetMyUserSeats(ctx: StateContext<UserSeatsModel>, action: UserSeatsActions.SetUserSeat) {
		const state = ctx.getState();
		const next = Object.assign({}, state.userSeatsForm, {
			model: Object.assign({}, action.payload, {name: action.payload.name})
		});
		ctx.patchState({
			userSeatsForm: next,
		})
	}

	@Action(UserSeatsActions.SaveFormNew)
	SaveFormNew(ctx: StateContext<UserSeatsModel>, action: UserSeatsActions.SaveFormNew) {
		return this.createUserSeat.mutate(action.payload, {errorPolicy: "all"})
			.subscribe((r) => {
				if (!r.errors) {
					ctx.dispatch([
						new UserSeatsListActions.FetchUserSeatsList(),
						new AppChangeRoute({path: Paths.SETTINGS_USERS, queryParams: {}}),
					]);
				} else {
					ctx.dispatch(new ShowGlobalSnackbar("Errors: " + JSON.stringify(r.errors)));
				}
			});
	}

	@Action(UserSeatsActions.SaveFormUpdate)
	SaveFormUpdate(ctx: StateContext<UserSeatsModel>, action: UserSeatsActions.SaveFormUpdate) {
		return this.updateUserSeat.mutate(action.payload, {errorPolicy: "all"})
			.subscribe((r) => {
				if (!r.errors) {
					ctx.dispatch([
						new UserSeatsListActions.FetchUserSeatsList(),
						new AppChangeRoute({path: Paths.SETTINGS_USERS, queryParams: {}}),
					]);
				} else {
					ctx.dispatch(new ShowGlobalSnackbar("Errors: " + JSON.stringify(r.errors)));
				}
			});
	}

	@Action(UserSeatsActions.SetSeatGroups)
	SetSeatGroups(ctx: StateContext<UserSeatsModel>, action: UserSeatsActions.SetSeatGroups) {
		ctx.patchState({
			seatGroups: action.payload,
		})
	}

	@Action(UserSeatsActions.UpdatePassword)
	UpdatePassword(ctx: StateContext<UserSeatsModel>, action: UserSeatsActions.UpdatePassword) {
		return this.updatePassword.mutate({id: action.payload.userID, input: action.payload.password})
			.subscribe((r) => {
				if (!r.errors) {
					ctx.dispatch(new ShowGlobalSnackbar("Password updated"));
				} else {
					ctx.dispatch(new ShowGlobalSnackbar("Errors: " + JSON.stringify(r.errors)));
				}
			});
	}

	@Action(UserSeatsActions.Clear)
	Clear(ctx: StateContext<UserSeatsModel>, action: UserSeatsActions.Clear) {
		ctx.setState(defaultState);
	}

}
