import {FetchCarrierEditDfQuery} from "./carrier-edit-df.generated";

export namespace CarrierEditDFActions {
	export class FetchCarrierEditDF {
		static readonly type = '[CarrierEditDF] fetch carrier edit Dao';
		constructor(public payload: string) {}
	}
	export class SetCarrier {
		static readonly type = '[CarrierEditDF] set carrier edit';
		constructor(public payload: EditResponse) {}
	}
	export class SaveForm {
		static readonly type = '[CarrierEditDF] save form';
	}
	export class Clear {
		static readonly type = '[CarrierEditDF] clear';
	}
	export type EditResponse = NonNullable<FetchCarrierEditDfQuery['carrier']>;
}
