import {Injectable} from "@angular/core";
import {Action, Selector, State, StateContext, Store} from "@ngxs/store";
import {ShipmentsListActions} from "./shipments-list.actions";
import ShipmentsResponse = ShipmentsListActions.ShipmentsResponse;
import {
	FetchShipmentsGQL,
	ShipmentSendOverviewEmailGQL,
	ShipmentsSearchCcLocationsGQL
} from "./shipments-list.generated";
import SetShipments = ShipmentsListActions.SetShipments;
import {toNotNullArray} from "../../functions/not-null-array";
import {filterCategoryList, searchInputList, selectedOptionList} from "../../shared/filter-bar/filter-bar.component";
import SearchCCLocations = ShipmentsListActions.SearchCCLocations;
import {AppActions} from "../../app.actions";
import ShowGlobalSnackbar = AppActions.ShowGlobalSnackbar;
import SetListOptions = ShipmentsListActions.SetListOptions;
import FetchShipments = ShipmentsListActions.FetchShipments;
import {ShipmentWhereInput} from "../../../generated/graphql";
import EmailTemplatesResponse = ShipmentsListActions.EmailTemplatesResponse;
import SetEmailTemplates = ShipmentsListActions.SetEmailTemplates;

export interface ShipmentsPagination {
	cursor: string | null;
	totalPages: number;
	pageIndex: number;
	hasNextPage: boolean;
	hasPreviousPage: boolean;
	startCursor: string | null;
	endCursor: string | null;
}

export interface ShipmentsListModel {
	shipments: ShipmentsResponse[];
	pagination: ShipmentsPagination;
	filterCategories: filterCategoryList;
	shipmentRowsSelected: Set<string>;
	searchListOptions: searchInputList;
	selectedFilters: selectedOptionList;
	emailTemplates: EmailTemplatesResponse[];
	loading: boolean;
}

enum ShipmentFilterCategories {
	cc_locations = "cc_locations"
}

const defaultState: ShipmentsListModel = {
	shipments: [],
	pagination: {
		cursor: null,
		totalPages: 0,
		pageIndex: 0,
		hasNextPage: true,
		hasPreviousPage: false,
		startCursor: null,
		endCursor: null,
	},
	shipmentRowsSelected: new Set(),
	filterCategories: [
		{
			name: 'C&C Locations',
			id: ShipmentFilterCategories.cc_locations.toString(),
			icon: 'location_city'
		},
	],
	searchListOptions: [],
	selectedFilters: [],
	emailTemplates: [],
	loading: false,
};

@Injectable()
@State<ShipmentsListModel>({
	name: 'shipmentsList',
	defaults: defaultState,
})
export class ShipmentsListState {

	constructor(
		private store: Store,
		private shipments: FetchShipmentsGQL,
		private searchCC: ShipmentsSearchCcLocationsGQL,
		private sendEmail: ShipmentSendOverviewEmailGQL,
	) {
	}

	@Selector()
	static state(state: ShipmentsListModel) {
		return state;
	}

	filtersToWhere(filters: selectedOptionList) {
		const where: ShipmentWhereInput = {and: []};
		filters.forEach((f) => {
/*			if (f.filterID === ShipmentFilterCategories.cc_locations) {
				where.and?.push({hasColliWith: [{hasClickCollectLocationWith: [{id: f.optionID}]}]});
			}*/
		})
		return where;
	}

	@Action(ShipmentsListActions.FetchShipments)
	FetchShipments(ctx: StateContext<ShipmentsListModel>, action: ShipmentsListActions.FetchShipments) {
		const filters = ctx.getState().selectedFilters;
		const where = this.filtersToWhere(filters);
		ctx.patchState({loading: true});
		return this.shipments.fetch({where: where})
			.subscribe((res) => {
				ctx.patchState({loading: false});
				const shipments = toNotNullArray(res.data.shipments.edges?.map((n) => n?.node));
				ctx.dispatch(new SetShipments(shipments));

				const templates = toNotNullArray(res.data.emailTemplates.edges?.map((n) => n?.node));
				ctx.dispatch(new SetEmailTemplates(templates));
			});
	}

