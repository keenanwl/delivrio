import {Component, OnDestroy, OnInit} from '@angular/core';
import {Observable, Subscription} from "rxjs";
import {CarrierEditPostNordModel, CarrierEditPostNordState} from "./carrier-edit-post-nord.ngxs";
import {Store} from "@ngxs/store";
import {FormControl, FormGroup} from "@angular/forms";
import {ActivatedRoute} from "@angular/router";
import {CarrierEditPostNordActions} from "./carrier-edit-post-nord.actions";
import SetID = CarrierEditPostNordActions.SetID;
import FetchCarrierPostNordEdit = CarrierEditPostNordActions.FetchCarrierPostNordEdit;
import Clear = CarrierEditPostNordActions.Clear;
import SaveForm = CarrierEditPostNordActions.SaveForm;

@Component({
	selector: 'app-carrier-edit-post-nord',
	templateUrl: './carrier-edit-post-nord.component.html',
	styleUrls: ['./carrier-edit-post-nord.component.scss']
})
export class CarrierEditPostNordComponent implements OnInit, OnDestroy {

	carrierEditPostNord$: Observable<CarrierEditPostNordModel>;

	editForm = new FormGroup({
		carrier: new FormGroup({
			name: new FormControl('', {nonNullable: true}),
		}),
		customerNumber: new FormControl('', {nonNullable: true}),
	});

	subscriptions$: Subscription[] = [];

	constructor(
		private route: ActivatedRoute,
		private store: Store
	) {
		this.carrierEditPostNord$ = store.select(CarrierEditPostNordState.get);
	}

	ngOnInit() {
		this.subscriptions$.push(this.route.queryParams
			.subscribe((params) => {
				this.store.dispatch([
					new SetID(!!params.id ? params.id : ''),
					new FetchCarrierPostNordEdit(),
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
