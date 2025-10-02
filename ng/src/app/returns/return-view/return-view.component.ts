import {Component, OnInit} from '@angular/core';
import {Observable, Subscription} from "rxjs";
import {Store} from "@ngxs/store";
import {MatDialog} from "@angular/material/dialog";
import {ReturnViewModel, ReturnViewState} from "./return-view.ngxs";
import {ReturnViewActions} from "./return-view.actions";
import FetchReturnView = ReturnViewActions.FetchReturnView;
import SetOrderID = ReturnViewActions.SetOrderID;
import {ActivatedRoute} from "@angular/router";
import ReturnOrderLinesResponse = ReturnViewActions.ReturnOrderLinesResponse;
import {Paths} from "../../app-routing.module";
import {PdfViewerDialogComponent} from "../../shared/pdf-viewer-dialog/pdf-viewer-dialog.component";
import ToggleShowDeleted = ReturnViewActions.ToggleShowDeleted;
import {ReturnColliStatus} from "../../../generated/graphql";
import MarkAccepted = ReturnViewActions.MarkAccepted;
import MarkDeclined = ReturnViewActions.MarkDeclined;

@Component({
	selector: 'app-return-view',
	templateUrl: './return-view.component.html',
	styleUrls: ['./return-view.component.scss']
})
export class ReturnViewComponent implements OnInit {
	returnView$: Observable<ReturnViewModel>;
	returnColliStatusDeleted = ReturnColliStatus.Deleted;

	subscriptions$: Subscription[] = [];

	constructor(
		private route: ActivatedRoute,
		private store: Store,
		private dialog: MatDialog,
	) {
		this.returnView$ = store.select(ReturnViewState.get);
	}

	ngOnInit(): void {
		this.subscriptions$.push(this.route.queryParams
			.subscribe((params) => {
				this.store.dispatch([
					new SetOrderID(!!params.orderID ? params.orderID : ''),
					new FetchReturnView(),
				]);
			}));
	}

	summarizeOrderLines(val: ReturnOrderLinesResponse) {
		return val.map((v) => v.orderLine)
	}

	showPDF(label: string) {
		const ref = this.dialog.open(PdfViewerDialogComponent, {data: {content: "boo"}})
		ref.componentInstance.labelsPDF = [label];
		ref.componentInstance.allPDFs = label;
		ref.componentInstance.title = "Return package label";
	}

	toggleDeleted() {
		this.store.dispatch(new ToggleShowDeleted());
	}

	markAccepted(returnColliID: string) {
		this.store.dispatch(new MarkAccepted(returnColliID))
	}

	markDeclined(returnColliID: string) {
		this.store.dispatch(new MarkDeclined(returnColliID));
	}
}
