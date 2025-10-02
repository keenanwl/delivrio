import {Pipe, PipeTransform} from '@angular/core';

@Pipe({
	name: 'deliveryOptionSelected',
	standalone: true
})
export class DeliveryOptionSelectedPipe implements PipeTransform {

	transform(checkID: string, selectedIDs?: {id: string}[]): boolean {
		if (!!selectedIDs) {
			return selectedIDs.some((opt) => {
				return opt.id === checkID;
			});
		}
		return false;
	}

}
