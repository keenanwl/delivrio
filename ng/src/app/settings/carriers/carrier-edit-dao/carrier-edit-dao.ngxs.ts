import {Injectable} from "@angular/core";
import {Action, Selector, State, StateContext} from "@ngxs/store";
import {GraphQLError} from "graphql";
import {CarrierEditDAOActions} from "./carrier-edit-dao.actions";
import EditResponse = CarrierEditDAOActions.EditResponse;
import {produce} from "immer";
import {FetchCarrierEditDaoGQL, UpdateCarrierAgreementDaoGQL} from "./carrier-edit-dao.generated";
import SetCarrier = CarrierEditDAOActions.SetCarrier;
import {AppActions} from "../../../app.actions";
import ShowGlobalSnackbar = AppActions.ShowGlobalSnackbar;
import AppChangeRoute = AppActions.AppChangeRoute;
import {Paths} from "../../../app-routing.module";

export interface CarrierEditDAOModel {
	form: {
		model: EditResponse | undefined;
		dirty: boolean;
		status: string;
		errors: readonly GraphQLError[];
	},
}

const defaultState: CarrierEditDAOModel = {
	form: {
		model: undefined,
		dirty: false,
		status: '',
		errors: [],
	},
};

@Injectable()
@State<CarrierEditDAOModel>({
	name: 'carrierEditDAO',
	defaults: defaultState,
})
export class CarrierEditDAOState {

	constructor(
		private update: UpdateCarrierAgreementDaoGQL,
		private fetchConn: FetchCarrierEditDaoGQL,
	) {}

	@Selector()
	static get(state: CarrierEditDAOModel) {
		return state;
	}

	@Action(CarrierEditDAOActions.FetchCarrierEditDAO)
	FetchCarrierEditDao(ctx: StateContext<CarrierEditDAOModel>, action: CarrierEditDAOActions.FetchCarrierEditDAO) {
		return this.fetchConn.fetch({id: action.payload}, {fetchPolicy: "no-cache"})
			.subscribe({next: (r) => {
				const car = r.data.carrier;
				if (!!car) {
					ctx.dispatch(new SetCarrier(car));
				}
			}});
	}

	@Action(CarrierEditDAOActions.SetCarrier)
	SetCarrierEdit(ctx: StateContext<CarrierEditDAOModel>, action: CarrierEditDAOActions.SetCarrier) {
		const state = ctx.getState();
		const next = produce(state.form, (st) => {
			st.model = action.payload;
		});
		ctx.patchState({form: next});
	}

	@Action(CarrierEditDAOActions.SaveForm)
	SaveForm(ctx: StateContext<CarrierEditDAOModel>, action: CarrierEditDAOActions.SaveForm) {
		const state = ctx.getState();

		const input = Object.assign({}, state.form.model?.carrierDAO,
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
