import {NgModule} from "@angular/core";
import {UserSeatsListComponent} from "./user-seats-list/user-seats-list.component";
import {UserSeatEditComponent} from "./user-seat-edit/user-seat-edit.component";
import {SeatGroupEditComponent} from "./seat-group-edit/seat-group-edit.component";
import {SeatGroupsListComponent} from "./seat-groups-list/seat-groups-list.component";
import {NgxsFormPluginModule} from "@ngxs/form-plugin";
import {CommonModule} from "@angular/common";
import {NgxsFormErrorsPluginModule} from "../../plugins/ngxs-form-errors/ngxs-form-errors.module";
import {FormsModule, ReactiveFormsModule} from "@angular/forms";
import {MaterialModule} from "../../modules/material.module";
import {NgxsModule} from "@ngxs/store";
import {DvoCardComponent} from "../../shared/dvo-card/dvo-card.component";
import {UserSeatsRoutingModule} from "./user-seats-routing.module";
import {UserSeatsState} from "./user-seat-edit/user-seat-edit.ngxs";
import {UserSeatsListState} from "./user-seats-list/user-seats-list.ngxs";
import {SeatGroupState} from "./seat-group-edit/seat-group-edit.ngxs";
import {SeatGroupsListState} from "./seat-groups-list/seat-groups-list.ngxs";

@NgModule({
	imports: [
		UserSeatsRoutingModule,
		NgxsModule.forFeature([
			UserSeatsListState,
			UserSeatsState,
			SeatGroupState,
			SeatGroupsListState,
		]),
		MaterialModule,
		CommonModule,
		FormsModule,
		ReactiveFormsModule,
		DvoCardComponent,
		NgxsFormPluginModule,
		NgxsFormErrorsPluginModule,
	],
	declarations: [
		UserSeatsListComponent,
		UserSeatEditComponent,
		SeatGroupEditComponent,
		SeatGroupsListComponent,
	]
})
export class UserSeatsModule { }
