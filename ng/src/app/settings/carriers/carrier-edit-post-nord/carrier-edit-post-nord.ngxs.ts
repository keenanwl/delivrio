import {Injectable} from "@angular/core";
import {Action, Selector, State, StateContext} from "@ngxs/store";
import SetCarrierEditPostNord = CarrierEditPostNordActions.SetCarrierPostNordEdit;
import SelectCarriersEditQueryResponse = CarrierEditPostNordActions.SelectCarriersEditQueryResponse;
import {AppActions} from "../../../app.actions";
import ShowGlobalSnackbar = AppActions.ShowGlobalSnackbar;
import AppChangeRoute = AppActions.AppChangeRoute;
import {GraphQLError} from "graphql";
import {FormErrorsActions} from "../../../plugins/ngxs-form-errors";
import SetFormErrors = FormErrorsActions.SetFormErrors;
import {CarrierEditPostNordActions} from "./carrier-edit-post-nord.actions";
import {FetchCarrierPostNordGQL, UpdateCarrierPostNordGQL} from "./carrier-edit-post-nord.generated";
import {Paths} from "../../../app-routing.module";

export interface CarrierEditPostNordModel {
	carrierEditPostNordForm: {
		model: SelectCarriersEditQueryResponse | undefined;
		dirty: boolean;
		status: string;
		errors: readonly GraphQLError[];
	},
	id: string;
}

const defaultState: CarrierEditPostNordModel = {
	carrierEditPostNordForm: {
		model: undefined,
		dirty: false,
		status: '',
		errors: [],
	},
	id: '',
};

@Injectable()
@State<CarrierEditPostNordModel>({
	name: 'carrierEditPostNord',
	defaults: defaultState,
})
export class CarrierEditPostNordState {

	constructor(
		private updateCarrierPostNord: UpdateCarrierPostNordGQL,
		private fetchConn: FetchCarrierPostNordGQL,
	) {}

	@Selector()
	static get(state: CarrierEditPostNordModel) {
		return state;
	}

	@Action(CarrierEditPostNordActions.FetchCarrierPostNordEdit)
	FetchCarrierPostNordEdit(ctx: StateContext<CarrierEditPostNordModel>, action: CarrierEditPostNordActions.FetchCarrierPostNordEdit) {
		return this.fetchConn.fetch({carrierID: ctx.getState().id})
			.subscribe({next: (r) => {
				const car = r.data.carrierPostNord;
				if (!!car) {
					ctx.dispatch(new SetCarrierEditPostNord(car));
				}
			}});
	}

	@Action(CarrierEditPostNordActions.SetID)
	SetID(ctx: StateContext<CarrierEditPostNordModel>, action: CarrierEditPostNordActions.SetID) {
		ctx.patchState({id: action.payload});
	}

	@Action(CarrierEditPostNordActions.Clear)
	Clear(ctx: StateContext<CarrierEditPostNordModel>, action: CarrierEditPostNordActions.Clear) {
		ctx.setState(defaultState);
	}

	@Action(CarrierEditPostNordActions.SetCarrierPostNordEdit)
	SetCarrierPostNordEdit(ctx: StateContext<CarrierEditPostNordModel>, action: CarrierEditPostNordActions.SetCarrierPostNordEdit) {
		const state = ctx.getState();
		const next = Object.assign({}, state.carrierEditPostNordForm, {
			model: Object.assign({}, action.payload, {name: action.payload.name})
		});
		ctx.patchState({
			carrierEditPostNordForm: next,
		})
	}

	@Action(CarrierEditPostNordActions.SaveForm)
	SaveForm(ctx: StateContext<CarrierEditPostNordModel>, action: CarrierEditPostNordActions.SaveForm) {
		const state = ctx.getState();

		const body = Object.assign({}, state.carrierEditPostNordForm.model,
		{
			id: undefined,
			carrier: undefined,
		});

		return this.updateCarrierPostNord.mutate({
			id: state.id,
			name: state.carrierEditPostNordForm.model?.carrier.name + '',
			input: body,
		}).subscribe({
			next: (resp) => {

				if (!!resp.errors) {
					ctx.dispatch(new SetFormErrors({errors: resp.errors, formPath: "carrierEditPostNord.carrierEditPostNordForm"}));
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
