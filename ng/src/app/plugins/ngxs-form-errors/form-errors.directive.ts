import {ChangeDetectorRef, Directive, Input, OnInit} from '@angular/core';
import {Subject, takeUntil} from "rxjs";
import {Actions, getValue, Store} from "@ngxs/store";
import {FormArray, FormGroupDirective} from "@angular/forms";
import {GraphQLError} from "graphql";
import {FormControl, FormGroup} from "@angular/forms";

@Directive({
	selector: '[ngxsFormErrors]'
})
export class FormErrorsDirective implements OnInit {

	@Input('ngxsFormErrors') path: string = null!;

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
			if (!!data && Array.isArray(data.rules)) {
				const control = this._formGroupDirective.form.get('rules');
				if (control instanceof FormArray) {
					if (control.controls.length === data.rules.length) {
						return;
					}
					control.clear();
					data.rules.forEach(() => {
						control.push(new FormGroup({
							id: new FormControl(''),
							name: new FormControl('')
						}));
					});
					this._cd.markForCheck();
				}
			}

		})

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

}
