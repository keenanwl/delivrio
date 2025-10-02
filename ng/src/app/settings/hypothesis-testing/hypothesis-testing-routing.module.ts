import {RouterModule, Routes} from "@angular/router";
import {NgModule} from "@angular/core";
import {HypothesisTestingListComponent} from "./hypothesis-testing-list/hypothesis-testing-list.component";
import {
	HypothesisTestingDeliveryOptionEditComponent
} from "./edit/hypothesis-testing-delivery-option-edit/hypothesis-testing-delivery-option-edit.component";

const routes: Routes = [
	{
		path: "",
		component: HypothesisTestingListComponent,
	},
	{
		path: "edit",
		component: HypothesisTestingDeliveryOptionEditComponent,
	},
]

@NgModule({
	imports: [RouterModule.forChild(routes)],
	exports: [RouterModule]
})
export class HypothesisTestingRoutingModule {
}
