import {Component, Input} from '@angular/core';

export interface PackingDimensions {
	length: number;
	height: number;
	width: number;
}

@Component({
	selector: 'app-dimensions-viewer',
	templateUrl: './dimensions-viewer.component.html',
	styleUrls: ['./dimensions-viewer.component.scss']
})
export class DimensionsViewerComponent {
	@Input()
	set dimensions(value: PackingDimensions) {
		this._dimensions = value;
		this.setDimensions(this._dimensions)
	}
	private _dimensions: PackingDimensions = {length: 0, height: 0, width: 0};

	@Input() showLabels = false;

	boxDisplayIndex = 0;
	boxDisplaySide = [
		"show-front",
		"show-back",
		"show-right",
		"show-left",
		"show-top",
		"show-bottom",
	]

	length = 300;
	height = 300;
	width = 250;

	constructor() {
	}

	setDimensions(dimensions: PackingDimensions) {
		const normalized = this.normalizeDimensions(dimensions, 300, 300);
		this.length = normalized.length;
		this.height = normalized.height;
		this.width = normalized.width;
	}

	normalizeDimensions(dim: PackingDimensions, minVal: number, maxVal: number): PackingDimensions {
		let maxDimension = Math.max(dim.length, dim.width, dim.height)

		let scaleFactor = maxVal / maxDimension
		if (maxDimension < minVal) {
			scaleFactor = minVal / maxDimension
		}

		return {
			length: dim.length * scaleFactor,
			width:  dim.width * scaleFactor,
			height: dim.height * scaleFactor,
		}
	}

	rotate() {
		if (this.boxDisplayIndex < this.boxDisplaySide.length -1) {
			this.boxDisplayIndex++;
		} else {
			this.boxDisplayIndex = 0;
		}
	}

}
