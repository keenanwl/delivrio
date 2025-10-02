import {Component, OnDestroy} from '@angular/core';
import {MatDialogRef} from "@angular/material/dialog";
import { Store } from '@ngxs/store';
import {APITokensActions} from "../../api-tokens.actions";
import DeleteToken = APITokensActions.DeleteToken;
import ClearDialogs = APITokensActions.ClearDialogs;
import {Observable} from "rxjs";
import {APITokensModel, APITokensState} from "../../api-tokens.ngxs";
import FetchAPITokensResponse = APITokensActions.FetchAPITokensResponse;

@Component({
  selector: 'app-api-token-confirm-delete',
  templateUrl: './api-token-confirm-delete.component.html',
  styleUrls: ['./api-token-confirm-delete.component.scss']
})
export class ApiTokenConfirmDeleteComponent implements OnDestroy {

	apiTokens$: Observable<APITokensModel>;

	constructor(private ref: MatDialogRef<any>, private store: Store) {
		this.apiTokens$ = store.select(APITokensState.get);
	}

	ngOnDestroy() {
		this.store.dispatch(new ClearDialogs());
	}

	confirm() {
		const ID = this.store.selectSnapshot(APITokensState.get).confirmDeleteID;
		this.store.dispatch(new DeleteToken(ID));
		this.cancel();
	}

	cancel() {
		this.ref.close();
	}

	getTokenName(confirmDeleteID: string, allMyTokens: FetchAPITokensResponse[]): string {
		return allMyTokens.find((t) => t.id === confirmDeleteID)?.name || 'not found';
	}

	getTokenLastUsed(confirmDeleteID: string, allMyTokens: FetchAPITokensResponse[]): string | null {
		return allMyTokens.find((t) => t.id === confirmDeleteID)?.lastUsed;
	}

}
