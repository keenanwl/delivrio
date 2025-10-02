import {Component} from '@angular/core';
import {Observable} from "rxjs";
import {ConsolidationEditModel, ConsolidationEditState} from "../../consolidation-edit.ngxs";
import {ActivatedRoute} from "@angular/router";
import {Actions, Store} from "@ngxs/store";
import {MatDialog} from "@angular/material/dialog";
import {ConsolidationEditActions} from "../../consolidation-edit.actions";
import IncrementLabelViewerOffset = ConsolidationEditActions.IncrementLabelViewerOffset;
import DecrementLabelViewerOffset = ConsolidationEditActions.DecrementLabelViewerOffset;

@Component({
	selector: 'app-create-shipment',
	templateUrl: './create-shipment.component.html',
	styleUrl: './create-shipment.component.scss'
})
export class CreateShipmentComponent {
	consolidationEdit$: Observable<ConsolidationEditModel>;

	constructor(
		private route: ActivatedRoute,
		private store: Store,
		private dialog: MatDialog,
		private actions$: Actions,
	) {
		this.consolidationEdit$ = store.select(ConsolidationEditState.get);
	}

	increment() {
		this.store.dispatch(new IncrementLabelViewerOffset());
	}

	decrement() {
		this.store.dispatch(new DecrementLabelViewerOffset());
	}
}
