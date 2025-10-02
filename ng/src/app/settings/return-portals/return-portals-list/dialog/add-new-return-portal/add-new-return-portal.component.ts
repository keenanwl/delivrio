import {Component} from '@angular/core';
import {MatDialogRef} from "@angular/material/dialog";
import {Store} from "@ngxs/store";
import {ReturnPortalsListActions} from "../../return-portals-list.actions";
import Create = ReturnPortalsListActions.Create;
import {Observable} from "rxjs";
import {ReturnPortalsListModel, ReturnPortalsListState} from "../../return-portals-list.ngxs";

@Component({
  selector: 'app-add-new-return-portal',
  templateUrl: './add-new-return-portal.component.html',
  styleUrls: ['./add-new-return-portal.component.scss']
})
export class AddNewReturnPortalComponent {

	portals$: Observable<ReturnPortalsListModel>;
	constructor(
		private ref: MatDialogRef<any>,
		private store: Store,
	) {
		this.portals$ = store.select(ReturnPortalsListState.get);
	}

	create(name: string, connection: string) {
		this.store.dispatch(new Create({name, connection}));
		this.cancel();
	}

	cancel() {
		this.ref.close();
	}

}
