import {Component, OnInit} from '@angular/core';
import {Select, Store} from "@ngxs/store";
import {LoginModel, LoginState} from "../login.ngxs";
import {Observable} from "rxjs";
import {LoginActions} from "../login.actions";
import {Paths} from "../../app-routing.module";

@Component({
	selector: 'app-request-password-reset',
	templateUrl: './request-password-reset.component.html',
	styleUrls: ['./request-password-reset.component.scss']
})
export class RequestPasswordResetComponent implements OnInit {

	login$: Observable<LoginModel>;

	constructor(private store: Store) {
		this.login$ = store.select(LoginState.getLoginState);
	}

	ngOnInit(): void {
	}

	requestResetEmail(email: string) {
		this.store.dispatch(new LoginActions.RequestEmail(email));
	}

	protected readonly Paths = Paths;
}

