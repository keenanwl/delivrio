import {Component, OnDestroy, OnInit} from '@angular/core';
import {Observable} from "rxjs";
import {Store} from "@ngxs/store";
import {MatDialog} from "@angular/material/dialog";
import {ProductsListModel, ProductsListState} from "./products-list.ngxs";
import {ProductsListActions} from "./products-list.actions";
import FetchProductsList = ProductsListActions.FetchProductsList;
import {Paths} from "../../app-routing.module";
import {AppActions} from "../../app.actions";
import AppChangeRoute = AppActions.AppChangeRoute;
import {AddProductDialogComponent} from "./dialogs/add-product-dialog/add-product-dialog.component";
import {ProductVariant} from "../../../generated/graphql";
import {PageEvent} from "@angular/material/paginator";
import NextPage = ProductsListActions.NextPage;
import PreviousPage = ProductsListActions.PreviousPage;
import Reset = ProductsListActions.Reset;

@Component({
	selector: 'app-products-list',
	templateUrl: './products-list.component.html',
	styleUrls: ['./products-list.component.scss']
})
export class ProductsListComponent implements OnInit, OnDestroy {

	productsList$: Observable<ProductsListModel>;

	displayedColumns: string[] = [
		'image',
		'name',
		'tags',
		'createdAt',
		'updatedAt',
		//'dimensions',
		'status',
	];

	constructor(
		private store: Store,
		private dialog: MatDialog,
	) {
		this.productsList$ = store.select(ProductsListState.get);
	}

	ngOnInit(): void {
		this.store.dispatch(new FetchProductsList());
	}

	ngOnDestroy() {
		this.store.dispatch(new Reset());
	}

	editOption(id: string) {
		this.store.dispatch([
			new AppChangeRoute({path: Paths.PRODUCTS_EDIT, queryParams: {id}}),
		]);
	}

	addNewProduct() {
		this.dialog.open(AddProductDialogComponent);
	}

	firstProductImage(variant: ProductVariant): string {

		if (!variant.productImage || !variant.product.productImage) {
			return "";
		}

		if (variant.productImage?.length > 0) {
			return variant.productImage[0]?.url;
		}

		if (variant.product.productImage?.length > 0) {
			return variant.product.productImage[0]?.url;
		}

		return "";
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
		this.store.dispatch(new AppChangeRoute({path: Paths.PRODUCTS_EDIT, queryParams: {id}}));
	}

}
