import {Component, OnDestroy, OnInit} from '@angular/core';
import {Observable, Subscription} from "rxjs";
import {Actions, ofActionCompleted, Store} from "@ngxs/store";
import {ActivatedRoute} from "@angular/router";
import {UpdateFormValue} from "@ngxs/form-plugin";
import {debounceTime} from "rxjs/operators";
import {DeliveryOptionEditUSPSActions} from "./delivery-option-edit-usps.actions";
import FetchAvailableAdditionalServices = DeliveryOptionEditUSPSActions.FetchAvailableAdditionalServices;
import SetID = DeliveryOptionEditUSPSActions.SetID;
import Fetch = DeliveryOptionEditUSPSActions.Fetch;
import Clear = DeliveryOptionEditUSPSActions.Clear;
import Save = DeliveryOptionEditUSPSActions.Save;
import {SelectedEmailTemplates} from "../delivery-option-email-templates/delivery-option-email-templates.component";
import SetAdditionalServiceEnabled = DeliveryOptionEditUSPSActions.SetAdditionalServiceEnabled;
import SetShowUnavailableAdditionalServices = DeliveryOptionEditUSPSActions.SetShowUnavailableAdditionalServices;
import SetSelectedEmailTemplates = DeliveryOptionEditUSPSActions.SetSelectedEmailTemplates;
import {DeliveryOptionEditUSPSModel, DeliveryOptionEditUSPSState} from "./delivery-option-edit-usps.ngxs";
import {UspsAdditionalServicesFragment} from "./delivery-option-edit-usps.generated";
import SetAdditionalServiceDisabled = DeliveryOptionEditUSPSActions.SetAdditionalServiceDisabled;
import {BaseDeliveryOptionFragment} from "../edit-common.generated";
import SetEditUSPS = DeliveryOptionEditUSPSActions.SetEditUSPS;

@Component({
	selector: 'app-delivery-option-edit-usps',
	templateUrl: './delivery-option-edit-usps.component.html',
	styleUrls: ['./delivery-option-edit-usps.component.scss']
})
export class DeliveryOptionEditUspsComponent implements OnInit, OnDestroy {

	deliveryOptionsUSPSEdit$: Observable<DeliveryOptionEditUSPSModel>;
	subscriptions$: Subscription[] = [];

	constructor(
		private store: Store,
		private route: ActivatedRoute,
		private actions$: Actions,
	) {
		this.deliveryOptionsUSPSEdit$ = store.select(DeliveryOptionEditUSPSState.get);
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

	ngOnDestroy(): void {
		this.subscriptions$.forEach((s) => s.unsubscribe());
		this.store.dispatch(new Clear());
	}

	save() {
		this.store.dispatch(new Save());
	}

	additionalServiceChange(addService: UspsAdditionalServicesFragment, checked: boolean) {
		if (checked) {
			this.store.dispatch(new SetAdditionalServiceEnabled(addService));
		} else {
			this.store.dispatch(new SetAdditionalServiceDisabled(addService));
		}
	}

	isAdditionalServiceCommon(addService: UspsAdditionalServicesFragment): boolean {
		const state = this.store.selectSnapshot(DeliveryOptionEditUSPSState.get).availableAdditionalServices;
		return !state.some((as) => {
			return as.internalID === addService.internalID;
		})
	}

	isAdditionalServiceEnabled(addService: UspsAdditionalServicesFragment): boolean {
		const state = this.store.selectSnapshot(DeliveryOptionEditUSPSState.get).enabledAdditionalServices;
		return state.some((as) => {
			return as.internalID === addService.internalID;
		})
	}

	toggleShowAdditional(checked: boolean) {
		this.store.dispatch(new SetShowUnavailableAdditionalServices(checked));
	}

	baseFormChange(val: BaseDeliveryOptionFragment) {
		this.store.dispatch(new SetEditUSPS(val));
	}

}
