import {Injectable} from "@angular/core";
import {Action, Selector, State, StateContext, Store} from "@ngxs/store";
import {AppActions} from "../../app.actions";
import AppChangeRoute = AppActions.AppChangeRoute;
import {Register3Actions} from "./register3.actions";
import SaveRegistration = Register3Actions.SaveRegistration;
import SetInvalidMessage = Register3Actions.SetInvalidMessage;
import { HttpErrorResponse } from "@angular/common/http";
import FetchCarrierPlatformLists = Register3Actions.FetchCarrierPlatformLists;
import {AppState} from "../../app.ngxs";
import {Paths} from "src/app/app-routing.module";
import {FetchPlatformsCarriersGQL, ReplaceCarriersPlatformsGQL} from "./register3.generated";
import {toNotNullArray} from "../../functions/not-null-array";

export interface Register3Model {
	invalid_message: string;
	users_id: string;
	selected_carriers: string[];
	selected_platforms: string[];
	available_carriers: Register3Actions.Carriers[];
	available_platforms: Register3Actions.Carriers[];
}

const defaultState: Register3Model = {
	invalid_message: "",
	users_id: "",
	selected_carriers: [],
	selected_platforms: [],
	available_carriers: [],
	available_platforms: [],
};

@Injectable()
@State<Register3Model>({
	name: 'register3',
	defaults: defaultState,
})
export class Register3State {

	constructor(
		private lists: FetchPlatformsCarriersGQL,
		private store: Store,
		private save: ReplaceCarriersPlatformsGQL,
	) { }

	@Selector()
	static state(state: Register3Model) {
		return state;
	}

	@Action(Register3Actions.SetUserID)
	SetUserID(ctx: StateContext<Register3Model>, action: Register3Actions.SetUserID) {
		ctx.patchState({users_id: action.payload});
	}

	@Action(SaveRegistration)
	SaveRegistration(ctx: StateContext<Register3Model>) {
		const state = ctx.getState();
		const myID = this.store.selectSnapshot(AppState.get).my_ids.my_pulid;
		return this.save.mutate({userID: myID, inputCarriers: state.selected_carriers, inputPlatforms: state.selected_platforms})
			.subscribe(() => {
				ctx.dispatch(new AppChangeRoute({path: Paths.ORDERS, queryParams: {}}))
			}, (e: HttpErrorResponse) => ctx.dispatch(new SetInvalidMessage(e.message)));
	}

	@Action(FetchCarrierPlatformLists)
	FetchCarrierPlatformLists(ctx: StateContext<Register3Model>) {
		return this.lists.fetch().subscribe((r) => {
			const c = toNotNullArray(r.data.connectOptionCarriers.edges?.map((c) => c?.node));
			const p = toNotNullArray(r.data.connectOptionPlatforms.edges?.map((p) => p?.node));
			ctx.dispatch(new Register3Actions.SetCarrierPlatformLists({carriers: c, platforms: p}));
		});
	}

	@Action(Register3Actions.SetCarrierPlatformLists)
	SetCarrierPlatformLists(ctx: StateContext<Register3Model>, action: Register3Actions.SetCarrierPlatformLists) {
		ctx.patchState({available_carriers: action.payload.carriers, available_platforms: action.payload.platforms});
	}

	@Action(Register3Actions.SetPlatforms)
	SetPlatforms(ctx: StateContext<Register3Model>, action: Register3Actions.SetPlatforms) {
		ctx.patchState({selected_platforms: action.payload});
	}

	@Action(Register3Actions.SetCarriers)
	SetCarriers(ctx: StateContext<Register3Model>, action: Register3Actions.SetCarriers) {
		ctx.patchState({selected_carriers: action.payload});
	}

}
