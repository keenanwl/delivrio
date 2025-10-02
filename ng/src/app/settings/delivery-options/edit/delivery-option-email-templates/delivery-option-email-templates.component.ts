import {Component, EventEmitter, Input, OnInit, Output} from '@angular/core';
import {FormControl, FormGroup} from "@angular/forms";
import {
	DeliveryOptionEditPostNordActions
} from "../delivery-option-edit-post-nord/delivery-option-edit-post-nord.actions";
import EmailTemplateResponse = DeliveryOptionEditPostNordActions.EmailTemplateResponse;

export interface SelectedEmailTemplates {
	clickCollectAtStore: string | undefined;
}

@Component({
	selector: 'app-delivery-option-email-templates',
	templateUrl: './delivery-option-email-templates.component.html',
	styleUrls: ['./delivery-option-email-templates.component.scss']
})
export class DeliveryOptionEmailTemplatesComponent implements OnInit {

	@Input() set selectedEmailTemplates(val: SelectedEmailTemplates) {
		this._selectedEmailTemplates = val;
		this.form.controls.clickCollectAtStore.setValue(val.clickCollectAtStore || '', {emitEvent: false});
	}
	get selectedEmailTemplates(): SelectedEmailTemplates {
		return this._selectedEmailTemplates;
	}
	_selectedEmailTemplates: SelectedEmailTemplates = {
		clickCollectAtStore: '',
	};

	@Input() set allEmailTemplates(val: EmailTemplateResponse[]) {
		this._allEmailTemplates = val;
	}
	get allEmailTemplates(): EmailTemplateResponse[] {
		return this._allEmailTemplates;
	}
	_allEmailTemplates: EmailTemplateResponse[] = [];

	@Output() emailsSelected = new EventEmitter<SelectedEmailTemplates>();

	form = new FormGroup({
		clickCollectAtStore: new FormControl('', {nonNullable: true}),
	});

	ngOnInit() {
		this.form.valueChanges
			.subscribe((v) => {
				this.emailsSelected.emit({
					clickCollectAtStore: v.clickCollectAtStore,
				})
			});
	}

}
