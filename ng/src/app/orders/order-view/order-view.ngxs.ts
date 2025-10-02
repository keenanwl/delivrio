import {Injectable} from "@angular/core";
import {Action, Selector, State, StateContext} from "@ngxs/store";
import {OrderViewActions} from "./order-view.actions";
import FetchOrderResponse = OrderViewActions.FetchOrderResponse;
import SetOrder = OrderViewActions.SetOrder;
import {
	CancelShipmentByColliIDsGQL,
	CreateLabelPrintJobsGQL, CreatePackingSlipPrintJobsGQL,
	CreateShipmentsGQL,
	DeletePackageGQL,
	DuplicateColliGQL, FetchLabelsGQL,
	FetchOrderViewGQL, FetchPackingSlipsGQL,
	MoveOrderLineGQL, PackingSlipsClearCacheGQL,
	UpdateOrderGQL
} from "./order-view.generated";
import TimelineResponse = OrderViewActions.TimelineResponse;
import {produce} from "immer";
import {toNotNullArray} from "../../functions/not-null-array";
import SetConnections = OrderViewActions.SetConnections;
import ConnectionResponse = OrderViewActions.ConnectionResponse;
import {AppActions} from "../../app.actions";
import ShowGlobalSnackbar = AppActions.ShowGlobalSnackbar;
import {GraphQLError} from "graphql";
import SetShipmentErrors = OrderViewActions.SetShipmentErrors;
import SetLabels = OrderViewActions.SetLabels;
import ShipmentStatusesResponse = OrderViewActions.ShipmentStatusesResponse;
import SetShipmentStatuses = OrderViewActions.SetShipmentStatuses;
import FetchOrder = OrderViewActions.FetchOrder;
import SetPackingSlips = OrderViewActions.SetPackingSlips;
import FetchSelectedWorkstation = AppActions.FetchSelectedWorkstation;

export interface OrderViewModel {
	order: FetchOrderResponse | undefined;
	orderID: string;
	selectedColliIDs: string[];
	timeline: TimelineResponse[];
	isDragging: boolean;
	connections: ConnectionResponse[];
	loading: boolean;
	shipmentErrors: readonly GraphQLError[];
	labelsPDF: string[];
	allLabelsPDF: string;
	packingSlipsPDF: string[];
	allPackingSlipsPDF: string;
	labelViewerOffset: number;
	shipmentStatuses: ShipmentStatusesResponse;
}

const defaultState: OrderViewModel = {
	order: undefined,
	orderID: '',
	selectedColliIDs: [],
	timeline: [],
	isDragging: false,
	connections: [],
	loading: false,
	shipmentErrors: [],
	labelsPDF: [],
	allLabelsPDF: '',
	packingSlipsPDF: [],
	allPackingSlipsPDF: '',
	labelViewerOffset: 0,
	shipmentStatuses: {
		shipmentStatuses: [],
		mayShipRemaining: false,
	}
};

@Injectable()
@State<OrderViewModel>({
	name: 'orderView',
	defaults: defaultState,
})
export class OrderViewState {

	constructor(
		private fetchOrder: FetchOrderViewGQL,
		private duplicateColli: DuplicateColliGQL,
		private move: MoveOrderLineGQL,
		private deletePackage: DeletePackageGQL,
		private updateOrder: UpdateOrderGQL,
		private createShipments: CreateShipmentsGQL,
		private fetchLabels: FetchLabelsGQL,
		private fetchPackingSlips: FetchPackingSlipsGQL,
		private cancelShipment: CancelShipmentByColliIDsGQL,
		private createPackingSlipPrintJobs: CreatePackingSlipPrintJobsGQL,
		private createLabelsPrintJobs: CreateLabelPrintJobsGQL,
		private packingSlipClearCache: PackingSlipsClearCacheGQL,
	) {}

	@Selector()
	static get(state: OrderViewModel) {
		return state;
	}

