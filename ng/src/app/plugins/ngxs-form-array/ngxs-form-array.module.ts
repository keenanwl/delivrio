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
export class NgxsFormArrayPluginModule {
	static forRoot(): ModuleWithProviders<NgxsFormArrayPluginModule> {
		return {
			ngModule: NgxsFormArrayPluginModule,
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
