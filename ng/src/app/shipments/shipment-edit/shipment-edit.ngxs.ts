import {Injectable} from "@angular/core";
import {Action, Selector, State, StateContext} from "@ngxs/store";
import {ShipmentEditActions} from "./shipment-edit.actions";

export interface ShipmentEditModel {
	/*shipmentForm: {
		model: FetchShipmentResponse | undefined;
		dirty: boolean;
		status: string;
		errors: readonly GraphQLError[];
	},*/
	shipmentID: string;
}

const defaultState: ShipmentEditModel = {
	/*shipmentForm: {
		model: undefined,
		dirty: false,
		status: '',
		errors: [],
	},*/
	shipmentID: '',
};

@Injectable()
@State<ShipmentEditModel>({
	name: 'shipmentEdit',
	defaults: defaultState,
})
export class ShipmentEditState {

	constructor() {}

	@Selector()
	static get(state: ShipmentEditModel) {
		return state;
	}

	@Action(ShipmentEditActions.FetchShipment)
	FetchMyShipment(ctx: StateContext<ShipmentEditModel>, action: ShipmentEditActions.FetchShipment) {

	}

}
