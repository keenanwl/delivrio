import {MatDialog, MatDialogRef} from "@angular/material/dialog";
import {ChangeDetectorRef, Component, OnDestroy, OnInit} from "@angular/core";
import {Actions, ofActionCompleted, Store} from "@ngxs/store";
import {FormArray, FormControl, FormGroup} from "@angular/forms";
import {Observable} from "rxjs";
import {
	DeliveryRuleConstraintComparison,
	DeliveryRuleConstraintGroupConstraintLogic,
	DeliveryRuleConstraintPropertyType,
	DeliveryRuleConstraintSelectedValueInput
} from "src/generated/graphql";
import FetchRuleConstrains = DeliveryOptionEditRulesActions.FetchRuleConstrains;
import FetchRuleOptions = DeliveryOptionEditRulesActions.FetchRuleOptions;
import SaveEditConstraintGroup = DeliveryOptionEditRulesActions.SaveEditConstraintGroup;
import SaveNewConstraintGroup = DeliveryOptionEditRulesActions.SaveNewConstraintGroup;
import DeleteConstraintGroup = DeliveryOptionEditRulesActions.DeleteConstraintGroup;
import {TimeRangeSelectorComponent} from "../../../../../shared/time-range-selector/time-range-selector.component";
import UpdateDayOfWeek = DeliveryOptionEditRulesActions.UpdateDayOfWeek;
import UpdateConstraintProperty = DeliveryOptionEditRulesActions.UpdateConstraintProperty;
import UpdateConstraintComparison = DeliveryOptionEditRulesActions.UpdateConstraintComparison;
import UpdateTimeOfDay = DeliveryOptionEditRulesActions.UpdateTimeOfDay;
import DeleteConstraintLine = DeliveryOptionEditRulesActions.DeleteConstraintLine;
import AddConstraintLine = DeliveryOptionEditRulesActions.AddConstraintLine;
import FetchRuleTagResponse = DeliveryOptionEditRulesActions.FetchRuleTagResponse;
import {COMMA, ENTER} from "@angular/cdk/keycodes";
import {MatAutocompleteSelectedEvent} from "@angular/material/autocomplete";
import SetProductTagSearchTerm = DeliveryOptionEditRulesActions.SetProductTagSearchTerm;
import FetchProductTags = DeliveryOptionEditRulesActions.FetchProductTags;
import {debounceTime} from "rxjs/operators";
import AddProductTag = DeliveryOptionEditRulesActions.AddProductTag;
import RemoveProductTag = DeliveryOptionEditRulesActions.RemoveProductTag;
import UpdateNumericValue = DeliveryOptionEditRulesActions.UpdateNumericValue;
import AddZipCode = DeliveryOptionEditRulesActions.AddPostalCode;
import RemovePostalCode = DeliveryOptionEditRulesActions.RemovePostalCode;
import {DeliveryOptionEditRulesModel, DeliveryOptionEditRulesState} from "../delivery-options-edit-rules.ngxs";
import {DeliveryOptionEditRulesActions} from "../delivery-option-edit-rules.actions";
import ClearRulesDialog = DeliveryOptionEditRulesActions.ClearRulesDialog;
import RemovePostalCodeString = DeliveryOptionEditRulesActions.RemovePostalCodeString;
import AddPostalCodeString = DeliveryOptionEditRulesActions.AddPostalCodeString;
import UpdateConstraintLogicType = DeliveryOptionEditRulesActions.UpdateConstraintLogicType;

enum ConstraintDataType {
	NUMERIC= "NUMERIC",
	NUMERIC_RANGE = "NUMERIC_RANGE",
	IDS = "IDS",
	DAY_OF_WEEK = "DAY_OF_WEEK",
	TIME_OF_DAY = "TIME_OF_DAY",
	TEXT = "TEXT",
	VALUES = "VALUES",
}

