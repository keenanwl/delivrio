import {NgModule} from "@angular/core";
import {FormsModule, ReactiveFormsModule } from "@angular/forms";
import {NgxsModule} from "@ngxs/store";
import {CommonModule} from "@angular/common";
import {Register2Component} from "./register2/register2.component";
import {Register3Component} from "./register3/register3.component";
import {RegisterComponent} from "./register.component";
import {NgxsFormArrayPluginModule} from "../plugins/ngxs-form-array/ngxs-form-array.module";
import {NgxsFormErrorsPluginModule} from "../plugins/ngxs-form-errors/ngxs-form-errors.module";
import {NgxsFormPluginModule} from "@ngxs/form-plugin";
import {MaterialModule} from "../modules/material.module";
import {Register1State} from "./register1/register1.ngxs";
import {Register2State} from "./register2/register2.ngxs";
import {Register3State} from "./register3/register3.ngxs";
import {RegisterRoutingModule} from "./register-routing.module";
import {RegisterPanelComponent} from "../shared/register-panel/register-panel.component";
import {RegisterService} from "./register1/register1.service.";
import {AlreadyCustomerHeaderComponent} from "./already-customer-header/already-customer-header.component";
import {Register1Component} from "./register1/register1.component";

@NgModule({
	imports: [
		RegisterRoutingModule,
		NgxsModule.forFeature([
			Register1State,
			Register2State,
			Register3State,
		]),
		MaterialModule,
		CommonModule,
		FormsModule,
		ReactiveFormsModule,
		NgxsFormPluginModule,
		NgxsFormErrorsPluginModule,
		NgxsFormArrayPluginModule,
	],
	declarations: [
		AlreadyCustomerHeaderComponent,
		RegisterComponent,
		Register1Component,
		Register2Component,
		Register3Component,
		RegisterPanelComponent,
	],
	providers: [
		RegisterService,
	]
})
export class RegisterModule { }
