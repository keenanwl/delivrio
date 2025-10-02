import {ProfileQuery} from "./profile.generated";

export namespace ProfileActions {
	export class FetchMyProfile {
		static readonly type = '[Profile] fetch my profile';
	}
	export class SetMyProfile {
		static readonly type = '[Profile] set my profile';
		constructor(public payload: SelectProfileQueryResponse | undefined) {}
	}
	export class SetLanguageList {
		static readonly type = '[Profile] set language list';
		constructor(public payload: LanguageListQueryResponse[]) {}
	}
	export class SetTenantList {
		static readonly type = '[Profile] set tenant list';
		constructor(public payload: AvailableTenant[]) {}
	}
	export class SaveForm {
		static readonly type = '[Profile] save form';
	}
	export type SelectProfileQueryResponse = NonNullable<ProfileQuery['user']>;
	export type AvailableTenant = NonNullable<ProfileQuery['availableTenants'][0]>;
	export type LanguageListQueryResponse = NonNullable<NonNullable<ProfileQuery['languages']['edges']>[0]>['node'];
}
