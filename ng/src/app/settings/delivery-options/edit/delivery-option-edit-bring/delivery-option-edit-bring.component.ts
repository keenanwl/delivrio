import {Component, OnDestroy, OnInit} from '@angular/core';
import {FormControl, FormGroup} from "@angular/forms";
import {Observable, Subscription} from "rxjs";
import {Actions, Store} from "@ngxs/store";
import {ActivatedRoute} from "@angular/router";
import {DeliveryOptionEditBringModel, DeliveryOptionEditBringState} from "./delivery-option-edit-bring.ngxs";
import {DeliveryOptionEditBringActions} from "./delivery-option-edit-bring.actions";
import Clear = DeliveryOptionEditBringActions.Clear;
import Save = DeliveryOptionEditBringActions.Save;
import Fetch = DeliveryOptionEditBringActions.Fetch;
import SetID = DeliveryOptionEditBringActions.SetID;

@Component({
	selector: 'app-delivery-option-edit-bring',
	templateUrl: './delivery-option-edit-bring.component.html',
	styleUrl: './delivery-option-edit-bring.component.scss'
})
export class DeliveryOptionEditBringComponent implements OnInit, OnDestroy {

	editForm = new FormGroup({
		clickCollect: new FormControl(true),
		overrideReturnAddress: new FormControl(true),
		overrideSenderAddress: new FormControl(true),
		hideDeliveryOption: new FormControl(true),
		name: new FormControl(''),
		paperless: new FormControl(true),
		senderPaysTax: new FormControl(true),
		requireSignature: new FormControl(true),
		carrierService: new FormGroup({
			id: new FormControl(""),
		}),
		deliveryOptionBring: new FormGroup({
			electronicCustoms: new FormControl(false),
		})
	});

	deliveryOptionsBringEdit$: Observable<DeliveryOptionEditBringModel>;
	subscriptions$: Subscription[] = [];

	constructor(
		private store: Store,
		private route: ActivatedRoute,
		private actions$: Actions,
	) {
		this.deliveryOptionsBringEdit$ = store.select(DeliveryOptionEditBringState.get);
	}

	ngOnDestroy(): void {
		this.subscriptions$.forEach((s) => s.unsubscribe());
		this.store.dispatch(new Clear());
	}

	ngOnInit(): void {

		this.subscriptions$.push(this.route.queryParams
			.subscribe((params) => {
				this.store.dispatch([
					new SetID(!!params.id ? params.id : ''),
					new Fetch(),
				]);
			}));
	}

	save() {
		this.store.dispatch(new Save());
	}

}
