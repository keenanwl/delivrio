import {Component, Input, output} from '@angular/core';
import {MatError, MatFormField, MatHint, MatLabel} from "@angular/material/form-field";
import {MatInput} from "@angular/material/input";
import {AsyncPipe, NgForOf, NgIf} from "@angular/common";
import {ReactiveFormsModule} from "@angular/forms";
import {DialogRef} from "@angular/cdk/dialog";
import {MatButton} from "@angular/material/button";
import {MatIcon} from "@angular/material/icon";
import {Observable} from "rxjs";
import {DeliveryOptionEditRulesModel, DeliveryOptionEditRulesState} from "../../delivery-options-edit-rules.ngxs";
import {MatOption} from "@angular/material/autocomplete";
import {MatSelect} from "@angular/material/select";
import {Store} from "@ngxs/store";
import {DeliveryOptionEditRulesActions} from "../../delivery-option-edit-rules.actions";
import CurrencyResponse = DeliveryOptionEditRulesActions.CurrencyResponse;
import {Paths} from "../../../../../../app-routing.module";

@Component({
	selector: 'app-delivery-rule-edit-price',
	standalone: true,
	imports: [
		MatError,
		MatFormField,
		MatInput,
		MatLabel,
		NgIf,
		ReactiveFormsModule,
		MatButton,
		MatIcon,
		AsyncPipe,
		MatOption,
		MatSelect,
		MatHint,
		NgForOf,
	],
	templateUrl: './delivery-rule-edit-price.component.html',
	styleUrl: './delivery-rule-edit-price.component.scss'
})
export class DeliveryRuleEditPriceComponent {
	@Input() name = "";
	@Input() price = 0.00;
	@Input() currencyID = "";

	out = output<{price: number; currency: CurrencyResponse}>();
	deliveryOptionEditRules$: Observable<DeliveryOptionEditRulesModel>;

	constructor(private ref: DialogRef, private store: Store) {
		this.deliveryOptionEditRules$ = store.select(DeliveryOptionEditRulesState.get);
	}

	save(price: string, currency: string, currencies: CurrencyResponse[]) {

		currencies.some((c) => {
			if (c.id == currency) {
				this.out.emit({price: parseFloat(price) || 0, currency: c});
				return true
			}
			return false;
		})

		this.close();
	}

	close() {
		this.ref.close();
	}

	protected readonly Paths = Paths;
}
