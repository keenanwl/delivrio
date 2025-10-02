import {Injectable} from "@angular/core";
import {Action, Selector, State, StateContext} from "@ngxs/store";
import {formErrors} from "../../../../account/company-info/company-info.ngxs";
import {
	CreateDeliveryRuleConstraintInput,
	DeliveryRuleConstraintComparison,
	DeliveryRuleConstraintGroupConstraintLogic,
	DeliveryRuleConstraintPropertyType
} from "src/generated/graphql";
import {produce} from "immer";
import FetchRuleConstraintsQueryResponse = DeliveryOptionEditRulesActions.FetchRuleConstraintsQueryResponse;
import SearchProductTagsResponse = DeliveryOptionEditRulesActions.SearchProductTagsResponse;
import CountriesResponse = DeliveryOptionEditRulesActions.CountriesResponse;
import SelectDeliveryOptionEditRulesRulesQueryResponse = DeliveryOptionEditRulesActions.SelectDeliveryOptionEditRulesRulesQueryResponse;
import {
	DeliveryOptionEditRulesActions
} from "./delivery-option-edit-rules.actions";
import {
	CreateConstraintGroupConstraintsGQL,
	CreateDeliveryRuleGQL, DeleteConstraintGroupConstraintsGQL, DeleteRuleGQL, DeliveryOptionsSearchCountriesGQL,
	FetchDeliveryOptionRulesGQL,
	FetchRuleConstraintsGQL,
	ReplaceConstraintGroupConstraintsGQL, ReplaceRuleCountriesGQL, SearchProductTagsGQL, UpdateDeliveryRulePriceGQL
} from "./delivery-option-edit-rules.generated";
import SetDeliveryOptionEditRulesRulesEdit = DeliveryOptionEditRulesActions.SetDeliveryOptionEditRulesRulesEdit;
import {toNotNullArray} from "src/app/functions/not-null-array";
import SetRuleConstraints = DeliveryOptionEditRulesActions.SetRuleConstraints;
import SetCurrencies = DeliveryOptionEditRulesActions.SetCurrencies;
import CurrencyResponse = DeliveryOptionEditRulesActions.CurrencyResponse;
import {state} from "@angular/animations";
import {AppActions} from "../../../../app.actions";
import ShowGlobalSnackbar = AppActions.ShowGlobalSnackbar;

export interface DeliveryOptionEditRulesModel {
	rules: SelectDeliveryOptionEditRulesRulesQueryResponse[];
	constraintLogicType: DeliveryRuleConstraintGroupConstraintLogic;
	constraints: FetchRuleConstraintsQueryResponse;
	rulesForm: {
		model: {rules: SelectDeliveryOptionEditRulesRulesQueryResponse[]} | undefined;
		dirty: boolean;
		status: string;
		errors: formErrors;
	},
	deliveryOptionID: string;
	selectedRule: string;
	selectedRuleIndex: number;
	selectedConstraintGroup: string;
	selectedConstraintGroupIndex: number;
	searchProductTags: SearchProductTagsResponse[],
	searchProductTagsTerm: string,
	searchCountries: CountriesResponse[],
	currencies: CurrencyResponse[],
}

const defaultState: DeliveryOptionEditRulesModel = {
	rules: [],
	constraintLogicType: DeliveryRuleConstraintGroupConstraintLogic.And,
	constraints: [],
	rulesForm: {
		model: undefined,
		dirty: false,
		status: '',
		errors: {}
	},
	deliveryOptionID: '',
	selectedRule: '',
	selectedRuleIndex: 0,
	selectedConstraintGroup: '',
	selectedConstraintGroupIndex: 0,
	searchProductTags: [],
	searchProductTagsTerm: "",
	searchCountries: [],
	currencies: [],
};

@Injectable()
@State<DeliveryOptionEditRulesModel>({
	name: 'deliveryOptionEditRules',
	defaults: defaultState,
})
export class DeliveryOptionEditRulesState {

