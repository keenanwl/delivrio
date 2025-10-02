import {Component, OnDestroy, OnInit} from '@angular/core';
import {Paths} from "../../app-routing.module";
import {ActivatedRoute} from "@angular/router";
import {Store} from "@ngxs/store";
import {Observable, Subscription} from "rxjs";
import {ShipmentViewModel, ShipmentViewState} from "./shipment-view.ngxs";
import {ShipmentViewActions} from "./shipment-view.actions";
import FetchShipment = ShipmentViewActions.FetchShipment;
import SetShipmentID = ShipmentViewActions.SetShipmentID;
import CancelShipment = ShipmentViewActions.CancelShipment;
import {AppActions} from "../../app.actions";
import AppChangeRoute = AppActions.AppChangeRoute;
import {MatDialog} from "@angular/material/dialog";
import CancelFulfillmentSync = ShipmentViewActions.CancelFulfillmentSync;
import CancelCancelSync = ShipmentViewActions.CancelCancelSync;
import {ShipmentStatus} from "../../../generated/graphql";
import Clear = ShipmentViewActions.Clear;

@Component({
	selector: 'app-shipment-view',
	templateUrl: './shipment-view.component.html',
	styleUrls: ['./shipment-view.component.scss']
})
export class ShipmentViewComponent implements OnInit, OnDestroy {

	shipmentView$: Observable<ShipmentViewModel>;

	displayedColumns: string[] = [
		'units',
		'description',
		'unitPrice',
		'weight',
	];

	subscriptions: Subscription[] = [];

	constructor(
		private route: ActivatedRoute,
		private store: Store,
		private dialog: MatDialog,
	) {
		this.shipmentView$ = store.select(ShipmentViewState.get);
	}

	ngOnDestroy() {
		this.store.dispatch(new Clear());
	}

	ngOnInit(): void {
		this.subscriptions.push(this.route.queryParams
			.subscribe((params) => {
				this.store.dispatch([
					new SetShipmentID(!!params.id ? params.id : ''),
					new FetchShipment(),
				]);
			}));
	}

	cancel() {
		this.store.dispatch(new CancelShipment());
	}

	edit(row: any) {

	}

	viewOrder(id: string) {
		this.store.dispatch(new AppChangeRoute({path: Paths.ORDERS_VIEW, queryParams: {id: id}}));
	}

	cancelFulfillmentSync(id: string) {
		this.store.dispatch(new CancelFulfillmentSync(id));
	}

	cancelCancelSync(id: string) {
		this.store.dispatch(new CancelCancelSync(id));
	}

	protected readonly ShipmentStatus = ShipmentStatus;
}
