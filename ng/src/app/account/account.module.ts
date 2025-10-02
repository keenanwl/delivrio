import {NgModule} from "@angular/core";
import {AccountRoutingModule} from "./account-routing.module";
import {NgxsModule} from "@ngxs/store";
import {MaterialModule} from "../modules/material.module";
import {CommonModule} from "@angular/common";
import {DvoCardComponent} from "../shared/dvo-card/dvo-card.component";
import {FormsModule, ReactiveFormsModule} from "@angular/forms";
import {NgxsFormPluginModule} from "@ngxs/form-plugin";
import {NgxsFormErrorsPluginModule} from "../plugins/ngxs-form-errors/ngxs-form-errors.module";
import {AccountComponent} from "./account.component";
import {CompanyInfoComponent} from "./company-info/company-info.component";
import {PlanComponent} from "./plan/plan.component";
import {ProfileComponent} from "./profile/profile.component";
import {CompanyInfoState} from "./company-info/company-info.ngxs";
import {LoginState} from "../login/login.ngxs";
import {PlanState} from "./plan/plan.ngxs";
import {ProfileState} from "./profile/profile.ngxs";

@NgModule({
	imports: [
		AccountRoutingModule,
		NgxsModule.forFeature([
			CompanyInfoState,
			LoginState,
			PlanState,
			ProfileState,
		]),
		MaterialModule,
		CommonModule,
		FormsModule,
		ReactiveFormsModule,
		DvoCardComponent,
		NgxsFormPluginModule,
		NgxsFormErrorsPluginModule,
	],
	declarations: [
		AccountComponent,
		CompanyInfoComponent,
		PlanComponent,
		ProfileComponent,
	]
})
export class AccountModule { }
