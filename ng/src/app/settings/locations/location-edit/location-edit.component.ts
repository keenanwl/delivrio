import {Component, OnDestroy, OnInit} from '@angular/core';
import {Observable} from "rxjs";
import {LocationEditModel, LocationEditState} from './location-edit.ngxs';
import {Actions, ofActionCompleted, Store} from "@ngxs/store";
import {ActivatedRoute} from "@angular/router";
import {LocationEditActions} from "./location-edit.actions";
import FetchLocationEdit = LocationEditActions.FetchLocationEdit;
import SetLocationID = LocationEditActions.SetLocationID;
import {FormArray, FormControl, FormGroup} from "@angular/forms";
import SetLocationEdit = LocationEditActions.SetLocationEdit;
import LocationTagResponse = LocationEditActions.LocationTagResponse;
import {MatListOption} from "@angular/material/list";
import SetSelectedLocationTags = LocationEditActions.SetSelectedLocationTags;
import Save = LocationEditActions.Save;
import Clear = LocationEditActions.Clear;
import SearchCountry = LocationEditActions.SearchCountry;
import CountriesResponse = LocationEditActions.CountriesResponse;
import ChangeCountry = LocationEditActions.ChangeCountry;

@Component({
	selector: 'app-location-edit',
	templateUrl: './location-edit.component.html',
	styleUrls: ['./location-edit.component.scss']
})
export class LocationEditComponent implements OnInit, OnDestroy {

	locationsEdit$: Observable<LocationEditModel>;

	editForm = new FormGroup({
		name: new FormControl('', {nonNullable: true}),
		address: new FormGroup({
			firstName: new FormControl('', {nonNullable: true}),
			lastName: new FormControl('', {nonNullable: true}),
			phoneNumber: new FormControl('', {nonNullable: true}),
			vatNumber: new FormControl('', {nonNullable: true}),
			email: new FormControl('', {nonNullable: true}),
			addressOne: new FormControl('', {nonNullable: true}),
			addressTwo: new FormControl('', {nonNullable: true}),
			zip: new FormControl('', {nonNullable: true}),
			city: new FormControl('', {nonNullable: true}),
			state: new FormControl('', {nonNullable: true}),
			country: new FormGroup({
				id: new FormControl('', {nonNullable: true}),
				label: new FormControl('', {nonNullable: true}),
				alpha2: new FormControl('', {nonNullable: true}),
			}),
			company: new FormControl('', {nonNullable: true}),
		}),
		locationTags: new FormArray<FormControl<LocationTagResponse>>([]),
	});

	constructor(
		private store: Store,
	    private route: ActivatedRoute,
		private actions$: Actions,
	) {
		this.locationsEdit$ = store.select(LocationEditState.get);
	}

	ngOnInit(): void {

		this.route.queryParams
			.subscribe((params) => {
				this.store.dispatch([
					new SetLocationID(!!params.id ? params.id : ''),
					new FetchLocationEdit(),
				]);
			});

		this.actions$.pipe(ofActionCompleted(SetLocationEdit))
			.subscribe((r) => {
				//const state = this.store.selectSnapshot(LocationEditState.get).locationEditForm.
			});
	}

	ngOnDestroy() {
		this.store.dispatch(new Clear());
	}

	isTagSelected(tagID: string, selectedTags: LocationTagResponse[] | null): boolean {
		if (!!selectedTags) {
			return selectedTags.some((t) => {
				if (t.id === tagID) {
					return true;
				}
				return false;
			});
		}
		return false;
	}

	tagsChanged(selection: MatListOption[]) {
		const tags: LocationTagResponse[] = [];
		selection.forEach((t) => {
			tags.push(t.value);
		})
		this.store.dispatch(new SetSelectedLocationTags(tags));
	}

	save() {
		this.store.dispatch(new Save());
	}

	searchCountries(term: string) {
		this.store.dispatch(new SearchCountry(term));
	}

	changeCountry(country: CountriesResponse) {
		this.store.dispatch(new ChangeCountry(country));
	}

}
