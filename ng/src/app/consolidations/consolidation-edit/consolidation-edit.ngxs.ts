import {Injectable} from "@angular/core";
import {Action, Selector, State, StateContext} from "@ngxs/store";
import {produce} from "immer";
import {
	ConsolidationSearchCountriesGQL,
	ConsolidationSearchOrdersGQL, CreateConsolidationShipmentGQL,
	FetchConsolidationGQL,
	UpdateConsolidationGQL
} from "./consolidation-edit.generated";
import {ConsolidationEditActions} from "./consolidation-edit.actions";
import ConsolidationResponse = ConsolidationEditActions.ConsolidationResponse;
import {AppActions} from "../../app.actions";
import ShowGlobalSnackbar = AppActions.ShowGlobalSnackbar;
import AppChangeRoute = AppActions.AppChangeRoute;
import {Paths} from "../../app-routing.module";
import {GraphQLError} from "graphql/index";
import PalletResponse = ConsolidationEditActions.PalletResponse;
import OrderResponse = ConsolidationEditActions.OrderResponse;
import SetPallets = ConsolidationEditActions.SetPallets;
import SetOrders = ConsolidationEditActions.SetOrders;
import SetSearchOrders = ConsolidationEditActions.SetSearchOrders;
import {toNotNullArray} from "../../functions/not-null-array";
import {CreateAddressInput, CreateOrUpdatePallet, UpdateConsolidationInput} from "../../../generated/graphql";
import DeliveryOptionItem = ConsolidationEditActions.DeliveryOptionItem;
import SetDeliveryOptions = ConsolidationEditActions.SetDeliveryOptions;
import PackagingItem = ConsolidationEditActions.PackagingItem;
import SetPackagings = ConsolidationEditActions.SetPackagings;
import {AddressInfoFragment} from "../../orders/order-edit/order-edit.generated";
import SetRecipient = ConsolidationEditActions.SetRecipient;
import SetSender = ConsolidationEditActions.SetSender;
import CountryResult = ConsolidationEditActions.CountryResult;
import SetSearchCountries = ConsolidationEditActions.SetSearchCountries;
import SetLabels = ConsolidationEditActions.SetLabels;
import SetShipmentErrors = ConsolidationEditActions.SetShipmentErrors;
import ConsolidationShipment = ConsolidationEditActions.ConsolidationShipment;
import SetShipmentInfo = ConsolidationEditActions.SetShipmentInfo;
import FetchConsolidationEdit = ConsolidationEditActions.FetchConsolidationEdit;

export interface ConsolidationEditModel {
	id: string;
	form: {
		model: ConsolidationResponse | undefined;
		dirty: boolean;
		status: string;
		errors: readonly GraphQLError[];
	}
	loading: boolean;
	pallets: PalletResponse[];
	orders: OrderResponse[];
	editPallet: PalletResponse | undefined;
	orderSearch: OrderResponse[];
	deliveryOptions: DeliveryOptionItem[];
	packagings: PackagingItem[];
	recipient: AddressInfoFragment | null;
	sender: AddressInfoFragment | null;
	searchCountries: CountryResult[];
	shipmentInfo: ConsolidationShipment;

	shipmentErrors: readonly GraphQLError[];
	labelsPDF: string[];
	allLabelsPDF: string;
	labelViewerOffset: number;
}

const defaultState: ConsolidationEditModel = {
	id: "",
	form: {
		model: undefined,
		dirty: false,
		status: '',
		errors: [],
	},
	pallets: [],
	orders: [],
	editPallet: undefined,
	orderSearch: [],
	deliveryOptions: [],
	packagings: [],
	loading: true,
	sender: null,
	recipient: null,
	searchCountries: [],
	shipmentInfo: {
		mayBook: false,
		mayPrebook: false,
	},

	shipmentErrors: [],
	labelsPDF: [],
	allLabelsPDF: '',
	labelViewerOffset: 0,
};

@Injectable()
@State<ConsolidationEditModel>({
	name: 'consolidationEdit',
	defaults: defaultState,
})
export class ConsolidationEditState {

