import {Component, EventEmitter, Output} from '@angular/core';
import {MatDialogRef} from "@angular/material/dialog";

@Component({
  selector: 'app-confirm-delete-return-colli',
  templateUrl: './confirm-delete-return-colli.component.html',
  styleUrls: ['./confirm-delete-return-colli.component.scss']
})
export class ConfirmDeleteReturnColliComponent {

	@Output() confirm = new EventEmitter<boolean>();

	constructor(private ref: MatDialogRef<any>) {
	}

	confirmButton() {
		this.confirm.emit(true);
		this.close();
	}

	close() {
		this.ref.close();
	}

}
