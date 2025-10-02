import {Injectable} from "@angular/core";
import {Action, Selector, State, StateContext, Store} from "@ngxs/store";
import {formErrors} from "../../../../account/company-info/company-info.ngxs";
import {toNotNullArray} from "../../../../functions/not-null-array";
import SetServices = DeliveryOptionEditDSVActions.SetServices;
import ServicesResponse = DeliveryOptionEditDSVActions.ServicesResponse;
import SetEditDSV = DeliveryOptionEditDSVActions.SetEditDSV;
import {AppActions} from "../../../../app.actions";
import ShowGlobalSnackbar = AppActions.ShowGlobalSnackbar;
import {DeliveryOptionEditDSVActions} from "./delivery-option-edit-dsv.actions";
import {FetchDeliveryOptionEditDsvGQL, SaveDeliveryOptionEditDsvGQL} from "./delivery-option-edit-dsv.generated";
import DeliveryOptionEditDSVResponse = DeliveryOptionEditDSVActions.DeliveryOptionEditDSVResponse;
import {produce} from "immer";
import {Paths} from "../../../../app-routing.module";
import AppChangeRoute = AppActions.AppChangeRoute;

export interface DeliveryOptionEditDSVModel {
	form: {
		model: DeliveryOptionEditDSVResponse | undefined;
		dirty: boolean;
		status: string;
		errors: formErrors;
	},
	carrierServices: ServicesResponse[],
	selectedID: string;
}

const defaultState: DeliveryOptionEditDSVModel = {
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
@State<DeliveryOptionEditDSVModel>({
	name: 'deliveryOptionEditDSV',
	defaults: defaultState,
})
export class DeliveryOptionEditDSVState {

	constructor(
		private fetch: FetchDeliveryOptionEditDsvGQL,
		private save: SaveDeliveryOptionEditDsvGQL,
		private store: Store,
	) {}

	@Selector()
	static get(state: DeliveryOptionEditDSVModel) {
		return state;
	}

	@Action(DeliveryOptionEditDSVActions.Fetch)
	Fetch(ctx: StateContext<DeliveryOptionEditDSVModel>, action: DeliveryOptionEditDSVActions.Fetch) {
		const id = ctx.getState().selectedID;
		return this.fetch.fetch({id})
			.subscribe((r) => {
				const services = toNotNullArray(r.data.carrierServices.edges?.map((s) => s?.node));
				if (!!services) {
					this.store.dispatch(new SetServices(services));
				}

				const edit = r.data.deliveryOptionDSV?.deliveryOption;
				if (!!edit) {
					this.store.dispatch(new SetEditDSV(edit));
				}
			});
	}

	@Action(DeliveryOptionEditDSVActions.SetServices)
	SetServices(ctx: StateContext<DeliveryOptionEditDSVModel>, action: DeliveryOptionEditDSVActions.SetServices) {
		ctx.patchState({carrierServices: action.payload});
	}

	@Action(DeliveryOptionEditDSVActions.SetID)
	SetID(ctx: StateContext<DeliveryOptionEditDSVModel>, action: DeliveryOptionEditDSVActions.SetID) {
		ctx.patchState({selectedID: action.payload});
	}

	@Action(DeliveryOptionEditDSVActions.SetEditDSV)
	SetEditDSV(ctx: StateContext<DeliveryOptionEditDSVModel>, action: DeliveryOptionEditDSVActions.SetEditDSV) {
		const state = ctx.getState();
		const next = produce(state, (st) => {
			st.form.model = action.payload;
		})
		ctx.setState(next);
	}

	@Action(DeliveryOptionEditDSVActions.Clear)
	Clear(ctx: StateContext<DeliveryOptionEditDSVModel>, action: DeliveryOptionEditDSVActions.Clear) {
		ctx.setState(defaultState);
	}

	@Action(DeliveryOptionEditDSVActions.Save)
	Save(ctx: StateContext<DeliveryOptionEditDSVModel>, action: DeliveryOptionEditDSVActions.Save) {
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
