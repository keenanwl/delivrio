import {Pipe, PipeTransform} from '@angular/core';

@Pipe({
  name: 'emailTemplateFilter'
})
export class EmailTemplateFilterPipe implements PipeTransform {

	transform(value: unknown, ...args: unknown[]): unknown {
		return value;
	}

}
