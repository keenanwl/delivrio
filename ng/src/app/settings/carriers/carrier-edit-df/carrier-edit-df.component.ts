import {Component, OnDestroy, OnInit} from '@angular/core';
import {Observable, Subscription} from "rxjs";
import {CarrierEditDFModel, CarrierEditDFState} from "../carrier-edit-df/carrier-edit-df.ngxs";
import {FormControl, FormGroup} from "@angular/forms";
import {Actions, ofActionCompleted, Store} from "@ngxs/store";
import {ActivatedRoute} from "@angular/router";
import {UpdateFormErrors} from "@ngxs/form-plugin";
import {CarrierEditDFActions} from "./carrier-edit-df.actions";
import FetchCarrierEditDF = CarrierEditDFActions.FetchCarrierEditDF;
import Clear = CarrierEditDFActions.Clear;

@Component({
	selector: 'app-carrier-edit-df',
	templateUrl: './carrier-edit-df.component.html',
	styleUrl: './carrier-edit-df.component.scss'
})
export class CarrierEditDfComponent implements OnInit, OnDestroy {
	state$: Observable<CarrierEditDFModel>;

	form = new FormGroup({
		id: new FormControl(''),
		name: new FormControl(''),
		carrierDF: new FormGroup({
			customerID: new FormControl(''),
			agreementNumber: new FormControl(''),
		}),
	});

	subscriptions$: Subscription[] = [];

	constructor(private store: Store,
				private route: ActivatedRoute,
				private actions$: Actions) {
		this.state$ = store.select(CarrierEditDFState.get);
	}

	ngOnInit(): void {

		this.route.queryParams
			.subscribe((params) => {
				this.store.dispatch([new FetchCarrierEditDF(!!params.id ? params.id : '')]);
			});

		this.actions$.pipe(ofActionCompleted(UpdateFormErrors))
			.subscribe((payload) => {
				Object.keys(payload.action.payload.errors || {}).forEach((e) => {
					const form = this.form.get(e);
					if (!!form) {
						form.setErrors({form: payload.action.payload.errors![e]});
						form.markAsTouched();
					}
				});
			});

	}

	ngOnDestroy(): void {
		this.subscriptions$.map((s) => s.unsubscribe());
		this.store.dispatch(new Clear());
	}

	save() {
		this.store.dispatch(new CarrierEditDFActions.SaveForm());
	}
}
