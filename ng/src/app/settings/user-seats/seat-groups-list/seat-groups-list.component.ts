import {Component, OnInit} from '@angular/core';
import {Observable} from "rxjs";
import {Store} from "@ngxs/store";
import {MatDialog} from "@angular/material/dialog";
import {SeatGroupsListActions} from "./seat-groups-list.actions";
import FetchSeatGroupsList = SeatGroupsListActions.FetchSeatGroupsList;
import {SeatGroupsListModel, SeatGroupsListState} from "./seat-groups-list.ngxs";
import {AppActions} from "../../../app.actions";
import AppChangeRoute = AppActions.AppChangeRoute;
import {Paths} from "../../../app-routing.module";

@Component({
	selector: 'app-seat-groups-list',
	templateUrl: './seat-groups-list.component.html',
	styleUrls: ['./seat-groups-list.component.scss']
})
export class SeatGroupsListComponent implements OnInit {

	seatGroupsList$: Observable<SeatGroupsListModel>;
	paths = Paths;

	displayedColumns: string[] = [
		'name',
		'userCount',
		'createdAt',
		'actions',
	];

	constructor(
		private store: Store,
		private dialog: MatDialog,
	) {
		this.seatGroupsList$ = store.select(SeatGroupsListState.get);
	}

	ngOnInit(): void {
		this.store.dispatch(new FetchSeatGroupsList());
	}

	editOption(id: string) {
		this.store.dispatch([
			new AppChangeRoute({path: Paths.SETTINGS_USERS_GROUPS_EDIT, queryParams: {id}}),
		]);
	}

}
