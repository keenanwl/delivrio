import {NgModule} from "@angular/core";
import {FormsModule, ReactiveFormsModule } from "@angular/forms";
import {NgxsModule} from "@ngxs/store";
import {CommonModule} from "@angular/common";
import {DvoCardComponent} from "../../shared/dvo-card/dvo-card.component";
import {NgxsFormPluginModule} from "@ngxs/form-plugin";
import {NgxsFormErrorsPluginModule} from "../../plugins/ngxs-form-errors/ngxs-form-errors.module";
import {APITokensState} from "./api-tokens.ngxs";
import {APITokensComponent} from "./api-tokens.component";
import {APITokensRoutingModule} from "./api-tokens-routing.module";
import {MaterialModule} from "../../modules/material.module";
import { AddNewApiTokenComponent } from './dialogs/add-new-api-token/add-new-api-token.component';
import {ClipboardModule} from "@angular/cdk/clipboard";
import { ApiTokenConfirmDeleteComponent } from './dialogs/api-token-confirm-delete/api-token-confirm-delete.component';

@NgModule({
	imports: [
		APITokensRoutingModule,
		NgxsModule.forFeature([
			APITokensState,
		]),
		MaterialModule,
		CommonModule,
		FormsModule,
		ReactiveFormsModule,
		DvoCardComponent,
		NgxsFormPluginModule,
		NgxsFormErrorsPluginModule,
		ClipboardModule,
	],
	declarations: [
		APITokensComponent,
  AddNewApiTokenComponent,
  ApiTokenConfirmDeleteComponent
	]
})
export class APITokensModule { }
