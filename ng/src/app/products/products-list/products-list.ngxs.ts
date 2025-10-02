import {Injectable} from "@angular/core";
import {Action, Selector, State, StateContext} from "@ngxs/store";
import {ProductsListActions} from "./products-list.actions";
import SetProductsList = ProductsListActions.SetProductsList;
import FetchProducts = ProductsListActions.FetchProducts;
import {CreateNewProductGQL, FetchProductsGQL, FetchProductsQueryVariables} from "./products-list.generated";
import {AppActions} from "../../app.actions";
import AppChangeRoute = AppActions.AppChangeRoute;
import {Paths} from "src/app/app-routing.module";
import {toNotNullArray} from "../../functions/not-null-array";
import FetchProductsList = ProductsListActions.FetchProductsList;

export interface ProductsListModel {
	list: FetchProducts[];
	pagination: ProductsPagination;
	loading: boolean;
}

export interface ProductsPagination {
	totalResults: number;
	pageIndex: number;
	hasNextPage: boolean;
	hasPreviousPage: boolean;
	startCursor: string | null;
	endCursor: string | null;
}

const defaultState: ProductsListModel = {
	list: [],
	loading: true,
	pagination: {
		totalResults: 0,
		pageIndex: 0,
		hasNextPage: true,
		hasPreviousPage: false,
		startCursor: null,
		endCursor: null,
	},
};

@Injectable()
@State<ProductsListModel>({
	name: 'productsList',
	defaults: defaultState,
})
export class ProductsListState {

	constructor(
		private list: FetchProductsGQL,
		private newProduct: CreateNewProductGQL,
	) {
	}

	@Selector()
	static get(state: ProductsListModel) {
		return state;
	}

	@Action(ProductsListActions.FetchProductsList)
	FetchMyProductsList(ctx: StateContext<ProductsListModel>, action: ProductsListActions.FetchProductsList) {
		ctx.patchState({loading: true});
		const state = ctx.getState();

		// Should correspond to paginator value
		const limit = 15;
		let variables: FetchProductsQueryVariables = {
			before: state.pagination.startCursor,
			last: limit,
		}
		if (action.payload === "next") {
			variables = {
				after: state.pagination.endCursor,
				first: limit,
			}
		}

		return this.list.fetch(variables, {fetchPolicy: 'no-cache'})
			.subscribe({next: (r) => {
				ctx.patchState({loading: false});
				const list = toNotNullArray(r.data.products.edges?.map((n) => n?.node));
				if (!!list) {
					ctx.dispatch(new SetProductsList(list));
				}

				const info = r.data.products;
				ctx.dispatch(new ProductsListActions.SetPagination({
					endCursor: info.pageInfo.endCursor as string,
					pageIndex: 0,
					startCursor: info.pageInfo.startCursor as string,
					totalResults: info.totalCount,
					hasNextPage: info.pageInfo.hasNextPage,
					hasPreviousPage: info.pageInfo.hasPreviousPage,
				}));
			}});
	}

	@Action(ProductsListActions.SetProductsList)
	SetMyProductsList(ctx: StateContext<ProductsListModel>, action: ProductsListActions.SetProductsList) {
		ctx.patchState({
			list: action.payload,
		});
	}

	@Action(ProductsListActions.AddNewProduct)
	AddNewProduct(ctx: StateContext<ProductsListModel>, action: ProductsListActions.AddNewProduct) {
		return this.newProduct.mutate({input: {title: action.payload}})
			.subscribe((r) => {
				ctx.dispatch(new AppChangeRoute({path: Paths.PRODUCTS_EDIT, queryParams: {id: r.data?.createProduct?.id}}));
			});
	}

	@Action(ProductsListActions.SetPagination)
	SetPagination(ctx: StateContext<ProductsListModel>, action: ProductsListActions.SetPagination) {
		ctx.patchState({pagination: action.payload});
	}

	@Action(ProductsListActions.PreviousPage)
	PreviousPage(ctx: StateContext<ProductsListModel>, action: ProductsListActions.PreviousPage) {
		ctx.dispatch(new FetchProductsList("previous"));
	}

	@Action(ProductsListActions.NextPage)
	NextPage(ctx: StateContext<ProductsListModel>, action: ProductsListActions.NextPage) {
		ctx.dispatch(new FetchProductsList());
	}

	@Action(ProductsListActions.Reset)
	Reset(ctx: StateContext<ProductsListModel>, action: ProductsListActions.Reset) {
		ctx.setState(defaultState);
	}

}
