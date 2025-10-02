import {Component, Input, OnInit} from '@angular/core';

@Component({
	selector: 'app-already-customer-header',
	templateUrl: './already-customer-header.component.html',
	styleUrls: ['./already-customer-header.component.scss']
})
export class AlreadyCustomerHeaderComponent implements OnInit {

	@Input() progressPercent = 0;
	@Input() progressStep = 1;

	constructor() { }

	ngOnInit(): void {
	}

}
