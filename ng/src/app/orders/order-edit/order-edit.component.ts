import {Component, OnDestroy, OnInit} from '@angular/core';
import {ActivatedRoute} from "@angular/router";
import {Actions, ofActionSuccessful, Store} from '@ngxs/store';
import {OrderEditActions} from "./order-edit.actions";
import FetchPackage = OrderEditActions.FetchPackage;
import SetColliID = OrderEditActions.SetColliID;
import {Observable, Subscription} from "rxjs";
import {OrderEditModel, OrderEditState} from "./order-edit.ngxs";
import {FormControl, FormGroup} from "@angular/forms";
import {debounceTime} from "rxjs/operators";
import SearchProducts = OrderEditActions.SearchProducts;
import RemoveOrderLine = OrderEditActions.RemoveOrderLine;
import SaveFormNew = OrderEditActions.SaveFormNew;
import SaveFormUpdate = OrderEditActions.SaveFormUpdate;
import OrderLineResponse = OrderEditActions.OrderLineResponse;
import ResetState = OrderEditActions.ResetState;
import {DeliveryOptionBrandNameStatus} from "../../../generated/graphql";
import SelectDeliveryOption = OrderEditActions.SelectDeliveryOption;
import {PriceEdited, RowIndex, UnitsEdited} from "../../shared/order-lines/order-lines.component";
import RowEditedUnits = OrderEditActions.RowEditedUnits;
import RowEditedPrice = OrderEditActions.RowEditedPrice;
import {MatDialog} from "@angular/material/dialog";
import {AddOrderLineComponent} from "./add-order-line/add-order-line.component";
import SearchCountry = OrderEditActions.SearchCountry;
import CountriesResponse = OrderEditActions.CountriesResponse;
import ChangeCountrySender = OrderEditActions.ChangeCountrySender;
import ChangeCountryRecipient = OrderEditActions.ChangeCountryRecipient;
import SetOrderID = OrderEditActions.SetOrderID;
import FetchDeliveryOptions = OrderEditActions.FetchDeliveryOptions;
import {EditDeliveryPointComponent} from "./dialogs/edit-delivery-point/edit-delivery-point.component";
import {UpdateFormValue} from "@ngxs/form-plugin";
import {EditCcLocationComponent} from "./dialogs/edit-cc-location/edit-cc-location.component";
import SetSelectedPackaging = OrderEditActions.SetSelectedPackaging;
import {SelectPackagingComponent} from "../../shared/select-packaging/select-packaging.component";

@Component({
	selector: 'app-order-edit',
	templateUrl: './order-edit.component.html',
	styleUrls: ['./order-edit.component.scss']
})
export class OrderEditComponent implements OnInit, OnDestroy {

	order$: Observable<OrderEditModel>;

	editForm = new FormGroup({
		// needs to be null to prevent foreign key issue
		order: new FormGroup({
			connection: new FormGroup({
				id: new FormControl('', {nonNullable: true}),
			}),
		}),
		deliveryOption: new FormGroup({
			id: new FormControl(null),
			clickCollect: new FormControl(false),
		}),
		recipient: new FormGroup({
			id: new FormControl<string>('', {nonNullable: true}),
			firstName: new FormControl<string>('', {nonNullable: true}),
			lastName: new FormControl<string>('', {nonNullable: true}),
			phoneNumber: new FormControl<string>('', {nonNullable: true}),
			email: new FormControl<string>('', {nonNullable: true}),
			addressOne: new FormControl<string>('', {nonNullable: true}),
			addressTwo: new FormControl<string>('', {nonNullable: true}),
			zip: new FormControl<string>('', {nonNullable: true}),
			city: new FormControl<string>('', {nonNullable: true}),
			state: new FormControl<string>('', {nonNullable: true}),
			country: new FormGroup({
				id: new FormControl('', {nonNullable: true}),
				label: new FormControl('', {nonNullable: true}),
				alpha2: new FormControl('', {nonNullable: true}),
			}),
			company: new FormControl<string>('', {nonNullable: true}),
		}),
		packaging: new FormGroup({
			id: new FormControl<string>('', {nonNullable: false}),
			name: new FormControl<string>('', {nonNullable: true}),
			lengthCm: new FormControl<string>('', {nonNullable: true}),
			heightCm: new FormControl<string>('', {nonNullable: true}),
			widthCm: new FormControl<string>('', {nonNullable: true}),
		}),
		sender: new FormGroup({
			firstName: new FormControl<string>('', {nonNullable: true}),
			lastName: new FormControl<string>('', {nonNullable: true}),
			phoneNumber: new FormControl<string>('', {nonNullable: true}),
			vatNumber: new FormControl<string>('', {nonNullable: true}),
			email: new FormControl<string>('', {nonNullable: true}),
			id: new FormControl<string>('', {nonNullable: true}),
			addressOne: new FormControl<string>('', {nonNullable: true}),
			addressTwo: new FormControl<string>('', {nonNullable: true}),
			zip: new FormControl<string>('', {nonNullable: true}),
			city: new FormControl<string>('', {nonNullable: true}),
			state: new FormControl<string>('', {nonNullable: true}),
			country: new FormGroup({
				id: new FormControl('', {nonNullable: true}),
				label: new FormControl('', {nonNullable: true}),
				alpha2: new FormControl('', {nonNullable: true}),
			}),
			company: new FormControl<string>('', {nonNullable: true}),
		}),
		// Not visible, but required so NGXS doesn't knock out the data
		orderLines: new FormControl<OrderLineResponse[]>([], {nonNullable: true}),
	});

