import {NgModule} from "@angular/core";
import {CommonModule, NgOptimizedImage} from "@angular/common";
import {NgxsModule} from "@ngxs/store";
import {NgxsFormPluginModule} from "@ngxs/form-plugin";
import {FormsModule, ReactiveFormsModule} from "@angular/forms";
import {NgxsFormErrorsPluginModule} from "../../plugins/ngxs-form-errors/ngxs-form-errors.module";
import {MaterialModule} from "../../modules/material.module";
import {DvoCardComponent} from "../../shared/dvo-card/dvo-card.component";
import {NgxsFormArrayPluginModule} from "../../plugins/ngxs-form-array/ngxs-form-array.module";
import {EmailTemplatesListComponent} from "./email-templates-list/email-templates-list.component";
import {EmailTemplateEditComponent} from "./email-template-edit/email-template-edit.component";
import {EmailTemplatesRoutingModule} from "./email-templates-routing.module";
import {EmailTemplatesListState} from "./email-templates-list/email-templates-list.ngxs";
import {AddEmailTemplateComponent} from './email-templates-list/dialogs/add-email-template/add-email-template.component';
import {EmailTemplateEditState} from "./email-template-edit/email-template-edit.ngxs";
import {
	TestEmailTemplateComponent
} from "./email-template-edit/dialogs/test-email-template/test-email-template.component";

@NgModule({
	imports: [
		EmailTemplatesRoutingModule,
		NgxsModule.forFeature([
			EmailTemplatesListState,
			EmailTemplateEditState,
		]),
		MaterialModule,
		CommonModule,
		FormsModule,
		ReactiveFormsModule,
		DvoCardComponent,
		NgxsFormPluginModule,
		NgxsFormErrorsPluginModule,
		NgxsFormArrayPluginModule,
		NgOptimizedImage,
	],
	providers: [
	],
	declarations: [
		EmailTemplatesListComponent,
		EmailTemplateEditComponent,
        AddEmailTemplateComponent,
        TestEmailTemplateComponent,
	]
})
export class EmailTemplatesModule { }
