
export namespace LoginActions {
	export type loginStage = 'token'| 'email-response' | 'valid';
	export type requestPasswordStage = 'request' | 'success' | 'error';
	export type resetPasswordStage = 'reset' | 'otk-not-present' | 'error' | 'success';
	export class Login {
		static readonly type = '[Login] login';
		constructor(public payload: {email: string, password: string}) {}
	}
	export class ChangeAutoLoginInfo {
		static readonly type = '[Login] change respondent code';
		constructor(public payload: {token: string; email: string}) {}
	}
	export class RequestEmail {
		static readonly type = '[Login] request email';
		constructor(public payload: string) {}
	}
	export class ResetPassword {
		static readonly type = '[Login] reset password';
		constructor(public payload: string) {}
	}
	export class SetOtk {
		static readonly type = '[Login] set otk';
		constructor(public payload: string) {}
	}
	export class ResetStage {
		static readonly type = '[Login] reset stage';
		constructor(public payload: resetPasswordStage) {}
	}
	export class SetJwt {
		static readonly type = '[Login] set jwt';
		constructor(public payload: string) {}
	}
	export class RefreshJwt {
		static readonly type = '[Login] refresh jwt';
	}
	export class SetRequestStage {
		static readonly type = '[Login] set request stage';
		constructor(public payload: requestPasswordStage) {}
	}
	export class SetMessage {
		static readonly type = '[Login] set message';
		constructor(public payload: string) {}
	}
	export class ClearAllLoginData {
		static readonly type = '[Login] clear all login data (all cookies)';
	}
}
