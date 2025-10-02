import {Component, OnDestroy, OnInit} from '@angular/core';
import {Observable, Subscription} from "rxjs";
import {NotificationsListModel, NotificationsListState} from "../../notifications-list.ngxs";
import {Store} from "@ngxs/store";
import {NotificationsListActions} from "../../notifications-list.actions";
import FetchNotificationsList = NotificationsListActions.FetchNotificationsList;
import {DialogRef} from "@angular/cdk/dialog";
import Create = NotificationsListActions.Create;

@Component({
	selector: 'app-create-notification',
	templateUrl: './create-notification.component.html',
	styleUrls: ['./create-notification.component.scss']
})
export class CreateNotificationComponent implements OnDestroy, OnInit {
	notifications$: Observable<NotificationsListModel>;
	subscriptions$: Subscription[] = []

	constructor(
		private store: Store,
		private ref: DialogRef,
	) {
		this.notifications$ = store.select(NotificationsListState.get);
	}

	ngOnDestroy(): void {
		this.subscriptions$.map((s) => s.unsubscribe());
		this.store.dispatch(new FetchNotificationsList());
	}

	ngOnInit(): void {

	}

	create(name: string, connectionID: string, emailTemplateID: string) {
		this.store.dispatch(new Create({name, connectionID, emailTemplateID}));
		this.close();
	}

	close() {
		this.ref.close();
	}
}
