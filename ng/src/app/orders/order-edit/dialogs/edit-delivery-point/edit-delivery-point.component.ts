import { Component, OnInit } from '@angular/core';
import { Store } from '@ngxs/store';
import {OrderEditActions} from "../../order-edit.actions";
import SearchDeliveryPoints = OrderEditActions.SearchDeliveryPoints;
import {Observable} from "rxjs";
import {OrderEditModel, OrderEditState} from "../../order-edit.ngxs";
import {MatDialogRef} from "@angular/material/dialog";
import DeliveryPointResponse = OrderEditActions.DeliveryPointResponse;
import SetDeliveryPoint = OrderEditActions.SetDeliveryPoint;

@Component({
	selector: 'app-edit-delivery-point',
	templateUrl: './edit-delivery-point.component.html',
	styleUrls: ['./edit-delivery-point.component.scss']
})
export class EditDeliveryPointComponent implements OnInit {
	order$: Observable<OrderEditModel>;
	constructor(private store: Store, private ref: MatDialogRef<any>) {
		this.order$ = store.select(OrderEditState.get);
	}

	ngOnInit(): void {
		this.store.dispatch(new SearchDeliveryPoints());
	}

	selectDeliveryPoint(dp: DeliveryPointResponse) {
		this.store.dispatch(new SetDeliveryPoint(dp));
		this.close();
	}

	close() {
		this.ref.close();
	}

}
