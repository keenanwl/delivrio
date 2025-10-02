import {Injectable} from "@angular/core";
import {Action, Selector, State, StateContext, Store} from "@ngxs/store";
import {ProfileActions} from "./profile.actions";
import SelectProfileQueryResponse = ProfileActions.SelectProfileQueryResponse;
import SetMyProfile = ProfileActions.SetMyProfile;
import LanguageListQueryResponse = ProfileActions.LanguageListQueryResponse;
import SetLanguageList = ProfileActions.SetLanguageList;
import {ProfileGQL, UpdateUserGQL} from "./profile.generated";
import {toNotNullArray} from "../../functions/not-null-array";
import AvailableTenant = ProfileActions.AvailableTenant;
import SetTenantList = ProfileActions.SetTenantList;
import {AppActions} from "../../app.actions";
import AppChangeRoute = AppActions.AppChangeRoute;
import {Paths} from "src/app/app-routing.module";
import ShowGlobalSnackbar = AppActions.ShowGlobalSnackbar;

export interface ProfileModel {
	myProfileForm: {
		model: SelectProfileQueryResponse | undefined;
		dirty: boolean;
		status: string;
		errors: {};
	},
	languages: LanguageListQueryResponse[];
	availableTenants: AvailableTenant[];
}

const defaultState: ProfileModel = {
	myProfileForm: {
		model: undefined,
		dirty: false,
		status: '',
		errors: {}
	},
	availableTenants: [],
	languages: [],
};

@Injectable()
@State<ProfileModel>({
	name: 'profile',
	defaults: defaultState,
})
export class ProfileState {

	constructor(
		private profile: ProfileGQL,
		private updateUser: UpdateUserGQL,
		private store: Store,
	) {
	}

	@Selector()
	static state(state: ProfileModel) {
		return state;
	}

	@Action(ProfileActions.FetchMyProfile)
	FetchMyProfile(ctx: StateContext<ProfileModel>, action: ProfileActions.FetchMyProfile) {
		return this.profile.fetch()
			.subscribe({next: (r) => {

				const tenants = toNotNullArray(r.data.availableTenants);
				ctx.dispatch(new SetTenantList(tenants));

				const languages = toNotNullArray(r.data.languages.edges?.map((n => n?.node)));
				ctx.dispatch(new SetLanguageList(languages));

				const user = r.data.user;
				if (!!user) {
					ctx.dispatch(new SetMyProfile(user));
				}
			}, error: (e) => {

			}});
	}

	@Action(ProfileActions.SetMyProfile)
	SetMyProfile(ctx: StateContext<ProfileModel>, action: ProfileActions.SetMyProfile) {
		const state = ctx.getState();
		const next = Object.assign({}, state.myProfileForm, {model: action.payload});
		ctx.patchState({
			myProfileForm: next,
		})
	}

	@Action(ProfileActions.SetLanguageList)
	SetLanguageList(ctx: StateContext<ProfileModel>, action: ProfileActions.SetLanguageList) {
		ctx.patchState({languages: action.payload})
	}

	@Action(ProfileActions.SetTenantList)
	SetTenantList(ctx: StateContext<ProfileModel>, action: ProfileActions.SetTenantList) {
		ctx.patchState({availableTenants: action.payload})
	}

	@Action(ProfileActions.SaveForm)
	SaveForm(ctx: StateContext<ProfileModel>, action: ProfileActions.SaveForm) {
		const state = ctx.getState();
		const body = Object.assign({}, state.myProfileForm.model, {
			languageID: state.myProfileForm.model?.language?.id,
			language: undefined,
			tenant: undefined,
		})
		return this.updateUser.mutate({input: body, newTenantID: state.myProfileForm.model?.tenant.id})
			.subscribe({
				next: () => {
					this.store.dispatch(new ShowGlobalSnackbar("Profile successfully saved"));
					window.location.reload();
				},
				error: (e) => {
					this.store.dispatch(new ShowGlobalSnackbar(JSON.stringify(e)));
				}
			});
	}

}
