import {NgModule, ModuleWithProviders} from '@angular/core';
import {NGXS_PLUGINS} from '@ngxs/store';
import {ReactiveFormsModule} from '@angular/forms';
import {NgxsFormErrorsPlugin} from "./ngxs-form-errors";
import {FormErrorsDirective} from "./form-errors.directive";

@NgModule({
	imports: [ReactiveFormsModule],
	declarations: [FormErrorsDirective],
	exports: [FormErrorsDirective]
})
export class NgxsFormErrorsPluginModule {
	static forRoot(): ModuleWithProviders<NgxsFormErrorsPluginModule> {
		return {
			ngModule: NgxsFormErrorsPluginModule,
			providers: [
				{
					provide: NGXS_PLUGINS,
					useClass: NgxsFormErrorsPlugin,
					multi: true
				}
			]
		};
	}
}
