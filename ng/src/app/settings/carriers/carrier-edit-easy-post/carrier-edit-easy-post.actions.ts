import {FetchCarrierEasyPostQuery} from "./carrier-edit-easy-post.generated";

export namespace CarrierEditEasyPostActions {
	export class SetID {
		static readonly type = '[CarrierEditEasyPost] set ID';
		constructor(public payload: string) {}
	}
	export class FetchCarrierEasyPostEdit {
		static readonly type = '[CarrierEditEasyPost] fetch CarrierEditEasyPost';
	}
	export class SetCarrierEasyPostEdit {
		static readonly type = '[CarrierEditEasyPost] set Carrier edit';
		constructor(public payload: SelectCarriersEditQueryResponse) {}
	}
	export class SaveForm {
		static readonly type = '[CarrierEditEasyPost] save form';
	}
	export class Clear {
		static readonly type = '[CarrierEditEasyPost] clear';
	}
	export type SelectCarriersEditQueryResponse = NonNullable<FetchCarrierEasyPostQuery['carrier']>;
}
