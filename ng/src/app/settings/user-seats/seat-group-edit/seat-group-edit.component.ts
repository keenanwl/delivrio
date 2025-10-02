import {ChangeDetectorRef, Component, OnDestroy, OnInit} from '@angular/core';
import {Actions, ofActionCompleted, Store} from "@ngxs/store";
import {ActivatedRoute} from "@angular/router";
import {SeatGroupActions} from "./seat-group-edit.actions";
import SetSeatGroupID = SeatGroupActions.SetSeatGroupID;
import FetchSeatGroup = SeatGroupActions.FetchSeatGroup;
import {Observable, Subscription} from "rxjs";
import {FormArray, FormControl, FormGroup} from "@angular/forms";
import {SeatGroupModel, SeatGroupState} from './seat-group-edit.ngxs';
import {CreateSeatGroupAccessRightInput, SeatGroupAccessRightLevel} from 'src/generated/graphql';
import SetAccessRights = SeatGroupActions.SetAccessRights;
import SaveFormEdit = SeatGroupActions.SaveFormEdit;
import SaveFormNew = SeatGroupActions.SaveFormNew;
import Clear = SeatGroupActions.Clear;

@Component({
	selector: 'app-seat-group-edit',
	templateUrl: './seat-group-edit.component.html',
	styleUrls: ['./seat-group-edit.component.scss']
})
export class SeatGroupEditComponent implements OnInit, OnDestroy {

	seatGroup$: Observable<SeatGroupModel>;
	accessRights$: Observable<Map<string, string>>;

	editForm = new FormGroup({
		id: new FormControl<string>('', {nonNullable: true}),
		name: new FormControl<string>('', {nonNullable: true}),
	});

	accessRightEditForm = new FormGroup({
		accessRights: new FormArray<FormGroup<{
			id: FormControl<string>;
			label: FormControl<string>,
			level: FormControl<SeatGroupAccessRightLevel>
		}>>([])
	});

	subscriptions$: Subscription[] = [];

	constructor(private store: Store,
	            private route: ActivatedRoute,
				private cd: ChangeDetectorRef,
	            private actions$: Actions) {
		this.seatGroup$ = store.select(SeatGroupState.get);
		this.accessRights$ = store.select(SeatGroupState.accessRights);
	}

	ngOnInit(): void {
		this.subscriptions$.push(this.route.queryParams
			.subscribe((params) => {
				this.store.dispatch([
					new SetSeatGroupID(!!params.id ? params.id : ''),
					new FetchSeatGroup(),
				]);
			}));
		this.subscriptions$.push(this.actions$.pipe(ofActionCompleted(SetAccessRights))
			.subscribe(() => {
				const state = this.store.selectSnapshot(SeatGroupState.get);
				const arState = this.store.selectSnapshot(SeatGroupState.accessRights);

				state.accessRights.forEach((c) => {

					const savedValue = arState.get(`${c?.internalID}`);

					this.accessRightEditForm.controls.accessRights.push(
						new FormGroup({
							id: new FormControl(`${c?.id}`, {nonNullable: true}),
							label: new FormControl(`${c?.label}`, {nonNullable: true}),
							level: new FormControl<SeatGroupAccessRightLevel>(
								!!savedValue ? savedValue : SeatGroupAccessRightLevel.None, {nonNullable: true}),
						})
					);
				});
				this.cd.detectChanges();
			}));
	}

	save() {

		const state = this.store.selectSnapshot(SeatGroupState.get);
		const form = this.accessRightEditForm.getRawValue();

		const ar: CreateSeatGroupAccessRightInput[] = [];
		form.accessRights.forEach((a) => ar.push({accessRightID: a.id, level: a.level, seatGroupID: ''}));

		if (state.seatGroupID.length > 0) {
			this.store.dispatch(new SaveFormEdit({
				input: {
					name: this.editForm.getRawValue().name,
				},
				accessRights: ar,
				id: state.seatGroupID,
			}));
		} else {
			this.store.dispatch(new SaveFormNew({
				input: {
					name: this.editForm.getRawValue().name,
				},
				accessRights: ar,
			}));
		}

	}

	ngOnDestroy(): void {
		this.store.dispatch(new Clear());
		this.subscriptions$.forEach(s => s.unsubscribe());
	}

}
