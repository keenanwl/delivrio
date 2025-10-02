import {Injectable} from "@angular/core";
import {Action, Selector, State, StateContext, Store} from "@ngxs/store";
import SetHypothesisTestingDeliveryOptionEdit = HypothesisTestingDeliveryOptionEditActions.SetHypothesisTestingDeliveryOptionEdit;
import {produce} from "immer";
import SetHypothesisTestingDeliveryOptionTags = HypothesisTestingDeliveryOptionEditActions.SetHypothesisTestingDeliveryOptionTags;
import ShowGlobalSnackbar = AppActions.ShowGlobalSnackbar;
import AppChangeRoute = AppActions.AppChangeRoute;
import {AppActions} from "../../../../app.actions";
import {
	HypothesisTestingDeliveryOptionEditActions
} from "./hypothesis-testing-delivery-option-edit.actions";
import {formErrors} from "../../../../account/company-info/company-info.ngxs";
import {
	FetchHypothesisTestGQL,
	UpdateHypothesisTestDeliveryOptionGQL
} from "./hypothesis-testing-delivery-option-edit.generated";
import HTResponse = HypothesisTestingDeliveryOptionEditActions.HTResponse;
import HTDeliveryOptionResponse = HypothesisTestingDeliveryOptionEditActions.HTDeliveryOptionResponse;
import SetAvailableDeliveryOptions = HypothesisTestingDeliveryOptionEditActions.SetAvailableDeliveryOptions;
import {toNotNullArray} from "../../../../functions/not-null-array";
import {Paths} from "../../../../app-routing.module";

export interface HypothesisTestingDeliveryOptionEditModel {
	editForm: {
		model: HTResponse | undefined;
		dirty: boolean;
		status: string;
		errors: formErrors;
	},
	hypothesisTestingDeliveryOptionID: string;
	availableDeliveryOptions: HTDeliveryOptionResponse[];
}

const defaultState: HypothesisTestingDeliveryOptionEditModel = {
	editForm: {
		model: undefined,
		dirty: false,
		status: '',
		errors: {}
	},
	hypothesisTestingDeliveryOptionID: '',
	availableDeliveryOptions: [],
};

@Injectable()
@State<HypothesisTestingDeliveryOptionEditModel>({
	name: 'hypothesisTestingDeliveryOptionEdit',
	defaults: defaultState,
})
export class HypothesisTestingDeliveryOptionEditState {

	constructor(
		private fetch: FetchHypothesisTestGQL,
		private update: UpdateHypothesisTestDeliveryOptionGQL,
		private store: Store,
	) {}

	@Selector()
	static get(state: HypothesisTestingDeliveryOptionEditModel) {
		return state;
	}

	@Action(HypothesisTestingDeliveryOptionEditActions.Fetch)
	FetchMyHypothesisTestingDeliveryOptionEdit(ctx: StateContext<HypothesisTestingDeliveryOptionEditModel>, action: HypothesisTestingDeliveryOptionEditActions.Fetch) {
		const id = ctx.getState().hypothesisTestingDeliveryOptionID;
		return this.fetch.fetch({id})
			.subscribe({next: (r) => {
				const availableDeliveryOptions = r.data.unassignedDeliveryOptions;
				if (!!availableDeliveryOptions) {
					ctx.dispatch(new SetAvailableDeliveryOptions(availableDeliveryOptions));
				}

				const hypothesisTestingDeliveryOption = r.data.hypothesisTest;
				if (!!hypothesisTestingDeliveryOption) {
					ctx.dispatch(new SetHypothesisTestingDeliveryOptionEdit(hypothesisTestingDeliveryOption));
				}
			}});
	}

	@Action(HypothesisTestingDeliveryOptionEditActions.SetHypothesisTestingDeliveryOptionID)
	SetHypothesisTestingDeliveryOptionID(ctx: StateContext<HypothesisTestingDeliveryOptionEditModel>, action: HypothesisTestingDeliveryOptionEditActions.SetHypothesisTestingDeliveryOptionID) {
		ctx.patchState({hypothesisTestingDeliveryOptionID: action.payload});
	}

