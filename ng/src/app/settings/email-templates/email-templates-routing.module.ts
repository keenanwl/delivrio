import {RouterModule, Routes} from "@angular/router";
import {NgModule} from "@angular/core";
import {EmailTemplatesListComponent} from "./email-templates-list/email-templates-list.component";
import {EmailTemplateEditComponent} from "./email-template-edit/email-template-edit.component";

const routes: Routes = [
	{
		path: "",
		component: EmailTemplatesListComponent,
	},
	{
		path: "edit",
		component: EmailTemplateEditComponent,
	},
]

@NgModule({
	imports: [RouterModule.forChild(routes)],
	exports: [RouterModule]
})
export class EmailTemplatesRoutingModule {
}
