import {Injectable} from '@angular/core';
import {CanActivate} from '@angular/router';
import {Actions, ofActionDispatched, Store} from '@ngxs/store';
import {AppActions} from "../app.actions";
import AppChangeRoute = AppActions.AppChangeRoute;
import {AppState} from "../app.ngxs";

@Injectable()
export class AuthGuardRegistered implements CanActivate {
	constructor(private store: Store, private actions$: Actions) {}

	canActivate() {
		const state = this.store.selectSnapshot(AppState.get);

		if (!state.isRegistered) {
			this.store.dispatch([
				new AppChangeRoute({path: "", queryParams: {}}),
			]);
			return false;
		}

		return true;
	}

}
