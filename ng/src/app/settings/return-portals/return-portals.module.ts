import {NgModule} from "@angular/core";
import {CommonModule} from "@angular/common";
import {NgxsModule} from "@ngxs/store";
import {NgxsFormPluginModule} from "@ngxs/form-plugin";
import {FormsModule, ReactiveFormsModule} from "@angular/forms";
import {NgxsFormErrorsPluginModule} from "../../plugins/ngxs-form-errors/ngxs-form-errors.module";
import {MaterialModule} from "../../modules/material.module";
import {DvoCardComponent} from "../../shared/dvo-card/dvo-card.component";
import {ReturnPortalsRoutingModule} from "./return-portals-routing.module";
import {ReturnPortalsListComponent} from './return-portals-list/return-portals-list.component';
import {ReturnPortalsListState} from "./return-portals-list/return-portals-list.ngxs";
import {AddNewReturnPortalComponent} from './return-portals-list/dialog/add-new-return-portal/add-new-return-portal.component';
import {ReturnPortalEditComponent} from './return-portal-edit/return-portal-edit.component';
import {ReturnPortalEditState} from "./return-portal-edit/return-portal-edit.ngxs";
import {NgxsFormArrayPluginModule} from "../../plugins/ngxs-form-array/ngxs-form-array.module";
import {DeliveryOptionSelectedPipe} from "./return-portal-edit/pipes/delivery-option-selected.pipe";
import {FilterEmailTemplatesPipe} from "./return-portal-edit/pipes/filter-email-templates.pipe";

@NgModule({
    imports: [
        ReturnPortalsRoutingModule,
        NgxsModule.forFeature([
            ReturnPortalsListState,
            ReturnPortalEditState,
        ]),
        MaterialModule,
        CommonModule,
        FormsModule,
        ReactiveFormsModule,
        DvoCardComponent,
        NgxsFormPluginModule,
        NgxsFormErrorsPluginModule,
        NgxsFormArrayPluginModule,
        DeliveryOptionSelectedPipe,
        FilterEmailTemplatesPipe,
    ],
	declarations: [
		ReturnPortalsListComponent,
		AddNewReturnPortalComponent,
		ReturnPortalEditComponent,
	],
})
export class ReturnPortalsModule { }
