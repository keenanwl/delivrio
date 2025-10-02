import {main, printers} from "../../../wailsjs/go/models";

export namespace SettingsActions {
	import USBDevice = main.USBDevice;
	import Printer = printers.Printer;

	export class SetUSBDevices {
		static readonly type = '[Settings] set USB devices';
		constructor(public payload: USBDevice[]) {}
	}
	export class SetPrinters {
		static readonly type = '[Settings] set printers';
		constructor(public payload: Printer[]) {}
	}
	export class Refresh {
		static readonly type = '[Settings] refresh';
	}
	export class RefreshPrinters {
		static readonly type = '[Settings] refresh printers';
	}
	export class TogglePrinter {
		static readonly type = '[Settings] toggle printer';
		constructor(public payload: {name: string; active: boolean}) {}
	}
	export class SaveNetworkPrinter {
		static readonly type = '[Settings] save network printer';
		constructor(public payload: string) {}
	}
	export class ArchivePrinter {
		static readonly type = '[Settings] archive printer';
		constructor(public payload: string) {}
	}
	export class SaveSelectedDevice {
		static readonly type = '[Settings] save selected device';
		constructor(public payload: string) {}
	}
	export class SearchNetworkPrinters {
		static readonly type = '[Settings] search network printers';
	}
	export class FetchRegistrationData {
		static readonly type = '[Settings] fetch registration data';
	}
	export class SetRegistrationData {
		static readonly type = '[Settings] set registration data';
		constructor(public payload: main.RemoteConnectionData) {}
	}
	export class RemoveAllRegistrationData {
		static readonly type = '[Settings] remove all registration data';
	}
	export class ClearRegistrationInfo {
		static readonly type = '[Settings] clear registration info';
	}
	export class SetSearchResults {
		static readonly type = '[Settings] set search results';
		constructor(public payload: main.SubnetSearch[]) {}
	}
}
