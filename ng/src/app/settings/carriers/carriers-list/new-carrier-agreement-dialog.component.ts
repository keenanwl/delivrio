import {MatDialogRef} from "@angular/material/dialog";
import {Component} from "@angular/core";
import {Observable} from "rxjs";
import {Store} from "@ngxs/store";
import {CarriersListModel, CarriersListState} from "./carriers-list.ngxs";
import {CarriersListActions} from "./carriers-list.actions";
import FetchCarrierBrands = CarriersListActions.FetchCarrierBrands;
import CreateNewAgreement = CarriersListActions.CreateNewAgreement;

@Component({
	selector: 'new-carrier-agreement-dialog',
	styleUrls: ['new-carrier-agreement-dialog.component.scss'],
	templateUrl: 'new-carrier-agreement.component.html',
})
export class NewCarrierAgreementDialogComponent {

	carrierList$: Observable<CarriersListModel>;

	constructor(
		private store: Store,
		private dialogRef: MatDialogRef<NewCarrierAgreementDialogComponent>,
	) {
		this.carrierList$ = store.select(CarriersListState.get);
		this.store.dispatch([new FetchCarrierBrands()]);
	}

	selected(name: string, carrierBrandID: string) {
		this.store.dispatch(new CreateNewAgreement({name, carrierBrandID}));
		this.dialogRef.close();
	}
}
