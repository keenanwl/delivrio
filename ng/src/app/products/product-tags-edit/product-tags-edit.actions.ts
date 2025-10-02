import {FetchProductTagsEditQuery} from "./product-tags-edit.generated";

export namespace ProductTagsEditTagsEditActions {
	export class FetchProductTagsEdit {
		static readonly type = '[ProductTagsEdit] fetch ProductTagsEdit';
	}
	export class SetProductTagsEdit {
		static readonly type = '[ProductTagsEdit] set ProductTagsEdit';
		constructor(public payload: FetchProductTagsEditTagsResponse[]) {}
	}
	export class SaveTagList {
		static readonly type = '[ProductTagsEdit] save tag list';
		constructor(public payload: string[]) {}
	}
	export class SaveTagListSuccess {
		static readonly type = '[ProductTagsEdit] save tag list success';
	}
	export class DeleteTag {
		static readonly type = '[ProductTagsEdit] save tag list success';
		constructor(public payload: string) {}
	}
	export type FetchProductTagsEditTagsResponse = NonNullable<NonNullable<NonNullable<FetchProductTagsEditQuery['productTags']>['edges']>[0]>['node'];
}
