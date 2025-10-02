import {Component, OnDestroy, OnInit} from '@angular/core';
import {Observable, Subscription} from "rxjs";
import {Store} from "@ngxs/store";
import {ConnectionsListModel, ConnectionsListState} from "./connections-list.ngxs";
import {ConnectionsListActions} from "./connections-list.actions";
import FetchConnectionsList = ConnectionsListActions.FetchConnectionsList;
import {AppActions} from "../../../app.actions";
import AppChangeRoute = AppActions.AppChangeRoute;
import Clear = ConnectionsListActions.Clear;
import {Paths} from "../../../app-routing.module";

@Component({
	selector: 'app-connections-list',
	templateUrl: './connections-list.component.html',
	styleUrls: ['./connections-list.component.scss']
})
export class ConnectionsListComponent implements OnInit, OnDestroy {

	connectionList$: Observable<ConnectionsListModel>;
	subscriptions$: Subscription[] = [];
	editPath = Paths.SETTINGS_CONNECTIONS_EDIT;

	constructor(private store: Store) {
		this.connectionList$ = store.select(ConnectionsListState.get);
	}

	ngOnInit(): void {
		this.store.dispatch(new FetchConnectionsList());
	}

	editConnection(id: string) {
		this.store.dispatch(new AppChangeRoute({path: Paths.SETTINGS_CONNECTIONS_EDIT, queryParams: {id}}));
	}

	ngOnDestroy(): void {
		this.subscriptions$.forEach((s) => s.unsubscribe());
		this.store.dispatch(new Clear());
	}

}
