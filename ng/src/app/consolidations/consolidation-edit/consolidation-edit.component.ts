import {Component, OnDestroy, OnInit} from '@angular/core';
import {Observable, Subscription} from "rxjs";
import {Paths} from "../../app-routing.module";
import {Actions, Store} from "@ngxs/store";
import {MatDialog} from "@angular/material/dialog";
import {AddConsolidationComponent} from "../consolidations-list/dialogs/add-consolidation/add-consolidation.component";
import {ConsolidationEditActions} from "./consolidation-edit.actions";
import FetchConsolidationEdit = ConsolidationEditActions.FetchConsolidationEdit;
import SetConsolidationID = ConsolidationEditActions.SetConsolidationID;
import {ActivatedRoute} from "@angular/router";
import {ConsolidationEditModel, ConsolidationEditState} from "./consolidation-edit.ngxs";
import Clear = ConsolidationEditActions.Clear;
import {FormControl, FormGroup} from "@angular/forms";
import {EditPalletComponent} from "./dialogs/edit-pallet/edit-pallet.component";
import EditPallet = ConsolidationEditActions.EditPallet;
import PalletResponse = ConsolidationEditActions.PalletResponse;
import SearchOrders = ConsolidationEditActions.SearchOrders;
import {CdkDragDrop} from "@angular/cdk/drag-drop";
import ListContainer = ConsolidationEditActions.ListContainer;
import ListContainerType = ConsolidationEditActions.ListContainerType;
import MoveOrder = ConsolidationEditActions.MoveOrder;
import OrderResponse = ConsolidationEditActions.OrderResponse;
import RemoveOrder = ConsolidationEditActions.RemoveOrder;
import Save = ConsolidationEditActions.Save;
import {AddressEditComponent} from "./dialogs/address-edit/address-edit.component";
import CreateShipment = ConsolidationEditActions.CreateShipment;
import {CreateShipmentComponent} from "./dialogs/create-shipment/create-shipment.component";
import ConsolidationShipment = ConsolidationEditActions.ConsolidationShipment;

@Component({
	selector: 'app-consolidation-edit',
	templateUrl: './consolidation-edit.component.html',
	styleUrl: './consolidation-edit.component.scss'
})
export class ConsolidationEditComponent implements OnInit, OnDestroy {

	form = new FormGroup({
		publicID: new FormControl(""),
		description: new FormControl(""),
		status: new FormControl(""),
		deliveryOption: new FormGroup({
			id: new FormControl(''),
		})
	});

	consolidationEdit$: Observable<ConsolidationEditModel>;
	subscriptions$: Subscription[] = [];
	paths = Paths;

	subscriptions: Subscription[] = [];

	searchOrders = new FormControl("", {nonNullable: true});
	isDragging = false;
	draggingID = "";

	containerTypes = ListContainerType;

	constructor(
		private route: ActivatedRoute,
		private store: Store,
		private dialog: MatDialog,
		private actions$: Actions,
	) {
		this.consolidationEdit$ = store.select(ConsolidationEditState.get);
	}

	ngOnInit(): void {
		this.subscriptions$.push(
			this.searchOrders.valueChanges
				.subscribe((v) => {
					this.store.dispatch(new SearchOrders(v));
				}));

		this.subscriptions$.push(this.route.queryParams
			.subscribe((params) => {
				this.store.dispatch([
					new SetConsolidationID(!!params.id ? params.id : ''),
					new FetchConsolidationEdit(),
				]);
			}));
	}

	ngOnDestroy(): void {
		this.subscriptions$.forEach((s) => s.unsubscribe());
		this.store.dispatch(new Clear());
	}

	add() {
		this.dialog.open(AddConsolidationComponent);
	}

	editPallet(val: PalletResponse | undefined) {
		this.store.dispatch(new EditPallet(val));
		this.dialog.open(EditPalletComponent)
	}

	editRecipient() {
		const ref = this.dialog.open(AddressEditComponent);
		ref.componentInstance.adr = this.store.selectSnapshot(ConsolidationEditState.get).recipient;
		ref.componentInstance.addressType = "recipient";
	}

	editSender() {
		const ref = this.dialog.open(AddressEditComponent);
		ref.componentInstance.adr = this.store.selectSnapshot(ConsolidationEditState.get).sender;
		ref.componentInstance.addressType = "sender";
	}

	startDrag(id: string) {
		this.draggingID = id;
		this.isDragging = true;
	}

	endDrag() {
		this.draggingID = "";
		this.isDragging = false;
	}

	drop(val: CdkDragDrop<ListContainer, ListContainer, OrderResponse>) {
		console.warn(val);
		this.store.dispatch(new MoveOrder({
			item: val.item.data,
			destination: val.container.data,
		}))
	}

	removeOrder(id: string) {
		this.store.dispatch(new RemoveOrder(id));
	}

	save() {
		this.store.dispatch(new Save());
	}

	createShipment(prebook: boolean) {
		this.dialog.open(CreateShipmentComponent);
		this.store.dispatch(new CreateShipment({prebook}));
	}

	hasShipment(palletID: string, allStatus: ConsolidationShipment): boolean {
		if (!allStatus.shipment || !allStatus.shipment.shipmentPallet) {
			return false;
		} else {
			return allStatus.shipment.shipmentPallet.some((sp) => {
				return sp.pallet?.id === palletID;
			})
		}
	}

}
