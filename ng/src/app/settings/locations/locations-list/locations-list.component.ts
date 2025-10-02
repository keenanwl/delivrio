import {Component, OnDestroy, OnInit} from '@angular/core';
import {Observable, Subscription} from "rxjs";
import {Store} from "@ngxs/store";
import {LocationsListModel, LocationsListState} from "./locations-list.ngxs";
import {Paths} from "../../../app-routing.module";
import {AppActions} from "../../../app.actions";
import AppChangeRoute = AppActions.AppChangeRoute;
import {LocationsListActions} from "./locations-list.actions";
import FetchLocationsList = LocationsListActions.FetchLocationsList;
import Reset = LocationsListActions.Reset;
import CreateNewLocation = LocationsListActions.CreateNewLocation;

export interface Person {
	name: string;
	lastName: string;
	age: number;
}

@Component({
	selector: 'app-locations-list',
	templateUrl: './locations-list.component.html',
	styleUrls: ['./locations-list.component.scss']
})
export class LocationsListComponent implements OnInit, OnDestroy {

	locations$: Observable<LocationsListModel>;

	displayedColumns: string[] = [
		'name',
		'address',
		'tags',
	];

	subscriptions: Subscription[] = [];

	name: number = 0;

	constructor(
		private store: Store,
	) {
		this.locations$ = store.select(LocationsListState.get);
	}

	ngOnInit() {
		this.store.dispatch(new FetchLocationsList());
	}

	ngOnDestroy() {
		this.store.dispatch(new Reset());
	}

	addNew() {
		this.store.dispatch(new CreateNewLocation());
	}

	edit(id: string) {
		this.store.dispatch(new AppChangeRoute({path: Paths.SETTINGS_LOCATION_EDIT, queryParams: {id}}));
	}

}



