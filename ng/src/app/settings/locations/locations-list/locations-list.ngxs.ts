import {Injectable} from "@angular/core";
import {Action, Selector, State, StateContext} from "@ngxs/store";
import {LocationsListActions} from "./locations-list.actions";
import SelectLocationsListQueryResponse = LocationsListActions.LocationsResponse;
import SetLocationsList = LocationsListActions.SetLocationsList;
import {CreateLocationGQL, ListLocationsGQL} from "./locations-list.generated";
import {toNotNullArray} from "../../../functions/not-null-array";
import {AppActions} from "../../../app.actions";
import AppChangeRoute = AppActions.AppChangeRoute;
import {Paths} from "../../../app-routing.module";
import ShowGlobalSnackbar = AppActions.ShowGlobalSnackbar;

export interface LocationsListModel {
	locationsList: SelectLocationsListQueryResponse[];
	loading: boolean;
}

const defaultState: LocationsListModel = {
	locationsList: [],
	loading: false,
};

@Injectable()
@State<LocationsListModel>({
	name: 'locationsList',
	defaults: defaultState,
})
export class LocationsListState {

	constructor(
		private list: ListLocationsGQL,
		private create: CreateLocationGQL,
	) {
	}

	@Selector()
	static get(state: LocationsListModel) {
		return state;
	}

	@Action(LocationsListActions.FetchLocationsList)
	FetchMyLocationsList(ctx: StateContext<LocationsListModel>, action: LocationsListActions.FetchLocationsList) {
		ctx.patchState({loading: true});
		return this.list.fetch({})
			.subscribe({
				next: (r) => {
					ctx.patchState({loading: false});
					const locations = toNotNullArray(r.data.locations.edges?.map((l) => l?.node));
					ctx.dispatch(new SetLocationsList(locations));
				}, error: () => {
					ctx.patchState({loading: false});
				},
			});
	}

	@Action(LocationsListActions.SetLocationsList)
	SetMyLocationsList(ctx: StateContext<LocationsListModel>, action: LocationsListActions.SetLocationsList) {
		ctx.patchState({locationsList: action.payload});
	}

	@Action(LocationsListActions.CreateNewLocation)
	CreateNewLocation(ctx: StateContext<LocationsListModel>, action: LocationsListActions.CreateNewLocation) {
		return this.create.mutate({
			input: {
				name: "New location",
				locationTagIDs: [],
			},
			inputAddress: {
				email: "",
				firstName: "",
				lastName: "",
				phoneNumber: "",
				addressOne: "",
				addressTwo: "",
				city: "",
				state: "",
				countryID: "",
				zip: "",
			}})
			.subscribe((res) => {
				if (!res.errors) {
					ctx.dispatch(new AppChangeRoute({path: Paths.SETTINGS_LOCATION_EDIT, queryParams: {id: res.data?.createLocation?.id}}))
				} else {
					ctx.dispatch(new ShowGlobalSnackbar("Error creating new location"));
				}
			});
	}

	@Action(LocationsListActions.Reset)
	Reset(ctx: StateContext<LocationsListModel>, action: LocationsListActions.Reset) {
		ctx.setState(defaultState);
	}

}
