import {RouterModule, Routes} from "@angular/router";
import {ProductsComponent} from "./products.component";
import {ProductEditComponent} from "./product-edit/product-edit.component";
import {ProductsListComponent} from "./products-list/products-list.component";
import { ProductTagsEditComponent } from "./product-tags-edit/product-tags-edit.component";
import {NgModule} from "@angular/core";

const routes: Routes = [
	{
		path: "",
		component: ProductsComponent,
		children: [
			{
				path: "edit",
				component: ProductEditComponent,
			},
			{
				path: "tags/edit",
				component: ProductTagsEditComponent,
			},
			{
				path: "",
				component: ProductsListComponent,
			},
		]
	},
]

@NgModule({
	imports: [RouterModule.forChild(routes)],
	exports: [RouterModule]
})
export class ProductsRoutingModule {

}
