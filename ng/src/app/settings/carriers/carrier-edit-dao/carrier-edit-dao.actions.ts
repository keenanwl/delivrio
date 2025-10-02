import {FetchCarrierEditDaoQuery} from "./carrier-edit-dao.generated";

export namespace CarrierEditDAOActions {
	export class FetchCarrierEditDAO {
		static readonly type = '[CarrierEditDAO] fetch carrier edit Dao';
		constructor(public payload: string) {}
	}
	export class SetCarrier {
		static readonly type = '[CarrierEditDAO] set carrier edit';
		constructor(public payload: EditResponse) {}
	}
	export class SaveForm {
		static readonly type = '[CarrierEditDAO] save form';
	}
	export class Clear {
		static readonly type = '[CarrierEditDAO] clear';
	}
	export type EditResponse = NonNullable<FetchCarrierEditDaoQuery['carrier']>;
}
