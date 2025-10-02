import {MatDialogRef} from "@angular/material/dialog";
import {Component, OnDestroy} from "@angular/core";
import {Store} from "@ngxs/store";
import CreateDeliveryRule = DeliveryOptionEditRulesActions.CreateDeliveryRule;
import {DeliveryOptionEditRulesActions} from "../delivery-option-edit-rules.actions";

@Component({
	selector: 'app-delivery-options-rule-dialog',
	styleUrls: ['new-delivery-options-rule-dialog.component.scss'],
	templateUrl: 'new-delivery-options-rule-dialog.component.html',
})
export class NewDeliveryOptionsRuleDialogComponent {

	constructor(
		private store: Store,
		private dialogRef: MatDialogRef<NewDeliveryOptionsRuleDialogComponent>,
	) {}

	selected(name: string) {
		this.store.dispatch([new CreateDeliveryRule(name)]);
		this.dialogRef.close();
	}

}
