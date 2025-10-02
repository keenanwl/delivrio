import {NgModule} from "@angular/core";
import {FormsModule, ReactiveFormsModule} from "@angular/forms";
import {NgxsModule} from "@ngxs/store";
import {CommonModule} from "@angular/common";
import {DvoCardComponent} from "../../shared/dvo-card/dvo-card.component";
import {NgxsFormPluginModule} from "@ngxs/form-plugin";
import {NgxsFormErrorsPluginModule} from "../../plugins/ngxs-form-errors/ngxs-form-errors.module";
import {DeliveryOptionGLSEditState} from "./edit/delivery-options-edit-gls/delivery-options-edit-gls.ngxs";
import {DeliveryOptionsRoutingModule} from "./delivery-options-routing.module";
import {DeliveryOptionsListState} from "./delivery-options-list/delivery-options-list.ngxs";
import {MaterialModule} from "../../modules/material.module";
import {DeliveryOptionsListComponent} from "./delivery-options-list/delivery-options-list.component";
import {
	DeliveryOptionEditPostNordComponent
} from "./edit/delivery-option-edit-post-nord/delivery-option-edit-post-nord.component";
import {DeliveryOptionEditPostNordState} from "./edit/delivery-option-edit-post-nord/delivery-option-edit-post-nord.ngxs";
import {TimeRangeSelectorModule} from "../../shared/time-range-selector/time-range-selector.module";
import {DeliveryOptionsEditGlsComponent} from "./edit/delivery-options-edit-gls/delivery-options-edit-gls.component";
import {DeliveryOptionEditRulesState} from "./edit/delivery-option-edit-rules/delivery-options-edit-rules.ngxs";
import {
	NewDeliveryOptionsRuleDialogComponent
} from "./edit/delivery-option-edit-rules/dialogs/new-delivery-options-rule-dialog.component";
import {
	NewDeliveryOptionsConstraintGroupDialogComponent
} from "./edit/delivery-option-edit-rules/dialogs/new-delivery-options-constraint-group-dialog.component";
import {DeliveryOptionEditRulesComponent} from "./edit/delivery-option-edit-rules/delivery-option-edit-rules.component";
import {NewDeliveryOptionDialogComponent} from "./delivery-options-list/new-delivery-option-dialog.component";
import {ConfirmDeleteRuleComponent} from './edit/delivery-option-edit-rules/dialogs/confirm-delete-rule/confirm-delete-rule.component';
import {CdkDrag, CdkDropList} from "@angular/cdk/drag-drop";
import {DeliveryOptionEmailTemplatesComponent} from './edit/delivery-option-email-templates/delivery-option-email-templates.component';
import {EmailTemplateFilterPipe} from './edit/delivery-option-email-templates/pipes/email-template-filter.pipe';
import {DeliveryOptionEditUspsComponent} from './edit/delivery-option-edit-usps/delivery-option-edit-usps.component';
import {DeliveryOptionEditUSPSState} from "./edit/delivery-option-edit-usps/delivery-option-edit-usps.ngxs";
import {CarrierServiceGrouperPipe} from "./edit/pipes/carrier-service-grouper.pipe";
import {DeliveryOptionEditDaoComponent} from './edit/delivery-option-edit-dao/delivery-option-edit-dao.component';
import {DeliveryOptionEditBringComponent} from './edit/delivery-option-edit-bring/delivery-option-edit-bring.component';
import {DeliveryOptionEditDAOState} from "./edit/delivery-option-edit-dao/delivery-option-edit-dao.ngxs";
import {DeliveryOptionEditBringState} from "./edit/delivery-option-edit-bring/delivery-option-edit-bring.ngxs";
import {DeliveryOptionEditDsvComponent} from './edit/delivery-option-edit-dsv/delivery-option-edit-dsv.component';
import {DeliveryOptionEditDSVState} from "./edit/delivery-option-edit-dsv/delivery-option-edit-dsv.ngxs";
import {DeliveryOptionEditBaseComponent} from "./edit/delivery-option-edit-base/delivery-option-edit-base.component";
import {DeliveryOptionEditDfComponent} from './edit/delivery-option-edit-df/delivery-option-edit-df.component';
import {DeliveryOptionEditDFState} from "./edit/delivery-option-edit-df/delivery-option-edit-df.ngxs";
import {ToggleContainerComponent} from "../../shared/toggle-container/toggle-container.component";
import {
	DeliveryOptionEditEasyPostState
} from "./edit/delivery-option-edit-easy-post/delivery-option-edit-easy-post.ngxs";
import {
	DeliveryOptionEditEasyPostComponent
} from "./edit/delivery-option-edit-easy-post/delivery-option-edit-easy-post.component";

@NgModule({
	imports: [
		DeliveryOptionsRoutingModule,
		DeliveryOptionEditEasyPostComponent,
		DeliveryOptionEditRulesComponent,
		DeliveryOptionEditBaseComponent,
		NgxsModule.forFeature([
			DeliveryOptionsListState,
			DeliveryOptionEditBringState,
			DeliveryOptionEditDAOState,
			DeliveryOptionEditDFState,
			DeliveryOptionEditDSVState,
			DeliveryOptionEditEasyPostState,
			DeliveryOptionGLSEditState,
			DeliveryOptionEditPostNordState,
			DeliveryOptionEditUSPSState,
			DeliveryOptionEditRulesState,
		]),
		MaterialModule,
		CommonModule,
		FormsModule,
		ReactiveFormsModule,
		DvoCardComponent,
		NgxsFormPluginModule,
		NgxsFormErrorsPluginModule,
		TimeRangeSelectorModule,
		CdkDrag,
		CdkDropList,
		CarrierServiceGrouperPipe,
		ToggleContainerComponent,
	],
    exports: [
        DeliveryOptionEditBaseComponent
    ],
	declarations: [
		DeliveryOptionsEditGlsComponent,
		DeliveryOptionEditPostNordComponent,
		DeliveryOptionsListComponent,

		NewDeliveryOptionsRuleDialogComponent,
		NewDeliveryOptionsConstraintGroupDialogComponent,

		NewDeliveryOptionDialogComponent,
		ConfirmDeleteRuleComponent,
		DeliveryOptionEmailTemplatesComponent,
		EmailTemplateFilterPipe,
		DeliveryOptionEditUspsComponent,
		DeliveryOptionEditDaoComponent,
		DeliveryOptionEditBringComponent,
		DeliveryOptionEditDsvComponent,
		DeliveryOptionEditDfComponent
	]
})
export class DeliveryOptionsModule { }
