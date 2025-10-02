import {Pipe, PipeTransform} from '@angular/core';
import {OrderEditActions} from "../order-edit.actions";
import FetchDeliveryOptionsResponse = OrderEditActions.FetchDeliveryOptionsResponse;

@Pipe({
  name: 'deliveryPointOptional'
})
export class DeliveryPointOptionalPipe implements PipeTransform {

  transform(selectedID: string | null, deliveryPoints: FetchDeliveryOptionsResponse[]): boolean {
	  return deliveryPoints.some((d) => {
		  if (d.deliveryOptionID === selectedID) {
			  return d.deliveryPoint;
		  }
		  return false;
	  })
  }

}
