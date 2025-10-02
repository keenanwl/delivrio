import {Injectable} from "@angular/core";
import {Action, Selector, State, StateContext} from "@ngxs/store";
import {ConnectionEditActions} from "./connection-edit.actions";
import SetConnectionEdit = ConnectionEditActions.SetConnectionEdit;
import {formErrors} from "../../../account/company-info/company-info.ngxs";
import SelectConnectionsEditQueryResponse = ConnectionEditActions.SelectConnectionsEditQueryResponse;
import {AppActions} from "../../../app.actions";
import AppChangeRoute = AppActions.AppChangeRoute;
import SelectConnectionBrandsQueryResponse = ConnectionEditActions.SelectConnectionBrandsQueryResponse;
import {Paths} from "../../../app-routing.module";
import {ConnectionsListActions} from "../connections-list/connections-list.actions";
import {
	ConnectionBrandsGQL, ConnectionDeliveryOptionFragment,
	CreateConnectionShopifyGQL,
	FetchConnectionGQL,
	UpdateConnectionShopifyGQL
} from "./connection-edit.generated";
import {toNotNullArray} from "../../../functions/not-null-array";
import SetLocations = ConnectionEditActions.SetLocations;
import LocationsResponse = ConnectionEditActions.LocationsResponse;
import {produce} from "immer";
import SetDeliveryOptions = ConnectionEditActions.SetDeliveryOptions;
import SetDocs = ConnectionEditActions.SetDocs;
import DocsResponse = ConnectionEditActions.DocsResponse;
import SetCurrencies = ConnectionEditActions.SetCurrencies;
import CurrencyResponse = ConnectionEditActions.CurrencyResponse;
import ShowGlobalSnackbar = AppActions.ShowGlobalSnackbar;

export interface ConnectionEditModel {
	connectionEditForm: {
		model: SelectConnectionsEditQueryResponse | undefined;
		dirty: boolean;
		status: string;
		errors: formErrors;
	},
	connectionShopifyID: string;
	brands: SelectConnectionBrandsQueryResponse[];
	locations: LocationsResponse[];
	deliveryOptions: ConnectionDeliveryOptionFragment[];
	packingSlipDocs: DocsResponse[];
	currencies: CurrencyResponse[];
}

const defaultState: ConnectionEditModel = {
	connectionEditForm: {
		model: undefined,
		dirty: false,
		status: '',
		errors: {}
	},
	connectionShopifyID: '',
	brands: [],
	locations: [],
	deliveryOptions: [],
	packingSlipDocs: [],
	currencies: [],
};

@Injectable()
@State<ConnectionEditModel>({
	name: 'connectionEdit',
	defaults: defaultState,
})
export class ConnectionEditState {

	constructor(
		private fetchConn: FetchConnectionGQL,
		private fetchBrands: ConnectionBrandsGQL,
		private createShopify: CreateConnectionShopifyGQL,
		private updateShopify: UpdateConnectionShopifyGQL,
	) {}

	@Selector()
	static get(state: ConnectionEditModel) {
		return state;
	}

	@Action(ConnectionEditActions.FetchConnectionEdit)
	FetchMyConnectionEdit(ctx: StateContext<ConnectionEditModel>, action: ConnectionEditActions.FetchConnectionEdit) {
		const id = ctx.getState().connectionShopifyID;
		return this.fetchConn.fetch({id})
			.subscribe({next: (r) => {
				const locations = toNotNullArray(r.data.locations.edges?.map((l) => l?.node));
				const docs = toNotNullArray(r.data.documents.edges?.map((n) => n?.node));
				const deliveryOptions = toNotNullArray(r.data.deliveryOptions.edges?.map((n) => n?.node));
				const currencies = toNotNullArray(r.data.currencies.edges?.map((n) => n?.node));
				ctx.dispatch([
					new SetLocations(locations),
					new SetDeliveryOptions(deliveryOptions),
					new SetDocs(docs),
					new SetCurrencies(currencies),
				]);

				const conn = r.data.connection;
				if (!!conn) {
					ctx.dispatch(new SetConnectionEdit(conn));
				}
			}});
	}

	@Action(ConnectionEditActions.SetDeliveryOptions)
	SetDeliveryOptions(ctx: StateContext<ConnectionEditModel>, action: ConnectionEditActions.SetDeliveryOptions) {
		ctx.patchState({deliveryOptions: action.payload});
	}

	@Action(ConnectionEditActions.SetDocs)
	SetDocs(ctx: StateContext<ConnectionEditModel>, action: ConnectionEditActions.SetDocs) {
		ctx.patchState({packingSlipDocs: action.payload});
	}

	@Action(ConnectionEditActions.SetConnectionEdit)
	SetMyConnectionEdit(ctx: StateContext<ConnectionEditModel>, action: ConnectionEditActions.SetConnectionEdit) {
		const state = ctx.getState();
		const next = Object.assign({}, state.connectionEditForm, {model: action.payload});
		ctx.patchState({
			connectionEditForm: next,
		})
	}

