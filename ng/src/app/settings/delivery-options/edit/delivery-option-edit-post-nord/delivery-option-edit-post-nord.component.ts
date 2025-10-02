import {Component, OnDestroy, OnInit} from '@angular/core';
import {Observable, Subscription} from "rxjs";
import {Actions, ofActionCompleted, Store} from "@ngxs/store";
import {ActivatedRoute} from "@angular/router";
import {DeliveryOptionEditPostNordModel, DeliveryOptionEditPostNordState} from "./delivery-option-edit-post-nord.ngxs";
import {DeliveryOptionEditPostNordActions} from "./delivery-option-edit-post-nord.actions";
import Fetch = DeliveryOptionEditPostNordActions.Fetch;
import SetID = DeliveryOptionEditPostNordActions.SetID;
import SetAdditionalServiceEnabled = DeliveryOptionEditPostNordActions.SetAdditionalServiceEnabled;
import Save = DeliveryOptionEditPostNordActions.Save;
import SetShowUnavailableAdditionalServices = DeliveryOptionEditPostNordActions.SetShowUnavailableAdditionalServices;
import {UpdateFormValue} from "@ngxs/form-plugin";
import FetchAvailableAdditionalServices = DeliveryOptionEditPostNordActions.FetchAvailableAdditionalServices;
import {debounceTime} from "rxjs/operators";
import Clear = DeliveryOptionEditPostNordActions.Clear;
import {MatAutocompleteSelectedEvent} from "@angular/material/autocomplete";
import AddLocation = DeliveryOptionEditPostNordActions.AddLocation;
import RemoveLocation = DeliveryOptionEditPostNordActions.RemoveLocation;
import SetSelectedEmailTemplates = DeliveryOptionEditPostNordActions.SetSelectedEmailTemplates;
import {SelectedEmailTemplates} from "../delivery-option-email-templates/delivery-option-email-templates.component";
import {BaseDeliveryOptionFragment} from "../edit-common.generated";
import SetEditPostNord = DeliveryOptionEditPostNordActions.SetEditPostNord;

@Component({
	selector: 'app-delivery-option-edit-post-nord',
	templateUrl: './delivery-option-edit-post-nord.component.html',
	styleUrls: ['./delivery-option-edit-post-nord.component.scss']
})
export class DeliveryOptionEditPostNordComponent implements OnInit, OnDestroy {
	deliveryOptionsPostNordEdit$: Observable<DeliveryOptionEditPostNordModel>;
	subscriptions$: Subscription[] = [];

	constructor(
		private store: Store,
		private route: ActivatedRoute,
		private actions$: Actions,
	) {
		this.deliveryOptionsPostNordEdit$ = store.select(DeliveryOptionEditPostNordState.get);
	}

	save() {
		this.store.dispatch(new Save());
	}

	selectLocation(event: MatAutocompleteSelectedEvent) {
		this.store.dispatch(new AddLocation(event.option.value));
	}

	additionalServiceChange(internalID: string, checked: boolean) {
		this.store.dispatch(new SetAdditionalServiceEnabled({internalID, checked}));
	}

	isAdditionalServiceDisabled(internalID: string): boolean {
		const state = this.store.selectSnapshot(DeliveryOptionEditPostNordState.get).availableAdditionalServices;
		return !state.includes(internalID);
	}

	isAdditionalServiceEnabled(checkInternalID: string): boolean {
		const state = this.store.selectSnapshot(DeliveryOptionEditPostNordState.get).enabledAdditionalServices;
		return !!state.find((obj) => obj.internalID === checkInternalID);
	}

	toggleShowAdditional(checked: boolean) {
		this.store.dispatch(new SetShowUnavailableAdditionalServices(checked));
	}

	ngOnDestroy(): void {
		this.subscriptions$.forEach((s) => s.unsubscribe());
		this.store.dispatch(new Clear());
	}

	removeLocation(id: string) {
		this.store.dispatch(new RemoveLocation(id));
	}

	baseFormChange(val: BaseDeliveryOptionFragment) {
		this.store.dispatch(new SetEditPostNord(val));
	}

	changeSelectedEmailTemplates(val: SelectedEmailTemplates) {
		this.store.dispatch(new SetSelectedEmailTemplates(val));
	}

	ngOnInit(): void {
		this.subscriptions$.push(this.actions$
			.pipe(ofActionCompleted(UpdateFormValue), debounceTime(200))
			.subscribe(() => {
				this.store.dispatch(new FetchAvailableAdditionalServices());

			}));

		this.subscriptions$.push(this.route.queryParams
			.subscribe((params) => {
				this.store.dispatch([
					new SetID(!!params.id ? params.id : ''),
					new Fetch(),
				]);
			}));
	}
}
