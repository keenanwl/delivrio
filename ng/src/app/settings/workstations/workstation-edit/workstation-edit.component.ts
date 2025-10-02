import {Component, OnDestroy, OnInit} from '@angular/core';
import {Observable, Subscription, timer} from "rxjs";
import {WorkstationEditModel, WorkstationEditState} from "./workstation-edit.ngxs";
import {Actions, ofActionCompleted, Store} from "@ngxs/store";
import {ActivatedRoute} from "@angular/router";
import {WorkstationEditActions} from "./workstation-edit.actions";
import {FormArray, FormControl, FormGroup} from "@angular/forms";
import {PrinterPrintSize} from "../../../../generated/graphql";
import SetWorkstationID = WorkstationEditActions.SetWorkstationID;
import FetchWorkstationEdit = WorkstationEditActions.FetchWorkstationEdit;
import SetWorkstationEdit = WorkstationEditActions.SetWorkstationEdit;
import Save = WorkstationEditActions.Save;
import PrinterResponse = WorkstationEditActions.PrinterResponse;
import Disable = WorkstationEditActions.Disable;
import Reset = WorkstationEditActions.Reset;

@Component({
	selector: 'app-workstation-edit',
	templateUrl: './workstation-edit.component.html',
	styleUrls: ['./workstation-edit.component.scss']
})
export class WorkstationEditComponent implements OnInit, OnDestroy {

	workstationEdit$: Observable<WorkstationEditModel>;

	editForm = new FormGroup({
		name: new FormControl('', {nonNullable: true}),
		autoPrintReceiver: new FormControl(false, {nonNullable: true}),
		lastPing: new FormControl('', {nonNullable: true}),
		printer: new FormArray<ReturnType<typeof this.newPrinterRow>>([])
	});

	paperSizeA4: PrinterPrintSize = PrinterPrintSize.A4;
	paperSize100x150: PrinterPrintSize = PrinterPrintSize.Cm_100_150;
	paperSize100x192: PrinterPrintSize = PrinterPrintSize.Cm_100_192;

	ticker = 0
	subscriptions$: Subscription[] = [];

	constructor(
		private store: Store,
		private route: ActivatedRoute,
		private actions$: Actions,
	) {
		this.workstationEdit$ = store.select(WorkstationEditState.get);
	}

	ngOnInit(): void {
		// TODO: fetch the last ping more frequently
		this.subscriptions$.push(timer(0, 1000).subscribe(() => this.ticker++));

		this.subscriptions$.push(this.route.queryParams
			.subscribe((params) => {
				this.store.dispatch([
					new SetWorkstationID(!!params.id ? params.id : ''),
					new FetchWorkstationEdit(),
				]);
			}));

		this.actions$.pipe(ofActionCompleted(SetWorkstationEdit))
			.subscribe((r) => {
				this.editForm.controls.printer.clear();
				const printers = this.store.selectSnapshot(WorkstationEditState.get).workstationEditForm.model?.printer
				printers?.forEach((p) => {
					this.editForm.controls.printer.push(this.newPrinterRow(p));
				});
			});
	}

	ngOnDestroy() {
		this.subscriptions$.forEach((s) => s.unsubscribe());
		this.store.dispatch(new Reset());
	}

	newPrinterRow(p: PrinterResponse) {
		return new FormGroup({
			id: new FormControl(p.id, {nonNullable: true}),
			name: new FormControl(p.name, {nonNullable: true}),
			lastPing: new FormControl<string>(p.lastPing, {nonNullable: true}),
			printSize: new FormControl(p.printSize, {nonNullable: true}),
			labelZpl: new FormControl<boolean>(p.labelZpl, {nonNullable: true}),
			labelPdf: new FormControl<boolean>(p.labelPdf, {nonNullable: true}),
			labelPng: new FormControl<boolean>(p.labelPng, {nonNullable: true}),
			useShell: new FormControl<boolean>(p.useShell, {nonNullable: true}),
			document: new FormControl<boolean>(p.document, {nonNullable: true}),
			rotate180: new FormControl<boolean>(p.rotate180, {nonNullable: true}),
		});
	}

	save() {
		this.store.dispatch(new Save());
	}

	disable() {
		this.store.dispatch(new Disable());
	}

}
