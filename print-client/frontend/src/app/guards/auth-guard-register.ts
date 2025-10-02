import {Injectable} from '@angular/core';
import {CanActivate} from '@angular/router';
import {Actions, ofActionDispatched, Store} from '@ngxs/store';
import {AppActions} from "../app.actions";
import AppChangeRoute = AppActions.AppChangeRoute;
import {AppState} from "../app.ngxs";
import {map} from "rxjs/operators";

@Injectable()
export class AuthGuardRegister implements CanActivate {
	constructor(private store: Store, private actions$: Actions) {}

	canActivate() {
		const state = this.store.selectSnapshot(AppState.get);

		if (state.isRegistered) {
			this.store.dispatch([
				new AppChangeRoute({path: "/dashboard", queryParams: {}}),
			]);
			return false;
		}

		return true;
	}

}
