import {FetchCarrierEditBringQuery} from "./carrier-edit-bring.generated";
import {environment} from "../../../../environments/environment";

const actionName = (val: string): string => {
	return environment.production ? val : "llll";
};

export namespace CarrierEditBringActions {
	export class FetchCarrierEditBring {
		static readonly type = actionName('[CarrierEditBring] fetch carrier edit bring');
		constructor(public payload: string) {}
	}
	export class SetCarrier {
		static readonly type = '[CarrierEditBring] set carrier edit';
		constructor(public payload: EditResponse) {}
	}
	export class SaveForm {
		static readonly type = '[CarrierEditBring] save form';
	}
	export class Clear {
		static readonly type = '[CarrierEditBring] clear';
	}
	export type EditResponse = NonNullable<FetchCarrierEditBringQuery['carrier']>;
}
