import {Component, EventEmitter, Input, Output, ViewEncapsulation} from '@angular/core';
import {Item, ReturnPortalViewResponse} from "../return-portal-frame.service";
import {ItemReturn} from "../return-portal-frame.ngxs";

@Component({
	encapsulation: ViewEncapsulation.ShadowDom,
	selector: 'app-return-portal-select-items',
	templateUrl: './return-portal-select-items.component.html',
	styleUrls: ['./return-portal-select-items.component.scss']
})
export class ReturnPortalSelectItemsComponent {

	@Input() returnOrder: ReturnPortalViewResponse | null = null;
	@Input() selectedItems: ItemReturn[] = [];
	@Output() toggleEvt = new EventEmitter<{item: Item, selected: boolean}>();
	@Output() incrementEvt = new EventEmitter<{orderLineID: string}>();
	@Output() decrementEvt = new EventEmitter<{orderLineID: string}>();
	@Output() reasonChangedEvt = new EventEmitter<{item: Item, reasonID: string}>();

	isSelected(selectedItems: ItemReturn[], orderLineID: string): boolean {
		const sel = selectedItems.find((i) => i.orderLineID === orderLineID && i.selected) !== undefined;
		return sel;
	}

	toggle(evt: MouseEvent, item: Item, selected: boolean) {
		/*evt.preventDefault();
		evt.stopPropagation();*/
		this.toggleEvt.next({item, selected});
	}

	totalCount(selectedItems: ItemReturn[]): number {
		let count = 0;
		selectedItems.forEach((i) => {
			if (i.selected) {
				count += i.quantity;
			}
		})
		return count;
	}

	selectedQuantity(selectedItems: ItemReturn[], idToCheck: string): number {
		let count = 0
		selectedItems.some((i) => {
			if (i.orderLineID == idToCheck) {
				count = i.quantity;
				return true;
			}
			return false;
		});
		return count;
	}

	mayAdjustQuantity(selectedItems: ItemReturn[], orderLineID: string): boolean {
		return selectedItems.some((i) => {
			if (i.orderLineID === orderLineID && i.selected) {
				return true;
			}
			return false;
		});
	}

	increment(orderLineID: string) {
		this.incrementEvt.next({orderLineID});
	}

	decrement(orderLineID: string) {
		this.decrementEvt.next({orderLineID});
	}

	reasonChanged(item: Item, reasonID: string) {
		//this.toggleContainer(orderLineID);
		this.reasonChangedEvt.next({item, reasonID});
	}

}
