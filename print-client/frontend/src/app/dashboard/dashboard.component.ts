import { Component, OnInit } from '@angular/core';
import {DashboardModel, DashboardState} from "./dashboard.ngxs";
import {Observable, timer} from "rxjs";
import {Store} from "@ngxs/store";
import {DashboardActions} from "./dashboard.actions";
import Refresh = DashboardActions.Refresh;
import SetPrintJobPendingCancel = DashboardActions.SetPrintJobPendingCancel;
import FetchRecentScans = DashboardActions.FetchRecentScans;
import {PrintJobStatus} from "../../../../../ng/src/generated/graphql";

@Component({
	selector: 'app-dashboard',
	templateUrl: './dashboard.component.html',
	styleUrls: ['./dashboard.component.scss']
})
export class DashboardComponent implements OnInit {

	dashboard$: Observable<DashboardModel>;
	ticker = 0;

	constructor(
		private store: Store,
	) {
		this.dashboard$ = store.select(DashboardState.get);
	}

	ngOnInit() {
		this.refresh();
		timer(0, 5000).subscribe(() => {
			this.ticker++;
			this.refresh();
		});

	}

	refresh() {
		this.store.dispatch([
			new Refresh(),
			new FetchRecentScans(),
		]);
	}

	dymoCheck() {
		//let result = dymo.label.framework.checkEnvironment()
		//console.warn(result);
	}

	cancelJob(id: string) {
		this.store.dispatch(new SetPrintJobPendingCancel({id: id, messages: ["Cancelled by user from workstation"]}));
	}

	protected readonly PrintJobStatus = PrintJobStatus;
}
