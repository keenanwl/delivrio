import {Injectable} from "@angular/core";
import {Action, Selector, State, StateContext} from "@ngxs/store";
import {ReturnsListActions} from "./returns-list.actions";
import {MutateReturnDeliveryOption, SearchResult} from "../../../generated/graphql";
import {
	CreateReturnOrderGQL, FetchReturnDeliveryOptionsGQL,
	FetchReturnsListGQL, MarkReturnDeletedGQL, ReturnAddDeliveryOptionGQL,
	SearchOrderLinesGQL,
	SearchOrdersGQL
} from "./returns-list.generated";
import {
	ReturnPortalViewResponse
} from "../../settings/return-portal-viewer/return-portal-frame/return-portal-frame.service";
import {
	ItemReturn
} from "../../settings/return-portal-viewer/return-portal-frame/return-portal-frame.ngxs";
import SetOrderLines = ReturnsListActions.SetOrderLines;
import {produce} from "immer";
import {AppActions} from "../../app.actions";
import ShowGlobalSnackbar = AppActions.ShowGlobalSnackbar;
import {toNotNullArray} from "../../functions/not-null-array";
import SetReturnsList = ReturnsListActions.SetReturnsList;
import FetchReturns = ReturnsListActions.FetchReturns;
import FetchReturnsList = ReturnsListActions.FetchReturnsList;
import SearchDeliveryOptionsResult = ReturnsListActions.SearchDeliveryOptionsResult;
import AppChangeRoute = AppActions.AppChangeRoute;
import {Paths} from "../../app-routing.module";

export interface ReturnsListModel {
	list: FetchReturns[];
	searchResults: SearchResult[];
	selectedOrder: SearchResult | null;
	itemsView: ReturnPortalViewResponse | undefined;
	selectedItems: ItemReturn[];
	newReturnPage: "select-order-lines" | "select-delivery-option";
	newReturnColliIDs: string[];
	newReturnColliDeliveryOptions: string[];
	newReturnColliDeliveryOptionsResults: SearchDeliveryOptionsResult;
	selectedDeliveryOptions: MutateReturnDeliveryOption[];
	loading: boolean;
}

const defaultState: ReturnsListModel = {
	list: [],
	searchResults: [],
	selectedOrder: null,
	itemsView: undefined,
	selectedItems: [],
	newReturnPage: "select-order-lines",
	newReturnColliIDs: [],
	newReturnColliDeliveryOptions: [],
	newReturnColliDeliveryOptionsResults: [],
	selectedDeliveryOptions: [],
	loading: false,
};

@Injectable()
@State<ReturnsListModel>({
	name: 'returnsList',
	defaults: defaultState,
})
export class ReturnsListState {

	constructor(
		private list: FetchReturnsListGQL,
		private search: SearchOrdersGQL,
		private searchOrderLines: SearchOrderLinesGQL,
		private create: CreateReturnOrderGQL,
		private del: MarkReturnDeletedGQL,
		private searchDelivery: FetchReturnDeliveryOptionsGQL,
		private addDeliveryOption: ReturnAddDeliveryOptionGQL,
	) {
	}

	@Selector()
	static get(state: ReturnsListModel) {
		return state;
	}

	@Action(ReturnsListActions.FetchReturnsList)
	FetchMyReturnsList(ctx: StateContext<ReturnsListModel>, action: ReturnsListActions.FetchReturnsList) {
		ctx.patchState({loading: true});
		return this.list.fetch()
			.subscribe((r) => {
				const list = toNotNullArray(r.data.returnCollis.edges?.map((n) => n?.node));
				ctx.patchState({loading: false});
				if (!!list) {
					ctx.dispatch(new SetReturnsList(list));
				}
			});
	}

	@Action(ReturnsListActions.SetReturnsList)
	SetReturnsList(ctx: StateContext<ReturnsListModel>, action: ReturnsListActions.SetReturnsList) {
		ctx.patchState({list: action.payload});
	}

