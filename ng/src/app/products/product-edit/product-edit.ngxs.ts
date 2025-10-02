import {Injectable} from "@angular/core";
import {Action, Selector, State, StateContext} from "@ngxs/store";
import {GraphQLError} from "graphql";
import {AppActions} from "../../app.actions";
import {ProductActions} from "./product-edit.actions";
import {Paths} from "../../app-routing.module";
import {ProductsListActions} from "../products-list/products-list.actions";
import AppChangeRoute = AppActions.AppChangeRoute;
import FetchProductResponse = ProductActions.FetchProductResponse;
import SetProduct = ProductActions.SetProduct;
import FetchProductTagsResponse = ProductActions.FetchProductTagsResponse;
import SetProductTags = ProductActions.SetProductTags;
import {produce} from "immer";
import {
	ArchiveVariantGQL,
	CreateProductGQL,
	CreateVariantGQL,
	FetchProductGQL, MustInventoryItemGQL, ProductDeleteImageGQL, ProductSearchCountriesGQL,
	ProductUploadImageGQL, UpdateInventoryItemGQL,
	UpdateProductGQL
} from "./product-edit.generated";
import {ProductStatus} from "../../../generated/graphql";
import SetProductImageInfo = ProductActions.SetProductImageInfo;
import ProductImageResponse = ProductActions.ProductImageResponse;
import VariantResponse = ProductActions.VariantResponse;
import SetVariants = ProductActions.SetVariants;
import {toNotNullArray} from "../../functions/not-null-array";
import FetchProduct = ProductActions.FetchProduct;
import MustInventoryItemResponse = ProductActions.MustInventoryItemResponse;
import CountriesResponse = ProductActions.CountriesResponse;
import SetCountries = ProductActions.SetCountries;
import SetInventoryItem = ProductActions.SetInventoryItem;
import ShowGlobalSnackbar = AppActions.ShowGlobalSnackbar;
import CloseInventoryForm = ProductActions.CloseInventoryForm;

export interface ProductModel {
	productForm: {
		model: FetchProductResponse;
		dirty: boolean;
		status: string;
		errors: readonly GraphQLError[];
	},
	variantsForm: {
		model: {variants: VariantResponse[]};
		dirty: boolean;
		status: string;
		errors: readonly GraphQLError[];
	},
	inventoryItemForm: {
		model: MustInventoryItemResponse | undefined;
		dirty: boolean;
		status: string;
		errors: readonly GraphQLError[];
	},
	productID: string;
	productTags: FetchProductTagsResponse[];
	imageIndex: number;
	productImages: ProductImageResponse;
	searchCountries: CountriesResponse[];
	loading: boolean;
}

const defaultState: ProductModel = {
	productForm: {
		model: {
			title: "",
			productTags: [],
			status: ProductStatus.Active,
			createdAt: "",
			bodyHTML: "",
		},
		dirty: false,
		status: '',
		errors: [],
	},
	variantsForm: {
		model: {variants: []},
		dirty: false,
		status: '',
		errors: [],
	},
	inventoryItemForm: {
		model: undefined,
		dirty: false,
		status: '',
		errors: [],
	},
	productID: '',
	productTags: [],
	imageIndex: 0,
	productImages: [],
	searchCountries: [],
	loading: false,
};

@Injectable()
@State<ProductModel>({
	name: 'productEdit',
	defaults: defaultState,
})
export class ProductState {

	constructor(
		private fetchProduct: FetchProductGQL,
		private createProduct: CreateProductGQL,
		private updateProduct: UpdateProductGQL,
		private createVariant: CreateVariantGQL,
		private uploadImage: ProductUploadImageGQL,
		private deleteImage: ProductDeleteImageGQL,
		private archiveVariant: ArchiveVariantGQL,
		private mustInventoryItem: MustInventoryItemGQL,
		private countrySearch: ProductSearchCountriesGQL,
		private updateInventory: UpdateInventoryItemGQL,
	) {}

	@Selector()
	static get(state: ProductModel) {
		return state;
	}

	@Action(ProductActions.FetchProduct)
	FetchMyProduct(ctx: StateContext<ProductModel>, action: ProductActions.FetchProduct) {
		ctx.patchState({loading: true});
		const state = ctx.getState();
		return this.fetchProduct.fetch({id: state.productID}, {fetchPolicy: "no-cache", errorPolicy: "all"})
			.subscribe({next: (r) => {
				ctx.patchState({loading: false});
				const images = r.data.productForImage?.productImage;
				if (!!images) {
					ctx.dispatch(new SetProductImageInfo(images));
				}
				const product = r.data.product;
				if (!!product) {
					ctx.dispatch(new SetProduct(product));
				}
				const variants = toNotNullArray(r.data.productVariants.edges?.map((n) => n?.node));
				if (!!variants) {
					ctx.dispatch(new SetVariants(variants));
				}
				const groups = r.data.productTags.edges?.map(g => g?.node);
				if (!!groups) {
					ctx.dispatch(new SetProductTags(groups));
				}
			}});
	}

