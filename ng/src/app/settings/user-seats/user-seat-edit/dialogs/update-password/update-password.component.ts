import { Component, EventEmitter, Output } from '@angular/core';
import { AsyncPipe, NgIf } from "@angular/common";
import { FormBuilder, FormGroup, Validators, AbstractControl, ValidationErrors, ReactiveFormsModule } from '@angular/forms';
import { MatButtonModule } from "@angular/material/button";
import { MatFormFieldModule } from "@angular/material/form-field";
import { MatIconModule } from "@angular/material/icon";
import { MatInputModule } from "@angular/material/input";
import { MatTooltipModule } from "@angular/material/tooltip";

@Component({
	selector: 'app-update-password',
	standalone: true,
	imports: [
		AsyncPipe,
		MatButtonModule,
		MatFormFieldModule,
		MatIconModule,
		MatInputModule,
		MatTooltipModule,
		NgIf,
		ReactiveFormsModule
	],
	templateUrl: './update-password.component.html',
	styleUrls: ['./update-password.component.scss']
})
export class UpdatePasswordComponent {
	@Output() out = new EventEmitter<string>();

	passwordForm: FormGroup;
	hidePassword = true;
	hideConfirmPassword = true;

	constructor(private fb: FormBuilder) {
		this.passwordForm = this.fb.group({
			password: ['', [Validators.required, Validators.minLength(8), this.passwordStrengthValidator]],
			confirmPassword: ['', Validators.required]
		}, { validators: this.passwordMatchValidator });
	}

	passwordStrengthValidator(control: AbstractControl): ValidationErrors | null {
		const value = control.value;
		if (!value) {
			return null;
		}

		const hasUpperCase = /[A-Z]+/.test(value);
		const hasLowerCase = /[a-z]+/.test(value);
		const hasNumeric = /[0-9]+/.test(value);

		const passwordValid = (hasUpperCase || hasLowerCase) && hasNumeric;

		return !passwordValid ? { passwordStrength: true } : null;
	}

	passwordMatchValidator(control: AbstractControl): ValidationErrors | null {
		const password = control.get('password');
		const confirmPassword = control.get('confirmPassword');

		if (password && confirmPassword && password.value !== confirmPassword.value) {
			confirmPassword.setErrors({ passwordMismatch: true });
			return { passwordMismatch: true };
		} else {
			confirmPassword?.setErrors(null);
			return null;
		}
	}

	onSubmit() {
		if (this.passwordForm.valid) {
			const password = this.passwordForm.get('password')?.value;
			if (password) {
				this.out.emit(password);
			}
		} else {
			this.passwordForm.markAllAsTouched();
		}
	}

	togglePasswordVisibility(field: 'password' | 'confirmPassword') {
		if (field === 'password') {
			this.hidePassword = !this.hidePassword;
		} else {
			this.hideConfirmPassword = !this.hideConfirmPassword;
		}
	}
}
