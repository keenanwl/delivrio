import {Component, OnDestroy, OnInit} from '@angular/core';
import {FormControl, FormGroup, FormsModule, ReactiveFormsModule} from "@angular/forms";
import {NgxsFormPluginModule} from "@ngxs/form-plugin";
import {MatAutocomplete, MatAutocompleteTrigger, MatOption} from "@angular/material/autocomplete";
import {MatChipGrid, MatChipInput, MatChipRow} from "@angular/material/chips";
import {MatError, MatFormField, MatLabel} from "@angular/material/form-field";
import {AsyncPipe, NgForOf, NgIf} from "@angular/common";
import {Observable, Subscription} from "rxjs";
import {ProductModel, ProductState} from "../../product-edit.ngxs";
import {Actions, ofActionDispatched, Store} from "@ngxs/store";
import {ActivatedRoute} from "@angular/router";
import {MatDialog, MatDialogRef} from "@angular/material/dialog";
import {ProductActions} from "../../product-edit.actions";
import SearchCountries = ProductActions.SearchCountries;
import CountriesResponse = ProductActions.CountriesResponse;
import ChangeCountry = ProductActions.ChangeCountry;
import {MatInput} from "@angular/material/input";
import {MatIconModule} from "@angular/material/icon";
import {MatButtonModule} from "@angular/material/button";
import ResetInventoryForm = ProductActions.ResetInventoryForm;
import SaveInventoryItem = ProductActions.SaveInventoryItem;
import CloseInventoryForm = ProductActions.CloseInventoryForm;

@Component({
	selector: 'app-edit-inventory',
	standalone: true,
	imports: [
		FormsModule,
		NgxsFormPluginModule,
		ReactiveFormsModule,
		MatAutocomplete,
		MatAutocompleteTrigger,
		MatChipGrid,
		MatChipInput,
		MatChipRow,
		MatFormField,
		MatLabel,
		MatOption,
		NgForOf,
		AsyncPipe,
		NgIf,
		MatError,
		MatInput,
		MatIconModule,
		MatButtonModule
	],
	templateUrl: './edit-inventory.component.html',
	styleUrl: './edit-inventory.component.scss'
})
export class EditInventoryComponent implements OnInit, OnDestroy {

	product$: Observable<ProductModel>;
	editForm = new FormGroup({
		id: new FormControl<string>('', {nonNullable: true}),
		code: new FormControl<string>('', {nonNullable: true}),
		sku: new FormControl<string>('', {nonNullable: true}),
		countryOfOrigin: new FormGroup({
			id: new FormControl(''),
			label: new FormControl(''),
			alpha2: new FormControl(''),
		})
	});

	subscriptions$: Subscription[] = [];

	constructor(private store: Store,
				private route: ActivatedRoute,
				private ref: MatDialogRef<any>,
				private actions$: Actions) {
		this.product$ = store.select(ProductState.get);
	}

	ngOnDestroy(): void {
        this.store.dispatch(new ResetInventoryForm());
		this.subscriptions$.forEach((s) => s.unsubscribe());
    }

	ngOnInit(): void {
		this.subscriptions$.push(this.actions$.pipe(ofActionDispatched(CloseInventoryForm))
			.subscribe(() => {
				this.ref.close();
			}))
	}

	save() {
		this.store.dispatch(new SaveInventoryItem({id: this.editForm.controls.id.value, input: {
			sku: this.editForm.controls.sku.value,
			code: this.editForm.controls.code.value,
			countryOfOriginID: this.editForm.controls.countryOfOrigin.controls.id.value,
		}}))
	}

	close() {
		this.ref.close();
	}

	searchCountries(term: string) {
		this.store.dispatch(new SearchCountries(term));
	}

	changeCountry(country: CountriesResponse) {
		this.store.dispatch(new ChangeCountry(country));
	}
}
