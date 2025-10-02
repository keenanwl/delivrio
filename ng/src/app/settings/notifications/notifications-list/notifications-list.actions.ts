import {FetchNotificationsQuery} from "./notifications-list.generated";

export namespace NotificationsListActions {
	export class FetchNotificationsList {
		static readonly type = '[NotificationsList] fetch NotificationsList';
	}
	export class SetNotificationsList {
		static readonly type = '[NotificationsList] set NotificationsList';
		constructor(public payload: NotificationsResponse[]) {}
	}
	export class SetConnections {
		static readonly type = '[NotificationsList] set connections';
		constructor(public payload: ConnectionsResponse[]) {}
	}
	export class SetEmailTemplates {
		static readonly type = '[NotificationsList] set email templates';
		constructor(public payload: EmailTemplatesResponse[]) {}
	}
	export class Clear {
		static readonly type = '[NotificationsList] clear';
	}
	export class Create {
		static readonly type = '[NotificationsList] create';
		constructor(public payload: {name: string; connectionID: string; emailTemplateID: string;}) {}
	}
	export class Toggle {
		static readonly type = '[NotificationsList] toggle';
		constructor(public payload: {notificationID: string; checked: boolean}) {}
	}
	export type NotificationsResponse = NonNullable<NonNullable<NonNullable<NonNullable<FetchNotificationsQuery['notifications']>['edges']>[0]>['node']>;
	export type ConnectionsResponse = NonNullable<NonNullable<NonNullable<NonNullable<FetchNotificationsQuery['connections']>['edges']>[0]>['node']>;
	export type EmailTemplatesResponse = NonNullable<NonNullable<NonNullable<NonNullable<FetchNotificationsQuery['emailTemplates']>['edges']>[0]>['node']>;
}
