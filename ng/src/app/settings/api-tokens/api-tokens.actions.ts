import {FetchMyApiTokensQuery} from "./api-tokens.generated";

export namespace APITokensActions {
	export class FetchAPITokens {
		static readonly type = '[APITokens] fetch APITokens';
	}
	export class SetAPITokens {
		static readonly type = '[APITokens] set APITokens';
		constructor(public payload: FetchAPITokensResponse[]) {}
	}
	export class CreateAPIToken {
		static readonly type = '[APITokens] create new API token';
		constructor(public payload: {name: string}) {}
	}
	export class CreateAPITokenSuccess {
		static readonly type = '[APITokens] create new API token success';
		constructor(public payload: {token: string}) {}
	}
	export class SetNewToken {
		static readonly type = '[APITokens] set new token';
		constructor(public payload: {token: string}) {}
	}
	export class ClearDialogs {
		static readonly type = '[APITokens] clear new token';
	}
	export class DeleteToken {
		static readonly type = '[APITokens] delete for real';
		constructor(public payload: string) {}
	}
	export class SetConfirmDeleteToken {
		static readonly type = '[APITokens] set confirm delete';
		constructor(public payload: string) {}
	}
	export type FetchAPITokensResponse = NonNullable<NonNullable<FetchMyApiTokensQuery['myAPITokens']>[0]>;
}
