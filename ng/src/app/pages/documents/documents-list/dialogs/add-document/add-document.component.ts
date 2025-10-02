import {Component} from '@angular/core';
import {Store} from "@ngxs/store";
import {DocumentsListActions} from "../../documents-list.actions";
import Create = DocumentsListActions.Create;
import {DialogRef} from "@angular/cdk/dialog";
import {DocumentMergeType} from "../../../../../../generated/graphql";

@Component({
	selector: 'app-add-document',
	templateUrl: './add-document.component.html',
	styleUrl: './add-document.component.scss'
})
export class AddDocumentComponent {

	unsortedKVFn = (a: any, b: any) => 0;

	constructor(private store: Store,
				private ref: DialogRef) {
	}

	create(name: string, mergeType: DocumentMergeType) {
		this.store.dispatch(new Create({name, mergeType}));
		this.close();
	}

	close() {
		this.ref.close();
	}

    protected readonly DocumentMergeType = DocumentMergeType;
}
