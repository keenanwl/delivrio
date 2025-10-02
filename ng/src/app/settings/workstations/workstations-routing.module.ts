import {RouterModule, Routes} from "@angular/router";
import {NgModule} from "@angular/core";
import {AuthGuard} from "../../guards/authGuard";
import {WorkstationsListComponent} from "./workstations-list/workstations-list.component";
import {WorkstationEditComponent} from "./workstation-edit/workstation-edit.component";

const routes: Routes = [
	{path: '', component: WorkstationsListComponent, canActivate: [AuthGuard]},
	{path: 'edit', component: WorkstationEditComponent, canActivate: [AuthGuard]},
]

@NgModule({
	imports: [RouterModule.forChild(routes)],
	exports: [RouterModule]
})
export class WorkstationsRoutingModule {
}
