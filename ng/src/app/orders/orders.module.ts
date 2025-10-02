import {NgModule} from "@angular/core";
import {MaterialModule} from "../modules/material.module";
import {CommonModule, NgOptimizedImage} from "@angular/common";
import {DvoCardComponent} from "../shared/dvo-card/dvo-card.component";
import {FormsModule, ReactiveFormsModule} from "@angular/forms";
import {OrdersComponent} from "./orders.component";
import {OrderEditComponent} from "./order-edit/order-edit.component";
import {OrderViewComponent} from "./order-view/order-view.component";
import {OrdersRoutingModule} from "./orders-routing.module";
import {NgxsModule} from "@ngxs/store";
import {OrdersState} from "./orders.ngxs";
import {OrderEditState} from "./order-edit/order-edit.ngxs";
import {OrderViewState} from "./order-view/order-view.ngxs";
import {NgxsFormPluginModule} from "@ngxs/form-plugin";
import {NgxsFormErrorsPluginModule} from "../plugins/ngxs-form-errors/ngxs-form-errors.module";
import {OrderLinesModule} from "../shared/order-lines/order-lines.module";
import {AddOrderLineComponent} from './order-edit/add-order-line/add-order-line.component';
import {A11yModule} from "@angular/cdk/a11y";
import {CdkMenu, CdkMenuItem, CdkMenuItemCheckbox, CdkMenuTrigger} from "@angular/cdk/menu";
import {DragDropModule} from "@angular/cdk/drag-drop";
import {CreateEditOrderDialogComponent} from './order-view/dialogs/create-edit-order-dialog/create-edit-order-dialog.component';
import {TableColumnSelectionComponent} from './dialogs/table-column-selection/table-column-selection.component';
import {CreateShipmentComponent} from './order-view/dialogs/create-shipment/create-shipment.component';
import {PdfViewerModule} from "ng2-pdf-viewer";
import {PreparePdfPipe} from './order-view/pipes/prepare-pdf.pipe';
import {MayCreateShipmentPipe} from './order-view/pipes/may-create-shipment.pipe';
import {DeliveryPointRequiredPipe} from './order-edit/pipes/delivery-point-required.pipe';
import {DeliveryPointOptionalPipe} from './order-edit/pipes/delivery-point-optional.pipe';
import {EditDeliveryPointComponent} from './order-edit/dialogs/edit-delivery-point/edit-delivery-point.component';
import {ColliViewModule} from "./order-view/colli-view/colli-view.module";
import {SignatureViewerComponent} from "./order-view/dialogs/signature-viewer/signature-viewer.component";
import {EditCcLocationComponent} from './order-edit/dialogs/edit-cc-location/edit-cc-location.component';
import {PackingSlipsComponent} from './order-view/dialogs/packing-slips/packing-slips.component';
import {RelativeTimePipe} from "../pipes/relative-time.pipe";
import {ColliStatusColorPipePipe} from "../shared/colli-status-color-pipe.pipe";
import {OrderStatusColorPipePipe} from "../shared/order-status-color-pipe.pipe";
import {TimelineViewerComponent} from "../shared/timeline-viewer/timeline-viewer.component";
import {ShipmentStatusColorPipePipe} from "../shared/shipment-status-color-pipe.pipe";
import {TotalOrderValuePipe} from "./pipes/total-order-value.pipe";
import {ShipmentParcelStatusColorPipePipe} from "../shared/shipment-parcel-status-color-pipe.pipe";

@NgModule({
	imports: [
		OrdersRoutingModule,
		NgxsModule.forFeature([
			OrdersState,
			OrderEditState,
			OrderViewState,
		]),
		MaterialModule,
		CommonModule,
		FormsModule,
		ReactiveFormsModule,
		DvoCardComponent,
		NgxsFormPluginModule,
		NgxsFormErrorsPluginModule,
		OrderLinesModule,
		A11yModule,
		CdkMenuItem,
		CdkMenu,
		CdkMenuTrigger,
		DragDropModule,
		CdkMenuItemCheckbox,
		PdfViewerModule,
		ColliViewModule,
		PreparePdfPipe,
		RelativeTimePipe,
		ColliStatusColorPipePipe,
		OrderStatusColorPipePipe,
		TimelineViewerComponent,
		NgOptimizedImage,
		ShipmentStatusColorPipePipe,
		TotalOrderValuePipe,
		ShipmentParcelStatusColorPipePipe,
	],
	declarations: [
		OrdersComponent,
		OrderEditComponent,
		OrderViewComponent,
		AddOrderLineComponent,
		CreateEditOrderDialogComponent,
		TableColumnSelectionComponent,
		CreateShipmentComponent,
		MayCreateShipmentPipe,
		DeliveryPointRequiredPipe,
		DeliveryPointOptionalPipe,
		EditDeliveryPointComponent,
		SignatureViewerComponent,
		EditCcLocationComponent,
		PackingSlipsComponent,
	]
})
export class OrdersModule { }
