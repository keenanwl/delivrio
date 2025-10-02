import {FetchDocumentQuery} from "./document-edit.generated";

export namespace DocumentEditActions {
	export class FetchDocumentEdit {
		static readonly type = '[DocumentEdit] fetch DocumentEdit';
	}
	export class SetDocumentEdit {
		static readonly type = '[DocumentEdit] set DocumentEdit';
		constructor(public payload: DocumentResponse) {}
	}
	export class SetDocumentID {
		static readonly type = '[DocumentEdit] set DocumentEdit ID';
		constructor(public payload: string) {}
	}
	export class SetCarrierBrands {
		static readonly type = '[DocumentEdit] set carrier brands';
		constructor(public payload: CarrierBrandResponse[]) {}
	}
	export class SetDateTimeRange {
		static readonly type = '[DocumentEdit] set date time range';
		constructor(public payload: {start: string; end: string}) {}
	}
	export class Download {
		static readonly type = '[DocumentEdit] download';
	}
	export class SetPDF {
		static readonly type = '[DocumentEdit] set PDF';
		constructor(public payload: string) {}
	}
	export class Clear {
		static readonly type = '[DocumentEdit] clear';
	}
	export class Save {
		static readonly type = '[DocumentEdit] create';
	}
	export type DocumentResponse = NonNullable<FetchDocumentQuery['document']>;
	export type CarrierBrandResponse = NonNullable<NonNullable<NonNullable<NonNullable<FetchDocumentQuery['carrierBrands']>['edges']>[0]>['node']>;
}
