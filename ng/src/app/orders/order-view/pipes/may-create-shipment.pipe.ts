import {Pipe, PipeTransform} from '@angular/core';
import {OrderViewActions} from "../order-view.actions";
import ShipmentStatusResponse = OrderViewActions.ShipmentStatusResponse;

@Pipe({
	name: 'mayCreateShipment'
})
export class MayCreateShipmentPipe implements PipeTransform {

	transform(id: string, statuses: ShipmentStatusResponse[]): string {

		let output = "";
		statuses.some((s) => {
			if (s.colliID === id) {
				if (!!s.shipmentID) {
					output = s.shipmentID;
				}
				return true;
			}
			return false;
		})
		return output;
	}

}
