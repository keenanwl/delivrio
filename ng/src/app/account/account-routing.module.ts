import {RouterModule, Routes} from "@angular/router";
import {NgModule} from "@angular/core";
import {AccountComponent} from "./account.component";
import {ProfileComponent} from "./profile/profile.component";
import {AuthGuard} from "../guards/authGuard";
import {PlanComponent} from "./plan/plan.component";
import {CompanyInfoComponent} from "./company-info/company-info.component";

const routes: Routes = [
	{
		path: "",
		component: AccountComponent,
		children: [
			{path: 'profile', component: ProfileComponent, canActivate: [AuthGuard]},
			{path: 'plan', component: PlanComponent, canActivate: [AuthGuard]},
			{path: 'company-info', component: CompanyInfoComponent, canActivate: [AuthGuard]},
			{path: '**', redirectTo: "profile"},
		],
	},
]

@NgModule({
	imports: [RouterModule.forChild(routes)],
	exports: [RouterModule]
})
export class AccountRoutingModule {

}
