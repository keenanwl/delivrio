import {AfterViewInit, ChangeDetectorRef, Component, OnDestroy, OnInit, ViewChild} from '@angular/core';
import {Actions, ofActionSuccessful, Store} from "@ngxs/store";
import {OrdersActions} from "./orders.actions";
import {Observable, Subscription, timer} from "rxjs";
import {OrdersModel, OrdersState, WhereChipOption} from "./orders.ngxs";
import {FormControl} from "@angular/forms";
import {MatAutocompleteTrigger} from "@angular/material/autocomplete";
import {MatPaginator, PageEvent} from "@angular/material/paginator";
import {AppActions} from "../app.actions";
import {Paths} from "../app-routing.module";
import {MatDialog, MatDialogRef} from "@angular/material/dialog";
import {
	CreateEditOrderDialogComponent
} from "./order-view/dialogs/create-edit-order-dialog/create-edit-order-dialog.component";
import {TableColumnSelectionComponent} from "./dialogs/table-column-selection/table-column-selection.component";
import {MatSort} from '@angular/material/sort';
import {OrderDirection} from "../../generated/graphql";
import {SelectPackagingComponent} from "../shared/select-packaging/select-packaging.component";
import {PdfViewerDialogComponent} from "../shared/pdf-viewer-dialog/pdf-viewer-dialog.component";
import SelectWhere = OrdersActions.SelectWhere;
import WhereChipClicked = OrdersActions.WhereChipClicked;
import WhereChipRemove = OrdersActions.WhereChipRemove;
import WhereChipRemoveAll = OrdersActions.WhereChipRemoveAll;
import SelectWhereTop = OrdersActions.SelectWhereTop;
import OrderRowsToggleAll = OrdersActions.OrderRowsToggleAll;
import OrderRowsToggleRows = OrdersActions.OrderRowsToggleRows;
import NextPage = OrdersActions.NextPage;
import PreviousPage = OrdersActions.PreviousPage;
import AppChangeRoute = AppActions.AppChangeRoute;
import ResetState = OrdersActions.ResetState;
import CreateNewOrder = OrdersActions.CreateNewOrder;
import ShowHideColumn = OrdersActions.ShowHideColumn;
import ChangeSortBy = OrdersActions.ChangeSortBy;
import BulkUpdatePackaging = OrdersActions.BulkUpdatePackaging;
import SetPackingSlips = OrdersActions.SetPackingSlips;
import BulkFetchPackingSlips = OrdersActions.BulkFetchPackingSlips;
import LocalFilterWhere = OrdersActions.LocalFilterWhere;
import CreatePackingSlipPrintJobs = OrdersActions.CreatePackingSlipPrintJobs;

@Component({
	selector: 'app-orders',
	templateUrl: './orders.component.html',
	styleUrls: ['./orders.component.scss']
})
export class OrdersComponent implements OnInit, OnDestroy, AfterViewInit {

	@ViewChild(MatPaginator) paginator: MatPaginator | undefined;
	@ViewChild("autoInput", {read: MatAutocompleteTrigger}) autocomplete: MatAutocompleteTrigger | undefined;
	@ViewChild(MatSort) sort: MatSort | undefined;

	filterControl = new FormControl('');
	orders$: Observable<OrdersModel>;
	openDialog: MatDialogRef<any> | null = null;

	displayedColumns: string[] = [
		'orderNumber',
		'creation',
		'connection',
		'recipient',
		'status',
		'shipments',
		'total',
		"country",
		"deliveryOption",
		'select',
	];

	subscriptions: Subscription[] = [];

	constructor(
		private store: Store,
		private dialog: MatDialog,
		private actions$: Actions,
	) {
		this.orders$ = store.select(OrdersState.state);
		if (!!this.paginator) {
			this.paginator.pageSize = 1;
		}
	}

	ngOnInit(): void {
		this.store.dispatch(new OrdersActions.FetchOrders());

		this.subscriptions.push(this.filterControl.valueChanges.subscribe((val) => {
			this.store.dispatch(new LocalFilterWhere(val || ''));
		}))
	}

	ngAfterViewInit() {
		this.subscriptions.push(this.sort!.sortChange.subscribe((s) => {
			this.store.dispatch(new ChangeSortBy(s.direction === 'asc' ? OrderDirection.Asc : OrderDirection.Desc));
		}));
	}

	ngOnDestroy() {
		this.subscriptions.map((s) => s.unsubscribe());
		this.store.dispatch(new ResetState());
		this.openDialog?.close();
	}

