import {Component, OnDestroy, OnInit} from '@angular/core';
import {FormControl, FormGroup} from "@angular/forms";
import {Observable, Subscription} from "rxjs";
import {Actions, Store} from "@ngxs/store";
import {ActivatedRoute} from "@angular/router";
import {DeliveryOptionEditDAOActions} from "./delivery-option-edit-dao.actions";
import SetID = DeliveryOptionEditDAOActions.SetID;
import Fetch = DeliveryOptionEditDAOActions.Fetch;
import Clear = DeliveryOptionEditDAOActions.Clear;
import Save = DeliveryOptionEditDAOActions.Save;
import {DeliveryOptionEditDAOModel, DeliveryOptionEditDAOState} from "./delivery-option-edit-dao.ngxs";

@Component({
	selector: 'app-delivery-option-edit-dao',
	templateUrl: './delivery-option-edit-dao.component.html',
	styleUrl: './delivery-option-edit-dao.component.scss'
})
export class DeliveryOptionEditDaoComponent implements OnInit, OnDestroy {

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

	deliveryOptionsDAOEdit$: Observable<DeliveryOptionEditDAOModel>;
	subscriptions$: Subscription[] = [];

	constructor(
		private store: Store,
		private route: ActivatedRoute,
		private actions$: Actions,
	) {
		this.deliveryOptionsDAOEdit$ = store.select(DeliveryOptionEditDAOState.get);
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

