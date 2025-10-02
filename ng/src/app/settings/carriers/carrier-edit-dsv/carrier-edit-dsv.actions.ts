import {FetchCarrierEditDsvQuery} from "./carrier-edit-dsv.generated";

export namespace CarrierEditDSVActions {
	export class FetchCarrierEditDSV {
		static readonly type = '[CarrierEditDSV] fetch carrier edit Dao';
		constructor(public payload: string) {}
	}
	export class SetCarrier {
		static readonly type = '[CarrierEditDSV] set carrier edit';
		constructor(public payload: EditResponse) {}
	}
	export class SaveForm {
		static readonly type = '[CarrierEditDSV] save form';
	}
	export class Clear {
		static readonly type = '[CarrierEditDSV] clear';
	}
	export type EditResponse = NonNullable<FetchCarrierEditDsvQuery['carrier']>;
}
