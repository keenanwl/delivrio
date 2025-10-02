import {Component, OnInit} from '@angular/core';
import {Observable} from "rxjs";
import {Store} from "@ngxs/store";
import {PlanModel, PlanState} from "./plan.ngxs";
import {PlanActions} from "./plan.actions";
import FetchMyPlan = PlanActions.FetchMyPlan;
import FetchPlanList = PlanActions.FetchPlanList;
import ChangePlan = PlanActions.ChangePlan;

@Component({
	selector: 'app-plan',
	templateUrl: './plan.component.html',
	styleUrls: ['./plan.component.scss']
})
export class PlanComponent implements OnInit {

	plan$: Observable<PlanModel>;

	constructor(private store: Store) {
		this.plan$ = store.select(PlanState.state);
	}

	ngOnInit(): void {
		this.store.dispatch([new FetchMyPlan(), new FetchPlanList()]);
	}

	changePlan(planId: string) {
		this.store.dispatch(new ChangePlan(planId));
	}

}
