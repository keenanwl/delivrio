import {Injectable} from "@angular/core";
import {Action, Selector, State, StateContext} from "@ngxs/store";
import {
	ArchivePrinter,
	Printers,
	SaveNetworkPrinter,
	TogglePrinter,
	USBDevices,
	SaveSelectedDevice,
	FindNetworkDevices, RegistrationData, RemoveRegistrationData
} from "../../../wailsjs/go/main/App";
import {main, printers} from "../../../wailsjs/go/models";
import USBDevice = main.USBDevice;
import {SettingsActions} from "./settings.actions";
import SetUSBDevices = SettingsActions.SetUSBDevices;
import Printer = printers.Printer;
import SetPrinters = SettingsActions.SetPrinters;
import SetRegistrationData = SettingsActions.SetRegistrationData;
import {AppActions} from "../app.actions";
import ShowGlobalSnackbar = AppActions.ShowGlobalSnackbar;
import ClearRegistrationInfo = SettingsActions.ClearRegistrationInfo;
import FetchIsRegistered = AppActions.FetchIsRegistered;
import FetchWorkstationName = AppActions.FetchWorkstationName;

export interface SettingsModel {
	devices: USBDevice[];
	printers: Printer[];
	loading: boolean;
	searchResults: main.SubnetSearch[];
	remoteData: main.RemoteConnectionData | undefined;
}

const defaultState: SettingsModel = {
	devices: [],
	printers: [],
	loading: false,
	searchResults: [],
	remoteData: undefined,
};

@Injectable()
@State<SettingsModel>({
	name: 'settings',
	defaults: defaultState,
})
export class SettingsState {

	constructor() {
	}

	@Selector()
	static get(state: SettingsModel) {
		return state;
	}

	@Action(SettingsActions.Refresh)
	Refresh(ctx: StateContext<SettingsModel>, action: SettingsActions.Refresh) {
		return USBDevices().then((devices) => {
			ctx.dispatch(new SetUSBDevices(devices));
		});
	}

	@Action(SettingsActions.RefreshPrinters)
	RefreshPrinters(ctx: StateContext<SettingsModel>, action: SettingsActions.RefreshPrinters) {
		return Printers().then((pl) => {
			ctx.dispatch(new SetPrinters(pl));
		})
	}

	@Action(SettingsActions.SetUSBDevices)
	SetUSBDevices(ctx: StateContext<SettingsModel>, action: SettingsActions.SetUSBDevices) {
		ctx.patchState({devices: action.payload});
	}

	@Action(SettingsActions.SetPrinters)
	SetPrinters(ctx: StateContext<SettingsModel>, action: SettingsActions.SetPrinters) {
		ctx.patchState({printers: action.payload});
	}

	@Action(SettingsActions.TogglePrinter)
	TogglePrinter(ctx: StateContext<SettingsModel>, action: SettingsActions.TogglePrinter) {
		return TogglePrinter(action.payload.name, action.payload.active)
			.then((r) => {

				let message = "Saved successfully";
				if (r.length > 0) {
					message = r;
				}
				ctx.dispatch(new ShowGlobalSnackbar(message));

			})
	}

	@Action(SettingsActions.FetchRegistrationData)
	FetchRegistrationData(ctx: StateContext<SettingsModel>, action: SettingsActions.FetchRegistrationData) {
		return RegistrationData()
			.then((r) => {
				ctx.dispatch(new SetRegistrationData(r));
			}, (e) => {
				ctx.dispatch(new ShowGlobalSnackbar("Error: " + e));
			})
	}

	@Action(SettingsActions.SaveNetworkPrinter)
	SaveNetworkPrinter(ctx: StateContext<SettingsModel>, action: SettingsActions.SaveNetworkPrinter) {
		return SaveNetworkPrinter(action.payload).then((r) => {
			if (r.length === 0) {
				ctx.dispatch(new ShowGlobalSnackbar("Saved successfully"));
			} else {
				ctx.dispatch(new ShowGlobalSnackbar("Error: " + r));
			}
		});
	}

	@Action(SettingsActions.ArchivePrinter)
	ArchivePrinter(ctx: StateContext<SettingsModel>, action: SettingsActions.ArchivePrinter) {
		return ArchivePrinter(action.payload).then((r) => {
			if (r.length === 0) {
				ctx.dispatch(new ShowGlobalSnackbar("Archived successfully"));
			} else {
				ctx.dispatch(new ShowGlobalSnackbar("Error: " + r));
			}
		});
	}

	@Action(SettingsActions.SaveSelectedDevice)
	SaveSelectedDevice(ctx: StateContext<SettingsModel>, action: SettingsActions.SaveSelectedDevice) {
		return SaveSelectedDevice(action.payload).then((r) => {
			if (r.length === 0) {
				ctx.dispatch(new ShowGlobalSnackbar("Saved successfully"));
			} else {
				ctx.dispatch(new ShowGlobalSnackbar("Error: " + r));
			}
		});
	}

	@Action(SettingsActions.SearchNetworkPrinters)
	SearchNetworkPrinters(ctx: StateContext<SettingsModel>, action: SettingsActions.SearchNetworkPrinters) {
		ctx.patchState({loading: true});
		return FindNetworkDevices("9100").then((r) => {
			ctx.dispatch(new SettingsActions.SetSearchResults(r));
			ctx.patchState({loading: false});
		}, (e) => {
			ctx.patchState({loading: false});
			ctx.dispatch(new ShowGlobalSnackbar("Error: " + e));
		})
	}

	@Action(SettingsActions.SetSearchResults)
	SetSearchResults(ctx: StateContext<SettingsModel>, action: SettingsActions.SetSearchResults) {
		ctx.patchState({searchResults: action.payload});
	}

	@Action(SettingsActions.SetRegistrationData)
	SetRegistrationData(ctx: StateContext<SettingsModel>, action: SettingsActions.SetRegistrationData) {
		ctx.patchState({remoteData: action.payload});
	}

	@Action(SettingsActions.ClearRegistrationInfo)
	ClearRegistrationInfo(ctx: StateContext<SettingsModel>, action: SettingsActions.ClearRegistrationInfo) {
		ctx.patchState({remoteData: undefined});
	}

	@Action(SettingsActions.RemoveAllRegistrationData)
	RemoveAllRegistrationData(ctx: StateContext<SettingsModel>, action: SettingsActions.RemoveAllRegistrationData) {
		return RemoveRegistrationData().then(() => {
			ctx.dispatch([
				new ClearRegistrationInfo(),
				new FetchIsRegistered(),
				new FetchWorkstationName(),
				new ShowGlobalSnackbar("Registration data removed!"),
			]);
		});
	}

}
