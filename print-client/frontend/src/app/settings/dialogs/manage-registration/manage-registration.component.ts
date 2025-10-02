import {Component, OnDestroy, OnInit} from '@angular/core';
import {MatButton} from "@angular/material/button";
import {MatIcon} from "@angular/material/icon";
import {MatDialogRef} from "@angular/material/dialog";
import {Store} from "@ngxs/store";
import {SettingsActions} from "../../settings.actions";
import FetchRegistrationData = SettingsActions.FetchRegistrationData;
import {AsyncPipe, NgIf} from "@angular/common";
import {SettingsModel, SettingsState} from "../../settings.ngxs";
import {Observable} from "rxjs";
import ClearRegistrationInfo = SettingsActions.ClearRegistrationInfo;
import RemoveAllRegistrationData = SettingsActions.RemoveAllRegistrationData;
import {MatDivider} from "@angular/material/divider";

@Component({
	selector: 'app-manage-registration',
	standalone: true,
	imports: [
		MatButton,
		MatIcon,
		AsyncPipe,
		NgIf,
		MatDivider,
	],
	templateUrl: './manage-registration.component.html',
	styleUrl: './manage-registration.component.scss'
})
export class ManageRegistrationComponent implements OnInit, OnDestroy {
	settings$: Observable<SettingsModel>;

	constructor(private ref: MatDialogRef<any>, private store: Store) {
		this.settings$ = store.select(SettingsState.get);
	}

	ngOnInit() {
		this.store.dispatch(new FetchRegistrationData());
	}

	ngOnDestroy() {
		this.store.dispatch(new ClearRegistrationInfo());
	}

	removeAllData() {
		this.store.dispatch(new RemoveAllRegistrationData());
		this.close();
	}

	close() {
		this.ref.close();
	}

}
