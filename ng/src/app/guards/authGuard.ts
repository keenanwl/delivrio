import {Injectable} from '@angular/core';
import {CanActivate} from '@angular/router';
import {Actions, ofActionDispatched, Store} from '@ngxs/store';
import {LoginState} from '../login/login.ngxs';
import {AppActions} from "../app.actions";
import AppChangeRoute = AppActions.AppChangeRoute;
import {AppState} from "../app.ngxs";
import {map} from "rxjs/operators";
import FetchLoggedInUser = AppActions.FetchLoggedInUser;
import SetLoggedInUser = AppActions.SetLoggedInUser;
import {LoginActions} from "../login/login.actions";
import ClearAllLoginData = LoginActions.ClearAllLoginData;
import {Paths} from '../app-routing.module';

@Injectable()
export class AuthGuard implements CanActivate {
	constructor(private store: Store, private actions$: Actions) {}

	canActivate() {
		const login = this.store.selectSnapshot(LoginState.getLoginState);
		const isLoggedIn = login.jwt.length > 0;

		const app = this.store.selectSnapshot(AppState.get);
		const isProfileLoaded = app.my_ids.my_pulid.length > 0;

		if (!isLoggedIn) {
			this.store.dispatch([
				new ClearAllLoginData(),
				new AppChangeRoute({path: Paths.LOGIN, queryParams: {}}),
			]);
			return false;
		}

		if (!isProfileLoaded) {
			this.store.dispatch(new FetchLoggedInUser());
			return this.actions$
				.pipe(ofActionDispatched(SetLoggedInUser), map(() => true));
		}

		return isLoggedIn;
	}

}
