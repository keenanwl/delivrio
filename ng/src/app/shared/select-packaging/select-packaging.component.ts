import {Component, EventEmitter, Input, OnInit, Output} from '@angular/core';
import {CommonModule} from '@angular/common';
import {MatButtonModule} from "@angular/material/button";
import {MatIconModule} from "@angular/material/icon";
import {MatListModule} from "@angular/material/list";
import {MatProgressSpinnerModule} from "@angular/material/progress-spinner";
import {Observable} from "rxjs";
import {Store} from "@ngxs/store";
import {MatDialogRef} from "@angular/material/dialog";
import FetchPackaging = SelectPackagingActions.FetchPackaging;
import PackagingResponse = SelectPackagingActions.PackagingResponse;
import {SelectPackagingActions} from "./select-packaging.actions";
import {SelectPackagingModel, SelectPackagingState} from "./select-packaging.ngxs";

@Component({
	imports: [
		CommonModule,
		MatButtonModule,
		MatIconModule,
		MatListModule,
		MatProgressSpinnerModule,
		// State imported at root until better option found
	],
	selector: 'app-select-packaging',
	standalone: true,
	styleUrl: './select-packaging.component.scss',
	templateUrl: './select-packaging.component.html'
})
export class SelectPackagingComponent implements OnInit {

	@Input() helpText = "Make a selection";
	@Output() selected: EventEmitter<PackagingResponse | null> = new EventEmitter();

	packaging$: Observable<SelectPackagingModel>;
	constructor(private store: Store, private ref: MatDialogRef<any>) {
		this.packaging$ = store.select(SelectPackagingState.get);
	}

	ngOnInit(): void {
		this.store.dispatch(new FetchPackaging());
	}

	selectPackaging(p: PackagingResponse | null) {
		this.selected.emit(p);
		this.close();
	}

	close() {
		this.ref.close();
	}
}
