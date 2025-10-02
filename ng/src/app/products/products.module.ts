import {NgModule} from "@angular/core";
import {ProductsRoutingModule} from "./products-routing.module";
import {ProductsComponent} from "./products.component";
import {ProductEditComponent} from "./product-edit/product-edit.component";
import {ProductsListComponent} from "./products-list/products-list.component";
import {ProductTagsEditComponent} from "./product-tags-edit/product-tags-edit.component";
import {NgxsModule} from "@ngxs/store";
import {ProductsListState} from "./products-list/products-list.ngxs";
import {ProductState} from "./product-edit/product-edit.ngxs";
import {ProductTagsEditState} from "./product-tags-edit/product-tags-edit.ngxs";
import {MaterialModule} from "../modules/material.module";
import {CommonModule, NgOptimizedImage} from "@angular/common";
import {DvoCardComponent} from "../shared/dvo-card/dvo-card.component";
import {FormsModule, ReactiveFormsModule} from "@angular/forms";
import { AddProductDialogComponent } from './products-list/dialogs/add-product-dialog/add-product-dialog.component';
import {NgxsFormPluginModule} from "@ngxs/form-plugin";
import {NgxsFormErrorsPluginModule} from "../plugins/ngxs-form-errors/ngxs-form-errors.module";

@NgModule({
    imports: [
        ProductsRoutingModule,
        NgxsModule.forFeature([
            ProductsListState,
            ProductState,
            ProductTagsEditState,
        ]),
        MaterialModule,
        CommonModule,
        FormsModule,
        ReactiveFormsModule,
        DvoCardComponent,
        NgxsFormPluginModule,
        NgxsFormErrorsPluginModule,
        NgOptimizedImage,
    ],
	declarations: [
		ProductsComponent,
		ProductEditComponent,
		ProductsListComponent,
		ProductTagsEditComponent,
  AddProductDialogComponent,
	]
})
export class ProductsModule { }
