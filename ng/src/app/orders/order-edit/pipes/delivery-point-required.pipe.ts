import {Pipe, PipeTransform} from '@angular/core';
import {OrderEditActions} from "../order-edit.actions";
import FetchDeliveryOptionsResponse = OrderEditActions.FetchDeliveryOptionsResponse;

@Pipe({
  name: 'deliveryPointRequired'
})
export class DeliveryPointRequiredPipe implements PipeTransform {

	transform(selectedID: string | null, deliveryPoints: FetchDeliveryOptionsResponse[]): boolean {
		return deliveryPoints.some((d) => {
			if (d.deliveryOptionID === selectedID) {
				return d.requiresDeliveryPoint;
			}
			return false;
		})
	}

}
