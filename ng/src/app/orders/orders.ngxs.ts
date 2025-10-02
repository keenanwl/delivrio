import {Injectable} from "@angular/core";
import {Action, Selector, State, StateContext} from "@ngxs/store";
import {OrdersActions} from "./orders.actions";
import {OrderDirection, OrderOrderField, OrderStatus, OrderWhereInput} from "src/generated/graphql";
import {
	BulkFetchPackingSlipsByOrderGQL,
	BulkUpdatePackagingGQL,
	CreateEmptyOrderGQL,
	FetchCountriesGQL,
	FetchOrdersGQL,
	FetchOrdersQueryVariables
} from "./orders.generated";
import {toNotNullArray} from "../functions/not-null-array";
import {AppActions} from "../app.actions";
import {Paths} from "../app-routing.module";
import FetchOrders = OrdersActions.FetchOrders;
import SetWhereOptions = OrdersActions.SetWhereOptions;
import FetchOptionsCountry = OrdersActions.FetchOptionsCountry;
import ResetWhereTop = OrdersActions.ResetWhereTop;
import SetPagination = OrdersActions.SetPagination;
import ResetPagination = OrdersActions.ResetPagination;
import ConnectionsResponse = OrdersActions.ConnectionsResponse;
import AppChangeRoute = AppActions.AppChangeRoute;
import LocationResponse = OrdersActions.LocationResponse;
import ShowGlobalSnackbar = AppActions.ShowGlobalSnackbar;
import SetPackingSlips = OrdersActions.SetPackingSlips;
import {CreatePackingSlipPrintJobsGQL} from "./order-view/order-view.generated";

export interface OrderFilter {
	property: "Name";
	where: OrderWhereInput;
	relation: "and" | "or";
}

type children = "status" | "country"

export interface WhereChipOption {
	id?: string;
	name: string;
	icon: string;
	orderStatus?: OrderStatus;
}
export interface WhereChip {
	name: "top" | children;
	selectedOption: WhereChipOption | null;
	options: WhereChipOption[];
	optionsAction: any;
}

const defaultWhereSelected: WhereChip = {
	name: "top",
	selectedOption: null,
	optionsAction: null,
	options: [
		{name:"status", icon: "local_shipping"},
		{name: "country", icon: "flag"},
	]
};

export interface OrderPagination {
	totalResults: number;
	pageIndex: number;
	hasNextPage: boolean;
	hasPreviousPage: boolean;
	startCursor: string | null;
	endCursor: string | null;
}

export interface OrdersModel {
	orders: OrdersActions.FetchOrdersQueryResponse[];
	filters: OrderFilter[];
	orderRowsSelected: {[key: string]: boolean};

	availableWheres: WhereChip[];
	filteredWhere: WhereChipOption[];
	localFilteredWhere: WhereChipOption[];
	selectedWheres: Map<string, WhereChip>;
	selectedWhere: WhereChip;

	pagination: OrderPagination;
	sortDirection: OrderDirection,
	connections: ConnectionsResponse[];
	senderLocations: LocationResponse[];
	displayedColumns: string[];

	packingSlipsPDF: string[];
	allPackingSlipsPDF: string;
	labelViewerOffset: number;

	loading: boolean;
}

const availableWheres: WhereChip[] = [
	{
		name: "status",
		optionsAction: null,
		options: [
			{orderStatus: OrderStatus.Cancelled, name: "Cancelled", icon: ""},
			{orderStatus: OrderStatus.Dispatched, name: "Dispatched", icon: ""},
			{orderStatus: OrderStatus.Pending, name: "Pending", icon: ""},
			{orderStatus: OrderStatus.PartiallyDispatched, name: "Partially Dispatched", icon: ""},
		],
		selectedOption: null,
	},
	{
		name: "country",
		optionsAction: new FetchOptionsCountry(),
		options: [],
		selectedOption: null,
	},
];