	@Action(OrderViewActions.FetchOrder)
	FetchMyOrder(ctx: StateContext<OrderViewModel>, action: OrderViewActions.FetchOrder) {
		const state = ctx.getState();
		return this.fetchOrder.fetch({id: state.orderID}, {fetchPolicy: "no-cache", errorPolicy: "all"})
			.subscribe({next: (r) => {
				const connections = toNotNullArray(r.data.connections.edges?.map((c) => c?.node));
				ctx.dispatch(new SetConnections(connections));

				const order = r.data.order;
				if (!!order) {
					ctx.dispatch(new SetOrder(order));
				}
				const timeline = r.data.orderTimeline;
				if (!!timeline) {
					ctx.dispatch(new OrderViewActions.SetTimeline(timeline));
				}
				const statuses = r.data.orderShipments;
				if (!!statuses) {
					ctx.dispatch(new SetShipmentStatuses(statuses));
				}
			}});
	}

	@Action(OrderViewActions.SetOrderID)
	SetOrderID(ctx: StateContext<OrderViewModel>, action: OrderViewActions.SetOrderID) {
		ctx.patchState({
			orderID: action.payload,
		})
	}

	@Action(OrderViewActions.SetOrder)
	SetMyOrder(ctx: StateContext<OrderViewModel>, action: OrderViewActions.SetOrder) {
		ctx.patchState({
			order: action.payload,
		})
	}

	@Action(OrderViewActions.SetTimeline)
	SetTimeline(ctx: StateContext<OrderViewModel>, action: OrderViewActions.SetTimeline) {
		ctx.patchState({timeline: action.payload});
	}

	@Action(OrderViewActions.ResetState)
	ResetState(ctx: StateContext<OrderViewModel>) {
		ctx.setState(defaultState);
	}

	@Action(OrderViewActions.DuplicatePackage)
	DuplicatePackage(ctx: StateContext<OrderViewModel>, action: OrderViewActions.DuplicatePackage) {
		return this.duplicateColli.mutate({fromColliID: action.payload.fromColliID})
			.subscribe((res) => {
				const order = res.data?.duplicateColli;
				if (!!order) {
					ctx.dispatch(new SetOrder(order));
				}
			});
	}

	@Action(OrderViewActions.AddPackage)
	AddPackage(ctx: StateContext<OrderViewModel>, action: OrderViewActions.AddPackage) {
		const state = produce(ctx.getState(),
		st => {
			st.order!.colli!.push(action.payload);
		});
		ctx.setState(state);
	}

	@Action(OrderViewActions.SetPackages)
	SetPackages(ctx: StateContext<OrderViewModel>, action: OrderViewActions.SetPackages) {
		const state = produce(ctx.getState(),
		st => {
			st.order!.colli = action.payload;
		});
		ctx.setState(state);
	}

	@Action(OrderViewActions.SetIsDragging)
	SetIsDragging(ctx: StateContext<OrderViewModel>, action: OrderViewActions.SetIsDragging) {
		ctx.patchState({isDragging: action.payload});
	}

	@Action(OrderViewActions.FireMoveOrderLine)
	FireMoveOrderLine(ctx: StateContext<OrderViewModel>, action: OrderViewActions.FireMoveOrderLine) {
		return this.move.mutate({colliID: action.payload.colliID, orderLineID: action.payload.orderLineID})
			.subscribe((res) => {
				const nextCollies = res.data?.moveOrderLine
				if (!!nextCollies) {
					ctx.dispatch(new OrderViewActions.SetPackages(nextCollies));
				}
			});
	}

	@Action(OrderViewActions.DeletePackage)
	DeletePackage(ctx: StateContext<OrderViewModel>, action: OrderViewActions.DeletePackage) {
		return this.deletePackage.mutate({colliID: action.payload.colliID}, {errorPolicy: "all"})
			.subscribe((res) => {
				const order = res.data?.deleteColli;
				if (!!order) {
					ctx.dispatch(new SetOrder(order));
				} else {
					ctx.dispatch([new ShowGlobalSnackbar("Error: " + JSON.stringify(res.errors))]);
				}
			});
	}

	@Action(OrderViewActions.SaveOrder)
	SaveOrder(ctx: StateContext<OrderViewModel>, action: OrderViewActions.SaveOrder) {
		const state = ctx.getState();
		return this.updateOrder.mutate({id: state.orderID, input: action.payload})
			.subscribe((res) => {
				const nextOrder = res.data?.updateOrder;
				if (!!nextOrder) {
					ctx.dispatch([
						new OrderViewActions.SetOrder(nextOrder),
						new OrderViewActions.SaveOrderSuccess(),
						new ShowGlobalSnackbar(`Order update success`),
					]);
				}
			});
	}

