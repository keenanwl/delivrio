import {FetchDashboardTilesQuery} from "./dashboard.generated";

export namespace DashboardActions {
	export class FetchTiles {
		static readonly type = '[Dashboard] fetch tiles';
	}
	export class SetTiles {
		static readonly type = '[Dashboard] set tiles';
		constructor(public payload: TileResponse[]) {}
	}
	export class SetHypothesisTest {
		static readonly type = '[Dashboard] set hypothesis tests';
		constructor(public payload: HypothesisTestResponse[]) {}
	}
	export class SetTrailingProductUpdateCounts {
		static readonly type = '[Dashboard] set trailing product update counts';
		constructor(public payload: ProductUpdateCount[]) {}
	}
	export class SetRateRequests {
		static readonly type = '[Dashboard] set rate requests';
		constructor(public payload: RateRequests) {}
	}

	export type TileResponse = NonNullable<NonNullable<FetchDashboardTilesQuery['dashboardTiles']>[0]>;
	export type HypothesisTestResponse = NonNullable<NonNullable<FetchDashboardTilesQuery['hypothesisTestResultsDashboard']>[0]>;
	export type RateRequests = NonNullable<FetchDashboardTilesQuery['rateRequests']>;
	export type ProductUpdateCount = number;
}
