import {Injectable} from "@angular/core";
import {Action, Selector, State, StateContext, Store} from "@ngxs/store";
import {formErrors} from "../../../../account/company-info/company-info.ngxs";
import {toNotNullArray} from "../../../../functions/not-null-array";
import SetServices = DeliveryOptionEditBringActions.SetServices;
import ServicesResponse = DeliveryOptionEditBringActions.ServicesResponse;
import SetEditBring = DeliveryOptionEditBringActions.SetEditBring;
import {AppActions} from "../../../../app.actions";
import ShowGlobalSnackbar = AppActions.ShowGlobalSnackbar;
import {DeliveryOptionEditBringActions} from "./delivery-option-edit-bring.actions";
import {FetchDeliveryOptionEditBringGQL, SaveDeliveryOptionEditBringGQL} from "./delivery-option-edit-bring.generated";
import DeliveryOptionEditBringResponse = DeliveryOptionEditBringActions.DeliveryOptionEditBringResponse;
import {produce} from "immer";
import {Paths} from "../../../../app-routing.module";
import AppChangeRoute = AppActions.AppChangeRoute;

export interface DeliveryOptionEditBringModel {
	form: {
		model: DeliveryOptionEditBringResponse | undefined;
		dirty: boolean;
		status: string;
		errors: formErrors;
	},
	carrierServices: ServicesResponse[],
	selectedID: string;
}

const defaultState: DeliveryOptionEditBringModel = {
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
@State<DeliveryOptionEditBringModel>({
	name: 'deliveryOptionEditBring',
	defaults: defaultState,
})
export class DeliveryOptionEditBringState {

	constructor(
		private fetch: FetchDeliveryOptionEditBringGQL,
		private save: SaveDeliveryOptionEditBringGQL,
		private store: Store,
	) {}

	@Selector()
	static get(state: DeliveryOptionEditBringModel) {
		return state;
	}

	@Action(DeliveryOptionEditBringActions.Fetch)
	Fetch(ctx: StateContext<DeliveryOptionEditBringModel>, action: DeliveryOptionEditBringActions.Fetch) {
		const id = ctx.getState().selectedID;
		return this.fetch.fetch({id})
			.subscribe((r) => {
				const services = toNotNullArray(r.data.carrierServices.edges?.map((s) => s?.node));
				if (!!services) {
					this.store.dispatch(new SetServices(services));
				}

				const edit = r.data.deliveryOptionBring?.deliveryOption;
				if (!!edit) {
					this.store.dispatch(new SetEditBring(edit));
				}
			});
	}

	@Action(DeliveryOptionEditBringActions.SetServices)
	SetServices(ctx: StateContext<DeliveryOptionEditBringModel>, action: DeliveryOptionEditBringActions.SetServices) {
		ctx.patchState({carrierServices: action.payload});
	}

	@Action(DeliveryOptionEditBringActions.SetID)
	SetID(ctx: StateContext<DeliveryOptionEditBringModel>, action: DeliveryOptionEditBringActions.SetID) {
		ctx.patchState({selectedID: action.payload});
	}

	@Action(DeliveryOptionEditBringActions.SetEditBring)
	SetEditBring(ctx: StateContext<DeliveryOptionEditBringModel>, action: DeliveryOptionEditBringActions.SetEditBring) {
		const state = ctx.getState();
		const next = produce(state, (st) => {
			st.form.model = action.payload;
		})
		ctx.setState(next);
	}

	@Action(DeliveryOptionEditBringActions.Clear)
	Clear(ctx: StateContext<DeliveryOptionEditBringModel>, action: DeliveryOptionEditBringActions.Clear) {
		ctx.setState(defaultState);
	}

	@Action(DeliveryOptionEditBringActions.Save)
	Save(ctx: StateContext<DeliveryOptionEditBringModel>, action: DeliveryOptionEditBringActions.Save) {
		const state = ctx.getState();

		return this.save.mutate({
			id: state.selectedID,
			input: {
				electronicCustoms: state.form.model?.deliveryOptionBring?.electronicCustoms,
			},
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
