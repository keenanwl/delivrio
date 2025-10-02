import {Pipe, PipeTransform} from '@angular/core';
import {environment} from "../../environments/environment";

// Duplicated from the core project because otherwise resolution was not working
@Pipe({
	name: 'printJobStatusColorPipe',
	standalone: true
})
export class PrintJobStatusColorPipe implements PipeTransform {

	// String for print-client
	transform(value: string | "pending" | "pending_success" | "success" | "pending_cancel" | "canceled"): string {
	  switch (value) {
		  case "success":
		  case "pending_success":
			  return "#4fc375"
		  case "pending":
			  return "#ffcf0a"
		  case "canceled":
		  case "pending_cancel":
			  return "#E00440";
	  }

	  if (!environment.production) {
		  console.warn("Unknown print job status: " + value)
	  }

	  return "#000000";
	}

}
