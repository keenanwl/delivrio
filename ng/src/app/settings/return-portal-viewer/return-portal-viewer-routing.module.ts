import {RouterModule, Routes} from "@angular/router";
import {NgModule} from "@angular/core";
import {ReturnPortalFrameComponent} from "./return-portal-frame/return-portal-frame.component";

const routes: Routes = [
	{
		path: "",
		component: ReturnPortalFrameComponent,
	},
]

@NgModule({
	imports: [RouterModule.forChild(routes)],
	exports: [RouterModule]
})
export class ReturnPortalViewerRoutingModule {
}