	constructor(
		private fetchRules: FetchDeliveryOptionRulesGQL,
		private fetchConstraints: FetchRuleConstraintsGQL,
		private createRules: CreateDeliveryRuleGQL,
		private createConstraints: CreateConstraintGroupConstraintsGQL,
		private replaceConstraints: ReplaceConstraintGroupConstraintsGQL,
		private deleteRule: DeleteRuleGQL,
		private deleteConstraintGroup: DeleteConstraintGroupConstraintsGQL,
		private tagSearch: SearchProductTagsGQL,
		private countrySearch: DeliveryOptionsSearchCountriesGQL,
		private replaceCountries: ReplaceRuleCountriesGQL,
		private updateRulePrice: UpdateDeliveryRulePriceGQL,
	) {}

	@Selector()
	static get(state: DeliveryOptionEditRulesModel) {
		return state;
	}

	@Action(DeliveryOptionEditRulesActions.SetDeliveryOptionEditRules)
	SetDeliveryOptionEditRules(ctx: StateContext<DeliveryOptionEditRulesModel>, action: DeliveryOptionEditRulesActions.SetDeliveryOptionEditRules) {
/*		const state = ctx.getState();
		const next = Object.assign({}, state.DeliveryOptionEditRulesForm, {model: action.payload});
		ctx.patchState({
			DeliveryOptionEditRulesForm: next,
		})*/
	}

	@Action(DeliveryOptionEditRulesActions.CreateDeliveryRule)
	CreateDeliveryRule(ctx: StateContext<DeliveryOptionEditRulesModel>, action: DeliveryOptionEditRulesActions.CreateDeliveryRule) {
		const state = ctx.getState();
		const optionId = state.deliveryOptionID;
		return this.createRules.mutate({input: {name: action.payload, deliveryOptionID: optionId}})
			.subscribe((r) => {
				const state2 = ctx.getState();
				ctx.dispatch(new DeliveryOptionEditRulesActions.FetchDeliveryOptionEditRulesRuleEdit({selectedIndex: state2.selectedRuleIndex}));
			});
	}

	@Action(DeliveryOptionEditRulesActions.SetRuleList)
	SetRuleList(ctx: StateContext<DeliveryOptionEditRulesModel>, action: DeliveryOptionEditRulesActions.SetRuleList) {
		ctx.patchState({
			rules: action.payload,
		})
	}

	@Action(DeliveryOptionEditRulesActions.FetchDeliveryOptionEditRulesRuleEdit)
	FetchDeliveryOptionEditRulesRuleEdit(ctx: StateContext<DeliveryOptionEditRulesModel>, action: DeliveryOptionEditRulesActions.FetchDeliveryOptionEditRulesRuleEdit) {
		const state = ctx.getState();
		return this.fetchRules.fetch({id: state.deliveryOptionID}, {fetchPolicy: "no-cache"})
			.subscribe({next: (r) => {
					const c = toNotNullArray(r.data.currencies.edges?.map((c) => c?.node));
					ctx.dispatch(new SetCurrencies(c));

					const d = toNotNullArray(r.data.deliveryRules.edges?.map((n) => n?.node));
					ctx.dispatch(new SetDeliveryOptionEditRulesRulesEdit(d));
					ctx.dispatch(new DeliveryOptionEditRulesActions.SetRuleList(d));
					const ruleID = d[action.payload.selectedIndex]?.id
					if (!!ruleID) {
						ctx.dispatch(new DeliveryOptionEditRulesActions.SetSelectedRule({ruleID, ruleIndex: action.payload.selectedIndex}))
					}
				}});
	}

	@Action(DeliveryOptionEditRulesActions.FetchRuleConstrains)
	FetchRuleConstrains(ctx: StateContext<DeliveryOptionEditRulesModel>, action: DeliveryOptionEditRulesActions.FetchRuleConstrains) {
		const state = ctx.getState();
		const id = state.selectedConstraintGroup;
		if (id.length > 0) {
			return this.fetchConstraints.fetch({id})
				.subscribe((r) => {

					const logic = r.data.constraintGroup?.constraintLogic;
					if (!!logic) {
						ctx.dispatch([
							new DeliveryOptionEditRulesActions.SetConstraintLogicType(logic),
						]);
					}

					const resp = r.data.constraints;
					if (!!resp) {
						ctx.dispatch([
							new DeliveryOptionEditRulesActions.SetRuleConstraintsNotify(resp),
						]);
					}
				});
		}

	}

