import {FetchEmailTemplateQuery} from "./email-template-edit.generated";

export namespace EmailTemplateEditActions {
	export class FetchEmailTemplateEdit {
		static readonly type = '[EmailTemplateEdit] fetch EmailTemplateEdit';
	}
	export class SetEmailTemplateEdit {
		static readonly type = '[EmailTemplateEdit] set EmailTemplateEdit';
		constructor(public payload: EmailTemplateResponse) {}
	}
	export class SetEmailTemplateID {
		static readonly type = '[EmailTemplateEdit] set EmailTemplateEdit ID';
		constructor(public payload: string) {}
	}
	export class SendTestEmail {
		static readonly type = '[EmailTemplateEdit] send test email';
		constructor(public payload: string) {}
	}
	export class Clear {
		static readonly type = '[EmailTemplateEdit] clear';
	}
	export class Save {
		static readonly type = '[EmailTemplateEdit] create';
	}
	export type EmailTemplateResponse = NonNullable<FetchEmailTemplateQuery['emailTemplate']>;
}
