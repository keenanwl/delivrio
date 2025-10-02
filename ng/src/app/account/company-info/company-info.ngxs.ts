import {Injectable} from "@angular/core";
import {Action, Selector, State, StateContext, Store} from "@ngxs/store";
import SelectCompanyInfoQueryResponse = CompanyInfoActions.SelectCompanyInfoQueryResponse;
import {AppState} from "../../app.ngxs";
import SetMyCompanyInfo = CompanyInfoActions.SetMyCompanyInfo;
import LanguageListQueryResponse = CompanyInfoActions.LanguageResponse;
import SetLanguageList = CompanyInfoActions.SetLanguageList;
import {CompanyInfoActions} from "./company-info.actions";
import {UpdateFormErrors} from "@ngxs/form-plugin";
import {CompanyInfoGQL, CompanyInfoSearchCountriesGQL, UpdateCompanyInfoGQL} from "./company-info.generated";
import {toNotNullArray} from "../../functions/not-null-array";
import CountryResponse = CompanyInfoActions.CountryResponse;
import {produce} from "immer";

export interface CompanyInfoModel {
	companyInfoForm: {
		model: SelectCompanyInfoQueryResponse | undefined;
		dirty: boolean;
		status: string;
		errors: formErrors;
	},
	languages: LanguageListQueryResponse[];
	countries: CountryResponse[];
	searchCountries: CountryResponse[],
}

const defaultState: CompanyInfoModel = {
	companyInfoForm: {
		model: undefined,
		dirty: false,
		status: '',
		errors: {}
	},
	countries: [],
	languages: [],
	searchCountries: [],
};

@Injectable()
@State<CompanyInfoModel>({
	name: 'companyInfo',
	defaults: defaultState,
})
export class CompanyInfoState {

	constructor(
		private companyInfo: CompanyInfoGQL,
		private updateCompany: UpdateCompanyInfoGQL,
		private countrySearch: CompanyInfoSearchCountriesGQL,
		private store: Store,
	) {
	}

	@Selector()
	static state(state: CompanyInfoModel) {
		return state;
	}

	@Action(CompanyInfoActions.FetchMyCompanyInfo)
	FetchMyCompanyInfo(ctx: StateContext<CompanyInfoModel>, action: CompanyInfoActions.FetchMyCompanyInfo) {
		const app = this.store.selectSnapshot(AppState.get);
		return this.companyInfo.fetch({id: app.my_ids.my_tenant_pulid})
			.subscribe({next: (r) => {
				const company = r.data.tenant;
				if (!!company) {
					ctx.dispatch(new SetMyCompanyInfo(company));
				}
				const languages = toNotNullArray(r.data.languages.edges?.map((n) => n?.node));
				ctx.dispatch(new SetLanguageList(languages));

				const countries = toNotNullArray(r.data.countries.edges?.map((n) => n?.node));
				ctx.dispatch(new SetLanguageList(languages));

			}, error: (e) => {

			}});
	}

	@Action(CompanyInfoActions.SetMyCompanyInfo)
	SetMyCompanyInfo(ctx: StateContext<CompanyInfoModel>, action: CompanyInfoActions.SetMyCompanyInfo) {
		const state = ctx.getState();
		const next = Object.assign({}, state.companyInfoForm, {model: action.payload});
		ctx.patchState({
			companyInfoForm: next,
		})
	}

	@Action(CompanyInfoActions.SetLanguageList)
	SetLanguageList(ctx: StateContext<CompanyInfoModel>, action: CompanyInfoActions.SetLanguageList) {
		ctx.patchState({
			languages: action.payload,
		})
	}

	@Action(CompanyInfoActions.SaveForm)
	SaveForm(ctx: StateContext<CompanyInfoModel>, action: CompanyInfoActions.SaveForm) {
		const state = ctx.getState();

		const bodyCompany = {name: state.companyInfoForm.model?.name, invoiceReference: state.companyInfoForm.model?.invoiceReference};
		const defaultLanguageID = state.companyInfoForm.model?.defaultLanguage.id;
		const adminContact = state.companyInfoForm.model?.adminContact;
		const billingContact = state.companyInfoForm.model?.billingContact;
		const address = Object.assign({},
			{firstName: "", lastName: "", email: "", phoneNumber: ""},
			state.companyInfoForm.model?.companyAddress,
			{countryID: state.companyInfoForm.model!.companyAddress!.country.id, country: undefined},
		);

		if (!!defaultLanguageID && !!adminContact && !!billingContact && !!address) {
			return this.updateCompany.mutate({
				input: bodyCompany,
				defaultLanguage: defaultLanguageID,
				adminContact,
				billingContact,
				address,
			});
		}

	}

	@Action(CompanyInfoActions.SetFormErrors)
	SetAdminFormErrors(ctx: StateContext<CompanyInfoModel>, action: CompanyInfoActions.SetFormErrors) {
		const state = ctx.getState();

		this.store.dispatch(new UpdateFormErrors({
			errors: mergeNestedErrors(state.companyInfoForm.errors, action.payload) as {[k: string]: string},
			path: "companyInfo.companyInfoForm"
		}))
	}

	@Action(CompanyInfoActions.SetCountries)
	SetCountries(ctx: StateContext<CompanyInfoModel>, action: CompanyInfoActions.SetCountries) {
		ctx.patchState({countries: action.payload});
	}

	@Action(CompanyInfoActions.SetSearchCountry)
	SetSearchCountry(ctx: StateContext<CompanyInfoModel>, action: CompanyInfoActions.SetSearchCountry) {
		ctx.patchState({searchCountries: action.payload});
	}

	@Action(CompanyInfoActions.SearchCountry)
	SearchCountry(ctx: StateContext<CompanyInfoModel>, action: CompanyInfoActions.SearchCountry) {
		return this.countrySearch.fetch({term: action.payload})
			.subscribe((res) => {
				const countries = toNotNullArray(res.data.countries.edges?.map((value) => value?.node));
				ctx.dispatch(new CompanyInfoActions.SetSearchCountry(countries));
			});
	}

	@Action(CompanyInfoActions.ChangeCountry)
	ChangeCountry(ctx: StateContext<CompanyInfoModel>, action: CompanyInfoActions.ChangeCountry) {
		const state = produce(ctx.getState(),
			st => {
				if (!!st.companyInfoForm.model?.companyAddress?.country) {
					st.companyInfoForm.model!.companyAddress!.country = action.payload;
				}
			});
		ctx.setState(state);
	}

}

export type formErrors = {[key: string]: string | {[key: string]: string}};

export function mergeNestedErrors(obj1: formErrors, obj2: formErrors): formErrors {

	const next: formErrors = {}
	for (const key in obj1) {
		const val = obj1[key];
		if (typeof val === "object") {
			next[key] = Object.assign({}, val);
		} else {
			next[key] = val;
		}
	}

	for (const key in obj2) {
		const val = obj2[key];
		if (typeof val === "object") {
			next[key] = Object.assign({}, val);
		} else {
			next[key] = val;
		}
	}

	return next;

}
