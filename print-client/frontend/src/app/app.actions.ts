import {Breakpoint} from "./app.ngxs";

export interface GQLFieldError {
	message: string;
	path: Array<string | number>;
}

export namespace AppActions {
	export class AppChangeRoute {
		static readonly type = '[App] change route';
		constructor(public payload: {path: string; queryParams: {[key: string]: any}}) {}
	}
	// Differs from above by only changing URL
	// in store. Then calls the .back() to maintain
	// browser navigation. Does not get listened to.
	export class AppSetRoute {
		static readonly type = '[App] set route';
		constructor(public payload: string) {}
	}
	export class AppGoBack {
		static readonly type = '[App] go back';
	}
	export class UpdateBreakpoint {
		static readonly type = '[App] update breakpoint';
		constructor(public payload: Breakpoint) {}
	}
	export class ShowGlobalSnackbar {
		static readonly type = '[App] show global snackbar';
		constructor(public payload: string) {}
	}
	export class Logout {
		static readonly type = '[App] logout';
	}
	export class FetchIsRegistered {
		static readonly type = '[App] fetch is registered';
	}
	export class SetIsRegistered {
		static readonly type = '[App] set is registered';
		constructor(public payload: boolean) {}
	}
	export class FetchWorkstationName {
		static readonly type = '[App] fetch workstation name';
	}
	export class SetWorkstationName {
		static readonly type = '[App] set workstation name';
		constructor(public payload: string) {}
	}
}
