import {Injectable} from "@angular/core";
import {Action, Selector, State, StateContext} from "@ngxs/store";
import {SelectPackagingActions} from "./select-packaging.actions";
import {FetchPackagingGQL} from "./select-packaging.generated";
import PackagingResponse = SelectPackagingActions.PackagingResponse;
import SetPackaging = SelectPackagingActions.SetPackaging;
import {toNotNullArray} from "../../functions/not-null-array";

export interface SelectPackagingModel {
	loading: boolean;
	packaging: PackagingResponse[];
}

const defaultState: SelectPackagingModel = {
	loading: false,
	packaging: [],
};

@Injectable()
@State<SelectPackagingModel>({
	name: 'selectPackaging',
	defaults: defaultState,
})
export class SelectPackagingState {

	constructor(
		private fetchPackaging: FetchPackagingGQL,
	) {}

	@Selector()
	static get(state: SelectPackagingModel) {
		return state;
	}

	@Action(SelectPackagingActions.FetchPackaging)
	FetchPackaging(ctx: StateContext<SelectPackagingModel>, action: SelectPackagingActions.FetchPackaging) {
		ctx.patchState({loading: true});
		return this.fetchPackaging.fetch()
			.subscribe((r) => {
				const packaging = toNotNullArray(r.data.packagings.edges?.map((n) => n?.node));
				ctx.dispatch(new SetPackaging(packaging));
		});
	}

	@Action(SelectPackagingActions.SetPackaging)
	SetPackaging(ctx: StateContext<SelectPackagingModel>, action: SelectPackagingActions.SetPackaging) {
		ctx.patchState({
			loading: false,
			packaging: action.payload
		});
	}

}