	isAllSelected(selected: {[key: string]: boolean}, all: OrdersActions.FetchOrdersQueryResponse[]) {
		if (all.length === 0) {
			return false;
		}
		return Object.keys(selected).length === all.length;
	}

	masterToggle() {
		this.store.dispatch(new OrderRowsToggleAll());
	}

	toggleRow(row: OrdersActions.FetchOrdersQueryResponse) {
		this.store.dispatch(new OrderRowsToggleRows([row]));
	}

	toChip() {
		this.filterControl.setValue('')
	}

	dropDownSelected(filterName: string, e: WhereChipOption) {
		this.filterControl.setValue(``)
		if (filterName === "top") {
			this.store.dispatch(new SelectWhereTop(e));
		} else {
			this.store.dispatch(new SelectWhere(e));
		}

		this.subscriptions.push(timer(0).subscribe(() => this.autocomplete?.openPanel()));
	}

	chipClickSelect(name: string) {
		this.store.dispatch(new WhereChipClicked(name));
		this.subscriptions.push(timer(0).subscribe(() => this.autocomplete?.openPanel()));
	}

	removeChip(name: string) {
		this.store.dispatch(new WhereChipRemove(name));
	}

	removeAllChips() {
		this.store.dispatch(new WhereChipRemoveAll());
	}

	selectPackaging() {
		const ref = this.dialog.open(SelectPackagingComponent)
		this.openDialog = ref;
		ref.componentInstance.helpText = "Bulk select packing for all selected collis. Dispatched collis will be ignored."
		this.subscriptions.push(ref.componentInstance.selected.subscribe((p) => {
			this.store.dispatch(new BulkUpdatePackaging(p?.id || null));
		}));
	}

	movePage(event: PageEvent) {
		const indexDiff = event.pageIndex - (event.previousPageIndex || 0);
		if (indexDiff === 1) {
			this.store.dispatch(new NextPage());
		} else {
			this.store.dispatch(new PreviousPage());
		}
	}

	edit(id: string) {
		this.store.dispatch(new AppChangeRoute({path: Paths.ORDERS_VIEW, queryParams: {id}}))
	}

	newOrder() {
		const state = this.store.selectSnapshot(OrdersState.state);
		const ref = this.dialog.open(CreateEditOrderDialogComponent);
		this.openDialog = ref;
		ref.componentInstance.isEdit = false;
		ref.componentInstance.connections = state.connections;
		ref.componentInstance.senderLocations = state.senderLocations;
		this.subscriptions.push(
			ref.componentInstance.saveEmit.subscribe((r) => {
				this.store.dispatch(new CreateNewOrder({
					input: r.input,
				}));
			})
		);
	}

	selectColumns() {
		const currentColumns = this.store.selectSnapshot(OrdersState.state).displayedColumns;
		const ref = this.dialog.open(TableColumnSelectionComponent);
		this.openDialog = ref;
		ref.componentInstance.availableColumns = this.displayedColumns;
		ref.componentInstance.selectedColumns = currentColumns;
		this.subscriptions.push(ref.componentInstance.nextSelectedColumns.subscribe((c) => {
			this.store.dispatch(new ShowHideColumn(c));
		}));

	}

	disableBulk(selectedRows: {[key: number]: string}) {
		return Object.keys(selectedRows).length === 0
	}

	bulkFetchPackingSlips() {
		this.store.dispatch(new BulkFetchPackingSlips());
		this.openDialog = this.dialog.open(PdfViewerDialogComponent);
		this.openDialog.componentInstance.loading = true;

		const action = this.actions$.pipe(ofActionSuccessful(SetPackingSlips))
			.subscribe(({payload}) => {

				// Weird change detection issue when we don't reset
				// the dialog from above. Zoneless related?
				this.openDialog?.close();
				const ref = this.dialog.open(PdfViewerDialogComponent);
				this.openDialog = ref;
				this.subscriptions.push(ref.componentInstance.pringBtn.subscribe(() => {
					console.warn("hhh")
					this.store.dispatch(new CreatePackingSlipPrintJobs())
				}))

				if (!!ref) {
					ref.componentInstance.allPDFs = payload.allPackingSlips;
					ref.componentInstance.labelsPDF = [...payload.packingSlips];
					ref.componentInstance.loading = false;
				}
			});
		this.subscriptions.push(action)

		this.subscriptions.push(this.openDialog.beforeClosed().subscribe(() => {
			action.unsubscribe();
		}));

	}

	protected readonly Object = Object;
}
