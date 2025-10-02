import {Injectable} from "@angular/core";
import {Action, Selector, State, StateContext, Store} from "@ngxs/store";
import {
	DeliveryOptionEditPostNordActions
} from "./delivery-option-edit-post-nord.actions";
import {formErrors} from "../../../../account/company-info/company-info.ngxs";
import SelectDeliveryOptionEditPostNordQueryResponse = DeliveryOptionEditPostNordActions.SelectDeliveryOptionEditPostNordQueryResponse;
import {
	AvailableAdditionalServicesPostNordGQL,
	FetchDeliveryOptionEditPostNordGQL, SaveDeliveryOptionEditPostNordGQL
} from "./delivery-option-edit-post-nord.generated";
import {toNotNullArray} from "../../../../functions/not-null-array";
import SetServices = DeliveryOptionEditPostNordActions.SetServices;
import ServicesResponse = DeliveryOptionEditPostNordActions.ServicesResponse;
import SetEditPostNord = DeliveryOptionEditPostNordActions.SetEditPostNord;
import AddedAdditionalServiceResponse = DeliveryOptionEditPostNordActions.AddedAdditionalServiceResponse;
import SetEnabledAdditionalServices = DeliveryOptionEditPostNordActions.SetEnabledAdditionalServices;
import {produce} from "immer";
import {AppActions} from "../../../../app.actions";
import ShowGlobalSnackbar = AppActions.ShowGlobalSnackbar;
import SetAvailableAdditionalServices = DeliveryOptionEditPostNordActions.SetAvailableAdditionalServices;
import AvailableAdditionalServiceResponse = DeliveryOptionEditPostNordActions.AvailableAdditionalServiceResponse;
import LocationResponse = DeliveryOptionEditPostNordActions.LocationResponse;
import SetLocations = DeliveryOptionEditPostNordActions.SetLocations;
import SetSelectedLocations = DeliveryOptionEditPostNordActions.SetSelectedLocations;
import EmailTemplateResponse = DeliveryOptionEditPostNordActions.EmailTemplateResponse;
import SetEmailTemplates = DeliveryOptionEditPostNordActions.SetEmailTemplates;
import {SelectedEmailTemplates} from "../delivery-option-email-templates/delivery-option-email-templates.component";
import SetSelectedEmailTemplates = DeliveryOptionEditPostNordActions.SetSelectedEmailTemplates;
import AppChangeRoute = AppActions.AppChangeRoute;
import {Paths} from "../../../../app-routing.module";

export interface DeliveryOptionEditPostNordModel {
	deliveryOptionEditPostNordForm: {
		model: SelectDeliveryOptionEditPostNordQueryResponse | undefined;
		dirty: boolean;
		status: string;
		errors: formErrors;
	},
	enabledAdditionalServices: AddedAdditionalServiceResponse[];
	availableAdditionalServices: AvailableAdditionalServiceResponse[];
	showUnavailableAdditionalServices: boolean;
	carrierServices: ServicesResponse[],
	selectedID: string;
	selectedLocations: LocationResponse[];
	locations: LocationResponse[];
	emailTemplates: EmailTemplateResponse[];
	selectedEmailTemplates: SelectedEmailTemplates;
}

