import {Pipe, PipeTransform} from '@angular/core';
import {CarrierServiceItemFragment} from "../edit-common.generated";

@Pipe({
	name: 'carrierServiceGrouper',
	standalone: true
})
export class CarrierServiceGrouperPipe implements PipeTransform {

  transform(value: CarrierServiceItemFragment[]): {outbound: CarrierServiceItemFragment[]; inbound: CarrierServiceItemFragment[]} {

	  const outbound: CarrierServiceItemFragment[] = [];
	  const inbound: CarrierServiceItemFragment[] = [];

	  value.forEach((v) => {
		  if (v.return) {
			  inbound.push(v);
		  } else {
			  outbound.push(v);
		  }
	  })

    return {outbound, inbound};
  }

}
