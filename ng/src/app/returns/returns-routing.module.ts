import {RouterModule, Routes} from "@angular/router";
import {ReturnsComponent} from "./returns.component";
import {NgModule} from "@angular/core";
import {ReturnEditComponent} from "./return-edit/return-edit.component";
import {ReturnsListComponent} from "./returns-list/returns-list.component";
import {ReturnViewComponent} from "./return-view/return-view.component";

const routes: Routes = [
	{
		path: "",
		component: ReturnsComponent,
		children: [
			{
				path: "edit",
				component: ReturnEditComponent,
			},
			{
				path: "view",
				component: ReturnViewComponent,
			},
			{
				path: "",
				component: ReturnsListComponent,
			},
		]
	},
]

@NgModule({
	imports: [RouterModule.forChild(routes)],
	exports: [RouterModule]
})
export class ReturnsRoutingModule {

}
