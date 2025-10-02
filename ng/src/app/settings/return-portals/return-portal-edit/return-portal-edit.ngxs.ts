import {Injectable} from "@angular/core";
import {Action, Selector, State, StateContext, Store} from "@ngxs/store";
import SetReturnPortalEdit = ReturnPortalEditActions.SetReturnPortalEdit;
import {formErrors} from "../../../account/company-info/company-info.ngxs";
import ReturnPortalEditResponse = ReturnPortalEditActions.ReturnPortalEditResponse;
import {AppActions} from "../../../app.actions";
import ShowGlobalSnackbar = AppActions.ShowGlobalSnackbar;
import AppChangeRoute = AppActions.AppChangeRoute;
import {Paths} from "../../../app-routing.module";
import {ReturnPortalEditActions} from "./return-portal-edit.actions";
import {FetchReturnPortalGQL, UpdateReturnPortalGQL} from "./return-portal-edit.generated";
import {produce} from "immer";
import ConnectionsResponse = ReturnPortalEditActions.ConnectionsResponse;
import SetConnections = ReturnPortalEditActions.SetConnections;
import {toNotNullArray} from "../../../functions/not-null-array";
import SetReturnDeliveryOptions = ReturnPortalEditActions.SetReturnDeliveryOptions;
import ReturnDeliveryOptionsResponse = ReturnPortalEditActions.ReturnDeliveryOptionsResponse;
import SetEmailTemplates = ReturnPortalEditActions.SetEmailTemplates;
import EmailTemplateResponse = ReturnPortalEditActions.EmailTemplateResponse;
import {UpdateReturnPortalInput} from "../../../../generated/graphql";

export interface ReturnPortalEditModel {
	returnPortalEditForm: {
		model: ReturnPortalEditResponse | undefined;
		dirty: boolean;
		status: string;
		errors: formErrors;
	},
	returnDeliveryOptions: ReturnDeliveryOptionsResponse[];
	connections: ConnectionsResponse[];
	returnPortalID: string;
	emailTemplates: EmailTemplateResponse[];
}

const defaultState: ReturnPortalEditModel = {
	returnPortalEditForm: {
		model: undefined,
		dirty: false,
		status: '',
		errors: {}
	},
	returnDeliveryOptions: [],
	connections: [],
	returnPortalID: '',
	emailTemplates: [],
};

@Injectable()
@State<ReturnPortalEditModel>({
	name: 'returnPortalEdit',
	defaults: defaultState,
})
export class ReturnPortalEditState {

	constructor(
		private fetch: FetchReturnPortalGQL,
		private store: Store,
		private update: UpdateReturnPortalGQL,
	) {}

	@Selector()
	static get(state: ReturnPortalEditModel) {
		return state;
	}

	@Action(ReturnPortalEditActions.FetchReturnPortalEdit)
	FetchMyReturnPortalEdit(ctx: StateContext<ReturnPortalEditModel>, action: ReturnPortalEditActions.FetchReturnPortalEdit) {
		const id = ctx.getState().returnPortalID;
		return this.fetch.fetch({id})
			.subscribe({next: (r) => {

				const templates = toNotNullArray(r.data.emailTemplates.edges?.map((rp) => rp?.node));
				ctx.dispatch(new SetEmailTemplates(templates));

				const portal = r.data.returnPortal;
				ctx.dispatch(new SetReturnPortalEdit(portal));

				const deliveryOptions = toNotNullArray(r.data.deliveryOptions.edges?.map((n) => n?.node));
				ctx.dispatch(new SetReturnDeliveryOptions(deliveryOptions));

				const connections = toNotNullArray(r.data.connections.edges?.map((rp) => rp?.node));
				ctx.dispatch(new SetConnections(connections));

			}});
	}

	@Action(ReturnPortalEditActions.SetReturnPortalID)
	SetReturnPortalID(ctx: StateContext<ReturnPortalEditModel>, action: ReturnPortalEditActions.SetReturnPortalID) {
		ctx.patchState({returnPortalID: action.payload});
	}

	@Action(ReturnPortalEditActions.SetConnections)
	SetConnections(ctx: StateContext<ReturnPortalEditModel>, action: ReturnPortalEditActions.SetConnections) {
		ctx.patchState({connections: action.payload});
	}

	@Action(ReturnPortalEditActions.Clear)
	Clear(ctx: StateContext<ReturnPortalEditModel>, action: ReturnPortalEditActions.Clear) {
		ctx.setState(defaultState);
	}

