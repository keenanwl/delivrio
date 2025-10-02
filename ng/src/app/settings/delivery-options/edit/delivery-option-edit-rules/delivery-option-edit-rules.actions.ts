import {
	DeliveryRuleConstraintComparison,
	DeliveryRuleConstraintGroupConstraintLogic, DeliveryRuleConstraintPropertyType
} from "../../../../../generated/graphql";
import {
	UpdateDeliveryOptionGlsMutationVariables
} from "../delivery-options-edit-gls/delivery-options-edit-gls.generated";
import {
	DeliveryOptionsSearchCountriesQuery,
	FetchDeliveryOptionRulesQuery, FetchRuleConstraintsQuery,
	SearchProductTagsQuery
} from "./delivery-option-edit-rules.generated";

export namespace DeliveryOptionEditRulesActions {

	export class FetchDeliveryOptionEditRules {
		static readonly type = '[deliveryOptionEditRules] fetch delivery options edit';
	}
	export class SetDeliveryOptionEditRules {
		static readonly type = '[deliveryOptionEditRules] set delivery options edit';
		//constructor(public payload: SelectDeliveryOptionEditRulesQueryResponse) {}
	}
	export class FetchDeliveryOptionEditRulesRuleEdit {
		static readonly type = '[deliveryOptionEditRules] fetch delivery options rules edit';
		constructor(public payload: {selectedIndex: number}) {}
	}
	export class SetDeliveryOptionEditRulesRulesEdit {
		static readonly type = '[deliveryOptionEditRules] set delivery options rules edit';
		constructor(public payload: SelectDeliveryOptionEditRulesRulesQueryResponse[]) {}
	}
	export class SaveForm {
		static readonly type = '[deliveryOptionEditRules] save form';
		constructor(public payload: UpdateDeliveryOptionGlsMutationVariables) {}
	}
	export class CreateDeliveryRule {
		static readonly type = '[deliveryOptionEditRules] create delivery rule';
		constructor(public payload: string) {}
	}
	export class SetRuleList {
		static readonly type = '[deliveryOptionEditRules] set rule list';
		constructor(public payload: SelectDeliveryOptionEditRulesRulesQueryResponse[]) {}
	}
	export class FetchRuleConstrains {
		static readonly type = '[deliveryOptionEditRules] fetch rule constraints';
	}
	// An alias for SetRuleConstraints that can subscribed to without creating a listening loop
	export class SetRuleConstraintsNotify {
		static readonly type = '[deliveryOptionEditRules] set rule constraints notify';
		constructor(public payload: FetchRuleConstraintsQueryResponse) {}
	}
	export class SetRuleConstraints {
		static readonly type = '[deliveryOptionEditRules] set rule constraints';
		constructor(public payload: FetchRuleConstraintsQueryResponse) {}
	}
	export class SetSelectedConstraintID {
		static readonly type = '[deliveryOptionEditRules] set selected constraint ID/index';
		constructor(public payload: {id: string, index: number}) {}
	}
	export class FetchRuleOptions {
		static readonly type = '[deliveryOptionEditRules] fetch rule options';
	}
	export class SaveNewConstraintGroup {
		static readonly type = '[deliveryOptionEditRules] save constraint new group';
	}
	export class SaveEditConstraintGroup {
		static readonly type = '[deliveryOptionEditRules] save constraint edit group';
	}
	export class SetSelectedRule {
		static readonly type = '[deliveryOptionEditRules] set selected rule';
		constructor(public payload: {ruleIndex: number, ruleID: string}) {}
	}
	export class DeleteConstraintGroup {
		static readonly type = '[deliveryOptionEditRules] delete constraint group';
		constructor(public payload: string) {}
	}
	export class SetSelectedOption {
		static readonly type = '[deliveryOptionEditRules] set selected option';
		constructor(public payload: string) {}
	}
	export class DeleteRule {
		static readonly type = '[deliveryOptionEditRules] delete rule';
		constructor(public payload: string) {}
	}
	export class UpdateDayOfWeek {
		static readonly type = '[deliveryOptionEditRules] update day of week';
		constructor(public payload: {index: number, dayOfWeek: string[]}) {}
	}
	export class AddConstraintLine {
		static readonly type = '[deliveryOptionEditRules] add constraint line';
	}
	export class DeleteConstraintLine {
		static readonly type = '[deliveryOptionEditRules] delete constraint line';
		constructor(public payload: {index: number}) {}
	}
	export class UpdateTimeOfDay {
		static readonly type = '[deliveryOptionEditRules] update time of day';
		constructor(public payload: {index: number, timeOfDay: string[]}) {}
	}
	export class UpdateConstraintProperty {
		static readonly type = '[deliveryOptionEditRules] update constraint property';
		constructor(public payload: {index: number, propertyType: DeliveryRuleConstraintPropertyType}) {}
	}
	export class UpdateConstraintComparison {
		static readonly type = '[deliveryOptionEditRules] update constraint comparison';
		constructor(public payload: {index: number, comparison: DeliveryRuleConstraintComparison}) {}
	}
	export class SetProductTagSearchTerm {
		static readonly type = '[deliveryOptionEditRules] set product tag search term';
		constructor(public payload: string) {}
	}
	export class SetProductTagSearch {
		static readonly type = '[deliveryOptionEditRules] set product tag search';
		constructor(public payload: SearchProductTagsResponse[]) {}
	}
	export class FetchProductTags {
		static readonly type = '[deliveryOptionEditRules] fetch product tags';
	}
	export class AddProductTag {
		static readonly type = '[deliveryOptionEditRules] add product tag';
		constructor(public payload: {index: number, tag: FetchRuleTagResponse}) {}
	}
	export class RemoveProductTag {
		static readonly type = '[deliveryOptionEditRules] remove product tag';
		constructor(public payload: {index: number, tag: FetchRuleTagResponse}) {}
	}
	export class RemovePostalCode {
		static readonly type = '[deliveryOptionEditRules] remove postal code numeric';
		constructor(public payload: {index: number, code: number}) {}
	}
	export class RemovePostalCodeString {
		static readonly type = '[deliveryOptionEditRules] remove postal code string';
		constructor(public payload: {index: number, code: string}) {}
	}
	export class AddPostalCode {
		static readonly type = '[deliveryOptionEditRules] add postal code numeric';
		constructor(public payload: {index: number, code: number}) {}
	}
	export class AddPostalCodeString {
		static readonly type = '[deliveryOptionEditRules] add postal code string';
		constructor(public payload: {index: number, code: string}) {}
	}
	export class UpdateNumericValue {
		static readonly type = '[deliveryOptionEditRules] update numeric value';
		constructor(public payload: {index: number, value: number}) {}
	}
	export class AddCountry {
		static readonly type = '[deliveryOptionEditRules] add country';
		constructor(public payload: {country: CountriesResponse}) {}
	}
	export class RemoveCountry {
		static readonly type = '[deliveryOptionEditRules] remove country';
		constructor(public payload: {id: string}) {}
	}
	export class SearchCountry {
		static readonly type = '[deliveryOptionEditRules] search country';
		constructor(public payload: string) {}
	}
	export class SetCountrySearch {
		static readonly type = '[deliveryOptionEditRules] set country search';
		constructor(public payload: CountriesResponse[]) {}
	}
	export class SetDeliveryOptionID {
		static readonly type = '[deliveryOptionEditRules] set delivery option ID';
		constructor(public payload: string) {}
	}
	export class SetConstraintLogicType {
		static readonly type = '[deliveryOptionEditRules] set constraint logic type';
		constructor(public payload: DeliveryRuleConstraintGroupConstraintLogic) {}
	}
	// Different actions to prevent loop
	export class UpdateConstraintLogicType {
		static readonly type = '[deliveryOptionEditRules] update constraint logic type';
		constructor(public payload: DeliveryRuleConstraintGroupConstraintLogic) {}
	}
	export class SetCurrencies {
		static readonly type = '[deliveryOptionEditRules] set currencies';
		constructor(public payload: CurrencyResponse[]) {}
	}
	export class UpdateRulePricing {
		static readonly type = '[deliveryOptionEditRules] update rule pricing';
		constructor(public payload: {price: number; currency: CurrencyResponse}) {}
	}
	export class ClearRulesDialog {
		static readonly type = '[deliveryOptionEditRules] clear rules dialog';
	}
	export class Clear {
		static readonly type = '[deliveryOptionEditRules] clear';
	}

