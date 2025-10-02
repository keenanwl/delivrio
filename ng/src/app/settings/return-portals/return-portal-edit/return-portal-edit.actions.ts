import {FetchReturnPortalQuery} from "./return-portal-edit.generated";

export namespace ReturnPortalEditActions {
	export class FetchReturnPortalEdit {
		static readonly type = '[ReturnPortalEdit] fetch return portal edit';
	}
	export class SetReturnPortalEdit {
		static readonly type = '[ReturnPortalEdit] set return portal edit';
		constructor(public payload: ReturnPortalEditResponse) {}
	}
	export class SetReturnDeliveryOptions {
		static readonly type = '[ReturnPortalEdit] set return portal delivery options';
		constructor(public payload: ReturnDeliveryOptionsResponse[]) {}
	}
	export class SetReturnPortalID {
		static readonly type = '[ReturnPortalEdit] set return portal ID';
		constructor(public payload: string) {}
	}
	export class SetConnections {
		static readonly type = '[ReturnPortalEdit] set connections';
		constructor(public payload: ConnectionsResponse[]) {}
	}
	export class SetDeliveryOptions {
		static readonly type = '[ReturnPortalEdit] set delivery options';
		constructor(public payload: string[]) {}
	}
	export class Save {
		static readonly type = '[ReturnPortalEdit] save';
	}
	export class Clear {
		static readonly type = '[ReturnPortalEdit] clear';
	}
	export class AddClaim {
		static readonly type = '[ReturnPortalEdit] add claim';
	}
	export class DeleteClaim {
		static readonly type = '[ReturnPortalEdit] delete claim';
		constructor(public payload: number) {}
	}
	export class SetEmailTemplates {
		static readonly type = '[ReturnPortalEdit] set email templates';
		constructor(public payload: EmailTemplateResponse[]) {}
	}
	export class SetSelectedDeliveryOptions {
		static readonly type = '[ReturnPortalEdit] set selected delivery options';
		constructor(public payload: string[]) {}
	}
	export type ReturnPortalEditResponse = NonNullable<FetchReturnPortalQuery['returnPortal']>;
	export type ReturnDeliveryOptionsResponse = NonNullable<NonNullable<NonNullable<FetchReturnPortalQuery['deliveryOptions']>['edges']>[0]>['node'];
	export type ClaimResponse = NonNullable<NonNullable<NonNullable<FetchReturnPortalQuery['returnPortal']>['returnPortalClaim']>[0]>;
	export type ConnectionsResponse = NonNullable<NonNullable<NonNullable<NonNullable<FetchReturnPortalQuery['connections']>['edges']>[0]>['node']>;
	export type EmailTemplateResponse = NonNullable<NonNullable<NonNullable<NonNullable<FetchReturnPortalQuery['emailTemplates']>['edges']>[0]>['node']>;
}
