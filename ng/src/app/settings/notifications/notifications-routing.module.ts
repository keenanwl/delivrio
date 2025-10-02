import {RouterModule, Routes} from "@angular/router";
import {NgModule} from "@angular/core";
import {NotificationsListComponent} from "./notifications-list/notifications-list.component";

const routes: Routes = [
	{
		path: "",
		component: NotificationsListComponent,
	},
]

@NgModule({
	imports: [RouterModule.forChild(routes)],
	exports: [RouterModule]
})
export class NotificationsRoutingModule {
}
