import {Component, OnInit} from "@angular/core";
import {Observable} from "rxjs";
import {Actions, ofActionDispatched, Store} from "@ngxs/store";
import {COMMA, ENTER} from "@angular/cdk/keycodes";
import {MatChipInputEvent} from "@angular/material/chips";
import {ProductTagsEditModel, ProductTagsEditState} from "./product-tags-edit.ngxs";
import {ProductTagsEditTagsEditActions} from "./product-tags-edit.actions";
import FetchProductTagsEdit = ProductTagsEditTagsEditActions.FetchProductTagsEdit;
import SaveTagList = ProductTagsEditTagsEditActions.SaveTagList;
import SaveTagListSuccess = ProductTagsEditTagsEditActions.SaveTagListSuccess;
import DeleteTag = ProductTagsEditTagsEditActions.DeleteTag;

@Component({
	selector: 'product-tags-edit',
	styleUrls: ['product-tags-edit.component.scss'],
	templateUrl: 'product-tags-edit.component.html',
})
export class ProductTagsEditComponent implements OnInit {

	readonly separatorKeysCodes = [ENTER, COMMA] as const;
	product$: Observable<ProductTagsEditModel>;

	tagsToSave = new Set<string>([]);

	constructor(
		private store: Store,
		private actions$: Actions,
	) {
		this.product$ = store.select(ProductTagsEditState.get);
		this.store.dispatch([new FetchProductTagsEdit()]);
	}

	ngOnInit() {
		this.actions$.pipe(ofActionDispatched(SaveTagListSuccess)).subscribe(() => this.tagsToSave.clear());
	}

	remove(tag: string) {
		if (tag) {
			this.tagsToSave.delete(tag);
		}
	}

	add(tag: MatChipInputEvent) {
		if (tag?.value) {
			const list = tag.value.split(',').map((t) => this.tagsToSave.add(t.trim()));
			tag.chipInput!.clear();
		}
	}

	save() {
		this.store.dispatch(new SaveTagList([...this.tagsToSave]));
	}

	delete(id: string) {
		this.store.dispatch(new DeleteTag(id));
	}

}
