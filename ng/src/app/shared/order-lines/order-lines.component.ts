import {Component, EventEmitter, Input, OnInit, Output} from '@angular/core';
import {OrderEditActions} from "../../orders/order-edit/order-edit.actions";
import OrderLineResponse = OrderEditActions.OrderLineResponse;
import {Paths} from "../../app-routing.module";

export type PriceEdited = {
	index: RowIndex;
	unitPrice: number;
}

export type UnitsEdited = {
	index: RowIndex;
	units: number;
}

export type DiscountEdited = {
	index: RowIndex;
	amount: number;
}

export type RowIndex = number;

@Component({
	selector: 'app-order-lines',
	templateUrl: './order-lines.component.html',
	styleUrls: ['./order-lines.component.scss']
})
export class OrderLinesComponent implements OnInit {

	displayedColumns: string[] = [
		'units',
		'description',
		'unitPrice',
		'discountAmount',
		'weight',
		'actions',
		'drag',
	];

	@Input() editable = false;
	@Input() dragable = false;
	@Input() orderLines: OrderLineResponse[] = [];
	@Input() showDragBoundary: boolean = false;

	@Output() rowClicked = new EventEmitter<string>();

	// Different outputs since templateRefs aren't available across cells
	@Output() priceEdited = new EventEmitter<PriceEdited>();
	@Output() discountEdited = new EventEmitter<DiscountEdited>();
	@Output() unitsEdited = new EventEmitter<UnitsEdited>();
	@Output() lineDeleted = new EventEmitter<RowIndex>();
	@Output() isDragging = new EventEmitter<boolean>();

	productEditPath = Paths.PRODUCTS_EDIT;

	ngOnInit() {
		if (!this.editable) {
			this.displayedColumns = this.displayedColumns.filter((v) => v !== 'actions');
		}
		if (!this.dragable) {
			this.displayedColumns = this.displayedColumns.filter((v) => v !== 'drag');
		}
	}

	trackByRows = (index: number, row: any) => {
		return index
	}

	edit(id: string) {
		//console.warn(id);
	}

	priceUpdated(index: RowIndex, newPrice: string) {
		let nextPrice = parseFloat(newPrice) || 0;
		this.priceEdited.emit({index, unitPrice: nextPrice});
	}

	discountUpdated(index: RowIndex, newDiscount: string) {
		let nextDiscount = parseFloat(newDiscount) || 0;
		this.discountEdited.emit({index, amount: nextDiscount});
	}

	unitsUpdated(index: RowIndex, newUnits: string) {
		let nextUnit = parseFloat(newUnits) || 0;
		this.unitsEdited.emit({index, units: parseFloat(newUnits)});
	}

	lineDelete(index: RowIndex) {
		this.lineDeleted.emit(index);
	}

	dragStart() {
		this.isDragging.emit(true);
	}

	dragEnd() {
		this.isDragging.emit(false);
	}

}
