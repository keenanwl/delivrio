import {HistoryLogsQuery} from "./history-logs.generated";

export namespace HistoryLogsActions {
	export class FetchHistoryLogs {
		static readonly type = '[HistoryLogs] fetch history logs';
	}
	export class SetHistories {
		static readonly type = '[HistoryLogs] set histories';
		constructor(public payload: HistoryRecord[]) {}
	}
	export class SetLogs {
		static readonly type = '[HistoryLogs] set logs';
		constructor(public payload: LogRecord[]) {}
	}

	export type HistoryRecord = NonNullable<NonNullable<NonNullable<HistoryLogsQuery['historyLogs']>['histories']>[0]>;
	export type LogRecord = NonNullable<NonNullable<NonNullable<HistoryLogsQuery['historyLogs']>['system_event']>[0]>;
}
