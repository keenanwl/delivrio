import {RouterModule, Routes} from "@angular/router";
import {NgModule} from "@angular/core";
import {ReturnPortalsListComponent} from "./return-portals-list/return-portals-list.component";
import {ReturnPortalEditComponent} from "./return-portal-edit/return-portal-edit.component";

const routes: Routes = [
	{
		path: "edit",
		component: ReturnPortalEditComponent,
	},
	{
		path: "",
		component: ReturnPortalsListComponent,
	},
]

@NgModule({
	imports: [RouterModule.forChild(routes)],
	exports: [RouterModule]
})
export class ReturnPortalsRoutingModule {
}
