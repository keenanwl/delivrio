import {Component} from '@angular/core';
import {MatDialogRef} from "@angular/material/dialog";
import {Store} from "@ngxs/store";
import {EmailTemplatesListActions} from "../../email-templates-list.actions";
import Create = EmailTemplatesListActions.Create;
import {EmailTemplateMergeType} from "../../../../../../generated/graphql";

@Component({
  selector: 'app-add-email-template',
  templateUrl: './add-email-template.component.html',
  styleUrls: ['./add-email-template.component.scss']
})
export class AddEmailTemplateComponent {

	constructor(
		private store: Store,
		private ref: MatDialogRef<any>,
	) {}

	create(name: string, merge: EmailTemplateMergeType) {
		this.store.dispatch(new Create({name, merge}));
		this.close();
	}

	close() {
		this.ref.close();
	}

    protected readonly mergeType = EmailTemplateMergeType;
}
