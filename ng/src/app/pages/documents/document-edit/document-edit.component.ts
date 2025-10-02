import {Component, OnDestroy, OnInit} from '@angular/core';
import {Observable, Subscription} from "rxjs";
import {Actions, ofActionCompleted, Store} from "@ngxs/store";
import {ActivatedRoute} from "@angular/router";
import {MatDialog} from "@angular/material/dialog";
import {debounceTime} from "rxjs/operators";
import {DocumentEditActions} from "./document-edit.actions";
import {FormControl, FormGroup} from "@angular/forms";
import {DocumentEditModel, DocumentEditState} from "./document-edit.ngxs";
import {DateTime} from "luxon";
import {DateTimeComponent} from "../../../shared/date-time/date-time.component";
import {DateRange} from "@angular/material/datepicker";
import {ViewDocumentComponent} from "./dialogs/view-document/view-document.component";
import {DocumentMergeType, DocumentPaperSize} from "../../../../generated/graphql";
import FetchDocumentEdit = DocumentEditActions.FetchDocumentEdit;
import SetDocumentID = DocumentEditActions.SetDocumentID;
import SetDocumentEdit = DocumentEditActions.SetDocumentEdit;
import Save = DocumentEditActions.Save;
import SetDateTimeRange = DocumentEditActions.SetDateTimeRange;
import Download = DocumentEditActions.Download;

type varStatus = {[key: string]: boolean};

@Component({
	selector: 'app-document-edit',
	templateUrl: './document-edit.component.html',
	styleUrl: './document-edit.component.scss'
})
export class DocumentEditComponent implements OnInit, OnDestroy {

	form = new FormGroup({
		name: new FormControl(""),
		startAt: new FormControl<string>(DateTime.local().toString()),
		endAt: new FormControl<string>(DateTime.now().toString()),
		htmlTemplate: new FormControl(""),
		htmlHeader: new FormControl(""),
		htmlFooter: new FormControl(""),
		paperSize: new FormControl<DocumentPaperSize>(DocumentPaperSize.A4),
		mergeType: new FormControl<DocumentMergeType>(DocumentMergeType.Orders),
		carrierBrand: new FormGroup({
			id: new FormControl<string | null>(null),
			label: new FormControl(''),
		})
	})

	genericVariables: varStatus = {
		"{{.DocCreatedDate}}": false,
		"{{.DocCreatedTime}}": false,
	}

	packingListVariables: varStatus = {
		"{{.DELIVRIOBarcodeImgTag}}": false,
		"{{.DELIVRIOBarcodeImgSrc}}": false,
		"{{.DELIVRIOBarcode}}": false,
		"{{.OrderPublicID}}": false,
		"{{.OrderCommentExternal}}": false,
		"{{.OrderCommentInternal}}": false,
		"{{index .OrderNoteAttributes \"key\"}}": false,
		"{{.DeliveryOptionName}}": false,
		"{{.DeliveryOptionCarrier}}": false,
	}

	orderLineRowVariables: varStatus = {
		"{{.ProductName}}": false,
		"{{.ProductVariantName}}": false,
		"{{.ProductFirstImageURL}}": false,
		"{{.Quantity}}": false,
		"{{.Price}}": false,
		"{{.Total}}": false,
	}

	customerAddressVariables = {
		"{{.CustomerFirstName}}": false,
		"{{.CustomerLastName}}": false,
		"{{.CustomerCompany}}": false,
		"{{.CustomerEmail}}": false,
		"{{.CustomerAddress1}}": false,
		"{{.CustomerAddress2}}": false,
		"{{.CustomerZip}}": false,
		"{{.CustomerCity}}": false,
		"{{.CustomerState}}": false,
		"{{.CustomerCountry}}": false,
	};

	senderAddressVariables = {
		"{{.SenderFirstName}}": false,
		"{{.SenderLastName}}": false,
		"{{.SenderCompany}}": false,
		"{{.SenderEmail}}": false,
		"{{.SenderAddress1}}": false,
		"{{.SenderAddress2}}": false,
		"{{.SenderZip}}": false,
		"{{.SenderCity}}": false,
		"{{.SenderState}}": false,
		"{{.SenderCountry}}": false,
	};

	ordersVariables: varStatus = {
		"{{.RangeFromDate}}": false,
		"{{.RangeToDate}}": false,
		"{{.ShipmentCount}}": false,
		"{{.CarrierName}}": false,
	}

	orderRowVariables: varStatus = {
		"{{.RowCount}}": false,
		"{{.OrderID}}": false,
		"{{.ShipmentID}}": false,
		"{{.ShipmentTrackingID}}": false,
		"{{.ShipmentCreatedDate}}": false,
	}

	ordersListLoopVariables: varStatus = Object.assign(
		{},
		{"{{range .Orders}}": false},
		this.orderRowVariables,
		{"{{end}}": false},
	);

	orderLineLoopVariables: varStatus = Object.assign(
		{},
		{"{{range .OrderLines}}": false},
		this.orderLineRowVariables,
		{"{{end}}": false},
	);

	allPackingSlipVariables: varStatus = Object.assign(
		{},
		this.genericVariables,
		this.packingListVariables,
		this.customerAddressVariables,
		this.senderAddressVariables,
	);

	allOrdersListVariables: varStatus = Object.assign(
		{},
		this.genericVariables,
		this.ordersVariables,
	);

