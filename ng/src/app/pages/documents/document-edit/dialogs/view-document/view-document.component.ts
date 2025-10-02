import {Component} from '@angular/core';
import {Observable} from "rxjs";
import {Store} from "@ngxs/store";
import {DocumentEditModel, DocumentEditState} from "../../document-edit.ngxs";
import {DialogRef} from "@angular/cdk/dialog";

@Component({
  selector: 'app-view-document',
  templateUrl: './view-document.component.html',
  styleUrl: './view-document.component.scss'
})
export class ViewDocumentComponent {
	state$: Observable<DocumentEditModel>;

	constructor(
		private store: Store,
		private ref: DialogRef,
	) {
		this.state$ = store.select(DocumentEditState.get);
	}

	prepareDownloadName(name?: string): string {
		if (!name) {
			return "delivrio_document.pdf"
		}

		let out = name?.replace(" ", "_") || '';
		out = out.substring(0, out.length < 30 ? out.length : 30);
		out = out.replace(/_\s*$/, '');

		return `deliverio${out.length > 0 ? '_' + out : '' }.pdf`
	}

	close() {
		this.ref.close();
	}
}
