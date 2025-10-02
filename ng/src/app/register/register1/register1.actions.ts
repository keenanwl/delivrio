import {CreateTenantInput, CreateUserInput} from "../../../generated/graphql";

export namespace Register1Actions {
  export class FetchMembershipInfo {
  	static readonly type = '[Register] fetch memberships in by id';
  }
  export class SetMembershipId {
  	static readonly type = '[Register] set memberships ID';
  	constructor(public payload: string) {}
  }
  export class SetInvalidParams {
  	static readonly type = '[Register] set invalid params';
  	constructor(public payload: boolean) {}
  }
  export class SetAddresses {
  	static readonly type = '[Register] set addresses';
  	//constructor(public payload: Address[]) {}
  }
  export class LookupAddress {
  	static readonly type = '[Register] lookup streetName';
  	constructor(public payload: string) {}
  }
  export class ClearAllRegister {
  	static readonly type = '[Register] clear all';
  }

  export class SetCompanyName {
  	static readonly type = '[Register] set company name';
  	constructor(public payload: string) {}
  }
  export class SetPhoneNumber {
  	static readonly type = '[Register] set phone number';
  	constructor(public payload: string) {}
  }
  export class SetVatNumber {
  	static readonly type = '[Register] set vat number';
  	constructor(public payload: string) {}
  }
  export class SetEmail {
  	static readonly type = '[Register] set email';
  	constructor(public payload: string) {}
  }
  export class SetPassword {
  	static readonly type = '[Register] set password';
  	constructor(public payload: string) {}
  }
  export class SetRepeatPassword {
  	static readonly type = '[Register] set repeat password';
  	constructor(public payload: string) {}
  }
  export class SubmitRegistrationInfo {
  	static readonly type = '[Register] submit person info';
  	constructor(public payload: {userInput: CreateUserInput, tenantInput: CreateTenantInput}) {}
  }
  export class SetInvalidMessage {
  	static readonly type = '[Register] set invalid message';
  	constructor(public payload: string) {}
  }
}
