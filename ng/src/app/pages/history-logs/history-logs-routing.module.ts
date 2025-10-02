import {RouterModule, Routes} from "@angular/router";
import {NgModule} from "@angular/core";
import {HistoryLogsComponent} from "./history-logs.component";

const routes: Routes = [
	{
		path: "",
		component: HistoryLogsComponent,
	},
]

@NgModule({
	imports: [RouterModule.forChild(routes)],
	exports: [RouterModule]
})
export class HistoryLogsRoutingModule {
}
