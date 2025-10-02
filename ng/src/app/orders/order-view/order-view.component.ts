import {Component, OnDestroy, OnInit} from '@angular/core';
import {ActivatedRoute} from "@angular/router";
import {Store} from "@ngxs/store";
import {Observable, Subscription} from "rxjs";
import {OrderViewModel, OrderViewState} from "./order-view.ngxs";
import {OrderViewActions} from "./order-view.actions";
import SetOrderID = OrderViewActions.SetOrderID;
import FetchOrder = OrderViewActions.FetchOrder;
import {AppActions} from "../../app.actions";
import AppGoBack = AppActions.AppGoBack;
import {Paths} from "../../app-routing.module";
import DuplicatePackage = OrderViewActions.DuplicatePackage;
import SetIsDragging = OrderViewActions.SetIsDragging;
import DeletePackage = OrderViewActions.DeletePackage;
import {MatDialog} from "@angular/material/dialog";
import {CreateEditOrderDialogComponent} from "./dialogs/create-edit-order-dialog/create-edit-order-dialog.component";
import Clear = OrderViewActions.Clear;
import CreateShipments = OrderViewActions.CreateShipments;
import {CreateShipmentComponent} from "./dialogs/create-shipment/create-shipment.component";
import FetchShipmentLabels = OrderViewActions.FetchShipmentLabels;
import {SignatureViewerComponent} from "./dialogs/signature-viewer/signature-viewer.component";
import {PackingSlipsComponent} from "./dialogs/packing-slips/packing-slips.component";
import FetchPackingSlips = OrderViewActions.FetchPackingSlips;
import {toNotNullArray} from "../../functions/not-null-array";
import SaveOrder = OrderViewActions.SaveOrder;
import {AppModel, AppState} from "../../app.ngxs";
import PackingSlipsClearCache = OrderViewActions.PackingSlipsClearCache;

@Component({
	selector: 'app-order-view',
	templateUrl: './order-view.component.html',
	styleUrls: ['./order-view.component.scss']
})
export class OrderViewComponent implements OnInit, OnDestroy {

	app$: Observable<AppModel>;
	orderView$: Observable<OrderViewModel>;

	displayedColumns: string[] = [
		'units',
		'description',
		'unitPrice',
		'weight',
	];

	editPath = Paths.ORDERS_PACKAGE_EDIT;
	subscriptions: Subscription[] = [];

	constructor(
		private route: ActivatedRoute,
		private store: Store,
		private dialog: MatDialog,
	) {
		this.orderView$ = store.select(OrderViewState.get);
		this.app$ = store.select(AppState.get);
	}

	ngOnInit(): void {
		this.subscriptions.push(this.route.queryParams
			.subscribe((params) => {
				this.store.dispatch([
					new SetOrderID(!!params.id ? params.id : ''),
					new FetchOrder(),
				]);
			}));
	}

	goBack() {
		this.store.dispatch(new AppGoBack());
	}

	edit(id: string) {
		console.warn(id);
	}

	addPackage() {
		const state = this.store.selectSnapshot(OrderViewState.get);
		if (!!state.order?.colli) {
			const fromColliID = state.order?.colli[0]?.id || '';
			if (!!fromColliID) {
				this.store.dispatch(new DuplicatePackage({fromColliID}));
			}
		}
	}

	isDragging(val: boolean) {
		this.store.dispatch(new SetIsDragging(val));
	}

	moveOrderLine(colliID: string, orderLineID: string) {
		this.store.dispatch(new OrderViewActions.FireMoveOrderLine({colliID, orderLineID}));
	}

	deletePackage(colliID: string) {
		this.store.dispatch(new DeletePackage({colliID}));
	}

	editOrder() {
		const state = this.store.selectSnapshot(OrderViewState.get);
		const ref = this.dialog.open(CreateEditOrderDialogComponent);
		ref.componentInstance.connections = state.connections
		ref.componentInstance.orderPublicID = state.order?.orderPublicID || ""
		ref.componentInstance.connectionID = state.order?.connection.id || ""
		ref.componentInstance.commentInternal = state.order?.commentInternal || ""
		ref.componentInstance.commentExternal = state.order?.commentExternal || ""
		this.subscriptions.push(ref.componentInstance.saveEmit.subscribe((c) => {
			this.store.dispatch(new SaveOrder(c.input))
		}));
	}

	ngOnDestroy(): void {
		this.subscriptions.forEach((s) => s.unsubscribe());
		this.store.dispatch(new Clear());
	}

	createShipment(id: string) {
		this.dialog.open(CreateShipmentComponent);
		this.store.dispatch(new CreateShipments({parcelIDs: [id]}));
	}

	viewShipment(colliID: string) {
		const ref = this.dialog.open(CreateShipmentComponent);
		ref.componentInstance.displayType = "view";
		this.store.dispatch(new FetchShipmentLabels([colliID]));
	}

	viewSignatures(colliID: string) {
		const ref = this.dialog.open(SignatureViewerComponent);
		ref.componentInstance.colliID = colliID;
	}

	packingSlips() {
		const state = this.store.selectSnapshot(OrderViewState.get);
		this.store.dispatch(new FetchPackingSlips({parcelIDs: toNotNullArray(state.order?.colli?.map((c) => c.id))}));
		this.dialog.open(PackingSlipsComponent);
	}

	packingSlipsClearCache() {
		const state = this.store.selectSnapshot(OrderViewState.get);
		this.store.dispatch(new PackingSlipsClearCache({orderIDs: [state.orderID]}));
	}

}
