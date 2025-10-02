import {BrowserModule, DomSanitizer} from '@angular/platform-browser';
import {
	CUSTOM_ELEMENTS_SCHEMA,
	DEFAULT_CURRENCY_CODE,
	LOCALE_ID,
	NgModule,
	provideExperimentalZonelessChangeDetection
} from '@angular/core';

import {AppRoutingModule} from './app-routing.module';
import {AppComponent} from './app.component';
import {LoginComponent} from './login/login.component';
import {LoginService} from './login/login.service';
import {MaterialModule} from "./modules/material.module";
import {BrowserAnimationsModule} from "@angular/platform-browser/animations";
import {RequestPasswordResetComponent} from './login/request-password-reset/request-password-reset.component';
import {PasswordResetComponent} from './login/password-reset/password-reset.component';
import {RequestPasswordResetService} from "./login/request-password-reset/request-password-reset.service";
import {PasswordResetService} from "./login/password-reset/password-reset.service";
import {LoginState} from "./login/login.ngxs";
import {NgxsModule} from "@ngxs/store";
import {environment} from "../environments/environment";
import {NgxsLoggerPluginModule} from "@ngxs/logger-plugin";
import { HTTP_INTERCEPTORS, provideHttpClient, withInterceptorsFromDi, withJsonpSupport } from "@angular/common/http";
import {JwtInterceptor} from "./interceptors/auth";
import {AppState} from "./app.ngxs";
import {CommonModule, DatePipe} from "@angular/common";
import {AppService} from "./app.service";
import {MAT_FORM_FIELD_DEFAULT_OPTIONS} from "@angular/material/form-field";
import {MenuItemsComponent} from './shared/menu-items/menu-items.component';
import {DialogComponent} from './shared/dialog/dialog.component';
import {FormsModule, ReactiveFormsModule} from "@angular/forms";
import {LoggedInWrapperComponent} from './logged-in-wrapper/logged-in-wrapper.component';
import {ToolbarComponent} from './shared/toolbar/toolbar.component';
import {OrdersService} from "./orders/orders.service";
import {GraphQLModule} from './modules/graphql.module';
import {DvoCardComponent} from './shared/dvo-card/dvo-card.component';
import {NgxsFormPluginModule} from "@ngxs/form-plugin";
import {SettingsComponent} from './settings/settings.component';
import {NgxsFormErrorsPluginModule} from "./plugins/ngxs-form-errors/ngxs-form-errors.module";
import {UserSeatsListState} from "./settings/user-seats/user-seats-list/user-seats-list.ngxs";
import {UserSeatsState} from "./settings/user-seats/user-seat-edit/user-seat-edit.ngxs";
import {SeatGroupsListState} from "./settings/user-seats/seat-groups-list/seat-groups-list.ngxs";
import {SeatGroupState} from "./settings/user-seats/seat-group-edit/seat-group-edit.ngxs";
import {DialogSelectPrinterComponent} from './logged-in-wrapper/dialog-select-printer/dialog-select-printer.component';
import {DialogViewPrintJobsComponent} from './logged-in-wrapper/dialog-view-print-jobs/dialog-view-print-jobs.component';
import {RelativeTimePipe} from './pipes/relative-time.pipe';
import {ActivatedRouteSnapshot, DetachedRouteHandle, RouteReuseStrategy} from "@angular/router";
import {MAT_ICON_DEFAULT_OPTIONS, MatIconRegistry} from "@angular/material/icon";
import {DateAdapter, MAT_DATE_FORMATS, MAT_DATE_LOCALE} from "@angular/material/core";
import {
	LuxonDateAdapter,
	MAT_LUXON_DATE_ADAPTER_OPTIONS,
	MAT_LUXON_DATE_FORMATS
} from "@angular/material-luxon-adapter";
import {SelectPackagingState} from "./shared/select-packaging/select-packaging.ngxs";
import {FilterStatusPipe} from "./logged-in-wrapper/filter-status-pipe";
import {PrintJobStatusColorPipe} from "./shared/printjob-status-color-pipe.pipe";
import {OrderStatusColorPipePipe} from "./shared/order-status-color-pipe.pipe";

export class NoopRouteReuseStrategy implements RouteReuseStrategy {
	shouldDetach(route: ActivatedRouteSnapshot): boolean {
		return false;
	}

	store(route: ActivatedRouteSnapshot, handle: DetachedRouteHandle | null): void {
	}

	shouldAttach(route: ActivatedRouteSnapshot): boolean {
		return false;
	}

	retrieve(route: ActivatedRouteSnapshot): DetachedRouteHandle | null {
		return null;
	}

	shouldReuseRoute(future: ActivatedRouteSnapshot, curr: ActivatedRouteSnapshot): boolean {
		// Implement your logic
		return false;
	}
}

@NgModule({
	declarations: [
        AppComponent,
        LoginComponent,
        RequestPasswordResetComponent,
        PasswordResetComponent,
        MenuItemsComponent,
        DialogComponent,
        LoggedInWrapperComponent,
        ToolbarComponent,
        SettingsComponent,
        DialogSelectPrinterComponent,
        DialogViewPrintJobsComponent,
    ],
    bootstrap: [AppComponent],
    schemas: [CUSTOM_ELEMENTS_SCHEMA],
	imports: [FilterStatusPipe,
        PrintJobStatusColorPipe,
        DvoCardComponent,
        CommonModule,
        BrowserModule,
        AppRoutingModule,
        MaterialModule,
        ReactiveFormsModule,
        BrowserAnimationsModule,
        NgxsFormErrorsPluginModule.forRoot(),
        NgxsFormPluginModule.forRoot(),
        NgxsModule.forRoot([
            AppState,
            LoginState,
            SelectPackagingState,
            UserSeatsListState,
            UserSeatsState,
            SeatGroupsListState,
            SeatGroupState,
        ], { developmentMode: !environment.production }),
        NgxsLoggerPluginModule.forRoot({ disabled: environment.production }),
        GraphQLModule,
        RelativeTimePipe,
        FormsModule,
        OrderStatusColorPipePipe
	],
	providers: [
		provideExperimentalZonelessChangeDetection(),
        AppService,
        OrdersService,
        LoginService,
        RequestPasswordResetService,
        PasswordResetService,
        { provide: HTTP_INTERCEPTORS, useClass: JwtInterceptor, multi: true },
        { provide: LOCALE_ID, useValue: 'en-GB' },
        { provide: DEFAULT_CURRENCY_CODE, useValue: 'DKK' },
        {
            provide: DateAdapter,
            useClass: LuxonDateAdapter,
            deps: [MAT_DATE_LOCALE, MAT_LUXON_DATE_ADAPTER_OPTIONS],
        },
        { provide: MAT_DATE_FORMATS, useValue: MAT_LUXON_DATE_FORMATS },
        { provide: MAT_DATE_LOCALE, useValue: 'en-US' },
        { provide: MAT_FORM_FIELD_DEFAULT_OPTIONS, useValue: { appearance: "outline" } },
        DatePipe,
        // Don't reuse routes to allow reloading on clicking current route
        { provide: RouteReuseStrategy, useClass: NoopRouteReuseStrategy },
        {
            provide: MAT_ICON_DEFAULT_OPTIONS,
            useFactory: (iconRegistry: MatIconRegistry, sanitizer: DomSanitizer) => {
                iconRegistry.setDefaultFontSetClass('material-symbols-outlined');
            },
            deps: [MatIconRegistry, DomSanitizer],
        },
        provideHttpClient(withInterceptorsFromDi(), withJsonpSupport()),
    ]
})
export class AppModule { }