	@Action(OrderViewActions.SetConnections)
	SetConnections(ctx: StateContext<OrderViewModel>, action: OrderViewActions.SetConnections) {
		ctx.patchState({connections: action.payload});
	}

	@Action(OrderViewActions.Clear)
	Clear(ctx: StateContext<OrderViewModel>, action: OrderViewActions.Clear) {
		ctx.setState(defaultState);
	}

	@Action(OrderViewActions.SetShipmentErrors)
	SetShipmentErrors(ctx: StateContext<OrderViewModel>, action: OrderViewActions.SetShipmentErrors) {
		ctx.patchState({shipmentErrors: action.payload.errors, loading: false});
	}

	@Action(OrderViewActions.SetShipmentStatuses)
	SetShipmentStatuses(ctx: StateContext<OrderViewModel>, action: OrderViewActions.SetShipmentStatuses) {
		ctx.patchState({shipmentStatuses: action.payload});
	}

	@Action(OrderViewActions.SetLabels)
	SetLabels(ctx: StateContext<OrderViewModel>, action: OrderViewActions.SetLabels) {
		ctx.patchState({
			labelsPDF: action.payload.labels,
			allLabelsPDF: action.payload.allLabels,
			labelViewerOffset: 0,
			loading: false,
		});
	}

	@Action(OrderViewActions.IncrementLabelViewerOffset)
	IncrementLabelViewerOffset(ctx: StateContext<OrderViewModel>, action: OrderViewActions.IncrementLabelViewerOffset) {
		const current = ctx.getState().labelViewerOffset;
		// Re-used for both packing slips and labels
		if (current < ctx.getState().labelsPDF.length-1 || current < ctx.getState().packingSlipsPDF.length-1) {
			ctx.patchState({labelViewerOffset: current+1});
		}
	}

	@Action(OrderViewActions.DecrementLabelViewerOffset)
	DecrementLabelViewerOffset(ctx: StateContext<OrderViewModel>, action: OrderViewActions.DecrementLabelViewerOffset) {
		const current = ctx.getState().labelViewerOffset;
		if (current > 0) {
			ctx.patchState({labelViewerOffset: current-1});
		}
	}

	@Action(OrderViewActions.FetchPackingSlips)
	FetchPackingSlips(ctx: StateContext<OrderViewModel>, action: OrderViewActions.FetchPackingSlips) {
		ctx.patchState({loading: true});
		return this.fetchPackingSlips.fetch({colliIDs: action.payload.parcelIDs}, {errorPolicy: "all"})
			.subscribe((r) => {
				ctx.patchState({loading: false});
				if (!!r.errors && r.errors.length > 0) {
					ctx.dispatch([
						new SetShipmentErrors({errors: r.errors})
					]);
				} else {
					ctx.dispatch(new SetPackingSlips({packingSlips: r.data.packingSlips.packingSlips, allPackingSlips: r.data.packingSlips.allPackingSlips}));
				}
			});
	}

	@Action(OrderViewActions.SetPackingSlips)
	SetPackingSlips(ctx: StateContext<OrderViewModel>, action: OrderViewActions.SetPackingSlips) {
		ctx.patchState({
			packingSlipsPDF: action.payload.packingSlips,
			allPackingSlipsPDF: action.payload.allPackingSlips,
		});
	}

	@Action(OrderViewActions.FetchShipmentLabels)
	FetchShipmentLabels(ctx: StateContext<OrderViewModel>, action: OrderViewActions.FetchShipmentLabels) {
		ctx.patchState({selectedColliIDs: action.payload});
		// Cleanup the ID vs IDs
		return this.fetchLabels.fetch({colliIDs: action.payload})
			.subscribe((r) => {
				if (!!r.errors && r.errors.length > 0) {
					ctx.dispatch([
						new SetShipmentErrors({errors: r.errors})
					]);
				} else {
					const labels = r.data?.shipmentLabels.labelsPDF;
					const allLabels = r.data?.shipmentLabels.allLabels || "";

					ctx.dispatch(new SetLabels({
						labels: labels,
						allLabels: allLabels,
					}));

				}
			});
	}

