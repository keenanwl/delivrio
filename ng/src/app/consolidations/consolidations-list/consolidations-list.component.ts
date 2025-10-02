import {Component, OnDestroy, OnInit} from '@angular/core';
import {Observable, Subscription} from "rxjs";
import {Paths} from "../../app-routing.module";
import {Store} from "@ngxs/store";
import {ConsolidationsListActions} from "./consolidations-list.actions";
import Clear = ConsolidationsListActions.Clear;
import FetchConsolidationsList = ConsolidationsListActions.FetchConsolidationsList;
import {ConsolidationsListModel, ConsolidationsListState} from "./consolidations-list.ngxs";
import {MatDialog} from "@angular/material/dialog";
import {AddConsolidationComponent} from "./dialogs/add-consolidation/add-consolidation.component";

@Component({
	selector: 'app-consolidations-list',
	templateUrl: './consolidations-list.component.html',
	styleUrl: './consolidations-list.component.scss'
})
export class ConsolidationsListComponent implements OnInit, OnDestroy {

	displayedColumns: string[] = [
		'createdAt',
		'name',
		'description',
		'status',
		'pallets',
		'orders',
	];

	consolidationList$: Observable<ConsolidationsListModel>;
	subscriptions$: Subscription[] = [];
	paths = Paths;

	constructor(
		private store: Store,
		private dialog: MatDialog,
	) {
		this.consolidationList$ = store.select(ConsolidationsListState.get);
	}

	ngOnInit(): void {
		this.store.dispatch(new FetchConsolidationsList());
	}

	ngOnDestroy(): void {
		this.subscriptions$.forEach((s) => s.unsubscribe());
		this.store.dispatch(new Clear());
	}

	add() {
		this.dialog.open(AddConsolidationComponent);
	}

}
