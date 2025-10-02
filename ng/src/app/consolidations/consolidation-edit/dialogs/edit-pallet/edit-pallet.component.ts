import {Component} from '@angular/core';
import {Observable} from "rxjs";
import {ConsolidationEditModel, ConsolidationEditState} from "../../consolidation-edit.ngxs";
import {Store} from "@ngxs/store";
import {DialogRef} from "@angular/cdk/dialog";
import {ConsolidationEditActions} from "../../consolidation-edit.actions";
import UpdatePallet = ConsolidationEditActions.UpdatePallet;
import {EmailTemplateMergeType} from "../../../../../generated/graphql";
import {Paths} from "../../../../app-routing.module";

@Component({
	selector: 'app-edit-pallet',
	templateUrl: './edit-pallet.component.html',
	styleUrl: './edit-pallet.component.scss'
})
export class EditPalletComponent {

	consolidationEdit$: Observable<ConsolidationEditModel>;
	paths = Paths;

	constructor(
		private store: Store,
		private ref: DialogRef,
	) {
		this.consolidationEdit$ = store.select(ConsolidationEditState.get);
	}

	addPalletUpdate(publicID: string, description: string, packagingID: string) {
		this.store.dispatch(new UpdatePallet({publicID, description, packagingID}));
		this.close();
	}

	close() {
		this.ref.close();
	}

	protected readonly mergeType = EmailTemplateMergeType;
}
