import {Pipe, PipeTransform} from '@angular/core';
import {ConsolidationEditActions} from "../consolidation-edit/consolidation-edit.actions";
import DeliveryOptionItem = ConsolidationEditActions.DeliveryOptionItem;

@Pipe({
	name: 'deliveryOptionGrouper',
	standalone: true
})
export class DeliveryOptionGrouperPipe implements PipeTransform {
	transform(value: DeliveryOptionItem[]): {outbound: DeliveryOptionItem[]; inbound: DeliveryOptionItem[]} {

		const outbound: DeliveryOptionItem[] = [];
		const inbound: DeliveryOptionItem[] = [];

		value.forEach((v) => {
			if (v.carrierService.return) {
				inbound.push(v);
			} else {
				outbound.push(v);
			}
		})

		return {outbound, inbound};
	}

}
