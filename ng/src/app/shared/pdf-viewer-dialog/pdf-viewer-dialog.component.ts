import {Component, ContentChild, EventEmitter, Inject, Input, Output} from '@angular/core';
import {MAT_DIALOG_DATA, MatDialogRef} from "@angular/material/dialog";

@Component({
	selector: 'app-pdf-viewer-dialog',
	templateUrl: './pdf-viewer-dialog.component.html',
	styleUrls: ['./pdf-viewer-dialog.component.scss']
})
export class PdfViewerDialogComponent {

	@Input() title = "";
	@Input() loading = false;
	@Input() loadingMessage = "";
	@Input() labelsPDF: string[] = [];
	@Input() errors: string[] = [];
	@Input() allPDFs: string = "";

	@Output() pringBtn = new EventEmitter<void>();

	labelViewerOffset = 0;

	constructor(@Inject(MAT_DIALOG_DATA) public data: any, private ref: MatDialogRef<any>) {
	}

	increment() {
		if (this.labelViewerOffset + 1 < this.labelsPDF.length) {
			this.labelViewerOffset += 1;
		}
	}

	decrement() {
		if (this.labelViewerOffset - 1 >= 0) {
			this.labelViewerOffset -= 1;
		}
	}

	print() {
		this.pringBtn.emit();
		this.ref.close();
	}

	close() {
		this.ref.close();
	}

}
