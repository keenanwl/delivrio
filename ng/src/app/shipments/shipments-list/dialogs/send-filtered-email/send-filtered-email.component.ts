import {Component} from '@angular/core';
import {ShipmentsListModel, ShipmentsListState} from "../../shipments-list.ngxs";
import {Store} from "@ngxs/store";
import {Observable} from "rxjs";
import {ShipmentsListActions} from "../../shipments-list.actions";
import SendOverviewEmail = ShipmentsListActions.SendOverviewEmail;

@Component({
	selector: 'app-send-filtered-email',
	templateUrl: './send-filtered-email.component.html',
	styleUrls: ['./send-filtered-email.component.scss']
})
export class SendFilteredEmailComponent {

	shipments$: Observable<ShipmentsListModel>;

	constructor(private store: Store) {
		this.shipments$ = store.select(ShipmentsListState.state);
	}

	send(email: string, templateID: string) {
		this.store.dispatch(new SendOverviewEmail({email, templateID}));
	}
}
