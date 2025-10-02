import {Pipe, PipeTransform} from '@angular/core';
import {EmailTemplateMergeType} from "../../../../../generated/graphql";

@Pipe({
	name: 'filterEmailTemplates',
	standalone: true
})
export class FilterEmailTemplatesPipe implements PipeTransform {
	transform<T extends {mergeType: EmailTemplateMergeType}>(value: T[], filterFor: EmailTemplateMergeType): T[] {
		return value.filter((o) => o.mergeType === filterFor)
	}
}
