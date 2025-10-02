import {NgModule} from "@angular/core";
import {FormsModule, ReactiveFormsModule} from "@angular/forms";
import {NgxsModule} from "@ngxs/store";
import {CommonModule} from "@angular/common";
import {DvoCardComponent} from "../../shared/dvo-card/dvo-card.component";
import {NgxsFormPluginModule} from "@ngxs/form-plugin";
import {NgxsFormErrorsPluginModule} from "../../plugins/ngxs-form-errors/ngxs-form-errors.module";
import {WorkstationsRoutingModule} from "./workstations-routing.module";
import {WorkstationsListState} from "./workstations-list/workstations-list.ngxs";
import {MaterialModule} from "../../modules/material.module";
import {NewWorkstationComponent} from './dialogs/new-workstation/new-workstation.component';
import {WorkstationsListComponent} from './workstations-list/workstations-list.component';
import {WorkstationEditComponent} from './workstation-edit/workstation-edit.component';
import {WorkstationEditState} from "./workstation-edit/workstation-edit.ngxs";
import {NgxsFormArrayPluginModule} from "../../plugins/ngxs-form-array/ngxs-form-array.module";
import {WorkstationStatusPipePipe} from "./pipes/workstation-status-pipe.pipe";
import {CdkCopyToClipboard} from "@angular/cdk/clipboard";
import {ToggleContainerComponent} from "../../shared/toggle-container/toggle-container.component";
import {RelativeTimePipe} from "../../pipes/relative-time.pipe";

@NgModule({
	imports: [
		WorkstationsRoutingModule,
		NgxsModule.forFeature([
			WorkstationsListState,
			WorkstationEditState,
		]),
		MaterialModule,
		CommonModule,
		FormsModule,
		ReactiveFormsModule,
		DvoCardComponent,
		NgxsFormPluginModule,
		NgxsFormErrorsPluginModule,
		NgxsFormArrayPluginModule,
		WorkstationStatusPipePipe,
		CdkCopyToClipboard,
		ToggleContainerComponent,
		RelativeTimePipe,
	],
	declarations: [
		NewWorkstationComponent,
		WorkstationsListComponent,
		WorkstationEditComponent,
	]
})
export class WorkstationsModule { }