	constructor(
		private fetch: FetchConsolidationGQL,
		private update: UpdateConsolidationGQL,
		private search: ConsolidationSearchOrdersGQL,
		private searchCountry: ConsolidationSearchCountriesGQL,
		private ship: CreateConsolidationShipmentGQL,
	) {
	}

	@Selector()
	static get(state: ConsolidationEditModel) {
		return state;
	}

	@Action(ConsolidationEditActions.FetchConsolidationEdit)
	FetchConsolidationEdit(ctx: StateContext<ConsolidationEditModel>, action: ConsolidationEditActions.FetchConsolidationEdit) {
		ctx.patchState({loading: true});
		return this.fetch.fetch({id: ctx.getState().id})
			.subscribe({next: (r) => {
				ctx.patchState({loading: false});

				const deliveryOptions = toNotNullArray(r.data.deliveryOptions.edges?.map((n) => n?.node));
				ctx.dispatch(new SetDeliveryOptions(deliveryOptions));

				const packagings = toNotNullArray(r.data.packagings.edges?.map((n) => n?.node));
				ctx.dispatch(new SetPackagings(packagings));

				const recipient = r.data.locationInfo.recipient;
				if (!!recipient) {
					ctx.dispatch(new SetRecipient(recipient));
				}

				const sender = r.data.locationInfo.sender;
				if (!!sender) {
					ctx.dispatch(new SetSender(sender));
				}

				const shipment = r.data.consolidationShipments;
				if (!!shipment) {
					ctx.dispatch(new SetShipmentInfo(shipment));
				}

				const pallets = r.data.consolidation.pallets;
				if (!!pallets) {
					ctx.dispatch([new SetPallets(pallets)]);
				}

				const orders = r.data.consolidation.orders;
				if (!!orders) {
					ctx.dispatch([new SetOrders(orders)]);
				}

				const consolidation = r.data.consolidation;
				if (!!consolidation) {
					ctx.dispatch(new ConsolidationEditActions.SetConsolidationEdit(consolidation));
				}
			}});
	}

	@Action(ConsolidationEditActions.SetConsolidationEdit)
	SetConsolidationEdit(ctx: StateContext<ConsolidationEditModel>, action: ConsolidationEditActions.SetConsolidationEdit) {
		const state = produce(ctx.getState(), st => {
			st.form.model = action.payload;
		});
		ctx.setState(state);
	}

	@Action(ConsolidationEditActions.SetConsolidationID)
	SetConsolidationID(ctx: StateContext<ConsolidationEditModel>, action: ConsolidationEditActions.SetConsolidationID) {
		ctx.patchState({id: action.payload})
	}

	@Action(ConsolidationEditActions.SetPallets)
	SetPallets(ctx: StateContext<ConsolidationEditModel>, action: ConsolidationEditActions.SetPallets) {
		ctx.patchState({pallets: action.payload});
	}

	@Action(ConsolidationEditActions.SetOrders)
	SetOrders(ctx: StateContext<ConsolidationEditModel>, action: ConsolidationEditActions.SetOrders) {
		ctx.patchState({orders: action.payload});
	}

	@Action(ConsolidationEditActions.SetShipmentInfo)
	SetShipmentInfo(ctx: StateContext<ConsolidationEditModel>, action: ConsolidationEditActions.SetShipmentInfo) {
		ctx.patchState({shipmentInfo: action.payload});
	}

	@Action(ConsolidationEditActions.SearchOrders)
	SearchOrders(ctx: StateContext<ConsolidationEditModel>, action: ConsolidationEditActions.SearchOrders) {
		return this.search.fetch({term: action.payload})
			.subscribe((r) => {
				ctx.dispatch(new SetSearchOrders(toNotNullArray(r.data.orders.edges?.map((n) => n?.node))));
			});
	}

	@Action(ConsolidationEditActions.SetSearchOrders)
	SetSearchOrders(ctx: StateContext<ConsolidationEditModel>, action: ConsolidationEditActions.SetSearchOrders) {
		ctx.patchState({orderSearch: action.payload});
	}

	@Action(ConsolidationEditActions.SetDeliveryOptions)
	SetDeliveryOptions(ctx: StateContext<ConsolidationEditModel>, action: ConsolidationEditActions.SetDeliveryOptions) {
		ctx.patchState({deliveryOptions: action.payload});
	}

