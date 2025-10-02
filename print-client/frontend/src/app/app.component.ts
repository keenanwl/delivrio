import {Component, OnInit} from '@angular/core';
import {Actions, ofActionDispatched, Store} from "@ngxs/store";
import {AppActions} from "./app.actions";
import AppChangeRoute = AppActions.AppChangeRoute;
import ShowGlobalSnackbar = AppActions.ShowGlobalSnackbar;
import {MatSnackBar} from "@angular/material/snack-bar";
import {Router} from "@angular/router";
import FetchIsRegistered = AppActions.FetchIsRegistered;
import {Observable, of, timer} from "rxjs";
import {AppModel, AppState} from "./app.ngxs";
import FetchWorkstationName = AppActions.FetchWorkstationName;
import {EventsOn} from "../../wailsjs/runtime";
import {DYMOPngRegisterStatus} from "../../wailsjs/go/main/App";
import {printer} from "../../wailsjs/go/models";

declare var dymo: any;

interface DYMOPngInput {
	base64Data: string;
	id: string;
}

interface DYMOPngOutput {
	msg: string;
	id: string;
	success: boolean;
}

@Component({
	selector: 'app-root',
	templateUrl: './app.component.html',
	styleUrls: ['./app.component.scss']
})
export class AppComponent implements OnInit {

	app$: Observable<AppModel>;

	constructor(
		private snackBar: MatSnackBar,
		private router: Router,
		private store: Store,
		private actions$: Actions
	) {
		this.app$ = store.select(AppState.get);
		this.printBase64PngLabel("123")
	}

	ngOnInit(): void {


		EventsOn("token-not-recognized", () => {
			this.store.dispatch(new ShowGlobalSnackbar("Authentication token for this workstation is invalid"))
		});

		EventsOn("registration-changed", () => {
			this.store.dispatch(new FetchIsRegistered())
		});

		this.store.dispatch([
			new FetchIsRegistered(),
			new FetchWorkstationName(),
		]);

		this.actions$
			.pipe(ofActionDispatched(AppChangeRoute))
			.subscribe(({payload}) => this.router.navigate(
				[payload.path],
				{queryParams: payload.queryParams}
			));

		this.actions$
			.pipe(ofActionDispatched(ShowGlobalSnackbar))
			.subscribe(({payload}) => {
				this.snackBar.open(payload, `close`,{duration: 3000});
			});


		(window as any).runtime.EventsOn("dymo-png", (input: DYMOPngInput) => {
			const res = this.printBase64PngLabel(input.base64Data)
			const out: DYMOPngOutput = {
				id: input.id,
				msg: res.msg,
				success: res.success,
			}
console.warn(out)
			DYMOPngRegisterStatus(JSON.stringify(out)).then(r => console.log(r))
		})
	}

	initPngLabel(base64Png: string): {success: boolean; msg: string} {

		timer(1000).subscribe(() => {
			return this.printBase64PngLabel(base64Png)
		})

		return {success: true, msg: ""}

	}

	printBase64PngLabel(base64Png: string): { success: boolean; msg: string } {

		try {
			dymo.label.framework.init()
			// Create a new label with 4" x 6" dimensions
			// 4 inches = 5760 twips (1440 twips per inch)
			// 6 inches = 8640 twips
			const label = dymo.label.framework.openLabelXml(
				// XML needs to be the same line
				`<?xml version="1.0" encoding="utf-8"?>
    <DieCutLabel Version="8.0" Units="twips">
        <PaperOrientation>Portrait</PaperOrientation>
        <Id>LargeShipping</Id>
        <PaperName>30256 Shipping</PaperName>
        <DrawCommands>
            <RoundRectangle X="0" Y="0" Width="5760" Height="8640" Rx="0" Ry="0" />
        </DrawCommands>
        <ObjectInfo>
         <ImageObject>
             <Name>Graphic</Name>
             <ForeColor Alpha="255" Red="0" Green="0" Blue="0" />
             <BackColor Alpha="0" Red="255" Green="255" Blue="255" />
             <LinkedObjectName></LinkedObjectName>
             <Rotation>Rotation0</Rotation>
             <IsMirrored>False</IsMirrored>
             <IsVariable>False</IsVariable>
             <Image></Image>
             <ScaleMode>Uniform</ScaleMode>
             <BorderWidth>0</BorderWidth>
             <BorderColor Alpha="255" Red="0" Green="0" Blue="0" />
             <HorizontalAlignment>Center</HorizontalAlignment>
             <VerticalAlignment>Center</VerticalAlignment>
         </ImageObject>
         <Bounds X="0" Y="0" Width="5760" Height="8640" />
     </ObjectInfo>
    </DieCutLabel>
			`);

			// Set the image source using base64 data
			label.setObjectText("Graphic", `${base64Png}`);
			label.render()

			const printers = dymo.label.framework.getPrinters();
			if (printers.length === 0) {
				return {success: false, msg: 'No DYMO printers found'}
			}
			// This disregards the selection made in DELIVRIO client
			const printer = printers[0];

			label.print(printer.name)

		} catch (error) {
			return {success: false, msg: 'error printing label: ' + error}
		}
		return {success: true, msg: ''}
	}
}
