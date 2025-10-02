
export namespace RegisterActions {
	export class SetRegistrationInfo {
		static readonly type = '[Register] set registration token';
		constructor(public payload: {registrationToken: string; url: string}) {}
	}
	export class Submit {
		static readonly type = '[Register] submit';
	}
	//export type FetchProducts = NonNullable<NonNullable<FetchProductsQuery['products']['edges']>[0]>['node'];
}
