import {Component, Input, ViewEncapsulation} from '@angular/core';

// Exists simply to break view encapsulation for the Custom Element

@Component({
	encapsulation: ViewEncapsulation.None,
	selector: 'app-return-portal-container',
	templateUrl: './return-portal-container.component.html',
	// Ah, can we add this to our own shadow dom to isolate the overlay?
	// The goal being 0% interface in the importing pages CSS
	styleUrls: ['./return-portal-container.component.scss', '../../../../return-portal-theme.scss']
})
export class ReturnPortalContainerComponent {
	@Input() url = "";
	@Input() portalid = "";

	@Input() page1Title: string | null = null;
	@Input() page2Title: string | null = null;
	@Input() page3Title: string | null = null;
	@Input() page4Title: string | null = null;

	@Input() page1Help: string | null = null;
	@Input() page4Help: string | null = null;

	@Input() page2HelpSelected: string | null = null;
	@Input() page2SelectedItem: string | null = null;
	@Input() page2SelectedItems: string | null = null;

}
