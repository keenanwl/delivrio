import {FetchHypothesisTestQuery} from "./hypothesis-testing-delivery-option-edit.generated";

export namespace HypothesisTestingDeliveryOptionEditActions {
	export class Fetch {
		static readonly type = '[HypothesisTestingDeliveryOptionEdit] fetch';
	}
	export class SetHypothesisTestingDeliveryOptionEdit {
		static readonly type = '[HypothesisTestingDeliveryOptionEdit] set HypothesisTestingDeliveryOption edit';
		constructor(public payload: HTResponse | undefined) {}
	}
	export class SetHypothesisTestingDeliveryOptionID {
		static readonly type = '[HypothesisTestingDeliveryOptionEdit] set HypothesisTestingDeliveryOption ID';
		constructor(public payload: string) {}
	}
	export class SetHypothesisTestingDeliveryOptionTags {
		static readonly type = '[HypothesisTestingDeliveryOptionEdit] set HypothesisTestingDeliveryOption tags';
		constructor(public payload: HTResponse[]) {}
	}
	export class SetSelectedHypothesisTestingDeliveryOptionTags {
		static readonly type = '[HypothesisTestingDeliveryOptionEdit] set selected HypothesisTestingDeliveryOption tags';
		constructor(public payload: HTResponse[]) {}
	}
	export class SetAvailableDeliveryOptions {
		static readonly type = '[HypothesisTestingDeliveryOptionEdit] set available delivery options';
		constructor(public payload: HTDeliveryOptionResponse[]) {}
	}
	export class MoveDeliveryOption {
		static readonly type = '[HypothesisTestingDeliveryOptionEdit] move delivery option';
		constructor(public payload: {container: "available" | "control" | "test"; deliveryOption: HTDeliveryOptionResponse;}) {}
	}
	export class Save {
		static readonly type = '[HypothesisTestingDeliveryOptionEdit] save';
	}
	export class Clear {
		static readonly type = '[HypothesisTestingDeliveryOptionEdit] clear';
	}
	export type HTResponse = NonNullable<FetchHypothesisTestQuery['hypothesisTest']>;
	export type HTDeliveryOptionResponse = NonNullable<NonNullable<FetchHypothesisTestQuery['unassignedDeliveryOptions']>[0]>;
}
