import {Pipe, PipeTransform} from '@angular/core';
import {ColliStatus} from "../../generated/graphql";
import {environment} from "../../environments/environment";

@Pipe({
  name: 'colliStatusColorPipe',
  standalone: true
})
export class ColliStatusColorPipePipe implements PipeTransform {

  transform(value: ColliStatus): string {

	  switch (value) {
		  case ColliStatus.Cancelled:
			  return "#E00440";
		  case ColliStatus.Dispatched:
			  return "#4fc375"
		  case ColliStatus.Pending:
			  return "#ffcf0a"
	  }

	  if (!environment.production) {
		  console.warn("Unknown colli status: " + value)
	  }

	  return "#000000";
  }

}
