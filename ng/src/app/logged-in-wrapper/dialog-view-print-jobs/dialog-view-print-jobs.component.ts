import {Component, OnDestroy, OnInit} from '@angular/core';
import {Observable, Subscription, timer} from "rxjs";
import {AppModel, AppState} from "../../app.ngxs";
import {Store} from "@ngxs/store";
import {MatDialogRef} from "@angular/material/dialog";
import {PrintJobStatus} from "../../../generated/graphql";

@Component({
  selector: 'app-dialog-view-print-jobs',
  templateUrl: './dialog-view-print-jobs.component.html',
  styleUrls: ['./dialog-view-print-jobs.component.scss']
})
export class DialogViewPrintJobsComponent implements OnInit, OnDestroy {

	app$: Observable<AppModel>;
	ticker = 0
	subscriptions$: Subscription[] = [];

	constructor(private store: Store, private ref: MatDialogRef<any>) {
		this.app$ = store.select(AppState.get);
	}

	ngOnInit() {
		this.subscriptions$.push(timer(0, 1000).subscribe(() => this.ticker++));
	}

	ngOnDestroy() {
		this.subscriptions$.map((s) => s.unsubscribe());
	}

	close() {
		this.ref.close();
	}

	protected readonly PrintJobStatus = PrintJobStatus;
}
