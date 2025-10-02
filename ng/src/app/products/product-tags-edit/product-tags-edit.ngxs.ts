import {Injectable} from "@angular/core";
import {Action, Selector, State, StateContext} from "@ngxs/store";
import {ProductTagsEditTagsEditActions} from "./product-tags-edit.actions";
import FetchProductTagsEditTagsResponse = ProductTagsEditTagsEditActions.FetchProductTagsEditTagsResponse;
import {ProductModel} from "../product-edit/product-edit.ngxs";
import SetProductTagsEdit = ProductTagsEditTagsEditActions.SetProductTagsEdit;
import SaveTagListSuccess = ProductTagsEditTagsEditActions.SaveTagListSuccess;
import {CreateTagsGQL, DeleteTagGQL, FetchProductTagsEditGQL} from "./product-tags-edit.generated";

export interface ProductTagsEditModel {
	productTags: FetchProductTagsEditTagsResponse[];
}

const defaultState: ProductTagsEditModel = {
	productTags: [],
};

@Injectable()
@State<ProductTagsEditModel>({
	name: 'productTagsEdit',
	defaults: defaultState,
})
export class ProductTagsEditState {

	constructor(
		private fetchTags: FetchProductTagsEditGQL,
		private createTags: CreateTagsGQL,
		private deleteTag: DeleteTagGQL,
	) {}

	@Selector()
	static get(state: ProductTagsEditModel) {
		return state;
	}

	@Action(ProductTagsEditTagsEditActions.FetchProductTagsEdit)
	FetchProductTagsEdit(ctx: StateContext<ProductModel>, action: ProductTagsEditTagsEditActions.FetchProductTagsEdit) {
		const state = ctx.getState();
		return this.fetchTags.fetch({}, {fetchPolicy: "no-cache", errorPolicy: "all"})
			.subscribe({next: (r) => {
					const tags = r.data.productTags.edges?.map(g => g?.node);
					if (!!tags) {
						ctx.dispatch([new SetProductTagsEdit(tags)]);
					}
				}, error: (e) => {

				}});
	}

	@Action(ProductTagsEditTagsEditActions.SaveTagList)
	SaveTagList(ctx: StateContext<ProductModel>, action: ProductTagsEditTagsEditActions.SaveTagList) {
		const state = ctx.getState();
		return this.createTags.mutate({input: action.payload})
			.subscribe((r) => {
				const tags = r.data?.createProductTags;
				if (!!tags) {
					ctx.dispatch([new SetProductTagsEdit(tags), new SaveTagListSuccess()]);
				}
			});
	}

	@Action(ProductTagsEditTagsEditActions.SetProductTagsEdit)
	SetProductTagsEdit(ctx: StateContext<ProductModel>, action: ProductTagsEditTagsEditActions.SetProductTagsEdit) {
		ctx.patchState({
			productTags: action.payload,
		});
	}

	@Action(ProductTagsEditTagsEditActions.DeleteTag)
	DeleteTag(ctx: StateContext<ProductModel>, action: ProductTagsEditTagsEditActions.DeleteTag) {
		return this.deleteTag.mutate({id: action.payload})
			.subscribe((r) => {
				const tags = r.data?.deleteTag;
				if (!!tags) {
					ctx.dispatch([new SetProductTagsEdit(tags)]);
				}
			});
	}

}
