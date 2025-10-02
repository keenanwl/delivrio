import {Component, OnDestroy, OnInit} from '@angular/core';
import {Observable, Subscription} from "rxjs";
import {ReturnsListModel, ReturnsListState} from "./returns-list.ngxs";
import {Store} from "@ngxs/store";
import {MatDialog} from "@angular/material/dialog";
import {NewReturnDialogComponent} from "./dialogs/new-return-dialog/new-return-dialog.component";
import {ReturnsListActions} from "./returns-list.actions";
import Clear = ReturnsListActions.Clear;
import FetchReturnsList = ReturnsListActions.FetchReturnsList;
import {AppActions} from "../../app.actions";
import AppChangeRoute = AppActions.AppChangeRoute;
import {Paths} from "../../app-routing.module";
import MarkReturnColliDeleted = ReturnsListActions.MarkReturnColliDeleted;
import {
	ConfirmDeleteReturnColliComponent
} from "./dialogs/confirm-delete-return-colli/confirm-delete-return-colli.component";

@Component({
	selector: 'app-returns-list',
	templateUrl: './returns-list.component.html',
	styleUrls: ['./returns-list.component.scss']
})
export class ReturnsListComponent implements OnInit, OnDestroy {

	returnsList$: Observable<ReturnsListModel>;

	subscriptions$: Subscription[] = []

	displayedColumns: string[] = [
		'orderPublicID',
		'createdAt',
		'name',
		'returnColliStatus',
		'countries',
		'actions',
	];

	constructor(
		private store: Store,
		private dialog: MatDialog,
	) {
		this.returnsList$ = store.select(ReturnsListState.get);
	}

	ngOnInit() {
		this.store.dispatch(new FetchReturnsList());
	}

	addNewReturn() {
		this.dialog.open(NewReturnDialogComponent);
	}

	editReturn(id: string) {
		this.store.dispatch(new AppChangeRoute({path: Paths.RETURN_VIEW, queryParams: {orderID: id}}));
	}

	ngOnDestroy(): void {
		this.store.dispatch(new Clear());
		this.subscriptions$.forEach((s) => s.unsubscribe());
	}

	confirmDelete(evt: MouseEvent, returnColliID: string) {
		evt.preventDefault();
		evt.cancelBubble = true;

		const ref = this.dialog.open(ConfirmDeleteReturnColliComponent);
		this.subscriptions$.push(ref.componentInstance.confirm.subscribe((r) => {
			if (r) {
				this.store.dispatch(new MarkReturnColliDeleted({returnColliID}));
			}
		}));
	}
}
