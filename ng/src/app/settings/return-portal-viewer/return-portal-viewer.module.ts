import {NgModule} from "@angular/core";
import {CommonModule, NgOptimizedImage} from "@angular/common";
import {NgxsModule} from "@ngxs/store";
import {NgxsFormPluginModule} from "@ngxs/form-plugin";
import {FormsModule, ReactiveFormsModule} from "@angular/forms";
import {NgxsFormErrorsPluginModule} from "../../plugins/ngxs-form-errors/ngxs-form-errors.module";
import {MaterialModule} from "../../modules/material.module";
import {DvoCardComponent} from "../../shared/dvo-card/dvo-card.component";
import {NgxsFormArrayPluginModule} from "../../plugins/ngxs-form-array/ngxs-form-array.module";
import {ReturnPortalViewerRoutingModule} from "./return-portal-viewer-routing.module";
import {ReturnPortalFrameState} from "./return-portal-frame/return-portal-frame.ngxs";
import {
	ReturnPortalSelectItemsModule
} from "./return-portal-frame/return-portal-select-items/return-portal-select-items.module";
import {GoogleMapsModule} from "@angular/google-maps";
import {ReturnPortalFrameModule} from "./return-portal-frame/return-portal-frame.module";

@NgModule({
	imports: [
		ReturnPortalViewerRoutingModule,
		NgxsModule.forFeature([
			ReturnPortalFrameState
		]),
		MaterialModule,
		CommonModule,
		FormsModule,
		ReactiveFormsModule,
		DvoCardComponent,
		NgxsFormPluginModule,
		NgxsFormErrorsPluginModule,
		NgxsFormArrayPluginModule,
		ReturnPortalSelectItemsModule,
		NgOptimizedImage,
		GoogleMapsModule,
		ReturnPortalFrameModule,
	],
})
export class ReturnPortalViewerModule { }
