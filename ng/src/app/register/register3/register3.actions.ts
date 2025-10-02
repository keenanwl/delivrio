import {FetchPlatformsCarriersQuery} from "./register3.generated";

export namespace Register3Actions {

	export class SetUserID {
		static readonly type = '[Register3] set user ID';
		constructor(public payload: string) {}
	}
	export class SetInvalidMessage {
		static readonly type = '[Register3] set invalid message';
		constructor(public payload: string) {}
	}
	export class SaveRegistration {
		static readonly type = '[Register3] save registration part 3';
	}
	export class SetCarriers {
		static readonly type = '[Register3] set carriers';
		constructor(public payload: string[]) {}
	}
	export class SetPlatforms {
		static readonly type = '[Register3] set platform';
		constructor(public payload: string[]) {}
	}
	export class FetchCarrierPlatformLists {
		static readonly type = '[Register3] fetch carrier platform lists';
	}
	export class SetCarrierPlatformLists {
		static readonly type = '[Register3] set carrier platform lists';
		constructor(public payload: {carriers: Carriers[], platforms: Platforms[]}) {}
	}

	export type Platforms = NonNullable<NonNullable<NonNullable<FetchPlatformsCarriersQuery['connectOptionPlatforms']>['edges']>[0]>['node'];
	export type Carriers = NonNullable<NonNullable<NonNullable<FetchPlatformsCarriersQuery['connectOptionCarriers']>['edges']>[0]>['node'];

}