	@Action(ProductActions.SetProductID)
	SetProductID(ctx: StateContext<ProductModel>, action: ProductActions.SetProductID) {
		ctx.patchState({
			productID: action.payload,
		})
	}

	@Action(ProductActions.SetProductTags)
	SetProductTags(ctx: StateContext<ProductModel>, action: ProductActions.SetProductTags) {
		ctx.patchState({
			productTags: action.payload,
		})
	}

	@Action(ProductActions.SetProduct)
	SetMyProduct(ctx: StateContext<ProductModel>, action: ProductActions.SetProduct) {
		const state = ctx.getState();
		const next = Object.assign({}, state.productForm, {
			model: Object.assign({}, action.payload)
		});
		ctx.patchState({
			productForm: next,
		})
	}

	@Action(ProductActions.AddTag)
	AddTag(ctx: StateContext<ProductModel>, action: ProductActions.AddTag) {
		const val = action.payload;
		if (val) {
			const state = produce(ctx.getState(), st => {
				st.productForm.model?.productTags?.push(val)
			});
			ctx.setState(state);
		}
	}

	@Action(ProductActions.RemoveTag)
	RemoveTag(ctx: StateContext<ProductModel>, action: ProductActions.RemoveTag) {
		const state = produce(ctx.getState(), st => {
			st.productForm.model.productTags = st.productForm.model?.productTags?.filter((t) => t.id !== action.payload);
		});
		ctx.setState(state);
	}

	@Action(ProductActions.SetCountries)
	SetCountries(ctx: StateContext<ProductModel>, action: ProductActions.SetCountries) {
		ctx.patchState({searchCountries: action.payload});
	}

	@Action(ProductActions.SaveFormNew)
	SaveFormNew(ctx: StateContext<ProductModel>, action: ProductActions.SaveFormNew) {
		return this.createProduct.mutate(action.payload)
			.subscribe(() => {
				ctx.dispatch([
					new ProductsListActions.FetchProductsList(),
					new AppChangeRoute({path: Paths.PRODUCTS, queryParams: {}}),
				]);
			});
	}

	@Action(ProductActions.SaveFormUpdate)
	SaveFormUpdate(ctx: StateContext<ProductModel>, action: ProductActions.SaveFormUpdate) {
		const state = ctx.getState();
		return this.updateProduct.mutate({
			id: state.productID,
			input: action.payload.input,
			variants: action.payload.variants,
			images: state.productImages.map((i) => {
				return {
					variantIDs: i.productVariant?.map((v) => v.id) || [],
					imageID: i.id,
				}
			})
		}).subscribe(() => {
			ctx.dispatch([
				new ProductsListActions.FetchProductsList(),
				new AppChangeRoute({path: Paths.PRODUCTS, queryParams: {}}),
			]);
		});
	}

	@Action(ProductActions.CreateVariant)
	CreateVariant(ctx: StateContext<ProductModel>, action: ProductActions.CreateVariant) {
		const state = ctx.getState();
		return this.createVariant.mutate({productID: state.productID})
			.subscribe((r) => {
				/*const state2 = produce(ctx.getState(), st => {
					const nextVariant = r.data?.createVariant
					if (!!nextVariant) {
						if (!!st.productForm.model.productVariant) {
							st.productForm.model.productVariant.push(nextVariant);
						} else {
							st.productForm.model.productVariant = [nextVariant];
						}
					}
				});*/
				//ctx.dispatch(new SetProduct(state2.productForm.model));
			});
	}

	@Action(ProductActions.UploadImage)
	UploadImage(ctx: StateContext<ProductModel>, action: ProductActions.UploadImage) {
		return this.uploadImage.mutate({productID: ctx.getState().productID, image: action.payload})
			.subscribe((r) => {
				const img = r.data?.uploadProductImage.productImage;
				if (!!img) {
					ctx.dispatch(new ProductActions.SetProductImageInfo(img));
				}
			});
	}

	@Action(ProductActions.SetProductImageInfo)
	SetProductImageInfo(ctx: StateContext<ProductModel>, action: ProductActions.SetProductImageInfo) {
		ctx.patchState({productImages: action.payload});
	}

	@Action(ProductActions.NextImage)
	NextImage(ctx: StateContext<ProductModel>, action: ProductActions.NextImage) {
		const state = ctx.getState();
		if (state.imageIndex + 1 < state.productImages.length) {
			ctx.patchState({imageIndex: state.imageIndex + 1});
		}
	}

