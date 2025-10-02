import {Component, OnDestroy, OnInit} from '@angular/core';
import {Store} from "@ngxs/store";
import {Observable, Subscription} from "rxjs";
import {Register2Model, Register2State} from "./register2.ngxs";
import {Register2Actions} from "./register2.actions";
import SaveRegistration2 = Register2Actions.SaveRegistration2;
import {ActivatedRoute} from "@angular/router";
import {FormControl, FormGroup} from "@angular/forms";
import {AppState} from "../../app.ngxs";

@Component({
	selector: 'app-register2',
	templateUrl: './register2.component.html',
	styleUrls: ['./register2.component.scss']
})
export class Register2Component implements OnInit, OnDestroy {

	register2$: Observable<Register2Model>;

	register2Form = new FormGroup({
		betterDeliveryOptions: new FormControl(false, {nonNullable: true}),
		improvePickPack: new FormControl(false, {nonNullable: true}),
		shippingLabel: new FormControl(false, {nonNullable: true}),
		customDocs: new FormControl(false, {nonNullable: true}),
		reducedCosts: new FormControl(true, {nonNullable: true}),
		easyReturns: new FormControl(false, {nonNullable: true}),
		clickCollect: new FormControl(false, {nonNullable: true}),
		numShipments: new FormControl(5000, {nonNullable: true}),
	});

	subscriptions: Subscription[] = [];

	constructor(private store: Store, private route: ActivatedRoute) {
	  this.register2$ = store.select(Register2State.state);
	}

	ngOnDestroy(): void {
		this.subscriptions.map((s) => s.unsubscribe());
	}

	ngOnInit(): void {
	}

	formatShipmentCount(value: number) {
		if (value === 50_000) {
			return Math.round(value / 1000) + 'k+';
		}
		if (value >= 1000) {
			return Math.round(value / 1000) + 'k';
		}

		return value.toString();
	}

	next() {
		const myID = this.store.selectSnapshot(AppState.get).my_ids.my_pulid;
		const next = Object.assign({}, this.register2Form.getRawValue(), {usersID: myID});
		this.store.dispatch(new SaveRegistration2(next));
	}

}
