import {Injectable} from "@angular/core";
import {Action, Selector, State, StateContext} from "@ngxs/store";
import {GraphQLError} from "graphql";
import {CarrierEditDFActions} from "./carrier-edit-df.actions";
import EditResponse = CarrierEditDFActions.EditResponse;
import {produce} from "immer";
import {
	FetchCarrierEditDfGQL,
	UpdateCarrierAgreementDfGQL
} from "./carrier-edit-df.generated";
import SetCarrier = CarrierEditDFActions.SetCarrier;
import {AppActions} from "../../../app.actions";
import ShowGlobalSnackbar = AppActions.ShowGlobalSnackbar;
import AppChangeRoute = AppActions.AppChangeRoute;
import {Paths} from "../../../app-routing.module";

export interface CarrierEditDFModel {
	form: {
		model: EditResponse | undefined;
		dirty: boolean;
		status: string;
		errors: readonly GraphQLError[];
	},
}

const defaultState: CarrierEditDFModel = {
	form: {
		model: undefined,
		dirty: false,
		status: '',
		errors: [],
	},
};

@Injectable()
@State<CarrierEditDFModel>({
	name: 'carrierEditDF',
	defaults: defaultState,
})
export class CarrierEditDFState {

	constructor(
		private update: UpdateCarrierAgreementDfGQL,
		private fetchConn: FetchCarrierEditDfGQL,
	) {}

	@Selector()
	static get(state: CarrierEditDFModel) {
		return state;
	}

	@Action(CarrierEditDFActions.FetchCarrierEditDF)
	FetchCarrierEditDao(ctx: StateContext<CarrierEditDFModel>, action: CarrierEditDFActions.FetchCarrierEditDF) {
		return this.fetchConn.fetch({id: action.payload}, {fetchPolicy: "no-cache"})
			.subscribe({next: (r) => {
				const car = r.data.carrier;
				if (!!car) {
					ctx.dispatch(new SetCarrier(car));
				}
			}});
	}

	@Action(CarrierEditDFActions.SetCarrier)
	SetCarrierEdit(ctx: StateContext<CarrierEditDFModel>, action: CarrierEditDFActions.SetCarrier) {
		const state = ctx.getState();
		const next = produce(state.form, (st) => {
			st.model = action.payload;
		});
		ctx.patchState({form: next});
	}

	@Action(CarrierEditDFActions.SaveForm)
	SaveForm(ctx: StateContext<CarrierEditDFModel>, action: CarrierEditDFActions.SaveForm) {
		const state = ctx.getState();
		const input = Object.assign({}, state.form.model?.carrierDF,
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
