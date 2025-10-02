import {Component, Input, input} from '@angular/core';
import {MatSlideToggle} from "@angular/material/slide-toggle";
import {NgIf} from "@angular/common";

@Component({
	selector: 'app-toggle-container',
	standalone: true,
	imports: [
		MatSlideToggle,
		NgIf
	],
	templateUrl: './toggle-container.component.html',
	styleUrl: './toggle-container.component.scss'
})
export class ToggleContainerComponent {
	@Input() name = "";
}
