import {Component, OnDestroy, OnInit} from '@angular/core';
import {FormControl, FormGroup} from "@angular/forms";
import {Observable, Subscription} from "rxjs";
import {
	DeliveryOptionEditDSVModel,
	DeliveryOptionEditDSVState
} from "./delivery-option-edit-dsv.ngxs";
import {Actions, Store} from "@ngxs/store";
import {ActivatedRoute} from "@angular/router";
import {DeliveryOptionEditDSVActions} from "./delivery-option-edit-dsv.actions";
import Clear = DeliveryOptionEditDSVActions.Clear;
import SetID = DeliveryOptionEditDSVActions.SetID;
import Fetch = DeliveryOptionEditDSVActions.Fetch;
import Save = DeliveryOptionEditDSVActions.Save;

@Component({
	selector: 'app-delivery-option-edit-dsv',
	templateUrl: './delivery-option-edit-dsv.component.html',
	styleUrl: './delivery-option-edit-dsv.component.scss'
})
export class DeliveryOptionEditDsvComponent implements OnInit, OnDestroy {
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
	});

	deliveryOptionsDSVEdit$: Observable<DeliveryOptionEditDSVModel>;
	subscriptions$: Subscription[] = [];

	constructor(
		private store: Store,
		private route: ActivatedRoute,
		private actions$: Actions,
	) {
		this.deliveryOptionsDSVEdit$ = store.select(DeliveryOptionEditDSVState.get);
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
