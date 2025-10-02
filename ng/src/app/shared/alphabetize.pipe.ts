import {Pipe, PipeTransform} from '@angular/core';

@Pipe({
	name: 'alphabetize',
	standalone: true
})
export class AlphabetizePipe implements PipeTransform {

	transform(array: any[], property: string): any[] {
		if (!array || array.length === 0 || !property) {
			return array;
		}

		// Copy the array to avoid mutating the original array
		const newArray = [...array];

		// Alphabetize the array based on the specified property
		newArray.sort((a, b) => {
			const propA = a[property].toLowerCase();
			const propB = b[property].toLowerCase();

			if (propA < propB) {
				return -1;
			} else if (propA > propB) {
				return 1;
			} else {
				return 0;
			}
		});

		return newArray;
	}

}
