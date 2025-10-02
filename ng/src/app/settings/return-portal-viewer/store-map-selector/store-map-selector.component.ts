import {Component, EventEmitter} from '@angular/core';
import { HttpClient } from "@angular/common/http";
import {DialogRef} from "@angular/cdk/dialog";

type markerPoint = {
	position: {lat: number; lng: number};
	label: string;
	address: string;
}

@Component({
  selector: 'app-store-map-selector',
  templateUrl: './store-map-selector.component.html',
  styleUrls: ['./store-map-selector.component.scss']
})
export class StoreMapSelectorComponent {
	apiLoaded = false;

	addressChanged = new EventEmitter<string>()

	options: google.maps.MapOptions = {
		center: {lat: 56.153729482069686, lng: 10.20627097426347},
		zoom: 15,
	};

	markers: markerPoint[] = [
		{position: {lat: 56.157410, lng: 10.207579}, label: 'Stor torv', address: "Store Torv 15 Aarhus 8000"},
		{position: {lat: 56.149426, lng: 10.203674}, label: 'Bruuns Galleri', address: "M. P. Bruuns Gade 25 Aarhus 8000"},
	];

/*	const customMarkerIcon: google.maps.Icon = {
		url: 'path/to/custom-marker.png', // Replace with the path to your custom marker icon
		scaledSize: new google.maps.Size(50, 50), // Adjust the size as needed
	};*/

	constructor(http: HttpClient, private ref: DialogRef) {
		http.jsonp('https://maps.googleapis.com/maps/api/js?key=<replace-me>', 'callback')
			.subscribe(() => {
				this.apiLoaded = true
			});
	}

	markerClicked(marker: markerPoint) {
		console.log(marker)
		this.addressChanged.emit(marker.address);
		this.ref.close();
	}

}
