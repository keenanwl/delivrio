import {Component, Input} from '@angular/core';
import {CommonModule} from '@angular/common';
import {MatButtonModule} from "@angular/material/button";
import {RelativeTimePipe} from "../../pipes/relative-time.pipe";
import {MatTooltipModule} from "@angular/material/tooltip";
import {TimelineViewerFragment} from "./timeline-viewer.generated";

@Component({
	selector: 'app-timeline-viewer',
	standalone: true,
	imports: [CommonModule, MatButtonModule, RelativeTimePipe, MatTooltipModule],
	templateUrl: './timeline-viewer.component.html',
	styleUrl: './timeline-viewer.component.scss'
})
export class TimelineViewerComponent {

	@Input() timelineData: TimelineViewerFragment[] = [];
	@Input() showOnlyOrdersFromID: string = "";

}
