import {Component, OnDestroy, OnInit} from '@angular/core';
import {Observable, Subscription} from "rxjs";
import {Store} from "@ngxs/store";
import {NotificationsListActions} from "./notifications-list.actions";
import FetchNotificationsList = NotificationsListActions.FetchNotificationsList;
import Clear = NotificationsListActions.Clear;
import {NotificationsListModel, NotificationsListState} from "./notifications-list.ngxs";
import {MatDialog} from "@angular/material/dialog";
import {CreateNotificationComponent} from "./dialogs/create-notification/create-notification.component";
import Toggle = NotificationsListActions.Toggle;

@Component({
	selector: 'app-notifications-list',
	templateUrl: './notifications-list.component.html',
	styleUrls: ['./notifications-list.component.scss']
})
export class NotificationsListComponent implements OnDestroy, OnInit {

	notifications$: Observable<NotificationsListModel>;
	subscriptions$: Subscription[] = []

	displayedColumns: string[] = [
		'name',
		'connection',
		'template',
		'trigger',
		'active',
	];

	constructor(
		private store: Store,
		private dialog: MatDialog,
	) {
		this.notifications$ = store.select(NotificationsListState.get);
	}

	ngOnDestroy(): void {
		this.subscriptions$.map((s) => s.unsubscribe());
		this.store.dispatch(new Clear());
	}

	ngOnInit(): void {
		this.store.dispatch(new FetchNotificationsList());
	}

	addNew() {
		this.dialog.open(CreateNotificationComponent);
	}

	toggle(id: string, checked: boolean) {
		this.store.dispatch(new Toggle({notificationID: id, checked}));
	}

}
