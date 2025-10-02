import {Injectable} from '@angular/core';
import {CanActivate} from '@angular/router';
import {Actions, ofActionDispatched, Store} from '@ngxs/store';
import {LoginState} from '../login/login.ngxs';
import {AppActions} from "../app.actions";
import AppChangeRoute = AppActions.AppChangeRoute;
import {Paths} from "../app-routing.module";

@Injectable()
export class NonAuthGuard implements CanActivate {
	constructor(private store: Store, private actions$: Actions) {}

	canActivate() {
		const login = this.store.selectSnapshot(LoginState.getLoginState);
		const isLoggedIn = login.jwt.length > 0;

		if (isLoggedIn) {
			this.store.dispatch(new AppChangeRoute({path: Paths.ORDERS, queryParams: {}}));
			return false;
		}

		return true;
	}

}
