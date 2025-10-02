import {Component, OnDestroy, OnInit} from '@angular/core';
import {Observable} from "rxjs";
import {ReturnsListModel, ReturnsListState} from "../../returns-list.ngxs";
import {Actions, ofActionDispatched, Store} from "@ngxs/store";
import {FormControl} from "@angular/forms";
import {debounceTime, filter} from "rxjs/operators";
import {ReturnsListActions} from "../../returns-list.actions";
import SearchOrders = ReturnsListActions.SearchOrders;
import {SearchResult} from "../../../../../generated/graphql";
import SetSelectedOrder = ReturnsListActions.SetSelectedOrder;
import ClearCreateOrder = ReturnsListActions.ClearCreateOrder;
import {ItemReturn} from "../../../../settings/return-portal-viewer/return-portal-frame/return-portal-frame.ngxs";
import IncrementQuantity = ReturnsListActions.IncrementQuantity;
import DecrementQuantity = ReturnsListActions.DecrementQuantity;
import SetSelectedItem = ReturnsListActions.SetSelectedItem;
import SetSelectedItemReason = ReturnsListActions.SetSelectedItemReason;
import CreateReturnOrder = ReturnsListActions.CreateReturnOrder;
import {MatDialogRef} from "@angular/material/dialog";
import Clear = ReturnsListActions.Clear;
import CreateReturnOrderPending = ReturnsListActions.CreateReturnOrderPending;
import {MatListOption} from "@angular/material/list";
import ChangeDeliveryOption = ReturnsListActions.ChangeDeliveryOption;

@Component({
  selector: 'app-new-return-dialog',
  templateUrl: './new-return-dialog.component.html',
  styleUrls: ['./new-return-dialog.component.scss']
})
export class NewReturnDialogComponent implements OnInit, OnDestroy {
	returnsList$: Observable<ReturnsListModel>;
	orderLookupControl = new FormControl('');

	constructor(
		private store: Store,
		private actions$: Actions,
		private ref: MatDialogRef<any>,
	) {
		this.returnsList$ = store.select(ReturnsListState.get);
	}

	ngOnInit() {
		this.orderLookupControl.valueChanges
			.pipe(debounceTime(100), filter((v) => !!v && v.length > 0))
			.subscribe((s) => {
				this.store.dispatch(new SearchOrders(s || ""));
			});

		this.actions$.pipe(ofActionDispatched(Clear))
			.subscribe(() => {
				this.ref.close();
			})
	}

	createOpen() {
		this.store.dispatch(new CreateReturnOrder());
	}

	createPending() {
		this.store.dispatch(new CreateReturnOrderPending());
	}

	resetSearch() {
		this.orderLookupControl.reset();
	}

	selectSearchResult(opt: SearchResult) {
		this.store.dispatch(new SetSelectedOrder(opt));
		this.resetSearch();
	}

	ngOnDestroy(): void {
		this.store.dispatch(new ClearCreateOrder());
	}

	toggle(event: {orderLineID: string, selected: boolean}) {
		this.store.dispatch(new SetSelectedItem(event))
	}

	increment(event: {orderLineID: string}) {
		this.store.dispatch(new IncrementQuantity(event));
	}

	decrement(event: {orderLineID: string}) {
		this.store.dispatch(new DecrementQuantity(event));
	}

	continueDisallowedMessage(selectedItems: ItemReturn[]): boolean {
		let mayContinue = false;
		selectedItems.forEach((i) => {
			let selectedQuantity = false;
			if (i.selected && i.quantity > 0) {
				selectedQuantity = true;
			}
			if (selectedQuantity && i.id.length > 0) {
				mayContinue = true;
			}
		});
		return mayContinue;
	}

	reasonChanged(event: {orderLineID: string, reasonID: string}) {
		this.store.dispatch(new SetSelectedItemReason(event));
	}

	changeDeliveryOption(returnColliID: string, option: MatListOption[]) {
		// Should only be 1 delivery option
		this.store.dispatch(new ChangeDeliveryOption({returnColliID, deliveryOptionID: option.pop()?.value || ''}));
	}

}
