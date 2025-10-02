import {NgModule} from "@angular/core";
import {MaterialModule} from "../modules/material.module";
import {NotFoundComponent} from "./not-found.component";
import {NotFoundRoutingModule} from "./not-found-routing.module";

@NgModule({
	imports: [
		NotFoundRoutingModule,
		MaterialModule,
	],
	declarations: [
		NotFoundComponent,
	]
})
export class NotFoundModule { }
