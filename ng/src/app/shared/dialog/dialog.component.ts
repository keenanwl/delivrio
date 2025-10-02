import {Component, Input, OnInit} from '@angular/core';

export interface ByrtieButton {
	text: string;
	callback: () => void;
}

@Component({
	selector: 'app-dialog',
	templateUrl: './dialog.component.html',
	styleUrls: ['./dialog.component.scss']
})
export class DialogComponent implements OnInit {

	@Input() title = '';
	@Input() body = '';
	@Input() buttons: ByrtieButton[] = [];

	constructor() { }

	ngOnInit(): void {
	}

	fireCallback(callback: () => void) {
		callback();
	}

}
