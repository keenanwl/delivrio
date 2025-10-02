import {Component, OnDestroy, OnInit} from '@angular/core';
import {FormControl, FormGroup} from "@angular/forms";
import {Observable, Subscription} from "rxjs";
import {
	DeliveryOptionEditDFModel,
	DeliveryOptionEditDFState
} from "./delivery-option-edit-df.ngxs";
import {Store} from "@ngxs/store";
import {ActivatedRoute} from "@angular/router";
import {DeliveryOptionEditDFActions} from "./delivery-option-edit-df.actions";
import Clear = DeliveryOptionEditDFActions.Clear;
import SetID = DeliveryOptionEditDFActions.SetID;
import Fetch = DeliveryOptionEditDFActions.Fetch;
import Save = DeliveryOptionEditDFActions.Save;
import {BaseDeliveryOptionFragment} from "../edit-common.generated";
import SetEditDF = DeliveryOptionEditDFActions.SetEditDF;

@Component({
	selector: 'app-delivery-option-edit-df',
	templateUrl: './delivery-option-edit-df.component.html',
	styleUrl: './delivery-option-edit-df.component.scss'
})
export class DeliveryOptionEditDfComponent implements OnInit, OnDestroy {

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

	deliveryOptionsDFEdit$: Observable<DeliveryOptionEditDFModel>;
	subscriptions$: Subscription[] = [];

	constructor(
		private store: Store,
		private route: ActivatedRoute,
	) {
		this.deliveryOptionsDFEdit$ = store.select(DeliveryOptionEditDFState.get);
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

	baseFormChange(val: BaseDeliveryOptionFragment) {
		this.store.dispatch(new SetEditDF(val));
	}

	save() {
		this.store.dispatch(new Save());
	}

}
