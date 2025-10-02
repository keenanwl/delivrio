import {NgModule} from "@angular/core";
import {CommonModule, NgOptimizedImage} from "@angular/common";
import {NgxsModule} from "@ngxs/store";
import {NgxsFormPluginModule} from "@ngxs/form-plugin";
import {FormsModule, ReactiveFormsModule} from "@angular/forms";
import {NgxsFormErrorsPluginModule} from "../../plugins/ngxs-form-errors/ngxs-form-errors.module";
import {MaterialModule} from "../../modules/material.module";
import {DvoCardComponent} from "../../shared/dvo-card/dvo-card.component";
import {NgxsFormArrayPluginModule} from "../../plugins/ngxs-form-array/ngxs-form-array.module";
import {PackagingRoutingModule} from "./packaging-routing.module";
import {PackagingListState} from "./packaging-list/packaging-list.ngxs";
import {PackagingListComponent} from "./packaging-list/packaging-list.component";
import {CreatePackagingComponent} from './packaging-list/dialogs/create-packaging/create-packaging.component';
import { DimensionsViewerComponent } from './shared/dimensions-viewer/dimensions-viewer.component';

@NgModule({
	imports: [
		PackagingRoutingModule,
		NgxsModule.forFeature([
			PackagingListState,
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
		PackagingListComponent,
  		CreatePackagingComponent,
    DimensionsViewerComponent,
	]
})
export class PackagingModule { }
