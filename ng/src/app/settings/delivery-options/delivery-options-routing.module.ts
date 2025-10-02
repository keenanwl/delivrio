import {RouterModule, Routes} from "@angular/router";
import {NgModule} from "@angular/core";
import {AuthGuard} from "../../guards/authGuard";
import {DeliveryOptionsEditGlsComponent} from "./edit/delivery-options-edit-gls/delivery-options-edit-gls.component";
import {
	DeliveryOptionEditPostNordComponent
} from "./edit/delivery-option-edit-post-nord/delivery-option-edit-post-nord.component";
import {DeliveryOptionsListComponent} from "./delivery-options-list/delivery-options-list.component";
import {DeliveryOptionEditUspsComponent} from "./edit/delivery-option-edit-usps/delivery-option-edit-usps.component";
import {DeliveryOptionEditDaoComponent} from "./edit/delivery-option-edit-dao/delivery-option-edit-dao.component";
import {DeliveryOptionEditBringComponent} from "./edit/delivery-option-edit-bring/delivery-option-edit-bring.component";
import {DeliveryOptionEditDsvComponent} from "./edit/delivery-option-edit-dsv/delivery-option-edit-dsv.component";
import {DeliveryOptionEditDfComponent} from "./edit/delivery-option-edit-df/delivery-option-edit-df.component";
import {
	DeliveryOptionEditEasyPostComponent
} from "./edit/delivery-option-edit-easy-post/delivery-option-edit-easy-post.component";

const routes: Routes = [
	{path: '', component: DeliveryOptionsListComponent, canActivate: [AuthGuard]},
	{path: 'edit/bring', component: DeliveryOptionEditBringComponent, canActivate: [AuthGuard]},
	{path: 'edit/dao', component: DeliveryOptionEditDaoComponent, canActivate: [AuthGuard]},
	{path: 'edit/danske-fragtmaend', component: DeliveryOptionEditDfComponent, canActivate: [AuthGuard]},
	{path: 'edit/dsv', component: DeliveryOptionEditDsvComponent, canActivate: [AuthGuard]},
	{path: 'edit/easy-post', component: DeliveryOptionEditEasyPostComponent, canActivate: [AuthGuard]},
	{path: 'edit/gls', component: DeliveryOptionsEditGlsComponent, canActivate: [AuthGuard]},
	{path: 'edit/post-nord', component: DeliveryOptionEditPostNordComponent, canActivate: [AuthGuard]},
	{path: 'edit/usps', component: DeliveryOptionEditUspsComponent, canActivate: [AuthGuard]},
]

@NgModule({
	imports: [RouterModule.forChild(routes)],
	exports: [RouterModule]
})
export class DeliveryOptionsRoutingModule {
}
