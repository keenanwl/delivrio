import { Component } from '@angular/core';

import { DialogService } from './services/dialog.service';
import {BrowserOpenURL} from "../../../wailsjs/runtime";
import { Store } from '@ngxs/store';
import {RegisterActions} from "./register.actions";
import SetRegistrationToken = RegisterActions.SetRegistrationInfo;
import Submit = RegisterActions.Submit;

@Component({
	selector: 'app-home',
	templateUrl: './register.component.html',
	styleUrls: ['./register.component.scss']
})
export class RegisterComponent {

	constructor(
	  private store: Store,
	  private dialogService: DialogService) {

	}

	openDialog(): void {
		this.dialogService.open();
	}

	openURL(link: string) {
		BrowserOpenURL(link);
	}

	submit(url: string, registrationToken: string) {
		this.store.dispatch([
			new SetRegistrationToken({url, registrationToken}),
		]);
		this.store.dispatch(new Submit());
	}

}
