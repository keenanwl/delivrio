import {NgModule} from "@angular/core";
import {CommonModule, NgOptimizedImage} from "@angular/common";
import {NgxsModule} from "@ngxs/store";
import {NgxsFormPluginModule} from "@ngxs/form-plugin";
import {FormsModule, ReactiveFormsModule} from "@angular/forms";
import {NgxsFormErrorsPluginModule} from "../../plugins/ngxs-form-errors/ngxs-form-errors.module";
import {MaterialModule} from "../../modules/material.module";
import {DvoCardComponent} from "../../shared/dvo-card/dvo-card.component";
import {NgxsFormArrayPluginModule} from "../../plugins/ngxs-form-array/ngxs-form-array.module";
import {NotificationsRoutingModule} from "./notifications-routing.module";
import {NotificationsListComponent} from "./notifications-list/notifications-list.component";
import {NotificationsListState} from "./notifications-list/notifications-list.ngxs";
import { CreateNotificationComponent } from './notifications-list/dialogs/create-notification/create-notification.component';

@NgModule({
	imports: [
		NotificationsRoutingModule,
		NgxsModule.forFeature([
			NotificationsListState,
		]),
		MaterialModule,
		CommonModule,
		FormsModule,
		ReactiveFormsModule,
		DvoCardComponent,
		NgxsFormPluginModule,
		NgxsFormErrorsPluginModule,
		NgxsFormArrayPluginModule,
		NgOptimizedImage,
	],
	providers: [
	],
	declarations: [
		NotificationsListComponent,
  CreateNotificationComponent,
	]
})
export class NotificationsModule { }
