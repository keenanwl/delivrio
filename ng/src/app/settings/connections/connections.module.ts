import {NgModule} from "@angular/core";
import {CommonModule} from "@angular/common";
import {ConnectionEditComponent} from "./connection-edit/connection-edit.component";
import {NgxsModule} from "@ngxs/store";
import {NgxsFormPluginModule} from "@ngxs/form-plugin";
import {ConnectionsListState} from "./connections-list/connections-list.ngxs";
import {ConnectionEditState} from "./connection-edit/connection-edit.ngxs";
import {FormsModule, ReactiveFormsModule} from "@angular/forms";
import {ConnectionsListComponent} from "./connections-list/connections-list.component";
import {NgxsFormErrorsPluginModule} from "../../plugins/ngxs-form-errors/ngxs-form-errors.module";
import {MaterialModule} from "../../modules/material.module";
import {DvoCardComponent} from "../../shared/dvo-card/dvo-card.component";
import {ConnectionsRoutingModule} from "./connections-routing.module";
import {LocationsSelectorComponent} from './connection-edit/locations-selector/locations-selector.component';
import {FilterTagsPipe} from "./connection-edit/locations-selector/pipes/filter-tags.pipe";
import {ToggleContainerComponent} from "../../shared/toggle-container/toggle-container.component";

@NgModule({
    imports: [
        ConnectionsRoutingModule,
        NgxsModule.forFeature([
            ConnectionsListState,
            ConnectionEditState,
        ]),
        MaterialModule,
        CommonModule,
        FormsModule,
        ReactiveFormsModule,
        DvoCardComponent,
        NgxsFormPluginModule,
        NgxsFormErrorsPluginModule,
        ToggleContainerComponent,
    ],
	declarations: [
		ConnectionsListComponent,
		ConnectionEditComponent,
        LocationsSelectorComponent,
		FilterTagsPipe,
	]
})
export class ConnectionsModule { }
