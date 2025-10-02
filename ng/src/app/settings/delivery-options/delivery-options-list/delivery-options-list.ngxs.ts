import {Injectable} from "@angular/core";
import {Action, Selector, State, StateContext} from "@ngxs/store";
import {DeliveryOptionsListActions} from "./delivery-options-list.actions";
import SelectDeliveryOptionsListQueryResponse = DeliveryOptionsListActions.SelectDeliveryOptionsListQueryResponse;
import SetDeliveryOptionsList = DeliveryOptionsListActions.SetDeliveryOptionsList;
import {AppActions} from "../../../app.actions";
import AppChangeRoute = AppActions.AppChangeRoute;
import SetCarrierAgreements = DeliveryOptionsListActions.SetCarrierAgreements;
import FetchCarrierAgreementsQueryResponse = DeliveryOptionsListActions.FetchCarrierAgreementsQueryResponse;
import {
	CreateDeliveryOptionGQL, DeliveryOptionArchiveGQL,
	FetchCarrierAgreementsAndConnectionsGQL,
	ListDeliveryOptionsGQL, UpdateDeliveryOptionSortOrderGQL
} from "./delivery-options-list.generated";
import {Paths} from "../../../app-routing.module";
import {toNotNullArray} from "../../../functions/not-null-array";
import {CarrierBrandInternalId, DeliveryOptionWhereInput} from "../../../../generated/graphql";
import FetchConnectionsResponse = DeliveryOptionsListActions.FetchConnectionsResponse;
import ShowGlobalSnackbar = AppActions.ShowGlobalSnackbar;
import FetchDeliveryOptionsList = DeliveryOptionsListActions.FetchDeliveryOptionsList;

export interface DeliveryOptionsListModel {
	deliveryOptionsList: SelectDeliveryOptionsListQueryResponse[];
	agreements: FetchCarrierAgreementsQueryResponse[];
	connections: FetchConnectionsResponse[];
	showArchived: boolean;
	loading: boolean;
}

const defaultState: DeliveryOptionsListModel = {
	deliveryOptionsList: [],
	connections: [],
	agreements: [],
	showArchived: false,
	loading: false,
};

@Injectable()
@State<DeliveryOptionsListModel>({
	name: 'deliveryOptionsList',
	defaults: defaultState,
})
export class DeliveryOptionsListState {

	constructor(
		private createDeliveryOption: CreateDeliveryOptionGQL,
		private list: ListDeliveryOptionsGQL,
		private agreementsConnections: FetchCarrierAgreementsAndConnectionsGQL,
		private updateSortOrder: UpdateDeliveryOptionSortOrderGQL,
		private archive: DeliveryOptionArchiveGQL,
	) {
	}

	@Selector()
	static get(state: DeliveryOptionsListModel) {
		return state;
	}

	@Action(DeliveryOptionsListActions.FetchDeliveryOptionsList)
	FetchMyDeliveryOptionsList(ctx: StateContext<DeliveryOptionsListModel>, action: DeliveryOptionsListActions.FetchDeliveryOptionsList) {
		ctx.patchState({loading: true});
		const state = ctx.getState();
		return this.list.fetch({showArchived: state.showArchived})
			.subscribe({next: (r) => {
				ctx.patchState({loading: false});
				const list = toNotNullArray(r.data.deliveryOptionsFiltered);
				ctx.dispatch(new SetDeliveryOptionsList(list));
			}});
	}

	@Action(DeliveryOptionsListActions.SetDeliveryOptionsList)
	SetMyDeliveryOptionsList(ctx: StateContext<DeliveryOptionsListModel>, action: DeliveryOptionsListActions.SetDeliveryOptionsList) {
		ctx.patchState({deliveryOptionsList: action.payload});
	}

