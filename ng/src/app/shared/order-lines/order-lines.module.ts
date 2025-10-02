import {NgModule} from "@angular/core";
import {OrderLinesComponent} from "./order-lines.component";
import {FormsModule, ReactiveFormsModule} from "@angular/forms";
import {CommonModule} from "@angular/common";
import {MaterialModule} from "../../modules/material.module";
import {DvoCardComponent} from "../dvo-card/dvo-card.component";
import {DragDropModule} from "@angular/cdk/drag-drop";
import {RouterLink} from "@angular/router";
import {TotalWeightPipe} from "./pipes/total-weight.pipe";

@NgModule({
    imports: [
        MaterialModule,
        CommonModule,
        FormsModule,
        ReactiveFormsModule,
        DvoCardComponent,
        DragDropModule,
        RouterLink,
        TotalWeightPipe,
    ],
	exports: [
		OrderLinesComponent
	],
	declarations: [
		OrderLinesComponent
	]
})
export class OrderLinesModule { }
