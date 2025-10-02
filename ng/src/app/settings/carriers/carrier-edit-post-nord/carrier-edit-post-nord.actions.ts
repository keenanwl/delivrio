import {FetchCarrierPostNordQuery} from "./carrier-edit-post-nord.generated";

export namespace CarrierEditPostNordActions {
	export class SetID {
		static readonly type = '[CarrierEditPostNord] set ID';
		constructor(public payload: string) {}
	}
	export class FetchCarrierPostNordEdit {
		static readonly type = '[CarrierEditPostNord] fetch CarrierEditPostNord';
	}
	export class SetCarrierPostNordEdit {
		static readonly type = '[CarrierEditPostNord] set Carrier edit';
		constructor(public payload: SelectCarriersEditQueryResponse) {}
	}
	export class SaveForm {
		static readonly type = '[CarrierEditPostNord] save form';
	}
	export class Clear {
		static readonly type = '[CarrierEditPostNord] clear';
	}
	export type SelectCarriersEditQueryResponse = NonNullable<FetchCarrierPostNordQuery['carrierPostNord']> & {name?: string};
}
