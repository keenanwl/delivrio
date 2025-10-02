import {Pipe, PipeTransform} from '@angular/core';
import {ShipmentStatus} from "../../generated/graphql";
import {environment} from "../../environments/environment";

@Pipe({
	name: 'shipmentStatusColorPipe',
	standalone: true
})
export class ShipmentStatusColorPipePipe implements PipeTransform {

	transform(value: ShipmentStatus): string {
		switch (value) {
			case ShipmentStatus.Pending:
				return "#ffcf0a";
			case ShipmentStatus.Prebooked:
				return "#FB8C00";
			case ShipmentStatus.Booked:
				return "#FB8C00"
			case ShipmentStatus.PartiallyDispatched:
				return "#1f78ff"
			case ShipmentStatus.Dispatched:
				return "#4fc375"
			case ShipmentStatus.Deleted:
				return "#000000"
		}

		if (!environment.production) {
			console.warn("Unknown shipment status: " + value)
		}

		return "rgba(0,0,0,0.38)";
	}

}
