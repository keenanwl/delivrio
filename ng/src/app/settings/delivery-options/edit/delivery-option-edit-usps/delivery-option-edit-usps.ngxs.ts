import {Injectable} from "@angular/core";
import {Action, Selector, State, StateContext, Store} from "@ngxs/store";
import {formErrors} from "../../../../account/company-info/company-info.ngxs";
import {BaseDeliveryOptionFragment} from "../edit-common.generated";
import {toNotNullArray} from "../../../../functions/not-null-array";
import SetServices = DeliveryOptionEditUSPSActions.SetServices;
import ServicesResponse = DeliveryOptionEditUSPSActions.ServicesResponse;
import SetEditUSPS = DeliveryOptionEditUSPSActions.SetEditUSPS;
import SetEnabledAdditionalServices = DeliveryOptionEditUSPSActions.SetEnabledAdditionalServices;
import {produce} from "immer";
import {AppActions} from "../../../../app.actions";
import ShowGlobalSnackbar = AppActions.ShowGlobalSnackbar;
import SetAvailableAdditionalServices = DeliveryOptionEditUSPSActions.SetAvailableAdditionalServices;
import LocationResponse = DeliveryOptionEditUSPSActions.LocationResponse;
import SetLocations = DeliveryOptionEditUSPSActions.SetLocations;
import SetSelectedLocations = DeliveryOptionEditUSPSActions.SetSelectedLocations;
import EmailTemplateResponse = DeliveryOptionEditUSPSActions.EmailTemplateResponse;
import SetEmailTemplates = DeliveryOptionEditUSPSActions.SetEmailTemplates;
import {SelectedEmailTemplates} from "../delivery-option-email-templates/delivery-option-email-templates.component";
import SetSelectedEmailTemplates = DeliveryOptionEditUSPSActions.SetSelectedEmailTemplates;
import {DeliveryOptionEditUSPSActions} from "./delivery-option-edit-usps.actions";
import {
	AvailableAdditionalServicesUspsGQL,
	FetchDeliveryOptionEditUspsGQL,
	SaveDeliveryOptionEditUspsGQL, UspsAdditionalServicesFragment
} from "./delivery-option-edit-usps.generated";
import {Paths} from "../../../../app-routing.module";
import AppChangeRoute = AppActions.AppChangeRoute;

export interface DeliveryOptionEditUSPSModel {
	deliveryOptionEditUSPSForm: {
		model: BaseDeliveryOptionFragment | undefined;
		dirty: boolean;
		status: string;
		errors: formErrors;
	},
	enabledAdditionalServices: UspsAdditionalServicesFragment[];
	availableAdditionalServices: UspsAdditionalServicesFragment[];
	showUnavailableAdditionalServices: boolean;
	carrierServices: ServicesResponse[],
	selectedID: string;
	selectedLocations: LocationResponse[];
	locations: LocationResponse[];
	emailTemplates: EmailTemplateResponse[];
	selectedEmailTemplates: SelectedEmailTemplates;
}

const defaultState: DeliveryOptionEditUSPSModel = {
	deliveryOptionEditUSPSForm: {
		model: undefined,
		dirty: false,
		status: '',
		errors: {}
	},
	enabledAdditionalServices: [],
	availableAdditionalServices: [],
	showUnavailableAdditionalServices: false,
	carrierServices: [],
	selectedID: "",
	locations: [],
	selectedLocations: [],
	emailTemplates: [],
	selectedEmailTemplates: {
		clickCollectAtStore: undefined,
	},
};

@Injectable()
@State<DeliveryOptionEditUSPSModel>({
	name: 'deliveryOptionEditUSPS',
	defaults: defaultState,
})
export class DeliveryOptionEditUSPSState {

	constructor(
		private fetch: FetchDeliveryOptionEditUspsGQL,
		private save: SaveDeliveryOptionEditUspsGQL,
		private fetchAvailable: AvailableAdditionalServicesUspsGQL,
		private store: Store,
	) {}

	@Selector()
	static get(state: DeliveryOptionEditUSPSModel) {
		return state;
	}

