import {FetchHypothesisTestingListQuery} from "./hypothesis-testing-list.generated";

export namespace HypothesisTestingListActions {
	export class FetchHypothesisTestingList {
		static readonly type = '[HypothesisTestingList] fetch HypothesisTestingList';
	}
	export class SetHypothesisTestingList {
		static readonly type = '[HypothesisTestingList] set HypothesisTestingList';
		constructor(public payload: HypothesisTestingResponse[]) {}
	}
	export class SetConnections {
		static readonly type = '[HypothesisTestingList] set connections';
		constructor(public payload: ConnectionsResponse[]) {}
	}
	export class CreateNewTest {
		static readonly type = '[HypothesisTestingList] create new test';
		constructor(public payload: {name: string; connectionID: string;}) {}
	}
	export class Reset {
		static readonly type = '[HypothesisTestingList] reset';
	}
	export type HypothesisTestingResponse = NonNullable<NonNullable<NonNullable<NonNullable<FetchHypothesisTestingListQuery['hypothesisTests']>['edges']>[0]>['node']>;
	export type ConnectionsResponse = NonNullable<NonNullable<NonNullable<NonNullable<FetchHypothesisTestingListQuery['connections']>['edges']>[0]>['node']>;
}
