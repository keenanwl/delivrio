import {RouterModule, Routes} from "@angular/router";
import {NgModule} from "@angular/core";
import {UserSeatsListComponent} from "./user-seats-list/user-seats-list.component";
import {UserSeatEditComponent} from "./user-seat-edit/user-seat-edit.component";
import {SeatGroupsListComponent} from "./seat-groups-list/seat-groups-list.component";
import {SeatGroupEditComponent} from "./seat-group-edit/seat-group-edit.component";
import {AuthGuard} from "../../guards/authGuard";

const routes: Routes = [
	{path: 'list', component: UserSeatsListComponent, canActivate: [AuthGuard]},
	{path: 'edit', component: UserSeatEditComponent, canActivate: [AuthGuard]},
	{path: 'groups/list', component: SeatGroupsListComponent, canActivate: [AuthGuard]},
	{path: 'groups/edit', component: SeatGroupEditComponent, canActivate: [AuthGuard]},
	{path: '', redirectTo: 'users', pathMatch: 'prefix'},
];

@NgModule({
	imports: [RouterModule.forChild(routes)],
	exports: [RouterModule]
})
export class UserSeatsRoutingModule {

}