	@Action(DeliveryOptionEditRulesActions.SetRuleConstraintsNotify)
	SetRuleConstraintsNotify(ctx: StateContext<DeliveryOptionEditRulesModel>, action: DeliveryOptionEditRulesActions.SetRuleConstraintsNotify) {
		ctx.dispatch(new SetRuleConstraints(action.payload));
	}

	@Action(DeliveryOptionEditRulesActions.SetRuleConstraints)
	SetRuleConstraints(ctx: StateContext<DeliveryOptionEditRulesModel>, action: DeliveryOptionEditRulesActions.SetRuleConstraints) {
		ctx.patchState({constraints: action.payload})
	}

	@Action(DeliveryOptionEditRulesActions.SetCurrencies)
	SetCurrencies(ctx: StateContext<DeliveryOptionEditRulesModel>, action: DeliveryOptionEditRulesActions.SetCurrencies) {
		ctx.patchState({currencies: action.payload});
	}

	@Action(DeliveryOptionEditRulesActions.SetSelectedConstraintID)
	SetSelectedConstraintID(ctx: StateContext<DeliveryOptionEditRulesModel>, action: DeliveryOptionEditRulesActions.SetSelectedConstraintID) {
		ctx.patchState({
			selectedConstraintGroup: action.payload.id,
			selectedConstraintGroupIndex: action.payload.index,
		})
	}

	@Action(DeliveryOptionEditRulesActions.SetSelectedRule)
	SetSelectedRule(ctx: StateContext<DeliveryOptionEditRulesModel>, action: DeliveryOptionEditRulesActions.SetSelectedRule) {
		ctx.patchState({
			selectedRule: action.payload.ruleID,
			selectedRuleIndex: action.payload.ruleIndex,
		})
	}

	@Action(DeliveryOptionEditRulesActions.SetSelectedOption)
	SetSelectedOption(ctx: StateContext<DeliveryOptionEditRulesModel>, action: DeliveryOptionEditRulesActions.SetSelectedOption) {
		ctx.patchState({deliveryOptionID: action.payload})
	}

	@Action(DeliveryOptionEditRulesActions.UpdateRulePricing)
	UpdateRulePricing(ctx: StateContext<DeliveryOptionEditRulesModel>, action: DeliveryOptionEditRulesActions.UpdateRulePricing) {
		const state = ctx.getState();
		return this.updateRulePrice.mutate({
			val: {price: action.payload.price, currencyID: action.payload.currency.id},
			deliveryRuleID: state.rules[state.selectedRuleIndex]?.id || "",
		}, {errorPolicy: "all"}).subscribe((r) => {
			if (!!r.errors) {
				ctx.dispatch(new ShowGlobalSnackbar("Errors: " + JSON.stringify(r.errors)));
			} else {
				ctx.dispatch(new DeliveryOptionEditRulesActions.FetchDeliveryOptionEditRulesRuleEdit({
					selectedIndex: state.selectedRuleIndex
				}));
				ctx.dispatch(new ShowGlobalSnackbar("Success: price updated"));
			}
		});
	}

	@Action(DeliveryOptionEditRulesActions.Clear)
	Clear(ctx: StateContext<DeliveryOptionEditRulesModel>, action: DeliveryOptionEditRulesActions.Clear) {
		ctx.setState(defaultState);
	}

	@Action(DeliveryOptionEditRulesActions.SaveForm)
	SaveForm(ctx: StateContext<DeliveryOptionEditRulesModel>, action: DeliveryOptionEditRulesActions.SaveForm) {
/*		return this.updateGLS.mutate(action.payload, {errorPolicy: "all"})
			.subscribe((res) => {
				if (!!res.errors) {
					ctx.dispatch([
						new SetFormErrors({errors: res.errors, formPath: 'DeliveryOptionEditRules.DeliveryOptionEditRulesForm'}),
						new SetFormErrors({errors: res.errors, formPath: 'DeliveryOptionEditRules.deliveryOptionPriceForm'}),
						new ShowGlobalSnackbar("Please fix the highlighted errors and try saving again"),
					]);
				} else {
					ctx.dispatch([
						new ShowGlobalSnackbar("Delivery options saved successfully"),
					]);
				}
			});*/
	}