	@Action(ReturnPortalEditActions.SetReturnPortalEdit)
	SetReturnPortalEdit(ctx: StateContext<ReturnPortalEditModel>, action: ReturnPortalEditActions.SetReturnPortalEdit) {
		const state = produce(ctx.getState(), st => {
			st.returnPortalEditForm.model = action.payload;
		});
		ctx.setState(state);
	}

	@Action(ReturnPortalEditActions.SetReturnDeliveryOptions)
	SetReturnDeliveryOptions(ctx: StateContext<ReturnPortalEditModel>, action: ReturnPortalEditActions.SetReturnDeliveryOptions) {
		ctx.patchState({returnDeliveryOptions: action.payload});
	}

	@Action(ReturnPortalEditActions.SetEmailTemplates)
	SetEmailTemplates(ctx: StateContext<ReturnPortalEditModel>, action: ReturnPortalEditActions.SetEmailTemplates) {
		ctx.patchState({emailTemplates: action.payload});
	}

	@Action(ReturnPortalEditActions.AddClaim)
	AddClaim(ctx: StateContext<ReturnPortalEditModel>, action: ReturnPortalEditActions.AddClaim) {
		const state = produce(ctx.getState(), st => {
			let current = st.returnPortalEditForm.model?.returnPortalClaim || [];
			current.push({
				id: "",
				name: "",
				description: "",
				restockable: false,
			});
			st.returnPortalEditForm.model!.returnPortalClaim = current;
		});
		ctx.setState(state);
	}

	@Action(ReturnPortalEditActions.DeleteClaim)
	DeleteClaim(ctx: StateContext<ReturnPortalEditModel>, action: ReturnPortalEditActions.DeleteClaim) {
		const state = produce(ctx.getState(), st => {
			let current = st.returnPortalEditForm.model?.returnPortalClaim || [];
			current = current.filter((c, ci) => ci !== action.payload);
			st.returnPortalEditForm.model!.returnPortalClaim = current;
		});
		ctx.setState(state);
	}

	@Action(ReturnPortalEditActions.Save)
	Save(ctx: StateContext<ReturnPortalEditModel>, action: ReturnPortalEditActions.Save) {
		const state = ctx.getState();

		const input: UpdateReturnPortalInput = {
			name: state.returnPortalEditForm.model?.name,
			returnOpenHours: state.returnPortalEditForm.model?.returnOpenHours,
			automaticallyAccept: state.returnPortalEditForm.model?.automaticallyAccept,
			clearEmailConfirmationLabel: true,
			emailConfirmationLabelID: state.returnPortalEditForm.model?.emailConfirmationLabel?.id,
			emailConfirmationQrCodeID: state.returnPortalEditForm.model?.emailConfirmationQrCode?.id,
			emailReceivedID: state.returnPortalEditForm.model?.emailReceived?.id,
			emailAcceptedID: state.returnPortalEditForm.model?.emailAccepted?.id,
			clearDeliveryOptions: true,
			addDeliveryOptionIDs: state.returnPortalEditForm.model?.deliveryOptions?.map((opt) => opt.id),
		}

		const inputClaims = state.returnPortalEditForm.model?.returnPortalClaim?.map((c) => {
			return {
				id: c.id,
				input: {
					name: c.name,
					description: c.description,
					restockable: c.restockable,
					returnPortalID: state.returnPortalID,
				}
			}
		}) || [];

		return this.update.mutate({id: state.returnPortalID, input: input, inputClaims})
			.subscribe((r) => {
				if (!!r.errors && r.errors.length > 0) {
					this.store.dispatch(new ShowGlobalSnackbar("Error saving"))
				} else {
					this.store.dispatch(new AppChangeRoute({path: Paths.SETTINGS_RETURN_PORTALS_LIST, queryParams: {}}));
				}
			});
	}

	@Action(ReturnPortalEditActions.SetSelectedDeliveryOptions)
	SetSelectedDeliveryOptions(ctx: StateContext<ReturnPortalEditModel>, action: ReturnPortalEditActions.SetSelectedDeliveryOptions) {
		const state = produce(ctx.getState(), st => {
			st.returnPortalEditForm.model!.deliveryOptions = action.payload.map((id) => {
				return {id: id};
			});
		});
		ctx.setState(state);
	}

}
