import {NgModule} from "@angular/core";
import {ReturnPortalFrameComponent} from "../return-portal-frame/return-portal-frame.component";
import {ReturnPortalContainerComponent} from "./return-portal-container.component";
import {ReturnPortalFrameModule} from "../return-portal-frame/return-portal-frame.module";
import {CommonModule} from "@angular/common";
import {RouterOutlet} from "@angular/router";
import {ReturnPortalViewerModule} from "../return-portal-viewer.module";
import { provideHttpClient, withInterceptorsFromDi } from "@angular/common/http";

@NgModule({ exports: [
        ReturnPortalFrameComponent
    ],
    declarations: [
        ReturnPortalContainerComponent,
    ], imports: [ReturnPortalFrameModule,
        CommonModule,
        RouterOutlet,
        ReturnPortalViewerModule], providers: [provideHttpClient(withInterceptorsFromDi())] })
export class ReturnPortalContainerModule { }
