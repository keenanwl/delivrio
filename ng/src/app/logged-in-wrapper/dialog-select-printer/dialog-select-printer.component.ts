import {Component, OnInit} from '@angular/core';
import {Store} from "@ngxs/store";
import {AppActions} from "../../app.actions";
import {Observable} from "rxjs";
import {AppModel, AppState} from "../../app.ngxs";
import {MatDialogRef} from "@angular/material/dialog";
import {UserPickupDay, WorkstationStatus} from "../../../generated/graphql";
import FetchSelectableWorkstations = AppActions.FetchSelectableWorkstations;
import SaveSelectedWorkstation = AppActions.SaveSelectedWorkstation;

@Component({
  selector: 'app-dialog-select-printer',
  templateUrl: './dialog-select-printer.component.html',
  styleUrls: ['./dialog-select-printer.component.scss']
})
export class DialogSelectPrinterComponent implements OnInit {

	app$: Observable<AppModel>;

	today: UserPickupDay = UserPickupDay.Today;
	tomorrow: UserPickupDay = UserPickupDay.Tomorrow;
	in2Days: UserPickupDay = UserPickupDay.In_2Days;
	in3Days: UserPickupDay = UserPickupDay.In_3Days;
	in4Days: UserPickupDay = UserPickupDay.In_4Days;
	in5Days: UserPickupDay = UserPickupDay.In_5Days;

	constructor(private store: Store, private ref: MatDialogRef<any>) {
		this.app$ = store.select(AppState.get);
	}

	ngOnInit() {
		this.store.dispatch(new FetchSelectableWorkstations())
	}

	close() {
		this.ref.close();
	}

	save(id: string, pickupDay: UserPickupDay) {
		this.store.dispatch(new SaveSelectedWorkstation({workstationID: id, pickupDay: pickupDay}));
		this.close();
	}

	protected readonly WorkstationStatus = WorkstationStatus;
}