	@Action(ConnectionEditActions.FetchConnectionBrands)
	FetchConnectionBrands(ctx: StateContext<ConnectionEditModel>, action: ConnectionEditActions.FetchConnectionBrands) {
		return this.fetchBrands.fetch()
			.subscribe({next: (r) => {
				const brands: SelectConnectionBrandsQueryResponse[] = [];
				r.data.connectionBrands.edges?.forEach((r) => {
					if (!!r) {
						brands.push(r)
					}
				});
				if (!!brands) {
					ctx.dispatch(new ConnectionEditActions.SetConnectionBrands(brands));
				}
			}});
	}

	@Action(ConnectionEditActions.SetConnectionBrands)
	SetConnectionBrands(ctx: StateContext<ConnectionEditModel>, action: ConnectionEditActions.SetConnectionBrands) {
		ctx.patchState({
			brands: action.payload,
		})
	}

	@Action(ConnectionEditActions.SetConnectionID)
	SetConnectionID(ctx: StateContext<ConnectionEditModel>, action: ConnectionEditActions.SetConnectionID) {
		ctx.patchState({
			connectionShopifyID: action.payload,
		})
	}

	@Action(ConnectionEditActions.SetCurrencies)
	SetCurrencies(ctx: StateContext<ConnectionEditModel>, action: ConnectionEditActions.SetCurrencies) {
		ctx.patchState({
			currencies: action.payload,
		})
	}

	@Action(ConnectionEditActions.SaveForm)
	SaveForm(ctx: StateContext<ConnectionEditModel>, action: ConnectionEditActions.SaveForm) {
		return this.createShopify.mutate(action.payload).subscribe((r) => {
			ctx.dispatch([
				new ConnectionsListActions.FetchConnectionsList(),
				new AppChangeRoute({path: Paths.SETTINGS_CONNECTIONS, queryParams: {}}),
			]);
		});
	}

	@Action(ConnectionEditActions.SaveFormUpdate)
	SaveFormUpdate(ctx: StateContext<ConnectionEditModel>, action: ConnectionEditActions.SaveFormUpdate) {
		return this.updateShopify.mutate(action.payload, {errorPolicy: "all"})
			.subscribe((r) => {
				if (!!r.errors) {
					ctx.dispatch(new ShowGlobalSnackbar("Errors: " + JSON.stringify(r.errors)));
				} else {
					ctx.dispatch([
						new ConnectionsListActions.FetchConnectionsList(),
						new AppChangeRoute({path: Paths.SETTINGS_CONNECTIONS, queryParams: {}}),
					]);
				}
			});
	}

	@Action(ConnectionEditActions.SetLocations)
	SetLocations(ctx: StateContext<ConnectionEditModel>, action: ConnectionEditActions.SetLocations) {
		ctx.patchState({locations: action.payload});
	}

	@Action(ConnectionEditActions.Clear)
	Clear(ctx: StateContext<ConnectionEditModel>, action: ConnectionEditActions.Clear) {
		ctx.setState(defaultState);
	}

	@Action(ConnectionEditActions.UpdateLocations)
	UpdateLocations(ctx: StateContext<ConnectionEditModel>, action: ConnectionEditActions.UpdateLocations) {
		const state = produce(ctx.getState(), st => {
			st.connectionEditForm.model!.pickupLocation = {id: action.payload.pickupID || ''};
			st.connectionEditForm.model!.sellerLocation = {id: action.payload.sellerID || ''};
			st.connectionEditForm.model!.senderLocation = {id: action.payload.senderID || ''};
			st.connectionEditForm.model!.returnLocation = {id: action.payload.returnID || ''};
		});
		ctx.setState(state);
	}

	@Action(ConnectionEditActions.AddFilterTag)
	AddFilterTag(ctx: StateContext<ConnectionEditModel>, action: ConnectionEditActions.AddFilterTag) {
		const state = produce(ctx.getState(), draft => {
			if (draft.connectionEditForm.model?.connectionShopify) {
				const filterTags = draft.connectionEditForm.model.connectionShopify.filterTags || [];
				if (!filterTags.includes(action.payload)) {
					filterTags.push(action.payload);
				}
				draft.connectionEditForm.model.connectionShopify.filterTags = filterTags;
			}
		});
		ctx.setState(state);
	}

	@Action(ConnectionEditActions.RemoveFilterTag)
	RemoveFilterTag(ctx: StateContext<ConnectionEditModel>, action: ConnectionEditActions.RemoveFilterTag) {
		const state = produce(ctx.getState(), draft => {
			if (draft.connectionEditForm.model?.connectionShopify?.filterTags) {
				draft.connectionEditForm.model.connectionShopify.filterTags =
					draft.connectionEditForm.model.connectionShopify.filterTags.filter(tag => tag !== action.payload);
			}
		});
		ctx.setState(state);
	}

}
