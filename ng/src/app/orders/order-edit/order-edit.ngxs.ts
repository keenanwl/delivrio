import {Injectable} from "@angular/core";
import {Action, Selector, State, StateContext} from "@ngxs/store";
import {GraphQLError} from "graphql";
import {AppActions} from "../../app.actions";
import {OrderEditActions} from "./order-edit.actions";
import {Paths} from "../../app-routing.module";
import {produce} from "immer";
import {
	CreateAddressInput,
	CreateColliInput,
	OrderStatus,
	ProductVariantQuantity,
	UpdateColliInput
} from "../../../generated/graphql";
import {
	AddColliToOrderGQL,
	FetchAvailableClickCollectLocationsGQL,
	FetchAvailableDeliveryPointsGQL,
	FetchColliGQL,
	FetchDeliveryOptionsGQL,
	OrdersSearchCountriesGQL,
	SearchProductsGQL,
	UpdateColliGQL
} from "./order-edit.generated";
import {toNotNullArray} from "../../functions/not-null-array";
import {timer} from "rxjs";
import AppChangeRoute = AppActions.AppChangeRoute;
import FetchOrderResponse = OrderEditActions.FetchOrderResponse;
import SetOrder = OrderEditActions.SetOrder;
import SearchProductsResponse = OrderEditActions.SearchProductsResponse;
import SetProducts = OrderEditActions.SetProducts;
import FetchDeliveryOptionsResponse = OrderEditActions.FetchDeliveryOptionsResponse;
import ConnectionListResponse = OrderEditActions.ConnectionListResponse;
import CountriesResponse = OrderEditActions.CountriesResponse;
import SetDeliveryPoint = OrderEditActions.SetDeliveryPoint;
import DeliveryPointResponse = OrderEditActions.DeliveryPointResponse;
import SetDeliveryPointsSearch = OrderEditActions.SetDeliveryPointsSearch;
import ClickCollectResponse = OrderEditActions.ClickCollectResponse;
import SetAvailableClickCollectLocations = OrderEditActions.SetAvailableClickCollectLocations;
import AvailableClickCollectResponse = OrderEditActions.AvailableClickCollectResponse;
import SetClickCollectLocation = OrderEditActions.SetClickCollectLocation;
import PackagingResponse = SelectPackagingActions.PackagingResponse;
import {SelectPackagingActions} from "../../shared/select-packaging/select-packaging.actions";

export interface OrderEditModel {
	orderForm: {
		model: FetchOrderResponse | undefined;
		dirty: boolean;
		status: string;
		errors: readonly GraphQLError[];
	},
	removeProductIDs: string[];
	orderID: string;
	colliID: string;
	searchProducts: SearchProductsResponse[];
	orderStatuses: OrderStatus[];
	deliveryOptions: FetchDeliveryOptionsResponse[];
	connections: ConnectionListResponse[];
	loading: boolean;
	searchCountries: CountriesResponse[];
	deliveryPoint: DeliveryPointResponse | null;
	deliveryPointsSearch: DeliveryPointResponse[];
	clickCollectLocation: ClickCollectResponse | null;
	availableClickCollectLocation: AvailableClickCollectResponse[];
	packaging: PackagingResponse[];
}

const defaultState: OrderEditModel = {
	orderForm: {
		model: undefined,
		dirty: false,
		status: '',
		errors: [],
	},
	removeProductIDs: [],
	colliID: '',
	orderID: '',
	searchProducts: [],
	orderStatuses: [
		OrderStatus.Dispatched,
		OrderStatus.Pending,
		OrderStatus.PartiallyDispatched,
		OrderStatus.Cancelled,
	],
	deliveryOptions: [],
	connections: [],
	loading: false,
	searchCountries: [],
	deliveryPoint: null,
	deliveryPointsSearch: [],
	clickCollectLocation: null,
	availableClickCollectLocation: [],
	packaging: [],
};

@Injectable()
@State<OrderEditModel>({
	name: 'orderEdit',
	defaults: defaultState,
})
export class OrderEditState {