	loopErr = "";
	htmlTagErr = false;
	subscriptions$: Subscription[] = [];
	state$: Observable<DocumentEditModel>;
	unsortedKVFn = (a: any, b: any) => 0;

	rangeExplanation = `Lines in between "range" and "end" will be repeated for each entity in the list.
		Add HTML around each property for custom styling. "range .**Entity**" and "end"
		lines are required if any sub variables are used from this section.
	`

	constructor(
		private store: Store,
		private route: ActivatedRoute,
		private dialog: MatDialog,
		private actions$: Actions,
	) {
		this.state$ = store.select(DocumentEditState.get);
		// Disallow changing this after creation.
		// Since it is connected to other entities with this type
		// expected.
		this.form.controls.mergeType.disable();
	}

	ngOnInit() {
		this.subscriptions$.push(this.route.queryParams
			.subscribe((params) => {
				this.store.dispatch([
					new SetDocumentID(!!params.id ? params.id : ''),
					new FetchDocumentEdit(),
				]);
			}));

		this.subscriptions$.push(this.actions$
			.pipe(ofActionCompleted(SetDocumentEdit))
			.subscribe(() => this.checkHTMLChanges()));

		this.subscriptions$.push(this.form.controls.htmlTemplate.valueChanges
			.pipe(debounceTime(500))
			.subscribe(() => {
				this.checkHTMLChanges();
			}));

		this.subscriptions$.push(this.form.controls.htmlHeader.valueChanges
			.pipe(debounceTime(500))
			.subscribe(() => {
				this.checkHTMLChanges();
			}));

		this.subscriptions$.push(this.form.controls.htmlFooter.valueChanges
			.pipe(debounceTime(500))
			.subscribe(() => {
				this.checkHTMLChanges();
			}));
	}

	checkHTMLChanges() {

		const header = this.form.controls.htmlHeader.value?.
			replace(/{{\s/g, "{{").replace(/\s}}/g, "}}") || ''
		const footer = this.form.controls.htmlFooter.value?.
			replace(/{{\s/g, "{{").replace(/\s}}/g, "}}") || ''
		const body = this.form.controls.htmlTemplate.value?.
			replace(/{{\s/g, "{{").replace(/\s}}/g, "}}") || '';
		const html = body.concat(header, footer);

		for (let mergeVariablesKey in this.allPackingSlipVariables) {
			this.allPackingSlipVariables[mergeVariablesKey] = !!html?.includes(mergeVariablesKey);
		}

		for (let mergeVariablesKey in this.allOrdersListVariables) {
			this.allOrdersListVariables[mergeVariablesKey] = !!html?.includes(mergeVariablesKey);
		}

		for (let orderLineKey in this.ordersListLoopVariables) {
			this.ordersListLoopVariables[orderLineKey] = !!html?.includes(orderLineKey);
		}

		for (let orderLineKey in this.orderLineLoopVariables) {
			this.orderLineLoopVariables[orderLineKey] = !!html?.includes(orderLineKey);
		}

		const rangeCount = this.countSubstrings(html || '', "{{range")
		const endCount = this.countSubstrings(html || '', "{{end")

		if (rangeCount !== endCount) {
			this.loopErr = `Error: found ${rangeCount} "range" placeholders and ${endCount} "end" placeholders. These should be equal.`
		} else {
			this.loopErr = "";
		}

		const headerHTMLCount = this.countSubstrings(header, "<html>");
		const footerHTMLCount = this.countSubstrings(footer, "<html>");
		const bodyHTMLCount = this.countSubstrings(body, "<html>");

		if ((header.length > 0 && headerHTMLCount === 0) || (footer.length > 0 && footerHTMLCount === 0) || (body.length > 0 && bodyHTMLCount === 0)) {
			this.htmlTagErr = true;
		} else {
			this.htmlTagErr = false;
		}

	}

	countSubstrings(str: string, subStr: string): number {
		let count = 0;
		let index = str.toLowerCase().indexOf(subStr.toLowerCase());
		while (index !== -1) {
			count++;
			index = str.toLowerCase().indexOf(subStr.toLowerCase(), index + subStr.length);
		}
		return count;
	}

	editTime() {
		const ref = this.dialog.open(DateTimeComponent);
		const state = this.store.selectSnapshot(DocumentEditState.get);
		ref.componentInstance.dateRange = new DateRange<DateTime>(DateTime.fromISO(state.form.model?.startAt), DateTime.fromISO(state.form.model?.endAt));
		this.subscriptions$.push(ref.componentInstance.dateRangeSelected
			.subscribe((r) => {
				console.warn(r)
				this.store.dispatch(new SetDateTimeRange({
					start: r.start?.toISO() || DateTime.now().toISO(),
					end: r.end?.toISO() || DateTime.now().toISO(),
				}))
			}));
	}

	save() {
		this.store.dispatch(new Save());
	}

	ngOnDestroy(): void {
		this.subscriptions$.forEach((s) => s.unsubscribe());
	}

	viewDocument() {
		this.dialog.open(ViewDocumentComponent);
		this.store.dispatch(new Download());
	}

	showCarrier(): boolean {
		return this.form.controls.mergeType.value === this.DocumentMergeType.Orders;
	}

	showDateRange(): boolean {
		return this.form.controls.mergeType.value === this.DocumentMergeType.Orders;
	}

	protected readonly DocumentPaperSize = DocumentPaperSize;
	protected readonly DocumentMergeType = DocumentMergeType;
}
