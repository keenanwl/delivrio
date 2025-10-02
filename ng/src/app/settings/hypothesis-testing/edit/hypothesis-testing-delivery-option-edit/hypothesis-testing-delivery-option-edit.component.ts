import {Component, OnInit} from '@angular/core';
import {FormArray, FormControl, FormGroup} from "@angular/forms";
import {Actions, ofActionCompleted, Store} from "@ngxs/store";
import {
	HypothesisTestingDeliveryOptionEditActions
} from "./hypothesis-testing-delivery-option-edit.actions";
import Fetch = HypothesisTestingDeliveryOptionEditActions.Fetch;
import {ActivatedRoute} from "@angular/router";
import {Observable} from "rxjs";
import {
	HypothesisTestingDeliveryOptionEditModel,
	HypothesisTestingDeliveryOptionEditState
} from "./hypothesis-testing-delivery-option-edit.ngxs";
import SetHypothesisTestingDeliveryOptionID = HypothesisTestingDeliveryOptionEditActions.SetHypothesisTestingDeliveryOptionID;
import {Paths} from "../../../../app-routing.module";
import SetHypothesisTestingDeliveryOptionEdit = HypothesisTestingDeliveryOptionEditActions.SetHypothesisTestingDeliveryOptionEdit;
import {CdkDragDrop} from "@angular/cdk/drag-drop";
import MoveDeliveryOption = HypothesisTestingDeliveryOptionEditActions.MoveDeliveryOption;
import HTDeliveryOptionResponse = HypothesisTestingDeliveryOptionEditActions.HTDeliveryOptionResponse;
import Save = HypothesisTestingDeliveryOptionEditActions.Save;

type formDeliveryOptionGroup = FormGroup<{
	id: FormControl<string>,
	name: FormControl<string>,
	description: FormControl<string>,
	hideDeliveryOption: FormControl<boolean>,
}>;

@Component({
  selector: 'app-hypothesis-testing-delivery-option-edit',
  templateUrl: './hypothesis-testing-delivery-option-edit.component.html',
  styleUrls: ['./hypothesis-testing-delivery-option-edit.component.scss']
})
export class HypothesisTestingDeliveryOptionEditComponent implements OnInit {

	state$: Observable<HypothesisTestingDeliveryOptionEditModel>;

	editForm = new FormGroup({
		name: new FormControl(""),
		active: new FormControl(false),
		hypothesisTestDeliveryOption: new FormGroup({
			byIntervalRotation: new FormControl(false),
			byOrder: new FormControl(false),
			randomizeWithinGroupSort: new FormControl(false),
			rotationIntervalHours: new FormControl(6),
			deliveryOptionGroupOne: new FormArray<formDeliveryOptionGroup>([]),
			deliveryOptionGroupTwo: new FormArray<formDeliveryOptionGroup>([]),
		})
	});

	deliveryOptionPath = Paths.SETTINGS_DELIVERY_OPTIONS;

	constructor(
		private store: Store,
		private route: ActivatedRoute,
		private actions$: Actions,
	) {
		this.state$ = store.select(HypothesisTestingDeliveryOptionEditState.get);
	}

	ngOnInit(): void {

		this.actions$.pipe(ofActionCompleted(SetHypothesisTestingDeliveryOptionEdit))
			.subscribe(() => {
				const state = this.store.selectSnapshot(HypothesisTestingDeliveryOptionEditState.get);
				this.editForm.controls.hypothesisTestDeliveryOption.controls.deliveryOptionGroupOne.clear();
				this.editForm.controls.hypothesisTestDeliveryOption.controls.deliveryOptionGroupTwo.clear();

				state.editForm.model?.hypothesisTestDeliveryOption?.deliveryOptionGroupOne?.forEach((opt) => {
					this.editForm.controls.hypothesisTestDeliveryOption.controls.deliveryOptionGroupOne.push(new FormGroup({
						id: new FormControl(opt.id || '', {nonNullable: true}),
						name: new FormControl(opt.name || '', {nonNullable: true}),
						description: new FormControl(opt.description || '', {nonNullable: true}),
						hideDeliveryOption: new FormControl<boolean>(opt.hideDeliveryOption || false, {nonNullable: true}),
					}))
				});
				state.editForm.model?.hypothesisTestDeliveryOption?.deliveryOptionGroupTwo?.forEach((opt) => {
					this.editForm.controls.hypothesisTestDeliveryOption.controls.deliveryOptionGroupTwo.push(new FormGroup({
						id: new FormControl(opt.id || '', {nonNullable: true}),
						name: new FormControl(opt.name || '', {nonNullable: true}),
						description: new FormControl(opt.description || '', {nonNullable: true}),
						hideDeliveryOption: new FormControl<boolean>(opt.hideDeliveryOption || false, {nonNullable: true}),
					}))
				});
			});

		this.route.queryParams
			.subscribe((params) => {
				this.store.dispatch([
					new SetHypothesisTestingDeliveryOptionID(!!params.id ? params.id : ''),
					new Fetch(),
				]);
			});
	}

	drop(evt: CdkDragDrop<any, any, HTDeliveryOptionResponse>, container: "available" | "control" | "test") {
		console.warn(evt);
		this.store.dispatch(new MoveDeliveryOption({container, deliveryOption: evt.item.data}));

	}

	save() {
		this.store.dispatch(new Save());
	}

}
