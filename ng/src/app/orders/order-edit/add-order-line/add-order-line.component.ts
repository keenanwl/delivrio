import {Component, OnInit} from '@angular/core';
import {Observable, Subscription} from "rxjs";
import {OrderEditModel, OrderEditState} from "../order-edit.ngxs";
import {FormControl} from "@angular/forms";
import {ActivatedRoute} from "@angular/router";
import {Store} from "@ngxs/store";
import {debounceTime} from "rxjs/operators";
import {OrderEditActions} from "../order-edit.actions";
import SearchProducts = OrderEditActions.SearchProducts;
import AddProduct = OrderEditActions.AddProduct;
import SearchProductsResponse = OrderEditActions.SearchProductsResponse;
import {Paths} from "../../../app-routing.module";
import {MatDialogRef} from "@angular/material/dialog";
import {defaultProductImg, ProductVariantImages} from "../../../functions/product-image";

@Component({
	selector: 'app-add-order-line',
	templateUrl: './add-order-line.component.html',
	styleUrls: ['./add-order-line.component.scss']
})
export class AddOrderLineComponent implements OnInit {

	order$: Observable<OrderEditModel>;
	searchProductsCtrl = new FormControl('', {nonNullable: true});

	productsURL = Paths.PRODUCTS_EDIT;

	subscriptions: Subscription[] = [];

	constructor(
		private route: ActivatedRoute,
		private store: Store,
		private ref: MatDialogRef<AddOrderLineComponent>,
	) {
		this.order$ = store.select(OrderEditState.get);
	}

	ngOnInit() {
		this.subscriptions.push(this.searchProductsCtrl.valueChanges.pipe(debounceTime(300))
			.subscribe((v) => {
				this.store.dispatch(new SearchProducts(v))
			}));
	}

	resetSearch() {
		this.searchProductsCtrl.reset();
	}

	addProduct(variant: SearchProductsResponse) {
		this.store.dispatch(new AddProduct(variant));
	}

	close() {
		this.ref.close();
	}

	image(variantData: ProductVariantImages): string {
		return defaultProductImg(variantData);
	}

}
