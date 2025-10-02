import {Injectable} from '@angular/core';
import {
	getActionTypeFromInstance,
	NgxsNextPluginFn,
	NgxsPlugin,
	setValue
} from '@ngxs/store';
import {FormErrorsActions} from "./form-errors.actions";
import SetFormErrors = FormErrorsActions.SetFormErrors;
import {UpdateFormValue} from "@ngxs/form-plugin";

@Injectable()
export class NgxsFormErrorsPlugin implements NgxsPlugin {
	handle(state: any, event: any, next: NgxsNextPluginFn) {
		const type = getActionTypeFromInstance(event);

		let nextState = state;

		if (type === UpdateFormValue.type) {

		}

		if (type === SetFormErrors.type) {
			nextState = setValue(nextState, `${event.payload.formPath}.errors`, [
				...event.payload.errors
			]);
		}

		return next(nextState, event);
	}
}
