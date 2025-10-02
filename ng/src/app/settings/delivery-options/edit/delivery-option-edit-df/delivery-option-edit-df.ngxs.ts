import {Injectable} from "@angular/core";
import {Action, Selector, State, StateContext, Store} from "@ngxs/store";
import {formErrors} from "../../../../account/company-info/company-info.ngxs";
import {BaseDeliveryOptionFragment} from "../edit-common.generated";
import {toNotNullArray} from "../../../../functions/not-null-array";
import SetServices = DeliveryOptionEditDFActions.SetServices;
import ServicesResponse = DeliveryOptionEditDFActions.ServicesResponse;
import SetEditDF = DeliveryOptionEditDFActions.SetEditDF;
import {AppActions} from "../../../../app.actions";
import ShowGlobalSnackbar = AppActions.ShowGlobalSnackbar;
import {DeliveryOptionEditDFActions} from "./delivery-option-edit-df.actions";
import {FetchDeliveryOptionEditDfGQL, SaveDeliveryOptionEditDfGQL} from "./delivery-option-edit-df.generated";
import {produce} from "immer";
import {Paths} from "../../../../app-routing.module";
import AppChangeRoute = AppActions.AppChangeRoute;

export interface DeliveryOptionEditDFModel {
	form: {
		model: BaseDeliveryOptionFragment | undefined;
		dirty: boolean;
		status: string;
		errors: formErrors;
	},
	carrierServices: ServicesResponse[],
	selectedID: string;
}

const defaultState: DeliveryOptionEditDFModel = {
	form: {
		model: undefined,
		dirty: false,
		status: '',
		errors: {}
	},
	carrierServices: [],
	selectedID: "",
};

@Injectable()
@State<DeliveryOptionEditDFModel>({
	name: 'deliveryOptionEditDF',
	defaults: defaultState,
})
export class DeliveryOptionEditDFState {

	constructor(
		private fetch: FetchDeliveryOptionEditDfGQL,
		private save: SaveDeliveryOptionEditDfGQL,
		private store: Store,
	) {}

	@Selector()
	static get(state: DeliveryOptionEditDFModel) {
		return state;
	}

	@Action(DeliveryOptionEditDFActions.Fetch)
	Fetch(ctx: StateContext<DeliveryOptionEditDFModel>, action: DeliveryOptionEditDFActions.Fetch) {
		const id = ctx.getState().selectedID;
		return this.fetch.fetch({id})
			.subscribe((r) => {
				const services = toNotNullArray(r.data.carrierServices.edges?.map((s) => s?.node));
				if (!!services) {
					this.store.dispatch(new SetServices(services));
				}

				const edit = r.data.deliveryOptionDF?.deliveryOption;
				if (!!edit) {
					this.store.dispatch(new SetEditDF(edit));
				}
			});
	}

	@Action(DeliveryOptionEditDFActions.SetServices)
	SetServices(ctx: StateContext<DeliveryOptionEditDFModel>, action: DeliveryOptionEditDFActions.SetServices) {
		ctx.patchState({carrierServices: action.payload});
	}

	@Action(DeliveryOptionEditDFActions.SetID)
	SetID(ctx: StateContext<DeliveryOptionEditDFModel>, action: DeliveryOptionEditDFActions.SetID) {
		ctx.patchState({selectedID: action.payload});
	}

	@Action(DeliveryOptionEditDFActions.SetEditDF)
	SetEditDF(ctx: StateContext<DeliveryOptionEditDFModel>, action: DeliveryOptionEditDFActions.SetEditDF) {
		const state = ctx.getState();
		const next = produce(state, (st) => {
			st.form.model = action.payload;
		})
		ctx.setState(next);
	}

	@Action(DeliveryOptionEditDFActions.Clear)
	Clear(ctx: StateContext<DeliveryOptionEditDFModel>, action: DeliveryOptionEditDFActions.Clear) {
		ctx.setState(defaultState);
	}

	@Action(DeliveryOptionEditDFActions.Save)
	Save(ctx: StateContext<DeliveryOptionEditDFModel>, action: DeliveryOptionEditDFActions.Save) {
		const state = ctx.getState();

		return this.save.mutate({
			id: state.selectedID,
			inputDeliveryOption: Object.assign({},
				state.form.model,
				{carrierServiceID: state.form.model?.carrierService.id},
				{carrierService: undefined},
				{
					clearDefaultPackaging: true,
					defaultPackaging: undefined,
					defaultPackagingID: state.form.model?.defaultPackaging?.id,
				}
			),
		}).subscribe((r) => {
			this.store.dispatch([
				new ShowGlobalSnackbar(`Delivery option saved successfully`),
				new AppChangeRoute({path: Paths.SETTINGS_DELIVERY_OPTIONS, queryParams: {}})
			]);
		})
	}

}
