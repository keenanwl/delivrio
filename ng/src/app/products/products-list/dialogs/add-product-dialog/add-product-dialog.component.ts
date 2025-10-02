import { Component, OnInit } from '@angular/core';
import {Observable} from "rxjs";
import {ProductsListModel, ProductsListState} from "../../products-list.ngxs";
import {Store} from "@ngxs/store";
import {MatDialog} from "@angular/material/dialog";
import {ProductsListActions} from "../../products-list.actions";
import AddNewProduct = ProductsListActions.AddNewProduct;

@Component({
	selector: 'app-add-product-dialog',
	templateUrl: './add-product-dialog.component.html',
	styleUrls: ['./add-product-dialog.component.scss']
})
export class AddProductDialogComponent implements OnInit {

	productsList$: Observable<ProductsListModel>;

	constructor(
		private store: Store,
		private dialog: MatDialog,
	) {
		this.productsList$ = store.select(ProductsListState.get);
	}

	ngOnInit(): void {
	}

	addProduct(name: string) {
		this.store.dispatch(new AddNewProduct(name));
		this.dialog.closeAll();
	}

}
