import {Injectable} from "@angular/core";
import {Action, Selector, State, StateContext} from "@ngxs/store";
import {GraphQLError} from "graphql";
import {SeatGroupActions} from "./seat-group-edit.actions";
import FetchSeatGroupResponse = SeatGroupActions.FetchSeatGroupResponse;
import SetSeatGroup = SeatGroupActions.SetSeatGroup;
import SetAccessRights = SeatGroupActions.SetAccessRights;
import FetchAccessRightsResponse = SeatGroupActions.FetchAccessRightsResponse;
import {SeatGroupsListActions} from "../seat-groups-list/seat-groups-list.actions";
import {AppActions} from "../../../app.actions";
import AppChangeRoute = AppActions.AppChangeRoute;
import {Paths} from "src/app/app-routing.module";
import {CreateSeatGroupGQL, FetchSeatGroupGQL, ReplaceSeatGroupGQL} from "./seat-group-edit.generated";
import {SeatGroupAccessRightLevel} from "../../../../generated/graphql";

export interface SeatGroupModel {
	seatGroupForm: {
		model: FetchSeatGroupResponse | undefined;
		dirty: boolean;
		status: string;
		errors: readonly GraphQLError[];
	},
	seatGroupID: string;
	accessRights: FetchAccessRightsResponse[];
}

const defaultState: SeatGroupModel = {
	seatGroupForm: {
		model: undefined,
		dirty: false,
		status: '',
		errors: [],
	},
	seatGroupID: '',
	accessRights: [],
};

@Injectable()
@State<SeatGroupModel>({
	name: 'seatGroup',
	defaults: defaultState,
})
export class SeatGroupState {

	constructor(
		private fetchSeatGroup: FetchSeatGroupGQL,
		private createSeatGroup: CreateSeatGroupGQL,
		private replaceSeatGroup: ReplaceSeatGroupGQL,
	) {}

	@Selector()
	static get(state: SeatGroupModel) {
		return state;
	}

	@Selector()
	static accessRights(state: SeatGroupModel) {
		const m = new Map<string, SeatGroupAccessRightLevel>()
		state.seatGroupForm.model?.seatGroupAccessRight?.forEach((a) => m.set(a.accessRight.internalID, a.level));
		return m;
	}

	@Action(SeatGroupActions.FetchSeatGroup)
	FetchMySeatGroup(ctx: StateContext<SeatGroupModel>, action: SeatGroupActions.FetchSeatGroup) {
		const state = ctx.getState();
		return this.fetchSeatGroup.fetch({id: state.seatGroupID}, {fetchPolicy: "no-cache", errorPolicy: "all"})
			.subscribe({next: (r) => {
				const edit = r.data.seatGroup;
				if (!!edit) {
					ctx.dispatch(new SetSeatGroup(edit));
				}
				const ar = r.data.accessRights.edges?.map(n => n?.node);
				if (!!ar) {
					ctx.dispatch(new SetAccessRights(ar));
				}

			}});
	}

	@Action(SeatGroupActions.SetSeatGroupID)
	SetSeatGroupID(ctx: StateContext<SeatGroupModel>, action: SeatGroupActions.SetSeatGroupID) {
		ctx.patchState({seatGroupID: action.payload})
	}

	@Action(SeatGroupActions.SetSeatGroup)
	SetMySeatGroup(ctx: StateContext<SeatGroupModel>, action: SeatGroupActions.SetSeatGroup) {
		const state = ctx.getState();
		const next = Object.assign({}, state.seatGroupForm, {
			model: Object.assign({}, action.payload, {name: action.payload.name})
		});
		ctx.patchState({seatGroupForm: next})
	}

	@Action(SeatGroupActions.SetAccessRights)
	SetAccessRights(ctx: StateContext<SeatGroupModel>, action: SeatGroupActions.SetAccessRights) {
		ctx.patchState({accessRights: action.payload});
	}

	@Action(SeatGroupActions.SaveFormNew)
	SaveFormNew(ctx: StateContext<SeatGroupModel>, action: SeatGroupActions.SaveFormNew) {
		return this.createSeatGroup.mutate(action.payload)
			.subscribe(() => {
				ctx.dispatch([
					new SeatGroupsListActions.FetchSeatGroupsList(),
					new AppChangeRoute({path: Paths.SETTINGS_USERS_GROUP, queryParams: {}}),
				]);
			});
	}

	@Action(SeatGroupActions.SaveFormEdit)
	SaveFormEdit(ctx: StateContext<SeatGroupModel>, action: SeatGroupActions.SaveFormEdit) {
		return this.replaceSeatGroup.mutate(action.payload)
			.subscribe(() => {
				ctx.dispatch([
					new SeatGroupsListActions.FetchSeatGroupsList(),
					new AppChangeRoute({path: Paths.SETTINGS_USERS_GROUP, queryParams: {}}),
				]);
			});
	}

	@Action(SeatGroupActions.Clear)
	Clear(ctx: StateContext<SeatGroupModel>, action: SeatGroupActions.Clear) {
		ctx.setState(defaultState);
	}

}
