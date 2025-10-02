import {formErrors} from "./company-info.ngxs";
import {CompanyInfoQuery} from "./company-info.generated";

export namespace CompanyInfoActions {
	export class FetchMyCompanyInfo {
		static readonly type = '[CompanyInfo] fetch my CompanyInfo';
	}
	export class SetMyCompanyInfo {
		static readonly type = '[CompanyInfo] set my CompanyInfo';
		constructor(public payload: SelectCompanyInfoQueryResponse | undefined) {}
	}
	export class SetLanguageList {
		static readonly type = '[CompanyInfo] set language list';
		constructor(public payload: LanguageResponse[]) {}
	}
	export class SetCountries {
		static readonly type = '[CompanyInfo] set countries';
		constructor(public payload: CountryResponse[]) {}
	}
	export class SearchCountry {
		static readonly type = '[CompanyInfo] search country';
		constructor(public payload: string) {}
	}
	export class SetSearchCountry {
		static readonly type = '[CompanyInfo] set search country';
		constructor(public payload: CountryResponse[]) {}
	}
	export class ChangeCountry {
		static readonly type = '[CompanyInfo] change country';
		constructor(public payload: CountryResponse) {}
	}
	export class SetFormErrors {
		static readonly type = '[CompanyInfo] set form errors';
		constructor(public payload: formErrors) {}
	}
	export class SaveForm {
		static readonly type = '[CompanyInfo] save form';
	}
	export type SelectCompanyInfoQueryResponse = NonNullable<CompanyInfoQuery['tenant']>;
	export type LanguageResponse = NonNullable<NonNullable<NonNullable<CompanyInfoQuery['languages']['edges']>[0]>['node']>;
	export type CountryResponse = NonNullable<NonNullable<NonNullable<CompanyInfoQuery['countries']['edges']>[0]>['node']>;
}
