import {FetchShipmentsQuery} from "./shipments-list.generated";
import {activeFilter, searchInputList, selectedOptionList} from "../../shared/filter-bar/filter-bar.component";

export namespace ShipmentsListActions {
	export class FetchShipments {
		static readonly type = '[ShipmentsList] fetch shipments';
	}
	export class SetShipments {
		static readonly type = '[ShipmentsList] set shipments';
		constructor(public payload: ShipmentsResponse[]) {}
	}
	export class ToggleAll {
		static readonly type = '[ShipmentsList] toggle all';
	}
	export class ToggleRows {
		static readonly type = '[ShipmentsList] toggle rows';
		constructor(public payload: ShipmentsResponse[]) {}
	}
	export class SearchFilterChanges {
		static readonly type = '[ShipmentsList] search filter changed';
		constructor(public payload: activeFilter) {}
	}
	export class SearchCCLocations {
		static readonly type = '[ShipmentsList] search cc locations';
		constructor(public payload: {lookup: string}) {}
	}
	export class SetListOptions {
		static readonly type = '[ShipmentsList] set list options';
		constructor(public payload: searchInputList) {}
	}
	export class SetSelectedFilters {
		static readonly type = '[ShipmentsList] set selected filters';
		constructor(public payload: selectedOptionList) {}
	}
	export class SetEmailTemplates {
		static readonly type = '[ShipmentsList] set email templates';
		constructor(public payload: EmailTemplatesResponse[]) {}
	}
	export class SendOverviewEmail {
		static readonly type = '[ShipmentsList] send overview email';
		constructor(public payload: {email: string; templateID: string;}) {}
	}

	export type ShipmentsResponse = NonNullable<NonNullable<NonNullable<FetchShipmentsQuery['shipments']>['edges']>[0]>['node'];
	export type EmailTemplatesResponse = NonNullable<NonNullable<NonNullable<NonNullable<FetchShipmentsQuery['emailTemplates']>['edges']>[0]>['node']>;
}
