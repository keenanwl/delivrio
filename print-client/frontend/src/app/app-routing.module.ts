import { NgModule } from '@angular/core';
import { RouterModule, Routes } from '@angular/router';
import {AuthGuardRegister} from "./guards/auth-guard-register";
import {AuthGuardRegistered} from "./guards/auth-guard-registered";

const routes: Routes = [
	{
		path: '',
		loadChildren: () => import('./register/register.module').then(m => m.RegisterModule),
		canActivate: [AuthGuardRegister],
	},
	{
		path: 'dashboard',
		loadChildren: () => import('./dashboard/dashboard.module').then(m => m.DashboardModule),
		canActivate: [AuthGuardRegistered],
	},
	{
		path: 'settings',
		loadChildren: () => import('./settings/settings.module').then(m => m.SettingsModule),
		canActivate: [AuthGuardRegistered],
	},
];

@NgModule({
	imports: [RouterModule.forRoot(routes, { useHash: true })],
	exports: [RouterModule],
	providers: [AuthGuardRegister, AuthGuardRegistered]
})
export class AppRoutingModule { }
