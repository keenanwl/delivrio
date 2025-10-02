import {NgModule} from '@angular/core';
import {Routes, RouterModule} from '@angular/router';
import {PasswordResetComponent} from "./login/password-reset/password-reset.component";
import {RequestPasswordResetComponent} from "./login/request-password-reset/request-password-reset.component";
import {LoginComponent} from "./login/login.component";
import {AuthGuard} from "./guards/authGuard";
import {LoggedInWrapperComponent} from "./logged-in-wrapper/logged-in-wrapper.component";
import {NonAuthGuard} from "./guards/nonAuthGuard.tsauthGuard";
import {SettingsComponent} from './settings/settings.component';

export namespace Paths {
	export const ACCOUNT = '/account';
	export const ACCOUNT_COMPANY_INFO = '/account/company-info';
	export const ACCOUNT_PLAN = '/account/plan';
	export const ACCOUNT_PROFILE = '/account/profile';
	export const API_DOCS = '/static/rest/v1/api-docs';
	export const CONSOLIDATIONS = '/consolidations';
	export const CONSOLIDATION_EDIT = '/consolidations/edit';
	export const DASHBOARD = '/dashboard';
	export const DOCUMENTS = '/documents';
	export const DOCUMENT_EDIT = '/documents/edit';
	export const LOGIN = '/login';
	export const NOT_FOUND = '/404';
	export const ORDERS = '/orders';
	export const ORDERS_PACKAGE_EDIT = '/orders/package/edit';
	export const ORDERS_VIEW = '/orders/view';
	export const PASSWORD_RESET = '/password-reset';
	export const PRODUCTS = '/products';
	export const PRODUCTS_EDIT = '/products/edit';
	export const PRODUCTS_TAGS_EDIT = '/products/tags/edit';
	export const REGISTER = '/register';
	export const REGISTER_2 = '/register/2';
	export const REGISTER_3 = '/register/3';
	export const REQUEST_PASSWORD = '/request-password';
	export const RETURN_EDIT = '/returns/edit';
	export const RETURN_PORTAL_VIEWER = '/return-portal-viewer';
	export const RETURNS = '/returns';
	export const HISTORY_LOGS = '/history-logs';
	export const RETURN_VIEW = '/returns/view';
	export const ROOT = '';
	export const SETTINGS = '/settings';
	export const SETTINGS_API_TOKENS = '/settings/api-tokens';
	export const SETTINGS_CARRIERS = '/settings/carriers';
	export const SETTINGS_CARRIERS_EDIT_GLS = '/settings/carriers/edit/gls';
	export const SETTINGS_CARRIERS_EDIT_POST_NORD = '/settings/carriers/edit/post-nord';
	export const SETTINGS_CARRIERS_EDIT_BRING = '/settings/carriers/edit/bring';
	export const SETTINGS_CARRIERS_EDIT_DAO = '/settings/carriers/edit/dao';
	export const SETTINGS_CARRIERS_EDIT_DF = '/settings/carriers/edit/danske-fragtmaend';
	export const SETTINGS_CARRIERS_EDIT_DSV = '/settings/carriers/edit/dsv';
	export const SETTINGS_CARRIERS_EDIT_EASY_POST = '/settings/carriers/edit/easy-post';
	export const SETTINGS_CARRIERS_EDIT_USPS = '/settings/carriers/edit/usps';
	export const SETTINGS_CONNECTIONS = '/settings/connections';
	export const SETTINGS_CONNECTIONS_EDIT = '/settings/connections/edit';
	export const SETTINGS_DELIVERY_OPTIONS = '/settings/delivery-options';
	export const SETTINGS_DELIVERY_OPTIONS_EDIT_BRING = '/settings/delivery-options/edit/bring';
	export const SETTINGS_DELIVERY_OPTIONS_EDIT_DAO = '/settings/delivery-options/edit/dao';
	export const SETTINGS_DELIVERY_OPTIONS_EDIT_DF = '/settings/delivery-options/edit/danske-fragtmaend';
	export const SETTINGS_DELIVERY_OPTIONS_EDIT_DSV = '/settings/delivery-options/edit/dsv';
	export const SETTINGS_DELIVERY_OPTIONS_EDIT_EASY_POST = '/settings/delivery-options/edit/easy-post';
	export const SETTINGS_DELIVERY_OPTIONS_EDIT_GLS = '/settings/delivery-options/edit/gls';
	export const SETTINGS_DELIVERY_OPTIONS_EDIT_POST_NORD = '/settings/delivery-options/edit/post-nord';
	export const SETTINGS_DELIVERY_OPTIONS_EDIT_USPS = '/settings/delivery-options/edit/usps';
	export const SETTINGS_EMAIL_TEMPLATES = '/settings/email-templates';
	export const SETTINGS_EMAIL_TEMPLATE_EDIT = '/settings/email-templates/edit';
	export const SETTINGS_NOTIFICATIONS = '/settings/notifications';
	export const SETTINGS_HYPOTHESIS_TESTING = '/settings/hypothesis-testing';
	export const SETTINGS_HYPOTHESIS_TESTING_EDIT = '/settings/hypothesis-testing/edit';
	export const SETTINGS_LOCATIONS = '/settings/locations';
	export const SETTINGS_LOCATION_EDIT = '/settings/locations/edit';
	export const SETTINGS_PACKAGING = '/settings/packaging';
	export const SETTINGS_PACKAGING_EDIT = '/settings/packaging/edit';
	export const SETTINGS_RATE_SHEETS_EDIT = '/settings/rate-sheets/edit';
	export const SETTINGS_RATE_SHEETS_LIST = '/settings/rate-sheets';
	export const SETTINGS_RETURN_PORTALS_EDIT = '/settings/return-portals/edit';
	export const SETTINGS_RETURN_PORTALS_LIST = '/settings/return-portals';
	export const SETTINGS_WORKSTATIONS = '/settings/workstations';
	export const SETTINGS_WORKSTATIONS_EDIT = '/settings/workstations/edit';
	export const SHIPMENT_EDIT = '/shipments/edit';
	export const SHIPMENT_VIEW = '/shipments/view';
	export const SHIPMENTS = '/shipments';
	export const SETTINGS_USERS = '/settings/users/list';
	export const SETTINGS_USERS_EDIT = '/settings/users/edit';
	export const SETTINGS_USERS_GROUP = '/settings/users/groups/list';
	export const SETTINGS_USERS_GROUPS_EDIT = '/settings/users/groups/edit';
}

