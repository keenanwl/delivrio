import {Component, OnInit} from '@angular/core';
import {ActivatedRoute} from "@angular/router";
import {Store} from "@ngxs/store";
import {LoginModel, LoginState} from "../login.ngxs";
import {Observable} from "rxjs";
import {LoginActions} from "../login.actions";
import {Paths} from "../../app-routing.module";

@Component({
	selector: 'app-password-reset',
	templateUrl: './password-reset.component.html',
	styleUrls: ['./password-reset.component.scss']
})
export class PasswordResetComponent implements OnInit {

	paths = Paths;
	login$: Observable<LoginModel>;

	constructor(public route: ActivatedRoute, private store: Store) {
		this.login$ = store.select(LoginState.getLoginState);
	}

	ngOnInit(): void {
		this.route.queryParams.subscribe(params => {
			if (typeof params['otk'] !== 'undefined') {
				this.store.dispatch(new LoginActions.SetOtk(params['otk'] + ''));
			} else {
				this.store.dispatch(new LoginActions.ResetStage(`otk-not-present`));
			}
		});
	}

	saveNewPassword(password: string) {
		this.store.dispatch(new LoginActions.ResetPassword(password));
	}

}