@Component({
	selector: 'app-new-delivery-options-constraint-group-dialog',
	styleUrls: ['new-delivery-options-constraint-group-dialog.component.scss'],
	templateUrl: 'new-delivery-options-constraint-group-dialog.component.html',
})
export class NewDeliveryOptionsConstraintGroupDialogComponent implements OnInit, OnDestroy {

/*	constraintEditForm = new FormGroup({
		constraintLogic: new FormControl<DeliveryRuleConstraintGroupConstraintLogic>(DeliveryRuleConstraintGroupConstraintLogic.And, {nonNullable: true}),
		constraints:
			new FormArray<FormGroup<{
				label: FormControl<string | null>,
				comparison: FormControl<string | null>,
				propertyType: FormControl<DeliveryRuleConstraintPropertyType | null>,
				values: FormControl<Array<string | null> | null>,
				value: FormControl<string | null>,
				dayOfWeek: FormControl<Array<string | null> | null>,
				timeOfDay: FormControl<Array<string | null> | null>,
				ids: FormControl<Array<string | null> | null>,
				tags: FormControl<Array<FetchRuleTagResponse | null> | null>,
				numericRange: FormControl<Array<number | null> | null>,
			}>>([])
	});*/

	trackBy = (index: number, row: FormGroup) => {
		return index;
	}
	productTagSearch = new FormControl("");

	deliveryOptionEditRules$: Observable<DeliveryOptionEditRulesModel>;
	separatorKeysCodes: number[] = [ENTER, COMMA];

	constructor(
		private store: Store,
		private dialogRef: MatDialogRef<NewDeliveryOptionsConstraintGroupDialogComponent>,
		private cd: ChangeDetectorRef,
		private actions$: Actions,
		private dialog: MatDialog,
	) {
		this.deliveryOptionEditRules$ = store.select(DeliveryOptionEditRulesState.get);
	}

	ngOnInit() {

		this.productTagSearch.valueChanges
			.pipe(debounceTime(200))
			.subscribe((v) => {
				if (typeof v === "string") {
					this.store.dispatch([
						new SetProductTagSearchTerm(v),
						new FetchProductTags(),
					])
				}
			});

		this.store.dispatch([new FetchRuleConstrains(), new FetchRuleOptions()]);

		/*this.constraintEditForm.controls.constraintLogic.valueChanges
			.subscribe((val) => {
				this.store.dispatch(new UpdateConstraintLogicType(val));
			});

		this.actions$.pipe(ofActionCompleted(DeliveryOptionEditRulesActions.SetConstraintLogicType))
			.subscribe(() => {
				const state = this.store.selectSnapshot(DeliveryOptionEditRulesState.get);
				this.constraintEditForm.controls.constraintLogic.setValue(state.constraintLogicType);
			});

		this.actions$.pipe(ofActionCompleted(DeliveryOptionEditRulesActions.SetRuleConstraintsNotify))
			.subscribe(() => {
				const state = this.store.selectSnapshot(DeliveryOptionEditRulesState.get);

				this.constraintEditForm.controls.constraints.clear();
				state.constraints?.forEach((c) => {
					this.constraintEditForm.controls.constraints.push(
						this.newConstraintLine(
							c.constraint?.selectedValue.dayOfWeek,
							c.constraint?.selectedValue.timeOfDay,
							c.constraint?.selectedValue.ids,
							c.constraint?.selectedValue.numericRange,
							c.constraint?.selectedValue.values,
							c.tags,
							this.valueFromBackendValue(c.constraint!.propertyType, c.constraint!.selectedValue),
							c.constraint?.comparison,
							c.constraint?.propertyType,
							this.isTimeProperty(c.constraint!.propertyType),
							c.constraint!.propertyType,
						),
					);
				});
				//this.cd.detectChanges();
			});*/
	}

	ngOnDestroy() {
		this.store.dispatch(new ClearRulesDialog());
	}

