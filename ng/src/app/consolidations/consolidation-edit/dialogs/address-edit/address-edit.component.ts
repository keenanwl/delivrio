import {Component, Input} from '@angular/core';
import {FormControl, FormGroup, FormGroupDirective, NgForm, ReactiveFormsModule} from "@angular/forms";
import {DvoCardComponent} from "../../../../shared/dvo-card/dvo-card.component";
import {AsyncPipe, NgForOf, NgIf} from "@angular/common";
import {MaterialModule} from "../../../../modules/material.module";
import {DialogRef} from "@angular/cdk/dialog";
import {AddressInfoFragment} from "../../../../orders/order-edit/order-edit.generated";
import {ConsolidationEditModel, ConsolidationEditState} from "../../consolidation-edit.ngxs";
import {Store} from "@ngxs/store";
import {Observable} from "rxjs";
import {ConsolidationEditActions} from "../../consolidation-edit.actions";
import SetSender = ConsolidationEditActions.SetSender;
import SetRecipient = ConsolidationEditActions.SetRecipient;
import SearchCountries = ConsolidationEditActions.SearchCountries;
import CountryResult = ConsolidationEditActions.CountryResult;
import {ErrorStateMatcher} from "@angular/material/core";

export class CountrySelectorErrorStateMatcher implements ErrorStateMatcher {
	isErrorState(control: FormControl | null, form: FormGroupDirective | NgForm | null): boolean {
		console.warn(control)
		return control?.invalid || false;
	}
}

@Component({
	selector: 'app-address-edit',
	standalone: true,
	imports: [
		DvoCardComponent,
		MaterialModule,
		NgForOf,
		NgIf,
		ReactiveFormsModule,
		AsyncPipe,
	],
	templateUrl: './address-edit.component.html',
	styleUrl: './address-edit.component.scss'
})
export class AddressEditComponent {

	@Input() addressType: "recipient" | "sender" | null = null;
	@Input()
	get adr(): AddressInfoFragment | null {
		return this._adr;
	}
	set adr(value: AddressInfoFragment | null) {
		this._adr = value;
		if (!!this._adr) {
			this.addressForm.patchValue(this._adr)
		}
	}
	private _adr: AddressInfoFragment | null = null;

	addressForm = new FormGroup({
		id: new FormControl<string>('', {nonNullable: true}),
		firstName: new FormControl<string>('', {nonNullable: true}),
		lastName: new FormControl<string>('', {nonNullable: true}),
		phoneNumber: new FormControl<string>('', {nonNullable: true}),
		email: new FormControl<string>('', {nonNullable: true}),
		addressOne: new FormControl<string>('', {nonNullable: true}),
		addressTwo: new FormControl<string>('', {nonNullable: true}),
		zip: new FormControl<string>('', {nonNullable: true}),
		city: new FormControl<string>('', {nonNullable: true}),
		state: new FormControl<string>(''),
		country: new FormGroup({
			id: new FormControl('', {nonNullable: true}),
			label: new FormControl('', {nonNullable: true}),
			alpha2: new FormControl('', {nonNullable: true}),
		}),
		company: new FormControl<string>(''),
	});

	dummyFormCtrl = new FormControl("");
	countryErr = new CountrySelectorErrorStateMatcher();

	consolidationEdit$: Observable<ConsolidationEditModel>;

	constructor(
		private ref: DialogRef,
		private store: Store,
	) {
		this.consolidationEdit$ = store.select(ConsolidationEditState.get);
	}

	save() {

		if (this.addressForm.controls.country.controls.id.value.length === 0) {
			this.dummyFormCtrl.setErrors(["Country must be set to save address"]);
			return;
		}

		if (this.addressType === "sender") {
			this.store.dispatch([new SetSender(this.addressForm.getRawValue())]);
		} else if (this.addressType === "recipient") {
			this.store.dispatch([new SetRecipient(this.addressForm.getRawValue())]);
		} else {
			throw new Error("Address type must be set");
		}
		this.close();
	}

	searchCountries(term: string) {
		this.store.dispatch(new SearchCountries(term));
	}

	changeCountry(country: CountryResult) {
		this.addressForm.patchValue({country: country});
	}

	close() {
		this.ref.close();
	}
}
