import {FetchProductsQuery} from "./products-list.generated";
import {ProductsPagination} from "./products-list.ngxs";

export namespace ProductsListActions {
	export class FetchProductsList {
		static readonly type = '[ProductsList] fetch ProductsList';
		constructor(public payload: "previous" | "next" = "next") {}
	}
	export class SetProductsList {
		static readonly type = '[ProductsList] set ProductsList';
		constructor(public payload: FetchProducts[]) {}
	}
	export class AddNewProduct {
		static readonly type = '[ProductsList] add new product';
		constructor(public payload: string) {}
	}
	export class SetPagination {
		static readonly type = '[ProductsList] set pagination';
		constructor(public payload: ProductsPagination) {}
	}
	export class NextPage {
		static readonly type = '[ProductsList] next page';
	}
	export class PreviousPage {
		static readonly type = '[ProductsList] previous page';
	}
	export class Reset {
		static readonly type = '[ProductsList] reset';
	}
	export type FetchProducts = NonNullable<NonNullable<NonNullable<FetchProductsQuery['products']['edges']>[0]>['node']>;
	export type ProductVariant = NonNullable<FetchProducts['productVariant']>[0];
}