	@Action(ConsolidationEditActions.EditPallet)
	EditPalletID(ctx: StateContext<ConsolidationEditModel>, action: ConsolidationEditActions.EditPallet) {
		ctx.patchState({editPallet: action.payload});
	}

	@Action(ConsolidationEditActions.SetPackagings)
	SetPackagings(ctx: StateContext<ConsolidationEditModel>, action: ConsolidationEditActions.SetPackagings) {
		ctx.patchState({packagings: action.payload});
	}

	@Action(ConsolidationEditActions.SetRecipient)
	SetRecipient(ctx: StateContext<ConsolidationEditModel>, action: ConsolidationEditActions.SetRecipient) {
		ctx.patchState({recipient: action.payload});
	}

	@Action(ConsolidationEditActions.SetSender)
	SetSender(ctx: StateContext<ConsolidationEditModel>, action: ConsolidationEditActions.SetSender) {
		ctx.patchState({sender: action.payload});
	}

	@Action(ConsolidationEditActions.SetSearchCountries)
	SetSearchCountries(ctx: StateContext<ConsolidationEditModel>, action: ConsolidationEditActions.SetSearchCountries) {
		ctx.patchState({searchCountries: action.payload});
	}

	@Action(ConsolidationEditActions.UpdatePallet)
	UpdatePallet(ctx: StateContext<ConsolidationEditModel>, action: ConsolidationEditActions.UpdatePallet) {
		const state = produce(ctx.getState(), st => {
			if (!!st.editPallet) {
				st.pallets.forEach((p, i) => {
					if (p.id === st.editPallet?.id) {
						st.pallets[i].publicID = action.payload.publicID;
						st.pallets[i].description = action.payload.description;
						st.pallets[i].packaging = {id: action.payload.packagingID};
					}
				});
			} else {
				st.pallets.push({
					id: "",
					publicID: action.payload.publicID,
					description: action.payload.description,
					packaging: {id: action.payload.packagingID},
					orders: [],
				})
			}
		})
		ctx.setState(state);
	}

	@Action(ConsolidationEditActions.MoveOrder)
	MoveOrder(ctx: StateContext<ConsolidationEditModel>, action: ConsolidationEditActions.MoveOrder) {
		ctx.setState(this.removeOrder(ctx, action.payload.item.id));
		const state = produce(ctx.getState(), st => {
				switch (action.payload.destination.type) {
					case ConsolidationEditActions.ListContainerType.ORDERS:
						st.orders.push(action.payload.item);
						break;
					case ConsolidationEditActions.ListContainerType.SEARCH:
						st.orderSearch.push(action.payload.item);
						break;
					case ConsolidationEditActions.ListContainerType.PALLET:
						st.pallets[action.payload.destination.index].orders!.push(action.payload.item);
						break;
				}
		})
		ctx.setState(state);
	}

	@Action(ConsolidationEditActions.RemoveOrder)
	RemoveOrder(ctx: StateContext<ConsolidationEditModel>, action: ConsolidationEditActions.RemoveOrder) {
		ctx.setState(this.removeOrder(ctx, action.payload));
	}

	removeOrder(ctx: StateContext<ConsolidationEditModel>, id: string) {
		return produce(ctx.getState(), st => {
			st.orders = st.orders.filter((o) => o.id !== id);
			st.orderSearch = st.orderSearch.filter((o) => o.id !== id);
			st.pallets.forEach((p, i) => {
				st.pallets[i].orders = st.pallets[i].orders!
					.filter((o) => o.id !== id);
			});
		});
	}

	@Action(ConsolidationEditActions.Clear)
	Clear(ctx: StateContext<ConsolidationEditModel>, action: ConsolidationEditActions.Clear) {
		ctx.setState(defaultState);
	}

	@Action(ConsolidationEditActions.SearchCountries)
	SearchCountries(ctx: StateContext<ConsolidationEditModel>, action: ConsolidationEditActions.SearchCountries) {
		return this.searchCountry.fetch({term: action.payload})
			.subscribe((resp) => {
				const c = toNotNullArray(resp.data.countries.edges?.map((n) => n?.node));
				ctx.dispatch(new SetSearchCountries(c));
			});
	}