const routes: Routes = [
	{path: '', component: LoggedInWrapperComponent, canActivate: [AuthGuard], children: [
		{
			path: 'dashboard',
			loadChildren: () => import('./pages/dashboard/dashboard.module').then(m => m.DashboardModule),
			canActivate: [AuthGuard],
			title: "Dashboard - DELIVRIO"
		},
		{
			path: 'orders',
			loadChildren: () => import('./orders/orders.module').then(m => m.OrdersModule),
			canActivate: [AuthGuard],
			title: "Orders - DELIVRIO",
		},
		{
			path: 'shipments',
			loadChildren: () => import('./shipments/shipments.module').then(m => m.ShipmentsModule),
			canActivate: [AuthGuard],
			title: "Shipments - DELIVRIO",
		},
		{
			path: 'products',
			loadChildren: () => import('./products/products.module').then(m => m.ProductsModule),
			canActivate: [AuthGuard],
			title: "Products - DELIVRIO",
		},
		{
			path: 'consolidations',
			loadChildren: () => import('./consolidations/consolidations.module').then(m => m.ConsolidationsModule),
			canActivate: [AuthGuard],
			title: "Consolidations - DELIVRIO",
		},
		{
			path: 'documents',
			loadChildren: () => import('./pages/documents/documents.module').then(m => m.DocumentsModule),
			canActivate: [AuthGuard],
			title: "Documents - DELIVRIO",
		},
		{
			path: 'returns',
			loadChildren: () => import('./returns/returns.module').then(m => m.ReturnsModule),
			canActivate: [AuthGuard],
			title: "Returns - DELIVRIO",
		},
		{
			path: 'account',
			loadChildren: () => import('./account/account.module').then(m => m.AccountModule),
			canActivate: [AuthGuard],
			title: "Account - DELIVRIO",
		},
		{
			path: 'history-logs',
			loadChildren: () => import('./pages/history-logs/history-logs.module').then(m => m.HistoryLogsModule),
			canActivate: [AuthGuard],
			title: "History & Logs",
		},
		{
			path: 'settings', component: SettingsComponent, canActivate: [AuthGuard], children: [
				{path: 'api-tokens', canActivate: [AuthGuard], loadChildren: () => import('./settings/api-tokens/api-tokens.module').then(m => m.APITokensModule)},
				{path: 'connections', canActivate: [AuthGuard], loadChildren: () => import('./settings/connections/connections.module').then(m => m.ConnectionsModule)},
				{path: 'carriers', canActivate: [AuthGuard], loadChildren: () => import('./settings/carriers/carriers.module').then(m => m.CarriersModule)},
				{path: 'email-templates', canActivate: [AuthGuard], loadChildren: () => import('./settings/email-templates/email-templates.module').then(m => m.EmailTemplatesModule)},
				{path: 'notifications', canActivate: [AuthGuard], loadChildren: () => import('./settings/notifications/notifications.module').then(m => m.NotificationsModule)},
				{path: 'hypothesis-testing', canActivate: [AuthGuard], loadChildren: () => import('./settings/hypothesis-testing/hypothesis-testing.module').then(m => m.HypothesisTestingModule)},
				{path: 'return-portals', canActivate: [AuthGuard], loadChildren: () => import('./settings/return-portals/return-portals.module').then(m => m.ReturnPortalsModule)},
				{path: 'delivery-options', canActivate: [AuthGuard], loadChildren: () => import('./settings/delivery-options/delivery-options.module').then(m => m.DeliveryOptionsModule)},
				{path: 'users', canActivate: [AuthGuard], loadChildren: () => import('./settings/user-seats/user-seats.module').then(m => m.UserSeatsModule)},
				{path: 'locations', canActivate: [AuthGuard], loadChildren: () => import('./settings/locations/locations.module').then(m => m.LocationsModule)},
				{path: 'packaging', canActivate: [AuthGuard], loadChildren: () => import('./settings/packaging/packaging.module').then(m => m.PackagingModule)},
				{path: 'workstations', canActivate: [AuthGuard], loadChildren: () => import('./settings/workstations/workstations.module').then(m => m.WorkstationsModule)},
				{path: '**', redirectTo: Paths.SETTINGS_CONNECTIONS},
			]
		},
		{
			path: '404',
			loadChildren: () => import('./not-found/not-found.module').then(m => m.NotFoundModule),
			canActivate: [AuthGuard],
			title: "404 - DELIVRIO",
		},
		{pathMatch: 'full', path: '', redirectTo: Paths.ORDERS},
	]},
	{
		path: 'login',
		component: LoginComponent,
		canActivate: [NonAuthGuard],
		title: "Login - DELIVRIO",
	},
	{
		path: 'return-portal-viewer',
		loadChildren: () => import('./settings/return-portal-viewer/return-portal-viewer.module').then(m => m.ReturnPortalViewerModule)
	},
	{path: 'password-reset', component: PasswordResetComponent},
	{path: 'request-password', component: RequestPasswordResetComponent},
	{path: 'register', loadChildren: () => import('./register/register.module').then(m => m.RegisterModule)},

];

@NgModule({
	imports: [RouterModule.forRoot(routes, {onSameUrlNavigation: "reload"})],
	exports: [RouterModule],
	providers: [AuthGuard, NonAuthGuard]
})
export class AppRoutingModule { }
