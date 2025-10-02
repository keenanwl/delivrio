import {Injectable} from "@angular/core";
import {Action, Selector, State, StateContext, Store} from "@ngxs/store";
import {formErrors} from "../../../account/company-info/company-info.ngxs";
import {produce} from "immer";
import {FetchEmailTemplateGQL, FireTestEmailGQL, UpdateEmailTemplateGQL} from "./email-template-edit.generated";
import {EmailTemplateEditActions} from "./email-template-edit.actions";
import EmailTemplateResponse = EmailTemplateEditActions.EmailTemplateResponse;
import {AppActions} from "../../../app.actions";
import ShowGlobalSnackbar = AppActions.ShowGlobalSnackbar;
import AppChangeRoute = AppActions.AppChangeRoute;
import {Paths} from "../../../app-routing.module";

export interface EmailTemplateEditModel {
	id: string;
	form: {
		model: EmailTemplateResponse | undefined;
		dirty: boolean;
		status: string;
		errors: formErrors;
	}
	loading: boolean;
}

const defaultState: EmailTemplateEditModel = {
	id: "",
	form: {
		model: undefined,
		dirty: false,
		status: '',
		errors: {}
	},
	loading: false,
};

@Injectable()
@State<EmailTemplateEditModel>({
	name: 'emailTemplateEdit',
	defaults: defaultState,
})
export class EmailTemplateEditState {

	constructor(
		private fetch: FetchEmailTemplateGQL,
		private update: UpdateEmailTemplateGQL,
		private send: FireTestEmailGQL,
	) {
	}

	@Selector()
	static get(state: EmailTemplateEditModel) {
		return state;
	}

	@Action(EmailTemplateEditActions.FetchEmailTemplateEdit)
	FetchEmailTemplateEdit(ctx: StateContext<EmailTemplateEditModel>, action: EmailTemplateEditActions.FetchEmailTemplateEdit) {
		return this.fetch.fetch({id: ctx.getState().id})
			.subscribe({next: (r) => {
				const emailTemplate = r.data.emailTemplate;
				if (!!emailTemplate) {
					ctx.dispatch(new EmailTemplateEditActions.SetEmailTemplateEdit(emailTemplate));
				}
			}});
	}

	@Action(EmailTemplateEditActions.SetEmailTemplateEdit)
	SetEmailTemplateEdit(ctx: StateContext<EmailTemplateEditModel>, action: EmailTemplateEditActions.SetEmailTemplateEdit) {
		const state = produce(ctx.getState(), st => {
			st.form.model = action.payload;
		});
		ctx.setState(state);
	}

	@Action(EmailTemplateEditActions.SetEmailTemplateID)
	SetEmailTemplateID(ctx: StateContext<EmailTemplateEditModel>, action: EmailTemplateEditActions.SetEmailTemplateID) {
		ctx.patchState({id: action.payload})
	}

	@Action(EmailTemplateEditActions.Clear)
	Clear(ctx: StateContext<EmailTemplateEditModel>, action: EmailTemplateEditActions.Clear) {
		ctx.setState(defaultState);
	}

	@Action(EmailTemplateEditActions.SendTestEmail)
	SendTestEmail(ctx: StateContext<EmailTemplateEditModel>, action: EmailTemplateEditActions.SendTestEmail) {
		ctx.patchState({loading: true});
		return this.send.fetch({id: ctx.getState().id, toEmail: action.payload}, {errorPolicy: "all"})
			.subscribe((r) => {
				ctx.patchState({loading: false});
				if (!!r.errors) {
					ctx.dispatch(new ShowGlobalSnackbar("An error occurred"));
				} else {
					ctx.dispatch(new ShowGlobalSnackbar("Success. Check your inbox."));
				}
			});
	}

	@Action(EmailTemplateEditActions.Save)
	Save(ctx: StateContext<EmailTemplateEditModel>, action: EmailTemplateEditActions.Save) {
		ctx.patchState({loading: true});
		return this.update.mutate({id: ctx.getState().id, input: ctx.getState().form?.model || {}}, {errorPolicy: "all"})
			.subscribe((r) => {
				ctx.patchState({loading: false});
				if (!!r.errors) {
					ctx.dispatch(new ShowGlobalSnackbar(`An error occurred: ${r.errors.map((e) => e.message).join(" ")}`));
				} else {
					ctx.dispatch(new AppChangeRoute({path: Paths.SETTINGS_EMAIL_TEMPLATES, queryParams: {}}));
				}
			});
	}

}
