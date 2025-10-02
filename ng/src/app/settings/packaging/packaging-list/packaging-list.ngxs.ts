import {Injectable} from "@angular/core";
import {Action, Selector, State, StateContext, Store} from "@ngxs/store";
import SetPackagingList = PackagingListActions.SetPackagingList;
import {toNotNullArray} from "../../../functions/not-null-array";
import PackagingResponse = PackagingListActions.PackagingResponse;
import {AppActions} from "../../../app.actions";
import ShowGlobalSnackbar = AppActions.ShowGlobalSnackbar;
import {
	PackagingListActions,
} from "./packaging-list.actions";
import {ArchivePackagingGQL, CreatePackagingGQL, FetchPackagingListGQL} from "./packaging-list.generated";
import CarrierBrandResponse = PackagingListActions.CarrierBrandResponse;
import USPSRateIndicatorResponse = PackagingListActions.USPSRateIndicatorResponse;
import USPSProcessingCategoryResponse = PackagingListActions.USPSProcessingCategoryResponse;
import SetProcessingCategories = PackagingListActions.SetProcessingCategories;
import SetCarrierBrands = PackagingListActions.SetCarrierBrands;
import SetRateIndicators = PackagingListActions.SetRateIndicators;
import FetchPackagingList = PackagingListActions.FetchPackagingList;

export interface PackagingListModel {
	packagingList: PackagingResponse[];
	carrierBrands: CarrierBrandResponse[];
	rateIndicators: USPSRateIndicatorResponse[];
	processingCategories: USPSProcessingCategoryResponse[];
	loading: boolean;
}

const defaultState: PackagingListModel = {
	carrierBrands: [],
	processingCategories: [],
	rateIndicators: [],
	packagingList: [],
	loading: false,
};

@Injectable()
@State<PackagingListModel>({
	name: 'PackagingList',
	defaults: defaultState,
})
export class PackagingListState {

	constructor(
		private list: FetchPackagingListGQL,
		private create: CreatePackagingGQL,
		private archive: ArchivePackagingGQL,
	) {
	}

	@Selector()
	static get(state: PackagingListModel) {
		return state;
	}

	@Action(PackagingListActions.FetchPackagingList)
	FetchMyPackagingList(ctx: StateContext<PackagingListModel>, action: PackagingListActions.FetchPackagingList) {
		ctx.patchState({loading: true});
		return this.list.fetch({})
			.subscribe({next: (r) => {
					ctx.patchState({loading: false});
				const packaging = toNotNullArray(r.data.packagingFiltered);
				ctx.dispatch(new SetPackagingList(packaging));

				const car = toNotNullArray(r.data.carrierBrands.edges?.map((l) => l?.node));
				ctx.dispatch(new SetCarrierBrands(car));

				const rate = toNotNullArray(r.data.packagingUSPSRateIndicators.edges?.map((l) => l?.node));
				ctx.dispatch(new SetRateIndicators(rate));

				const cat = toNotNullArray(r.data.packagingUSPSProcessingCategories.edges?.map((l) => l?.node));
				ctx.dispatch(new SetProcessingCategories(cat));
			}});
	}

	@Action(PackagingListActions.SetCarrierBrands)
	SetCarrierBrands(ctx: StateContext<PackagingListModel>, action: PackagingListActions.SetCarrierBrands) {
		ctx.patchState({carrierBrands: action.payload});
	}

	@Action(PackagingListActions.SetRateIndicators)
	SetRateIndicators(ctx: StateContext<PackagingListModel>, action: PackagingListActions.SetRateIndicators) {
		ctx.patchState({rateIndicators: action.payload});
	}

	@Action(PackagingListActions.Archive)
	Archive(ctx: StateContext<PackagingListModel>, action: PackagingListActions.Archive) {
		ctx.patchState({loading: true});
		return this.archive.mutate({id: action.payload})
			.subscribe((resp) => {
				ctx.patchState({loading: false});
				if (!!resp.errors) {
					ctx.dispatch(new ShowGlobalSnackbar("Errors: " + JSON.stringify(resp.errors)));
				} else {
					ctx.dispatch(new FetchPackagingList());
				}
			});
	}

	@Action(PackagingListActions.SetProcessingCategories)
	SetProcessingCategories(ctx: StateContext<PackagingListModel>, action: PackagingListActions.SetProcessingCategories) {
		ctx.patchState({processingCategories: action.payload});
	}

	@Action(PackagingListActions.SetPackagingList)
	SetPackagingList(ctx: StateContext<PackagingListModel>, action: PackagingListActions.SetPackagingList) {
		ctx.patchState({packagingList: action.payload});
	}

	@Action(PackagingListActions.Clear)
	Clear(ctx: StateContext<PackagingListModel>, action: PackagingListActions.Clear) {
		ctx.setState(defaultState);
	}

	@Action(PackagingListActions.Create)
	Create(ctx: StateContext<PackagingListModel>, action: PackagingListActions.Create) {
		return this.create.mutate({
				input: action.payload.packaging,
				inputPackagingUSPS: action.payload.uspsPackaging,
				inputPackagingDF: action.payload.dfPackaging,
			},
			{errorPolicy: "all"})
			.subscribe((r) => {
				if (!!r.errors) {
					ctx.dispatch(new ShowGlobalSnackbar("An error occurred"));
				} else {
					const packaging = toNotNullArray(r.data?.createPackaging);
					ctx.dispatch(new SetPackagingList(packaging));
				}
			});
	}

}
