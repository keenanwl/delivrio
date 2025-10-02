import {RouterModule, Routes} from "@angular/router";
import {NgModule} from "@angular/core";
import {OrdersComponent} from "./orders.component";
import {OrderViewComponent} from "./order-view/order-view.component";
import {OrderEditComponent} from "./order-edit/order-edit.component";

const routes: Routes = [
	{
		path: "package/edit",
		component: OrderEditComponent,
	},
	{
		path: "view",
		component: OrderViewComponent,
	},
	{
		path: "",
		component: OrdersComponent,
	},
]

@NgModule({
	imports: [RouterModule.forChild(routes)],
	exports: [RouterModule]
})
export class OrdersRoutingModule {
}
