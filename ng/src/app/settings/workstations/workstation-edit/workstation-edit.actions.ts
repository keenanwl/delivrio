import {FetchWorkstationQuery} from "./workstation-edit.generated";

export namespace WorkstationEditActions {
	export class FetchWorkstationEdit {
		static readonly type = '[WorkstationEdit] fetch WorkstationEdit';
	}
	export class SetWorkstationEdit {
		static readonly type = '[WorkstationEdit] set Workstation edit';
		constructor(public payload: WorkstationEditResponse) {}
	}
	export class SetWorkstationID {
		static readonly type = '[WorkstationEdit] set Workstation ID';
		constructor(public payload: string) {}
	}
	export class Save {
		static readonly type = '[WorkstationEdit] save';
	}
	export class Disable {
		static readonly type = '[WorkstationEdit] disable';
	}
	export class Reset {
		static readonly type = '[WorkstationEdit] reset';
	}
	export type WorkstationEditResponse = NonNullable<FetchWorkstationQuery['workstation']>;
	export type PrinterResponse = NonNullable<NonNullable<NonNullable<FetchWorkstationQuery['workstation']>['printer']>[0]>;
}
