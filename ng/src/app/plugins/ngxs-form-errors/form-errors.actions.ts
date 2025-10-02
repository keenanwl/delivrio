import {GraphQLError} from "graphql";

export namespace FormErrorsActions {
	export class SetFormErrors {
		static readonly type = '[NGXSFormErrorsPlugin] set form errors';
		constructor(public payload: {errors: readonly GraphQLError[], formPath: string}) {}
	}
}
