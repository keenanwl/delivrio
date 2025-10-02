import {Pipe, PipeTransform} from '@angular/core';
import {DatePipe} from '@angular/common';

@Pipe({
	name: 'relativeTime'
})
export class RelativeTimePipe implements PipeTransform {
	constructor(private datePipe: DatePipe) {}

	transform(value: string | Date, ticker: number): string {
		const inputDate = new Date(value);
		const currentDate = new Date();
		const diffInSeconds = Math.floor((currentDate.getTime() - inputDate.getTime()) / 1000);

		if (diffInSeconds < 60) {
			return 'just now';
		} else if (diffInSeconds < 3600) {
			const minutes = Math.floor(diffInSeconds / 60);
			return `${minutes} minute${minutes > 1 ? 's' : ''} ago`;
		} else if (diffInSeconds < 86400) {
			return this.datePipe.transform(inputDate, 'h:mm a') || '';
		} else if (diffInSeconds < 604800) {
			const days = Math.floor(diffInSeconds / 86400);
			return `${days} day${days > 1 ? 's' : ''} ago`;
		} else {
			return this.datePipe.transform(inputDate, 'mediumDate') || '';
		}
	}
}
