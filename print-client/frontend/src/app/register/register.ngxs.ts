import {Injectable} from "@angular/core";
import {Action, Selector, State, StateContext} from "@ngxs/store";
import {RegisterActions} from "./register.actions";
import {Register, ShowDialog} from "../../../wailsjs/go/main/App";
import {AppActions} from "../app.actions";
import FetchIsRegistered = AppActions.FetchIsRegistered;
import FetchWorkstationName = AppActions.FetchWorkstationName;

export interface RegisterModel {
	registrationURL: string;
	registrationToken: string;
}

const defaultState: RegisterModel = {
	registrationURL: '',
	registrationToken: '',
};

@Injectable()
@State<RegisterModel>({
	name: 'register',
	defaults: defaultState,
})
export class RegisterState {

	constructor() {
	}

	@Selector()
	static get(state: RegisterModel) {
		return state;
	}

	@Action(RegisterActions.Submit)
	FetchMyProductsList(ctx: StateContext<RegisterModel>, action: RegisterActions.Submit) {
		const state = ctx.getState();
		return Register(state.registrationURL, state.registrationToken).then((r) => {
			if (r.length > 0) {
				ShowDialog("Error connecting to remote server", r);
			}
			ctx.dispatch([
				new FetchIsRegistered(),
				new FetchWorkstationName(),
			]);
		});
	}

	@Action(RegisterActions.SetRegistrationInfo)
	SetRegistrationInfo(ctx: StateContext<RegisterModel>, action: RegisterActions.SetRegistrationInfo) {
		ctx.patchState({registrationToken: action.payload.registrationToken, registrationURL: action.payload.url});
	}

}
