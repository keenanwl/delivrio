import {Injectable} from "@angular/core";
import {Action, Selector, State, StateContext} from "@ngxs/store";
import SetCarrierEditUSPS = CarrierEditUSPSActions.SetCarrierUSPSEdit;
import SelectCarriersEditQueryResponse = CarrierEditUSPSActions.SelectCarriersEditQueryResponse;
import {AppActions} from "../../../app.actions";
import ShowGlobalSnackbar = AppActions.ShowGlobalSnackbar;
import AppChangeRoute = AppActions.AppChangeRoute;
import {GraphQLError} from "graphql";
import {FormErrorsActions} from "../../../plugins/ngxs-form-errors";
import SetFormErrors = FormErrorsActions.SetFormErrors;
import {Paths} from "../../../app-routing.module";
import {CarrierEditUSPSActions} from "./carrier-edit-usps.actions";
import {FetchCarrierUspsGQL, UpdateCarrierUspsGQL} from "./carrier-edit-usps.generated";

export interface CarrierEditUSPSModel {
	carrierEditUSPSForm: {
		model: SelectCarriersEditQueryResponse | undefined;
		dirty: boolean;
		status: string;
		errors: readonly GraphQLError[];
	},
	id: string;
}

const defaultState: CarrierEditUSPSModel = {
	carrierEditUSPSForm: {
		model: undefined,
		dirty: false,
		status: '',
		errors: [],
	},
	id: '',
};

@Injectable()
@State<CarrierEditUSPSModel>({
	name: 'carrierEditUSPS',
	defaults: defaultState,
})
export class CarrierEditUSPSState {

	constructor(
		private update: UpdateCarrierUspsGQL,
		private fetch: FetchCarrierUspsGQL,
	) {}

	@Selector()
	static get(state: CarrierEditUSPSModel) {
		return state;
	}

	@Action(CarrierEditUSPSActions.FetchCarrierUSPSEdit)
	FetchCarrierUSPSEdit(ctx: StateContext<CarrierEditUSPSModel>, action: CarrierEditUSPSActions.FetchCarrierUSPSEdit) {
		return this.fetch.fetch({carrierID: ctx.getState().id})
			.subscribe({next: (r) => {
				const car = r.data.carrierUSPS;
				if (!!car) {
					ctx.dispatch(new SetCarrierEditUSPS(car));
				}
			}});
	}

	@Action(CarrierEditUSPSActions.SetID)
	SetID(ctx: StateContext<CarrierEditUSPSModel>, action: CarrierEditUSPSActions.SetID) {
		ctx.patchState({id: action.payload});
	}

	@Action(CarrierEditUSPSActions.Clear)
	Clear(ctx: StateContext<CarrierEditUSPSModel>, action: CarrierEditUSPSActions.Clear) {
		ctx.setState(defaultState);
	}

	@Action(CarrierEditUSPSActions.SetCarrierUSPSEdit)
	SetCarrierUSPSEdit(ctx: StateContext<CarrierEditUSPSModel>, action: CarrierEditUSPSActions.SetCarrierUSPSEdit) {
		const state = ctx.getState();
		const next = Object.assign({}, state.carrierEditUSPSForm, {
			model: Object.assign({}, action.payload, {name: action.payload.name})
		});
		ctx.patchState({
			carrierEditUSPSForm: next,
		})
	}

	@Action(CarrierEditUSPSActions.SaveForm)
	SaveForm(ctx: StateContext<CarrierEditUSPSModel>, action: CarrierEditUSPSActions.SaveForm) {
		const state = ctx.getState();

		const body = Object.assign({}, state.carrierEditUSPSForm.model,
		{
			id: undefined,
			carrier: undefined,
		});

		return this.update.mutate({
			id: state.id,
			name: state.carrierEditUSPSForm.model?.carrier.name + '',
			input: body,
		}).subscribe({
			next: (resp) => {

				if (!!resp.errors) {
					ctx.dispatch(new SetFormErrors({errors: resp.errors, formPath: "carrierEditUSPS.carrierEditUSPSForm"}));
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
