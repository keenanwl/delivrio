import {Component, OnInit} from '@angular/core';
import {Observable} from "rxjs";
import {Store} from "@ngxs/store";
import {MatDialog} from "@angular/material/dialog";
import {UserSeatsListModel, UserSeatsListState} from "./user-seats-list.ngxs";
import {UserSeatsListActions} from "./user-seats-list.actions";
import FetchUserSeatsList = UserSeatsListActions.FetchUserSeatsList;
import {AppActions} from "../../../app.actions";
import AppChangeRoute = AppActions.AppChangeRoute;
import {Paths} from 'src/app/app-routing.module';

@Component({
	selector: 'app-user-seats-list',
	templateUrl: './user-seats-list.component.html',
	styleUrls: ['./user-seats-list.component.scss']
})
export class UserSeatsListComponent implements OnInit {

	userSeatsList$: Observable<UserSeatsListModel>;

	displayedColumns: string[] = [
		'name',
		'createdAt',
		'seatGroup',
		'isAccountOwner',
		'actions',
	];

	constructor(
		private store: Store,
		private dialog: MatDialog,
	) {
		this.userSeatsList$ = store.select(UserSeatsListState.get);
	}

	ngOnInit(): void {
		this.store.dispatch(new FetchUserSeatsList());
	}

	createNew() {

	}

	editOption(id: string) {
		this.store.dispatch([
			new AppChangeRoute({path: Paths.SETTINGS_USERS_EDIT, queryParams: {id}}),
		]);
	}

	protected readonly Paths = Paths;
}