	@Action(ReturnsListActions.CreateReturnOrder)
	CreateReturnOrder(ctx: StateContext<ReturnsListModel>, action: ReturnsListActions.CreateReturnOrder) {
		const orderID = ctx.getState().selectedOrder?.id || "";
		const orderLines = ctx.getState().selectedItems.map((l) => {
			return {
				orderLineID: l.orderLineID,
				claimID: l.id,
				units: l.quantity,
			}
		});
		return this.create.mutate({orderID, orderLines, portalID: ""})
			.subscribe((r) => {
				const ids = r.data?.createReturnOrder;
				if (!r.errors && !!ids) {
					ctx.dispatch([
						new ReturnsListActions.SetNewReturnColliIDs(ids),
						new ReturnsListActions.SearchReturnDeliveryOptions(),
					]);
				} else {
					ctx.dispatch(new ShowGlobalSnackbar("An error occurred"));
				}
			});
	}

	@Action(ReturnsListActions.SetSearchOrders)
	SetSearchOrders(ctx: StateContext<ReturnsListModel>, action: ReturnsListActions.SetSearchOrders) {
		ctx.patchState({searchResults: action.payload});
	}

	@Action(ReturnsListActions.SearchOrders)
	SearchOrders(ctx: StateContext<ReturnsListModel>, action: ReturnsListActions.SearchOrders) {
		return this.search.fetch({term: action.payload})
			.subscribe((s) => {
				ctx.dispatch(new ReturnsListActions.SetSearchOrders(s.data.search))
			});
	}

	@Action(ReturnsListActions.SetSelectedOrder)
	SetSelectedOrder(ctx: StateContext<ReturnsListModel>, action: ReturnsListActions.SetSelectedOrder) {
		ctx.patchState({selectedOrder: action.payload, searchResults: []});
		ctx.dispatch(new ReturnsListActions.SearchOrderLines());
	}

	@Action(ReturnsListActions.ClearCreateOrder)
	ClearCreateOrder(ctx: StateContext<ReturnsListModel>, action: ReturnsListActions.ClearCreateOrder) {
		ctx.patchState({
			selectedOrder: null,
			searchResults: [],
			selectedItems: [],
			itemsView: undefined,
			newReturnColliIDs: [],
			newReturnPage: "select-order-lines",
		});
	}

	@Action(ReturnsListActions.SetOrderLines)
	SetOrderLines(ctx: StateContext<ReturnsListModel>, action: ReturnsListActions.SetOrderLines) {
		ctx.patchState({itemsView: action.payload.view, selectedItems: action.payload.selected});
	}

	@Action(ReturnsListActions.SearchOrderLines)
	SearchOrderLines(ctx: StateContext<ReturnsListModel>, action: ReturnsListActions.SearchOrderLines) {
		const orderID = ctx.getState().selectedOrder?.id || '';
		return this.searchOrderLines.fetch({order: orderID})
			.subscribe((o) => {
				const itemsView: ReturnPortalViewResponse = {
					order_date: new Date(),
					return_reasons: [],
					packages: [],
					order_id: orderID,
				}

				const selected: ItemReturn[] = [];

				o.data.order?.colli?.forEach((c) => {
					itemsView.packages.push({
						items: c.orderLines?.map((l) => {

							selected.push({
								id: "",
								orderLineID: l.id,
								quantity: l.units,
								availableQuantity: l.units,
								selected: false,
							});

							return {
								order_line_id: l.id,
								name:          l.productVariant.product.title,
								variant_name:  l.productVariant.description || '',
								quantity:      l.units,
								image_url:     l.productVariant.productImage?.pop()?.url || '',
							}
						}) || []
					});
				});

				itemsView.return_reasons = o.data.returnClaimsByOrder || [];

				ctx.dispatch(new SetOrderLines({view: itemsView, selected}));
			});
	}

	@Action(ReturnsListActions.SetSelectedItem)
	SetSelectedItem(ctx: StateContext<ReturnsListModel>, action: ReturnsListActions.SetSelectedItem) {
		const next = ctx.getState().selectedItems?.map((i) => {
			if (i.orderLineID === action.payload.orderLineID) {
				return Object.assign({}, i, {selected: action.payload.selected});
			}
			return i;
		})
		ctx.patchState({
			selectedItems: next,
		});
	}

	@Action(ReturnsListActions.SetSelectedItemReason)
	SetSelectedItemReason(ctx: StateContext<ReturnsListModel>, action: ReturnsListActions.SetSelectedItemReason) {
		const next = ctx.getState().selectedItems?.map((i) => {
			if (i.orderLineID === action.payload.orderLineID) {
				return Object.assign({}, i, {id: action.payload.reasonID});
			}
			return i;
		})
		ctx.patchState({
			selectedItems: next,
		});
	}

