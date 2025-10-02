import {FetchCarrierUspsQuery} from "./carrier-edit-usps.generated";

export namespace CarrierEditUSPSActions {
	export class SetID {
		static readonly type = '[CarrierEditUSPS] set ID';
		constructor(public payload: string) {}
	}
	export class FetchCarrierUSPSEdit {
		static readonly type = '[CarrierEditUSPS] fetch CarrierEditUSPS';
	}
	export class SetCarrierUSPSEdit {
		static readonly type = '[CarrierEditUSPS] set Carrier edit';
		constructor(public payload: SelectCarriersEditQueryResponse) {}
	}
	export class SaveForm {
		static readonly type = '[CarrierEditUSPS] save form';
	}
	export class Clear {
		static readonly type = '[CarrierEditUSPS] clear';
	}
	export type SelectCarriersEditQueryResponse = NonNullable<FetchCarrierUspsQuery['carrierUSPS']> & {name?: string};
}
