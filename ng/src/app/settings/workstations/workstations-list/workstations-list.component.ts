import {Component, OnInit} from '@angular/core';
import {Observable} from "rxjs";
import {WorkstationsModel, WorkstationsListState} from "./workstations-list.ngxs";
import {Store} from "@ngxs/store";
import {MatDialog} from "@angular/material/dialog";
import {NewWorkstationComponent} from "../dialogs/new-workstation/new-workstation.component";
import {WorkstationsListActions} from "./workstations-list.actions";
import FetchWorkstations = WorkstationsListActions.FetchWorkstations;
import {AppActions} from "../../../app.actions";
import AppChangeRoute = AppActions.AppChangeRoute;
import {Paths} from "../../../app-routing.module";
import ToggleArchived = WorkstationsListActions.ToggleArchived;

@Component({
	selector: 'app-workstations-list',
	templateUrl: './workstations-list.component.html',
	styleUrls: ['./workstations-list.component.scss']
})
export class WorkstationsListComponent implements OnInit {
	workstations$: Observable<WorkstationsModel>;
	displayedColumns: string[] = [
		'name',
		'printers',
		'selectedUser',
		'status',
	];

	constructor(
		private store: Store,
		private dialog: MatDialog,
	) {
		this.workstations$ = store.select(WorkstationsListState.get);
	}

	ngOnInit(): void {
		this.store.dispatch(new FetchWorkstations());
	}

	addNew() {
		const ref = this.dialog.open(NewWorkstationComponent);
	}

	edit(id: string) {
		this.store.dispatch(new AppChangeRoute({path: Paths.SETTINGS_WORKSTATIONS_EDIT, queryParams: {id}}));
	}

	toggleArchived() {
		this.store.dispatch(new ToggleArchived());
	}
}
