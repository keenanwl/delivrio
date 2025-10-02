import {Pipe, PipeTransform} from '@angular/core';
import {OrderEditActions} from "../../../orders/order-edit/order-edit.actions";
import OrderLineResponse = OrderEditActions.OrderLineResponse;

@Pipe({
	name: 'totalWeight',
	standalone: true
})
export class TotalWeightPipe implements PipeTransform {

	transform(value: OrderLineResponse[]): string {
		let totalGrams = 0;
		value.forEach((ol) => {
			totalGrams += ol.units * (ol.productVariant.weightG || 0);
		})
		return `${(totalGrams / 1000).toFixed(2)}/${(totalGrams / 453.592).toFixed(2)}`;
	}

}
