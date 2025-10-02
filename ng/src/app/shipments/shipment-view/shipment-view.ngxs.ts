import {Injectable} from "@angular/core";
import {Action, Selector, State, StateContext} from "@ngxs/store";
import {ShipmentViewActions} from "./shipment-view.actions";
import {
	CancelCancelSyncGQL, CancelFulfillmentSyncGQL,
	DebugUpdateLabelIDsGQL,
	FetchShipmentGQL,
	ShipmentViewCancelShipmentGQL
} from "./shipment-view.generated";
import FetchShipmentResponse = ShipmentViewActions.FetchShipmentResponse;
import {AppActions} from "../../app.actions";
import ShowGlobalSnackbar = AppActions.ShowGlobalSnackbar;

export interface ShipmentViewModel {
	shipment: FetchShipmentResponse | undefined;
	shipmentID: string;
}

const defaultState: ShipmentViewModel = {
	shipment: undefined,
	shipmentID: '',
};

@Injectable()
@State<ShipmentViewModel>({
	name: 'shipmentView',
	defaults: defaultState,
})
export class ShipmentViewState {

	constructor(
		private shipment: FetchShipmentGQL,
		private cancel: ShipmentViewCancelShipmentGQL,
		private debugIDs: DebugUpdateLabelIDsGQL,
		private cancelCancelSync: CancelCancelSyncGQL,
		private cancelFulfillmentSync: CancelFulfillmentSyncGQL,
	) {}

	@Selector()
	static get(state: ShipmentViewModel) {
		return state;
	}

	@Action(ShipmentViewActions.FetchShipment)
	FetchShipment(ctx: StateContext<ShipmentViewModel>, action: ShipmentViewActions.FetchShipment) {
		return this.shipment.fetch({id: ctx.getState().shipmentID})
			.subscribe((res) => {
				const shipment = res.data.shipment;
				if (!!shipment) {
					ctx.dispatch(new ShipmentViewActions.SetShipment(shipment));
				}
			});
	}

	@Action(ShipmentViewActions.SetShipment)
	SetShipment(ctx: StateContext<ShipmentViewModel>, action: ShipmentViewActions.SetShipment) {
		ctx.patchState({shipment: action.payload});
	}

	@Action(ShipmentViewActions.SetShipmentID)
	SetShipmentID(ctx: StateContext<ShipmentViewModel>, action: ShipmentViewActions.SetShipmentID) {
		ctx.patchState({shipmentID: action.payload});
	}

	@Action(ShipmentViewActions.CancelShipment)
	CancelShipment(ctx: StateContext<ShipmentViewModel>, action: ShipmentViewActions.CancelShipment) {
		return this.cancel.mutate({shipmentID: ctx.getState().shipmentID})
			.subscribe((res) => {
				ctx.dispatch(new ShipmentViewActions.FetchShipment());
			});
	}

	@Action(ShipmentViewActions.CancelCancelSync)
	CancelCancelSync(ctx: StateContext<ShipmentViewModel>, action: ShipmentViewActions.CancelCancelSync) {
		return this.cancelCancelSync.mutate({shipmentParcelID: action.payload})
			.subscribe((res) => {
				if (!!res.errors) {
					ctx.dispatch(new ShowGlobalSnackbar("Error: " + JSON.stringify(res.errors)));
				} else {
					ctx.dispatch(new ShipmentViewActions.FetchShipment());
				}
			});
	}

	@Action(ShipmentViewActions.CancelFulfillmentSync)
	CancelFulfillmentSync(ctx: StateContext<ShipmentViewModel>, action: ShipmentViewActions.CancelFulfillmentSync) {
		return this.cancelFulfillmentSync.mutate({shipmentParcelID: action.payload})
			.subscribe((res) => {
				if (!!res.errors) {
					ctx.dispatch(new ShowGlobalSnackbar("Error: " + JSON.stringify(res.errors)));
				} else {
					ctx.dispatch(new ShipmentViewActions.FetchShipment());
				}
			});
	}

	@Action(ShipmentViewActions.DebugUpdateLabelIDs)
	DebugUpdateLabelIDs(ctx: StateContext<ShipmentViewModel>, action: ShipmentViewActions.DebugUpdateLabelIDs) {
		return this.debugIDs.mutate({parcelID: action.payload.parcelID, itemID: action.payload.itemID})
			.subscribe((res) => {
				ctx.dispatch(new ShipmentViewActions.FetchShipment());
			});
	}

	@Action(ShipmentViewActions.Clear)
	Clear(ctx: StateContext<ShipmentViewModel>, action: ShipmentViewActions.Clear) {
		ctx.setState(defaultState);
	}

}
