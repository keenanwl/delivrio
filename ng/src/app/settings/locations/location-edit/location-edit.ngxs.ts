import {Injectable} from "@angular/core";
import {Action, Selector, State, StateContext, Store} from "@ngxs/store";
import {LocationEditActions} from "./location-edit.actions";
import SetLocationEdit = LocationEditActions.SetLocationEdit;
import {formErrors} from "../../../account/company-info/company-info.ngxs";
import LocationEditResponse = LocationEditActions.LocationEditResponse;
import {
	FetchLocationGQL,
	LocationSearchCountriesGQL,
	UpdateLocationGQL
} from "./location-edit.generated";
import {produce} from "immer";
import {toNotNullArray} from "../../../functions/not-null-array";
import SetLocationTags = LocationEditActions.SetLocationTags;
import LocationTagResponse = LocationEditActions.LocationTagResponse;
import {AppActions} from "../../../app.actions";
import ShowGlobalSnackbar = AppActions.ShowGlobalSnackbar;
import {CreateAddressInput, UpdateLocationInput} from "../../../../generated/graphql";
import AppChangeRoute = AppActions.AppChangeRoute;
import {Paths} from "../../../app-routing.module";
import {OrderEditModel} from "../../../orders/order-edit/order-edit.ngxs";
import CountriesResponse = LocationEditActions.CountriesResponse;

export interface LocationEditModel {
	locationEditForm: {
		model: LocationEditResponse | undefined;
		dirty: boolean;
		status: string;
		errors: formErrors;
	},
	locationID: string;
	locationTags: LocationTagResponse[];
	searchCountries: CountriesResponse[];
}

const defaultState: LocationEditModel = {
	locationEditForm: {
		model: undefined,
		dirty: false,
		status: '',
		errors: {}
	},
	locationID: '',
	locationTags: [],
	searchCountries: [],
};

@Injectable()
@State<LocationEditModel>({
	name: 'locationEdit',
	defaults: defaultState,
})
export class LocationEditState {

	constructor(
		private fetchLocation: FetchLocationGQL,
		private updateLocation: UpdateLocationGQL,
		private countrySearch: LocationSearchCountriesGQL,
		private store: Store,
	) {}

	@Selector()
	static get(state: LocationEditModel) {
		return state;
	}

	@Action(LocationEditActions.FetchLocationEdit)
	FetchMyLocationEdit(ctx: StateContext<LocationEditModel>, action: LocationEditActions.FetchLocationEdit) {
		const id = ctx.getState().locationID;
		return this.fetchLocation.fetch({id})
			.subscribe({next: (r) => {

				const tags = toNotNullArray(r.data.locationTags.edges?.map((t) => t?.node));
				if (!!tags) {
					ctx.dispatch(new SetLocationTags(tags));
				}

				const location = r.data.location;
				if (!!location) {
					ctx.dispatch(new SetLocationEdit(location));
				}
			}});
	}

	@Action(LocationEditActions.SetLocationID)
	SetLocationID(ctx: StateContext<LocationEditModel>, action: LocationEditActions.SetLocationID) {
		ctx.patchState({locationID: action.payload});
	}

	@Action(LocationEditActions.SetLocationEdit)
	SetLocationEdit(ctx: StateContext<LocationEditModel>, action: LocationEditActions.SetLocationEdit) {
		const state = produce(ctx.getState(), st => {
			st.locationEditForm.model = action.payload;
		});
		ctx.setState(state);
	}

	@Action(LocationEditActions.SetSelectedLocationTags)
	SetSelectedLocationTags(ctx: StateContext<LocationEditModel>, action: LocationEditActions.SetSelectedLocationTags) {
		const state = produce(ctx.getState(), st => {
			st.locationEditForm.model!.locationTags = action.payload;
		});
		ctx.setState(state);
	}

	@Action(LocationEditActions.SetLocationTags)
	SetLocationTags(ctx: StateContext<LocationEditModel>, action: LocationEditActions.SetLocationTags) {
		ctx.patchState({locationTags: action.payload});
	}

	@Action(LocationEditActions.Clear)
	Clear(ctx: StateContext<LocationEditModel>, action: LocationEditActions.Clear) {
		ctx.setState(defaultState);
	}

	@Action(LocationEditActions.Save)
	Save(ctx: StateContext<LocationEditModel>, action: LocationEditActions.Save) {
		const state = ctx.getState();

		const nextAddress: CreateAddressInput = Object.assign({},
			state.locationEditForm.model?.address, {
				countryID: state.locationEditForm.model!.address.country.id,
				country: undefined,
			}
		);

		const next: UpdateLocationInput = Object.assign({},
			state.locationEditForm.model,
			{
				locationTags: undefined,
				address: undefined,
				addLocationTagIDs: state.locationEditForm.model?.locationTags.map((t) => t.id),
			}
		)

		return this.updateLocation.mutate({id: state.locationID, input: next, inputAddress: nextAddress})
			.subscribe((r) => {
				this.store.dispatch([
					new ShowGlobalSnackbar("Location saved successfully"),
					new AppChangeRoute({path: Paths.SETTINGS_LOCATIONS, queryParams: {}}),
				]);
			});

	}

	@Action(LocationEditActions.SearchCountry)
	SearchCountry(ctx: StateContext<OrderEditModel>, action: LocationEditActions.SearchCountry) {
		return this.countrySearch.fetch({term: action.payload})
			.subscribe((res) => {
				const countries = toNotNullArray(res.data.countries.edges?.map((value) => value?.node));
				if (!!countries) {
					ctx.dispatch(new LocationEditActions.SetCountrySearch(countries));
				}
			});
	}

	@Action(LocationEditActions.SetCountrySearch)
	SetCountrySearch(ctx: StateContext<LocationEditModel>, action: LocationEditActions.SetCountrySearch) {
		ctx.patchState({searchCountries: action.payload});
	}

	@Action(LocationEditActions.ChangeCountry)
	ChangeCountrySender(ctx: StateContext<LocationEditModel>, action: LocationEditActions.ChangeCountry) {
		const state = produce(ctx.getState(),
			st => {
				if (!!st.locationEditForm.model?.address.country) {
					st.locationEditForm.model!.address.country = action.payload;
				}
			});
		ctx.setState(state);
	}

}
