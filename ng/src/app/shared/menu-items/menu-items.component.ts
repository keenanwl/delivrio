import {ChangeDetectorRef, Component, Input, OnInit} from '@angular/core';
import {Store} from "@ngxs/store";
import {AppModel, AppState} from "../../app.ngxs";
import {Observable, timer} from "rxjs";
import {AppActions} from "../../app.actions";
import {Paths} from "../../app-routing.module";

@Component({
	selector: 'app-menu-items',
	templateUrl: './menu-items.component.html',
	styleUrls: ['./menu-items.component.scss']
})
export class MenuItemsComponent implements OnInit {

	@Input() tenantName = "";
	paths = Paths;
	app$: Observable<AppModel>;

	constructor(private store: Store, private cd: ChangeDetectorRef) {
		this.app$ = store.select(AppState.get);
	}

	ngOnInit(): void {
	}

	logout() {
		this.store.dispatch(new AppActions.Logout());
	}

}
