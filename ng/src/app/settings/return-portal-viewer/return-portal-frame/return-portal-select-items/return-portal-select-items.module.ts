import {NgModule} from "@angular/core";
import {CommonModule, NgOptimizedImage} from "@angular/common";
import {ReturnPortalSelectItemsComponent} from "./return-portal-select-items.component";
import {FormsModule, ReactiveFormsModule} from "@angular/forms";
import {MaterialModule} from "../../../../modules/material.module";

@NgModule({
    imports: [
        MaterialModule,
        CommonModule,
        FormsModule,
        ReactiveFormsModule,
        NgOptimizedImage,
    ],
	exports: [
		ReturnPortalSelectItemsComponent,
	],
	declarations: [
		ReturnPortalSelectItemsComponent,
	]
})
export class ReturnPortalSelectItemsModule { }
