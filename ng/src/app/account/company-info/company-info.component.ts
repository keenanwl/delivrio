import {Component, OnInit} from '@angular/core';
import {Observable} from "rxjs";
import {CompanyInfoModel, CompanyInfoState} from "./company-info.ngxs";
import {Actions, ofActionCompleted, Store} from "@ngxs/store";
import {CompanyInfoActions} from "./company-info.actions";
import FetchMyCompanyInfo = CompanyInfoActions.FetchMyCompanyInfo;
import {FormControl, FormGroup} from "@angular/forms";
import SetFormErrors = CompanyInfoActions.SetFormErrors;
import LanguageListQueryResponse = CompanyInfoActions.LanguageResponse;
import SearchCountry = CompanyInfoActions.SearchCountry;
import CountryResponse = CompanyInfoActions.CountryResponse;
import ChangeCountry = CompanyInfoActions.ChangeCountry;

@Component({
	selector: 'app-company-info',
	templateUrl: './company-info.component.html',
	styleUrls: ['./company-info.component.scss']
})
export class CompanyInfoComponent implements OnInit {

	companyInfo$: Observable<CompanyInfoModel>;

	companyAddressForm = new FormGroup({
		addressOne: new FormControl(''),
		addressTwo: new FormControl(''),
		city: new FormControl(''),
		state: new FormControl(''),
		zip: new FormControl(''),
		country: new FormGroup({
			id: new FormControl('', {nonNullable: true}),
			label: new FormControl('', {nonNullable: true}),
			alpha2: new FormControl('', {nonNullable: true}),
		}),
	});

	billingContactForm = new FormGroup({
		name: new FormControl(''),
		surname: new FormControl(''),
		email: new FormControl(''),
		phoneNumber: new FormControl(''),
	});

	adminContactForm = new FormGroup({
		name: new FormControl(''),
		surname: new FormControl(''),
		email: new FormControl(''),
		phoneNumber: new FormControl(''),
	});

	companyInfoForm = new FormGroup({
		name: new FormControl(''),
		invoiceReference: new FormControl(''),
		defaultLanguage: new FormControl(''),
		companyAddress: this.companyAddressForm,
		billingContact: this.billingContactForm,
		adminContact: this.adminContactForm,
	});

	languageComparisonFunction = (option: LanguageListQueryResponse, value: LanguageListQueryResponse): boolean => {
		return option?.id === value?.id;
	}

	constructor(
		private store: Store,
		private actions$: Actions,
	) {
		this.companyInfo$ = store.select(CompanyInfoState.state);
	}

	ngOnInit(): void {
		this.store.dispatch([new FetchMyCompanyInfo()]);
	}

	onSubmit() {
		this.store.dispatch(new CompanyInfoActions.SaveForm());
	}

	searchCountries(term: string) {
		this.store.dispatch(new SearchCountry(term));
	}

	changeCountry(country: CountryResponse) {
		this.store.dispatch(new ChangeCountry(country));
	}

}
