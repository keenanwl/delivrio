import {RouterModule, Routes} from "@angular/router";
import {NgModule} from "@angular/core";
import {Register1Component} from "./register1/register1.component";
import {NonAuthGuard} from "../guards/nonAuthGuard.tsauthGuard";
import {Register2Component} from "./register2/register2.component";
import {Register3Component} from "./register3/register3.component";
import {AuthGuard} from "../guards/authGuard";
import {RegisterComponent} from "./register.component";

const routes: Routes = [
	{path: '', component: RegisterComponent, children: [
		{path: '', component: Register1Component, canActivate: [NonAuthGuard]},
		{path: '2', component: Register2Component, canActivate: [AuthGuard]},
		{path: '3', component: Register3Component, canActivate: [AuthGuard]},
	]}
]

@NgModule({
	imports: [RouterModule.forChild(routes)],
	exports: [RouterModule]
})
export class RegisterRoutingModule {
}
