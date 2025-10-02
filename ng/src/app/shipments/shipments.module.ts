import {NgModule} from "@angular/core";
import {MaterialModule} from "../modules/material.module";
import {CommonModule} from "@angular/common";
import {DvoCardComponent} from "../shared/dvo-card/dvo-card.component";
import {ShipmentsComponent} from "./shipments.component";
import {ShipmentEditComponent} from "./shipment-edit/shipment-edit.component";
import {ShipmentViewComponent} from "./shipment-view/shipment-view.component";
import {ShipmentsRoutingModule} from "./shipments-routing.module";
import {NgxsModule} from "@ngxs/store";
import {NgxsFormPluginModule} from "@ngxs/form-plugin";
import {NgxsFormErrorsPluginModule} from "../plugins/ngxs-form-errors/ngxs-form-errors.module";
import {ShipmentsListState} from "./shipments-list/shipments-list.ngxs";
import {ShipmentEditState} from "./shipment-edit/shipment-edit.ngxs";
import {ShipmentViewState} from "./shipment-view/shipment-view.ngxs";
import {FormsModule, ReactiveFormsModule} from "@angular/forms";
import {ShipmentsListComponent} from "./shipments-list/shipments-list.component";
import {OrderLinesModule} from "../shared/order-lines/order-lines.module";
import {ShipmentViewPostNordComponent} from './shipment-view/shipment-view-post-nord/shipment-view-post-nord.component';
import {PdfViewerModule} from "ng2-pdf-viewer";
import {ShipmentPreparePdfPipe} from "./shipment-view/shipment-view-post-nord/pipes/prepare-pdf.pipe";
import {PreparePdfPipe} from "../orders/order-view/pipes/prepare-pdf.pipe";
import {DebugSetIdsComponent} from './shipment-view/dialogs/debug-set-ids/debug-set-ids.component';
import {FilterBarModule} from "../shared/filter-bar/filter-bar.module";
import {SendFilteredEmailComponent} from './shipments-list/dialogs/send-filtered-email/send-filtered-email.component';
import {ShipmentStatusColorPipePipe} from "../shared/shipment-status-color-pipe.pipe";
import {RelativeTimePipe} from "../pipes/relative-time.pipe";
import {ShipmentParcelStatusColorPipePipe} from "../shared/shipment-parcel-status-color-pipe.pipe";

@NgModule({
	imports: [
		ShipmentsRoutingModule,
		NgxsModule.forFeature([
			ShipmentsListState,
			ShipmentEditState,
			ShipmentViewState,
		]),
		MaterialModule,
		CommonModule,
		FormsModule,
		ReactiveFormsModule,
		DvoCardComponent,
		NgxsFormPluginModule,
		NgxsFormErrorsPluginModule,
		OrderLinesModule,
		PdfViewerModule,
		ShipmentPreparePdfPipe,
		PreparePdfPipe,
		FilterBarModule,
		ShipmentStatusColorPipePipe,
		RelativeTimePipe,
		ShipmentParcelStatusColorPipePipe,
	],
	declarations: [
		ShipmentsComponent,
		ShipmentsListComponent,
		ShipmentEditComponent,
		ShipmentViewComponent,
  		ShipmentViewPostNordComponent,
    DebugSetIdsComponent,
    SendFilteredEmailComponent,
	]
})
export class ShipmentsModule { }
