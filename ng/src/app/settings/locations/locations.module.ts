import {NgModule} from "@angular/core";
import {CommonModule} from "@angular/common";
import {MaterialModule} from "../../modules/material.module";
import {LocationsListComponent} from "./locations-list/locations-list.component";
import {NgxsFormPluginModule} from "@ngxs/form-plugin";
import {NgxsFormErrorsPluginModule} from "../../plugins/ngxs-form-errors/ngxs-form-errors.module";
import {FormsModule, ReactiveFormsModule} from "@angular/forms";
import {NgxsModule} from "@ngxs/store";
import {LocationEditComponent} from "./location-edit/location-edit.component";
import {LocationsRoutingModule} from "./locations-routing.module";
import {DvoCardComponent} from "../../shared/dvo-card/dvo-card.component";
import {LocationsListState} from "./locations-list/locations-list.ngxs";
import {LocationEditState} from "./location-edit/location-edit.ngxs";
import {NgxsFormArrayPluginModule} from "../../plugins/ngxs-form-array/ngxs-form-array.module";

@NgModule({
	imports: [
		LocationsRoutingModule,
		NgxsModule.forFeature([
			LocationsListState,
			LocationEditState,
		]),
		MaterialModule,
		CommonModule,
		FormsModule,
		ReactiveFormsModule,
		DvoCardComponent,
		NgxsFormPluginModule,
		NgxsFormErrorsPluginModule,
		NgxsFormArrayPluginModule,
	],
	declarations: [
		LocationsListComponent,
		LocationEditComponent,
	]
})
export class LocationsModule { }
