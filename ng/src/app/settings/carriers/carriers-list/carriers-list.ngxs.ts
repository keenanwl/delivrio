import {Injectable} from "@angular/core";
import {Action, Selector, State, StateContext} from "@ngxs/store";
import {CarriersListActions} from "./carriers-list.actions";
import SelectCarriersListQueryResponse = CarriersListActions.SelectCarriersListQueryResponse;
import SetCarriersList = CarriersListActions.SetCarriersList;
import SelectCarrierBrandsQueryResponse = CarriersListActions.SelectCarrierBrandsQueryResponse;
import {
	CarrierBrandsGQL,
	CreateCarrierAgreementGQL,
	ListCarriersGQL
} from "./carriers-list.generated";
import {toNotNullArray} from "../../../functions/not-null-array";
import {Paths} from "../../../app-routing.module";
import {CarrierBrandInternalId} from "../../../../generated/graphql";
import {AppActions} from "../../../app.actions";
import AppChangeRoute = AppActions.AppChangeRoute;

export interface CarriersListModel {
	carriersList: SelectCarriersListQueryResponse[];
	brands: SelectCarrierBrandsQueryResponse[];
	loading: boolean;
}

const defaultState: CarriersListModel = {
	carriersList: [],
	brands: [],
	loading: false,
};

@Injectable()
@State<CarriersListModel>({
	name: 'carriersList',
	defaults: defaultState,
})
export class CarriersListState {

	constructor(
		private createCarrier: CreateCarrierAgreementGQL,
		private list: ListCarriersGQL,
		private fetchBrands: CarrierBrandsGQL,
	) {
	}

	@Selector()
	static get(state: CarriersListModel) {
		return state;
	}

	@Action(CarriersListActions.FetchCarriersList)
	FetchCarriersList(ctx: StateContext<CarriersListModel>, action: CarriersListActions.FetchCarriersList) {
		ctx.patchState({loading: true});
		return this.list.fetch({})
			.subscribe({next: (r) => {
				ctx.patchState({loading: false});

				ctx.dispatch(new SetCarriersList(toNotNullArray(r.data.carriers.edges?.map((c) => c?.node))));
			}});
	}

	@Action(CarriersListActions.SetCarriersList)
	SetMyCarriersList(ctx: StateContext<CarriersListModel>, action: CarriersListActions.SetCarriersList) {
		ctx.patchState({carriersList: action.payload})
	}

	@Action(CarriersListActions.FetchCarrierBrands)
	FetchCarrierBrands(ctx: StateContext<CarriersListModel>, action: CarriersListActions.FetchCarrierBrands) {
		ctx.patchState({loading: true});
		return this.fetchBrands.fetch()
			.subscribe({next: (r) => {
				ctx.patchState({loading: false});
				const brands = toNotNullArray(r.data.carrierBrands.edges?.map((r) => r?.node));
				ctx.dispatch(new CarriersListActions.SetCarrierBrands(brands));
			}});
	}

	@Action(CarriersListActions.SetCarrierBrands)
	SetCarrierBrands(ctx: StateContext<CarriersListModel>, action: CarriersListActions.SetCarrierBrands) {
		ctx.patchState({brands: action.payload});
	}

	@Action(CarriersListActions.Clear)
	Clear(ctx: StateContext<CarriersListModel>, action: CarriersListActions.Clear) {
		ctx.setState(defaultState);
	}

	@Action(CarriersListActions.CreateNewAgreement)
	CreateNewAgreement(ctx: StateContext<CarriersListModel>, action: CarriersListActions.CreateNewAgreement) {
		ctx.patchState({loading: true});
		return this.createCarrier.mutate({
			name: action.payload.name,
			carrierBrand: action.payload.carrierBrandID,
		}).subscribe((r) => {
			ctx.patchState({loading: false});
			const resp = r.data?.createCarrierAgreement;
			if (!!resp) {

				let path = "";
				switch (resp.carrier) {
					case CarrierBrandInternalId.Bring:
						path = Paths.SETTINGS_CARRIERS_EDIT_BRING;
						break;
					case CarrierBrandInternalId.Dao:
						path = Paths.SETTINGS_CARRIERS_EDIT_DAO;
						break;
					case CarrierBrandInternalId.Df:
						path = Paths.SETTINGS_CARRIERS_EDIT_DF;
						break;
					case CarrierBrandInternalId.Dsv:
						path = Paths.SETTINGS_CARRIERS_EDIT_DSV;
						break;
					case CarrierBrandInternalId.EasyPost:
						path = Paths.SETTINGS_CARRIERS_EDIT_EASY_POST;
						break;
					case CarrierBrandInternalId.Gls:
						path = Paths.SETTINGS_CARRIERS_EDIT_GLS;
						break;
					case CarrierBrandInternalId.PostNord:
						path = Paths.SETTINGS_CARRIERS_EDIT_POST_NORD;
						break;
					case CarrierBrandInternalId.Usps:
						path = Paths.SETTINGS_CARRIERS_EDIT_USPS;
						break;
					default:
						throw new Error("carrier ID not recognized: " + resp.carrier);
				}

				ctx.dispatch(new AppChangeRoute({
					path,
					queryParams: {id: resp.id}
				}));
			}
		});
	}

}
