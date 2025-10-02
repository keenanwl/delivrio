import {Component, Input, OnInit} from '@angular/core';
import {Observable} from "rxjs";
import {AppModel, AppState} from "../app.ngxs";
import {Store} from "@ngxs/store";
import {Paths} from "../app-routing.module";

@Component({
	selector: 'app-account',
	templateUrl: './account.component.html',
	styleUrls: ['./account.component.scss']
})
export class AccountComponent implements OnInit {

	app$: Observable<AppModel>;
@Input() path: string = "";
	constructor(private store: Store) {
		this.app$ = store.select(AppState.get);
	}

	ngOnInit(): void {
	}

	protected readonly Paths = Paths;
}