const defaultState: OrdersModel = {
	orders: [],
	orderRowsSelected: {},
	filters: [],
	availableWheres: availableWheres,
	filteredWhere: defaultWhereSelected.options,
	localFilteredWhere: defaultWhereSelected.options,
	selectedWheres: new Map<string, WhereChip>(),
	selectedWhere: defaultWhereSelected,

	pagination: {
		totalResults: 0,
		pageIndex: 0,
		hasNextPage: true,
		hasPreviousPage: false,
		startCursor: null,
		endCursor: null,
	},
	connections: [],
	senderLocations: [],
	sortDirection: OrderDirection.Desc,
	displayedColumns: [
		'orderNumber',
		'creation',
		'connection',
		'recipient',
		'status',
		'shipments',
		'total',
		"country",
		"deliveryOption",
		'select',
	],
	packingSlipsPDF: [],
	allPackingSlipsPDF: "",
	labelViewerOffset: 0,
	loading: false,
};

@Injectable()
@State<OrdersModel>({
	name: 'orders',
	defaults: defaultState,
})
export class OrdersState {

	constructor(
		private fetchOrders: FetchOrdersGQL,
		private fetchCountries: FetchCountriesGQL,
		private createOrder: CreateEmptyOrderGQL,
		private bulkPackaging: BulkUpdatePackagingGQL,
		private fetchPackingSlips: BulkFetchPackingSlipsByOrderGQL,
		private createPackingSlipPrintJobs: CreatePackingSlipPrintJobsGQL,
	) {
	}

	@Selector()
	static state(state: OrdersModel) {
		return state;
	}

	@Action(OrdersActions.ResetWhereTop)
	ResetWhereTop(ctx: StateContext<OrdersModel>, action: OrdersActions.ResetWhereTop) {
		const state = ctx.getState();
		const nextFilteredOptions = defaultWhereSelected.options.filter((v) => !state.selectedWheres.has(v.name));
		ctx.patchState({
			filteredWhere: nextFilteredOptions,
			localFilteredWhere: nextFilteredOptions,
			selectedWhere: defaultWhereSelected,
		})
	}

	@Action(OrdersActions.SelectWhereTop)
	SelectWhereTop(ctx: StateContext<OrdersModel>, action: OrdersActions.SelectWhereTop) {
		const state = ctx.getState();
		let nextWhere = state.selectedWhere;
		let nextSelectedWheres = new Map(state.selectedWheres);
		state.availableWheres.forEach((where) => {
			if (where.name === action.payload.name) {
				nextSelectedWheres.set(where.name, where)
				nextWhere = where;
			}
		});

		ctx.patchState({
			selectedWhere: nextWhere,
			selectedWheres: nextSelectedWheres,
			filteredWhere: [],
			localFilteredWhere: [],
		});

		ctx.dispatch(new SetWhereOptions(nextWhere.options));

		if (!!nextWhere.optionsAction) {
			ctx.dispatch(nextWhere.optionsAction);
		}
	}

	@Action(OrdersActions.SelectWhere)
	SelectWhere(ctx: StateContext<OrdersModel>, action: OrdersActions.SelectWhere) {
		const state = ctx.getState();
		let nextSelectedWheres = new Map(state.selectedWheres);
		state.filteredWhere.forEach((o) => {
			if (o.name === action.payload.name) {
				let setOption = Object.assign({}, nextSelectedWheres.get(state.selectedWhere.name));
				if (!!setOption) {
					setOption.selectedOption = o;
					nextSelectedWheres.set(state.selectedWhere.name, setOption);
				}
			}
		});

		ctx.patchState({
			selectedWheres: nextSelectedWheres,
		});

		ctx.dispatch([new ResetWhereTop(), new ResetPagination(), new FetchOrders()]);

	}

