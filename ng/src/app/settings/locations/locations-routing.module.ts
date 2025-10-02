import {RouterModule, Routes} from "@angular/router";
import {NgModule} from "@angular/core";
import {LocationEditComponent} from "./location-edit/location-edit.component";
import { LocationsListComponent } from "./locations-list/locations-list.component";

const routes: Routes = [
	{
		path: "edit",
		component: LocationEditComponent,
	},
	{
		path: "",
		component: LocationsListComponent,
	},
]

@NgModule({
	imports: [RouterModule.forChild(routes)],
	exports: [RouterModule]
})
export class LocationsRoutingModule {
}
