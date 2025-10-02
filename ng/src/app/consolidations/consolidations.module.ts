import {NgModule} from "@angular/core";
import {MaterialModule} from "../modules/material.module";
import {CommonModule} from "@angular/common";
import {DvoCardComponent} from "../shared/dvo-card/dvo-card.component";
import {FormsModule, ReactiveFormsModule} from "@angular/forms";
import {ConsolidationsRoutingModule} from "./consolidations-routing.module";
import {NgxsModule} from "@ngxs/store";
import {NgxsFormPluginModule} from "@ngxs/form-plugin";
import {NgxsFormErrorsPluginModule} from "../plugins/ngxs-form-errors/ngxs-form-errors.module";
import {CdkMenu, CdkMenuItem, CdkMenuItemCheckbox, CdkMenuTrigger} from "@angular/cdk/menu";
import {DragDropModule} from "@angular/cdk/drag-drop";
import {RouterOutlet} from "@angular/router";
import {ConsolidationsListState} from "./consolidations-list/consolidations-list.ngxs";
import {ConsolidationsListComponent} from "./consolidations-list/consolidations-list.component";
import {AddConsolidationComponent} from './consolidations-list/dialogs/add-consolidation/add-consolidation.component';
import {ConsolidationEditComponent} from './consolidation-edit/consolidation-edit.component';
import {ConsolidationEditState} from "./consolidation-edit/consolidation-edit.ngxs";
import {NgxsFormArrayPluginModule} from "../plugins/ngxs-form-array/ngxs-form-array.module";
import {EditPalletComponent} from './consolidation-edit/dialogs/edit-pallet/edit-pallet.component';
import {OrderStatusColorPipePipe} from "../shared/order-status-color-pipe.pipe";
import {CarrierServiceGrouperPipe} from "../settings/delivery-options/edit/pipes/carrier-service-grouper.pipe";
import {DeliveryOptionGrouperPipe} from "./pipes/delivery-option-grouper.pipe";
import {ShipmentStatusColorPipePipe} from "../shared/shipment-status-color-pipe.pipe";
import {CreateShipmentComponent} from './consolidation-edit/dialogs/create-shipment/create-shipment.component';
import {PdfViewerModule} from "ng2-pdf-viewer";
import {PreparePdfPipe} from "../orders/order-view/pipes/prepare-pdf.pipe";

@NgModule({
	imports: [
		ConsolidationsRoutingModule,
		NgxsModule.forFeature([
			ConsolidationsListState,
			ConsolidationEditState,
		]),
		MaterialModule,
		CommonModule,
		FormsModule,
		ReactiveFormsModule,
		DvoCardComponent,
		NgxsFormPluginModule,
		NgxsFormErrorsPluginModule,
		CdkMenuItem,
		CdkMenu,
		CdkMenuTrigger,
		DragDropModule,
		CdkMenuItemCheckbox,
		RouterOutlet,
		NgxsFormArrayPluginModule,
		OrderStatusColorPipePipe,
		CarrierServiceGrouperPipe,
		DeliveryOptionGrouperPipe,
		ShipmentStatusColorPipePipe,
		PdfViewerModule,
		PreparePdfPipe,
	],
	declarations: [
		ConsolidationsListComponent,
  AddConsolidationComponent,
  ConsolidationEditComponent,
  EditPalletComponent,
  CreateShipmentComponent,
	]
})
export class ConsolidationsModule { }
