import {Component, OnDestroy, OnInit} from '@angular/core';
import {Observable, Subscription} from "rxjs";
import {Paths} from "../../../app-routing.module";
import {Store} from "@ngxs/store";
import {MatDialog} from "@angular/material/dialog";
import {DocumentsListModel, DocumentsListState} from "./documents-list.ngxs";
import {AddDocumentComponent} from "./dialogs/add-document/add-document.component";
import {DocumentsListActions} from "./documents-list.actions";
import FetchDocumentsList = DocumentsListActions.FetchDocumentsList;
import Clear = DocumentsListActions.Clear;

@Component({
  selector: 'app-documents-list',
  templateUrl: './documents-list.component.html',
  styleUrl: './documents-list.component.scss'
})
export class DocumentsListComponent implements OnInit, OnDestroy {

	documents$: Observable<DocumentsListModel>;

	displayedColumns: string[] = [
		'name',
		'carrier',
		'mergeType',
		'createdAt',
	];

	subscriptions: Subscription[] = [];
	paths = Paths;

	constructor(
		private store: Store,
		private dialog: MatDialog,
	) {
		this.documents$ = store.select(DocumentsListState.get);
	}

	ngOnInit() {
		this.store.dispatch(new FetchDocumentsList());
	}

	ngOnDestroy() {
		this.subscriptions.forEach((s) => s.unsubscribe());
		this.store.dispatch(new Clear());
	}

	addNew() {
		this.dialog.open(AddDocumentComponent);
	}
}
