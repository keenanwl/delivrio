import {Injectable} from "@angular/core";
import {Action, Selector, State, StateContext} from "@ngxs/store";
import {SeatGroupsListActions} from "./seat-groups-list.actions";
import SetSeatGroupsList = SeatGroupsListActions.SetSeatGroupsList;
import FetchSeatGroups = SeatGroupsListActions.FetchSeatGroups;
import {FetchSeatGroupsGQL} from "./seat-groups-list.generated";

export interface SeatGroupsListModel {
	list: FetchSeatGroups[];
	loading: boolean;
}

const defaultState: SeatGroupsListModel = {
	list: [],
	loading: false,
};

@Injectable()
@State<SeatGroupsListModel>({
	name: 'seatGroupsList',
	defaults: defaultState,
})
export class SeatGroupsListState {

	constructor(
		private list: FetchSeatGroupsGQL,
	) {
	}

	@Selector()
	static get(state: SeatGroupsListModel) {
		return state;
	}

	@Action(SeatGroupsListActions.FetchSeatGroupsList)
	FetchMySeatGroupsList(ctx: StateContext<SeatGroupsListModel>, action: SeatGroupsListActions.FetchSeatGroupsList) {
		ctx.patchState({loading: true});
		return this.list.fetch({}, {fetchPolicy: 'no-cache'})
			.subscribe({next: (r) => {
				ctx.patchState({loading: false});
				const list = r.data.seatGroups.edges?.map((n) => n?.node);
				if (!!list) {
					ctx.dispatch(new SetSeatGroupsList(list));
				}
			}, error: (e) => {
				ctx.patchState({loading: false});
			}});
	}

	@Action(SeatGroupsListActions.SetSeatGroupsList)
	SetMySeatGroupsList(ctx: StateContext<SeatGroupsListModel>, action: SeatGroupsListActions.SetSeatGroupsList) {
		ctx.patchState({
			list: action.payload,
		})
	}
	/*
	@Action(SeatGroupsListActions.CreateNewGLSDeliveryOption)
	CreateNewGLSDeliveryOption(ctx: StateContext<SeatGroupsListModel>, action: SeatGroupsListActions.CreateNewGLSDeliveryOption) {
		return this.createGLS.mutate({input: {
			name: action.payload.name, carrierID: action.payload.agreementId, groupID: ""}}).subscribe((r) => {
			ctx.dispatch(new AppChangeRoute({path: `/settings/delivery-options/edit`, queryParams: {id: r.data?.createDeliveryOptionGLS.id}}));
		});
	}

	@Action(SeatGroupsListActions.FetchCarrierAgreements)
	FetchCarrierAgreements(ctx: StateContext<SeatGroupsListModel>, action: SeatGroupsListActions.FetchCarrierAgreements) {
		const app = this.store.selectSnapshot(AppState.get);
		return this.agreements.fetch({groupId: app.my_ids.my_group_pulid})
			.subscribe({next: (r) => {
					const conns = r.data.carriers.edges;
					if (!!conns) {
						ctx.dispatch(new SetCarrierAgreements(conns));
					}
				}, error: (e) => {

				}});
	}

	@Action(SeatGroupsListActions.SetCarrierAgreements)
	SetCarrierAgreements(ctx: StateContext<SeatGroupsListModel>, action: SeatGroupsListActions.SetCarrierAgreements) {
		ctx.patchState({
			agreements: action.payload,
		})
	}*/

}
