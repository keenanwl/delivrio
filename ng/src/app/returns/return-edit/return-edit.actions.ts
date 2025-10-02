
export namespace ReturnEditActions {
	export class FetchReturn {
		static readonly type = '[ReturnEdit] fetch Return';
	}
	export class SetReturnID {
		static readonly type = '[ReturnEdit] set product ID';
		constructor(public payload: string) {}
	}
	export class SetReturn {
		static readonly type = '[ReturnEdit] set product';
		//constructor(public payload: FetchReturnResponse) {}
	}
	export class SetReturnTags {
		static readonly type = '[ReturnEdit] set product tags';
		//constructor(public payload: FetchReturnTagsResponse[]) {}
	}
	export class AddTag {
		static readonly type = '[ReturnEdit] add tag';
		//constructor(public payload: NonNullable<FetchReturnResponse['productTags']>[0]) {}
	}
	export class RemoveTag {
		static readonly type = '[ReturnEdit] remove tag';
		constructor(public payload: string) {}
	}
	export class CreateVariant {
		static readonly type = '[ReturnEdit] create variant';
	}
	/*export type FetchReturnResponse = NonNullable<FetchReturnQuery['product']>;
	export type FetchReturnTagsResponse = NonNullable<NonNullable<NonNullable<FetchReturnQuery['productTags']>['edges']>[0]>['node'];*/
}
