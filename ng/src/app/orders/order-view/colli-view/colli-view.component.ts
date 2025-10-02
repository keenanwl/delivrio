import {Component, Input} from '@angular/core';
import {OrderViewActions} from "../order-view.actions";
import ColliResponse = OrderViewActions.ColliResponse;

@Component({
	selector: 'app-colli-view',
	templateUrl: './colli-view.component.html',
	styleUrls: ['./colli-view.component.scss']
})
export class ColliViewComponent {
	@Input() colli: ColliResponse | undefined;
}
