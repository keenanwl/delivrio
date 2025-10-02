import {Component, OnDestroy, OnInit} from '@angular/core';
import {Observable, Subscription} from "rxjs";
import {HypothesisTestingListModel, HypothesisTestingListState} from "./hypothesis-testing-list.ngxs";
import {Store} from "@ngxs/store";
import {HypothesisTestingListActions} from "./hypothesis-testing-list.actions";
import FetchHypothesisTestingList = HypothesisTestingListActions.FetchHypothesisTestingList;
import Reset = HypothesisTestingListActions.Reset;
import {MatDialog} from "@angular/material/dialog";
import {
	AddNewHypothesisTestDialogComponent
} from "./dialogs/add-new-hypothesis-test-dialog/add-new-hypothesis-test-dialog.component";
import {AppActions} from "../../../app.actions";
import AppChangeRoute = AppActions.AppChangeRoute;
import {Paths} from "../../../app-routing.module";

@Component({
	selector: 'app-hypothesis-testing-list',
	templateUrl: './hypothesis-testing-list.component.html',
	styleUrls: ['./hypothesis-testing-list.component.scss']
})
export class HypothesisTestingListComponent implements OnInit, OnDestroy {

	state$: Observable<HypothesisTestingListModel>;

	displayedColumns: string[] = [
		'name',
		'connection',
		'active',
	];

	subscriptions: Subscription[] = [];

	name: number = 0;

	constructor(
		private store: Store,
		private dialog: MatDialog,
	) {
		this.state$ = store.select(HypothesisTestingListState.get);
	}

	ngOnInit() {
		this.store.dispatch(new FetchHypothesisTestingList());
	}

	ngOnDestroy() {
		this.store.dispatch(new Reset());
	}

	addNew() {
		this.dialog.open(AddNewHypothesisTestDialogComponent);
	}

	edit(id: string) {
		this.store.dispatch(new AppChangeRoute({path: Paths.SETTINGS_HYPOTHESIS_TESTING_EDIT, queryParams: {id}}));
	}

}
