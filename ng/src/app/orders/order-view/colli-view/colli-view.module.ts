import {NgModule} from "@angular/core";
import {MaterialModule} from "../../../modules/material.module";
import {CommonModule} from "@angular/common";
import {DvoCardComponent} from "../../../shared/dvo-card/dvo-card.component";
import {ColliViewComponent} from "./colli-view.component";
import {CdkMenu, CdkMenuItem, CdkMenuTrigger} from "@angular/cdk/menu";
import {ColliStatusColorPipePipe} from "../../../shared/colli-status-color-pipe.pipe";

@NgModule({
	imports: [
		MaterialModule,
		CommonModule,
		DvoCardComponent,
		CdkMenu,
		CdkMenuItem,
		CdkMenuTrigger,
		ColliStatusColorPipePipe,
	],
	exports: [
		ColliViewComponent,
	],
	declarations: [
		ColliViewComponent,
	]
})
export class ColliViewModule { }
