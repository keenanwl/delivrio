import {Component, OnDestroy, OnInit} from '@angular/core';
import {Observable} from "rxjs";
import {FormArray, FormControl, FormGroup} from "@angular/forms";
import {Actions, ofActionSuccessful, Store} from "@ngxs/store";
import {ActivatedRoute} from "@angular/router";
import {ProductActions} from "./product-edit.actions";
import {ProductModel, ProductState} from './product-edit.ngxs';
import SetProductID = ProductActions.SetProductID;
import FetchProduct = ProductActions.FetchProduct;
import FetchProductResponse = ProductActions.FetchProductResponse;
import {
	ProductStatus,
	UpdateProductInput,
	UpdateProductVariantIdInput,
	UpdateProductVariantInput
} from 'src/generated/graphql';
import {COMMA, ENTER} from "@angular/cdk/keycodes";
import {MatChipInputEvent} from '@angular/material/chips';
import {MatAutocompleteSelectedEvent} from "@angular/material/autocomplete";
import {MatDialog} from '@angular/material/dialog';
import {AppActions} from "../../app.actions";
import AppChangeRoute = AppActions.AppChangeRoute;
import {Paths} from "../../app-routing.module";
import AddTag = ProductActions.AddTag;
import RemoveTag = ProductActions.RemoveTag;
import SaveFormUpdate = ProductActions.SaveFormUpdate;
import CreateVariant = ProductActions.CreateVariant;
import UploadImage = ProductActions.UploadImage;
import NextImage = ProductActions.NextImage;
import PreviousImage = ProductActions.PreviousImage;
import AddImageVariant = ProductActions.AddImageVariant;
import RemoveImageVariant = ProductActions.RemoveImageVariant;
import Reset = ProductActions.Reset;
import DeleteImage = ProductActions.DeleteImage;
import SetVariants = ProductActions.SetVariants;
import ArchiveVariant = ProductActions.ArchiveVariant;
import {EditInventoryComponent} from "./dialogs/edit-inventory/edit-inventory.component";
import MustInventoryItem = ProductActions.MustInventoryItem;

@Component({
	selector: 'app-product-edit',
	templateUrl: './product-edit.component.html',
	styleUrls: ['./product-edit.component.scss']
})
export class ProductEditComponent implements OnInit, OnDestroy {

	readonly separatorKeysCodes = [ENTER, COMMA] as const;
	product$: Observable<ProductModel>;

	editForm = new FormGroup({
		title: new FormControl<string>('', {nonNullable: true}),
		bodyHTML: new FormControl<string>('', {nonNullable: true}),
		productTags: new FormControl<FetchProductResponse['productTags']>([], {nonNullable: true}),
		status: new FormControl<FetchProductResponse['status']>(ProductStatus.Active, {nonNullable: true}),
	});

	variantsForm = new FormGroup({
		variants: new FormArray<ReturnType<typeof this.newProductVariantForm>>([]),
	});

	constructor(private store: Store,
	            private route: ActivatedRoute,
				private dialog: MatDialog,
	            private actions$: Actions) {
		this.product$ = store.select(ProductState.get);
	}

	ngOnInit(): void {
		this.route.queryParams
			.subscribe((params) => {
				this.store.dispatch([
					new SetProductID(!!params.id ? params.id : ''),
					new FetchProduct(),
				]);
			});

		this.actions$.pipe(ofActionSuccessful(SetVariants))
			.subscribe(() => {
				const variants = this.store.selectSnapshot(ProductState.get).variantsForm.model.variants;
				if (!!variants) {
					this.variantsForm.controls.variants.clear();
					variants.forEach((v, i) => {
						//this.imagePreview[i] = new ImageSnippet(v.imgURL || '', 'loaded');
						this.variantsForm.controls.variants.push(this.newProductVariantForm(
							v.id,
							v.dimensionHeight,
							v.dimensionWidth,
							v.dimensionLength,
							v.weightG,
							v.eanNumber,
							v.description,
						));
					})
				}
			});
	}

	ngOnDestroy() {
		this.store.dispatch(new Reset());
	}

	newProductVariantForm(
		id: string,
		height: number | null = null,
	    width: number | null = null,
		length: number | null = null,
		weight: number | null = null,
		ean: string | null = '',
		description: string | null = ''
	) {
		return new FormGroup({
			dimensionHeight: new FormControl<number | null>(height),
			dimensionWidth: new FormControl<number | null>(width),
			dimensionLength: new FormControl<number | null>(length),
			weightG: new FormControl<number | null>(weight),
			eanNumber: new FormControl<string | null>(ean),
			description: new FormControl<string | null>(description),
			id: new FormControl<string>(id, {nonNullable: true}),
		});
	}

	processFile(imageInput: any) {
		const file: File = imageInput.files[0];
		const reader = new FileReader();

		// Check for dupe requests
		reader.addEventListener('load', (event: any) => {
			this.store.dispatch(new UploadImage(event.target.result));
		});

		reader.readAsDataURL(file);
	}

	clearImage(id: string) {
		this.store.dispatch(new DeleteImage({imageID: id}));
	}

	manageTags() {
		this.store.dispatch(new AppChangeRoute({path: Paths.PRODUCTS_TAGS_EDIT, queryParams: {}}));
	}

	remove(tag: string) {
		this.store.dispatch(new RemoveTag(tag));
	}

	add(tag: MatChipInputEvent) {

	}

	select(tag: MatAutocompleteSelectedEvent) {
		this.store.dispatch(new AddTag(tag.option.value));
	}

	addVariant() {
		this.store.dispatch(new CreateVariant());
	}

	save() {

		//const images = this.imagePreview.map((i) => i.src);
		const variants: UpdateProductVariantIdInput[] = this.variantsForm.getRawValue().variants.map((v) => {
			return {
				id: v.id,
				variant: Object.assign({}, v, {id: undefined}) as UpdateProductVariantInput
			};
		});

		const values = Object.assign<{}, UpdateProductInput, {productTags: undefined, addProductTagIDs: string[] | undefined}>(
			{},
			this.editForm.getRawValue(),
			{
				productTags: undefined,
				addProductTagIDs: this.editForm.getRawValue().productTags?.map((t) => t?.id)
			});

		this.store.dispatch(new SaveFormUpdate({input: values, variants: variants}));

	}

	nextImage() {
		this.store.dispatch(new NextImage());
	}

	previousImage() {
		this.store.dispatch(new PreviousImage());
	}

	updateImageVariant(checked: boolean, variantIndex: number) {
		const variant = this.variantsForm.controls.variants.at(variantIndex);
		const variantID = variant.controls.id.value;
		if (checked) {
			this.store.dispatch(new AddImageVariant({variantID}));
		} else {
			this.store.dispatch(new RemoveImageVariant({variantID}));
		}
	}

	variantConnectedToImage(variantIndex: number): boolean {
		const state = this.store.selectSnapshot(ProductState.get);
		const variant = this.variantsForm.controls.variants.at(variantIndex);
		const variantID = variant.controls.id.value;

		return state.productImages.some((i, imageIndex) => {
			if (imageIndex === state.imageIndex) {
				return i.productVariant?.some((v) => {
					return v.id === variantID;
				});
			}
			return false;
		});
	}

	archiveVariant(id: string) {
		this.store.dispatch(new ArchiveVariant(id));
	}

	openInventory(productVariantID: string) {
		this.store.dispatch(new MustInventoryItem(productVariantID));
		this.dialog.open(EditInventoryComponent);
	}

}
