import {NgModule} from "@angular/core";
import {TimeRangeSelectorComponent} from "./time-range-selector.component";
import {MaterialModule} from "../../modules/material.module";
import {ZeroPadPipe} from "./zero-pad.pipe";
import {CommonModule} from "@angular/common";

@NgModule({
	imports: [MaterialModule, CommonModule],
	declarations: [
		ZeroPadPipe,
		TimeRangeSelectorComponent,
	],
	exports: [TimeRangeSelectorComponent, TimeRangeSelectorComponent],
})
export class TimeRangeSelectorModule { }
