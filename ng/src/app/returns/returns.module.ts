import {NgModule} from "@angular/core";
import {NgxsFormPluginModule} from "@ngxs/form-plugin";
import {NgxsFormErrorsPluginModule} from "../plugins/ngxs-form-errors/ngxs-form-errors.module";
import {ReturnEditComponent} from './return-edit/return-edit.component';
import {ReturnsRoutingModule} from "./returns-routing.module";
import {NgxsModule} from "@ngxs/store";
import {MaterialModule} from "../modules/material.module";
import {CommonModule} from "@angular/common";
import {FormsModule, ReactiveFormsModule} from "@angular/forms";
import {DvoCardComponent} from "../shared/dvo-card/dvo-card.component";
import {ReturnsComponent} from "./returns.component";
import {ReturnsListComponent} from "./returns-list/returns-list.component";
import {ReturnsListState} from "./returns-list/returns-list.ngxs";
import {ReturnEditState} from "./return-edit/return-edit.ngxs";
import {ReturnViewComponent} from './return-view/return-view.component';
import {ReturnViewState} from "./return-view/return-view.ngxs";
import {NewReturnDialogComponent} from './returns-list/dialogs/new-return-dialog/new-return-dialog.component';
import {
	ReturnPortalSelectItemsModule
} from "../settings/return-portal-viewer/return-portal-frame/return-portal-select-items/return-portal-select-items.module";
import {ColliViewModule} from "../orders/order-view/colli-view/colli-view.module";
import {OrderLinesModule} from "../shared/order-lines/order-lines.module";
import {CdkMenu, CdkMenuItem, CdkMenuTrigger} from "@angular/cdk/menu";
import {ConfirmDeleteReturnColliComponent} from './returns-list/dialogs/confirm-delete-return-colli/confirm-delete-return-colli.component';
import {PdfViewerDialogModule} from "../shared/pdf-viewer-dialog/pdf-viewer-dialog.module";
import {TimelineViewerComponent} from "../shared/timeline-viewer/timeline-viewer.component";

@NgModule({
    imports: [
        ReturnsRoutingModule,
        NgxsModule.forFeature([
            ReturnsListState,
            ReturnEditState,
            ReturnViewState,
        ]),
        MaterialModule,
        CommonModule,
        FormsModule,
        ReactiveFormsModule,
        DvoCardComponent,
        NgxsFormPluginModule,
        NgxsFormErrorsPluginModule,
        ReturnPortalSelectItemsModule,
        ColliViewModule,
        OrderLinesModule,
        CdkMenuItem,
        CdkMenu,
        CdkMenuTrigger,
        PdfViewerDialogModule,
        TimelineViewerComponent,
    ],
	declarations: [
		ReturnsComponent,
		ReturnEditComponent,
		ReturnsListComponent,
		ReturnViewComponent,
		NewReturnDialogComponent,
        ConfirmDeleteReturnColliComponent,
	]
})
export class ReturnsModule { }
