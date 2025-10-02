import {Component, OnInit} from '@angular/core';
import {Store} from "@ngxs/store";
import {Observable} from "rxjs";
import {ReturnPortalsListModel, ReturnPortalsListState} from "./return-portals-list.ngxs";
import {Paths} from "../../../app-routing.module";
import {ReturnPortalsListActions} from "./return-portals-list.actions";
import FetchReturnPortalsList = ReturnPortalsListActions.FetchReturnPortalsList;
import {MatDialog} from '@angular/material/dialog';
import {AddNewReturnPortalComponent} from "./dialog/add-new-return-portal/add-new-return-portal.component";
import {AppActions} from "../../../app.actions";
import AppChangeRoute = AppActions.AppChangeRoute;

@Component({
	selector: 'app-return-portals-list',
	templateUrl: './return-portals-list.component.html',
	styleUrls: ['./return-portals-list.component.scss']
})
export class ReturnPortalsListComponent implements OnInit {

	portals$: Observable<ReturnPortalsListModel>;

	displayedColumns: string[] = [
		'name',
		'horizon',
		'automaticallyAccept',
		'connection',
	];

	constructor(
		private store: Store,
		private dialog: MatDialog,
	) {
		this.portals$ = store.select(ReturnPortalsListState.get);
	}

	ngOnInit(): void {
		this.store.dispatch(new FetchReturnPortalsList());
	}

	addNew() {
		this.dialog.open(AddNewReturnPortalComponent);
	}

	delete(id: string) {

	}

	edit(id: string) {
		this.store.dispatch(new AppChangeRoute({path: Paths.SETTINGS_RETURN_PORTALS_EDIT, queryParams: {id}}))
	}
}
