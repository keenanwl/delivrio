import {CreateSignupOptionsInput} from "../../../generated/graphql";

export namespace Register2Actions {
	export class SetInvalidMessage {
		static readonly type = '[Register2] set invalid message';
		constructor(public payload: string) {}
	}
	export class SaveRegistration2 {
		static readonly type = '[Register2] save registration part 2';
		constructor(public payload: CreateSignupOptionsInput) {}
	}
}
