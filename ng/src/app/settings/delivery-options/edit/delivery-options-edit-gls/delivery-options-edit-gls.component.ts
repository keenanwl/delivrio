import {Component, OnDestroy, OnInit} from '@angular/core';
import {Store} from "@ngxs/store";
import {ActivatedRoute} from "@angular/router";
import SaveForm = DeliveryOptionsGLSEditActions.SaveForm;
import {DeliveryOptionsGLSEditActions} from "./delivery-options-edit-gls.actions";
import FetchDeliveryOptionsGLSEdit = DeliveryOptionsGLSEditActions.FetchDeliveryOptionsGLSEdit;
import {Observable, Subscription} from "rxjs";
import {DeliveryOptionsGLSEditModel, DeliveryOptionGLSEditState} from "./delivery-options-edit-gls.ngxs";
import SetSelectedOption = DeliveryOptionsGLSEditActions.SetSelectedOption;
import Clear = DeliveryOptionsGLSEditActions.Clear;
import {BaseDeliveryOptionFragment} from "../edit-common.generated";
import SetDeliveryOptionsGLSEdit = DeliveryOptionsGLSEditActions.SetDeliveryOptionsGLSEdit;

@Component({
	selector: 'app-delivery-options-edit-gls',
	templateUrl: './delivery-options-edit-gls.component.html',
	styleUrls: ['./delivery-options-edit-gls.component.scss']
})
export class DeliveryOptionsEditGlsComponent implements OnInit, OnDestroy {

	deliveryOptionsGLSEdit$: Observable<DeliveryOptionsGLSEditModel>;
	subscriptions$: Subscription[] = [];

	constructor(
		private store: Store,
		private route: ActivatedRoute,
	) {
		this.deliveryOptionsGLSEdit$ = store.select(DeliveryOptionGLSEditState.get);
	}

	ngOnDestroy(): void {
        this.subscriptions$.forEach((s) => s.unsubscribe());
		this.store.dispatch(new Clear());
    }

	ngOnInit(): void {
		this.subscriptions$.push(this.route.queryParams
			.subscribe((params) => {
				this.store.dispatch([
					new SetSelectedOption(!!params.id ? params.id : ''),
					new FetchDeliveryOptionsGLSEdit(),
				]);
			}));
	}

	baseFormChange(val: BaseDeliveryOptionFragment) {
		this.store.dispatch(new SetDeliveryOptionsGLSEdit(val));
	}

	save() {
		const state = this.store.selectSnapshot(DeliveryOptionGLSEditState.get);

		this.store.dispatch(new SaveForm({
			id: state.selectedOption,
			inputDeliveryOption: Object.assign({},
				state.deliveryOptionsGLSEditForm.model,
				{carrierServiceID: state.deliveryOptionsGLSEditForm.model?.carrierService.id},
				{carrierService: undefined},
				{clearDefaultPackaging: true,
					defaultPackaging: undefined,
					defaultPackagingID: state.deliveryOptionsGLSEditForm.model?.defaultPackaging?.id,}),
		}));
	}

}
