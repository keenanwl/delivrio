import {FetchPackagingQuery} from "../../orders/order-edit/order-edit.generated";

export namespace SelectPackagingActions {

	export class FetchPackaging {
		static readonly type = '[SelectPackaging] fetch packaging';
	}
	export class SetPackaging {
		static readonly type = '[SelectPackaging] set packaging';
		constructor(public payload: PackagingResponse[]) {}
	}

	export type PackagingResponse = NonNullable<NonNullable<NonNullable<NonNullable<FetchPackagingQuery['packagings']>['edges']>[0]>['node']>;
}
