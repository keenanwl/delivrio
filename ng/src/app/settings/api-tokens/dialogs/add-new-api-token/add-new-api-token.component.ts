import {Component, OnDestroy} from '@angular/core';
import {Store} from '@ngxs/store';
import {AppActions} from "../../../../app.actions";
import ShowGlobalSnackbar = AppActions.ShowGlobalSnackbar;
import {Observable} from "rxjs";
import {APITokensModel, APITokensState} from "../../api-tokens.ngxs";
import {APITokensActions} from "../../api-tokens.actions";
import CreateAPIToken = APITokensActions.CreateAPIToken;
import ClearNewToken = APITokensActions.ClearDialogs;
import {DialogRef} from "@angular/cdk/dialog";

@Component({
	selector: 'app-add-new-api-token',
	templateUrl: './add-new-api-token.component.html',
	styleUrls: ['./add-new-api-token.component.scss']
})
export class AddNewApiTokenComponent implements OnDestroy {

	apiTokens$: Observable<APITokensModel>;

	constructor(
		private store: Store,
		private dialog: DialogRef<any>,
	) {
		this.apiTokens$ = store.select(APITokensState.get);
	}

	ngOnDestroy() {
		this.store.dispatch(new ClearNewToken());
	}

	copySuccess() {
		this.store.dispatch(new ShowGlobalSnackbar(`Copied to clipboard`));
	}

	create(name: string) {
		this.store.dispatch(new CreateAPIToken({name: name}));
	}

	close() {
		this.dialog.close();
	}

}
