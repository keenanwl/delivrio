import {Injectable} from "@angular/core";
import {Action, Selector, State, StateContext} from "@ngxs/store";
import {GraphQLError} from "graphql";
import {CarrierEditBringActions} from "./carrier-edit-bring.actions";
import EditResponse = CarrierEditBringActions.EditResponse;
import {produce} from "immer";
import {FetchCarrierEditBringGQL, UpdateCarrierAgreementBringGQL} from "./carrier-edit-bring.generated";
import SetCarrier = CarrierEditBringActions.SetCarrier;
import {AppActions} from "../../../app.actions";
import ShowGlobalSnackbar = AppActions.ShowGlobalSnackbar;
import AppChangeRoute = AppActions.AppChangeRoute;
import {Paths} from "../../../app-routing.module";

export interface CarrierEditBringModel {
	form: {
		model: EditResponse | undefined;
		dirty: boolean;
		status: string;
		errors: readonly GraphQLError[];
	},
}

const defaultState: CarrierEditBringModel = {
	form: {
		model: undefined,
		dirty: false,
		status: '',
		errors: [],
	},
};

@Injectable()
@State<CarrierEditBringModel>({
	name: 'carrierEditBring',
	defaults: defaultState,
})
export class CarrierEditBringState {

	constructor(
		private update: UpdateCarrierAgreementBringGQL,
		private fetchConn: FetchCarrierEditBringGQL,
	) {}

	@Selector()
	static get(state: CarrierEditBringModel) {
		return state;
	}

	@Action(CarrierEditBringActions.FetchCarrierEditBring)
	FetchCarrierEditBring(ctx: StateContext<CarrierEditBringModel>, action: CarrierEditBringActions.FetchCarrierEditBring) {
		return this.fetchConn.fetch({id: action.payload}, {fetchPolicy: "no-cache"})
			.subscribe({next: (r) => {
				const car = r.data.carrier;
				if (!!car) {
					ctx.dispatch(new SetCarrier(car));
				}
			}});
	}

	@Action(CarrierEditBringActions.SetCarrier)
	SetCarrierEdit(ctx: StateContext<CarrierEditBringModel>, action: CarrierEditBringActions.SetCarrier) {
		const state = ctx.getState();
		const next = produce(state.form, (st) => {
			st.model = action.payload;
		});
		ctx.patchState({form: next});
	}

	@Action(CarrierEditBringActions.SaveForm)
	SaveForm(ctx: StateContext<CarrierEditBringModel>, action: CarrierEditBringActions.SaveForm) {
		const state = ctx.getState();

		const input = Object.assign({}, state.form.model?.carrierBring,
			{
				id: undefined,
			});

		return this.update.mutate(
			{
				id: state.form.model?.id + '',
				name: state.form.model?.name + '',
				input: input,
			}, {
				errorPolicy: "all",
			}).subscribe({
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
