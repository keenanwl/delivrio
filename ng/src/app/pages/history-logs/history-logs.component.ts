import {Component, OnInit} from '@angular/core';
import {CommonModule} from '@angular/common';
import {Observable} from "rxjs";
import {Store} from "@ngxs/store";
import {HistoryLogsModel, HistoryLogsState} from "./history-logs.ngxs";
import {HistoryLogsActions} from "./history-logs.actions";
import FetchHistoryLogs = HistoryLogsActions.FetchHistoryLogs;
import {RelativeTimePipe} from "../../pipes/relative-time.pipe";
import {DvoCardComponent} from "../../shared/dvo-card/dvo-card.component";
import {MaterialModule} from "../../modules/material.module";
import {SystemEventsEventType} from "../../../generated/graphql";

@Component({
	selector: 'app-history-logs',
	standalone: true,
	imports: [CommonModule, RelativeTimePipe, DvoCardComponent, MaterialModule],
	templateUrl: './history-logs.component.html',
	styleUrl: './history-logs.component.scss'
})
export class HistoryLogsComponent implements OnInit {
	state$: Observable<HistoryLogsModel>;

	logEventType = SystemEventsEventType

	constructor(private store: Store) {
		this.state$ = store.select(HistoryLogsState.get);
	}

	ngOnInit() {
		this.store.dispatch(new FetchHistoryLogs());
	}

	refresh() {
		this.store.dispatch(new FetchHistoryLogs());
	}
}
