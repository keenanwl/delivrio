import {Component, OnDestroy, OnInit} from '@angular/core';
import {Observable, Subscription} from "rxjs";
import {ActivatedRoute} from "@angular/router";
import {Store} from "@ngxs/store";
import {FormControl, FormGroup} from "@angular/forms";
import {CarrierEditUSPSModel, CarrierEditUSPSState} from "./carrier-edit-usps.ngxs";
import {CarrierEditUSPSActions} from "./carrier-edit-usps.actions";
import FetchCarrierUSPSEdit = CarrierEditUSPSActions.FetchCarrierUSPSEdit;
import SetID = CarrierEditUSPSActions.SetID;
import Clear = CarrierEditUSPSActions.Clear;
import SaveForm = CarrierEditUSPSActions.SaveForm;

@Component({
	selector: 'app-carrier-edit-usps',
	templateUrl: './carrier-edit-usps.component.html',
	styleUrls: ['./carrier-edit-usps.component.scss']
})
export class CarrierEditUspsComponent implements OnInit, OnDestroy {
	carrierEditUSPS$: Observable<CarrierEditUSPSModel>;

	editForm = new FormGroup({
		carrier: new FormGroup({
			name: new FormControl('', {nonNullable: true}),
		}),
		consumerKey: new FormControl('', {nonNullable: true}),
		consumerSecret: new FormControl('', {nonNullable: true}),
		crid: new FormControl('', {nonNullable: true}),
		mid: new FormControl('', {nonNullable: true}),
		manifestMid: new FormControl('', {nonNullable: true}),
		epsAccountNumber: new FormControl('', {nonNullable: true}),
		isTestAPI: new FormControl(false, {nonNullable: true}),
	});

	subscriptions$: Subscription[] = [];

	constructor(
		private route: ActivatedRoute,
		private store: Store
	) {
		this.carrierEditUSPS$ = store.select(CarrierEditUSPSState.get);
	}

	ngOnInit() {
		this.subscriptions$.push(this.route.queryParams
			.subscribe((params) => {
				this.store.dispatch([
					new SetID(!!params.id ? params.id : ''),
					new FetchCarrierUSPSEdit(),
				]);
			}));
	}

	ngOnDestroy(): void {
		this.subscriptions$.forEach((s) => s.unsubscribe());
		this.store.dispatch(new Clear());
	}

	save() {
		this.store.dispatch(new SaveForm());
	}
}
