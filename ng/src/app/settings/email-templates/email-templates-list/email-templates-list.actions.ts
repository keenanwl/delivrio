import {FetchEmailTemapltesQuery} from "./email-templates-list.generated";
import {EmailTemplateMergeType} from "../../../../generated/graphql";

export namespace EmailTemplatesListActions {
	export class FetchEmailTemplatesList {
		static readonly type = '[EmailTemplatesList] fetch EmailTemplatesList';
	}
	export class SetEmailTemplatesList {
		static readonly type = '[EmailTemplatesList] set EmailTemplatesList';
		constructor(public payload: EmailTemplatesResponse[]) {}
	}
	export class Clear {
		static readonly type = '[EmailTemplatesList] clear';
	}
	export class Create {
		static readonly type = '[EmailTemplatesList] create';
		constructor(public payload: {name: string; merge: EmailTemplateMergeType}) {}
	}
	export type EmailTemplatesResponse = NonNullable<NonNullable<NonNullable<NonNullable<FetchEmailTemapltesQuery['emailTemplates']>['edges']>[0]>['node']>;
}
