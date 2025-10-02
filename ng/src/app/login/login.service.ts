import {Injectable} from "@angular/core";
import { HttpClient } from "@angular/common/http";
import {Observable} from "rxjs";

@Injectable()
export class LoginService {

	constructor(public http: HttpClient) {
	}

	postLogin(email: string, password: string): Observable<LoginResponse> {
		return this.http.post<LoginResponse>(`/api/login`, {
			email: email,
			password: password,
		});
	}

	refreshJwt(): Observable<LoginResponse> {
		return this.http.get<LoginResponse>(`/restricted/api/refresh_token`);
	}

}

export interface LoginResponse {
	code: number;
	expire: string;
	token: string;
}
