import {NgModule} from "@angular/core";
import {FilterBarComponent} from "./filter-bar.component";
import {MaterialModule} from "../../modules/material.module";
import {FormsModule, ReactiveFormsModule} from "@angular/forms";
import {CommonModule} from "@angular/common";

@NgModule({
	imports: [
		MaterialModule,
		FormsModule,
		ReactiveFormsModule,
		CommonModule,
	],
	exports: [
		FilterBarComponent,
	],
	declarations: [
		FilterBarComponent,
	]
})
export class FilterBarModule { }
