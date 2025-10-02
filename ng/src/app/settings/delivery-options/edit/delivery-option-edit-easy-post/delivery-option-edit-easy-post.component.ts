import {Component, OnDestroy, OnInit} from '@angular/core';
import {FormControl, FormGroup, ReactiveFormsModule} from "@angular/forms";
import {Observable, Subscription} from "rxjs";
import {
	DeliveryOptionEditEasyPostModel,
	DeliveryOptionEditEasyPostState
} from "./delivery-option-edit-easy-post.ngxs";
import {Actions, Store} from "@ngxs/store";
import {ActivatedRoute} from "@angular/router";
import {DeliveryOptionEditEasyPostActions} from "./delivery-option-edit-easy-post.actions";
import Clear = DeliveryOptionEditEasyPostActions.Clear;
import SetID = DeliveryOptionEditEasyPostActions.SetID;
import Fetch = DeliveryOptionEditEasyPostActions.Fetch;
import Save = DeliveryOptionEditEasyPostActions.Save;
import {AsyncPipe, NgForOf, NgIf} from "@angular/common";
import {NgxsFormPluginModule} from "@ngxs/form-plugin";
import {NgxsFormErrorsPluginModule} from "../../../../plugins/ngxs-form-errors/ngxs-form-errors.module";
import {DvoCardComponent} from "../../../../shared/dvo-card/dvo-card.component";
import {MatError, MatFormField, MatLabel} from "@angular/material/form-field";
import {CarrierServiceGrouperPipe} from "../pipes/carrier-service-grouper.pipe";
import {MatOptgroup, MatOption, MatSelect} from "@angular/material/select";
import {MatSlideToggle} from "@angular/material/slide-toggle";
import {MatIcon} from "@angular/material/icon";
import {MatFabButton, MatIconButton} from "@angular/material/button";
import {DeliveryOptionEditRulesComponent} from "../delivery-option-edit-rules/delivery-option-edit-rules.component";
import {MatInput} from "@angular/material/input";
import {ToggleContainerComponent} from "../../../../shared/toggle-container/toggle-container.component";
import {BaseDeliveryOptionFragment} from "../edit-common.generated";
import SetEditEasyPost = DeliveryOptionEditEasyPostActions.SetEditEasyPost;
import {DeliveryOptionEditBaseComponent} from "../delivery-option-edit-base/delivery-option-edit-base.component";
import ToggleAdditionalService = DeliveryOptionEditEasyPostActions.ToggleAdditionalService;

@Component({
	selector: 'app-delivery-option-edit-easy-post',
	standalone: true,
	imports: [
		AsyncPipe,
		NgxsFormPluginModule,
		NgxsFormErrorsPluginModule,
		DvoCardComponent,
		MatFormField,
		NgIf,
		CarrierServiceGrouperPipe,
		MatSelect,
		MatOptgroup,
		MatOption,
		NgForOf,
		MatLabel,
		MatError,
		ReactiveFormsModule,
		MatSlideToggle,
		MatIcon,
		MatFabButton,
		DeliveryOptionEditRulesComponent,
		MatInput,
		ToggleContainerComponent,
		DeliveryOptionEditBaseComponent,
		MatSelect,
		MatIconButton,
	],
	templateUrl: './delivery-option-edit-easy-post.component.html',
	styleUrl: './delivery-option-edit-easy-post.component.scss'
})
export class DeliveryOptionEditEasyPostComponent implements OnInit, OnDestroy {

	editForm = new FormGroup({
		clickCollect: new FormControl(true),
		overrideReturnAddress: new FormControl(true),
		overrideSenderAddress: new FormControl(true),
		hideDeliveryOption: new FormControl(true),
		name: new FormControl(''),
		paperless: new FormControl(true),
		senderPaysTax: new FormControl(true),
		requireSignature: new FormControl(true),
		carrierService: new FormGroup({
			id: new FormControl(""),
		}),
	});

	deliveryOptionsEasyPostEdit$: Observable<DeliveryOptionEditEasyPostModel>;
	subscriptions$: Subscription[] = [];

	constructor(
		private store: Store,
		private route: ActivatedRoute,
		private actions$: Actions,
	) {
		this.deliveryOptionsEasyPostEdit$ = store.select(DeliveryOptionEditEasyPostState.get);
	}

	ngOnDestroy(): void {
		this.subscriptions$.forEach((s) => s.unsubscribe());
		this.store.dispatch(new Clear());
	}

	ngOnInit(): void {

		this.subscriptions$.push(this.route.queryParams
			.subscribe((params) => {
				this.store.dispatch([
					new SetID(!!params.id ? params.id : ''),
					new Fetch(),
				]);
			}));
	}

	save() {
		this.store.dispatch(new Save());
	}

	baseFormChange(val: BaseDeliveryOptionFragment) {
		this.store.dispatch(new SetEditEasyPost(val));
	}

	additionalServiceSelected(selectedIDs: string[], id: string): boolean {
		return selectedIDs.some((i) => {
			if (id === i) {
				return true;
			}
			return false;
		});
	}

	additionalServiceToggled(id: string, isAdd: boolean) {
		this.store.dispatch(new ToggleAdditionalService({id: id, isAdd}));
	}

}
