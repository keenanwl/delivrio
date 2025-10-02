import {Injectable} from "@angular/core";
import {Action, Selector, State, StateContext} from "@ngxs/store";
import {PlanActions} from "./plan.actions";
import SelectPlanQueryResponse = PlanActions.SelectPlanQueryResponse;
import SetMyPlan = PlanActions.SetMyPlan;
import PlanListQueryResponse = PlanActions.PlanListQueryResponse;
import SetPlanList = PlanActions.SetPlanList;
import FetchMyPlan = PlanActions.FetchMyPlan;
import {FetchPlanListGQL, FetchUserPlanGQL, UpdatePlanGQL} from "./plan.generated";

export interface PlanModel {
	myPlan: SelectPlanQueryResponse | null;
	plans: PlanListQueryResponse[];
}

const defaultState: PlanModel = {
	myPlan: null,
	plans: [],
};

@Injectable()
@State<PlanModel>({
	name: 'plan',
	defaults: defaultState,
})
export class PlanState {

	constructor(
		private fetchPlans: FetchPlanListGQL,
		private fetchUserPlan: FetchUserPlanGQL,
		private updatePlan: UpdatePlanGQL,
	) {
	}

	@Selector()
	static state(state: PlanModel) {
		return state;
	}

	@Action(PlanActions.FetchMyPlan)
	FetchMyPlan(ctx: StateContext<PlanModel>, action: PlanActions.FetchMyPlan) {
		return this.fetchUserPlan.fetch({}, {fetchPolicy: "no-cache"})
			.subscribe({next: (r) => {
				const user = r.data.user;
				if (!!user) {
					ctx.dispatch(new SetMyPlan(user));
				}
			}, error: (e) => {

			}});
	}

	@Action(PlanActions.FetchPlanList)
	FetchPlanList(ctx: StateContext<PlanModel>, action: PlanActions.FetchPlanList) {
		return this.fetchPlans.fetch()
			.subscribe((r) => {
				const plans: PlanListQueryResponse[] = [];
				r.data.plans.edges?.forEach((p) => {
					if (!!p) {
						plans.push(p);
					}
				})
				ctx.dispatch(new SetPlanList(plans));
			});
	}

	@Action(PlanActions.SetMyPlan)
	SetMyPlan(ctx: StateContext<PlanModel>, action: PlanActions.SetMyPlan) {
		ctx.patchState({
			myPlan: action.payload,
		})
	}

	@Action(PlanActions.SetPlanList)
	SetLanguageList(ctx: StateContext<PlanModel>, action: PlanActions.SetPlanList) {
		ctx.patchState({
			plans: action.payload,
		})
	}

	@Action(PlanActions.ChangePlan)
	SaveForm(ctx: StateContext<PlanModel>, action: PlanActions.ChangePlan) {
		return this.updatePlan.mutate({planID: action.payload}).subscribe({
			next: () => ctx.dispatch(new FetchMyPlan()),
			error: (e) => {
				console.log(e)
			}
		});
	}

}
