import {NgModule} from "@angular/core";
import {CommonModule, NgOptimizedImage} from "@angular/common";
import {NgxsModule} from "@ngxs/store";
import {NgxsFormPluginModule} from "@ngxs/form-plugin";
import {FormsModule, ReactiveFormsModule} from "@angular/forms";
import {NgxsFormErrorsPluginModule} from "../../plugins/ngxs-form-errors/ngxs-form-errors.module";
import {MaterialModule} from "../../modules/material.module";
import {DvoCardComponent} from "../../shared/dvo-card/dvo-card.component";
import {NgxsFormArrayPluginModule} from "../../plugins/ngxs-form-array/ngxs-form-array.module";
import {DocumentsListComponent} from "./documents-list/documents-list.component";
import {DocumentsRoutingModule} from "./documents-routing.module";
import {DocumentEditComponent} from "./document-edit/document-edit.component";
import {DocumentsListState} from "./documents-list/documents-list.ngxs";
import {AddDocumentComponent} from './documents-list/dialogs/add-document/add-document.component';
import {RelativeTimePipe} from "../../pipes/relative-time.pipe";
import {DocumentEditState} from "./document-edit/document-edit.ngxs";
import {AlphabetizePipe} from "../../shared/alphabetize.pipe";
import {ViewDocumentComponent} from './document-edit/dialogs/view-document/view-document.component';
import {PdfViewerModule} from "ng2-pdf-viewer";
import {PreparePdfPipe} from "../../orders/order-view/pipes/prepare-pdf.pipe";

@NgModule({
    imports: [
        DocumentsRoutingModule,
        NgxsModule.forFeature([
            DocumentsListState,
            DocumentEditState,
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
        RelativeTimePipe,
        AlphabetizePipe,
        PdfViewerModule,
        PreparePdfPipe,
    ],
	providers: [
	],
	declarations: [
		DocumentsListComponent,
		DocumentEditComponent,
  		AddDocumentComponent,
    	ViewDocumentComponent,
	]
})
export class DocumentsModule { }
