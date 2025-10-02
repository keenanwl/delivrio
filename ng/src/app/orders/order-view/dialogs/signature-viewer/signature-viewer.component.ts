import {ChangeDetectionStrategy, ChangeDetectorRef, Component, Input} from '@angular/core';
import {OrderViewState} from "../../order-view.ngxs";
import {Store} from "@ngxs/store";
import {MatDialogRef} from "@angular/material/dialog";

@Component({
	selector: 'app-signature-viewer',
	templateUrl: './signature-viewer.component.html',
	styleUrls: ['./signature-viewer.component.scss'],
	changeDetection: ChangeDetectionStrategy.OnPush,
})
export class SignatureViewerComponent {

	@Input() colliID = "";

	rotateDeg = 0;

	constructor(
		private store: Store,
		private dialog: MatDialogRef<any>,
		private cd: ChangeDetectorRef,
	) {}

	getImageURLs(colliID: string): string[] {

		const ship = this.store.selectSnapshot(OrderViewState.get).shipmentStatuses;
		let output: string[] = [];
		ship.shipmentStatuses.some((s) => {
			if (s.colliID == colliID) {
				output = s.ccSignatures;
				return true;
			}
			return false;
		});

		return output;

	}

	close() {
		this.dialog.close();
	}

	rotate() {

		if (this.rotateDeg < 280) {
			this.rotateDeg += 90;
		} else {
			this.rotateDeg = 0;
		}

		this.cd.detectChanges();

	}

}