	@Action(ProductActions.PreviousImage)
	PreviousImage(ctx: StateContext<ProductModel>, action: ProductActions.PreviousImage) {
		const state = ctx.getState();
		if (state.imageIndex > 0) {
			ctx.patchState({imageIndex: state.imageIndex - 1});
		}
	}

	@Action(ProductActions.AddImageVariant)
	AddImageVariant(ctx: StateContext<ProductModel>, action: ProductActions.AddImageVariant) {
		const state = produce(ctx.getState(), st => {
			st.productImages[st.imageIndex].productVariant = st.productImages[st.imageIndex].productVariant!
				.filter((i) => i.id !== action.payload.variantID);
			st.productImages[st.imageIndex].productVariant!.push({id: action.payload.variantID});
		});
		ctx.setState(state);
	}

	@Action(ProductActions.RemoveImageVariant)
	RemoveImageVariant(ctx: StateContext<ProductModel>, action: ProductActions.RemoveImageVariant) {
		const state = produce(ctx.getState(), st => {
			st.productImages[st.imageIndex].productVariant = st.productImages[st.imageIndex].productVariant!
				.filter((i) => i.id !== action.payload.variantID);
		});
		ctx.setState(state);
	}

	@Action(ProductActions.Reset)
	Reset(ctx: StateContext<ProductModel>, action: ProductActions.Reset) {
		ctx.setState(defaultState);
	}

	@Action(ProductActions.DeleteImage)
	DeleteImage(ctx: StateContext<ProductModel>, action: ProductActions.DeleteImage) {
		return this.deleteImage.mutate({imageID: action.payload.imageID})
			.subscribe((i) => {
				const imgs = i.data?.deleteProductImage.productImage;
				if (!!imgs) {
					ctx.patchState({imageIndex: 0});
					ctx.dispatch(new SetProductImageInfo(imgs));
				}
			});
	}

	@Action(ProductActions.SetVariants)
	SetVariants(ctx: StateContext<ProductModel>, action: ProductActions.SetVariants) {
		const state = produce(ctx.getState(), st => {
			st.variantsForm.model.variants = action.payload;
		});
		ctx.setState(state);
	}

	@Action(ProductActions.ArchiveVariant)
	ArchiveVariant(ctx: StateContext<ProductModel>, action: ProductActions.ArchiveVariant) {
		return this.archiveVariant.mutate({variantID: action.payload})
			.subscribe((r) => {
				ctx.dispatch(new FetchProduct());
			});
	}

	@Action(ProductActions.MustInventoryItem)
	MustInventoryItem(ctx: StateContext<ProductModel>, action: ProductActions.MustInventoryItem) {
		return this.mustInventoryItem.mutate({productVariantID: action.payload})
			.subscribe((r) => {
				ctx.dispatch(new SetInventoryItem(r.data?.mustInventory));
			});
	}

	@Action(ProductActions.SearchCountries)
	SearchCountries(ctx: StateContext<ProductModel>, action: ProductActions.SearchCountries) {
		return this.countrySearch.fetch({term: action.payload}, {errorPolicy: "all"})
			.subscribe((r) => {
				ctx.dispatch(new SetCountries(toNotNullArray(r.data.countries.edges?.map((n) => n?.node))));
			});
	}

	@Action(ProductActions.ChangeCountry)
	ChangeCountry(ctx: StateContext<ProductModel>, action: ProductActions.ChangeCountry) {
		const state = produce(ctx.getState(), st => {
			st.inventoryItemForm.model!.countryOfOrigin = action.payload;
		});
		ctx.setState(state);
	}

	@Action(ProductActions.SetInventoryItem)
	SetInventoryItem(ctx: StateContext<ProductModel>, action: ProductActions.SetInventoryItem) {
		const state = produce(ctx.getState(), st => {
			st.inventoryItemForm.model = action.payload;
		});
		ctx.setState(state);
	}

	@Action(ProductActions.ResetInventoryForm)
	ResetInventoryForm(ctx: StateContext<ProductModel>, action: ProductActions.ResetInventoryForm) {
		const state = produce(ctx.getState(), st => {
			st.inventoryItemForm.model = undefined;
		});
		ctx.setState(state);
	}

	@Action(ProductActions.SaveInventoryItem)
	SaveInventoryItem(ctx: StateContext<ProductModel>, action: ProductActions.SaveInventoryItem) {
		return this.updateInventory.mutate(action.payload, {errorPolicy: "all"})
			.subscribe((r) => {
				if (!!r.errors) {
					ctx.dispatch(new ShowGlobalSnackbar("Error saving: " + JSON.stringify(r.errors)));
				} else {
					ctx.dispatch(new CloseInventoryForm())
					ctx.dispatch(new ShowGlobalSnackbar("Saved successfully"));
				}
			});
	}

}
