import {Component, Input, OnDestroy} from '@angular/core';
import {Observable, Subscription} from "rxjs";
import {OrderViewModel, OrderViewState} from "../../order-view.ngxs";
import {Store} from "@ngxs/store";
import {DialogRef} from "@angular/cdk/dialog";
import {OrderViewActions} from "../../order-view.actions";
import DecrementLabelViewerOffset = OrderViewActions.DecrementLabelViewerOffset;
import IncrementLabelViewerOffset = OrderViewActions.IncrementLabelViewerOffset;
import CancelShipment = OrderViewActions.CancelShipment;
import ClearShipmentView = OrderViewActions.ClearDialogs;
import {toNotNullArray} from "../../../../functions/not-null-array";
import CreateLabelPrintJobs = OrderViewActions.CreateLabelPrintJobs;

@Component({
	selector: 'app-create-shipment',
	templateUrl: './create-shipment.component.html',
	styleUrls: ['./create-shipment.component.scss']
})
export class CreateShipmentComponent implements OnDestroy {

	orderView$: Observable<OrderViewModel>;

	subscriptions: Subscription[] = [];

	@Input() displayType: "view" | "create" = "create";

	constructor(
		private store: Store,
		private dialog: DialogRef,
	) {
		this.orderView$ = store.select(OrderViewState.get);
	}

	close() {
		this.dialog.close();
	}

	increment() {
		this.store.dispatch(new IncrementLabelViewerOffset());
	}

	decrement() {
		this.store.dispatch(new DecrementLabelViewerOffset());
	}

	cancelShipment() {
		this.store.dispatch(new CancelShipment());
		this.close();
	}

	ngOnDestroy(): void {
		this.store.dispatch(new ClearShipmentView());
		this.subscriptions.forEach((s) => s.unsubscribe());
	}

	print() {
		const state = this.store.selectSnapshot(OrderViewState.get);
		this.store.dispatch(new CreateLabelPrintJobs({parcelIDs: state.selectedColliIDs}));
	}

}
