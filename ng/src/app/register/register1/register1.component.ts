import {Component, OnInit} from '@angular/core';
import {Actions, ofActionDispatched, Store} from "@ngxs/store";
import {Observable, Subject, Subscription} from "rxjs";
import {AppModel, AppState} from "../../app.ngxs";
import {ActivatedRoute} from "@angular/router";
import {debounceTime, distinctUntilChanged} from "rxjs/operators";
import {Register1Actions} from "./register1.actions";
import FetchMembershipInfo = Register1Actions.FetchMembershipInfo;
import {RegisterModel, Register1State} from "./register1.ngxs";
import LookupAddress = Register1Actions.LookupAddress;
import SubmitRegistrationInfo = Register1Actions.SubmitRegistrationInfo;
import SetMembershipId = Register1Actions.SetMembershipId;
import SetInvalidParams = Register1Actions.SetInvalidParams;
import {AppActions} from "../../app.actions";
import AppChangeRoute = AppActions.AppChangeRoute;
import {FormControl, FormGroup} from "@angular/forms";
import SetLoggedInUser = AppActions.SetLoggedInUser;
import {LoginActions} from "../../login/login.actions";

@Component({
	selector: 'app-register1',
	templateUrl: './register1.component.html',
	styleUrls: ['./register1.component.scss'],
})
export class Register1Component implements OnInit {

	register$: Observable<RegisterModel>;
	app$: Observable<AppModel>;

	addressChanged: Subject<string> = new Subject<string>();
	subscriptions: Subscription[] = [];

	registerForm = new FormGroup({
		name: new FormControl('', {nonNullable: true}),
		surname: new FormControl('', {nonNullable: true}),
		companyName: new FormControl('', {nonNullable: true}),
		phoneNumber: new FormControl('', {nonNullable: true}),
		vatNumber: new FormControl('', {nonNullable: true}),
		email: new FormControl('', {nonNullable: true}),
		password: new FormControl('', {nonNullable: true}),
		repeatPassword: new FormControl('', {nonNullable: true}),
	});

	repeatPasswordControl = new FormControl("");

	constructor(
		public route: ActivatedRoute,
		private store: Store,
		private actions$: Actions,
	) {
	  this.register$ = store.select(Register1State.state);
	  this.app$ = store.select(AppState.get);
	}

	ngOnInit(): void {

		// Clear any lingering tokens.
		this.store.dispatch(new LoginActions.ClearAllLoginData());

		this.subscriptions.push(this.actions$.pipe(ofActionDispatched(SetLoggedInUser)).subscribe(() => {
			const app = this.store.selectSnapshot(AppState.get);
			this.store.dispatch(new AppChangeRoute({path: `/register/2`, queryParams: {}}));
		}));

		this.subscriptions.push(this.route.queryParams.subscribe(params => {
			if (typeof params['membership_id'] !== 'undefined') {
				this.store.dispatch([
					new SetMembershipId(params['membership_id']),
					new FetchMembershipInfo(),
				]);
			} else {
				this.store.dispatch(new SetInvalidParams(true));
			}
		}));

		this.subscriptions.push(this.addressChanged
			.pipe(
				debounceTime(500),
				distinctUntilChanged(),
			).subscribe(model => this.store.dispatch(new LookupAddress(model))));

	}

	lookup(query: string) {
		this.addressChanged.next(query);
	}

	next() {
		const values = this.registerForm.getRawValue();
		if (values.password === "" || (values.password !== values.repeatPassword)) {
			this.registerForm.controls.password.markAsTouched();
			this.registerForm.controls.password.setErrors({0: `Passwords don't match`})
			this.registerForm.controls.repeatPassword!.setErrors({0: `Passwords don't match`})
			this.registerForm.controls.repeatPassword!.markAsTouched()
			return
		}

		this.store.dispatch([
			new SubmitRegistrationInfo({userInput: values, tenantInput: {name: values.companyName, vatNumber: values.vatNumber}}),
		]);
	}

}