	@Action(DeliveryOptionEditRulesActions.SaveNewConstraintGroup)
	SaveNewConstraintGroup(ctx: StateContext<DeliveryOptionEditRulesModel>, action: DeliveryOptionEditRulesActions.SaveNewConstraintGroup) {
		const state1 = ctx.getState();
		const input = state1.constraints.map((c) => c!.constraint).filter((x): x is CreateDeliveryRuleConstraintInput => x !== null);
		if (!!input) {
			return this.createConstraints.mutate({ruleID: state1.selectedRule, logicType: state1.constraintLogicType, input})
				.subscribe((r) => {
					const state = ctx.getState();
					ctx.dispatch(
						new DeliveryOptionEditRulesActions.FetchDeliveryOptionEditRulesRuleEdit({
							selectedIndex: state.selectedRuleIndex
						}))
				});
		}
	}

	@Action(DeliveryOptionEditRulesActions.SaveEditConstraintGroup)
	SaveEditConstraintGroup(ctx: StateContext<DeliveryOptionEditRulesModel>, action: DeliveryOptionEditRulesActions.SaveEditConstraintGroup) {
		const state1 = ctx.getState();
		const input = state1.constraints.map((c) => c!.constraint).filter((x): x is CreateDeliveryRuleConstraintInput => x !== null);
		if (!!input) {
			return this.replaceConstraints.mutate({
				deliveryGroupId: state1.selectedConstraintGroup, logicType: state1.constraintLogicType, input})
				.subscribe((r) => {
					const state = ctx.getState();
					ctx.dispatch(
						new DeliveryOptionEditRulesActions.FetchDeliveryOptionEditRulesRuleEdit({
							selectedIndex: state.selectedRuleIndex
						}))
				});
		}

	}

	@Action(DeliveryOptionEditRulesActions.DeleteConstraintGroup)
	DeleteConstraintGroup(ctx: StateContext<DeliveryOptionEditRulesModel>, action: DeliveryOptionEditRulesActions.DeleteConstraintGroup) {
		return this.deleteConstraintGroup.mutate({groupID: action.payload})
			.subscribe((r) => {
				const state = ctx.getState();
				ctx.dispatch(
					new DeliveryOptionEditRulesActions.FetchDeliveryOptionEditRulesRuleEdit({
						selectedIndex: state.selectedRuleIndex
					}))
			});
	}

	@Action(DeliveryOptionEditRulesActions.UpdateDayOfWeek)
	UpdateDayOfWeek(ctx: StateContext<DeliveryOptionEditRulesModel>, action: DeliveryOptionEditRulesActions.UpdateDayOfWeek) {
		const state = produce(ctx.getState(), st => {
			st.constraints![action.payload.index].constraint!.selectedValue.dayOfWeek = action.payload.dayOfWeek;
		});
		ctx.dispatch(new DeliveryOptionEditRulesActions.SetRuleConstraints(state.constraints));
	}

	@Action(DeliveryOptionEditRulesActions.UpdateTimeOfDay)
	UpdateTimeOfDay(ctx: StateContext<DeliveryOptionEditRulesModel>, action: DeliveryOptionEditRulesActions.UpdateTimeOfDay) {
		const state = produce(ctx.getState(), st => {
			st.constraints![action.payload.index].constraint!.selectedValue.timeOfDay = action.payload.timeOfDay;
		});
		ctx.dispatch(new DeliveryOptionEditRulesActions.SetRuleConstraints(state.constraints));
	}

	@Action(DeliveryOptionEditRulesActions.UpdateConstraintProperty)
	UpdateConstraintProperty(ctx: StateContext<DeliveryOptionEditRulesModel>, action: DeliveryOptionEditRulesActions.UpdateConstraintProperty) {
		const state = produce(ctx.getState(), st => {
			st.constraints![action.payload.index].constraint!.propertyType = action.payload.propertyType;
		});
		ctx.dispatch(new DeliveryOptionEditRulesActions.SetRuleConstraints(state.constraints));
	}

