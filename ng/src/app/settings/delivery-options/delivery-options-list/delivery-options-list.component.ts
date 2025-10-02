import {Component, OnDestroy, OnInit} from '@angular/core';
import {Observable, Subscription} from "rxjs";
import {Store} from "@ngxs/store";
import {DeliveryOptionsListActions} from "./delivery-options-list.actions";
import FetchDeliveryOptionsList = DeliveryOptionsListActions.FetchDeliveryOptionsList;
import {AppActions} from "../../../app.actions";
import AppChangeRoute = AppActions.AppChangeRoute;
import {DeliveryOptionsListModel, DeliveryOptionsListState} from "./delivery-options-list.ngxs";
import {MatDialog} from "@angular/material/dialog";
import {NewDeliveryOptionDialogComponent} from "./new-delivery-option-dialog.component";
import {Paths} from "../../../app-routing.module";
import {CarrierBrandInternalId} from "../../../../generated/graphql";
import Clear = DeliveryOptionsListActions.Clear;
import UpdateSortOrder = DeliveryOptionsListActions.UpdateSortOrder;
import {CdkDragDrop} from "@angular/cdk/drag-drop";
import SelectDeliveryOptionsListQueryResponse = DeliveryOptionsListActions.SelectDeliveryOptionsListQueryResponse;
import {ArchiveConfirmationComponent} from "./dialogs/archive-confirmation/archive-confirmation.component";
import ToggleShowArchive = DeliveryOptionsListActions.ToggleShowArchive;

@Component({
	selector: 'app-delivery-options-list',
	templateUrl: './delivery-options-list.component.html',
	styleUrls: ['./delivery-options-list.component.scss']
})
export class DeliveryOptionsListComponent implements OnInit, OnDestroy {

	deliveryOptionsList$: Observable<DeliveryOptionsListModel>;
	subscriptions$: Subscription[] = [];

	displayedColumns: string[] = [
		'logo',
		'carrierAgreement',
		'carrierService',
		'deliveryOptionName',
		'connection',
		'actions',
	];

	constructor(
		private store: Store,
		private dialog: MatDialog,
	) {
		this.deliveryOptionsList$ = store.select(DeliveryOptionsListState.get);
	}

	ngOnInit(): void {
		this.store.dispatch(new FetchDeliveryOptionsList());
	}

	ngOnDestroy() {
		this.subscriptions$.forEach((s) => s.unsubscribe());
		this.store.dispatch(new Clear());
	}

	editOption(id: string, internalID: CarrierBrandInternalId) {
		let path = Paths.SETTINGS_DELIVERY_OPTIONS_EDIT_GLS;
		switch (internalID) {
			case CarrierBrandInternalId.PostNord:
				path = Paths.SETTINGS_DELIVERY_OPTIONS_EDIT_POST_NORD
				break;
			case CarrierBrandInternalId.Usps:
				path = Paths.SETTINGS_DELIVERY_OPTIONS_EDIT_USPS
				break;
			case CarrierBrandInternalId.Dao:
				path = Paths.SETTINGS_DELIVERY_OPTIONS_EDIT_DAO
				break;
			case CarrierBrandInternalId.Df:
				path = Paths.SETTINGS_DELIVERY_OPTIONS_EDIT_DF
				break;
			case CarrierBrandInternalId.Dsv:
				path = Paths.SETTINGS_DELIVERY_OPTIONS_EDIT_DSV
				break;
			case CarrierBrandInternalId.EasyPost:
				path = Paths.SETTINGS_DELIVERY_OPTIONS_EDIT_EASY_POST
				break;
			case CarrierBrandInternalId.Bring:
				path = Paths.SETTINGS_DELIVERY_OPTIONS_EDIT_BRING
				break;
		}
		this.store.dispatch(new AppChangeRoute({path, queryParams: {id}}));
	}

	createNew() {
		const ref = this.dialog.open(NewDeliveryOptionDialogComponent);
	}

	drop(evt: CdkDragDrop<string>) {
		console.warn(evt);
		this.store.dispatch(new UpdateSortOrder({deliveryOptionID: evt.item.data, nextIndex: evt.currentIndex}));
	}

	toggleArchived() {
		this.store.dispatch(new ToggleShowArchive());
	}

	archive(event: Event, dopt: SelectDeliveryOptionsListQueryResponse) {
		event.stopPropagation()
		const ref = this.dialog.open(ArchiveConfirmationComponent);
		ref.componentRef!.instance.doptName = `${dopt.name} - ${dopt.carrier.name}`
		ref.componentRef!.instance.dOptID = dopt.id;
	}

}