	@Action(DeliveryOptionEditUSPSActions.Fetch)
	Fetch(ctx: StateContext<DeliveryOptionEditUSPSModel>, action: DeliveryOptionEditUSPSActions.Fetch) {
		const id = ctx.getState().selectedID;
		return this.fetch.fetch({id})
			.subscribe((r) => {
				const services = toNotNullArray(r.data.carrierServices.edges?.map((s) => s?.node));
				if (!!services) {
					this.store.dispatch(new SetServices(services));
				}

				const locations = toNotNullArray(r.data.locations.edges?.map((s) => s?.node));
				if (!!locations) {
					this.store.dispatch(new SetLocations(locations));
				}

				const selectedLocations = toNotNullArray(r.data.clickCollectLocations?.deliveryOption.clickCollectLocation?.map((s) => s));
				if (!!selectedLocations) {
					this.store.dispatch(new SetSelectedLocations(selectedLocations));
				}

				const emails = toNotNullArray(r.data.emailTemplates?.edges?.map((s) => s?.node));
				if (!!emails) {
					this.store.dispatch(new SetEmailTemplates(emails));
				}

				const selectedEmails = r.data.selectedEmailTemplates?.deliveryOption;
				if (!!selectedEmails) {
					this.store.dispatch(new SetSelectedEmailTemplates({clickCollectAtStore: selectedEmails.emailClickCollectAtStore?.id}));
				}

				const additionalServices = toNotNullArray(r.data.carrierAdditionalServiceUspSs.edges?.map((s) => s?.node));
				if (!!additionalServices) {
					this.store.dispatch(new SetEnabledAdditionalServices(additionalServices));
				}

				const edit = r.data.deliveryOptionUSPS?.deliveryOption;
				if (!!edit) {
					this.store.dispatch(new SetEditUSPS(edit));
				}

				this.store.dispatch(new DeliveryOptionEditUSPSActions.FetchAvailableAdditionalServices());

			});
	}

	@Action(DeliveryOptionEditUSPSActions.SetAdditionalServiceEnabled)
	SetAdditionalServiceEnabled(ctx: StateContext<DeliveryOptionEditUSPSModel>, action: DeliveryOptionEditUSPSActions.SetAdditionalServiceEnabled) {
		const state = produce(ctx.getState(), st => {
			st.enabledAdditionalServices = st.enabledAdditionalServices.filter((val) => val.internalID !== action.payload.internalID);
			st.enabledAdditionalServices.push(action.payload)
		});
		ctx.patchState({enabledAdditionalServices: state.enabledAdditionalServices});
	}

	@Action(DeliveryOptionEditUSPSActions.SetAdditionalServiceDisabled)
	SetAdditionalServiceDisabled(ctx: StateContext<DeliveryOptionEditUSPSModel>, action: DeliveryOptionEditUSPSActions.SetAdditionalServiceDisabled) {
		const state = produce(ctx.getState(), st => {
			st.enabledAdditionalServices = st.enabledAdditionalServices.filter((val) => val.internalID !== action.payload.internalID);
		});
		ctx.patchState({enabledAdditionalServices: state.enabledAdditionalServices});
	}

	@Action(DeliveryOptionEditUSPSActions.SetServices)
	SetServices(ctx: StateContext<DeliveryOptionEditUSPSModel>, action: DeliveryOptionEditUSPSActions.SetServices) {
		ctx.patchState({carrierServices: action.payload});
	}

	@Action(DeliveryOptionEditUSPSActions.SetID)
	SetID(ctx: StateContext<DeliveryOptionEditUSPSModel>, action: DeliveryOptionEditUSPSActions.SetID) {
		ctx.patchState({selectedID: action.payload});
	}

	@Action(DeliveryOptionEditUSPSActions.SetEnabledAdditionalServices)
	SetEnabledAdditionalServices(ctx: StateContext<DeliveryOptionEditUSPSModel>, action: DeliveryOptionEditUSPSActions.SetEnabledAdditionalServices) {
		ctx.patchState({enabledAdditionalServices: action.payload});
	}

	@Action(DeliveryOptionEditUSPSActions.SetSelectedLocations)
	SetSelectedLocations(ctx: StateContext<DeliveryOptionEditUSPSModel>, action: DeliveryOptionEditUSPSActions.SetSelectedLocations) {
		ctx.patchState({selectedLocations: action.payload});
	}

	@Action(DeliveryOptionEditUSPSActions.SetLocations)
	SetLocations(ctx: StateContext<DeliveryOptionEditUSPSModel>, action: DeliveryOptionEditUSPSActions.SetLocations) {
		ctx.patchState({locations: action.payload});
	}

	@Action(DeliveryOptionEditUSPSActions.SetEditUSPS)
	SetEditUSPS(ctx: StateContext<DeliveryOptionEditUSPSModel>, action: DeliveryOptionEditUSPSActions.SetEditUSPS) {
		const state = produce(ctx.getState(), st => {
			st.deliveryOptionEditUSPSForm.model = action.payload;
		});
		ctx.setState(state);
	}

	@Action(DeliveryOptionEditUSPSActions.SetAvailableAdditionalServices)
	SetAvailableAdditionalServices(ctx: StateContext<DeliveryOptionEditUSPSModel>, action: DeliveryOptionEditUSPSActions.SetAvailableAdditionalServices) {
		ctx.patchState({
			// Enable this filtering when we have it on the backend
			// So only available services are still active
			// enabledAdditionalServices: [],
			availableAdditionalServices: action.payload
		});
	}

	@Action(DeliveryOptionEditUSPSActions.SetShowUnavailableAdditionalServices)
	SetShowUnavailableAdditionalServices(ctx: StateContext<DeliveryOptionEditUSPSModel>, action: DeliveryOptionEditUSPSActions.SetShowUnavailableAdditionalServices) {
		ctx.patchState({showUnavailableAdditionalServices: action.payload});
	}

