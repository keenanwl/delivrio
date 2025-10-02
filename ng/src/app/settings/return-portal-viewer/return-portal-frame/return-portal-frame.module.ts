import {NgModule} from "@angular/core";
import {CommonModule, NgOptimizedImage} from "@angular/common";
import {NgxsModule} from "@ngxs/store";
import {NgxsFormPluginModule} from "@ngxs/form-plugin";
import {FormsModule, ReactiveFormsModule} from "@angular/forms";
import {GoogleMapsModule} from "@angular/google-maps";
import {ReturnPortalViewerRoutingModule} from "../return-portal-viewer-routing.module";
import {ReturnPortalFrameState} from "./return-portal-frame.ngxs";
import {MaterialModule} from "../../../modules/material.module";
import {DvoCardComponent} from "../../../shared/dvo-card/dvo-card.component";
import {NgxsFormErrorsPluginModule} from "../../../plugins/ngxs-form-errors/ngxs-form-errors.module";
import {NgxsFormArrayPluginModule} from "../../../plugins/ngxs-form-array/ngxs-form-array.module";
import {ReturnPortalSelectItemsModule} from "./return-portal-select-items/return-portal-select-items.module";
import {ReturnPortalFrameService} from "./return-portal-frame.service";
import {ReturnPortalFrameComponent} from "./return-portal-frame.component";
import {ReturnFrameErrorDialogComponent} from "./dialogs/return-frame-error-dialog/return-frame-error-dialog.component";
import {StoreMapSelectorComponent} from "../store-map-selector/store-map-selector.component";
import {MAT_ICON_DEFAULT_OPTIONS, MatIconRegistry} from "@angular/material/icon";
import {DomSanitizer} from "@angular/platform-browser";
import {returnsIconSet} from "./frame-icons";

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
	],
	providers: [
		ReturnPortalFrameService,
		{
			provide: MAT_ICON_DEFAULT_OPTIONS,
			useFactory: (iconRegistry: MatIconRegistry, sanitizer: DomSanitizer) => {
				iconRegistry.addSvgIconSetLiteral(sanitizer.bypassSecurityTrustHtml(returnsIconSet()))
			},
			deps: [MatIconRegistry, DomSanitizer],
		},
	],
	exports: [
		ReturnPortalFrameComponent,
	],
	declarations: [
		ReturnPortalFrameComponent,
		ReturnFrameErrorDialogComponent,
		StoreMapSelectorComponent,
	]
})
export class ReturnPortalFrameModule { }
