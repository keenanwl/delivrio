import {FetchReturnPortalsQuery} from "./return-portals-list.generated";

export namespace ReturnPortalsListActions {
	export class FetchReturnPortalsList {
		static readonly type = '[ReturnPortalsList] fetch ReturnPortalsList';
	}
	export class SetReturnPortalsList {
		static readonly type = '[ReturnPortalsList] set ReturnPortalsList';
		constructor(public payload: ReturnPortalsListResponse[]) {}
	}
	export class SetConnections {
		static readonly type = '[ReturnPortalsList] set connections';
		constructor(public payload: ConnectionsResponse[]) {}
	}
	export class Create {
		static readonly type = '[ReturnPortalsList] create';
		constructor(public payload: {name: string, connection: string}) {}
	}
	export class Clear {
		static readonly type = '[ReturnPortalsList] clear';
	}
	export type ReturnPortalsListResponse = NonNullable<NonNullable<NonNullable<NonNullable<FetchReturnPortalsQuery['returnPortals']>['edges']>[0]>['node']>;
	export type ConnectionsResponse = NonNullable<NonNullable<NonNullable<NonNullable<FetchReturnPortalsQuery['connections']>['edges']>[0]>['node']>;
}
