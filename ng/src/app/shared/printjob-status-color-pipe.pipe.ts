import {Pipe, PipeTransform} from '@angular/core';
import {OrderStatus, PrintJobStatus} from "../../generated/graphql";
import {environment} from "../../environments/environment";

@Pipe({
  name: 'printJobStatusColorPipe',
  standalone: true
})
export class PrintJobStatusColorPipe implements PipeTransform {

	// String for print-client
  transform(value: PrintJobStatus | string): string {
	  switch (value) {
		  case PrintJobStatus.Canceled:
			  return "#E00440";
		  case PrintJobStatus.Success:
			  return "#4fc375"
		  case PrintJobStatus.AtPrinter:
			  return "#1f78ff"
		  case PrintJobStatus.Pending:
			  return "#ffcf0a"
	  }

	  if (!environment.production) {
		  console.warn("Unknown print job status: " + value)
	  }

	  return "#000000";
  }

}
