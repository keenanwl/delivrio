import {createApplication} from '@angular/platform-browser';
import {createCustomElement} from '@angular/elements';
import { ApplicationRef } from '@angular/core';
import {
	ReturnPortalContainerComponent
} from "../../../src/app/settings/return-portal-viewer/return-portal-container/return-portal-container.component";
import {appConfig} from "./app/app.config";

(async () => {
	const app: ApplicationRef = await createApplication(appConfig);

	// Define Web Components
	const el = createCustomElement(ReturnPortalContainerComponent, {injector: app.injector });
	customElements.define('delivrio-return-portal', el);
})();
