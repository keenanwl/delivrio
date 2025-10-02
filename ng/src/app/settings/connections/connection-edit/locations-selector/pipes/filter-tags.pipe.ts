import {Pipe, PipeTransform} from '@angular/core';
import {GraphQLError} from "graphql/index";
import {FormControl} from "@angular/forms";
import {ConnectionEditActions} from "../../connection-edit.actions";
import LocationsResponse = ConnectionEditActions.LocationsResponse;

@Pipe({
	name: 'filterTags',
	pure: true,
})
export class FilterTagsPipe implements PipeTransform {

	constructor() {
	}

	transform(locations: LocationsResponse[], internalTagID: string): LocationsResponse[] {
		return locations.filter((l) =>
			l.locationTags.some((t) => t.internalID === internalTagID));
	}
}
