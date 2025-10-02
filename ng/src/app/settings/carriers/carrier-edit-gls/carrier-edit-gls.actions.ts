import {FetchCarrierEditGlsQuery} from "./carrier-edit-gls.generated";

export namespace CarrierEditGLSActions {
	export class FetchCarrierEdit {
		static readonly type = '[CarrierEditGLS] fetch carrier edit GLS';
		constructor(public payload: string) {}
	}
	export class SetCarrierEdit {
		static readonly type = '[CarrierEditGLS] set carrier edit GLS';
		constructor(public payload: CarriersEditGLSQueryResponse) {}
	}
	export class SaveForm {
		static readonly type = '[CarrierEditGLS] save form';
	}
	export class Clear {
		static readonly type = '[CarrierEditGLS] clear';
	}
	export type CarriersEditGLSQueryResponse = NonNullable<FetchCarrierEditGlsQuery['carrier']>;
}
