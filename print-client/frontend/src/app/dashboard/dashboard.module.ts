import { NgModule } from '@angular/core';
import {CommonModule, DatePipe, NgIf} from '@angular/common';
import {MatInputModule} from "@angular/material/input";
import {MatButtonModule} from "@angular/material/button";
import {MatIcon, MatIconModule} from "@angular/material/icon";
import {NgxsModule} from "@ngxs/store";
import {DashboardRoutingModule} from "./dashboard-routing.module";
import {DashboardComponent} from "./dashboard.component";
import {DashboardState} from "./dashboard.ngxs";
import {MatListModule} from "@angular/material/list";
import {RelativeTimePipe} from "../pipes/relative-time.pipe";
import {MatTooltipModule} from "@angular/material/tooltip";
import {MatChip, MatChipListbox} from "@angular/material/chips";
import {PrintJobStatusColorPipe} from "../pipes/printjob-status-color-pipe.pipe";

@NgModule({
	declarations: [
		DashboardComponent,
		RelativeTimePipe,
	],
	imports: [
		NgxsModule.forFeature([
			DashboardState,
		]),
		PrintJobStatusColorPipe,
		CommonModule,
		DashboardRoutingModule,
		MatInputModule,
		MatButtonModule,
		MatIconModule,
		MatListModule,
		MatTooltipModule,
		MatChip,
		MatChipListbox,
		MatIcon,
		NgIf,
		DatePipe,
	],
	providers: [DatePipe],
})
export class DashboardModule { }
