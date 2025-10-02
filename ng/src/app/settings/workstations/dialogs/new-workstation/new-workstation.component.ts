import {Component, OnDestroy, OnInit} from '@angular/core';
import {Observable, Subscription} from "rxjs";
import {WorkstationsModel, WorkstationsListState} from "../../workstations-list/workstations-list.ngxs";
import {Actions, ofActionCompleted, Store} from "@ngxs/store";
import {WorkstationsListActions} from "../../workstations-list/workstations-list.actions";
import CreateNewWorkstation = WorkstationsListActions.CreateNewWorkstation;
import SetRegistrationToken = WorkstationsListActions.SetRegistrationToken;
import {MatDialogRef} from "@angular/material/dialog";
import FetchWorkstations = WorkstationsListActions.FetchWorkstations;
import {WorkstationDeviceType} from "../../../../../generated/graphql";

@Component({
	selector: 'app-new-workstation',
	templateUrl: './new-workstation.component.html',
	styleUrls: ['./new-workstation.component.scss']
})
export class NewWorkstationComponent implements OnInit, OnDestroy {
	workstations$: Observable<WorkstationsModel>;
	subscriptions$: Subscription[] = [];
	showToken = false;

	constructor(
		private store: Store,
		private actions$: Actions,
		private dialogRef: MatDialogRef<void>,
	) {
		this.workstations$ = store.select(WorkstationsListState.get);
	}

	ngOnInit() {
		this.subscriptions$.push(this.actions$.pipe(ofActionCompleted(SetRegistrationToken))
			.subscribe(() => {
				this.showToken = true;
				this.store.dispatch(new FetchWorkstations());
			}));
	}

	ngOnDestroy() {
		this.subscriptions$.forEach((s) => s.unsubscribe());
	}

	next(name: string) {
		// Hardcoded until app goes live
		this.store.dispatch(new CreateNewWorkstation({name, deviceType: WorkstationDeviceType.LabelStation}));
	}

	close() {
		this.dialogRef.close();
	}
}
