import {Injectable} from "@angular/core";
import { HttpClient } from "@angular/common/http";
import {CreateTenantInput, CreateUserInput} from "../../../generated/graphql";
import {Observable} from "rxjs";

interface RegisterReponse {
	success: boolean;
	message: string;
	user_pulid: string;
}

@Injectable()
export class RegisterService {

	constructor(
		private http: HttpClient,
	) {}

	public initialRegistration(inputUser: CreateUserInput, inputTenant: CreateTenantInput): Observable<RegisterReponse> {
		return this.http.post<RegisterReponse>(`/api/register`, {user_input: inputUser, tenant_input: inputTenant});
	}

}
