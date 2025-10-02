import {CreateReturnColliOutput, Item, ReturnPortalViewResponse} from "./return-portal-frame.service";

export namespace ReturnPortalFrameActions {
	export class FetchReturnPortalFrame {
		static readonly type = '[ReturnPortalFrame] fetch return portal edit';
	}
	export class SetReturnPortalFrame {
		static readonly type = '[ReturnPortalFrame] set return portal edit';
		constructor(public payload: ReturnPortalViewResponse) {}
	}
	export class SetReturnPortalID {
		static readonly type = '[ReturnPortalFrame] set return portal ID';
		constructor(public payload: string) {}
	}
	export class SetSelectedItem {
		static readonly type = '[ReturnPortalFrame] set selected item';
		constructor(public payload: { item: Item; selected: boolean }) {}
	}
	export class SetSelectedItemReason {
		static readonly type = '[ReturnPortalFrame] set selected item reason';

		constructor(public payload: { item: Item; reasonID: string }) {}
	}
	export class IncrementQuantity {
		static readonly type = '[ReturnPortalFrame] increment quantity';
		constructor(public payload: {orderLineID: string}) {}
	}
	export class DecrementQuantity {
		static readonly type = '[ReturnPortalFrame] decrement quantity';
		constructor(public payload: {orderLineID: string}) {}
	}
	export class SetOrderInfo {
		static readonly type = '[ReturnPortalFrame] set order info';
		constructor(public payload: {email: string; orderPublicID: string}) {}
	}
	export class SetReturnCollis {
		static readonly type = '[ReturnPortalFrame] set return collis';
		constructor(public payload: CreateReturnColliOutput[]) {}
	}
	export class SelectDeliveryOption {
		static readonly type = '[ReturnPortalFrame] select delivery option';
		constructor(public payload: {returnColliID: string; deliveryOptionID: string}) {}
	}
	export class Save {
		static readonly type = '[ReturnPortalFrame] save';
	}
	export class Clear {
		static readonly type = '[ReturnPortalFrame] clear';
	}
	export class AddClaim {
		static readonly type = '[ReturnPortalFrame] add claim';
	}
	export class DeleteClaim {
		static readonly type = '[ReturnPortalFrame] delete claim';
		constructor(public payload: number) {}
	}
	export class CreateOrder {
		static readonly type = '[ReturnPortalFrame] create order';
	}
	export class SubmitDeliveryOptions {
		static readonly type = '[ReturnPortalFrame] submit delivery options';
	}
	export class ShowErrorDialog {
		static readonly type = '[ReturnPortalFrame] show error dialog';
		constructor(public payload: {title: string; body: string}) {}
	}
	export class LoadingRundown {
		static readonly type = '[ReturnPortalFrame] loading rundown';
	}
	export class StopLoading {
		static readonly type = '[ReturnPortalFrame] stop loading';
	}
	export class UpdateComment {
		static readonly type = '[ReturnPortalFrame] update comment';
		constructor(public payload: string) {}
	}
	export class SetBaseURL {
		static readonly type = '[ReturnPortalFrame] set base URL';
		constructor(public payload: string) {}
	}
}
