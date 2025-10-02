import {Component, OnDestroy, OnInit} from '@angular/core';
import {Store} from "@ngxs/store";
import {AppModel, AppState} from "../app.ngxs";
import {Observable, Subscription, timer} from "rxjs";
import {AppActions} from "../app.actions";
import {FormControl} from '@angular/forms';
import {debounceTime} from "rxjs/operators";
import {EntityType, SearchResult} from "../../generated/graphql";
import Search = AppActions.Search;
import AppChangeRoute = AppActions.AppChangeRoute;
import {Paths} from "../app-routing.module";
import ClearSearchResults = AppActions.ClearSearchResults;
import {MatDialog} from "@angular/material/dialog";
import {DialogSelectPrinterComponent} from "./dialog-select-printer/dialog-select-printer.component";
import FetchSelectedWorkstation = AppActions.FetchSelectedWorkstation;
import {DateTime} from "luxon";
import {DialogViewPrintJobsComponent} from "./dialog-view-print-jobs/dialog-view-print-jobs.component";

@Component({
	selector: 'app-logged-in-wrapper',
	templateUrl: './logged-in-wrapper.component.html',
	styleUrls: ['./logged-in-wrapper.component.scss']
})
export class LoggedInWrapperComponent implements OnInit, OnDestroy {

	app$: Observable<AppModel>;

	searchControl = new FormControl('', {nonNullable: true});
	subscriptions$: Subscription[] = [];

	constructor(
		private store: Store,
		private dialog: MatDialog,
	) {
		this.app$ = store.select(AppState.get);
	}

	ngOnInit(): void {
		// Disabled for limited system

		this.subscriptions$.push(
			timer(0, 15_000).subscribe(() => {
				this.store.dispatch(new FetchSelectedWorkstation());
			}))

		this.subscriptions$.push(this.searchControl.valueChanges.pipe(debounceTime(100))
			.subscribe((s) => {
				if (s.length > 0) {
					this.store.dispatch(new Search(s));
				}
			}));
	}

	ngOnDestroy() {
		this.subscriptions$.forEach((s) => s.unsubscribe());
	}

	selectSearchResult(result: SearchResult) {
		this.resetSearch();

		switch (result.entity) {
			case EntityType.Order:
				this.store.dispatch(new AppChangeRoute({
					path: Paths.ORDERS_VIEW, queryParams: {id: result.id},
				}));
				break;
			case EntityType.Product:
				this.store.dispatch(new AppChangeRoute({
					path: Paths.PRODUCTS_EDIT, queryParams: {id: result.id},
				}));
				break;
		}
	}

	resetSearch() {
		this.searchControl.reset();
		this.store.dispatch(new ClearSearchResults());
	}

	showPrinterDialog() {
		this.dialog.open(DialogSelectPrinterComponent)
	}

	showPrintJobsDialog() {
		this.dialog.open(DialogViewPrintJobsComponent)
	}

	isPrinterActive(lastPing?: string): boolean {
		if (!lastPing) {
			return false
		}

		const date = DateTime.fromISO(lastPing)
		if (!date) {
			return false
		}

		const currentDate = DateTime.now();
		const secondsDifference = currentDate.diff(date, 'seconds').seconds;
		return secondsDifference >= 0 && secondsDifference <= 30;
	}

}
