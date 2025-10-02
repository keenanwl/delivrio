import {NgModule} from "@angular/core";
import {FormsModule, ReactiveFormsModule} from "@angular/forms";
import {MaterialModule} from "src/app/modules/material.module";
import {CarriersRoutingModule} from "./carriers-routing.module";
import {NgxsModule} from "@ngxs/store";
import {CommonModule} from "@angular/common";
import {DvoCardComponent} from "../../shared/dvo-card/dvo-card.component";
import {CarriersListState} from "./carriers-list/carriers-list.ngxs";
import {CarriersListComponent} from "./carriers-list/carriers-list.component";
import {NewCarrierAgreementDialogComponent} from "./carriers-list/new-carrier-agreement-dialog.component";
import {NgxsFormPluginModule} from "@ngxs/form-plugin";
import {NgxsFormErrorsPluginModule} from "../../plugins/ngxs-form-errors/ngxs-form-errors.module";
import {CarrierEditPostNordComponent} from './carrier-edit-post-nord/carrier-edit-post-nord.component';
import {CarrierEditPostNordState} from "./carrier-edit-post-nord/carrier-edit-post-nord.ngxs";
import {CarrierEditUspsComponent} from './carrier-edit-usps/carrier-edit-usps.component';
import {CarrierEditUSPSState} from "./carrier-edit-usps/carrier-edit-usps.ngxs";
import {CarrierEditGLSState} from "./carrier-edit-gls/carrier-edit-gls.ngxs";
import {CarrierEditGLSComponent} from "./carrier-edit-gls/carrier-edit-gls.component";
import {CarrierEditBringComponent} from "./carrier-edit-bring/carrier-edit-bring.component";
import {CarrierEditBringState} from "./carrier-edit-bring/carrier-edit-bring.ngxs";
import {CarrierEditDaoComponent} from "./carrier-edit-dao/carrier-edit-dao.component";
import {CarrierEditDAOState} from "./carrier-edit-dao/carrier-edit-dao.ngxs";
import {CarrierEditDsvComponent} from './carrier-edit-dsv/carrier-edit-dsv.component';
import {AlphabetizePipe} from "../../shared/alphabetize.pipe";
import {CarrierEditDSVState} from "./carrier-edit-dsv/carrier-edit-dsv.ngxs";
import {CarrierEditDfComponent} from './carrier-edit-df/carrier-edit-df.component';
import {CarrierEditDFState} from "./carrier-edit-df/carrier-edit-df.ngxs";
import {CarrierEditEasyPostState} from "./carrier-edit-easy-post/carrier-edit-easy-post.ngxs";

@NgModule({
    imports: [
        CarriersRoutingModule,
        NgxsModule.forFeature([
            CarriersListState,
            CarrierEditBringState,
            CarrierEditDAOState,
            CarrierEditDFState,
            CarrierEditDSVState,
            CarrierEditEasyPostState,
            CarrierEditGLSState,
            CarrierEditPostNordState,
            CarrierEditUSPSState,
        ]),
        MaterialModule,
        CommonModule,
        FormsModule,
        ReactiveFormsModule,
        DvoCardComponent,
        NgxsFormPluginModule,
        NgxsFormErrorsPluginModule,
        AlphabetizePipe,
    ],
	declarations: [
		CarrierEditBringComponent,
		CarrierEditDaoComponent,
		CarrierEditGLSComponent,
		CarriersListComponent,
		NewCarrierAgreementDialogComponent,
        CarrierEditPostNordComponent,
        CarrierEditUspsComponent,
        CarrierEditDsvComponent,
        CarrierEditDfComponent,
	]
})
export class CarriersModule { }
