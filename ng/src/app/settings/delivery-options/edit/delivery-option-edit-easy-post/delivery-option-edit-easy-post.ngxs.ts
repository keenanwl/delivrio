import {Injectable} from "@angular/core";
import {Action, Selector, State, StateContext, Store} from "@ngxs/store";
import {formErrors} from "../../../../account/company-info/company-info.ngxs";
import {toNotNullArray} from "../../../../functions/not-null-array";
import SetServices = DeliveryOptionEditEasyPostActions.SetServices;
import ServicesResponse = DeliveryOptionEditEasyPostActions.ServicesResponse;
import SetEditEasyPost = DeliveryOptionEditEasyPostActions.SetEditEasyPost;
import {AppActions} from "../../../../app.actions";
import ShowGlobalSnackbar = AppActions.ShowGlobalSnackbar;
import DeliveryOptionEditEasyPostResponse = DeliveryOptionEditEasyPostActions.DeliveryOptionEditEasyPostResponse;
import {produce} from "immer";
import {Paths} from "../../../../app-routing.module";
import AppChangeRoute = AppActions.AppChangeRoute;
import {
	FetchAdditionalServiceEasyPostGQL,
	FetchDeliveryOptionEditEasyPostGQL,
	SaveDeliveryOptionEditEasyPostGQL
} from "./delivery-option-edit-easy-post.generated";
import {
	DeliveryOptionEditEasyPostActions
} from "./delivery-option-edit-easy-post.actions";
import AdditionalServiceResponse = DeliveryOptionEditEasyPostActions.AdditionalServiceResponse;
import SetAdditionalServices = DeliveryOptionEditEasyPostActions.SetAdditionalServices;
import SetSelectedAdditionalService = DeliveryOptionEditEasyPostActions.SetSelectedAdditionalService;
import FetchAdditionalServices = DeliveryOptionEditEasyPostActions.FetchAdditionalServices;

export interface DeliveryOptionEditEasyPostModel {
	form: {
		model: DeliveryOptionEditEasyPostResponse | undefined;
		dirty: boolean;
		status: string;
		errors: formErrors;
	},
	carrierServices: ServicesResponse[],
	additionalServices: AdditionalServiceResponse[],
	selectedAdditionalServices: string[],
	selectedID: string;
}

const defaultState: DeliveryOptionEditEasyPostModel = {
	form: {
		model: undefined,
		dirty: false,
		status: '',
		errors: {}
	},
	carrierServices: [],
	additionalServices: [],
	selectedAdditionalServices: [],
	selectedID: "",
};

@Injectable()
@State<DeliveryOptionEditEasyPostModel>({
	name: 'deliveryOptionEditEasyPost',
	defaults: defaultState,
})
export class DeliveryOptionEditEasyPostState {

	constructor(
		private fetch: FetchDeliveryOptionEditEasyPostGQL,
		private additionalServices: FetchAdditionalServiceEasyPostGQL,
		private save: SaveDeliveryOptionEditEasyPostGQL,
		private store: Store,
	) {}

	@Selector()
	static get(state: DeliveryOptionEditEasyPostModel) {
		return state;
	}

	@Action(DeliveryOptionEditEasyPostActions.Fetch)
	Fetch(ctx: StateContext<DeliveryOptionEditEasyPostModel>, action: DeliveryOptionEditEasyPostActions.Fetch) {
		const id = ctx.getState().selectedID;
		return this.fetch.fetch({id})
			.subscribe((r) => {
				const services = toNotNullArray(r.data.carrierServices.edges?.map((s) => s?.node));
				this.store.dispatch(new SetServices(services));

				const edit = r.data.deliveryOptionEasyPost?.deliveryOption;
				if (!!edit) {
					this.store.dispatch(new SetEditEasyPost(edit));
					// This should be called on each change to the service,
					// But right now, all services have all additional services
					ctx.dispatch(new FetchAdditionalServices(edit?.carrierService.id))
				}

				const addServices = toNotNullArray(r.data.deliveryOptionEasyPost?.carrierAddServEasyPost?.map((n) => n.id));
				this.store.dispatch(new SetSelectedAdditionalService(addServices));

			});
	}

