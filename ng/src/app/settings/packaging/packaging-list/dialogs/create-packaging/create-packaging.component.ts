import {Component, OnDestroy, OnInit} from '@angular/core';
import {Observable, Subscription} from "rxjs";
import {Store} from "@ngxs/store";
import {DialogRef} from "@angular/cdk/dialog";
import {PackagingListModel, PackagingListState} from '../../packaging-list.ngxs';
import {PackagingListActions} from "../../packaging-list.actions";
import FetchPackagingList = PackagingListActions.FetchPackagingList;
import Create = PackagingListActions.Create;
import {FormControl} from "@angular/forms";
import {
	CarrierBrandInternalId,
	CreatePackagingDfInput,
	CreatePackagingUspsInput, PackagingDfapiType
} from "../../../../../../generated/graphql";
import CarrierBrandResponse = PackagingListActions.CarrierBrandResponse;

@Component({
	selector: 'app-create-packaging',
	templateUrl: './create-packaging.component.html',
	styleUrls: ['./create-packaging.component.scss']
})
export class CreatePackagingComponent implements OnDestroy, OnInit {
	packaging$: Observable<PackagingListModel>;
	subscriptions$: Subscription[] = [];
	carrierBrands = CarrierBrandInternalId;

	dfAPITypes = PackagingDfapiType;

	// USPS
	rateCtrl = new FormControl("", {nonNullable: true});
	categoryCtrl = new FormControl("", {nonNullable: true});

	// DF
	dfAPITypeCtrl = new FormControl<PackagingDfapiType>(this.dfAPITypes.Pl1, {nonNullable: true});
	stackableCtrl = new FormControl(false, {nonNullable: true});

	constructor(
		private store: Store,
		private ref: DialogRef,
	) {
		this.packaging$ = store.select(PackagingListState.get);
	}

	ngOnDestroy(): void {
		this.subscriptions$.map((s) => s.unsubscribe());
		this.store.dispatch(new FetchPackagingList());
	}

	ngOnInit(): void {

	}

	create(name: string, height: string, width: string, length: string, brand: CarrierBrandResponse) {

		let uspsPackaging: CreatePackagingUspsInput | undefined = undefined;
		if (brand.internalID === this.carrierBrands.Usps) {
			uspsPackaging = {
				packagingUSPSProcessingCategoryID: this.categoryCtrl.value,
				packagingUSPSRateIndicatorID: this.rateCtrl.value,
			}
		}

		let dfPackaging: CreatePackagingDfInput | undefined = undefined;
		if (brand.internalID === this.carrierBrands.Df) {
			dfPackaging = {
				stackable: this.stackableCtrl.value,
				apiType: this.dfAPITypeCtrl.value,
			}
		}

		this.store.dispatch(new Create({
			packaging: {
				name,
				heightCm: parseInt(height),
				widthCm: parseInt(width),
				lengthCm: parseInt(length),
				carrierBrandID: brand.id === "" ? null : brand.id,
			},
			uspsPackaging,
			dfPackaging,
		}));
		this.close();
	}

	close() {
		this.ref.close();
	}
}
