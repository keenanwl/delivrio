import {FetchReturnCollisViewQuery} from "./return-view.generated";
import {TimelineViewerFragment} from "../../shared/timeline-viewer/timeline-viewer.generated";

export namespace ReturnViewActions {
	export class FetchReturnView {
		static readonly type = '[ReturnView] fetch Return view';
	}
	export class SetOrderID {
		static readonly type = '[ReturnView] set order ID';
		constructor(public payload: string) {}
	}
	export class SetReturnView {
		static readonly type = '[ReturnView] set return view';
		constructor(public payload: ReturnColliResponse[]) {}
	}
	export class SetOrderPublicID {
		static readonly type = '[ReturnView] set order public ID';
		constructor(public payload: string) {}
	}
	export class SetTimeline {
		static readonly type = '[ReturnView] set timeline';
		constructor(public payload: TimelineViewerFragment[]) {}
	}
	export class ToggleShowDeleted {
		static readonly type = '[ReturnView] toggle show deleted';
	}
	export class MarkAccepted {
		static readonly type = '[ReturnView] mark accepted';
		constructor(public payload: string) {}
	}
	export class MarkDeclined {
		static readonly type = '[ReturnView] mark declined';
		constructor(public payload: string) {}
	}
	export type ReturnColliResponse = NonNullable<NonNullable<NonNullable<FetchReturnCollisViewQuery['returnColli']>['collis']>[0]>;
	export type ReturnOrderLinesResponse = NonNullable<NonNullable<ReturnColliResponse['colli']>['returnOrderLine']>;
}
