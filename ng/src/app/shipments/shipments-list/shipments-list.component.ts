import {Component, OnDestroy, OnInit} from '@angular/core';
import {Store} from '@ngxs/store';
import {Observable, Subscription} from "rxjs";
import {PageEvent} from "@angular/material/paginator";
import {ShipmentsListModel, ShipmentsListState} from './shipments-list.ngxs';
import {ShipmentsListActions} from "./shipments-list.actions";
import FetchShipments = ShipmentsListActions.FetchShipments;
import ToggleRows = ShipmentsListActions.ToggleRows;
import ToggleAll = ShipmentsListActions.ToggleAll;
import ShipmentsResponse = ShipmentsListActions.ShipmentsResponse;
import {Paths} from "../../app-routing.module";
import {AppActions} from "../../app.actions";
import AppChangeRoute = AppActions.AppChangeRoute;
import {activeFilter, selectedOptionList} from "../../shared/filter-bar/filter-bar.component";
import SearchFilterChanges = ShipmentsListActions.SearchFilterChanges;
import SetSelectedFilters = ShipmentsListActions.SetSelectedFilters;
import {MatDialog} from "@angular/material/dialog";
import {SendFilteredEmailComponent} from "./dialogs/send-filtered-email/send-filtered-email.component";

@Component({
	selector: 'app-shipments-list',
	templateUrl: './shipments-list.component.html',
	styleUrls: ['./shipments-list.component.scss']
})
export class ShipmentsListComponent implements OnInit, OnDestroy {

	shipments$: Observable<ShipmentsListModel>;
	paths = Paths;

	displayedColumns: string[] = [
		'createdAt',
		'status',
		'fulfillmentSync',
		'cancelSync',
		'collis',
		'pallets',
		'select',
	];

	subscriptions: Subscription[] = [];

	constructor(
		private store: Store,
		private dialog: MatDialog,
	) {
		this.shipments$ = store.select(ShipmentsListState.state);
	}

	ngOnInit(): void {
		this.store.dispatch(new FetchShipments());
	}

	ngOnDestroy() {
		/*this.subscriptions.map((s) => s.unsubscribe());
		this.store.dispatch(new ResetState());*/
	}

	isAllSelected(selected: Set<string>, all: any) {
		return selected.size === all.length;
	}

	masterToggle() {
		this.store.dispatch(new ToggleAll());
	}

	toggleRow(row: ShipmentsResponse) {
		this.store.dispatch(new ToggleRows([row]));
	}

	showEmailDialog() {
		this.dialog.open(SendFilteredEmailComponent);
	}

	movePage(event: PageEvent) {
		/*const indexDiff = event.pageIndex - (event.previousPageIndex || 0);
		if (indexDiff === 1) {
			this.store.dispatch(new NextPage());
		} else {
			this.store.dispatch(new PreviousPage());
		}*/
	}

	edit(id: string) {
		this.store.dispatch(new AppChangeRoute({path: Paths.SHIPMENT_VIEW, queryParams: {id}}))
	}

	goToOrder(orderID: string) {
		this.store.dispatch(new AppChangeRoute({path: Paths.ORDERS_VIEW, queryParams: {id: orderID}}))
	}

	searchFilter(val: activeFilter) {
		this.store.dispatch(new SearchFilterChanges(val));
	}

	selectedOptionsChanged(val: selectedOptionList) {
		this.store.dispatch(new SetSelectedFilters(val));
	}

}
