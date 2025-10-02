import {Injectable} from "@angular/core";
import {Action, Selector, State, StateContext, Store} from "@ngxs/store";
import SetNotificationsList = NotificationsListActions.SetNotificationsList;
import {toNotNullArray} from "../../../functions/not-null-array";
import NotificationsResponse = NotificationsListActions.NotificationsResponse;
import {AppActions} from "../../../app.actions";
import ShowGlobalSnackbar = AppActions.ShowGlobalSnackbar;
import {CreateNotificationGQL, FetchNotificationsGQL, ToggleNotificationGQL} from "./notifications-list.generated";
import {NotificationsListActions} from "./notifications-list.actions";
import ConnectionsResponse = NotificationsListActions.ConnectionsResponse;
import EmailTemplatesResponse = NotificationsListActions.EmailTemplatesResponse;
import SetEmailTemplates = NotificationsListActions.SetEmailTemplates;
import SetConnections = NotificationsListActions.SetConnections;

export interface NotificationsListModel {
	notificationsList: NotificationsResponse[];
	connections: ConnectionsResponse[];
	emailTemplates: EmailTemplatesResponse[];
	loading: boolean;
}

const defaultState: NotificationsListModel = {
	notificationsList: [],
	connections: [],
	emailTemplates: [],
	loading: false,
};

@Injectable()
@State<NotificationsListModel>({
	name: 'NotificationsList',
	defaults: defaultState,
})
export class NotificationsListState {

	constructor(
		private list: FetchNotificationsGQL,
		private create: CreateNotificationGQL,
		private toggle: ToggleNotificationGQL,
	) {
	}

	@Selector()
	static get(state: NotificationsListModel) {
		return state;
	}

	@Action(NotificationsListActions.FetchNotificationsList)
	FetchMyNotificationsList(ctx: StateContext<NotificationsListModel>, action: NotificationsListActions.FetchNotificationsList) {
		ctx.patchState({loading: true});
		return this.list.fetch({})
			.subscribe({next: (r) => {
				ctx.patchState({loading: false});

				const notifications = toNotNullArray(r.data.notifications.edges?.map((l) => l?.node));
				ctx.dispatch(new SetNotificationsList(notifications));

				const templates = toNotNullArray(r.data.emailTemplates.edges?.map((l) => l?.node));
				ctx.dispatch(new SetEmailTemplates(templates))

				const connections = toNotNullArray(r.data.connections.edges?.map((l) => l?.node));
				ctx.dispatch(new SetConnections(connections))
			}});
	}

	@Action(NotificationsListActions.SetNotificationsList)
	SetNotificationsList(ctx: StateContext<NotificationsListModel>, action: NotificationsListActions.SetNotificationsList) {
		ctx.patchState({notificationsList: action.payload});
	}

	@Action(NotificationsListActions.SetConnections)
	SetConnections(ctx: StateContext<NotificationsListModel>, action: NotificationsListActions.SetConnections) {
		ctx.patchState({connections: action.payload});
	}

	@Action(NotificationsListActions.SetEmailTemplates)
	SetEmailTemplates(ctx: StateContext<NotificationsListModel>, action: NotificationsListActions.SetEmailTemplates) {
		ctx.patchState({emailTemplates: action.payload});
	}

	@Action(NotificationsListActions.Clear)
	Clear(ctx: StateContext<NotificationsListModel>, action: NotificationsListActions.Clear) {
		ctx.setState(defaultState);
	}

	@Action(NotificationsListActions.Toggle)
	Toggle(ctx: StateContext<NotificationsListModel>, action: NotificationsListActions.Toggle) {
		return this.toggle.mutate(action.payload).subscribe((resp) => {
			if (!!resp.errors) {
				ctx.dispatch(new ShowGlobalSnackbar("An error occurred"));
			} else {
				let msg = "Enabled: success"
				if (!action.payload.checked) {
					msg = "Disabled: success"
				}
				ctx.patchState({
					notificationsList: ctx.getState().notificationsList.map((n) => {
						let out = Object.assign({}, n)
						if (n.id === action.payload.notificationID) {
							out.active = action.payload.checked;
						}
						return out;
					})
				});
				ctx.dispatch(new ShowGlobalSnackbar(msg));
			}
		})
	}

	@Action(NotificationsListActions.Create)
	Create(ctx: StateContext<NotificationsListModel>, action: NotificationsListActions.Create) {
		ctx.patchState({loading: true});
		return this.create.mutate({name: action.payload.name, connectionID: action.payload.connectionID, emailTemplateID: action.payload.emailTemplateID})
			.subscribe((r) => {
				ctx.patchState({loading: false});
				if (!!r.errors) {
					ctx.dispatch(new ShowGlobalSnackbar("An error occurred"));
				} else {
					const notifications = toNotNullArray(r.data?.createNotification);
					ctx.dispatch(new SetNotificationsList(notifications));
				}
			});
	}

}
