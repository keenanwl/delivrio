import {Component, OnDestroy, OnInit} from '@angular/core';
import {AppActions} from "../../../app.actions";
import AppChangeRoute = AppActions.AppChangeRoute;
import {Store} from "@ngxs/store";
import {Observable, Subscription} from "rxjs";
import {CarriersListModel, CarriersListState} from "./carriers-list.ngxs";
import {CarriersListActions} from "./carriers-list.actions";
import FetchCarriersList = CarriersListActions.FetchCarriersList;
import {MatDialog} from "@angular/material/dialog";
import {NewCarrierAgreementDialogComponent} from "./new-carrier-agreement-dialog.component";
import {Paths} from "../../../app-routing.module";
import {CarrierBrandInternalId} from "../../../../generated/graphql";
import Clear = CarriersListActions.Clear;

@Component({
	selector: 'app-carriers-list',
	templateUrl: './carriers-list.component.html',
	styleUrls: ['./carriers-list.component.scss']
})
export class CarriersListComponent implements OnInit, OnDestroy {

	carrierList$: Observable<CarriersListModel>;
	subscriptions$: Subscription[] = [];

	constructor(
		private store: Store,
		private dialog: MatDialog,
	) {
		this.carrierList$ = store.select(CarriersListState.get);
	}

	ngOnInit(): void {
		this.store.dispatch(new FetchCarriersList());
	}

	editCarrier(id: string, internalID: CarrierBrandInternalId) {
		switch (internalID) {
			case CarrierBrandInternalId.Bring:
				this.store.dispatch(new AppChangeRoute({path: Paths.SETTINGS_CARRIERS_EDIT_BRING, queryParams: {id}}));
				break;
			case CarrierBrandInternalId.Dao:
				this.store.dispatch(new AppChangeRoute({path: Paths.SETTINGS_CARRIERS_EDIT_DAO, queryParams: {id}}));
				break;
			case CarrierBrandInternalId.Df:
				this.store.dispatch(new AppChangeRoute({path: Paths.SETTINGS_CARRIERS_EDIT_DF, queryParams: {id}}));
				break;
			case CarrierBrandInternalId.Dsv:
				this.store.dispatch(new AppChangeRoute({path: Paths.SETTINGS_CARRIERS_EDIT_DSV, queryParams: {id}}));
				break;
			case CarrierBrandInternalId.EasyPost:
				this.store.dispatch(new AppChangeRoute({path: Paths.SETTINGS_CARRIERS_EDIT_EASY_POST, queryParams: {id}}));
				break;
			case CarrierBrandInternalId.Gls:
				this.store.dispatch(new AppChangeRoute({path: Paths.SETTINGS_CARRIERS_EDIT_GLS, queryParams: {id}}));
				break;
			case CarrierBrandInternalId.PostNord:
				this.store.dispatch(new AppChangeRoute({path: Paths.SETTINGS_CARRIERS_EDIT_POST_NORD, queryParams: {id}}));
				break;
			case CarrierBrandInternalId.Usps:
				this.store.dispatch(new AppChangeRoute({path: Paths.SETTINGS_CARRIERS_EDIT_USPS, queryParams: {id}}));
				break;
			default:
				throw new Error("internal ID not recognized: " + internalID);
		}

	}

	addNew() {
		this.dialog.open(NewCarrierAgreementDialogComponent);
	}

	ngOnDestroy(): void {
		this.subscriptions$.forEach((s) => s.unsubscribe());
		this.store.dispatch(new Clear());
	}

}
