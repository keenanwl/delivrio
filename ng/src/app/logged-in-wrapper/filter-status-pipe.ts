import {Pipe, PipeTransform} from '@angular/core';
import {AppActions} from "../app.actions";
import {PrintJobStatus} from "../../generated/graphql";
import PrintJob = AppActions.PrintJob;

@Pipe({
	standalone: true,
	name: 'filterStatus'
})
export class FilterStatusPipe implements PipeTransform {

	transform(items: PrintJob[]): any[] {
		if (!items) {
			return [];
		}
		return items.filter(item =>
			item.status === PrintJobStatus.Pending || item.status === PrintJobStatus.AtPrinter);
	}

}
