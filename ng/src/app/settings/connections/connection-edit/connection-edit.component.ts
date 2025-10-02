import {Component, OnDestroy, OnInit} from '@angular/core';
import {Observable, Subscription} from "rxjs";
import {Actions, ofActionCompleted, Store} from "@ngxs/store";
import {ConnectionEditModel, ConnectionEditState} from "./connection-edit.ngxs";
import {FormControl, FormGroup} from "@angular/forms";
import {ConnectionEditActions} from "./connection-edit.actions";
import SaveForm = ConnectionEditActions.SaveForm;
import {ActivatedRoute} from "@angular/router";
import FetchConnectionEdit = ConnectionEditActions.FetchConnectionEdit;
import {UpdateFormErrors} from "@ngxs/form-plugin";
import FetchConnectionBrands = ConnectionEditActions.FetchConnectionBrands;
import SaveFormUpdate = ConnectionEditActions.SaveFormUpdate;
import SetConnectionID = ConnectionEditActions.SetConnectionID;
import {SelectedLocations} from "./locations-selector/locations-selector.component";
import ConnectionsEditResponse = ConnectionEditActions.SelectConnectionsEditQueryResponse;
import UpdateLocations = ConnectionEditActions.UpdateLocations;
import {
	CreateConnectionInput,
	UpdateConnectionInput, UpdateConnectionShopifyInput
} from "../../../../generated/graphql";
import Clear = ConnectionEditActions.Clear;
import {Paths} from "../../../app-routing.module";
import {MatChipInputEvent} from "@angular/material/chips";
import {COMMA, ENTER} from "@angular/cdk/keycodes";

@Component({
	selector: 'app-connection-edit',
	templateUrl: './connection-edit.component.html',
	styleUrls: ['./connection-edit.component.scss']
})
export class ConnectionEditComponent implements OnInit, OnDestroy {

	connectionEdit$: Observable<ConnectionEditModel>;
	separatorKeysCodes: number[] = [ENTER, COMMA];

	editForm = new FormGroup({
		name: new FormControl('', {nonNullable: true}),
		syncOrders: new FormControl(true, {nonNullable: true}),
		syncProducts: new FormControl(true, {nonNullable: true}),
		fulfillAutomatically: new FormControl(false, {nonNullable: true}),
		dispatchAutomatically: new FormControl(false, {nonNullable: true}),
		autoPrintParcelSlip: new FormControl(false, {nonNullable: true}),
		convertCurrency: new FormControl(false, {nonNullable: true}),
		currency: new FormGroup({
			id: new FormControl('', {nonNullable: true}),
			display: new FormControl('', {nonNullable: true}),
		}),
		pickupLocation: new FormControl({id: ''}, {nonNullable: true}),
		sellerLocation: new FormControl({id: ''}, {nonNullable: true}),
		senderLocation: new FormControl({id: ''}, {nonNullable: true}),
		returnLocation: new FormControl({id: ''}, {nonNullable: true}),
		connectionShopify: new FormGroup({
			storeURL: new FormControl('', {nonNullable: true}),
			apiKey: new FormControl('', {nonNullable: true}),
			rateIntegration: new FormControl(false, {nonNullable: true}),
			syncFrom: new FormControl((new Date()).toISOString(), {nonNullable: true}),
			filterTags: new FormControl<string[]>([], {nonNullable: true}),
		}),
		defaultDeliveryOption: new FormGroup({
			id: new FormControl<string | null>(null, {nonNullable: false}),
		}),
		packingSlipTemplate: new FormGroup({
			id: new FormControl<string | null>(null, {nonNullable: false}),
		})
	});

	subscriptions$: Subscription[] = [];

	constructor(private store: Store,
	            private route: ActivatedRoute,
	            private actions$: Actions) {
		this.connectionEdit$ = store.select(ConnectionEditState.get);
	}

	ngOnInit(): void {
		this.route.queryParams
			.subscribe((params) => {
				this.store.dispatch([
					new SetConnectionID(!!params.id ? params.id : ''),
					new FetchConnectionEdit(),
				]);
			});

		this.store.dispatch(new FetchConnectionBrands());
		this.actions$.pipe(ofActionCompleted(UpdateFormErrors))
			.subscribe((payload) => {
				Object.keys(payload.action.payload.errors || {}).forEach((e) => {
					const form = this.editForm.get(e);
					if (!!form) {
						form.setErrors({form: payload.action.payload.errors![e]});
						form.markAsTouched();
					}
				});
			});
	}

	ngOnDestroy(): void {
		this.subscriptions$.forEach((s) => s.unsubscribe());
		this.store.dispatch(new Clear());
	}

	onSubmit() {
		const id = this.store.selectSnapshot(ConnectionEditState.get).connectionShopifyID;
		const form = this.editForm.getRawValue();

		const input: UpdateConnectionShopifyInput = {
			apiKey: form.connectionShopify.apiKey,
			storeURL: form.connectionShopify.storeURL,
			rateIntegration: form.connectionShopify.rateIntegration,
			syncFrom: form.connectionShopify.syncFrom,
			filterTags: form.connectionShopify.filterTags,
		};
		const connection: CreateConnectionInput = {
			name: this.editForm.getRawValue().name,
			autoPrintParcelSlip: form.autoPrintParcelSlip,
			sellerLocationID: form.sellerLocation.id,
			senderLocationID: form.senderLocation.id,
			pickupLocationID: form.pickupLocation.id,
			returnLocationID: form.returnLocation.id,
			syncProducts: form.syncProducts,
			syncOrders: form.syncOrders,
			convertCurrency: form.convertCurrency,
			currencyID: form.currency.id,
			fulfillAutomatically: form.fulfillAutomatically,
			dispatchAutomatically: form.dispatchAutomatically,
			defaultDeliveryOptionID: form.defaultDeliveryOption.id,
			packingSlipTemplateID: form.packingSlipTemplate.id,
		}

		if (id.length === 0) {
			this.store.dispatch(new SaveForm({input: input, inputConnection: connection}));
		} else {

			const connectionUpdate: UpdateConnectionInput = Object.assign(
				connection, {clearDefaultDeliveryOption: true});

			this.store.dispatch(new SaveFormUpdate({id, input: input, inputConnection: connectionUpdate}));
		}
	}

	toSelectedLocations(connection: ConnectionsEditResponse): SelectedLocations {
		return {
			sellerID: connection?.sellerLocation?.id,
			senderID: connection?.senderLocation?.id,
			pickupID: connection?.pickupLocation?.id,
			returnID: connection?.returnLocation?.id,
		}
	}

	updateLocations(val: SelectedLocations) {
		this.store.dispatch(new UpdateLocations(val));
	}

	addTag(event: MatChipInputEvent): void {
		const value = (event.value || '').trim();
		if (value) {
			this.store.dispatch(new ConnectionEditActions.AddFilterTag(value));
		}
		event.chipInput!.clear();
	}

	removeTag(tag: string): void {
		this.store.dispatch(new ConnectionEditActions.RemoveFilterTag(tag));
	}

	protected readonly Paths = Paths;
}
