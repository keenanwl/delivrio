import {Component, EventEmitter, Input, OnDestroy, OnInit, Output} from '@angular/core';
import {Actions, ofActionSuccessful, Store} from "@ngxs/store";
import {OrderViewActions} from "../../order-view.actions";
import {MatDialogRef} from "@angular/material/dialog";
import {Subscription} from "rxjs";
import {OrdersActions} from "../../../orders.actions";
import {CreateOrderInput, UpdateOrderInput} from "../../../../../generated/graphql";
import SaveOrderSuccess = OrderViewActions.SaveOrderSuccess;
import ConnectionsResponse = OrdersActions.ConnectionsResponse;
import LocationResponse = OrdersActions.LocationResponse;

@Component({
	selector: 'app-create-edit-order-dialog',
	templateUrl: './create-edit-order-dialog.component.html',
	styleUrls: ['./create-edit-order-dialog.component.scss']
})
export class CreateEditOrderDialogComponent implements OnInit, OnDestroy {

	@Input() connections: ConnectionsResponse[] = [];
	@Input() senderLocations: LocationResponse[] = [];
	@Input() orderPublicID = "";
	@Input() commentInternal = "";
	@Input() commentExternal = "";
	@Input() connectionID = "";

	@Output() saveEmit: EventEmitter<{input: CreateOrderInput}> = new EventEmitter();

	isEdit = true;

	subscriptions$: Subscription[] = [];

	constructor(
		private actions$: Actions,
		private store: Store,
		private dialogRef: MatDialogRef<any>,
	) {

	}

	ngOnInit() {
		this.subscriptions$.push(this.actions$.pipe(ofActionSuccessful(SaveOrderSuccess))
			.subscribe(() => {
				this.close();
			}));
	}

	save(orderPublicID: string, connectionID: string, commentInternal: string, commentExternal: string) {
		this.saveEmit.emit({
			input: {
				connectionID: connectionID,
				orderPublicID: orderPublicID,
				commentInternal,
				commentExternal,
			},
		});
	}

	close() {
		this.dialogRef.close();
	}

	ngOnDestroy(): void {
		this.subscriptions$.forEach((s) => s.unsubscribe());
	}

}
