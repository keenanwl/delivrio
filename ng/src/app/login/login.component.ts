import {Component, OnDestroy, OnInit} from '@angular/core';
import {Actions, ofActionDispatched, Store} from "@ngxs/store";
import {LoginModel, LoginState} from "./login.ngxs";
import {Observable, Subscription} from "rxjs";
import {LoginActions} from "./login.actions";
import {AppActions} from "../app.actions";
import AppChangeRoute = AppActions.AppChangeRoute;
import SetLoggedInUser = AppActions.SetLoggedInUser;
import {Paths} from '../app-routing.module';

@Component({
	selector: 'app-login',
	templateUrl: './login.component.html',
	styleUrls: ['./login.component.scss']
})
export class LoginComponent implements OnInit, OnDestroy {

	paths = Paths;
	login$: Observable<LoginModel>;
	subscriptions$: Subscription[] = [];

	constructor(
		public store: Store,
		private actions$: Actions,
	) {
		this.login$ = store.select(LoginState.getLoginState);
	}

	ngOnInit(): void {
		this.subscriptions$.push(this.actions$.pipe(ofActionDispatched(SetLoggedInUser)).subscribe(() => {
			this.store.dispatch(new AppChangeRoute({path: Paths.DASHBOARD, queryParams: {}}));
		}));
	}

	fireLogin(email: string, password: string) {
		this.store.dispatch(new LoginActions.Login({email, password}))
	}

	ngOnDestroy(): void {
		this.subscriptions$.map((s) => s.unsubscribe());
	}

}
