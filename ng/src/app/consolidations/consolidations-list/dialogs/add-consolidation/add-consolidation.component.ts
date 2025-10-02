import {Component} from '@angular/core';
import {DialogRef} from "@angular/cdk/dialog";
import {Store} from "@ngxs/store";
import {ConsolidationsListActions} from "../../consolidations-list.actions";
import AddConsolidation = ConsolidationsListActions.AddConsolidation;

@Component({
	selector: 'app-add-consolidation',
	templateUrl: './add-consolidation.component.html',
	styleUrl: './add-consolidation.component.scss'
})
export class AddConsolidationComponent {

	constructor(
		private ref: DialogRef,
		private store: Store,
	) {}

	create(publicID: string, description: string) {
		this.store.dispatch(new AddConsolidation({publicID, description}));
		this.close();
	}

	close() {
		this.ref.close();
	}

}