	@Action(HypothesisTestingDeliveryOptionEditActions.SetHypothesisTestingDeliveryOptionEdit)
	SetHypothesisTestingDeliveryOptionEdit(ctx: StateContext<HypothesisTestingDeliveryOptionEditModel>, action: HypothesisTestingDeliveryOptionEditActions.SetHypothesisTestingDeliveryOptionEdit) {
		const state = produce(ctx.getState(), st => {
			st.editForm.model = action.payload;

		});
		ctx.setState(state);
	}

	@Action(HypothesisTestingDeliveryOptionEditActions.Clear)
	Clear(ctx: StateContext<HypothesisTestingDeliveryOptionEditModel>, action: HypothesisTestingDeliveryOptionEditActions.Clear) {
		ctx.setState(defaultState);
	}

	@Action(HypothesisTestingDeliveryOptionEditActions.SetAvailableDeliveryOptions)
	SetAvailableDeliveryOptions(ctx: StateContext<HypothesisTestingDeliveryOptionEditModel>, action: HypothesisTestingDeliveryOptionEditActions.SetAvailableDeliveryOptions) {
		ctx.patchState({availableDeliveryOptions: action.payload});
	}

	@Action(HypothesisTestingDeliveryOptionEditActions.MoveDeliveryOption)
	MoveDeliveryOption(ctx: StateContext<HypothesisTestingDeliveryOptionEditModel>, action: HypothesisTestingDeliveryOptionEditActions.MoveDeliveryOption) {
		const state = produce(ctx.getState(), st => {
			const next = st.editForm.model;
			st.editForm.model!.hypothesisTestDeliveryOption!.deliveryOptionGroupOne
				= next?.hypothesisTestDeliveryOption?.deliveryOptionGroupOne?.filter((i) => i.id !== action.payload.deliveryOption.id);
			st.editForm.model!.hypothesisTestDeliveryOption!.deliveryOptionGroupTwo
				= next?.hypothesisTestDeliveryOption?.deliveryOptionGroupTwo?.filter((i) => i.id !== action.payload.deliveryOption.id);

			st.availableDeliveryOptions
				= st.availableDeliveryOptions.filter((i) => i.id !== action.payload.deliveryOption.id);

			switch (action.payload.container) {
				case "available":
					st.availableDeliveryOptions.push(action.payload.deliveryOption);
					break;
				case "control":
					st.editForm.model!.hypothesisTestDeliveryOption!.deliveryOptionGroupOne!.push(action.payload.deliveryOption);
					break;
				case "test":
					st.editForm.model!.hypothesisTestDeliveryOption!.deliveryOptionGroupTwo!.push(action.payload.deliveryOption);
					break;
			}

		});
		ctx.dispatch([
			new SetHypothesisTestingDeliveryOptionEdit(state.editForm.model),
			new SetAvailableDeliveryOptions(state.availableDeliveryOptions),
		]);
	}

	@Action(HypothesisTestingDeliveryOptionEditActions.Save)
	Save(ctx: StateContext<HypothesisTestingDeliveryOptionEditModel>, action: HypothesisTestingDeliveryOptionEditActions.Save) {
		const state = ctx.getState();
		return this.update.mutate({id: state.hypothesisTestingDeliveryOptionID,
			input: {
				name: state.editForm.model?.name,
				active: state.editForm.model?.active,
			},
			inputDeliveryOption: {
				clearDeliveryOptionGroupOne: true,
				addDeliveryOptionGroupOneIDs: state.editForm.model?.hypothesisTestDeliveryOption?.deliveryOptionGroupOne?.map((i) => i.id),
				clearDeliveryOptionGroupTwo: true,
				addDeliveryOptionGroupTwoIDs: state.editForm.model?.hypothesisTestDeliveryOption?.deliveryOptionGroupTwo?.map((i) => i.id),
				byOrder: state.editForm.model?.hypothesisTestDeliveryOption?.byOrder,
				byIntervalRotation: state.editForm.model?.hypothesisTestDeliveryOption?.byIntervalRotation,
				rotationIntervalHours: state.editForm.model?.hypothesisTestDeliveryOption?.rotationIntervalHours,
			},
		}).subscribe((r) => {
			if (!!r.errors) {
				ctx.dispatch(new ShowGlobalSnackbar("An error occurred"));
			} else {
				ctx.dispatch(new AppChangeRoute({path: Paths.SETTINGS_HYPOTHESIS_TESTING, queryParams: {}}))
			}
		});
	}

}
