import {NgModule} from "@angular/core";
import {CommonModule} from "@angular/common";
import {MaterialModule} from "../../modules/material.module";
import {EmailTemplateEditorComponent} from "./email-template-editor.component";
import {PaperbitsDesignerComponent} from './paperbits/paperbits.component';

@NgModule({
	imports: [
		MaterialModule,
		CommonModule,
	],
	exports: [
		EmailTemplateEditorComponent,
		PaperbitsDesignerComponent
	],
	declarations: [
		EmailTemplateEditorComponent,
		PaperbitsDesignerComponent
	]
})
export class EmailTemplateEditorModule { }
