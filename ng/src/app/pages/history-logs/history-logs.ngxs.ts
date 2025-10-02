import {Injectable} from "@angular/core";
import {Action, Selector, State, StateContext, Store} from "@ngxs/store";
import {HistoryLogsActions} from "./history-logs.actions";
import HistoryRecord = HistoryLogsActions.HistoryRecord;
import LogRecord = HistoryLogsActions.LogRecord;
import {HistoryLogsGQL} from "./history-logs.generated";
import SetHistories = HistoryLogsActions.SetHistories;
import SetLogs = HistoryLogsActions.SetLogs;
import {toNotNullArray} from "../../functions/not-null-array";

export interface HistoryLogsModel {
	histories: HistoryRecord[];
	logs: LogRecord[];
	loading: boolean;
}

const defaultState: HistoryLogsModel = {
	histories: [],
	logs: [],
	loading: false,
};

@Injectable()
@State<HistoryLogsModel>({
	name: 'historyLogs',
	defaults: defaultState,
})
export class HistoryLogsState {

	constructor(
		private store: Store,
		private fetch: HistoryLogsGQL,
	) {}

	@Selector()
	static get(state: HistoryLogsModel) {
		return state;
	}

	@Action(HistoryLogsActions.FetchHistoryLogs)
	FetchHistoryLogs(ctx: StateContext<HistoryLogsModel>, action: HistoryLogsActions.FetchHistoryLogs) {
		ctx.patchState({loading: true});
		return this.fetch.fetch({}, {errorPolicy: "all"})
			.subscribe((r) => {
				ctx.dispatch([
					new SetHistories(toNotNullArray(r.data.historyLogs.histories)),
					new SetLogs(toNotNullArray(r.data.historyLogs.system_event)),
				])
				ctx.patchState({loading: false});
			});
	}

	@Action(HistoryLogsActions.SetHistories)
	SetHypothesisTest(ctx: StateContext<HistoryLogsModel>, action: HistoryLogsActions.SetHistories) {
		ctx.patchState({histories: action.payload})
	}

	@Action(HistoryLogsActions.SetLogs)
	SetLogs(ctx: StateContext<HistoryLogsModel>, action: HistoryLogsActions.SetLogs) {
		ctx.patchState({logs: action.payload})
	}

}
