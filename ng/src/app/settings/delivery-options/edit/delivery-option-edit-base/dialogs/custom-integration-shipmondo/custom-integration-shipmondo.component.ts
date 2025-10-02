import {Component, Input, output} from '@angular/core';
import {MatButton} from "@angular/material/button";
import {MatFormField, MatHint, MatLabel} from "@angular/material/form-field";
import {MatIcon} from "@angular/material/icon";
import {MatInput} from "@angular/material/input";
import {DialogRef} from "@angular/cdk/dialog";

@Component({
	selector: 'app-custom-integration-shipmondo',
	standalone: true,
	imports: [
		MatButton,
		MatFormField,
		MatHint,
		MatIcon,
		MatInput,
		MatLabel
	],
	templateUrl: './custom-integration-shipmondo.component.html',
	styleUrl: './custom-integration-shipmondo.component.scss'
})
export class CustomIntegrationShipmondoComponent {

	@Input()
	set initialVal(value: string) {
		this._initialVal = value;
		this.updateVal(value);
	}
	protected _initialVal = "";
	outputVal = output<string>();

	demoVal = ""

	constructor(private ref: DialogRef) {
	}

	updateVal(nextVal: string) {
		this.demoVal = nextVal.replace("{{.DropPointID}}", "99387");
	}

	save(nextVal: string) {
		this.outputVal.emit(nextVal);
		this.close();
	}

	close() {
		this.ref.close();
	}
}
