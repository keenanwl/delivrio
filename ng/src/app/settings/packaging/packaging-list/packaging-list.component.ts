import {Component, OnDestroy, OnInit} from '@angular/core';
import {Observable, Subscription} from "rxjs";
import {Store} from "@ngxs/store";
import {MatDialog} from "@angular/material/dialog";
import {PackagingListActions} from "./packaging-list.actions";
import Clear = PackagingListActions.Clear;
import {PackagingListModel, PackagingListState} from "./packaging-list.ngxs";
import {CreatePackagingComponent} from "./dialogs/create-packaging/create-packaging.component";
import FetchPackagingList = PackagingListActions.FetchPackagingList;
import PackagingResponse = PackagingListActions.PackagingResponse;
import {ArchiveConfirmationComponent} from "./dialogs/archive-confirmation/archive-confirmation.component";

@Component({
  	selector: 'app-packaging-list',
	templateUrl: './packaging-list.component.html',
	styleUrls: ['./packaging-list.component.scss']
})
export class PackagingListComponent implements OnInit, OnDestroy {
	packaging$: Observable<PackagingListModel>;
	subscriptions$: Subscription[] = []

	displayedColumns: string[] = [
		'name',
		'dimensions',
		'carrier',
		'actions',
	];

	boxDisplayIndex = 0;
	boxDisplaySide = [
		"show-front",
		"show-back",
		"show-right",
		"show-left",
		"show-top",
		"show-bottom",
	]

	constructor(
		private store: Store,
		private dialog: MatDialog,
	) {
		this.packaging$ = store.select(PackagingListState.get);
	}

	ngOnDestroy(): void {
		this.subscriptions$.map((s) => s.unsubscribe());
		this.store.dispatch(new Clear());
	}

	ngOnInit(): void {
		this.store.dispatch(new FetchPackagingList());
	}

	addNew() {
		this.dialog.open(CreatePackagingComponent);
	}

	rotate() {
		if (this.boxDisplayIndex < this.boxDisplaySide.length -1) {
			this.boxDisplayIndex++;
		} else {
			this.boxDisplayIndex = 0;
		}
	}

	archive(event: Event, p: PackagingResponse) {
		event.stopPropagation()
		const ref = this.dialog.open(ArchiveConfirmationComponent);
		ref.componentRef!.instance.packagingName = `${p.name} - ${p.carrierBrand?.label}`
		ref.componentRef!.instance.id = p.id;
	}
}