	@Action(DeliveryOptionEditRulesActions.UpdateConstraintComparison)
	UpdateConstraintComparison(ctx: StateContext<DeliveryOptionEditRulesModel>, action: DeliveryOptionEditRulesActions.UpdateConstraintComparison) {
		const state = produce(ctx.getState(), st => {
			st.constraints![action.payload.index].constraint!.comparison = action.payload.comparison;
		});
		ctx.dispatch(new DeliveryOptionEditRulesActions.SetRuleConstraints(state.constraints));
	}

	@Action(DeliveryOptionEditRulesActions.AddConstraintLine)
	AddConstraintLine(ctx: StateContext<DeliveryOptionEditRulesModel>, action: DeliveryOptionEditRulesActions.AddConstraintLine) {
		const state = produce(ctx.getState(), st => {
			st.constraints.push({
				constraint: {
					comparison: DeliveryRuleConstraintComparison.GreaterThan,
					propertyType: DeliveryRuleConstraintPropertyType.CartTotal,
					selectedValue: {numeric: 0},
				},
				tags: [],
			});
		});
		ctx.dispatch(new DeliveryOptionEditRulesActions.SetRuleConstraintsNotify(state.constraints));
	}

	@Action(DeliveryOptionEditRulesActions.DeleteConstraintLine)
	DeleteConstraintLine(ctx: StateContext<DeliveryOptionEditRulesModel>, action: DeliveryOptionEditRulesActions.DeleteConstraintLine) {
		const state = produce(ctx.getState(), st => {
			st.constraints = st.constraints!.filter((r, i) => i !== action.payload.index);
		});
		ctx.dispatch(new DeliveryOptionEditRulesActions.SetRuleConstraintsNotify(state.constraints));
	}

	@Action(DeliveryOptionEditRulesActions.SetProductTagSearchTerm)
	SetProductTagSearchTerm(ctx: StateContext<DeliveryOptionEditRulesModel>, action: DeliveryOptionEditRulesActions.SetProductTagSearchTerm) {
		ctx.patchState({searchProductTagsTerm: action.payload});
	}

	@Action(DeliveryOptionEditRulesActions.FetchProductTags)
	FetchProductTags(ctx: StateContext<DeliveryOptionEditRulesModel>, action: DeliveryOptionEditRulesActions.FetchProductTags) {
		return this.tagSearch.fetch({term: ctx.getState().searchProductTagsTerm})
			.subscribe((r) => {
				const s = r.data.productTags.edges?.map((g) => g?.node);
				ctx.dispatch(new DeliveryOptionEditRulesActions.SetProductTagSearch(toNotNullArray(s)));
			})
	}

	@Action(DeliveryOptionEditRulesActions.SetProductTagSearch)
	SetProductTagSearch(ctx: StateContext<DeliveryOptionEditRulesModel>, action: DeliveryOptionEditRulesActions.SetProductTagSearch) {
		ctx.patchState({searchProductTags: action.payload});
	}

	@Action(DeliveryOptionEditRulesActions.AddProductTag)
	AddProductTag(ctx: StateContext<DeliveryOptionEditRulesModel>, action: DeliveryOptionEditRulesActions.AddProductTag) {
		const state = produce(ctx.getState(), st => {
			if (!!st.constraints![action.payload.index]!.constraint!.selectedValue?.ids) {
				st.constraints![action.payload.index]!.constraint!.selectedValue?.ids?.push(action.payload.tag.id);
			} else {
				st.constraints![action.payload.index]!.constraint!.selectedValue!.ids = [action.payload.tag.id];
			}

			if (!!st.constraints![action.payload.index]?.tags) {
				st.constraints![action.payload.index]!.tags!.push(action.payload.tag);
			} else {
				st.constraints![action.payload.index]!.tags = [action.payload.tag];
			}
		});
		ctx.dispatch(new DeliveryOptionEditRulesActions.SetRuleConstraints(state.constraints));
	}

	@Action(DeliveryOptionEditRulesActions.RemoveProductTag)
	RemoveProductTag(ctx: StateContext<DeliveryOptionEditRulesModel>, action: DeliveryOptionEditRulesActions.RemoveProductTag) {

		const state = produce(ctx.getState(), st => {
			st.constraints![action.payload.index]!.constraint!.selectedValue.ids =
				st.constraints![action.payload.index]!.constraint!.selectedValue?.ids?.filter((id) => id !== action.payload.tag.id);
			st.constraints![action.payload.index]!.tags =
				st.constraints![action.payload.index]!.tags?.filter((t) => t.id !== action.payload.tag.id);
		});
		ctx.dispatch(new DeliveryOptionEditRulesActions.SetRuleConstraints(state.constraints));
	}

