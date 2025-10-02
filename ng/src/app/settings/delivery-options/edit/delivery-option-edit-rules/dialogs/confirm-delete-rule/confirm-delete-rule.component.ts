import {Component, EventEmitter, Input, OnInit, Output} from '@angular/core';
import {MatDialogRef} from "@angular/material/dialog";

@Component({
  selector: 'app-confirm-delete-rule',
  templateUrl: './confirm-delete-rule.component.html',
  styleUrls: ['./confirm-delete-rule.component.scss']
})
export class ConfirmDeleteRuleComponent implements OnInit {

	@Input() name: string = "";
	@Output() confirmed = new EventEmitter<boolean>();

	constructor(private ref: MatDialogRef<any>) {
	}

	ngOnInit() {

	}

	confirm() {
		this.confirmed.emit(true);
		this.close();
	}

	close() {
		this.ref.close();
	}

}