	@Action(ConsolidationEditActions.CreateShipment)
	CreateShipment(ctx: StateContext<ConsolidationEditModel>, action: ConsolidationEditActions.CreateShipment) {
		ctx.patchState({loading: true});
		const state = ctx.getState();
		return this.ship.fetch({consolidationID: state.id, prebook: action.payload.prebook}, {errorPolicy: "all"})
			.subscribe((resp) => {
				if (!!resp.errors) {
					ctx.dispatch(new SetShipmentErrors({errors: resp.errors}));
				} else {
					ctx.dispatch([
						new SetLabels(resp.data.createConsolidationShipment),
						new FetchConsolidationEdit(),
					]);
				}
				ctx.patchState({loading: false});
			});
	}

	@Action(ConsolidationEditActions.SetShipmentErrors)
	SetShipmentErrors(ctx: StateContext<ConsolidationEditModel>, action: ConsolidationEditActions.SetShipmentErrors) {
		ctx.patchState({shipmentErrors: action.payload.errors, loading: false});
	}

	@Action(ConsolidationEditActions.SetLabels)
	SetLabels(ctx: StateContext<ConsolidationEditModel>, action: ConsolidationEditActions.SetLabels) {
		ctx.patchState({
			labelsPDF: action.payload.labelsPDF,
			allLabelsPDF: action.payload.allLabels,
			labelViewerOffset: 0,
			loading: false,
		});
	}

	@Action(ConsolidationEditActions.IncrementLabelViewerOffset)
	IncrementLabelViewerOffset(ctx: StateContext<ConsolidationEditModel>, action: ConsolidationEditActions.IncrementLabelViewerOffset) {
		const current = ctx.getState().labelViewerOffset;
		if (current < ctx.getState().labelsPDF.length-1) {
			ctx.patchState({labelViewerOffset: current+1});
		}
	}

	@Action(ConsolidationEditActions.DecrementLabelViewerOffset)
	DecrementLabelViewerOffset(ctx: StateContext<ConsolidationEditModel>, action: ConsolidationEditActions.DecrementLabelViewerOffset) {
		const current = ctx.getState().labelViewerOffset;
		if (current > 0) {
			ctx.patchState({labelViewerOffset: current-1});
		}
	}

	@Action(ConsolidationEditActions.Save)
	Save(ctx: StateContext<ConsolidationEditModel>, action: ConsolidationEditActions.Save) {

		const state = ctx.getState();
		const input: UpdateConsolidationInput = {
			publicID: state.form?.model?.publicID,
			description: state.form?.model?.description,
			deliveryOptionID: state.form.model?.deliveryOption?.id,
			clearOrders: true,
			addOrderIDs: state.orders.map((o) => o.id),
		}

		const inputPallets: CreateOrUpdatePallet[] = state.pallets.map((p) => {
			return {
				id: p.id,
				create: {
					consolidationID: state.id,
					orderIDs: p.orders?.map((o) => o.id),
					description: p.description,
					publicID: p.publicID,
					packagingID: p.packaging?.id,
				}
			}
		});

		let sender: CreateAddressInput | undefined = undefined;
		if (!!state.sender) {
			sender = Object.assign({}, state.sender, {
				countryID: state.sender?.country.id,
				id: undefined,
				country: undefined});
		}

		let recipient: CreateAddressInput | undefined = undefined;
		if (!!state.recipient) {
			recipient = Object.assign({}, state.recipient, {
				countryID: state.recipient?.country.id,
				id: undefined,
				country: undefined});
		}

		return this.update.mutate(
			{
				id: state.id,
				input: input,
				sender: sender,
				recipient: recipient,
				inputPallets: inputPallets,
			},
			{errorPolicy: "all"},
		).subscribe((r) => {
			if (!!r.errors) {
				ctx.dispatch(new ShowGlobalSnackbar(`An error occurred: ${r.errors.map((e) => e.message).join(" ")}`));
			} else {
				ctx.dispatch(new AppChangeRoute({path: Paths.CONSOLIDATIONS, queryParams: {}}));
			}
		});
	}

}
