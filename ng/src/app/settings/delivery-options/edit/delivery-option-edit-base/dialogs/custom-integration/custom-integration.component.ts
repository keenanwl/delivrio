import {Component, Input, output} from '@angular/core';
import {FormsModule} from "@angular/forms";
import {MatError, MatFormField, MatHint, MatLabel} from "@angular/material/form-field";
import {MatInput} from "@angular/material/input";
import {NgIf} from "@angular/common";
import {MatButton} from "@angular/material/button";
import {MatIcon} from "@angular/material/icon";
import {DialogRef} from "@angular/cdk/dialog";

@Component({
	selector: 'app-custom-integration',
	standalone: true,
	imports: [
		FormsModule,
		MatError,
		MatFormField,
		MatInput,
		MatLabel,
		MatHint,
		NgIf,
		MatButton,
		MatIcon,
	],
	templateUrl: './custom-integration.component.html',
	styleUrl: './custom-integration.component.scss'
})
export class CustomIntegrationComponent {

	@Input()
	set initialVal(value: number) {
		this._initialVal = value;
		this.updateVal(value.toString(10));
	}
	protected _initialVal = 1;

	outputVal = output<number>();

	webshipperCode = {
		"shipping_rate_id": 1,
		"drop_point": {
			"drop_point_id": "99387",
			"name": "Shell Frederiksværk",
			"address_1": "Hillerødvej 34",
			"zip": "3300",
			"city": "Frederiksværk",
			"country_code": "DK",
			"distance": 0.5154479040557662
		}
	};

	constructor(private ref: DialogRef) {
	}

	updateVal(nextVal: string) {
		this.webshipperCode.shipping_rate_id = parseInt(nextVal) || 1;
	}

	save(nextVal: string) {
		const val = parseInt(nextVal) || 1;
		this.outputVal.emit(val);
		this.close();
	}

	close() {
		this.ref.close();
	}

	protected readonly JSON = JSON;
}
