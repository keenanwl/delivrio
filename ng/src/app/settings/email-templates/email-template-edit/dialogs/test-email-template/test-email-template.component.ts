import { Component } from '@angular/core';
import {Store} from "@ngxs/store";
import {MatDialogRef} from "@angular/material/dialog";
import {EmailTemplateEditActions} from "../../email-template-edit.actions";
import SendTestEmail = EmailTemplateEditActions.SendTestEmail;

@Component({
	selector: 'app-test-email-template',
	templateUrl: './test-email-template.component.html',
	styleUrls: ['./test-email-template.component.scss']
})
export class TestEmailTemplateComponent {

	constructor(
		private store: Store,
		private ref: MatDialogRef<any>,
	) {}

	send(toEmail: string) {
		this.store.dispatch(new SendTestEmail(toEmail));
		this.close();
	}

	close() {
		this.ref.close();
	}

}
