import {ChangeDetectorRef, Directive, Input, OnDestroy, OnInit} from '@angular/core';
import {Observable, Subject, takeUntil} from "rxjs";
import {Actions, getValue, Store} from "@ngxs/store";
import {FormArray, FormGroupDirective} from "@angular/forms";
import {GraphQLError} from "graphql";
import {FormControl, FormGroup} from "@angular/forms";
import {debounceTime, distinctUntilChanged, filter} from "rxjs/operators";
import {UpdateForm, UpdateFormDirty, UpdateFormErrors, UpdateFormStatus, UpdateFormValue} from "@ngxs/form-plugin";

@Directive({
	selector: '[ngxsFormArray]'
})
export class FormErrorsDirective implements OnInit, OnDestroy {

	@Input('ngxsFormArray') path: string = null!;
	@Input('ngxsFormDebounce') debounce = 100;
	@Input('ngxsFormClearOnDestroy')
	set clearDestroy(val: boolean) {
		this._clearDestroy = val != null && `${val}` !== 'false';
	}

	get clearDestroy(): boolean {
		return this._clearDestroy;
	}

	_clearDestroy = false;

	private readonly _destroy$ = new Subject<void>();
	private _updating = false;

	constructor(
		private _actions$: Actions,
		private _store: Store,
		private _formGroupDirective: FormGroupDirective,
		private _cd: ChangeDetectorRef,
	) { }

	ngOnInit() {

		this.getStateStream(`${this.path}.model`).subscribe((data) => {
			if (!data) {
				return;
			}

			for (let controlsKey in data) {
				if (Array.isArray(data[controlsKey])) {
					let arr = this._formGroupDirective.form.controls[controlsKey] as FormArray;
					arr.controls = [];
					arr.controls.push(new FormControl());
				}
			}
			this._formGroupDirective.form.patchValue(data, {emitEvent: false});
		});

		this._formGroupDirective
			.valueChanges!.pipe(
			distinctUntilChanged((a, b) => JSON.stringify(a) === JSON.stringify(b)),
			this.debounceChange()
		).subscribe(() => {
			this.updateFormStateWithRawValue();
		});

		this.getStateStream(`${this.path}.errors`).subscribe((errors: readonly GraphQLError[] | null) => {

			this._formGroupDirective.form.setErrors(null);

			if (!!errors && Array.isArray(errors)) {
				errors.forEach((e) => {
					const path = e.path;
					if (!!path) {
						const pathWithoutMutation = path.slice(2, path.length);
						console.warn(pathWithoutMutation)
						const control = this._formGroupDirective.form.get(pathWithoutMutation);
						if (!!control) {
							control.setErrors(e.message as any);
							control.markAllAsTouched()
						}
					}
				});
			}
			this._cd.markForCheck();
		});
	}

	private getStateStream(path: string) {
		return this._store.select(state => getValue(state, path)).pipe(takeUntil(this._destroy$));
	}

	updateFormStateWithRawValue(withFormStatus?: boolean) {
		if (this._updating) return;

		const value = this._formGroupDirective.control.getRawValue();
console.warn("RAW VAL", value)
		const actions: any[] = [
			new UpdateFormValue({
				path: this.path,
				value
			}),
			new UpdateFormDirty({
				path: this.path,
				dirty: this._formGroupDirective.dirty
			}),
			new UpdateFormErrors({
				path: this.path,
				errors: this._formGroupDirective.errors
			})
		];

		if (withFormStatus) {
			actions.push(
				new UpdateFormStatus({
					path: this.path,
					status: this._formGroupDirective.status
				})
			);
		}

		this._updating = true;
		this._store.dispatch(actions).subscribe({
			error: () => (this._updating = false),
			complete: () => (this._updating = false)
		});
	}

	ngOnDestroy() {
		this._destroy$.next();
		this._destroy$.complete();

		if (this.clearDestroy) {
			this._store.dispatch(
				new UpdateForm({
					path: this.path,
					value: null,
					dirty: null,
					status: null,
					errors: null
				})
			);
		}
	}

	private debounceChange() {
		const skipDebounceTime =
			this._formGroupDirective.control.updateOn !== 'change' || this.debounce < 0;

		return skipDebounceTime
			? (change: Observable<any>) => change.pipe(takeUntil(this._destroy$))
			: (change: Observable<any>) =>
				change.pipe(debounceTime(this.debounce), takeUntil(this._destroy$));
	}

	private get form(): FormGroup {
		return this._formGroupDirective.form;
	}

}
