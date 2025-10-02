import {Component, EventEmitter, Input, Output} from '@angular/core';
import {MaterialModule} from "../../../../../modules/material.module";
import {DialogRef} from "@angular/cdk/dialog";
import {Store} from "@ngxs/store";
import {DeliveryOptionsListActions} from "../../delivery-options-list.actions";
import Archive = DeliveryOptionsListActions.Archive;

@Component({
  selector: 'app-archive-confirmation',
  standalone: true,
	imports: [
		MaterialModule
	],
  templateUrl: './archive-confirmation.component.html',
  styleUrl: './archive-confirmation.component.scss'
})
export class ArchiveConfirmationComponent {

	@Input() dOptID = "";
	@Input() doptName = "";

	constructor(private ref: DialogRef, private store: Store) {
	}

	confirm() {
		this.store.dispatch(new Archive({deliveryOptionID: this.dOptID}));
		this.cancel();
	}

	cancel() {
		this.ref.close();
	}
}