	@Action(DeliveryOptionEditRulesActions.UpdateNumericValue)
	UpdateNumericValue(ctx: StateContext<DeliveryOptionEditRulesModel>, action: DeliveryOptionEditRulesActions.UpdateNumericValue) {
		const state = produce(ctx.getState(), st => {
			st.constraints![action.payload.index]!.constraint!.selectedValue!.numeric = action.payload.value;
		});
		ctx.dispatch(new DeliveryOptionEditRulesActions.SetRuleConstraints(state.constraints));
	}

	@Action(DeliveryOptionEditRulesActions.RemovePostalCode)
	RemovePostalCode(ctx: StateContext<DeliveryOptionEditRulesModel>, action: DeliveryOptionEditRulesActions.RemovePostalCode) {
		const state = produce(ctx.getState(), st => {
			st.constraints![action.payload.index]!.constraint!.selectedValue!.numericRange =
				st.constraints![action.payload.index]!.constraint!.selectedValue!.numericRange!.filter((r) => r !== action.payload.code)
		});
		ctx.dispatch(new DeliveryOptionEditRulesActions.SetRuleConstraints(state.constraints));
	}

	@Action(DeliveryOptionEditRulesActions.RemovePostalCodeString)
	RemovePostalCodeString(ctx: StateContext<DeliveryOptionEditRulesModel>, action: DeliveryOptionEditRulesActions.RemovePostalCodeString) {
		const state = produce(ctx.getState(), st => {
			st.constraints![action.payload.index]!.constraint!.selectedValue!.values =
				st.constraints![action.payload.index]!.constraint!.selectedValue!.values!.filter((r) => r !== action.payload.code)
		});
		ctx.dispatch(new DeliveryOptionEditRulesActions.SetRuleConstraints(state.constraints));
	}

	@Action(DeliveryOptionEditRulesActions.AddPostalCode)
	AddZipCode(ctx: StateContext<DeliveryOptionEditRulesModel>, action: DeliveryOptionEditRulesActions.AddPostalCode) {
		const state = produce(ctx.getState(), st => {
			if (!!st.constraints![action.payload.index].constraint?.selectedValue.numericRange) {
				st.constraints![action.payload.index]!.constraint!.selectedValue!.numericRange!.push(action.payload.code);
			} else {
				st.constraints![action.payload.index]!.constraint!.selectedValue = {numericRange: [action.payload.code]};
			}
		});
		ctx.dispatch(new DeliveryOptionEditRulesActions.SetRuleConstraints(state.constraints));
	}

	@Action(DeliveryOptionEditRulesActions.AddPostalCodeString)
	AddPostalCodeString(ctx: StateContext<DeliveryOptionEditRulesModel>, action: DeliveryOptionEditRulesActions.AddPostalCodeString) {
		const state = produce(ctx.getState(), st => {
			if (!!st.constraints![action.payload.index].constraint?.selectedValue.values) {
				st.constraints![action.payload.index]!.constraint!.selectedValue!.values!.push(action.payload.code);
			} else {
				st.constraints![action.payload.index]!.constraint!.selectedValue = {values: [action.payload.code]};
			}
		});
		ctx.dispatch(new DeliveryOptionEditRulesActions.SetRuleConstraints(state.constraints));
	}

	@Action(DeliveryOptionEditRulesActions.SetConstraintLogicType)
	SetConstraintLogicType(ctx: StateContext<DeliveryOptionEditRulesModel>, action: DeliveryOptionEditRulesActions.SetConstraintLogicType) {
		ctx.patchState({constraintLogicType: action.payload});
	}

	@Action(DeliveryOptionEditRulesActions.UpdateConstraintLogicType)
	UpdateConstraintLogicType(ctx: StateContext<DeliveryOptionEditRulesModel>, action: DeliveryOptionEditRulesActions.UpdateConstraintLogicType) {
		ctx.patchState({constraintLogicType: action.payload});
	}

