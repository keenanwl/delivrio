import {Component, OnDestroy, OnInit} from '@angular/core';
import {FormControl, FormGroup} from "@angular/forms";
import {UpdateFormErrors} from "@ngxs/form-plugin";
import {Observable, Subscription} from "rxjs";
import {Actions, ofActionCompleted, Store} from "@ngxs/store";
import {ActivatedRoute} from "@angular/router";
import {CarrierEditBringModel, CarrierEditBringState} from "./carrier-edit-bring.ngxs";
import FetchCarrierEditBring = CarrierEditBringActions.FetchCarrierEditBring;
import {CarrierEditBringActions} from "./carrier-edit-bring.actions";
import Clear = CarrierEditBringActions.Clear;

@Component({
	selector: 'app-carrier-edit-bring',
	templateUrl: './carrier-edit-bring.component.html',
	styleUrl: './carrier-edit-bring.component.scss'
})
export class CarrierEditBringComponent implements OnInit, OnDestroy {

	state$: Observable<CarrierEditBringModel>;

	form = new FormGroup({
		id: new FormControl(''),
		name: new FormControl(''),
		carrierBring: new FormGroup({
			test: new FormControl(true),
			customerNumber: new FormControl(''),
		}),
	});

	constructor(private store: Store,
				private route: ActivatedRoute,
				private actions$: Actions) {
		this.state$ = store.select(CarrierEditBringState.get);
	}

	subscriptions$: Subscription[] = [];

	ngOnDestroy(): void {
		this.subscriptions$.map((s) => s.unsubscribe());
		this.store.dispatch(new Clear());
	}

	ngOnInit(): void {

		this.subscriptions$.push(this.route.queryParams
			.subscribe((params) => {
				this.store.dispatch([new FetchCarrierEditBring(!!params.id ? params.id : '')]);
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
		this.store.dispatch(new CarrierEditBringActions.SaveForm());
	}

}
