import {RouterModule, Routes} from "@angular/router";
import {NgModule} from "@angular/core";
import {DocumentEditComponent} from "./document-edit/document-edit.component";
import {DocumentsListComponent} from "./documents-list/documents-list.component";

const routes: Routes = [
	{
		path: "",
		component: DocumentsListComponent,
	},
	{
		path: "edit",
		component: DocumentEditComponent,
	},
]

@NgModule({
	imports: [RouterModule.forChild(routes)],
	exports: [RouterModule]
})
export class DocumentsRoutingModule {
}
