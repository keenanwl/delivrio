import {Component, Input, ViewEncapsulation} from '@angular/core';
import {MatDialogRef} from "@angular/material/dialog";

@Component({
	encapsulation: ViewEncapsulation.ShadowDom,
	selector: 'app-return-frame-error-dialog',
	templateUrl: './return-frame-error-dialog.component.html',
	styleUrls: ['./return-frame-error-dialog.component.scss']
})
export class ReturnFrameErrorDialogComponent {

	@Input() title = "";
	@Input() body = "";

	constructor(private ref: MatDialogRef<any>) {
	}

	close() {
		this.ref.close();
	}

}
