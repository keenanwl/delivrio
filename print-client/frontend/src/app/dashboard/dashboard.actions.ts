import {main} from "../../../wailsjs/go/models";

export namespace DashboardActions {
	import USBDevice = main.USBDevice;
	import PrintJobDisplay = main.PrintJobDisplay;
	import RecentScan = main.RecentScan;

	export class SetUSBDevices {
		static readonly type = '[Dashboard] set USB devices';
		constructor(public payload: USBDevice[]) {}
	}
	export class SetActivePrintJobs {
		static readonly type = '[Dashboard] set active print jobs';
		constructor(public payload: {current: PrintJobDisplay[], recent: PrintJobDisplay[]}) {}
	}
	export class Refresh {
		static readonly type = '[Dashboard] refresh';
	}
	export class SetPrintJobPendingCancel {
		static readonly type = '[Dashboard] set print job pending cancel';
		constructor(public payload: {id: string, messages: string[]}) {}
	}
	export class FetchRecentScans {
		static readonly type = '[Dashboard] fetch recent scans';
	}
	export class SetRecentScans {
		static readonly type = '[Dashboard] set recent scans';
		constructor(public payload: RecentScan[]) {}
	}
}
