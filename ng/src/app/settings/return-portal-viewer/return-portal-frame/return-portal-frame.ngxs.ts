import {Injectable} from "@angular/core";
import {Action, Selector, State, StateContext, Store} from "@ngxs/store";
import {ReturnPortalFrameActions} from "./return-portal-frame.actions";
import {
	CreateReturnColliOutput,
	CreateReturnOrderItemInput,
	ReturnPortalFrameService,
	ReturnPortalViewResponse
} from "./return-portal-frame.service";
import {produce} from "immer"
import SetReturnCollis = ReturnPortalFrameActions.SetReturnCollis;
import ShowErrorDialog = ReturnPortalFrameActions.ShowErrorDialog;

export interface ItemReturn {
	quantity: number;
	availableQuantity: number;
	id: string;
	orderLineID: string;
	selected: boolean;
}

export interface ReturnPortalFrameModel {
	step: 0 | 1 | 2 | 3;
	itemsView: ReturnPortalViewResponse | undefined;
	selectedItems: ItemReturn[];
	returnPortalID: string;
	returnCollis: CreateReturnColliOutput[];
	email: string;
	orderPublicID: string;
	loading: boolean;
	comment: string;
	baseURL: string;
}

const defaultState: ReturnPortalFrameModel = {
	step: 0,
	itemsView: undefined,
	selectedItems: [],
	returnPortalID: '',
	returnCollis: [],
	email: '',
	orderPublicID: '',
	loading: false,
	comment: "",
	baseURL: "",
};

@Injectable()
@State<ReturnPortalFrameModel>({
	name: 'returnPortalFrame',
	defaults: defaultState,
})
export class ReturnPortalFrameState {

	constructor(
		private store: Store,
		private service: ReturnPortalFrameService,
	) {}

	@Selector()
	static get(state: ReturnPortalFrameModel) {
		return state;
	}

	@Action(ReturnPortalFrameActions.FetchReturnPortalFrame)
	FetchMyReturnPortalFrame(ctx: StateContext<ReturnPortalFrameModel>, action: ReturnPortalFrameActions.FetchReturnPortalFrame) {
		ctx.patchState({loading: true});
		const state = ctx.getState();
		return this.service.getOrderOverview(state.returnPortalID, state.email, state.orderPublicID)
			.subscribe({
				next: (resp) => {
					ctx.dispatch([
						new ReturnPortalFrameActions.SetReturnPortalFrame(resp),
					]);
				},
				error: (err) => {
					ctx.dispatch([
						new ShowErrorDialog({title: `Error fetching`, body: err.error.message}),
						new ReturnPortalFrameActions.LoadingRundown(),
					]);
				}
			});
	}

	@Action(ReturnPortalFrameActions.SetReturnPortalFrame)
	SetReturnPortalFrame(ctx: StateContext<ReturnPortalFrameModel>, action: ReturnPortalFrameActions.SetReturnPortalFrame) {
		ctx.patchState({
			itemsView: action.payload,
			loading: false,
			step: 1,
		});
	}

	@Action(ReturnPortalFrameActions.SetSelectedItem)
	SetSelectedItem(ctx: StateContext<ReturnPortalFrameModel>, action: ReturnPortalFrameActions.SetSelectedItem) {
		const state = ctx.getState();
		const found = state.selectedItems.find((i) => i.orderLineID === action.payload.item.order_line_id);

		const next = state.selectedItems.filter((i) => i.orderLineID !== action.payload.item.order_line_id);
		if (!!found) {
			next.push(Object.assign({}, found, {selected: action.payload.selected}));
		} else {
			next.push({
				quantity: action.payload.item.quantity,
				availableQuantity: action.payload.item.quantity,
				id: "",
				orderLineID: action.payload.item.order_line_id,
				selected: action.payload.selected,
			});
		}
		ctx.patchState({
			selectedItems: next,
		});
	}

	@Action(ReturnPortalFrameActions.SetSelectedItemReason)
	SetSelectedItemReason(ctx: StateContext<ReturnPortalFrameModel>, action: ReturnPortalFrameActions.SetSelectedItemReason) {
		const state = ctx.getState();
		const found = state.selectedItems.find((i) => i.orderLineID === action.payload.item.order_line_id);

		const next = state.selectedItems.filter((i) => i.orderLineID !== action.payload.item.order_line_id);
		if (!!found) {
			next.push(Object.assign({}, found, {id: action.payload.reasonID}));
		} else {
			next.push({
				quantity: action.payload.item.quantity,
				availableQuantity: action.payload.item.quantity,
				id: action.payload.reasonID,
				orderLineID: action.payload.item.order_line_id,
				selected: true,
			});
		}
		ctx.patchState({
			selectedItems: next,
		});
	}

	@Action(ReturnPortalFrameActions.StopLoading)
	StopLoading(ctx: StateContext<ReturnPortalFrameModel>, action: ReturnPortalFrameActions.StopLoading) {
		ctx.patchState({loading: false});
	}

