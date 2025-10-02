import {Component, OnInit} from '@angular/core';
import {Observable} from "rxjs";
import {OrderEditModel, OrderEditState} from "../../order-edit.ngxs";
import {Store} from "@ngxs/store";
import {MatDialogRef} from "@angular/material/dialog";
import {OrderEditActions} from "../../order-edit.actions";
import AvailableClickCollectResponse = OrderEditActions.AvailableClickCollectResponse;
import FetchAvailableClickCollectLocations = OrderEditActions.FetchAvailableClickCollectLocations;
import SetClickCollectLocation = OrderEditActions.SetClickCollectLocation;

@Component({
	selector: 'app-edit-cc-location',
	templateUrl: './edit-cc-location.component.html',
	styleUrls: ['./edit-cc-location.component.scss']
})
export class EditCcLocationComponent implements OnInit {

	order$: Observable<OrderEditModel>;
	constructor(private store: Store, private ref: MatDialogRef<any>) {
		this.order$ = store.select(OrderEditState.get);
	}

	ngOnInit() {
		this.store.dispatch(new FetchAvailableClickCollectLocations());
	}

	selectCCLocation(opt: AvailableClickCollectResponse) {
		this.store.dispatch(new SetClickCollectLocation(opt));
		this.close();
	}

	close() {
		this.ref.close();
	}
}