	@Action(DeliveryOptionsListActions.CreateNewDeliveryOption)
	CreateNewDeliveryOption(ctx: StateContext<DeliveryOptionsListModel>, action: DeliveryOptionsListActions.CreateNewDeliveryOption) {
		return this.createDeliveryOption.mutate({
			name: action.payload.name,
			agreementID: action.payload.agreementId,
			connectionID: action.payload.connectionID,
		}).subscribe((r) => {

			const resp = r.data?.createDeliveryOption;
			if (!!resp) {
				let path = "";
				switch (resp.carrier) {
					case CarrierBrandInternalId.Bring:
						path = Paths.SETTINGS_DELIVERY_OPTIONS_EDIT_BRING;
						break;
					case CarrierBrandInternalId.Dao:
						path = Paths.SETTINGS_DELIVERY_OPTIONS_EDIT_DAO;
						break;
					case CarrierBrandInternalId.Df:
						path = Paths.SETTINGS_DELIVERY_OPTIONS_EDIT_DF;
						break;
					case CarrierBrandInternalId.Dsv:
						path = Paths.SETTINGS_DELIVERY_OPTIONS_EDIT_DSV;
						break;
					case CarrierBrandInternalId.EasyPost:
						path = Paths.SETTINGS_DELIVERY_OPTIONS_EDIT_EASY_POST;
						break;
					case CarrierBrandInternalId.Gls:
						path = Paths.SETTINGS_DELIVERY_OPTIONS_EDIT_GLS;
						break;
					case CarrierBrandInternalId.PostNord:
						path = Paths.SETTINGS_DELIVERY_OPTIONS_EDIT_POST_NORD;
						break;
					case CarrierBrandInternalId.Usps:
						path = Paths.SETTINGS_DELIVERY_OPTIONS_EDIT_USPS;
						break
					default:
						ctx.dispatch(new ShowGlobalSnackbar("Carrier not found: " + resp.carrier));
						return;
				}

				ctx.dispatch(new AppChangeRoute({
					path,
					queryParams: {id: resp.id}
				}));
			}

		});
	}

	@Action(DeliveryOptionsListActions.FetchCarrierAgreements)
	FetchCarrierAgreements(ctx: StateContext<DeliveryOptionsListModel>, action: DeliveryOptionsListActions.FetchCarrierAgreements) {
		return this.agreementsConnections.fetch({})
			.subscribe({next: (r) => {
				const agreements = toNotNullArray(r.data.carriers.edges?.map((c) => c?.node));
				if (!!agreements) {
					ctx.dispatch(new SetCarrierAgreements(agreements));
				}
				const conns = toNotNullArray(r.data.connections.edges?.map((c) => c?.node));
				if (!!conns) {
					ctx.dispatch(new DeliveryOptionsListActions.SetConnections(conns));
				}
			}});
	}

	@Action(DeliveryOptionsListActions.SetCarrierAgreements)
	SetCarrierAgreements(ctx: StateContext<DeliveryOptionsListModel>, action: DeliveryOptionsListActions.SetCarrierAgreements) {
		ctx.patchState({agreements: action.payload});
	}

	@Action(DeliveryOptionsListActions.SetConnections)
	SetConnections(ctx: StateContext<DeliveryOptionsListModel>, action: DeliveryOptionsListActions.SetConnections) {
		ctx.patchState({connections: action.payload});
	}

	@Action(DeliveryOptionsListActions.Clear)
	Clear(ctx: StateContext<DeliveryOptionsListModel>, action: DeliveryOptionsListActions.Clear) {
		ctx.setState(defaultState);
	}

	@Action(DeliveryOptionsListActions.ToggleShowArchive)
	ToggleShowArchive(ctx: StateContext<DeliveryOptionsListModel>, action: DeliveryOptionsListActions.ToggleShowArchive) {
		ctx.patchState({showArchived: !ctx.getState().showArchived});
		ctx.dispatch(new FetchDeliveryOptionsList());
	}

	@Action(DeliveryOptionsListActions.UpdateSortOrder)
	UpdateSortOrder(ctx: StateContext<DeliveryOptionsListModel>, action: DeliveryOptionsListActions.UpdateSortOrder) {
		const state = ctx.getState();
		const nextOrder: string[] = [];
		state.deliveryOptionsList.forEach((li, index) => {
			if (li.id !== action.payload.deliveryOptionID) {
				nextOrder.push(li.id);
			}
		});

		nextOrder.splice(action.payload.nextIndex, 0, action.payload.deliveryOptionID);

		return this.updateSortOrder.mutate({nextSortOrder: nextOrder})
			.subscribe((r) => {
				const list = r.data?.updateDeliveryOptionSortOrder;
				if (!!list) {
					ctx.dispatch(new SetDeliveryOptionsList(list));
				}
			});
	}

	@Action(DeliveryOptionsListActions.Archive)
	Archive(ctx: StateContext<DeliveryOptionsListModel>, action: DeliveryOptionsListActions.Archive) {
		ctx.patchState({loading: true});
		return this.archive.mutate({id: action.payload.deliveryOptionID}, {errorPolicy: "all"})
			.subscribe((resp) => {
				ctx.patchState({loading: false});
				if (!!resp.errors) {
					ctx.dispatch(new ShowGlobalSnackbar("Error archiving: " + JSON.stringify(resp.errors)));
				} else {
					ctx.dispatch(new FetchDeliveryOptionsList());
				}
			});
	}

}
