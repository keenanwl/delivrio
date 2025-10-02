import {Component, OnDestroy, OnInit} from '@angular/core';
import {Observable, Subscription} from "rxjs";
import {CarrierEditDSVModel, CarrierEditDSVState} from "../carrier-edit-dsv/carrier-edit-dsv.ngxs";
import {FormControl, FormGroup} from "@angular/forms";
import {Actions, ofActionCompleted, Store} from "@ngxs/store";
import {ActivatedRoute} from "@angular/router";
import {UpdateFormErrors} from "@ngxs/form-plugin";
import {CarrierEditDSVActions} from "./carrier-edit-dsv.actions";
import FetchCarrierEditDSV = CarrierEditDSVActions.FetchCarrierEditDSV;
import Clear = CarrierEditDSVActions.Clear;

@Component({
	selector: 'app-carrier-edit-dsv',
	templateUrl: './carrier-edit-dsv.component.html',
	styleUrl: './carrier-edit-dsv.component.scss'
})
export class CarrierEditDsvComponent implements OnInit, OnDestroy {
	state$: Observable<CarrierEditDSVModel>;

	form = new FormGroup({
		id: new FormControl(''),
		name: new FormControl(''),
		carrierDSV: new FormGroup({
			customerID: new FormControl(''),
			apiKey: new FormControl(''),
		}),
	});

	constructor(private store: Store,
				private route: ActivatedRoute,
				private actions$: Actions) {
		this.state$ = store.select(CarrierEditDSVState.get);
	}

	subscriptions$: Subscription[] = [];

	ngOnDestroy(): void {
        this.subscriptions$.map((s) => s.unsubscribe());
		this.store.dispatch(new Clear());
    }

	ngOnInit(): void {

		this.subscriptions$.push(this.route.queryParams
			.subscribe((params) => {
				this.store.dispatch([new FetchCarrierEditDSV(!!params.id ? params.id : '')]);
			}));

		this.subscriptions$.push(this.actions$.pipe(ofActionCompleted(UpdateFormErrors))
			.subscribe((payload) => {
				Object.keys(payload.action.payload.errors || {}).forEach((e) => {
					const form = this.form.get(e);
					if (!!form) {
						form.setErrors({form: payload.action.payload.errors![e]});
						form.markAsTouched();
					}
				});
			}));

	}

	save() {
		this.store.dispatch(new CarrierEditDSVActions.SaveForm());
	}
}