	@Action(DeliveryOptionEditUSPSActions.FetchAvailableAdditionalServices)
	FetchAvailableAdditionalServices(ctx: StateContext<DeliveryOptionEditUSPSModel>, action: DeliveryOptionEditUSPSActions.FetchAvailableAdditionalServices) {
		return this.fetchAvailable.fetch({
			carrierServiceID: ctx.getState().deliveryOptionEditUSPSForm.model?.carrierService?.id || ""})
			.subscribe((r) => {
				const available = r.data.availableAdditionalServicesUSPS;
				this.store.dispatch(new SetAvailableAdditionalServices(available));
			});
	}

	@Action(DeliveryOptionEditUSPSActions.AddLocation)
	AddLocation(ctx: StateContext<DeliveryOptionEditUSPSModel>, action: DeliveryOptionEditUSPSActions.AddLocation) {
		const state = ctx.getState();
		const next = state.selectedLocations.filter((l) => l.id !== action.payload.id);
		next.push(action.payload);
		ctx.dispatch(new SetSelectedLocations(next));
	}

	@Action(DeliveryOptionEditUSPSActions.RemoveLocation)
	RemoveLocation(ctx: StateContext<DeliveryOptionEditUSPSModel>, action: DeliveryOptionEditUSPSActions.RemoveLocation) {
		const state = ctx.getState();
		const next = state.selectedLocations.filter((l) => l.id !== action.payload);
		ctx.dispatch(new SetSelectedLocations(next));
	}

	@Action(DeliveryOptionEditUSPSActions.SetEmailTemplates)
	SetEmailTemplates(ctx: StateContext<DeliveryOptionEditUSPSModel>, action: DeliveryOptionEditUSPSActions.SetEmailTemplates) {
		ctx.patchState({emailTemplates: action.payload});
	}

	@Action(DeliveryOptionEditUSPSActions.SetSelectedEmailTemplates)
	SetSelectedEmailTemplates(ctx: StateContext<DeliveryOptionEditUSPSModel>, action: DeliveryOptionEditUSPSActions.SetSelectedEmailTemplates) {
		ctx.patchState({selectedEmailTemplates: action.payload});
	}

	@Action(DeliveryOptionEditUSPSActions.Clear)
	Clear(ctx: StateContext<DeliveryOptionEditUSPSModel>, action: DeliveryOptionEditUSPSActions.Clear) {
		ctx.setState(defaultState);
	}

	@Action(DeliveryOptionEditUSPSActions.Save)
	Save(ctx: StateContext<DeliveryOptionEditUSPSModel>, action: DeliveryOptionEditUSPSActions.Save) {
		const state = ctx.getState();

		return this.save.mutate({
			id: state.selectedID,
			input: {},
			inputDeliveryOption: {
				carrierServiceID: state.deliveryOptionEditUSPSForm.model?.carrierService.id,
				name: state.deliveryOptionEditUSPSForm.model?.name || '',
				description: state.deliveryOptionEditUSPSForm.model?.description,
				hideDeliveryOption: state.deliveryOptionEditUSPSForm.model?.hideDeliveryOption || false,
				clickOptionDisplayCount: state.deliveryOptionEditUSPSForm.model?.clickOptionDisplayCount || 3,
				clickCollect: state.deliveryOptionEditUSPSForm.model?.clickCollect || false,
				clearClickCollectLocation: true,
				addClickCollectLocationIDs: state.selectedLocations.map((l) => l.id),
				overrideReturnAddress: state.deliveryOptionEditUSPSForm.model?.overrideReturnAddress || false,
				overrideSenderAddress: state.deliveryOptionEditUSPSForm.model?.overrideSenderAddress || false,
				emailClickCollectAtStoreID: state.selectedEmailTemplates.clickCollectAtStore,
				deliveryEstimateFrom: state.deliveryOptionEditUSPSForm.model?.deliveryEstimateFrom,
				deliveryEstimateTo: state.deliveryOptionEditUSPSForm.model?.deliveryEstimateTo,
				clearDefaultPackaging: true,
				defaultPackagingID: state.deliveryOptionEditUSPSForm.model?.defaultPackaging?.id,
				shipmondoIntegration: state.deliveryOptionEditUSPSForm.model?.shipmondoIntegration,
				shipmondoDeliveryOption: state.deliveryOptionEditUSPSForm.model?.shipmondoDeliveryOption,
				webshipperIntegration: state.deliveryOptionEditUSPSForm.model?.webshipperIntegration,
				webshipperID: state.deliveryOptionEditUSPSForm.model?.webshipperID,
			},
			inputAdditionalServices: state.enabledAdditionalServices.map((a) => a.id),
		}).subscribe((r) => {
			this.store.dispatch([
				new ShowGlobalSnackbar(`Delivery option saved successfully`),
				new AppChangeRoute({path: Paths.SETTINGS_DELIVERY_OPTIONS, queryParams: {}})
			]);
		})
	}

}