	constructor(
		private fetchColli: FetchColliGQL,
		private searchProducts: SearchProductsGQL,
		private createOrder: AddColliToOrderGQL,
		private updateColli: UpdateColliGQL,
		private fetchDeliveryOptions: FetchDeliveryOptionsGQL,
		private countrySearch: OrdersSearchCountriesGQL,
		private dpSearch: FetchAvailableDeliveryPointsGQL,
		private availableCCLocations: FetchAvailableClickCollectLocationsGQL,
	) {}

	@Selector()
	static get(state: OrderEditModel) {
		return state;
	}

	@Action(OrderEditActions.FetchPackage)
	FetchPackage(ctx: StateContext<OrderEditModel>, action: OrderEditActions.FetchPackage) {
		const state = ctx.getState();
		return this.fetchColli.fetch({id: state.colliID}, {fetchPolicy: "no-cache", errorPolicy: "all"})
			.subscribe({next: (r) => {

				// Order is important here
				const connections = toNotNullArray(r.data?.connections?.edges?.map((c) => c?.node));
				if (!!connections) {
					ctx.dispatch(new OrderEditActions.SetConnectionList(connections));
				}

				const order = r.data?.colli;
				if (!!order) {
					ctx.dispatch(new SetOrder(order));
				}

				const deliveryPoint = r.data?.deliveryPoint;
				if (!!deliveryPoint) {
					ctx.dispatch(new SetDeliveryPoint(deliveryPoint));
				}

				const ccLocation = r.data?.clickCollectLocation;
				if (!!ccLocation) {
					ctx.dispatch(new SetClickCollectLocation(ccLocation));
				}

			}});
	}

	@Action(OrderEditActions.SetOrderID)
	SetOrderID(ctx: StateContext<OrderEditModel>, action: OrderEditActions.SetOrderID) {
		ctx.patchState({orderID: action.payload})
	}

	@Action(OrderEditActions.SetColliID)
	SetColliID(ctx: StateContext<OrderEditModel>, action: OrderEditActions.SetColliID) {
		ctx.patchState({colliID: action.payload})
	}

	@Action(OrderEditActions.SetProducts)
	SetOrderTags(ctx: StateContext<OrderEditModel>, action: OrderEditActions.SetProducts) {
		ctx.patchState({searchProducts: action.payload})
	}

	@Action(OrderEditActions.SetOrder)
	SetOrder(ctx: StateContext<OrderEditModel>, action: OrderEditActions.SetOrder) {
		const state = ctx.getState();
		const next = Object.assign({}, state.orderForm, {
			model: Object.assign({}, action.payload)
		});
		ctx.patchState({
			orderForm: next,
		})
	}

	@Action(OrderEditActions.SearchProducts)
	SearchProducts(ctx: StateContext<OrderEditModel>, action: OrderEditActions.SearchProducts) {

		ctx.patchState({loading: true});

		if (action.payload.length === 0) {
			ctx.dispatch(new SetProducts([]));
			return;
		}

		return this.searchProducts.fetch({term: action.payload}, {fetchPolicy: "no-cache"})
			.subscribe((r) => {
				const products = toNotNullArray(r.data.productVariants.edges?.map((p) => p?.node))
				ctx.dispatch(new SetProducts(products));

				ctx.dispatch(new OrderEditActions.StopLoading());
			});
	}

	@Action(OrderEditActions.AddProduct)
	AddProduct(ctx: StateContext<OrderEditModel>, action: OrderEditActions.AddProduct) {
		const val = action.payload;
		if (val) {
			const state = produce(ctx.getState(), st => {
				const next = st.orderForm.model?.orderLines || [];
				const variantWithUnits = {
					id: '', // new order line
					units: 1,
					unitPrice: 0.00,
					discountAllocationAmount: 0.00,
					currency: {display: "Shop default"},
					productVariant: val,
				}
				next?.push(variantWithUnits);
				st.orderForm.model!.orderLines = next;
			});
			ctx.setState(state);
		}
	}

