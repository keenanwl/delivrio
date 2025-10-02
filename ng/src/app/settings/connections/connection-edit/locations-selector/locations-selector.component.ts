import {Component, EventEmitter, Input, OnInit, Output} from '@angular/core';
import {FormControl, FormGroup} from "@angular/forms";
import {ConnectionEditActions} from "../connection-edit.actions";
import LocationsResponse = ConnectionEditActions.LocationsResponse;

export interface SelectedLocations {
	senderID: string | undefined;
	returnID: string | undefined;
	sellerID: string | undefined;
	pickupID: string | undefined;
}

@Component({
	selector: 'app-locations-selector',
	templateUrl: './locations-selector.component.html',
	styleUrls: ['./locations-selector.component.scss']
})
export class LocationsSelectorComponent implements OnInit {

	@Input() set selectedLocations(val: SelectedLocations) {
		this._selectedLocations = val;
		this.form.controls.seller.setValue(val.sellerID || '', {emitEvent: false});
		this.form.controls.sender.setValue(val.senderID || '', {emitEvent: false});
		this.form.controls.return.setValue(val.returnID || '', {emitEvent: false});
		this.form.controls.pickup.setValue(val.pickupID || '', {emitEvent: false});
	}
	get selectedLocations(): SelectedLocations {
		return this._selectedLocations;
	}
	_selectedLocations: SelectedLocations = {
		sellerID: '',
		senderID: '',
		returnID: '',
		pickupID: '',
	};

	@Input() set allLocations(val: LocationsResponse[]) {
		this._allLocations = val;
	}
	get allLocations(): LocationsResponse[] {
		return this._allLocations;
	}
	_allLocations: LocationsResponse[] = [];

	@Output() locationsSelected = new EventEmitter<SelectedLocations>();

	form = new FormGroup({
		seller: new FormControl('', {nonNullable: true}),
		sender: new FormControl('', {nonNullable: true}),
		return: new FormControl('', {nonNullable: true}),
		pickup: new FormControl('', {nonNullable: true}),
	});

	ngOnInit() {
		this.form.valueChanges
			.subscribe((v) => {
				this.locationsSelected.emit({
					sellerID: v.seller,
					senderID: v.sender,
					pickupID: v.pickup,
					returnID: v.return,
				})
			});
	}

}
