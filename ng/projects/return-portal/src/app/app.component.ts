import {Component, DoBootstrap, Injector} from '@angular/core';
import {CommonModule} from '@angular/common';
import {RouterOutlet} from '@angular/router';
import {createCustomElement} from "@angular/elements";
import {ReturnPortalViewerModule} from "../../../../src/app/settings/return-portal-viewer/return-portal-viewer.module";
import {HttpClientModule} from "@angular/common/http";
import {
	ReturnPortalContainerModule
} from "../../../../src/app/settings/return-portal-viewer/return-portal-container/return-portal-container.module";
import {
	ReturnPortalContainerComponent
} from "../../../../src/app/settings/return-portal-viewer/return-portal-container/return-portal-container.component";

// This is just the elements wrapper around

@Component({
	selector: 'app-root',
	standalone: true,
	imports: [
		CommonModule,
		RouterOutlet,
		ReturnPortalViewerModule,
		HttpClientModule,
		ReturnPortalContainerModule,
	],
	templateUrl: './app.component.html',
	styleUrls: ['./app.component.scss'],
})
export class AppComponent implements DoBootstrap {
	constructor(private injector: Injector) {
	}
	ngDoBootstrap() {
		const el = createCustomElement(ReturnPortalContainerComponent, {injector: this.injector });
		customElements.define('delivrio-return-portal', el);
	}
}