	@Action(OrderEditActions.RemoveOrderLine)
	RemoveProduct(ctx: StateContext<OrderEditModel>, action: OrderEditActions.RemoveOrderLine) {
		const state = produce(ctx.getState(),
			st => {
				if (!!st.orderForm.model?.orderLines) {
					st.orderForm.model!.orderLines = st.orderForm.model?.orderLines?.filter((ol, i) => {
						const keep = i !== action.payload;
						if (!keep && ol.id.length > 0) {
							st.removeProductIDs.push(ol.id);
						}
						return keep;
					});
				}
			});
		ctx.setState(state);
	}

	selectedProductVariants(state: OrderEditModel): ProductVariantQuantity[] {
		let nextProducts: ProductVariantQuantity[] = [];
		if (!!state.orderForm.model?.orderLines) {
			nextProducts = state.orderForm.model?.orderLines?.map((l) => {
				return {
					units: l.units,
					orderLineID: l.id,
					variantID: l.productVariant.id,
					currency: l.currency.display,
					price: l.unitPrice,
					discount: l.discountAllocationAmount,
				}
			});
		}
		return nextProducts;
	}

	nextRecipientAddress(state: OrderEditModel): CreateAddressInput {
		return Object.assign({}, state.orderForm.model?.recipient,
			{id: undefined, country: undefined},
			{countryID: state.orderForm.model?.recipient.country.id}
		);
	}

	nextSenderAddress(state: OrderEditModel): CreateAddressInput {
		return Object.assign({},
			state.orderForm.model?.sender,
			{id: undefined, country: undefined},
			{countryID: state.orderForm.model?.sender.country.id}
		);
	}

	@Action(OrderEditActions.SaveFormNew)
	SaveFormNew(ctx: StateContext<OrderEditModel>) {

		const state = ctx.getState();
		//const statusID = state.orderForm.model?.orderStatus.id;
		const nextOrder: CreateColliInput = Object.assign(
			{},
			state.orderForm.model!,
			{
		//		orderStatusID: !!statusID ? statusID : '',
				orderStatus: undefined,
				recipient: undefined,
				orderSender: undefined,
				orderLines: undefined,
				deliveryOption: undefined,
			},
		);

		return this.createOrder.mutate({
			orderID: '',
			input: nextOrder,
			deliveryOptionID: state.orderForm.model?.deliveryOption?.id,
			deliveryPointID: state.deliveryPoint?.id,
			ccLocationID: state.clickCollectLocation?.id,
			products: this.selectedProductVariants(state),
			recipientAddress: this.nextRecipientAddress(state),
			senderAddress: this.nextSenderAddress(state),
		}).subscribe(() => {
			ctx.dispatch([
				new AppChangeRoute({path: Paths.ORDERS_VIEW, queryParams: {id: ctx.getState().orderID}}),
			]);
		});
	}

	@Action(OrderEditActions.SaveFormUpdate)
	SaveFormUpdate(ctx: StateContext<OrderEditModel>) {
		const state = ctx.getState();
		const nextOrder: UpdateColliInput = Object.assign(
			{},
			state.orderForm.model!,
			{
				connection: undefined,
				orderStatus: undefined,
				recipient: undefined,
				sender: undefined,
				orderLines: undefined,
				deliveryOption: undefined,
				packaging: undefined,
				order: undefined,
			},
		);

		return this.updateColli.mutate({
			id: state.colliID,
			input: nextOrder,
			deliveryOptionID: state.orderForm.model?.deliveryOption?.id,
			deliveryPointID: state.deliveryPoint?.id,
			ccLocationID: state.clickCollectLocation?.id,
			packagingID: state.orderForm.model?.packaging?.id,
			products: this.selectedProductVariants(state),
			removeProducts: state.removeProductIDs,
			recipientAddressID: state.orderForm.model?.recipient.id + '',
			recipientAddress: this.nextRecipientAddress(state),
			senderAddressID: state.orderForm.model?.sender.id + '',
			senderAddress: this.nextSenderAddress(state),
			updateExistingRecipient: false,
		}).subscribe(() => {
			ctx.dispatch([
				new AppChangeRoute({path: Paths.ORDERS_VIEW, queryParams: {id: ctx.getState().orderID}}),
			]);
		});
	}