	valueFromBackendValue(
	    propertyType: DeliveryRuleConstraintPropertyType,
	    value: DeliveryRuleConstraintSelectedValueInput
	): string {

		switch (this.constraintDataType(propertyType)) {
			case ConstraintDataType.NUMERIC:
				return value.numeric + '';
			case ConstraintDataType.NUMERIC_RANGE:
				return value.numericRange?.join(" - ") || '';
			case ConstraintDataType.VALUES:
				return value.values?.join(", ") || '';
			case ConstraintDataType.TEXT:
				return value.text || '';
			case ConstraintDataType.DAY_OF_WEEK:
				return value.dayOfWeek?.join(', ') || '';
			case ConstraintDataType.IDS:
				return value.ids?.join(', ') || '';
			case ConstraintDataType.TIME_OF_DAY:
				return value.timeOfDay?.join(' - ') || '';
		}

		return 'error';
	}

	selected(name: string) {
		this.dialogRef.close();
	}

	newConstraintLine(
		dayOfWeek: Array<string | null> | null = null,
		timeOfDay: Array<string | null> | null = null,
		ids: Array<string | null> | null = null,
		numericRange: Array<number | null> | null = null,
		values: Array<string | null> | null = null,
		tags: Array<FetchRuleTagResponse | null> | null | undefined = null,
		value: string | null = null,
	    comparison: string | null = DeliveryRuleConstraintComparison.Equals,
	    label: string | null = null,
		disabled: boolean = false,
		propertyType: DeliveryRuleConstraintPropertyType,
	) {
		const valForm = new FormControl(value);
		if (disabled) {
			valForm.disable();
		}

		return new FormGroup({
			label: new FormControl(label),
			comparison: new FormControl(comparison),
			values: new FormControl(values),
			value: valForm,
			dayOfWeek: new FormControl(dayOfWeek),
			timeOfDay: new FormControl(timeOfDay),
			ids: new FormControl(ids),
			tags: new FormControl(tags),
			numericRange: new FormControl(numericRange),
			propertyType: new FormControl(propertyType),
		});
	}

	constraintDataType(propertyType: DeliveryRuleConstraintPropertyType): ConstraintDataType {
		switch (propertyType) {
			case DeliveryRuleConstraintPropertyType.CartTotal:
			case DeliveryRuleConstraintPropertyType.OrderLines:
			case DeliveryRuleConstraintPropertyType.TotalWeight:
				return ConstraintDataType.NUMERIC;
			case DeliveryRuleConstraintPropertyType.PostalCodeNumeric:
				return ConstraintDataType.NUMERIC_RANGE;
			case DeliveryRuleConstraintPropertyType.Sku:
			case DeliveryRuleConstraintPropertyType.PostalCodeString:
				return ConstraintDataType.VALUES;
			case DeliveryRuleConstraintPropertyType.DayOfWeek:
				return ConstraintDataType.DAY_OF_WEEK;
			case DeliveryRuleConstraintPropertyType.ProductTag:
			case DeliveryRuleConstraintPropertyType.AllProductsTagged:
				return ConstraintDataType.IDS;
			case DeliveryRuleConstraintPropertyType.TimeOfDay:
				return ConstraintDataType.TIME_OF_DAY;
		}
		return ConstraintDataType.TEXT;
	}

	addConstraint() {
		this.store.dispatch(new AddConstraintLine());
	}

	removeLine(index: number) {
		this.store.dispatch(new DeleteConstraintLine({index}));
	}

	isStringProperty(propertyType: DeliveryRuleConstraintPropertyType): boolean {
		return this.constraintDataType(propertyType) === ConstraintDataType.TEXT;
	}

	isNumericProperty(propertyType: DeliveryRuleConstraintPropertyType): boolean {
		return this.constraintDataType(propertyType) === ConstraintDataType.NUMERIC;
	}

	isNumericRangeProperty(propertyType: DeliveryRuleConstraintPropertyType): boolean {
		return this.constraintDataType(propertyType) === ConstraintDataType.NUMERIC_RANGE;
	}

	isValuesProperty(propertyType: DeliveryRuleConstraintPropertyType): boolean {
		return this.constraintDataType(propertyType) === ConstraintDataType.VALUES;
	}

	isWeekdayProperty(propertyType: DeliveryRuleConstraintPropertyType): boolean {
		return this.constraintDataType(propertyType) === ConstraintDataType.DAY_OF_WEEK;
	}

