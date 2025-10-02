import {Component, OnInit} from '@angular/core';
import {AsyncPipe, NgForOf, NgIf} from "@angular/common";
import {Observable} from "rxjs";
import {SettingsModel, SettingsState} from "../../settings.ngxs";
import {Store} from "@ngxs/store";
import {MatProgressSpinnerModule} from "@angular/material/progress-spinner";
import {MatButtonModule} from "@angular/material/button";
import {MatIconModule} from "@angular/material/icon";
import {MatDialogRef} from "@angular/material/dialog";
import {SettingsActions} from "../../settings.actions";
import SaveNetworkPrinter = SettingsActions.SaveNetworkPrinter;
import RefreshPrinters = SettingsActions.RefreshPrinters;
import {MatDividerModule} from "@angular/material/divider";

@Component({
	selector: 'app-network-search-results',
	standalone: true,
	imports: [
		AsyncPipe,
		NgIf,
		MatProgressSpinnerModule,
		NgForOf,
		MatButtonModule,
		MatIconModule,
		MatDividerModule
	],
	templateUrl: './network-search-results.component.html',
	styleUrl: './network-search-results.component.scss'
})
export class NetworkSearchResultsComponent implements OnInit {
	settings$: Observable<SettingsModel>;

	constructor(
		private store: Store,
		private dialog: MatDialogRef<any>,
	) {
		this.settings$ = store.select(SettingsState.get);
	}

	ngOnInit() {
	}

	addNetworkLocation(host: string, port: string) {
		this.store.dispatch([
			new SaveNetworkPrinter(host + ":" + port),
			new RefreshPrinters(),
		]);
	}

	close() {
		this.dialog.close();
	}
}
