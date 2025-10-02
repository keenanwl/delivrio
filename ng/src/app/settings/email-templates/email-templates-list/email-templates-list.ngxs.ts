import {Injectable} from "@angular/core";
import {Action, Selector, State, StateContext, Store} from "@ngxs/store";
import SetEmailTemplatesList = EmailTemplatesListActions.SetEmailTemplatesList;
import {toNotNullArray} from "../../../functions/not-null-array";
import {EmailTemplatesListActions} from "./email-templates-list.actions";
import EmailTemplatesResponse = EmailTemplatesListActions.EmailTemplatesResponse;
import {CreateEmailTemplateGQL, FetchEmailTemapltesGQL} from "./email-templates-list.generated";
import {AppActions} from "../../../app.actions";
import ShowGlobalSnackbar = AppActions.ShowGlobalSnackbar;
import AppChangeRoute = AppActions.AppChangeRoute;
import {Paths} from "../../../app-routing.module";

export interface EmailTemplatesListModel {
	emailTemplatesList: EmailTemplatesResponse[];
	loading: boolean;
}

const defaultState: EmailTemplatesListModel = {
	emailTemplatesList: [],
	loading: false,
};

@Injectable()
@State<EmailTemplatesListModel>({
	name: 'emailTemplatesList',
	defaults: defaultState,
})
export class EmailTemplatesListState {

	constructor(
		private list: FetchEmailTemapltesGQL,
		private create: CreateEmailTemplateGQL,
	) {
	}

	@Selector()
	static get(state: EmailTemplatesListModel) {
		return state;
	}

	@Action(EmailTemplatesListActions.FetchEmailTemplatesList)
	FetchMyEmailTemplatesList(ctx: StateContext<EmailTemplatesListModel>, action: EmailTemplatesListActions.FetchEmailTemplatesList) {
		ctx.patchState({loading: true});
		return this.list.fetch({})
			.subscribe({
				next: (r) => {
					ctx.patchState({loading: false});
					const emailTemplates = toNotNullArray(r.data.emailTemplates.edges?.map((l) => l?.node));
					ctx.dispatch(new SetEmailTemplatesList(emailTemplates));
				},
				error: () => {
					ctx.patchState({loading: false});
				},
			});
	}

	@Action(EmailTemplatesListActions.SetEmailTemplatesList)
	SetMyEmailTemplatesList(ctx: StateContext<EmailTemplatesListModel>, action: EmailTemplatesListActions.SetEmailTemplatesList) {
		ctx.patchState({emailTemplatesList: action.payload})
	}

	@Action(EmailTemplatesListActions.Clear)
	Clear(ctx: StateContext<EmailTemplatesListModel>, action: EmailTemplatesListActions.Clear) {
		ctx.setState(defaultState);
	}

	@Action(EmailTemplatesListActions.Create)
	Create(ctx: StateContext<EmailTemplatesListModel>, action: EmailTemplatesListActions.Create) {
		ctx.patchState({loading: true});
		return this.create.mutate({name: action.payload.name, merge: action.payload.merge})
			.subscribe({
				next: (r) => {
					ctx.patchState({loading: false});
					if (!!r.errors) {
						ctx.dispatch(new ShowGlobalSnackbar("An error occurred"));
					} else {
						ctx.dispatch(new AppChangeRoute({path: Paths.SETTINGS_EMAIL_TEMPLATE_EDIT, queryParams: {id: r.data?.createEmailTemplates}}));
					}
				},
				error: () => {
					ctx.patchState({loading: false});
				}
			});
	}

}
