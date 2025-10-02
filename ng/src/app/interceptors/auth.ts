import {Injectable} from '@angular/core';
import { HttpRequest, HttpHandler, HttpEvent, HttpInterceptor, HttpErrorResponse } from '@angular/common/http';
import {Observable, throwError} from 'rxjs';
import {Store} from '@ngxs/store';
import {LoginState} from '../login/login.ngxs';
import {catchError} from 'rxjs/operators';
import {LoginActions} from "../login/login.actions";

@Injectable()
export class JwtInterceptor implements HttpInterceptor {

	constructor(private store: Store) {
	}

	intercept(request: HttpRequest<any>, next: HttpHandler): Observable<HttpEvent<any>> {

		const login = this.store.selectSnapshot(LoginState.getLoginState);

		if (login.jwt.length > 0) {
			request = request.clone({
				setHeaders: {
					Authorization: `Bearer ${login.jwt}`
				}
			});
		}

		return next.handle(request).pipe(catchError((err, httpEvent) => {

			if (err instanceof HttpErrorResponse) {
				if (err.status === 401 && !!err.url && err.url.indexOf(`/api/refresh_token`) < 0 && err.url.indexOf(`/api/login`) < 0) {
					this.store.dispatch(new LoginActions.RefreshJwt());
				}
			}

			return throwError(err);
		}));
	}

}
