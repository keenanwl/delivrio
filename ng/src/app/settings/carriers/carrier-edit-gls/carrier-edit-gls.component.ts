import {Component, OnDestroy, OnInit} from '@angular/core';
import {FormControl, FormGroup} from "@angular/forms";
import {Observable, Subscription} from "rxjs";
import {CarrierEditModel, CarrierEditGLSState} from "./carrier-edit-gls.ngxs";
import {Actions, ofActionCompleted, Store} from "@ngxs/store";
import {ActivatedRoute} from "@angular/router";
import {CarrierEditGLSActions} from "./carrier-edit-gls.actions";
import {UpdateFormErrors} from "@ngxs/form-plugin";
import FetchCarrierEdit = CarrierEditGLSActions.FetchCarrierEdit;
import Clear = CarrierEditGLSActions.Clear;

@Component({
	selector: 'app-carrier-edit',
	templateUrl: './carrier-edit-gls.component.html',
	styleUrls: ['./carrier-edit-gls.component.scss']
})
export class CarrierEditGLSComponent implements OnInit, OnDestroy {

	carrierEdit$: Observable<CarrierEditModel>;

	editForm = new FormGroup({
		id: new FormControl(''),
		name: new FormControl(''),
		carrierGLS: new FormGroup({
			contactID: new FormControl(''),
			glsUsername: new FormControl(''),
			glsPassword: new FormControl(''),
			glsCountryCode: new FormControl(''),
			customerID: new FormControl(''),
			syncShipmentCancellation: new FormControl(false),
			printErrorOnLabel: new FormControl(false),
		}),
	});

	constructor(
		private store: Store,
	    private route: ActivatedRoute,
	    private actions$: Actions,
	) {
		this.carrierEdit$ = store.select(CarrierEditGLSState.get);
	}

	subscriptions$: Subscription[] = [];

	ngOnDestroy(): void {
		this.subscriptions$.map((s) => s.unsubscribe());
		this.store.dispatch(new Clear());
	}

	ngOnInit(): void {

		this.subscriptions$.push(this.route.queryParams
			.subscribe((params) => {
				this.store.dispatch([new FetchCarrierEdit(!!params.id ? params.id : '')]);
			}));

		this.subscriptions$.push(this.actions$.pipe(ofActionCompleted(UpdateFormErrors))
			.subscribe((payload) => {
				Object.keys(payload.action.payload.errors || {}).forEach((e) => {
					const form = this.editForm.get(e);
					if (!!form) {
						form.setErrors({form: payload.action.payload.errors![e]});
						form.markAsTouched();
					}
				});
			}));

	}

	onSubmit() {
		this.store.dispatch(new CarrierEditGLSActions.SaveForm());
	}

}
