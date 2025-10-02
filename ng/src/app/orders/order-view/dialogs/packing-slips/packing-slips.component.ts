import {Component, OnDestroy} from '@angular/core';
import {Observable} from "rxjs";
import {OrderViewModel, OrderViewState} from "../../order-view.ngxs";
import {Store} from "@ngxs/store";
import {DialogRef} from "@angular/cdk/dialog";
import {OrderViewActions} from "../../order-view.actions";
import IncrementLabelViewerOffset = OrderViewActions.IncrementLabelViewerOffset;
import DecrementLabelViewerOffset = OrderViewActions.DecrementLabelViewerOffset;
import CreatePackingSlipPrintJobs = OrderViewActions.CreatePackingSlipPrintJobs;
import {toNotNullArray} from "../../../../functions/not-null-array";
import ClearDialogs = OrderViewActions.ClearDialogs;

@Component({
	selector: 'app-packing-slips',
	templateUrl: './packing-slips.component.html',
	styleUrls: ['./packing-slips.component.scss']
})
export class PackingSlipsComponent implements OnDestroy {

	orderView$: Observable<OrderViewModel>;

	constructor(
		private store: Store,
		private dialog: DialogRef,
	) {
		this.orderView$ = store.select(OrderViewState.get);
	}

	ngOnDestroy(): void {
        this.store.dispatch(new ClearDialogs);
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

	print() {
		const state = this.store.selectSnapshot(OrderViewState.get);
		this.store.dispatch(new CreatePackingSlipPrintJobs({parcelIDs: toNotNullArray(state.order?.colli?.map((c) => c.id))}));
	}

}
