import {RouterModule, Routes} from "@angular/router";
import {NgModule} from "@angular/core";
import {ConnectionEditComponent} from "./connection-edit/connection-edit.component";
import {ConnectionsListComponent} from "./connections-list/connections-list.component";

const routes: Routes = [
	{
		path: "edit",
		component: ConnectionEditComponent,
	},
	{
		path: "",
		component: ConnectionsListComponent,
	},
]

@NgModule({
	imports: [RouterModule.forChild(routes)],
	exports: [RouterModule]
})
export class ConnectionsRoutingModule {
}
