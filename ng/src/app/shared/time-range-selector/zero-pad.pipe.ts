import {Pipe, PipeTransform} from '@angular/core';
import {GraphQLError} from "graphql/index";
import {FormControl} from "@angular/forms";

@Pipe({
	name: 'zeroPad',
	pure: true,
})
export class ZeroPadPipe implements PipeTransform {

	constructor() {
	}

	transform(value: number, totalLength: number): string {
		return String(value).padStart(totalLength, '0')
	}
}
