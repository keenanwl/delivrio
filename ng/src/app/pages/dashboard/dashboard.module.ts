import {NgModule} from "@angular/core";
import {NgxsModule} from "@ngxs/store";
import {FormsModule, ReactiveFormsModule} from "@angular/forms";
import {DashboardComponent} from "./dashboard.component";
import {DashboardRoutingModule} from "./dashboard-routing.module";
import {MaterialModule} from "../../modules/material.module";
import {CommonModule} from "@angular/common";
import {DvoCardComponent} from "../../shared/dvo-card/dvo-card.component";
import {DashboardState} from "./dashboard.ngxs";
import {NgApexchartsModule} from "ng-apexcharts";

@NgModule({
	imports: [
		DashboardRoutingModule,
		NgxsModule.forFeature([
			DashboardState,
		]),
		MaterialModule,
		CommonModule,
		FormsModule,
		ReactiveFormsModule,
		DvoCardComponent,
		NgApexchartsModule,
	],
	declarations: [
		DashboardComponent,
	]
})
export class DashboardModule { }