	@Action(OrderEditActions.FetchDeliveryOptions)
	FetchDeliveryOptions(ctx: StateContext<OrderEditModel>) {
		const state = ctx.getState();

		return this.fetchDeliveryOptions.fetch({orderInfo: {
			connectionID: state.orderForm.model?.order.connection.id || '',
			zip: state.orderForm.model?.recipient.zip || "",
			country: state.orderForm.model?.recipient.country.id || "",
			productLines: state.orderForm.model?.orderLines?.map((l) => {
				return {
					productVariantID: l.productVariant.id,
					unitPrice: l.unitPrice,
					units: l.units,
				};
			})
		}}).subscribe((r) => {

			const dopt = toNotNullArray(r.data?.deliveryOptionsList);

			ctx.dispatch([
				new OrderEditActions.SetDeliveryOptions(dopt),
			]);

		});
	}

	@Action(OrderEditActions.SetDeliveryOptions)
	SetDeliveryOptions(ctx: StateContext<OrderEditModel>, action: OrderEditActions.SetDeliveryOptions) {
		ctx.patchState({deliveryOptions: action.payload})
	}

	@Action(OrderEditActions.ResetState)
	ResetState(ctx: StateContext<OrderEditModel>) {
		ctx.setState(defaultState);
	}

	@Action(OrderEditActions.SelectDeliveryOption)
	SelectDeliveryOption(ctx: StateContext<OrderEditModel>, action: OrderEditActions.SelectDeliveryOption) {
		const state = produce(ctx.getState(), st => {
			st.orderForm.model!.deliveryOption = {
				id: action.payload.id,
				clickCollect: action.payload.clickCollect,
			};
			// Reset in case we are changing CarrierBrands
			st.deliveryPoint = null;
		});
		ctx.setState(state);
	}

	@Action(OrderEditActions.SetConnectionList)
	SetConnectionList(ctx: StateContext<OrderEditModel>, action: OrderEditActions.SetConnectionList) {
		ctx.patchState({connections: action.payload});
	}

	@Action(OrderEditActions.RowEditedPrice)
	RowEditedPrice(ctx: StateContext<OrderEditModel>, action: OrderEditActions.RowEditedPrice) {
		const state = produce(ctx.getState(),
			st => {
				if (!!st.orderForm.model?.orderLines) {
					st.orderForm.model!.orderLines.some((o, i) => {
						if (i === action.payload.index && !!st.orderForm.model?.orderLines) {
							st.orderForm.model!.orderLines[i].unitPrice = action.payload.unitPrice;
							return true;
						}
						return false;
					})
				}
			});
		ctx.setState(state);
	}

	@Action(OrderEditActions.RowEditedUnits)
	RowEditedUnits(ctx: StateContext<OrderEditModel>, action: OrderEditActions.RowEditedUnits) {
		const state = produce(ctx.getState(),
			st => {
				if (!!st.orderForm.model?.orderLines) {
					st.orderForm.model!.orderLines.some((o, i) => {
						if (i === action.payload.index && !!st.orderForm.model?.orderLines) {
							st.orderForm.model!.orderLines[i].units = action.payload.units;
							return true;
						}
						return false;
					})
				}
			});
		ctx.setState(state);
	}

	@Action(OrderEditActions.StopLoading)
	StopLoading(ctx: StateContext<OrderEditModel>, action: OrderEditActions.StopLoading) {
		return timer(500).pipe((s) => {
			ctx.patchState({loading: false})
			return s;
		});
	}

	@Action(OrderEditActions.SearchCountry)
	SearchCountry(ctx: StateContext<OrderEditModel>, action: OrderEditActions.SearchCountry) {
		return this.countrySearch.fetch({term: action.payload})
			.subscribe((res) => {
				const countries = toNotNullArray(res.data.countries.edges?.map((value) => value?.node));
				if (!!countries) {
					ctx.dispatch(new OrderEditActions.SetCountrySearch(countries));
				}
			});
	}

