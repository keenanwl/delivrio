import { Component, OnInit } from '@angular/core';
import {Observable} from "rxjs";
import {SettingsModel, SettingsState} from "./settings.ngxs";
import {Actions, ofActionDispatched, Store} from "@ngxs/store";
import {SettingsActions} from "./settings.actions";
import Refresh = SettingsActions.Refresh;
import RefreshPrinters = SettingsActions.RefreshPrinters;
import TogglePrinter = SettingsActions.TogglePrinter;
import {AppActions} from "../app.actions";
import AppChangeRoute = AppActions.AppChangeRoute;
import SaveNetworkPrinter = SettingsActions.SaveNetworkPrinter;
import ArchivePrinter = SettingsActions.ArchivePrinter;
import SaveSelectedDevice = SettingsActions.SaveSelectedDevice;
import SearchNetworkPrinters = SettingsActions.SearchNetworkPrinters;
import {MatDialog} from "@angular/material/dialog";
import {NetworkSearchResultsComponent} from "./dialogs/network-search-results/network-search-results.component";
import {ManageRegistrationComponent} from "./dialogs/manage-registration/manage-registration.component";

@Component({
	selector: 'app-settings',
	templateUrl: './settings.component.html',
	styleUrls: ['./settings.component.scss']
})
export class SettingsComponent implements OnInit {
	settings$: Observable<SettingsModel>;

	constructor(
		private store: Store,
		private dialog: MatDialog,
		private actions: Actions,
	) {
		this.settings$ = store.select(SettingsState.get);
	}

	ngOnInit() {
		this.refresh();
	}

	goHome() {
		this.store.dispatch(new AppChangeRoute({path: "dashboard", queryParams: {}}))
	}

	refresh() {
		this.store.dispatch([
			new Refresh(),
			new RefreshPrinters(),
		]);
	}

	togglePrinter(name: string, active: boolean) {
		this.store.dispatch(new TogglePrinter({name, active}));
	}

	saveNetworkPrinter(addr: string) {
		// Probably needs to await the response
		this.store.dispatch([new SaveNetworkPrinter(addr), new RefreshPrinters()]);
	}

	archivePrinter(evt: MouseEvent, id: string) {
		evt.preventDefault()
		evt.cancelBubble = true;
		this.store.dispatch(new ArchivePrinter(id));
	}

	saveSelectedDevice(id: string) {
		this.store.dispatch(new SaveSelectedDevice(id));
	}

	searchNetworkPrinters() {
		this.store.dispatch(new SearchNetworkPrinters());
		this.dialog.open(NetworkSearchResultsComponent, {})
	}

	registrationDialog() {
		this.dialog.open(ManageRegistrationComponent);
	}
}