	@Action(ReturnPortalFrameActions.SetOrderInfo)
	SetOrderInfo(ctx: StateContext<ReturnPortalFrameModel>, action: ReturnPortalFrameActions.SetOrderInfo) {
		ctx.patchState({email: action.payload.email, orderPublicID: action.payload.orderPublicID});
	}

	@Action(ReturnPortalFrameActions.SetReturnPortalID)
	SetReturnPortalID(ctx: StateContext<ReturnPortalFrameModel>, action: ReturnPortalFrameActions.SetReturnPortalID) {
		ctx.patchState({returnPortalID: action.payload});
	}

	@Action(ReturnPortalFrameActions.Clear)
	Clear(ctx: StateContext<ReturnPortalFrameModel>, action: ReturnPortalFrameActions.Clear) {
		ctx.setState(defaultState);
	}

	@Action(ReturnPortalFrameActions.IncrementQuantity)
	IncrementQuantity(ctx: StateContext<ReturnPortalFrameModel>, action: ReturnPortalFrameActions.IncrementQuantity) {
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

	@Action(ReturnPortalFrameActions.DecrementQuantity)
	DecrementQuantity(ctx: StateContext<ReturnPortalFrameModel>, action: ReturnPortalFrameActions.DecrementQuantity) {
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

	@Action(ReturnPortalFrameActions.Save)
	Save(ctx: StateContext<ReturnPortalFrameModel>, action: ReturnPortalFrameActions.Save) {

	}

	@Action(ReturnPortalFrameActions.SetReturnCollis)
	SetReturnCollis(ctx: StateContext<ReturnPortalFrameModel>, action: ReturnPortalFrameActions.SetReturnCollis) {
		ctx.patchState({returnCollis: action.payload, step: 2, loading: false});
	}

	@Action(ReturnPortalFrameActions.UpdateComment)
	UpdateComment(ctx: StateContext<ReturnPortalFrameModel>, action: ReturnPortalFrameActions.UpdateComment) {
		ctx.patchState({comment: action.payload});
	}

	@Action(ReturnPortalFrameActions.SelectDeliveryOption)
	SelectDeliveryOption(ctx: StateContext<ReturnPortalFrameModel>, action: ReturnPortalFrameActions.SelectDeliveryOption) {
		const state = produce(ctx.getState(), st => {
			let next = st.returnCollis;
			for (let i = 0; i < next.length; i++) {
				if (next[i].return_colli_id === action.payload.returnColliID) {
					next[i].selected_delivery_option_id = action.payload.deliveryOptionID;
					break
				}
			}
			st.returnCollis = next;
		});
		ctx.setState(state);
	}

	@Action(ReturnPortalFrameActions.CreateOrder)
	CreateOrder(ctx: StateContext<ReturnPortalFrameModel>, action: ReturnPortalFrameActions.CreateOrder) {
		ctx.patchState({loading: true});
		const state = ctx.getState();
		const body = {
			portal_id: state.returnPortalID,
			order_lines: state.selectedItems.reduce(
				(acc: CreateReturnOrderItemInput[], i) => {
					if (i.selected) {
						acc.push({
							claim_id: i.id,
							order_line_id: i.orderLineID,
							units: i.quantity,
						});
					}
					return acc;
				},
				[]
			),
			comment: state.comment,
		}
		return this.service.createOrder(state.returnPortalID, state.email, state.orderPublicID, body)
			.subscribe({
				next: (r) => {
					const returnCollis = r.return_collis;
					ctx.dispatch(new SetReturnCollis(returnCollis));
				},
				error: (err) => {
					ctx.dispatch([
						new ShowErrorDialog({title: `Error creating return`, body: err.error.message}),
						new ReturnPortalFrameActions.LoadingRundown(),
					]);
				}
			})
	}

	@Action(ReturnPortalFrameActions.SubmitDeliveryOptions)
	SubmitDeliveryOptions(ctx: StateContext<ReturnPortalFrameModel>, action: ReturnPortalFrameActions.SubmitDeliveryOptions) {
		ctx.patchState({loading: true});
		const state = ctx.getState();
		const body = {
			delivery_options:
				state.returnCollis.map((c) => {
					return {
						return_colli_id: c.return_colli_id,
						delivery_option_id: c.selected_delivery_option_id,
					}
				})
		};
		return this.service.submitDeliveryOptions(
			state.returnPortalID,
			state.email,
			state.orderPublicID,
			body
		).subscribe({
			next: (r) => {
				if (r.success) {
					ctx.patchState({step: 3, loading: false});
				}
			},
			error: (err) => {
				ctx.dispatch([
					new ShowErrorDialog({title: `Error selecting shipping option`, body: err.error.message}),
					new ReturnPortalFrameActions.LoadingRundown(),
				]);
			},
		})
	}

	@Action(ReturnPortalFrameActions.SetBaseURL)
	SetBaseURL(ctx: StateContext<ReturnPortalFrameModel>, action: ReturnPortalFrameActions.SetBaseURL) {
		ctx.patchState({baseURL: action.payload});
		this.service.baseURL = action.payload;
	}

}
