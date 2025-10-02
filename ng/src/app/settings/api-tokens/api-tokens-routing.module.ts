import {RouterModule, Routes} from "@angular/router";
import {NgModule} from "@angular/core";
import {AuthGuard} from "../../guards/authGuard";
import {APITokensComponent} from "./api-tokens.component";

const routes: Routes = [
	{path: '', component: APITokensComponent, canActivate: [AuthGuard]},
]

@NgModule({
	imports: [RouterModule.forChild(routes)],
	exports: [RouterModule]
})
export class APITokensRoutingModule {
}
