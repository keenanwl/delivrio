import { HttpClient } from "@angular/common/http";
import {Injectable} from "@angular/core";
import {Observable} from "rxjs";

@Injectable()
export class RequestPasswordResetService {

	constructor(public http: HttpClient) {
	}

	postRequestEmail(email: string): Observable<RequestPasswordResetResponse> {
		return this.http.post<RequestPasswordResetResponse>(`/api/requestEmail`, {
			email,
		});
	}

}

export interface RequestPasswordResetResponse {
	success: boolean;
}