const defaultState: DeliveryOptionEditPostNordModel = {
	deliveryOptionEditPostNordForm: {
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
@State<DeliveryOptionEditPostNordModel>({
	name: 'deliveryOptionEditPostNord',
	defaults: defaultState,
})
export class DeliveryOptionEditPostNordState {

	constructor(
		private fetch: FetchDeliveryOptionEditPostNordGQL,
		private save: SaveDeliveryOptionEditPostNordGQL,
		private fetchAvailable: AvailableAdditionalServicesPostNordGQL,
		private store: Store,
	) {}

	@Selector()
	static get(state: DeliveryOptionEditPostNordModel) {
		return state;
	}

	@Action(DeliveryOptionEditPostNordActions.Fetch)
	Fetch(ctx: StateContext<DeliveryOptionEditPostNordModel>, action: DeliveryOptionEditPostNordActions.Fetch) {
		const id = ctx.getState().selectedID;
		return this.fetch.fetch({id})
			.subscribe((r) => {
				const services = toNotNullArray(r.data.carrierServices.edges?.map((s) => s?.node));
				this.store.dispatch(new SetServices(services));

				const locations = toNotNullArray(r.data.locations.edges?.map((s) => s?.node));
				this.store.dispatch(new SetLocations(locations));

				const selectedLocations = toNotNullArray(r.data.clickCollectLocations?.deliveryOption.clickCollectLocation?.map((s) => s));
				this.store.dispatch(new SetSelectedLocations(selectedLocations));

				const emails = toNotNullArray(r.data.emailTemplates?.edges?.map((s) => s?.node));
				this.store.dispatch(new SetEmailTemplates(emails));

				const selectedEmails = r.data.selectedEmailTemplates?.deliveryOption;
				if (!!selectedEmails) {
					this.store.dispatch(new SetSelectedEmailTemplates({clickCollectAtStore: selectedEmails.emailClickCollectAtStore?.id}));
				}

				const additionalServices = toNotNullArray(r.data.carrierAdditionalServicePostNords.edges?.map((s) => s?.node));
				this.store.dispatch(new SetEnabledAdditionalServices(additionalServices));

				const edit = r.data.deliveryOptionPostNord?.deliveryOption;
				if (!!edit) {
					this.store.dispatch(new SetEditPostNord(edit));
				}

				this.store.dispatch(new DeliveryOptionEditPostNordActions.FetchAvailableAdditionalServices());

			});
	}

	@Action(DeliveryOptionEditPostNordActions.SetAdditionalServiceEnabled)
	SetAdditionalServiceEnabled(ctx: StateContext<DeliveryOptionEditPostNordModel>, action: DeliveryOptionEditPostNordActions.SetAdditionalServiceEnabled) {
		const state = produce(ctx.getState(), st => {
			st.enabledAdditionalServices = st.enabledAdditionalServices.filter((val) => val.internalID !== action.payload.internalID);
			if (action.payload.checked) {
				st.enabledAdditionalServices.push({internalID: action.payload.internalID})
			}
		});
		ctx.patchState({enabledAdditionalServices: state.enabledAdditionalServices});
	}

	@Action(DeliveryOptionEditPostNordActions.SetServices)
	SetServices(ctx: StateContext<DeliveryOptionEditPostNordModel>, action: DeliveryOptionEditPostNordActions.SetServices) {
		ctx.patchState({carrierServices: action.payload});
	}

	@Action(DeliveryOptionEditPostNordActions.SetID)
	SetID(ctx: StateContext<DeliveryOptionEditPostNordModel>, action: DeliveryOptionEditPostNordActions.SetID) {
		ctx.patchState({selectedID: action.payload});
	}

	@Action(DeliveryOptionEditPostNordActions.SetEnabledAdditionalServices)
	SetEnabledAdditionalServices(ctx: StateContext<DeliveryOptionEditPostNordModel>, action: DeliveryOptionEditPostNordActions.SetEnabledAdditionalServices) {
		ctx.patchState({enabledAdditionalServices: action.payload});
	}

	@Action(DeliveryOptionEditPostNordActions.SetSelectedLocations)
	SetSelectedLocations(ctx: StateContext<DeliveryOptionEditPostNordModel>, action: DeliveryOptionEditPostNordActions.SetSelectedLocations) {
		ctx.patchState({selectedLocations: action.payload});
	}

	@Action(DeliveryOptionEditPostNordActions.SetLocations)
	SetLocations(ctx: StateContext<DeliveryOptionEditPostNordModel>, action: DeliveryOptionEditPostNordActions.SetLocations) {
		ctx.patchState({locations: action.payload});
	}

	@Action(DeliveryOptionEditPostNordActions.SetEditPostNord)
	SetEditPostNord(ctx: StateContext<DeliveryOptionEditPostNordModel>, action: DeliveryOptionEditPostNordActions.SetEditPostNord) {
		const state = produce(ctx.getState(), st => {
			st.deliveryOptionEditPostNordForm.model = action.payload;
		});
		ctx.setState(state);
	}

	@Action(DeliveryOptionEditPostNordActions.SetAvailableAdditionalServices)
	SetAvailableAdditionalServices(ctx: StateContext<DeliveryOptionEditPostNordModel>, action: DeliveryOptionEditPostNordActions.SetAvailableAdditionalServices) {
		ctx.patchState({
			// Remove all selected services not available to this service
			enabledAdditionalServices: ctx.getState().enabledAdditionalServices.filter(item => action.payload.includes(item.internalID)),
			availableAdditionalServices: action.payload
		});
	}

	@Action(DeliveryOptionEditPostNordActions.SetShowUnavailableAdditionalServices)
	SetShowUnavailableAdditionalServices(ctx: StateContext<DeliveryOptionEditPostNordModel>, action: DeliveryOptionEditPostNordActions.SetShowUnavailableAdditionalServices) {
		ctx.patchState({showUnavailableAdditionalServices: action.payload});
	}

	@Action(DeliveryOptionEditPostNordActions.FetchAvailableAdditionalServices)
	FetchAvailableAdditionalServices(ctx: StateContext<DeliveryOptionEditPostNordModel>, action: DeliveryOptionEditPostNordActions.FetchAvailableAdditionalServices) {
		return this.fetchAvailable.fetch({
			carrierServiceID: ctx.getState().deliveryOptionEditPostNordForm.model!.carrierService.id})
			.subscribe((r) => {
				const available = r.data.availableAdditionalServicesPostNord;
				this.store.dispatch(new SetAvailableAdditionalServices(available));
			});
	}

	@Action(DeliveryOptionEditPostNordActions.AddLocation)
	AddLocation(ctx: StateContext<DeliveryOptionEditPostNordModel>, action: DeliveryOptionEditPostNordActions.AddLocation) {
		const state = ctx.getState();
		const next = state.selectedLocations.filter((l) => l.id !== action.payload.id);
		next.push(action.payload);
		ctx.dispatch(new SetSelectedLocations(next));
	}

	@Action(DeliveryOptionEditPostNordActions.RemoveLocation)
	RemoveLocation(ctx: StateContext<DeliveryOptionEditPostNordModel>, action: DeliveryOptionEditPostNordActions.RemoveLocation) {
		const state = ctx.getState();
		const next = state.selectedLocations.filter((l) => l.id !== action.payload);
		ctx.dispatch(new SetSelectedLocations(next));
	}

	@Action(DeliveryOptionEditPostNordActions.SetEmailTemplates)
	SetEmailTemplates(ctx: StateContext<DeliveryOptionEditPostNordModel>, action: DeliveryOptionEditPostNordActions.SetEmailTemplates) {
		ctx.patchState({emailTemplates: action.payload});
	}

	@Action(DeliveryOptionEditPostNordActions.SetSelectedEmailTemplates)
	SetSelectedEmailTemplates(ctx: StateContext<DeliveryOptionEditPostNordModel>, action: DeliveryOptionEditPostNordActions.SetSelectedEmailTemplates) {
		ctx.patchState({selectedEmailTemplates: action.payload});
	}

	@Action(DeliveryOptionEditPostNordActions.Clear)
	Clear(ctx: StateContext<DeliveryOptionEditPostNordModel>, action: DeliveryOptionEditPostNordActions.Clear) {
		ctx.setState(defaultState);
	}

	@Action(DeliveryOptionEditPostNordActions.Save)
	Save(ctx: StateContext<DeliveryOptionEditPostNordModel>, action: DeliveryOptionEditPostNordActions.Save) {
		const state = ctx.getState();

		return this.save.mutate({
			id: state.selectedID,
			input: {},
			inputDeliveryOption: Object.assign({},
				state.deliveryOptionEditPostNordForm.model,
				{carrierServiceID: state.deliveryOptionEditPostNordForm.model?.carrierService.id},
				{carrierService: undefined},
				{
					clearDefaultPackaging: true,
					defaultPackaging: undefined,
					defaultPackagingID: state.deliveryOptionEditPostNordForm.model?.defaultPackaging?.id,
				}),
			inputAdditionalServices: state.enabledAdditionalServices.map((a) => a.internalID),
		}, {errorPolicy: "all"}).subscribe((r) => {
			if (!!r.errors) {
				this.store.dispatch([
					new ShowGlobalSnackbar(`Error: ` + JSON.stringify(r.errors)),
				]);
			} else {
				this.store.dispatch([
					new ShowGlobalSnackbar(`Delivery option saved successfully`),
					new AppChangeRoute({path: Paths.SETTINGS_DELIVERY_OPTIONS, queryParams: {}})
				]);
			}
		})
	}

}