	@Action(OrderViewActions.CreateShipments)
	CreateShipments(ctx: StateContext<OrderViewModel>, action: OrderViewActions.CreateShipments) {
		ctx.patchState({loading: true, selectedColliIDs: action.payload.parcelIDs});
		const state = ctx.getState();
		return this.createShipments.fetch({orderID:state.orderID, parcelIDs: action.payload.parcelIDs}, {errorPolicy: "all"})
			.subscribe((r) => {
				if (!!r.errors && r.errors.length > 0) {
					ctx.dispatch([
						new SetShipmentErrors({errors: r.errors})
					]);
				} else {
					const labels = r.data?.createShipments.labelsPDF;
					const allLabels = r.data?.createShipments.allLabels;
					if (!!labels && !!allLabels) {
						ctx.dispatch(new SetLabels({
							labels: labels,
							allLabels: allLabels,
						}));
					}
				}
				ctx.dispatch(new FetchOrder());
			});
	}

	@Action(OrderViewActions.CancelShipment)
	CancelShipment(ctx: StateContext<OrderViewModel>, action: OrderViewActions.CancelShipment) {
		ctx.patchState({loading: true});
		const state = ctx.getState()
		return this.cancelShipment.mutate({colliIDs: state.selectedColliIDs}, {errorPolicy: "all"})
			.subscribe((r) => {
				if (!!r.errors) {
					ctx.dispatch([new ShowGlobalSnackbar("Error cancelling shipment")]);
				} else {
					ctx.dispatch([
						new ShowGlobalSnackbar("Shipment cancelled"),
						new FetchOrder(),
					]);
				}
			});
	}

	@Action(OrderViewActions.ClearDialogs)
	ClearShipmentView(ctx: StateContext<OrderViewModel>, action: OrderViewActions.ClearDialogs) {
		ctx.patchState({
			loading: false,
			allLabelsPDF: '',
			packingSlipsPDF: [],
			allPackingSlipsPDF: '',
			labelViewerOffset: 0,
			selectedColliIDs: [],
			labelsPDF: [],
			shipmentErrors: [],
		});
	}

	@Action(OrderViewActions.CreatePackingSlipPrintJobs)
	CreatePackingSlipPrintJobs(ctx: StateContext<OrderViewModel>, action: OrderViewActions.CreatePackingSlipPrintJobs) {
		return this.createPackingSlipPrintJobs.fetch({colliIDs: action.payload.parcelIDs}, {errorPolicy: "all"})
			.subscribe((resp) => {
				if (!!resp.errors) {
					ctx.dispatch(new ShowGlobalSnackbar("An error occurred: " + resp.errors.join(" ")))
				} else {
					ctx.dispatch([new ShowGlobalSnackbar("Jobs created"), new FetchSelectedWorkstation()]);
				}
			});
	}

	@Action(OrderViewActions.CreateLabelPrintJobs)
	CreateLabelPrintJobs(ctx: StateContext<OrderViewModel>, action: OrderViewActions.CreateLabelPrintJobs) {
		return this.createLabelsPrintJobs.fetch({colliIDs: action.payload.parcelIDs}, {errorPolicy: "all"})
			.subscribe((resp) => {
				if (!!resp.errors) {
					ctx.dispatch(new ShowGlobalSnackbar("An error occurred: " + resp.errors.join(" ")));
				} else {
					ctx.dispatch([new ShowGlobalSnackbar("Jobs created"), new FetchSelectedWorkstation()]);
				}
			});
	}

	@Action(OrderViewActions.PackingSlipsClearCache)
	PackingSlipsClearCache(ctx: StateContext<OrderViewModel>, action: OrderViewActions.PackingSlipsClearCache) {
		return this.packingSlipClearCache.fetch({orderIDs: action.payload.orderIDs}, {errorPolicy: "all"})
			.subscribe((resp) => {
				if (!!resp.errors) {
					ctx.dispatch(new ShowGlobalSnackbar("An error occurred: " + resp.errors.join(" ")));
				} else {
					ctx.dispatch([new ShowGlobalSnackbar("Cache cleared")]);
				}
			});
	}

}
