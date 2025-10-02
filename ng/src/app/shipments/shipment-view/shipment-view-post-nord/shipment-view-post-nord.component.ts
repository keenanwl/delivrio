import {Component, Input} from '@angular/core';
import {ShipmentViewActions} from "../shipment-view.actions";
import ShipmentParcelResponse = ShipmentViewActions.ShipmentParcelResponse;
import {DebugSetIdsComponent} from "../dialogs/debug-set-ids/debug-set-ids.component";
import {MatDialog} from "@angular/material/dialog";

@Component({
  selector: 'app-shipment-view-post-nord',
  templateUrl: './shipment-view-post-nord.component.html',
  styleUrls: ['./shipment-view-post-nord.component.scss']
})
export class ShipmentViewPostNordComponent {

	@Input() bookingID: string = '1';
	@Input() labelsItemIDs: ShipmentParcelResponse[] = [];
	@Input() shipmentReferenceNumber: string = '3';

	constructor(private dialog: MatDialog) {
	}


	editIDs(parcelID: string) {
		const ref = this.dialog.open(DebugSetIdsComponent);
		ref.componentInstance.parcelID = parcelID;
	}

}
