import {Injectable} from "@angular/core";
import {Action, Selector, State, StateContext} from "@ngxs/store";
import SelectCarriersEditQueryResponse = CarrierEditEasyPostActions.SelectCarriersEditQueryResponse;
import SetCarrierEasyPostEdit = CarrierEditEasyPostActions.SetCarrierEasyPostEdit;
import {AppActions} from "../../../app.actions";
import ShowGlobalSnackbar = AppActions.ShowGlobalSnackbar;
import AppChangeRoute = AppActions.AppChangeRoute;
import {GraphQLError} from "graphql";
import {FormErrorsActions} from "../../../plugins/ngxs-form-errors";
import {Paths} from "../../../app-routing.module";
import {FetchCarrierEasyPostGQL, UpdateCarrierEasyPostGQL} from "./carrier-edit-easy-post.generated";
import {CarrierEditEasyPostActions} from "./carrier-edit-easy-post.actions";

export interface CarrierEditEasyPostModel {
	carrierEditEasyPostForm: {
		model: SelectCarriersEditQueryResponse | undefined;
		dirty: boolean;
		status: string;
		errors: readonly GraphQLError[];
	},
	id: string;
}

const defaultState: CarrierEditEasyPostModel = {
	carrierEditEasyPostForm: {
		model: undefined,
		dirty: false,
		status: '',
		errors: [],
	},
	id: '',
};

@Injectable()
@State<CarrierEditEasyPostModel>({
	name: 'carrierEditEasyPost',
	defaults: defaultState,
})
export class CarrierEditEasyPostState {

	constructor(
		private updateCarrierEasyPost: UpdateCarrierEasyPostGQL,
		private fetchConn: FetchCarrierEasyPostGQL,
	) {}

	@Selector()
	static get(state: CarrierEditEasyPostModel) {
		return state;
	}

	@Action(CarrierEditEasyPostActions.FetchCarrierEasyPostEdit)
	FetchCarrierEasyPostEdit(ctx: StateContext<CarrierEditEasyPostModel>, action: CarrierEditEasyPostActions.FetchCarrierEasyPostEdit) {
		return this.fetchConn.fetch({carrierID: ctx.getState().id})
			.subscribe({next: (r) => {
				const car = r.data.carrier;
				if (!!car) {
					ctx.dispatch(new SetCarrierEasyPostEdit(car));
				}
			}});
	}

	@Action(CarrierEditEasyPostActions.SetID)
	SetID(ctx: StateContext<CarrierEditEasyPostModel>, action: CarrierEditEasyPostActions.SetID) {
		ctx.patchState({id: action.payload});
	}

	@Action(CarrierEditEasyPostActions.Clear)
	Clear(ctx: StateContext<CarrierEditEasyPostModel>, action: CarrierEditEasyPostActions.Clear) {
		ctx.setState(defaultState);
	}

	@Action(CarrierEditEasyPostActions.SetCarrierEasyPostEdit)
	SetCarrierEasyPostEdit(ctx: StateContext<CarrierEditEasyPostModel>, action: CarrierEditEasyPostActions.SetCarrierEasyPostEdit) {
		const state = ctx.getState();
		const next = Object.assign({}, state.carrierEditEasyPostForm, {
			model: Object.assign({}, action.payload, {name: action.payload.name})
		});
		ctx.patchState({
			carrierEditEasyPostForm: next,
		})
	}

	@Action(CarrierEditEasyPostActions.SaveForm)
	SaveForm(ctx: StateContext<CarrierEditEasyPostModel>, action: CarrierEditEasyPostActions.SaveForm) {
		const state = ctx.getState();

		return this.updateCarrierEasyPost.mutate({
			id: state.id,
			name: state.carrierEditEasyPostForm.model?.name || '',
			input: {
				apiKey: state.carrierEditEasyPostForm.model?.carrierEasyPost?.apiKey,
				// It's a hack array vs string, seems to work with just 1 ID right now
				carrierAccounts: state.carrierEditEasyPostForm.model?.carrierEasyPost?.carrierAccounts,
				test: state.carrierEditEasyPostForm.model?.carrierEasyPost?.test,
			},
		}).subscribe({
			next: (resp) => {

				if (!!resp.errors) {
					ctx.dispatch(new ShowGlobalSnackbar("Error: " + JSON.stringify(resp.errors)));
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
