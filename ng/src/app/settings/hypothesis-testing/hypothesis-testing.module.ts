import {NgModule} from "@angular/core";
import {CommonModule} from "@angular/common";
import {MaterialModule} from "../../modules/material.module";
import {NgxsFormPluginModule} from "@ngxs/form-plugin";
import {NgxsFormErrorsPluginModule} from "../../plugins/ngxs-form-errors/ngxs-form-errors.module";
import {FormsModule, ReactiveFormsModule} from "@angular/forms";
import {NgxsModule} from "@ngxs/store";
import {DvoCardComponent} from "../../shared/dvo-card/dvo-card.component";
import {NgxsFormArrayPluginModule} from "../../plugins/ngxs-form-array/ngxs-form-array.module";
import {HypothesisTestingRoutingModule} from "./hypothesis-testing-routing.module";
import {HypothesisTestingListComponent} from "./hypothesis-testing-list/hypothesis-testing-list.component";
import {HypothesisTestingListState} from "./hypothesis-testing-list/hypothesis-testing-list.ngxs";
import {AddNewHypothesisTestDialogComponent} from './hypothesis-testing-list/dialogs/add-new-hypothesis-test-dialog/add-new-hypothesis-test-dialog.component';
import {HypothesisTestingDeliveryOptionEditComponent} from './edit/hypothesis-testing-delivery-option-edit/hypothesis-testing-delivery-option-edit.component';
import {
	HypothesisTestingDeliveryOptionEditState
} from "./edit/hypothesis-testing-delivery-option-edit/hypothesis-testing-delivery-option-edit.ngxs";
import {CdkDrag, CdkDropList, CdkDropListGroup} from "@angular/cdk/drag-drop";

@NgModule({
	imports: [
		HypothesisTestingRoutingModule,
		NgxsModule.forFeature([
			HypothesisTestingListState,
			HypothesisTestingDeliveryOptionEditState,
		]),
		MaterialModule,
		CommonModule,
		FormsModule,
		ReactiveFormsModule,
		DvoCardComponent,
		NgxsFormPluginModule,
		NgxsFormErrorsPluginModule,
		NgxsFormArrayPluginModule,
		CdkDrag,
		CdkDropList,
		CdkDropListGroup,
	],
	declarations: [
		HypothesisTestingListComponent,
		AddNewHypothesisTestDialogComponent,
		HypothesisTestingDeliveryOptionEditComponent,
	]
})
export class HypothesisTestingModule { }
