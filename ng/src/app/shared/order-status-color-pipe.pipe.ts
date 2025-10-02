import {Pipe, PipeTransform} from '@angular/core';
import {OrderStatus} from "../../generated/graphql";
import {environment} from "../../environments/environment";

@Pipe({
  name: 'orderStatusColorPipe',
  standalone: true
})
export class OrderStatusColorPipePipe implements PipeTransform {

  transform(value: OrderStatus): string {
	  switch (value) {
		  case OrderStatus.Cancelled:
			  return "#E00440";
		  case OrderStatus.Dispatched:
			  return "#4fc375"
		  case OrderStatus.PartiallyDispatched:
			  return "#1f78ff"
		  case OrderStatus.Pending:
			  return "#ffcf0a"
	  }

	  if (!environment.production) {
		  console.warn("Unknown order status: " + value)
	  }

	  return "#000000";
  }

}