	export type SearchProductTagsResponse = NonNullable<NonNullable<NonNullable<SearchProductTagsQuery['productTags']['edges']>[0]>['node']>;
	//export type SelectDeliveryOptionEditRulesQueryResponse = NonNullable<FetchDeliveryOptionEditRulesQuery['deliveryOptionGLS']>;
	export type SelectDeliveryOptionEditRulesRulesQueryResponse = NonNullable<NonNullable<FetchDeliveryOptionRulesQuery['deliveryRules']['edges']>[0]>['node'];

	export type CountriesResponse = NonNullable<NonNullable<NonNullable<NonNullable<DeliveryOptionsSearchCountriesQuery['countries']>['edges']>[0]>['node']>;

	export type FetchRuleConstraintGroupQueryResponse = NonNullable<FetchRuleConstraintsQuery['constraintGroup']>;
	export type FetchRuleConstraintsQueryResponse = NonNullable<FetchRuleConstraintsQuery['constraints']>;
	export type FetchRuleTagResponse = NonNullable<NonNullable<NonNullable<NonNullable<FetchRuleConstraintsQuery['constraints']>[0]>['tags']>[0]>;
	export type CurrencyResponse = NonNullable<NonNullable<NonNullable<NonNullable<FetchDeliveryOptionRulesQuery['currencies']>['edges']>[0]>['node']>;
}