	@Action(OrdersActions.WhereChipClicked)
	WhereChipClicked(ctx: StateContext<OrdersModel>, action: OrdersActions.WhereChipClicked) {
		const state = ctx.getState();
		let nextWhere = state.selectedWhere;
		let nextFilteredWheres = state.filteredWhere;
		state.selectedWheres.forEach((where) => {
			if (where.name === action.payload) {
				nextWhere = where;
				nextFilteredWheres = nextWhere.options.filter((v) => !state.selectedWheres.has(v.name));
			}
		});

		ctx.patchState({
			selectedWhere: nextWhere,
			filteredWhere: nextFilteredWheres,
			localFilteredWhere: nextFilteredWheres,
		});
		if (!!nextWhere.optionsAction) {
			ctx.dispatch(nextWhere.optionsAction);
		}
	}

	@Action(OrdersActions.WhereChipRemove)
	WhereChipRemove(ctx: StateContext<OrdersModel>, action: OrdersActions.WhereChipRemove) {
		const next = new Map(ctx.getState().selectedWheres);
		next.delete(action.payload);
		ctx.patchState({selectedWheres: next});
		ctx.dispatch([new ResetWhereTop(), new ResetPagination(), new FetchOrders()]);
	}

	@Action(OrdersActions.WhereChipRemoveAll)
	WhereChipRemoveAll(ctx: StateContext<OrdersModel>, action: OrdersActions.WhereChipRemoveAll) {
		ctx.patchState({selectedWheres: new Map<string, WhereChip>()});
		ctx.dispatch([new ResetWhereTop(), new ResetPagination(), new FetchOrders()]);
	}

	@Action(OrdersActions.FetchOrders)
	FetchOrders(ctx: StateContext<OrdersModel>, action: OrdersActions.FetchOrders) {

		ctx.patchState({loading: true});
		const state = ctx.getState()
		let andish: OrderWhereInput = {and: []};
		state.selectedWheres.forEach((value, key) => {
			if (value.name === "status") {
				andish.and?.push({status: value.selectedOption?.orderStatus});
			} else if (value.name === "country") {
				andish.and?.push({
					hasColliWith: [
						{
							hasRecipientWith: [
								{hasCountryWith: [{id: `${value.selectedOption?.id}`}]}
							]
						}
					]
				});
			}
		});

		// Should correspond to paginator value
		const limit = 15;
		let variables: FetchOrdersQueryVariables = {
			where: andish,
			before: state.pagination.startCursor,
			last: limit,
			orderBy: {
				direction: state.sortDirection,
				field: OrderOrderField.CreatedAt,
			}
		}
		if (action.payload === "next") {
			variables = {
				where: andish,
				after: state.pagination.endCursor,
				first: limit,
				orderBy: {
					direction: state.sortDirection,
					field: OrderOrderField.CreatedAt,
				}
			}
		}

		return this.fetchOrders.fetch(variables)
			.subscribe((o) => {
				const connections = toNotNullArray(o.data?.connections.edges?.map((o) => o?.node));
				ctx.dispatch(new OrdersActions.SetConnections(connections));

				const senderLocations = toNotNullArray(o.data?.locations.edges?.map((o) => o?.node));
				ctx.dispatch(new OrdersActions.SetSenderLocations(senderLocations));

				const orders = toNotNullArray(o.data?.orders.edges?.map((o) => o?.node));
				ctx.dispatch(new OrdersActions.SetOrders(orders));

				const info = o.data?.orders;
				ctx.dispatch(new SetPagination({
					endCursor: info.pageInfo.endCursor as string,
					pageIndex: 0,
					startCursor: info.pageInfo.startCursor as string,
					totalResults: info.totalCount,
					hasNextPage: info.pageInfo.hasNextPage,
					hasPreviousPage: info.pageInfo.hasPreviousPage,
				}));
				ctx.patchState({loading: false});

			});
	}

	@Action(OrdersActions.FetchOptionsCountry)
	FetchOptionsCountry(ctx: StateContext<OrdersModel>, action: OrdersActions.FetchOptionsCountry) {
		return this.fetchCountries.fetch()
			.subscribe((r) => {
				const orderStatusLabels: WhereChipOption[] = [];
				r.data.countries?.edges?.forEach((c) => {
					orderStatusLabels.push({id: c?.node?.id, name: `${c?.node?.label}`, icon: ""})
				})
				ctx.dispatch(new SetWhereOptions(orderStatusLabels))
			});
	}

