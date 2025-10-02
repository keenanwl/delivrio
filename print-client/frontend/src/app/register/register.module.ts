import { NgModule } from '@angular/core';
import { CommonModule } from '@angular/common';

import { RegisterRoutingModule } from './register-routing.module';

import { RegisterComponent } from './register.component';
import {MatInputModule} from "@angular/material/input";
import {MatButtonModule} from "@angular/material/button";
import {MatIconModule} from "@angular/material/icon";
import {RegisterState} from "./register.ngxs";
import {NgxsModule} from "@ngxs/store";


@NgModule({
	declarations: [
		RegisterComponent
	],
	imports: [
		NgxsModule.forFeature([
			RegisterState,
		]),
		CommonModule,
		RegisterRoutingModule,
		MatInputModule,
		MatButtonModule,
		MatIconModule
	]
})
export class RegisterModule { }
