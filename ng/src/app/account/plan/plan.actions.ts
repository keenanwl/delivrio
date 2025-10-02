import {FetchPlanListQuery, FetchUserPlanQuery} from "./plan.generated";

export namespace PlanActions {
	export class FetchMyPlan {
		static readonly type = '[Plan] fetch my Plan';
	}
	export class FetchPlanList {
		static readonly type = '[Plan] fetch plan list';
	}
	export class SetMyPlan {
		static readonly type = '[Plan] set my Plan';
		constructor(public payload: SelectPlanQueryResponse | undefined) {}
	}
	export class SetPlanList {
		static readonly type = '[Plan] set language list';
		constructor(public payload: PlanListQueryResponse[]) {}
	}
	export class ChangePlan {
		static readonly type = '[Plan] change plan';
		constructor(public payload: string) {}
	}
	export type SelectPlanQueryResponse = NonNullable<FetchUserPlanQuery['user']>;
	export type PlanListQueryResponse = NonNullable<NonNullable<FetchPlanListQuery['plans']['edges']>[0]>;
}