	@Action(DeliveryOptionEditRulesActions.AddCountry)
	AddCountry(ctx: StateContext<DeliveryOptionEditRulesModel>, action: DeliveryOptionEditRulesActions.AddCountry) {
		const state = ctx.getState();
		const ruleID = state.rules[state.selectedRuleIndex]!.id;
		const countries = (state.rules[state.selectedRuleIndex]!.country || []).map((c) => c.id);
		countries.push(action.payload.country.id);
		return this.replaceCountries.mutate({ruleID, countries})
			.subscribe((r) => {
				const state = produce(ctx.getState(), st => {
					const nextCountries = r.data?.replaceDeliveryRuleCountries.country;
					if (!!nextCountries) {
						st.rules[st.selectedRuleIndex]!.country = nextCountries;
					}
				});
				ctx.dispatch([
					new DeliveryOptionEditRulesActions.SetRuleList(state.rules),
					new ShowGlobalSnackbar("Success: country added"),
				]);
			});
	}

	@Action(DeliveryOptionEditRulesActions.RemoveCountry)
	RemoveCountry(ctx: StateContext<DeliveryOptionEditRulesModel>, action: DeliveryOptionEditRulesActions.RemoveCountry) {
		const state = ctx.getState();
		const ruleID = state.rules[state.selectedRuleIndex]!.id;
		const countries = (state.rules[state.selectedRuleIndex]!.country || [])
			.filter((c) => c.id !== action.payload.id)
			.map((c) => c.id);

		return this.replaceCountries.mutate({ruleID, countries})
			.subscribe((r) => {
				const state = produce(ctx.getState(), st => {
					const nextCountries = r.data?.replaceDeliveryRuleCountries.country;
					if (!!nextCountries) {
						st.rules[st.selectedRuleIndex]!.country = nextCountries;
					}
				});
				ctx.dispatch([
					new DeliveryOptionEditRulesActions.SetRuleList(state.rules),
					new ShowGlobalSnackbar("Success: country removed"),
				]);
			});
	}

	@Action(DeliveryOptionEditRulesActions.SearchCountry)
	SearchCountry(ctx: StateContext<DeliveryOptionEditRulesModel>, action: DeliveryOptionEditRulesActions.SearchCountry) {
		return this.countrySearch.fetch({term: action.payload})
			.subscribe((res) => {
				const countries = toNotNullArray(res.data.countries.edges?.map((value) => value?.node));
				if (!!countries) {
					ctx.dispatch(new DeliveryOptionEditRulesActions.SetCountrySearch(countries));
				}
			});
	}

	@Action(DeliveryOptionEditRulesActions.SetCountrySearch)
	SetCountrySearch(ctx: StateContext<DeliveryOptionEditRulesModel>, action: DeliveryOptionEditRulesActions.SetCountrySearch) {
		ctx.patchState({searchCountries: action.payload});
	}

	@Action(DeliveryOptionEditRulesActions.SetDeliveryOptionID)
	SetDeliveryOptionID(ctx: StateContext<DeliveryOptionEditRulesModel>, action: DeliveryOptionEditRulesActions.SetDeliveryOptionID) {
		ctx.patchState({deliveryOptionID: action.payload});
	}

	@Action(DeliveryOptionEditRulesActions.DeleteRule)
	DeleteRule(ctx: StateContext<DeliveryOptionEditRulesModel>, action: DeliveryOptionEditRulesActions.DeleteRule) {
		return this.deleteRule.mutate({ruleID: action.payload})
			.subscribe((res) => {
				const state = ctx.getState();
				ctx.dispatch(
					new DeliveryOptionEditRulesActions.FetchDeliveryOptionEditRulesRuleEdit({
						selectedIndex: state.selectedRuleIndex
					}))
			});
	}

	@Action(DeliveryOptionEditRulesActions.ClearRulesDialog)
	ClearRulesDialog(ctx: StateContext<DeliveryOptionEditRulesModel>, action: DeliveryOptionEditRulesActions.ClearRulesDialog) {
		const state = produce(ctx.getState(), st => {
			st.constraints = [];
		});
		ctx.dispatch(new DeliveryOptionEditRulesActions.SetRuleConstraints(state.constraints));
	}

}
