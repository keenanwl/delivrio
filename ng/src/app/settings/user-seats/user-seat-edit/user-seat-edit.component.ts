import {Component, OnDestroy, OnInit} from '@angular/core';
import {Observable, Subscription} from "rxjs";
import {FormControl, FormGroup} from "@angular/forms";
import {Actions, Store} from "@ngxs/store";
import {ActivatedRoute} from "@angular/router";
import {UserSeatsModel, UserSeatsState} from "./user-seat-edit.ngxs";
import {UserSeatsActions} from "./user-seat-edit.actions";
import FetchUserSeat = UserSeatsActions.FetchUserSeat;
import SetUserSeatID = UserSeatsActions.SetUserSeatID;
import FetchSeatGroupsResponse = UserSeatsActions.FetchSeatGroupsResponse;
import SaveFormNew = UserSeatsActions.SaveFormNew;
import FetchUserSeatResponse = UserSeatsActions.FetchUserSeatResponse;
import SaveFormUpdate = UserSeatsActions.SaveFormUpdate;
import Clear = UserSeatsActions.Clear;
import {MatDialog} from "@angular/material/dialog";
import {UpdatePasswordComponent} from "./dialogs/update-password/update-password.component";
import UpdatePassword = UserSeatsActions.UpdatePassword;

@Component({
	selector: 'app-user-seat-edit',
	templateUrl: './user-seat-edit.component.html',
	styleUrls: ['./user-seat-edit.component.scss']
})
export class UserSeatEditComponent implements OnInit, OnDestroy {

	userSeats$: Observable<UserSeatsModel>;

	editForm = new FormGroup({
		name: new FormControl<string>('', {nonNullable: true}),
		surname: new FormControl<string>('', {nonNullable: true}),
		email: new FormControl<string>('', {nonNullable: true}),
		seatGroup: new FormControl<string | null>(null),
	});

	subscriptions$: Subscription[] = [];

	groupComparisonFunction = (option: FetchSeatGroupsResponse, value: FetchSeatGroupsResponse): boolean => {
		return option?.id === value?.id;
	}

	constructor(private store: Store,
	            private route: ActivatedRoute,
				private dialog: MatDialog,
	            private actions$: Actions) {
		this.userSeats$ = store.select(UserSeatsState.get);
	}

	ngOnInit(): void {
		this.subscriptions$.push(this.route.queryParams
			.subscribe((params) => {
				this.store.dispatch([
					new SetUserSeatID(!!params.id ? params.id : ''),
					new FetchUserSeat(),
				]);
			}));
	}

	save() {
		const id = this.store.selectSnapshot(UserSeatsState.get).userSeatID;
		const values = this.editForm.getRawValue();
		const output = Object.assign({}, values, {
			seatGroupID: values.seatGroup,
			seatGroup: undefined,
		});
		if (id.length === 0) {
			this.store.dispatch(new SaveFormNew({input: output}));
		} else {
			this.store.dispatch(new SaveFormUpdate({id, input: output}));
		}

	}

	updatePassword() {
		const ref = this.dialog.open(UpdatePasswordComponent);
		ref.componentInstance.out.subscribe((out) => {
			this.store.dispatch(new UpdatePassword({userID: this.store.selectSnapshot(UserSeatsState.get).userSeatID, password: out}));
			ref.close();
		});
	}

	ngOnDestroy(): void {
		this.store.dispatch(new Clear());
		this.subscriptions$.forEach(s => s.unsubscribe());
	}

}
