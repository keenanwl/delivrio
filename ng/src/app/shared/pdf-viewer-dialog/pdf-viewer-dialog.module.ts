import {NgModule} from "@angular/core";
import {CommonModule} from "@angular/common";
import {MaterialModule} from "../../modules/material.module";
import {PdfViewerDialogComponent} from "./pdf-viewer-dialog.component";
import {PdfViewerModule} from "ng2-pdf-viewer";
import {PreparePdfPipe} from "../../orders/order-view/pipes/prepare-pdf.pipe";

@NgModule({
	imports: [
		MaterialModule,
		CommonModule,
		PdfViewerModule,
		PreparePdfPipe,
	],
	exports: [
		PdfViewerDialogComponent,
	],
	declarations: [
		PdfViewerDialogComponent,
	]
})
export class PdfViewerDialogModule { }