	@Action(ShipmentsListActions.SetShipments)
	SetShipments(ctx: StateContext<ShipmentsListModel>, action: ShipmentsListActions.SetShipments) {
		ctx.patchState({shipments: action.payload});
	}

	@Action(ShipmentsListActions.ToggleAll)
	ToggleAll(ctx: StateContext<ShipmentsListModel>, action: ShipmentsListActions.ToggleAll) {
		const state = ctx.getState();
		let next = new Set(state.shipmentRowsSelected);
		if (next.size !== state.shipments.length) {
			//state.shipments.forEach((o) => next.add(`${o.id}`))
		} else {
			next.clear();
		}

		ctx.patchState({
			shipmentRowsSelected: next,
		});
	}

	@Action(ShipmentsListActions.ToggleRows)
	ToggleRows(ctx: StateContext<ShipmentsListModel>, action: ShipmentsListActions.ToggleRows) {
		const state = ctx.getState();
		const next = new Set(state.shipmentRowsSelected);
		//action.payload.forEach((o) => next.has(`${o.id}`) ? next.delete(`${o.id}`) : next.add(`${o.id}`));
		ctx.patchState({shipmentRowsSelected: next});
	}

	@Action(ShipmentsListActions.SetListOptions)
	SetListOptions(ctx: StateContext<ShipmentsListModel>, action: ShipmentsListActions.SetListOptions) {
		ctx.patchState({searchListOptions: action.payload});
	}

	@Action(ShipmentsListActions.SetSelectedFilters)
	SetSelectedFilters(ctx: StateContext<ShipmentsListModel>, action: ShipmentsListActions.SetSelectedFilters) {
		ctx.patchState({selectedFilters: action.payload});
		ctx.dispatch(new FetchShipments());
	}

	@Action(ShipmentsListActions.SetEmailTemplates)
	SetEmailTemplates(ctx: StateContext<ShipmentsListModel>, action: ShipmentsListActions.SetEmailTemplates) {
		ctx.patchState({emailTemplates: action.payload});
	}

	@Action(ShipmentsListActions.SearchCCLocations)
	SearchCCLocations(ctx: StateContext<ShipmentsListModel>, action: ShipmentsListActions.SearchCCLocations) {
		return this.searchCC.fetch({lookup: action.payload.lookup})
			.subscribe((r) => {
				const list = toNotNullArray(r.data.locations.edges?.map((n) => n?.node));
				ctx.dispatch(new SetListOptions(list));
			});
	}

	@Action(ShipmentsListActions.SendOverviewEmail)
	SendOverviewEmail(ctx: StateContext<ShipmentsListModel>, action: ShipmentsListActions.SendOverviewEmail) {
		const filters = ctx.getState().selectedFilters;
		const where = this.filtersToWhere(filters);
		return this.sendEmail.fetch({email: action.payload.email, templateID: action.payload.templateID, where})
			.subscribe((r) => {
				if (!!r.errors) {
					ctx.dispatch(new ShowGlobalSnackbar("Email sent to " + action.payload.email));
				} else {
					ctx.dispatch(new ShowGlobalSnackbar("Error " + JSON.stringify(r.errors)));
				}
			});
	}

	@Action(ShipmentsListActions.SearchFilterChanges)
	SearchFilterChanges(ctx: StateContext<ShipmentsListModel>, action: ShipmentsListActions.SearchFilterChanges) {
		switch (action.payload.filterID) {
			case ShipmentFilterCategories.cc_locations:
				ctx.dispatch(new SearchCCLocations({lookup: action.payload.lookup}));
				break;
			default:
				ctx.dispatch(new ShowGlobalSnackbar("Unknow filter: " + action.payload.filterID));
				break;
		}
	}

}
