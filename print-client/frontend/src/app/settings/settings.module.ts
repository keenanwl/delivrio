import { NgModule } from '@angular/core';
import { CommonModule } from '@angular/common';
import {MatInputModule} from "@angular/material/input";
import {MatButtonModule} from "@angular/material/button";
import {MatIconModule} from "@angular/material/icon";
import {NgxsModule} from "@ngxs/store";
import {SettingsRoutingModule} from "./settings-routing.module";
import {SettingsComponent} from "./settings.component";
import {SettingsState} from "./settings.ngxs";
import {MatListModule} from "@angular/material/list";
import {MatProgressSpinnerModule} from "@angular/material/progress-spinner";
import {MatDialogModule} from "@angular/material/dialog";

@NgModule({
	declarations: [
		SettingsComponent,
	],
	imports: [
		NgxsModule.forFeature([
			SettingsState,
		]),
		CommonModule,
		SettingsRoutingModule,
		MatInputModule,
		MatButtonModule,
		MatIconModule,
		MatListModule,
		MatProgressSpinnerModule,
		MatDialogModule,
	]
})
export class SettingsModule { }
