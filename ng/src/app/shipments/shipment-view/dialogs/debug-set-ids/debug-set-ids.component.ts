import {Component, Input} from '@angular/core';
import {Store} from "@ngxs/store";
import {ShipmentViewActions} from "../../shipment-view.actions";
import DebugUpdateLabelIDs = ShipmentViewActions.DebugUpdateLabelIDs;
import {MatDialogRef} from "@angular/material/dialog";

@Component({
	selector: 'app-debug-set-ids',
	templateUrl: './debug-set-ids.component.html',
	styleUrls: ['./debug-set-ids.component.scss']
})
export class DebugSetIdsComponent {

	@Input() parcelID: string = "";

	constructor(
		private store: Store,
		private ref: MatDialogRef<any>,
	) {}

	save(itemID: string) {
		this.store.dispatch(new DebugUpdateLabelIDs({parcelID: this.parcelID, itemID}));
		this.ref.close();
	}

}
