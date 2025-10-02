import {Action, Selector, State, StateContext} from '@ngxs/store';
import {Injectable} from "@angular/core";
import { HttpErrorResponse } from "@angular/common/http";
import {Register1Actions} from "./register1.actions"
import {LoginActions} from "../../login/login.actions";
import Login = LoginActions.Login;
import {GraphQLError} from "graphql/index";
import {CreateUserInput} from "../../../generated/graphql";
import {RegisterService} from "./register1.service.";
import {AppActions} from "../../app.actions";
import ShowGlobalSnackbar = AppActions.ShowGlobalSnackbar;

export interface RegisterModel {
	register1EditForm: {
		model: CreateUserInput | undefined;
		dirty: boolean;
		status: string;
		errors: readonly GraphQLError[];
	},

	invalid_params: boolean,
	invalid_message: string;
	loading: boolean;
}

const defaultState: RegisterModel = {
	register1EditForm: {
		model: undefined,
		dirty: false,
		status: '',
		errors: [],
	},

	invalid_message: "",
	invalid_params: false,
	loading: false,
};

@Injectable()
@State<RegisterModel>({
	name: 'register',
	defaults: defaultState,
})
export class Register1State {

	constructor(
		private userCreate: RegisterService,
	) {}

	@Selector()
	static state(state: RegisterModel) {
		return state;
	}

	@Action(Register1Actions.SetInvalidParams)
	SetInvalidParams(ctx: StateContext<RegisterModel>, action: Register1Actions.SetInvalidParams) {
		ctx.patchState({
			invalid_params: action.payload,
			loading: false,
		});
	}

	@Action(Register1Actions.ClearAllRegister)
	ClearAllLoginData(ctx: StateContext<RegisterModel>) {
		ctx.patchState({
			...defaultState
		});
	}

	@Action(Register1Actions.SetInvalidMessage)
	SetInvalidMessage(ctx: StateContext<RegisterModel>, action: Register1Actions.SetInvalidMessage) {
		ctx.patchState({
			invalid_message: action.payload,
		});
	}

	@Action(Register1Actions.SubmitRegistrationInfo)
	SubmitRegistrationInfo(ctx: StateContext<RegisterModel>, action: Register1Actions.SubmitRegistrationInfo) {
		ctx.patchState({
		    loading: true,
		});
		return this.userCreate.initialRegistration(action.payload.userInput, action.payload.tenantInput).subscribe({
			next: (r) => {
				ctx.dispatch(new Login({email: action.payload.userInput.email, password: action.payload.userInput.password || ''}));
			},
			error: (e: HttpErrorResponse) => {
		        ctx.dispatch(new ShowGlobalSnackbar(e.error.message));
			}});

	}

}
