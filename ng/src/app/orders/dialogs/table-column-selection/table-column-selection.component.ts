import {Component, EventEmitter, Input, Output} from '@angular/core';
import {MatListOption} from "@angular/material/list";
import {MatDialogRef} from "@angular/material/dialog";

@Component({
	selector: 'app-table-column-selection',
	templateUrl: './table-column-selection.component.html',
	styleUrls: ['./table-column-selection.component.scss']
})
export class TableColumnSelectionComponent {

	@Input() availableColumns: string[] = [];
	@Input() selectedColumns: string[] = [];
	@Output() nextSelectedColumns: EventEmitter<string[]> = new EventEmitter<string[]>();

	constructor(private ref: MatDialogRef<any>) {
	}

	changeSelection(val: MatListOption[]) {
		this.nextSelectedColumns.emit(val.map(c => c.value));
	}

	close() {
		this.ref.close();
	}

}
