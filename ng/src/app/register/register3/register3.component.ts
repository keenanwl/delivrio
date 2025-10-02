import {Component, OnDestroy, OnInit} from '@angular/core';
import {Store} from "@ngxs/store";
import {Observable, Subscription} from "rxjs";
import {Register3Actions} from "./register3.actions";
import FetchCarrierPlatformLists = Register3Actions.FetchCarrierPlatformLists;
import {Register3Model, Register3State} from "./register3.ngxs";
import SaveRegistration = Register3Actions.SaveRegistration;
import SetPlatforms = Register3Actions.SetPlatforms;
import SetCarriers = Register3Actions.SetCarriers;

@Component({
	selector: 'app-register3',
	templateUrl: './register3.component.html',
	styleUrls: ['./register3.component.scss']
})
export class Register3Component implements OnInit, OnDestroy {

	register3$: Observable<Register3Model>;

	subscriptions: Subscription[] = [];

	constructor(
		private store: Store,
	) {
		this.register3$ = store.select(Register3State.state);
	}

	ngOnDestroy(): void {
		this.subscriptions.map((s) => s.unsubscribe());
	}

	ngOnInit(): void {
		this.store.dispatch([
			new FetchCarrierPlatformLists(),
		]);
	}

	platformChange(val: string[]) {
		this.store.dispatch(new SetPlatforms(val));
	}

	carrierChange(val: string[]) {
		this.store.dispatch(new SetCarriers(val));
	}

	signup() {
		this.store.dispatch(new SaveRegistration());
	}

}
