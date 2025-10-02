import {Component, Input, OnDestroy, OnInit} from '@angular/core';
import {FormArray, FormControl, FormGroup, ReactiveFormsModule} from "@angular/forms";
import {Store} from "@ngxs/store";
import {
	NewDeliveryOptionsRuleDialogComponent
} from "./dialogs/new-delivery-options-rule-dialog.component";
import {
	NewDeliveryOptionsConstraintGroupDialogComponent
} from "./dialogs/new-delivery-options-constraint-group-dialog.component";
import {DeliveryOptionEditRulesActions} from "./delivery-option-edit-rules.actions";
import SetSelectedConstraintID = DeliveryOptionEditRulesActions.SetSelectedConstraintID;
import SetSelectedRule = DeliveryOptionEditRulesActions.SetSelectedRule;
import {MatDialog} from "@angular/material/dialog";
import RemoveCountry = DeliveryOptionEditRulesActions.RemoveCountry;
import SearchCountry = DeliveryOptionEditRulesActions.SearchCountry;
import CountriesResponse = DeliveryOptionEditRulesActions.CountriesResponse;
import AddCountry = DeliveryOptionEditRulesActions.AddCountry;
import {Observable, take} from "rxjs";
import {DeliveryOptionEditRulesModel, DeliveryOptionEditRulesState} from "./delivery-options-edit-rules.ngxs";
import SetDeliveryOptionID = DeliveryOptionEditRulesActions.SetSelectedOption;
import FetchDeliveryOptionEditRulesRuleEdit = DeliveryOptionEditRulesActions.FetchDeliveryOptionEditRulesRuleEdit;
import Clear = DeliveryOptionEditRulesActions.Clear;
import {ConfirmDeleteRuleComponent} from "./dialogs/confirm-delete-rule/confirm-delete-rule.component";
import DeleteRule = DeliveryOptionEditRulesActions.DeleteRule;
import {DeliveryRuleEditPriceComponent} from "./dialogs/delivery-rule-edit-price/delivery-rule-edit-price.component";
import UpdateRulePricing = DeliveryOptionEditRulesActions.UpdateRulePricing;
import {DvoCardComponent} from "../../../../shared/dvo-card/dvo-card.component";
import {AsyncPipe, NgForOf, NgIf} from "@angular/common";
import {MatMenu, MatMenuItem, MatMenuTrigger} from "@angular/material/menu";
import {MatButton, MatIconButton} from "@angular/material/button";
import {MatOption, MatRipple} from "@angular/material/core";
import {MatFormField, MatLabel} from "@angular/material/form-field";
import {MatChipGrid, MatChipInput, MatChipRow} from "@angular/material/chips";
import {MatIcon} from "@angular/material/icon";
import {MatAutocomplete, MatAutocompleteTrigger} from "@angular/material/autocomplete";
import {NgxsFormPluginModule} from "@ngxs/form-plugin";
import {NgxsFormErrorsPluginModule} from "../../../../plugins/ngxs-form-errors/ngxs-form-errors.module";

@Component({
	standalone: true,
	selector: 'app-delivery-option-edit-rules',
	templateUrl: './delivery-option-edit-rules.component.html',
	imports: [
		DvoCardComponent,
		ReactiveFormsModule,
		AsyncPipe,
		MatMenuTrigger,
		MatIconButton,
		MatMenuItem,
		MatButton,
		MatRipple,
		NgForOf,
		NgIf,
		MatFormField,
		MatChipGrid,
		MatChipRow,
		MatIcon,
		MatAutocompleteTrigger,
		MatChipInput,
		MatAutocomplete,
		MatOption,
		MatMenu,
		NgxsFormPluginModule,
		NgxsFormErrorsPluginModule,
		MatLabel,
	],
	styleUrls: ['./delivery-option-edit-rules.component.scss']
})
export class DeliveryOptionEditRulesComponent implements OnInit, OnDestroy {

	deliveryOptionEditRules$: Observable<DeliveryOptionEditRulesModel>;

	editRules = new FormGroup({
		ruleGroups: new FormArray<FormGroup<{id: FormControl, name: FormControl}>>([]),
	});

	@Input()
	get deliveryOptionID(): string {
		return this._deliveryOptionID;
	}
	set deliveryOptionID(id: string) {
		this._deliveryOptionID = id;
		this.store.dispatch([
			new SetDeliveryOptionID(id),
			new FetchDeliveryOptionEditRulesRuleEdit({selectedIndex: 0}),
		]);
	}
	_deliveryOptionID: string = '';

	constructor(
		private store: Store,
		private dialog: MatDialog
	) {
		this.deliveryOptionEditRules$ = store.select(DeliveryOptionEditRulesState.get);
	}

	ngOnInit() {

	}

	newRule() {
		this.dialog.open(NewDeliveryOptionsRuleDialogComponent);
	}

	groupEditor(id: string, index: number) {
		this.store.dispatch(new SetSelectedConstraintID({id, index}));
		this.dialog.open(NewDeliveryOptionsConstraintGroupDialogComponent);
	}

	selectRule(ruleID: string, index: number) {
		this.store.dispatch(new SetSelectedRule({ruleIndex: index, ruleID: ruleID}));
	}

	removeCountry(id: string) {
		this.store.dispatch(new RemoveCountry({id}));
	}

	searchCountries(term: string) {
		this.store.dispatch(new SearchCountry(term));
	}

	addCountry(country: CountriesResponse) {
		this.store.dispatch(new AddCountry({country: country}));
	}

	ngOnDestroy(): void {
		this.store.dispatch(new Clear());
	}

	confirmDeleteRule(name: string, id: string) {
		const ref = this.dialog.open(ConfirmDeleteRuleComponent)
		ref.componentInstance.name = name;
		ref.componentInstance.confirmed.pipe(take(1))
			.subscribe(() => {
				this.store.dispatch(new DeleteRule(id));
			});
	}

	editPrice(name: string, currentPrice: number, currentCurrencyID: string) {
		const ref = this.dialog.open(DeliveryRuleEditPriceComponent);
		ref.componentInstance.name = name;
		ref.componentInstance.price = currentPrice;
		ref.componentInstance.currencyID = currentCurrencyID;
		ref.componentInstance.out.subscribe((p) => {
			this.store.dispatch(new UpdateRulePricing({price: p.price, currency: p.currency}));
		});

	}

}
