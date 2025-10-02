import {Component, OnInit} from '@angular/core';
import {Observable} from "rxjs";
import {CarriersListModel} from "../carriers/carriers-list/carriers-list.ngxs";
import {Store} from "@ngxs/store";
import {MatDialog} from "@angular/material/dialog";
import {APITokensModel, APITokensState} from "./api-tokens.ngxs";
import {APITokensActions} from "./api-tokens.actions";
import FetchAPITokens = APITokensActions.FetchAPITokens;
import {AddNewApiTokenComponent} from "./dialogs/add-new-api-token/add-new-api-token.component";
import {ApiTokenConfirmDeleteComponent} from "./dialogs/api-token-confirm-delete/api-token-confirm-delete.component";
import SetConfirmDeleteToken = APITokensActions.SetConfirmDeleteToken;
import {Paths} from "../../app-routing.module";

@Component({
	selector: 'app-api-tokens',
	templateUrl: './api-tokens.component.html',
	styleUrls: ['./api-tokens.component.scss']
})
export class APITokensComponent implements OnInit {

	apiTokens$: Observable<APITokensModel>;
	docPath = Paths.API_DOCS;

	displayedColumns: string[] = [
		'name',
		'createdAt',
		'lastUsed',
		'actions',
	];

	constructor(
		private store: Store,
		private dialog: MatDialog,
	) {
		this.apiTokens$ = store.select(APITokensState.get);
	}

	ngOnInit(): void {
		this.store.dispatch(new FetchAPITokens());
	}

	addNew() {
		this.dialog.open(AddNewApiTokenComponent);
	}

	deleteToken(id: string) {
		this.store.dispatch(new SetConfirmDeleteToken(id));
		this.dialog.open(ApiTokenConfirmDeleteComponent);
	}

}
