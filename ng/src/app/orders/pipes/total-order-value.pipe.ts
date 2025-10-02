import {Pipe, PipeTransform} from '@angular/core';
import {OrdersActions} from "../orders.actions";
import OrderLineResponse = OrdersActions.OrderLineResponse;

@Pipe({
	name: 'totalOrderValue',
	standalone: true
})
export class TotalOrderValuePipe implements PipeTransform {

	transform(collis: OrderLineResponse[]): string {

		let total = 0;
		let currency: string | null = null;
		collis.forEach((c) => {
			c.orderLines?.forEach((v) => {
				total += v.unitPrice * v.units;
				if (!currency) {
					currency = v.currency.display
				} else if (currency !== v.currency.display) {
					return `-`
				}
			})
		})

		if (currency === null) {
			return "-";
		}

		return `${total.toFixed(0)} ${currency}`
	}

}
