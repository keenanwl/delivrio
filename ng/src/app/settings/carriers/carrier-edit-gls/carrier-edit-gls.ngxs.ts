import {Injectable} from "@angular/core";
import {Action, Selector, State, StateContext} from "@ngxs/store";
import {CarrierEditGLSActions} from "./carrier-edit-gls.actions";
import SetCarrierEdit = CarrierEditGLSActions.SetCarrierEdit;
import SelectCarriersEditQueryResponse = CarrierEditGLSActions.CarriersEditGLSQueryResponse;
import {AppActions} from "../../../app.actions";
import ShowGlobalSnackbar = AppActions.ShowGlobalSnackbar;
import AppChangeRoute = AppActions.AppChangeRoute;
import {GraphQLError} from "graphql";
import {FormErrorsActions} from "../../../plugins/ngxs-form-errors";
import SetFormErrors = FormErrorsActions.SetFormErrors;
import {FetchCarrierEditGlsGQL, UpdateCarrierAgreementGlsGQL} from "./carrier-edit-gls.generated";
import {Paths} from "src/app/app-routing.module";

export interface CarrierEditModel {
	carrierEditForm: {
		model: SelectCarriersEditQueryResponse | undefined;
		dirty: boolean;
		status: string;
		errors: readonly GraphQLError[];
	},
}

const defaultState: CarrierEditModel = {
	carrierEditForm: {
		model: undefined,
		dirty: false,
		status: '',
		errors: [],
	},
};

@Injectable()
@State<CarrierEditModel>({
	name: 'carrierEditGLS',
	defaults: defaultState,
})
export class CarrierEditGLSState {

	constructor(
		private updateCarrierGLS: UpdateCarrierAgreementGlsGQL,
		private fetchConn: FetchCarrierEditGlsGQL,
	) {}

	@Selector()
	static get(state: CarrierEditModel) {
		return state;
	}

	@Action(CarrierEditGLSActions.FetchCarrierEdit)
	FetchMyCarrierEdit(ctx: StateContext<CarrierEditModel>, action: CarrierEditGLSActions.FetchCarrierEdit) {
		return this.fetchConn.fetch({id: action.payload}, {fetchPolicy: "no-cache"})
			.subscribe({next: (r) => {
				const car = r.data.carrier;
				if (!!car) {
					ctx.dispatch(new SetCarrierEdit(car));
				}
			}});
	}

	@Action(CarrierEditGLSActions.SetCarrierEdit)
	SetMyCarrierEdit(ctx: StateContext<CarrierEditModel>, action: CarrierEditGLSActions.SetCarrierEdit) {
		const state = ctx.getState();
		const next = Object.assign({}, state.carrierEditForm, {
			model: Object.assign({}, action.payload, {name: action.payload.name})
		});
		ctx.patchState({
			carrierEditForm: next,
		});
	}

	@Action(CarrierEditGLSActions.SaveForm)
	SaveForm(ctx: StateContext<CarrierEditModel>, action: CarrierEditGLSActions.SaveForm) {
		const state = ctx.getState();
		const input = Object.assign({}, state.carrierEditForm.model?.carrierGLS,
		{
			id: undefined,
			name: undefined,
		});

		return this.updateCarrierGLS.mutate({
			id: state.carrierEditForm.model?.id + '',
			name: state.carrierEditForm.model?.name + '',
			input: input,
		}, {errorPolicy: "all"}).subscribe({
			next: (resp) => {

				if (!!resp.errors) {
					ctx.dispatch([
						new SetFormErrors({errors: resp.errors, formPath: "carrierEditGLS.carrierEditForm"}),
						new ShowGlobalSnackbar(resp.errors.join(" "))
					]);
				} else {
					ctx.dispatch([
						new ShowGlobalSnackbar(`Carrier saved`),
						new AppChangeRoute({path: Paths.SETTINGS_CARRIERS, queryParams: {}}),
					]);
				}
			},
		});

	}

}
