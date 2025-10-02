import {Component, OnDestroy, OnInit} from '@angular/core';
import {Observable, Subscription} from "rxjs";
import {Store} from "@ngxs/store";
import {EmailTemplatesListModel, EmailTemplatesListState} from "./email-templates-list.ngxs";
import {EmailTemplatesListActions} from "./email-templates-list.actions";
import FetchEmailTemplatesList = EmailTemplatesListActions.FetchEmailTemplatesList;
import {MatDialog} from "@angular/material/dialog";
import {AddEmailTemplateComponent} from "./dialogs/add-email-template/add-email-template.component";
import {Paths} from "../../../app-routing.module";
import {AppActions} from "../../../app.actions";

@Component({
	selector: 'app-email-templates-list',
	templateUrl: './email-templates-list.component.html',
	styleUrls: ['./email-templates-list.component.scss']
})
export class EmailTemplatesListComponent implements OnInit, OnDestroy {

	emails$: Observable<EmailTemplatesListModel>;

	displayedColumns: string[] = [
		'name',
		'subject',
		'mergeType',
	];

	subscriptions: Subscription[] = [];
	editPath = Paths.SETTINGS_EMAIL_TEMPLATE_EDIT;

	constructor(
		private store: Store,
		private dialog: MatDialog,
	) {
		this.emails$ = store.select(EmailTemplatesListState.get);
	}

	ngOnInit() {
		this.store.dispatch(new FetchEmailTemplatesList());
	}

	ngOnDestroy() {
		this.subscriptions.forEach((s) => s.unsubscribe());
		//this.store.dispatch(new Reset());
	}

	addNew() {
		this.dialog.open(AddEmailTemplateComponent)
	}

}
