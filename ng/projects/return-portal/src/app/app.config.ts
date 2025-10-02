import {ApplicationConfig, importProvidersFrom} from '@angular/core';
import { provideRouter } from '@angular/router';

import { routes } from './app.routes';
import {NgxsModule} from "@ngxs/store";
import {environment} from "../../../../src/environments/environment";
import {provideAnimations} from "@angular/platform-browser/animations";
import {NgxsLoggerPluginModule} from "@ngxs/logger-plugin";
import {MAT_FORM_FIELD_DEFAULT_OPTIONS} from "@angular/material/form-field";
import {
	ReturnPortalContainerModule
} from "../../../../src/app/settings/return-portal-viewer/return-portal-container/return-portal-container.module";

export const appConfig: ApplicationConfig = {
  providers: [
	  provideRouter(routes),
	  importProvidersFrom(ReturnPortalContainerModule),
	  importProvidersFrom(NgxsModule.forRoot([], {developmentMode: !environment.production})),
	  importProvidersFrom(NgxsLoggerPluginModule.forRoot({disabled: environment.production})),
	  provideAnimations(),
	  {provide: MAT_FORM_FIELD_DEFAULT_OPTIONS, useValue: {appearance: "outline"}},
  ]
};