	@Action(OrdersActions.SetWhereOptions)
	SetWhereOptions(ctx: StateContext<OrdersModel>, action: OrdersActions.SetWhereOptions) {
		const state = ctx.getState();
		const nextFilteredOptions = action.payload.filter((v) => !state.selectedWheres.has(v.name));
		ctx.patchState({
			filteredWhere: nextFilteredOptions,
			localFilteredWhere: nextFilteredOptions,
		});
	}

	@Action(OrdersActions.SetOrders)
	SetOrders(ctx: StateContext<OrdersModel>, action: OrdersActions.SetOrders) {
		ctx.patchState({orders: action.payload})
	}

	@Action(OrdersActions.SetConnections)
	SetConnections(ctx: StateContext<OrdersModel>, action: OrdersActions.SetConnections) {
		ctx.patchState({connections: action.payload})
	}

	@Action(OrdersActions.AddOrderFilter)
	AddOrderFilter(ctx: StateContext<OrdersModel>, action: OrdersActions.AddOrderFilter) {
		const next = [...ctx.getState().filters];
		next.push(action.payload);
		ctx.patchState({filters: next});
	}

	@Action(OrdersActions.OrderRowsToggleRows)
	OrderRowsToggleRows(ctx: StateContext<OrdersModel>, action: OrdersActions.OrderRowsToggleRows) {
		const state = ctx.getState();
		const next: {[key: string]: boolean} = Object.assign({}, state.orderRowsSelected);
		action.payload.forEach((o) => {
			if (next[o.id] === true) {
				delete next[o.id]
			} else {
				next[o.id] = true
			}
		})

		ctx.patchState({
			orderRowsSelected: next,
		});
	}

	@Action(OrdersActions.OrderRowsToggleAll)
	OrderRowsToggleAll(ctx: StateContext<OrdersModel>, action: OrdersActions.OrderRowsToggleAll) {
		const state = ctx.getState();
		const next: {[key: string]: boolean} = {};
		if (state.orders.length !== Object.keys(state.orderRowsSelected).length) {
			state.orders.forEach((o) => {
				next[o.id] = true
			})
		}

		ctx.patchState({orderRowsSelected: next});
	}

	@Action(OrdersActions.PreviousPage)
	PreviousPage(ctx: StateContext<OrdersModel>, action: OrdersActions.PreviousPage) {
		ctx.dispatch(new FetchOrders("previous"));
	}

	@Action(OrdersActions.NextPage)
	NextPage(ctx: StateContext<OrdersModel>, action: OrdersActions.NextPage) {
		ctx.dispatch(new FetchOrders());
	}

	@Action(OrdersActions.SetPagination)
	SetPagination(ctx: StateContext<OrdersModel>, action: OrdersActions.SetPagination) {
		ctx.patchState({
			orderRowsSelected: {},
			pagination: action.payload});
	}

	@Action(OrdersActions.ResetPagination)
	ResetPagination(ctx: StateContext<OrdersModel>, action: OrdersActions.ResetPagination) {
		ctx.patchState({pagination: defaultState.pagination});
	}

	@Action(OrdersActions.ResetState)
	ResetState(ctx: StateContext<OrdersModel>) {
		ctx.setState(defaultState);
	}

	@Action(OrdersActions.CreateNewOrder)
	CreateNewOrder(ctx: StateContext<OrdersModel>, action: OrdersActions.CreateNewOrder) {
		return this.createOrder.mutate({input: action.payload.input}, {errorPolicy: "all"})
			.subscribe((r) => {
				if (!!r.errors) {
					ctx.dispatch(new ShowGlobalSnackbar("Error: " + JSON.stringify(r.errors)));
				} else {
					ctx.dispatch([
						new AppChangeRoute({path: Paths.ORDERS_VIEW, queryParams: {id: r.data?.createEmptyOrder?.id}}),
					])
				}
			});
	}

