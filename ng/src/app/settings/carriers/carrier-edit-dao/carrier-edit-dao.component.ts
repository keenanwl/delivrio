import {Component, OnDestroy, OnInit} from '@angular/core';
import {Observable, Subscription} from "rxjs";
import {CarrierEditDAOModel, CarrierEditDAOState} from "./carrier-edit-dao.ngxs";
import {FormControl, FormGroup} from "@angular/forms";
import {Actions, ofActionCompleted, Store} from "@ngxs/store";
import {ActivatedRoute} from "@angular/router";
import {UpdateFormErrors} from "@ngxs/form-plugin";
import FetchCarrierEditDAO = CarrierEditDAOActions.FetchCarrierEditDAO;
import {CarrierEditDAOActions} from "./carrier-edit-dao.actions";
import Clear = CarrierEditDAOActions.Clear;

@Component({
	selector: 'app-carrier-edit-dao',
	templateUrl: './carrier-edit-dao.component.html',
	styleUrl: './carrier-edit-dao.component.scss'
})
export class CarrierEditDaoComponent implements OnInit, OnDestroy {
	state$: Observable<CarrierEditDAOModel>;

	form = new FormGroup({
		id: new FormControl(''),
		name: new FormControl(''),
		carrierDAO: new FormGroup({
			customerID: new FormControl(''),
			apiKey: new FormControl(''),
		}),
	});

	subscriptions$: Subscription[] = [];

	constructor(private store: Store,
				private route: ActivatedRoute,
				private actions$: Actions) {
		this.state$ = store.select(CarrierEditDAOState.get);
	}

	ngOnInit(): void {

		this.route.queryParams
			.subscribe((params) => {
				this.store.dispatch([new FetchCarrierEditDAO(!!params.id ? params.id : '')]);
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
		this.store.dispatch(new CarrierEditDAOActions.SaveForm());
	}
}