	@Action(ReturnsListActions.IncrementQuantity)
	IncrementQuantity(ctx: StateContext<ReturnsListModel>, action: ReturnsListActions.IncrementQuantity) {
		const state = produce(ctx.getState(), st => {
			let next = st.selectedItems;
			for (let i = 0; i < next.length; i++) {
				const nextQuantity = next[i].quantity + 1;
				if (nextQuantity <= next[i].availableQuantity && next[i].orderLineID === action.payload.orderLineID) {
					next[i].quantity = nextQuantity;
					break;
				}
			}
			st.selectedItems = next;
		});
		ctx.setState(state);
	}

	@Action(ReturnsListActions.DecrementQuantity)
	DecrementQuantity(ctx: StateContext<ReturnsListModel>, action: ReturnsListActions.DecrementQuantity) {
		const state = produce(ctx.getState(), st => {
			let next = st.selectedItems;
			for (let i = 0; i < next.length; i++) {
				const nextQuantity = next[i].quantity - 1;
				if (nextQuantity > 0 && next[i].orderLineID === action.payload.orderLineID) {
					next[i].quantity = nextQuantity;
					break;
				}
			}
			st.selectedItems = next;
		});
		ctx.setState(state);
	}

	@Action(ReturnsListActions.MarkReturnColliDeleted)
	MarkReturnColliDeleted(ctx: StateContext<ReturnsListModel>, action: ReturnsListActions.MarkReturnColliDeleted) {
		return this.del.mutate({returnColliID: action.payload.returnColliID})
			.subscribe((r) => {
				if (!!r.errors) {
					ctx.dispatch(new ShowGlobalSnackbar("An error occurred"));
				} else {
					ctx.dispatch(new FetchReturnsList());
				}
			})
	}

	@Action(ReturnsListActions.SetNewReturnColliIDs)
	SetNewReturnColliIDs(ctx: StateContext<ReturnsListModel>, action: ReturnsListActions.SetNewReturnColliIDs) {
		ctx.patchState({newReturnPage: "select-delivery-option", newReturnColliIDs: action.payload});
	}

	@Action(ReturnsListActions.SetSearchReturnDeliveryOptions)
	SetSearchReturnDeliveryOptions(ctx: StateContext<ReturnsListModel>, action: ReturnsListActions.SetSearchReturnDeliveryOptions) {

		const state = ctx.getState();
		const selected: MutateReturnDeliveryOption[] = state.newReturnColliIDs.map((colliID, index) => {
			return {
				returnColliID: colliID,
				deliveryOptionID: action.payload[index]?.[0]?.deliveryOptionID || "",
			}
		})

		ctx.patchState({
			newReturnColliDeliveryOptionsResults: action.payload,
			selectedDeliveryOptions: selected,
		});
	}

	@Action(ReturnsListActions.SearchReturnDeliveryOptions)
	SearchReturnDeliveryOptions(ctx: StateContext<ReturnsListModel>, action: ReturnsListActions.SearchReturnDeliveryOptions) {
		return this.searchDelivery.fetch({returnColliIDs: ctx.getState().newReturnColliIDs})
			.subscribe((r) => {
				const options = r.data.returnDeliveryOptions;
				ctx.dispatch(new ReturnsListActions.SetSearchReturnDeliveryOptions(options));
			});
	}

	@Action(ReturnsListActions.ChangeDeliveryOption)
	ChangeDeliveryOption(ctx: StateContext<ReturnsListModel>, action: ReturnsListActions.ChangeDeliveryOption) {
		let nextState = ctx.getState().selectedDeliveryOptions;
		nextState = nextState.filter((o) => o.returnColliID !== action.payload.returnColliID)
		nextState.push(action.payload);
		ctx.patchState({selectedDeliveryOptions: nextState});
	}

	@Action(ReturnsListActions.CreateReturnOrderPending)
	CreateReturnOrderPending(ctx: StateContext<ReturnsListModel>, action: ReturnsListActions.CreateReturnOrderPending) {
		return this.addDeliveryOption.fetch({deliveryOptions: ctx.getState().selectedDeliveryOptions})
			.subscribe((r) => {
				if (!!r.errors) {
					ctx.dispatch(new ShowGlobalSnackbar("An error occurred"));
				} else {
					ctx.dispatch(new AppChangeRoute({path: Paths.RETURN_VIEW, queryParams: {orderID: r.data.addReturnDeliveryOption}}))
				}
			});
	}

}
