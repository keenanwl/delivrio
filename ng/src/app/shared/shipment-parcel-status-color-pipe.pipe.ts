import {Pipe, PipeTransform} from '@angular/core';
import {ShipmentParcelStatus} from "../../generated/graphql";
import {environment} from "../../environments/environment";

@Pipe({
	name: 'shipmentParcelStatusColorPipe',
	standalone: true
})
export class ShipmentParcelStatusColorPipePipe implements PipeTransform {

  transform(value: ShipmentParcelStatus): string {
	  switch (value) {
		case ShipmentParcelStatus.Pending:
			return "#ffcf0a"
		case ShipmentParcelStatus.Printed:
		  	return "#FB8C00"
		case ShipmentParcelStatus.InTransit:
		  	return "#1f78ff"
		case ShipmentParcelStatus.OutForDelivery:
		  	return "#1f78ff"
		case ShipmentParcelStatus.Delivered:
		  	return "#4fc375"
		case ShipmentParcelStatus.AwaitingCcPickup:
		  	return "#1f78ff"
		case ShipmentParcelStatus.PickedUp:
		  	return "#4fc375"
	  }

	  if (!environment.production) {
		  console.warn("Unknown shipment parcel status: " + value)
	  }

	  return "#908e8e";
  }

}