	@Action(DeliveryOptionEditEasyPostActions.FetchAdditionalServices)
	FetchAdditionalServices(ctx: StateContext<DeliveryOptionEditEasyPostModel>, action: DeliveryOptionEditEasyPostActions.FetchAdditionalServices) {
		return this.additionalServices.fetch({carrierServiceID: action.payload})
			.subscribe((r) => {
				const services = toNotNullArray(r.data.carrierAdditionalServiceEasyPosts.edges?.map((s) => s?.node));
				this.store.dispatch(new SetAdditionalServices(services));
			});
	}

	@Action(DeliveryOptionEditEasyPostActions.SetAdditionalServices)
	SetAdditionalServices(ctx: StateContext<DeliveryOptionEditEasyPostModel>, action: DeliveryOptionEditEasyPostActions.SetAdditionalServices) {
		ctx.patchState({additionalServices: action.payload});
	}

	@Action(DeliveryOptionEditEasyPostActions.SetSelectedAdditionalService)
	SetSelectedAdditionalService(ctx: StateContext<DeliveryOptionEditEasyPostModel>, action: DeliveryOptionEditEasyPostActions.SetSelectedAdditionalService) {
		ctx.patchState({selectedAdditionalServices: action.payload});
	}

	@Action(DeliveryOptionEditEasyPostActions.ToggleAdditionalService)
	ToggleAdditionalService(ctx: StateContext<DeliveryOptionEditEasyPostModel>, action: DeliveryOptionEditEasyPostActions.ToggleAdditionalService) {
		const state = ctx.getState();
		if (action.payload.isAdd) {
			const next = [...state.selectedAdditionalServices];
			next.push(action.payload.id);
			ctx.patchState({
				selectedAdditionalServices: next,
			});
		} else {
			ctx.patchState({
				selectedAdditionalServices: state.selectedAdditionalServices
					.filter((a) => a !== action.payload.id)});
		}
	}

	@Action(DeliveryOptionEditEasyPostActions.SetServices)
	SetServices(ctx: StateContext<DeliveryOptionEditEasyPostModel>, action: DeliveryOptionEditEasyPostActions.SetServices) {
		ctx.patchState({carrierServices: action.payload});
	}

	@Action(DeliveryOptionEditEasyPostActions.SetID)
	SetID(ctx: StateContext<DeliveryOptionEditEasyPostModel>, action: DeliveryOptionEditEasyPostActions.SetID) {
		ctx.patchState({selectedID: action.payload});
	}

	@Action(DeliveryOptionEditEasyPostActions.SetEditEasyPost)
	SetEditEasyPost(ctx: StateContext<DeliveryOptionEditEasyPostModel>, action: DeliveryOptionEditEasyPostActions.SetEditEasyPost) {
		const state = ctx.getState();
		const next = produce(state, (st) => {
			st.form.model = action.payload;
		})
		ctx.setState(next);
	}

	@Action(DeliveryOptionEditEasyPostActions.Clear)
	Clear(ctx: StateContext<DeliveryOptionEditEasyPostModel>, action: DeliveryOptionEditEasyPostActions.Clear) {
		ctx.setState(defaultState);
	}

	@Action(DeliveryOptionEditEasyPostActions.Save)
	Save(ctx: StateContext<DeliveryOptionEditEasyPostModel>, action: DeliveryOptionEditEasyPostActions.Save) {
		const state = ctx.getState();
		return this.save.mutate({
			id: state.selectedID,
			input: {
				clearCarrierAddServEasyPost: true,
				addCarrierAddServEasyPostIDs: state.selectedAdditionalServices,
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
				defaultPackagingID: state.form.model?.defaultPackaging?.id,
				customsEnabled: state.form.model?.customsEnabled,
				customsSigner: state.form.model?.customsSigner,
			},
		}).subscribe((r) => {
			this.store.dispatch([
				new ShowGlobalSnackbar(`Delivery option saved successfully`),
				new AppChangeRoute({path: Paths.SETTINGS_DELIVERY_OPTIONS, queryParams: {}})
			]);
		})
	}

}
