import {ListConnectionsQuery} from "./connections-list.generated";

export namespace ConnectionsListActions {
	export class FetchConnectionsList {
		static readonly type = '[ConnectionsList] fetch ConnectionsList';
	}
	export class SetConnectionsList {
		static readonly type = '[ConnectionsList] set ConnectionsList';
		constructor(public payload: ConnectionResponse[]) {}
	}
	export class Clear {
		static readonly type = '[ConnectionsList] clear';
	}
	export type ConnectionResponse = NonNullable<NonNullable<NonNullable<ListConnectionsQuery['connections']>['edges']>[0]>['node'];
}
