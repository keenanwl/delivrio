import {Component, EventEmitter, Input, OnDestroy, OnInit, Output} from '@angular/core';
import {FormControl, FormGroup, ReactiveFormsModule} from "@angular/forms";
import {BaseDeliveryOptionFragment, CarrierServiceItemFragment} from "../edit-common.generated";
import {MatDialog} from "@angular/material/dialog";
import {SelectPackagingComponent} from "../../../../shared/select-packaging/select-packaging.component";
import {Subscription} from "rxjs";
import {CustomIntegrationComponent} from "./dialogs/custom-integration/custom-integration.component";
import {
	CustomIntegrationShipmondoComponent
} from "./dialogs/custom-integration-shipmondo/custom-integration-shipmondo.component";
import {MatError, MatFormField, MatLabel} from "@angular/material/form-field";
import {MatInput} from "@angular/material/input";
import {MatOptgroup, MatOption} from "@angular/material/autocomplete";
import {NgForOf, NgIf} from "@angular/common";
import {CarrierServiceGrouperPipe} from "../pipes/carrier-service-grouper.pipe";
import {ToggleContainerComponent} from "../../../../shared/toggle-container/toggle-container.component";
import {MatIconButton} from "@angular/material/button";
import {MatSlideToggle} from "@angular/material/slide-toggle";
import {MatIcon} from "@angular/material/icon";
import {MatSelect} from "@angular/material/select";

@Component({
	selector: 'app-delivery-option-edit-base',
	templateUrl: './delivery-option-edit-base.component.html',
	styleUrl: './delivery-option-edit-base.component.scss',
	standalone: true,
	imports: [
		MatFormField,
		ReactiveFormsModule,
		MatInput,
		MatOptgroup,
		MatOption,
		NgForOf,
		CarrierServiceGrouperPipe,
		NgIf,
		ToggleContainerComponent,
		MatIconButton,
		MatSlideToggle,
		MatIcon,
		MatError,
		MatLabel,
		MatSelect,
	]
})
export class DeliveryOptionEditBaseComponent implements OnInit, OnDestroy {

	@Input() carrierName = "";
	@Input() availableServices: CarrierServiceItemFragment[] = [];
	@Input()
	get formValues(): BaseDeliveryOptionFragment | undefined {
		return this._formValues;
	}
	set formValues(value: BaseDeliveryOptionFragment) {
		this._formValues = value;
		if (!!value) {
			this.form.patchValue(value, {emitEvent: false});
		}

	}
	private _formValues: BaseDeliveryOptionFragment | undefined = undefined;

	@Output() formChange: EventEmitter<BaseDeliveryOptionFragment> = new EventEmitter();

	subscriptions$: Subscription[] = [];

	form = new FormGroup({
		clickCollect: new FormControl(true),
		overrideReturnAddress: new FormControl(true),
		overrideSenderAddress: new FormControl(true),
		hideDeliveryOption: new FormControl(true),
		name: new FormControl('', {nonNullable: true}),
		description: new FormControl(''),
		clickOptionDisplayCount: new FormControl(0),
		deliveryEstimateTo: new FormControl(0),
		deliveryEstimateFrom: new FormControl(0),
		webshipperIntegration: new FormControl(false, {nonNullable: true}),
		webshipperID: new FormControl(0),
		shipmondoIntegration: new FormControl(false, {nonNullable: true}),
		shipmondoDeliveryOption: new FormControl(""),
		customsEnabled: new FormControl(false, {nonNullable: true}),
		customsSigner: new FormControl(""),
		hideIfCompanyEmpty: new FormControl(false, {nonNullable: true}),
		carrierService: new FormGroup({
			id: new FormControl("", {nonNullable: true}),
		}),
		defaultPackaging: new FormGroup<{id: FormControl; name: FormControl} | null>({
			id: new FormControl<string | null>(null),
			name: new FormControl<string | null>(null),
		}), // Hack :(
	});

	constructor(private dialog: MatDialog) {
		this.formChange.emit(this.form.getRawValue());
	}

	ngOnInit() {
		this.form.valueChanges.subscribe(() => {
			this.formChange.emit(this.form.getRawValue());
		});
	}

	ngOnDestroy() {
		this.subscriptions$.forEach(s => s.unsubscribe());
	}

	selectPackaging() {
		const ref = this.dialog.open(SelectPackagingComponent);
		ref.componentInstance.helpText = `If this delivery option is selected and the colli does not have packaging, this packaging will be used.`
		this.subscriptions$.push(ref.componentInstance.selected.subscribe((p) => {
			this.form.patchValue({defaultPackaging: {id: p?.id || null, name: p?.name || null}});
		}));
	}

	enableWebshipper() {
		const ref = this.dialog.open(CustomIntegrationComponent)
		ref.componentInstance.initialVal = this.form.controls.webshipperID.value || 1;
		ref.componentInstance.outputVal.subscribe((o) =>
			this.form.controls.webshipperID.setValue(o));
	}

	enableShipmondo() {
		const ref = this.dialog.open(CustomIntegrationShipmondoComponent)
		ref.componentInstance.initialVal = this.form.controls.shipmondoDeliveryOption.value || '';
		ref.componentInstance.outputVal.subscribe((o) =>
			this.form.controls.shipmondoDeliveryOption.setValue(o));
	}

}
