import {FetchConsolidationsQuery} from "./consolidations-list.generated";

export namespace ConsolidationsListActions {
	export class FetchConsolidationsList {
		static readonly type = '[ConsolidationsList] fetch consolidations';
	}
	export class SetConsolidationsList {
		static readonly type = '[ConsolidationsList] set consolidations';
		constructor(public payload: ConsolidationRecord[]) {}
	}
	export class AddConsolidation {
		static readonly type = '[ConsolidationsList] add consolidation';
		constructor(public payload: {publicID: string; description: string}) {}
	}
	export class Clear {
		static readonly type = '[ConsolidationsList] clear';
	}

	export type ConsolidationRecord = NonNullable<NonNullable<NonNullable<FetchConsolidationsQuery['consolidations']>['edges']>[0]>['node'];
}
