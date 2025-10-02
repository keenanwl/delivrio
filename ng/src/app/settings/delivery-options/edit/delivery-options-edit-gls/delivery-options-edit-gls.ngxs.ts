import {Injectable} from "@angular/core";
import {Action, Selector, State, StateContext} from "@ngxs/store";
import {formErrors} from "../../../../account/company-info/company-info.ngxs";
import {DeliveryOptionsGLSEditActions} from "./delivery-options-edit-gls.actions";
import {FormErrorsActions} from "../../../../plugins/ngxs-form-errors";
import {AppActions} from "../../../../app.actions";
import SetDeliveryOptionsGLSEdit = DeliveryOptionsGLSEditActions.SetDeliveryOptionsGLSEdit;
import SelectDeliveryOptionsGLSEditQueryResponse = DeliveryOptionsGLSEditActions.SelectDeliveryOptionsGLSEditQueryResponse;
import SetFormErrors = FormErrorsActions.SetFormErrors;
import ShowGlobalSnackbar = AppActions.ShowGlobalSnackbar;
import {
	FetchDeliveryOptionsGlsEditGQL,
	UpdateDeliveryOptionGlsGQL
} from "./delivery-options-edit-gls.generated";
import GLSServicesResponse = DeliveryOptionsGLSEditActions.GLSServicesResponse;
import {toNotNullArray} from "../../../../functions/not-null-array";
import SetServices = DeliveryOptionsGLSEditActions.SetServices;
import {DeliveryOptionEditPostNordModel} from "../delivery-option-edit-post-nord/delivery-option-edit-post-nord.ngxs";
import {Paths} from "../../../../app-routing.module";
import AppChangeRoute = AppActions.AppChangeRoute;
import {produce} from "immer";

export interface DeliveryOptionsGLSEditModel {
	deliveryOptionsGLSEditForm: {
		model: SelectDeliveryOptionsGLSEditQueryResponse | undefined;
		dirty: boolean;
		status: string;
		errors: formErrors;
	},
	carrierServices: GLSServicesResponse[],
	selectedOption: string;
}

const defaultState: DeliveryOptionsGLSEditModel = {
	deliveryOptionsGLSEditForm: {
		model: undefined,
		dirty: false,
		status: '',
		errors: {}
	},
	selectedOption: '',
	carrierServices: [],
};

@Injectable()
@State<DeliveryOptionsGLSEditModel>({
	name: 'deliveryOptionsEditGLS',
	defaults: defaultState,
})
export class DeliveryOptionGLSEditState {

	constructor(
		private fetchEdit: FetchDeliveryOptionsGlsEditGQL,
		private updateGLS: UpdateDeliveryOptionGlsGQL,
	) {}

	@Selector()
	static get(state: DeliveryOptionsGLSEditModel) {
		return state;
	}

	@Action(DeliveryOptionsGLSEditActions.FetchDeliveryOptionsGLSEdit)
	FetchMyDeliveryOptionsGLSEdit(ctx: StateContext<DeliveryOptionsGLSEditModel>, action: DeliveryOptionsGLSEditActions.FetchDeliveryOptionsGLSEdit) {
		const state = ctx.getState();
		return this.fetchEdit.fetch({id: state.selectedOption})
			.subscribe({next: (r) => {

				const services = toNotNullArray(r.data.carrierServices.edges?.map((s) => s?.node));
				ctx.dispatch(new SetServices(services));

				const d = r.data.deliveryOptionGLS?.deliveryOption;
				if (!!d) {
					ctx.dispatch([
						new SetDeliveryOptionsGLSEdit(d),
					]);
				}
			}});
	}

	@Action(DeliveryOptionsGLSEditActions.SetDeliveryOptionsGLSEdit)
	SetDeliveryOptionsGLSEdit(ctx: StateContext<DeliveryOptionsGLSEditModel>, action: DeliveryOptionsGLSEditActions.SetDeliveryOptionsGLSEdit) {
		const state = produce(ctx.getState(), st => {
			st.deliveryOptionsGLSEditForm.model = action.payload;
		});
		ctx.setState(state);
	}

	@Action(DeliveryOptionsGLSEditActions.SetSelectedOption)
	SetSelectedOption(ctx: StateContext<DeliveryOptionsGLSEditModel>, action: DeliveryOptionsGLSEditActions.SetSelectedOption) {
		ctx.patchState({selectedOption: action.payload})
	}

	@Action(DeliveryOptionsGLSEditActions.Clear)
	Clear(ctx: StateContext<DeliveryOptionsGLSEditModel>, action: DeliveryOptionsGLSEditActions.Clear) {
		ctx.setState(defaultState);
	}

	@Action(DeliveryOptionsGLSEditActions.SetServices)
	SetServices(ctx: StateContext<DeliveryOptionEditPostNordModel>, action: DeliveryOptionsGLSEditActions.SetServices) {
		ctx.patchState({carrierServices: action.payload});
	}

	@Action(DeliveryOptionsGLSEditActions.SaveForm)
	SaveForm(ctx: StateContext<DeliveryOptionsGLSEditModel>, action: DeliveryOptionsGLSEditActions.SaveForm) {
		return this.updateGLS.mutate(action.payload, {errorPolicy: "all"})
			.subscribe((res) => {
				if (!!res.errors) {
					ctx.dispatch([
						new SetFormErrors({errors: res.errors, formPath: 'deliveryOptionsGLSEdit.deliveryOptionsGLSEditForm'}),
						new SetFormErrors({errors: res.errors, formPath: 'deliveryOptionsGLSEdit.deliveryOptionPriceForm'}),
						new ShowGlobalSnackbar("Please fix the highlighted errors and try saving again"),
					]);
				} else {
					ctx.dispatch([
						new ShowGlobalSnackbar(`Delivery option saved successfully`),
						new AppChangeRoute({path: Paths.SETTINGS_DELIVERY_OPTIONS, queryParams: {}})
					]);
				}
			});
	}

}