	isTimeProperty(propertyType: DeliveryRuleConstraintPropertyType): boolean {
		return this.constraintDataType(propertyType) === ConstraintDataType.TIME_OF_DAY;
	}

	isIDsProperty(propertyType: DeliveryRuleConstraintPropertyType): boolean {
		return this.constraintDataType(propertyType) === ConstraintDataType.IDS;
	}

	saveNew() {
		this.store.dispatch(new SaveNewConstraintGroup());
		this.close();
	}

	saveEdit() {
		this.store.dispatch(new SaveEditConstraintGroup());
		this.close();
	}

	close() {
		this.dialogRef.close();
	}

	delete() {
		const state = this.store.selectSnapshot(DeliveryOptionEditRulesState.get);
		this.store.dispatch(new DeleteConstraintGroup(state.selectedConstraintGroup));
		this.close();
	}

	selectTimeRange(index: number) {
		const ref = this.dialog.open(TimeRangeSelectorComponent);
		ref.componentInstance.selectedRange.subscribe((range) => {
			this.store.dispatch(new UpdateTimeOfDay({index, timeOfDay: [range.start, range.end]}));
		});
	}

	updateProperty(index: number, propertyType: DeliveryRuleConstraintPropertyType) {
		this.store.dispatch(new UpdateConstraintProperty({index, propertyType}));
	}

	updateComparison(index: number, comparison: DeliveryRuleConstraintComparison) {
		this.store.dispatch(new UpdateConstraintComparison({index, comparison}));
	}

	updateDaysOfWeek(index: number, dayOfWeek: string[]) {
		this.store.dispatch(new UpdateDayOfWeek({index, dayOfWeek}));
	}

	removeTag(index: number, tag: FetchRuleTagResponse) {
		this.store.dispatch(new RemoveProductTag({index, tag}))
	}

	removePostalCode(index: number, code: number) {
		this.store.dispatch(new RemovePostalCode({index, code}));
	}

	removePostalCodeString(index: number, code: string) {
		this.store.dispatch(new RemovePostalCodeString({index, code}));
	}

	addZip(index: number, code: string, elem: HTMLInputElement) {
		elem.value = '';
		const parsed = parseInt(code);
		if (!!parsed) {
			this.store.dispatch(new AddZipCode({index, code: parsed}));
		}
	}

	addPostalCodeBlur(index: number, code: FocusEvent, elem: HTMLInputElement) {
		const input = code.target as HTMLInputElement;
		const value = input.value;
		elem.value = '';

		const parsed = parseInt(value);
		if (!!parsed) {
			this.store.dispatch(new AddZipCode({index, code: parsed}));
		}
	}

	addPostalCodeString(index: number, code: string, elem: HTMLInputElement) {
		elem.value = '';
		this.store.dispatch(new AddPostalCodeString({index, code: code.trim()}));
	}

	addPostalCodeBlurString(index: number, code: FocusEvent, elem: HTMLInputElement) {
		const input = code.target as HTMLInputElement;
		const value = input.value;
		elem.value = '';

		if ((value || '').trim().length > 0) {
			this.store.dispatch(new AddPostalCodeString({index, code: value.trim()}));
		}
	}

	selectTag(index: number, event: MatAutocompleteSelectedEvent, elem: HTMLInputElement): void {
		elem.value = '';
		this.store.dispatch(new AddProductTag({index, tag: event.option.value}))
	}

	updateNumericValue(index: number, value: string | null): void {
		const parsed = parseInt(value || "0", 10)
		if (!!parsed) {
			this.store.dispatch(new UpdateNumericValue({index, value: parsed}));
		}
	}

	updateConstraintLogic(val: DeliveryRuleConstraintGroupConstraintLogic): void {
		this.store.dispatch(new UpdateConstraintLogicType(val));
	}

	protected readonly DeliveryRuleConstraintPropertyType = DeliveryRuleConstraintPropertyType;
	protected readonly DeliveryRuleConstraintComparison = DeliveryRuleConstraintComparison;
	protected readonly DeliveryRuleConstraintGroupConstraintLogic = DeliveryRuleConstraintGroupConstraintLogic;
}
