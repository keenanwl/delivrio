import {Component, Input, OnInit} from '@angular/core';

@Component({
	selector: 'app-dvo-card',
	templateUrl: './dvo-card.component.html',
	styleUrls: ['./dvo-card.component.scss'],
	standalone: true,
})
export class DvoCardComponent implements OnInit {

	@Input() background: string = "#FFFFFF";

	constructor() { }

	ngOnInit(): void {
	}

}