	@Action(OrdersActions.ShowHideColumn)
	ShowHideColumn(ctx: StateContext<OrdersModel>, action: OrdersActions.ShowHideColumn) {
		ctx.patchState({displayedColumns: action.payload.sort(
			(a, b) => defaultState.displayedColumns.indexOf(a) - defaultState.displayedColumns.indexOf(b))
		});
	}

	@Action(OrdersActions.ChangeSortBy)
	ChangeSortBy(ctx: StateContext<OrdersModel>, action: OrdersActions.ChangeSortBy) {
		ctx.patchState({sortDirection: action.payload});
		ctx.dispatch([
			new OrdersActions.ResetPagination(),
			new FetchOrders(),
		]);
	}

	@Action(OrdersActions.BulkUpdatePackaging)
	BulkUpdatePackaging(ctx: StateContext<OrdersModel>, action: OrdersActions.BulkUpdatePackaging) {
		const state = ctx.getState();
		return this.bulkPackaging.mutate({orderIDs: Object.keys(state.orderRowsSelected), packagingID: action.payload}, {errorPolicy: "all"})
			.subscribe((r) => {
				if (!!r.errors) {
					ctx.dispatch(new ShowGlobalSnackbar("Error: " + JSON.stringify(r.errors)));
				} else {
					ctx.dispatch(new ShowGlobalSnackbar(r.data?.bulkUpdatePackaging.msg || ""));
				}
			});
	}

	@Action(OrdersActions.BulkFetchPackingSlips)
	BulkFetchPackingSlips(ctx: StateContext<OrdersModel>, action: OrdersActions.BulkFetchPackingSlips) {
		ctx.patchState({loading: true});
		const state = ctx.getState();
		return this.fetchPackingSlips.fetch({orderIDs: Object.keys(state.orderRowsSelected)}, {errorPolicy: "all"})
			.subscribe((r) => {
				ctx.patchState({loading: false});
				if (!!r.errors) {
					ctx.dispatch(new ShowGlobalSnackbar("Error: " + JSON.stringify(r.errors)));
				} else {
					ctx.dispatch(new SetPackingSlips(r.data.packingSlipsByOrder));
				}
			});
	}

	@Action(OrdersActions.CreatePackingSlipPrintJobs)
	CreatePackingSlipPrintJobs(ctx: StateContext<OrdersModel>, action: OrdersActions.CreatePackingSlipPrintJobs) {
		ctx.patchState({loading: true});
		const state = ctx.getState();
		const colliIDs: string[] = []
		state.orders.forEach((order) => {
			order.colli?.forEach((colli) => {
				if (state.orderRowsSelected[order.id]) {
					colliIDs.push(colli.id);
				}
			});
		})

		return this.createPackingSlipPrintJobs.fetch({colliIDs: colliIDs}, {errorPolicy: "all"})
			.subscribe((r) => {
				ctx.patchState({loading: false});
				if (!!r.errors) {
					ctx.dispatch(new ShowGlobalSnackbar("Error: " + JSON.stringify(r.errors)));
				} else {
					ctx.dispatch(new ShowGlobalSnackbar("Print jobs created"));
				}
			});
	}

	@Action(OrdersActions.SetPackingSlips)
	SetPackingSlips(ctx: StateContext<OrdersModel>, action: OrdersActions.SetPackingSlips) {
		ctx.patchState({
			packingSlipsPDF: action.payload.packingSlips,
			allPackingSlipsPDF: action.payload.allPackingSlips,
		});
	}

	@Action(OrdersActions.LocalFilterWhere)
	LocalFilterWhere(ctx: StateContext<OrdersModel>, action: OrdersActions.LocalFilterWhere) {
		const state = ctx.getState();
		ctx.patchState({
			localFilteredWhere: state.filteredWhere.filter((where) => {
				if (where.name.toLowerCase().includes(action.payload)) {
					return true
				}
				return false;
			})
		});
	}
}