	searchProductsCtrl = new FormControl('', {nonNullable: true});

	subscriptions: Subscription[] = [];

	constructor(
		private route: ActivatedRoute,
		private store: Store,
		private dialog: MatDialog,
		private actions$: Actions,
	) {
		this.order$ = store.select(OrderEditState.get);
	}

	ngOnInit(): void {
		this.subscriptions.push(this.route.queryParams
			.subscribe((params) => {
				this.store.dispatch([
					new SetColliID(!!params.colliID ? params.colliID : ''),
					new SetOrderID(!!params.orderID ? params.orderID : ''),
					new FetchPackage(),
				]);
			}));

		this.subscriptions.push(this.searchProductsCtrl.valueChanges.pipe(debounceTime(300))
			.subscribe((v) => {
				this.store.dispatch(new SearchProducts(v))
			}));

		this.subscriptions.push(this.actions$.pipe(
			ofActionSuccessful(UpdateFormValue),
			debounceTime(200),
		).subscribe(() => {
			this.store.dispatch(new FetchDeliveryOptions());
		}));

	}

	ngOnDestroy() {
		this.subscriptions.map((s) => s.unsubscribe());
		this.store.dispatch(new ResetState());
	}

	addOrderLine() {
		this.dialog.open(AddOrderLineComponent, {width: "400px", height: "500px"});
	}

	isDeliveryOptionDisabled(status: DeliveryOptionBrandNameStatus): boolean {
		return status === DeliveryOptionBrandNameStatus.NotAvailable;
	}

	selectDeliveryOption(id: string, clickCollect: boolean) {
		this.store.dispatch(new SelectDeliveryOption({id, clickCollect}));
	}

	searchCountries(term: string) {
		this.store.dispatch(new SearchCountry(term));
	}

	changeCountrySender(country: CountriesResponse) {
		this.store.dispatch(new ChangeCountrySender(country));
	}

	changeCountryRecipient(country: CountriesResponse) {
		this.store.dispatch(new ChangeCountryRecipient(country));
	}

	save() {
		const id = this.store.selectSnapshot(OrderEditState.get).colliID;
		if (id.length === 0) {
			this.store.dispatch(new SaveFormNew());
		} else {
			this.store.dispatch(new SaveFormUpdate());
		}
	}

	priceEdited(edit: PriceEdited) {
		this.store.dispatch(new RowEditedPrice(edit));
	}

	unitsEdited(edit: UnitsEdited) {
		this.store.dispatch(new RowEditedUnits(edit));
	}
	lineDeleted(index: RowIndex) {
		this.store.dispatch(new RemoveOrderLine(index));
	}

	editDeliveryPoint() {
		this.dialog.open(EditDeliveryPointComponent);
	}

	editCCLocation() {
		this.dialog.open(EditCcLocationComponent);
	}

	editPackaging() {
		const ref = this.dialog.open(SelectPackagingComponent);
		ref.componentInstance.helpText = `Packaging to assign when shipment is created. May be overwritten if Workstation has packaging selected.`
		this.subscriptions.push(ref.componentInstance.selected.subscribe((p) => {
			console.warn(p);
			this.store.dispatch(new SetSelectedPackaging(p));
		}));
	}

	clearPackaging() {
		this.store.dispatch(new SetSelectedPackaging(null));
	}

}
