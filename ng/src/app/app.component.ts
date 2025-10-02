import {Component, OnInit} from '@angular/core';
import {Observable} from "rxjs";
import {AppModel, AppState} from "./app.ngxs";
import {Actions, ofActionDispatched, Store} from "@ngxs/store";
import {ActivatedRoute, Router} from "@angular/router";
import {CookieService} from "ngx-cookie-service";
import {BreakpointObserver, Breakpoints} from "@angular/cdk/layout";
import {AppActions} from "./app.actions";
import UpdateBreakpoint = AppActions.UpdateBreakpoint;
import AppChangeRoute = AppActions.AppChangeRoute;
import {LoginActions} from "./login/login.actions";
import ShowGlobalSnackbar = AppActions.ShowGlobalSnackbar;
import {MatSnackBar} from '@angular/material/snack-bar';
import FetchLoggedInUser = AppActions.FetchLoggedInUser;

@Component({
	selector: 'app-root',
	templateUrl: './app.component.html',
	styleUrls: ['./app.component.scss'],
})
export class AppComponent implements OnInit {

	app$: Observable<AppModel>;

	constructor(
	    private router: Router,
	    private route: ActivatedRoute,
	    private store: Store,
	    private cookieService: CookieService,
	    private breakpointObserver: BreakpointObserver,
	    private actions$: Actions,
	    private snackBar: MatSnackBar,
	) {
		this.app$ = store.select(AppState.get);
	}

	ngOnInit(): void {

	    this.breakpointObserver.observe([
	        Breakpoints.XSmall,
	        Breakpoints.Small,
	        Breakpoints.Medium,
	        Breakpoints.Large,
	    ]).subscribe(result => {
	        if (result.breakpoints[Breakpoints.XSmall] || result.breakpoints[Breakpoints.Small]) {
	            this.store.dispatch(new UpdateBreakpoint('mobile'));
	        } else if (result.breakpoints[Breakpoints.Medium]) {
	            this.store.dispatch(new UpdateBreakpoint('tablet'));
	        } else {
	            this.store.dispatch(new UpdateBreakpoint('desktop'));
	        }
	    });

	    this.actions$
	        .pipe(ofActionDispatched(AppChangeRoute))
	        .subscribe(({payload}) => this.router.navigate(
	            [payload.path],
	            {queryParams: payload.queryParams}
	        ));

	    let set = false;

	    this.route.queryParams
	        .subscribe((params) => {
	            if (!set && this.cookieService.check('token')) {
	                this.store.dispatch([
	                    new LoginActions.SetJwt(this.cookieService.get('token')),
	                    new FetchLoggedInUser(),
	                ]);
	                set = true;
	            }
	        });

		this.actions$
			.pipe(ofActionDispatched(ShowGlobalSnackbar))
			.subscribe(({payload}) => {
				this.snackBar.open(payload, `close`,{duration: 3000});
			});

	}

}
