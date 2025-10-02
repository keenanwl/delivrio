import {Action, Selector, State, StateContext} from '@ngxs/store';
import {CookieService} from 'ngx-cookie-service';
import {Injectable} from "@angular/core";
import {LoginService} from "./login.service";
import {RequestPasswordResetService} from "./request-password-reset/request-password-reset.service";
import {PasswordResetService} from "./password-reset/password-reset.service";
import {AppActions} from "../app.actions";
import AppChangeRoute = AppActions.AppChangeRoute;
import FetchAppConfig = AppActions.FetchAppConfig;
import {LoginActions} from './login.actions';
import ShowGlobalSnackbar = AppActions.ShowGlobalSnackbar;
import {Paths} from "../app-routing.module";

export interface LoginModel {
	respondentCode: string;
	respondentCodeError: string;
	jwt: string;
	otk: string;
	nextPasswordBase64: string;
	email: string;
	resetStage: LoginActions.resetPasswordStage;
	loading: boolean;
	message: string;
	stage: LoginActions.loginStage;
	requestStage: LoginActions.requestPasswordStage;
}

const defaultState: LoginModel = {
	respondentCode: '',
	respondentCodeError: '',
	jwt: '',
	otk: '',
	nextPasswordBase64: '',
	email: '',
	resetStage: 'reset',
	loading: false,
	message: '',
	stage: 'token',
	requestStage: "request",
};

@Injectable()
@State<LoginModel>({
	name: 'login',
	defaults: defaultState,
})
export class LoginState {

	constructor(
		public loginService: LoginService,
		public resetPasswordService: PasswordResetService,
		public requestResetPasswordService: RequestPasswordResetService,
		private authService: LoginService,
		private cookieService: CookieService,
	) {}

	@Selector()
	static getLoginState(state: LoginModel) {
		return state;
	}

	@Action(LoginActions.Login)
	Login(ctx: StateContext<LoginModel>, action: LoginActions.Login) {
		return this.authService.postLogin(action.payload.email, action.payload.password)
			.subscribe({
				next: (r) => {
					const oneMonth = new Date();
					oneMonth.setMonth(oneMonth.getMonth() + 1);

					this.cookieService.set('token', r.token!, oneMonth);

					ctx.dispatch([
						new LoginActions.SetJwt(r.token!),
						new AppActions.FetchLoggedInUser(),
					]);
				},
				error: (e) => ctx.dispatch(new ShowGlobalSnackbar(`Your login credentials were not accepted`)),
			});
	}

	@Action(LoginActions.ChangeAutoLoginInfo)
	ChangeAutoLoginInfo(ctx: StateContext<LoginModel>, action: LoginActions.ChangeAutoLoginInfo) {
		ctx.patchState({
			respondentCode: action.payload.token,
			email: action.payload.email,
		});
	}

	@Action(LoginActions.ResetPassword)
	ResetPassword(ctx: StateContext<LoginModel>, action: LoginActions.ResetPassword) {
		ctx.patchState({
			loading: true,
		});

		return this.resetPasswordService.postNextPassword({new_password: action.payload, otk: ctx.getState().otk})
			.subscribe((r) => {
				if (!r.success) {
					ctx.dispatch([new LoginActions.ResetStage(`error`), new LoginActions.SetMessage(r.message)])
				} else {
					ctx.dispatch([new LoginActions.ResetStage(`success`), new LoginActions.SetMessage('')])
				}
			});

	}

	@Action(LoginActions.RefreshJwt)
	RefreshJwt(ctx: StateContext<LoginModel>, action: LoginActions.RefreshJwt) {
		ctx.patchState({
			loading: true
		});

		// Refreshing the JWT is not easily supported in the new Go router.
		return this.loginService.refreshJwt()
			.subscribe({next:(d) => {

				let error = '';
				let stage: LoginActions.loginStage = 'valid';

				const oneMonth = new Date();
				oneMonth.setMonth(oneMonth.getMonth() + 1);

				this.cookieService.set(
					'token',
					d.token,
					oneMonth,
					undefined,
					undefined,
					true,
					'Strict'
					);

				ctx.patchState({
					respondentCodeError: error,
					stage: stage,
					jwt: d.token,
					loading: false
				});

				ctx.dispatch([new AppChangeRoute({path: '/settings', queryParams: {}}), new FetchAppConfig()]);

				window.location = window.location;

			}, error: (d) => {

				ctx.dispatch([new LoginActions.ClearAllLoginData(), new AppChangeRoute({path: Paths.LOGIN, queryParams: {}})]);

			}});

	}

	@Action(LoginActions.RequestEmail)
	RequestEmail(ctx: StateContext<LoginModel>, action: LoginActions.RequestEmail) {
		ctx.patchState({
			loading: true,
		});

		return this.requestResetPasswordService.postRequestEmail(action.payload)
			.subscribe((resp) => {
				ctx.dispatch(new LoginActions.SetRequestStage(resp.success ? 'success' : 'error'));
			});
	}

	@Action(LoginActions.SetOtk)
	SetOtk(ctx: StateContext<LoginModel>, action: LoginActions.SetOtk) {
		ctx.patchState({
			otk: action.payload,
		});
	}

	@Action(LoginActions.SetJwt)
	SetJwt(ctx: StateContext<LoginModel>, action: LoginActions.SetJwt) {
		ctx.patchState({
			jwt: action.payload,
			stage: 'valid',
		});
	}

	@Action(LoginActions.SetMessage)
	SetMessage(ctx: StateContext<LoginModel>, action: LoginActions.SetMessage) {
		ctx.patchState({
			message: action.payload,
		});
	}

	@Action(LoginActions.ResetStage)
	OtkNotPresent(ctx: StateContext<LoginModel>, action: LoginActions.ResetStage) {
		ctx.patchState({
			resetStage: action.payload,
		});
	}

	@Action(LoginActions.SetRequestStage)
	SetRequestStage(ctx: StateContext<LoginModel>, action: LoginActions.SetRequestStage) {
		ctx.patchState({
			requestStage: action.payload,
		});
	}

	@Action(LoginActions.ClearAllLoginData)
	ClearAllLoginData(ctx: StateContext<LoginModel>, action: LoginActions.ClearAllLoginData) {
		this.cookieService.deleteAll(undefined, undefined, true);

		ctx.patchState({
			...defaultState
		});

		//ctx.dispatch(new AppChangeRoute({path: '/login', queryParams: {}}));
	}

}
