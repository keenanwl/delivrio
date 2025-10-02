import {RouterModule, Routes} from "@angular/router";
import {NgModule} from "@angular/core";
import {ConsolidationsListComponent} from "./consolidations-list/consolidations-list.component";
import {ConsolidationEditComponent} from "./consolidation-edit/consolidation-edit.component";

const routes: Routes = [
	{
		path: "",
		component: ConsolidationsListComponent,
	},
	{
		path: "edit",
		component: ConsolidationEditComponent,
	},
]

@NgModule({
	imports: [RouterModule.forChild(routes)],
	exports: [RouterModule]
})
export class ConsolidationsRoutingModule {
}
