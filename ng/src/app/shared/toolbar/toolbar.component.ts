import {Component, Input, OnInit} from '@angular/core';
import {Paths} from "../../app-routing.module";

@Component({
	selector: 'app-toolbar',
	templateUrl: './toolbar.component.html',
	styleUrls: ['./toolbar.component.scss']
})
export class ToolbarComponent implements OnInit {

	paths = Paths;
	@Input() showMenuIcon: boolean = false;

	constructor() { }

	ngOnInit(): void {
	}

}
