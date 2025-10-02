import {Injectable} from "@angular/core";
import {Action, Selector, State, StateContext} from "@ngxs/store";
import {GraphQLError} from "graphql";
import {CarrierEditDSVActions} from "./carrier-edit-dsv.actions";
import EditResponse = CarrierEditDSVActions.EditResponse;
import {produce} from "immer";
import {FetchCarrierEditDsvGQL, UpdateCarrierAgreementDsvGQL} from "./carrier-edit-dsv.generated";
import SetCarrier = CarrierEditDSVActions.SetCarrier;
import {AppActions} from "../../../app.actions";
import ShowGlobalSnackbar = AppActions.ShowGlobalSnackbar;
import AppChangeRoute = AppActions.AppChangeRoute;
import {Paths} from "../../../app-routing.module";

export interface CarrierEditDSVModel {
	form: {
		model: EditResponse | undefined;
		dirty: boolean;
		status: string;
		errors: readonly GraphQLError[];
	},
}

const defaultState: CarrierEditDSVModel = {
	form: {
		model: undefined,
		dirty: false,
		status: '',
		errors: [],
	},
};

@Injectable()
@State<CarrierEditDSVModel>({
	name: 'carrierEditDSV',
	defaults: defaultState,
})
export class CarrierEditDSVState {

	constructor(
		private update: UpdateCarrierAgreementDsvGQL,
		private fetchConn: FetchCarrierEditDsvGQL,
	) {}

	@Selector()
	static get(state: CarrierEditDSVModel) {
		return state;
	}

	@Action(CarrierEditDSVActions.FetchCarrierEditDSV)
	FetchCarrierEditDsv(ctx: StateContext<CarrierEditDSVModel>, action: CarrierEditDSVActions.FetchCarrierEditDSV) {
		return this.fetchConn.fetch({id: action.payload}, {fetchPolicy: "no-cache"})
			.subscribe({next: (r) => {
				const car = r.data.carrier;
				if (!!car) {
					ctx.dispatch(new SetCarrier(car));
				}
			}});
	}

	@Action(CarrierEditDSVActions.SetCarrier)
	SetCarrierEdit(ctx: StateContext<CarrierEditDSVModel>, action: CarrierEditDSVActions.SetCarrier) {
		const state = ctx.getState();
		const next = produce(state.form, (st) => {
			st.model = action.payload;
		});
		ctx.patchState({form: next});
	}

	@Action(CarrierEditDSVActions.SaveForm)
	SaveForm(ctx: StateContext<CarrierEditDSVModel>, action: CarrierEditDSVActions.SaveForm) {
		const state = ctx.getState();

		const input = Object.assign({}, state.form.model?.carrierDSV,
			{
				id: undefined,
			});

		return this.update.mutate(
			{
				id: state.form.model?.id + '',
				name: state.form.model?.name + '',
				input: input,
			},{errorPolicy: "all"}
		).subscribe({
			next: (resp) => {
				if (!!resp.errors) {
					ctx.dispatch([
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