	@Action(OrderEditActions.SetCountrySearch)
	SetCountrySearch(ctx: StateContext<OrderEditModel>, action: OrderEditActions.SetCountrySearch) {
		ctx.patchState({searchCountries: action.payload});
	}

	@Action(OrderEditActions.ChangeCountrySender)
	ChangeCountrySender(ctx: StateContext<OrderEditModel>, action: OrderEditActions.ChangeCountrySender) {
		const state = produce(ctx.getState(),
		st => {
			if (!!st.orderForm.model?.sender.country) {
				st.orderForm.model!.sender.country = action.payload;
			}
		});
		ctx.setState(state);
	}

	@Action(OrderEditActions.ChangeCountryRecipient)
	ChangeCountryRecipient(ctx: StateContext<OrderEditModel>, action: OrderEditActions.ChangeCountryRecipient) {
		const state = produce(ctx.getState(),
		st => {
			if (!!st.orderForm.model?.recipient.country) {
				st.orderForm.model!.recipient.country = action.payload;
			}
		});
		ctx.setState(state);
	}

	@Action(OrderEditActions.SetDeliveryPoint)
	SetDeliveryPoint(ctx: StateContext<OrderEditModel>, action: OrderEditActions.SetDeliveryPoint) {
		ctx.patchState({deliveryPoint: action.payload});
	}

	@Action(OrderEditActions.SetDeliveryPointsSearch)
	SetDeliveryPointsSearch(ctx: StateContext<OrderEditModel>, action: OrderEditActions.SetDeliveryPointsSearch) {
		ctx.patchState({deliveryPointsSearch: action.payload, loading: false});
	}

	@Action(OrderEditActions.SetClickCollectLocation)
	SetClickCollectLocation(ctx: StateContext<OrderEditModel>, action: OrderEditActions.SetClickCollectLocation) {
		ctx.patchState({clickCollectLocation: action.payload});
	}

	@Action(OrderEditActions.FetchAvailableClickCollectLocations)
	FetchAvailableClickCollectLocations(ctx: StateContext<OrderEditModel>, action: OrderEditActions.FetchAvailableClickCollectLocations) {
		const deliveryOptionID = ctx.getState().orderForm.model?.deliveryOption?.id;
		return this.availableCCLocations.fetch({deliveryOptionID: deliveryOptionID || ''})
			.subscribe((r) => {
				const locs = toNotNullArray(r.data.locations.edges?.map((n) => n?.node));
				if (!!locs) {
					ctx.dispatch(new SetAvailableClickCollectLocations(locs));
				}
			});
	}

	@Action(OrderEditActions.SetAvailableClickCollectLocations)
	SetAvailableClickCollectLocations(ctx: StateContext<OrderEditModel>, action: OrderEditActions.SetAvailableClickCollectLocations) {
		ctx.patchState({availableClickCollectLocation: action.payload});
	}

	@Action(OrderEditActions.SearchDeliveryPoints)
	SearchDeliveryPoints(ctx: StateContext<OrderEditModel>, action: OrderEditActions.SearchDeliveryPoints) {
		const state = ctx.getState();
		ctx.patchState({loading: true});
		return this.dpSearch.fetch({
			deliveryOptionID: state.orderForm.model?.deliveryOption?.id,
			lookupAddress: Object.assign({},
				state.orderForm.model?.recipient,
				{
					countryID: state.orderForm.model?.recipient.country.id || '',
					country: undefined,
					id: undefined,
				})
		}).subscribe((r) => {
			const deliveryPoints = r.data.availableDeliveryPoints;
			if (!!deliveryPoints) {
				ctx.dispatch(new SetDeliveryPointsSearch(deliveryPoints));
			}
		});
	}

	@Action(OrderEditActions.SetSelectedPackaging)
	SetSelectedPackaging(ctx: StateContext<OrderEditModel>, action: OrderEditActions.SetSelectedPackaging) {
		const state = produce(ctx.getState(),
			st => {
				st.orderForm.model!.packaging = action.payload;
			});
		ctx.setState(state);
	}

}
