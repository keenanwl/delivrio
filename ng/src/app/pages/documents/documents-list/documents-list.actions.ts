import {FetchDocumentsQuery} from "./documents-list.generated";
import {DocumentMergeType} from "../../../../generated/graphql";

export namespace DocumentsListActions {
	export class FetchDocumentsList {
		static readonly type = '[DocumentsList] fetch DocumentsList';
	}
	export class SetDocumentsList {
		static readonly type = '[DocumentsList] set DocumentsList';
		constructor(public payload: DocumentsResponse[]) {}
	}
	export class Clear {
		static readonly type = '[DocumentsList] clear';
	}
	export class Create {
		static readonly type = '[DocumentsList] create';
		constructor(public payload: {name: string, mergeType: DocumentMergeType}) {}
	}
	export type DocumentsResponse = NonNullable<NonNullable<NonNullable<NonNullable<FetchDocumentsQuery['documents']>['edges']>[0]>['node']>;
}
