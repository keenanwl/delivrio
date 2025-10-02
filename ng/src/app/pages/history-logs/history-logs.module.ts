import {NgModule} from "@angular/core";
import {NgxsModule} from "@ngxs/store";
import {FormsModule, ReactiveFormsModule} from "@angular/forms";
import {MaterialModule} from "../../modules/material.module";
import {CommonModule} from "@angular/common";
import {DvoCardComponent} from "../../shared/dvo-card/dvo-card.component";
import {HistoryLogsState} from "./history-logs.ngxs";
import {HistoryLogsRoutingModule} from "./history-logs-routing.module";
import {HistoryLogsComponent} from "./history-logs.component";

@NgModule({
	imports: [
		HistoryLogsRoutingModule,
		NgxsModule.forFeature([
			HistoryLogsState,
		]),
		MaterialModule,
		CommonModule,
		FormsModule,
		ReactiveFormsModule,
		DvoCardComponent,
		HistoryLogsComponent,
	],
	declarations: [
	]
})
export class HistoryLogsModule { }
