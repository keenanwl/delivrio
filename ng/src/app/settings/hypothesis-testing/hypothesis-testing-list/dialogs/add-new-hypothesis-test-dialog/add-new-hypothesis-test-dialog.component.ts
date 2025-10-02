import {Component} from '@angular/core';
import {Observable, Subscription} from "rxjs";
import {HypothesisTestingListModel, HypothesisTestingListState} from "../../hypothesis-testing-list.ngxs";
import {Store} from "@ngxs/store";
import {MatDialogRef} from "@angular/material/dialog";
import {HypothesisTestingListActions} from "../../hypothesis-testing-list.actions";
import CreateNewTest = HypothesisTestingListActions.CreateNewTest;

@Component({
	selector: 'app-add-new-hypothesis-test-dialog',
	templateUrl: './add-new-hypothesis-test-dialog.component.html',
	styleUrls: ['./add-new-hypothesis-test-dialog.component.scss']
})
export class AddNewHypothesisTestDialogComponent {
	state$: Observable<HypothesisTestingListModel>;

	subscriptions: Subscription[] = [];

	constructor(
		private store: Store,
		private ref: MatDialogRef<any>,
	) {
		this.state$ = store.select(HypothesisTestingListState.get);
	}

	create(name: string, connectionID: string) {
		this.store.dispatch(new CreateNewTest({name, connectionID}));
		this.ref.close();
	}
}
