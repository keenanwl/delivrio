import {MatDialogRef} from "@angular/material/dialog";
import {Component, EventEmitter, Output} from "@angular/core";
import {Observable} from "rxjs";
import {DeliveryOptionsListModel, DeliveryOptionsListState} from "./delivery-options-list.ngxs";
import {Store} from "@ngxs/store";
import {DeliveryOptionsListActions} from "./delivery-options-list.actions";
import CreateNewGLSDeliveryOption = DeliveryOptionsListActions.CreateNewDeliveryOption;
import FetchCarrierAgreements = DeliveryOptionsListActions.FetchCarrierAgreements;

type carriers = "gls";

@Component({
	selector: 'new-delivery-option-dialog',
	styleUrls: ['new-delivery-option-dialog.component.scss'],
	templateUrl: 'new-delivery-option-dialog.component.html',
})
export class NewDeliveryOptionDialogComponent {

	@Output() selectedCarrier: EventEmitter<carriers> = new EventEmitter();

	deliveryOptionsList$: Observable<DeliveryOptionsListModel>;

	constructor(
		private store: Store,
		private dialogRef: MatDialogRef<NewDeliveryOptionDialogComponent>,
	) {
		this.deliveryOptionsList$ = store.select(DeliveryOptionsListState.get);
		this.store.dispatch([new FetchCarrierAgreements()]);
	}

	selected(name: string, agreementId: string, connectionID: string) {
		this.store.dispatch(new CreateNewGLSDeliveryOption({name, agreementId, connectionID}));
		this.dialogRef.close();
	}
}
