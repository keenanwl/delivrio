import {Injectable} from "@angular/core";
import {Action, Selector, State, StateContext, Store} from "@ngxs/store";
import {formErrors} from "../../../../account/company-info/company-info.ngxs";
import {toNotNullArray} from "../../../../functions/not-null-array";
import SetServices = DeliveryOptionEditDAOActions.SetServices;
import ServicesResponse = DeliveryOptionEditDAOActions.ServicesResponse;
import SetEditDAO = DeliveryOptionEditDAOActions.SetEditDAO;
import {AppActions} from "../../../../app.actions";
import ShowGlobalSnackbar = AppActions.ShowGlobalSnackbar;
import {DeliveryOptionEditDAOActions} from "./delivery-option-edit-dao.actions";
import {FetchDeliveryOptionEditDaoGQL, SaveDeliveryOptionEditDaoGQL} from "./delivery-option-edit-dao.generated";
import DeliveryOptionEditDAOResponse = DeliveryOptionEditDAOActions.DeliveryOptionEditDAOResponse;
import {produce} from "immer";
import {Paths} from "../../../../app-routing.module";
import AppChangeRoute = AppActions.AppChangeRoute;

export interface DeliveryOptionEditDAOModel {
	form: {
		model: DeliveryOptionEditDAOResponse | undefined;
		dirty: boolean;
		status: string;
		errors: formErrors;
	},
	carrierServices: ServicesResponse[],
	selectedID: string;
}

const defaultState: DeliveryOptionEditDAOModel = {
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
@State<DeliveryOptionEditDAOModel>({
	name: 'deliveryOptionEditDAO',
	defaults: defaultState,
})
export class DeliveryOptionEditDAOState {

	constructor(
		private fetch: FetchDeliveryOptionEditDaoGQL,
		private save: SaveDeliveryOptionEditDaoGQL,
		private store: Store,
	) {}

	@Selector()
	static get(state: DeliveryOptionEditDAOModel) {
		return state;
	}

	@Action(DeliveryOptionEditDAOActions.Fetch)
	Fetch(ctx: StateContext<DeliveryOptionEditDAOModel>, action: DeliveryOptionEditDAOActions.Fetch) {
		const id = ctx.getState().selectedID;
		return this.fetch.fetch({id})
			.subscribe((r) => {
				const services = toNotNullArray(r.data.carrierServices.edges?.map((s) => s?.node));
				if (!!services) {
					this.store.dispatch(new SetServices(services));
				}

				const edit = r.data.deliveryOptionDAO?.deliveryOption;
				if (!!edit) {
					this.store.dispatch(new SetEditDAO(edit));
				}
			});
	}

	@Action(DeliveryOptionEditDAOActions.SetServices)
	SetServices(ctx: StateContext<DeliveryOptionEditDAOModel>, action: DeliveryOptionEditDAOActions.SetServices) {
		ctx.patchState({carrierServices: action.payload});
	}

	@Action(DeliveryOptionEditDAOActions.SetID)
	SetID(ctx: StateContext<DeliveryOptionEditDAOModel>, action: DeliveryOptionEditDAOActions.SetID) {
		ctx.patchState({selectedID: action.payload});
	}

	@Action(DeliveryOptionEditDAOActions.SetEditDAO)
	SetEditDAO(ctx: StateContext<DeliveryOptionEditDAOModel>, action: DeliveryOptionEditDAOActions.SetEditDAO) {
		const state = ctx.getState();
		const next = produce(state, (st) => {
			st.form.model = action.payload;
		})
		ctx.setState(next);
	}

	@Action(DeliveryOptionEditDAOActions.Clear)
	Clear(ctx: StateContext<DeliveryOptionEditDAOModel>, action: DeliveryOptionEditDAOActions.Clear) {
		ctx.setState(defaultState);
	}

	@Action(DeliveryOptionEditDAOActions.Save)
	Save(ctx: StateContext<DeliveryOptionEditDAOModel>, action: DeliveryOptionEditDAOActions.Save) {
		const state = ctx.getState();
		return this.save.mutate({
			id: state.selectedID,
			inputDeliveryOption: {
				carrierServiceID: state.form.model?.carrierService.id,
				name: state.form.model?.name || '',
				description: state.form.model?.description,
				hideDeliveryOption: state.form.model?.hideDeliveryOption || false,
				clickOptionDisplayCount: state.form.model?.clickOptionDisplayCount || 3,
				clickCollect: state.form.model?.clickCollect || false,
				clearClickCollectLocation: true,
				overrideReturnAddress: state.form.model?.overrideReturnAddress || false,
				overrideSenderAddress: state.form.model?.overrideSenderAddress || false,
				deliveryEstimateFrom: state.form.model?.deliveryEstimateFrom,
				deliveryEstimateTo: state.form.model?.deliveryEstimateTo,
			},
		}).subscribe((r) => {
			this.store.dispatch([
				new ShowGlobalSnackbar(`Delivery option saved successfully`),
				new AppChangeRoute({path: Paths.SETTINGS_DELIVERY_OPTIONS, queryParams: {}})
			]);
		})
	}

}
