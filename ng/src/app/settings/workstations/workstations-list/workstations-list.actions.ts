import {FetchWorkstationsQuery} from "./workstations-list.generated";
import {WorkstationDeviceType} from "../../../../generated/graphql";

export namespace WorkstationsListActions {
	export class FetchWorkstations {
		static readonly type = '[WorkstationsList] fetch workstations';
	}
	export class SetWorkstations {
		static readonly type = '[WorkstationsList] set workstations';
		constructor(public payload: WorkstationsResponse[]) {}
	}
	export class CreateNewWorkstation {
		static readonly type = '[WorkstationsList] create new workstation';
		constructor(public payload: {name: string, deviceType: WorkstationDeviceType}) {}
	}
	export class SetRegistrationToken {
		static readonly type = '[WorkstationsList] set registration token';
		constructor(public payload: {registrationToken: string, registrationTokenImg: string}) {}
	}
	export class ToggleArchived {
		static readonly type = '[WorkstationsList] toggle archived';
	}
	export type WorkstationsResponse = NonNullable<NonNullable<FetchWorkstationsQuery['filteredWorkstations']>[0]>;
}
