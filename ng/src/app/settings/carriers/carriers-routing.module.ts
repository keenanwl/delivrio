import {RouterModule, Routes} from "@angular/router";
import {NgModule} from "@angular/core";
import {CarriersListComponent} from "./carriers-list/carriers-list.component";
import {AuthGuard} from "../../guards/authGuard";
import {CarrierEditGLSComponent} from "./carrier-edit-gls/carrier-edit-gls.component";
import {CarrierEditPostNordComponent} from "./carrier-edit-post-nord/carrier-edit-post-nord.component";
import {CarrierEditUspsComponent} from "./carrier-edit-usps/carrier-edit-usps.component";
import {CarrierEditBringComponent} from "./carrier-edit-bring/carrier-edit-bring.component";
import {CarrierEditDaoComponent} from "./carrier-edit-dao/carrier-edit-dao.component";
import {CarrierEditDsvComponent} from "./carrier-edit-dsv/carrier-edit-dsv.component";
import {CarrierEditDfComponent} from "./carrier-edit-df/carrier-edit-df.component";
import {CarrierEditEasyPostComponent} from "./carrier-edit-easy-post/carrier-edit-easy-post.component";

const routes: Routes = [
	{path: '', component: CarriersListComponent, canActivate: [AuthGuard]},
	{path: 'edit/bring', component: CarrierEditBringComponent, canActivate: [AuthGuard]},
	{path: 'edit/dao', component: CarrierEditDaoComponent, canActivate: [AuthGuard]},
	{path: 'edit/danske-fragtmaend', component: CarrierEditDfComponent, canActivate: [AuthGuard]},
	{path: 'edit/dsv', component: CarrierEditDsvComponent, canActivate: [AuthGuard]},
	{path: 'edit/easy-post', component: CarrierEditEasyPostComponent, canActivate: [AuthGuard]},
	{path: 'edit/gls', component: CarrierEditGLSComponent, canActivate: [AuthGuard]},
	{path: 'edit/post-nord', component: CarrierEditPostNordComponent, canActivate: [AuthGuard]},
	{path: 'edit/usps', component: CarrierEditUspsComponent, canActivate: [AuthGuard]},
]

@NgModule({
	imports: [RouterModule.forChild(routes)],
	exports: [RouterModule]
})
export class CarriersRoutingModule {
}
