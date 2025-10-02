import { HttpClient } from "@angular/common/http";
import {Injectable} from "@angular/core";
import {Observable} from "rxjs";

@Injectable()
export class PasswordResetService {

	constructor(public http: HttpClient) {
	}

	postNextPassword(request: ResetPasswordRequest): Observable<ResetPasswordResponse> {
		return this.http.post<ResetPasswordResponse>(`/api/resetPassword`, request);
	}

}

export interface ResetPasswordRequest {
	new_password: string;
	otk: string;
}

export interface ResetPasswordResponse {
	success: boolean;
	message: string;
}

