import {RouterModule, Routes} from "@angular/router";
import {NgModule} from "@angular/core";
import {ShipmentViewComponent} from "./shipment-view/shipment-view.component";
import {ShipmentEditComponent} from "./shipment-edit/shipment-edit.component";
import {ShipmentsListComponent } from "./shipments-list/shipments-list.component";

const routes: Routes = [
	{
		path: "edit",
		component: ShipmentEditComponent,
	},
	{
		path: "view",
		component: ShipmentViewComponent,
	},
	{
		path: "",
		component: ShipmentsListComponent,
	},
]

@NgModule({
	imports: [RouterModule.forChild(routes)],
	exports: [RouterModule]
})
export class ShipmentsRoutingModule {
}
