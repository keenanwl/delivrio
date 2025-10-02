import {Component, OnInit} from '@angular/core';
import {Store} from "@ngxs/store";
import {LoginActions} from "../login/login.actions";
import ClearAllLoginData = LoginActions.ClearAllLoginData;

@Component({
	selector: 'app-register',
	templateUrl: './register.component.html',
	styleUrls: ['./register.component.scss']
})
export class RegisterComponent implements OnInit {

	constructor(private store: Store) { }

	ngOnInit(): void {
		this.store.dispatch(new ClearAllLoginData());
	}

}
