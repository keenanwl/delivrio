import {Component, OnInit} from '@angular/core';
import {FormControl, FormGroup} from "@angular/forms";
import {Store} from "@ngxs/store";
import {ProfileActions} from "./profile.actions";
import FetchMyProfile = ProfileActions.FetchMyProfile;
import {Observable} from "rxjs";
import {ProfileModel, ProfileState} from "./profile.ngxs";
import LanguageListQueryResponse = ProfileActions.LanguageListQueryResponse;
import SaveForm = ProfileActions.SaveForm;

@Component({
	selector: 'app-profile',
	templateUrl: './profile.component.html',
	styleUrls: ['./profile.component.scss']
})
export class ProfileComponent implements OnInit {

	profile$: Observable<ProfileModel>;

	myProfileForm = new FormGroup({
		name: new FormControl(''),
		surname: new FormControl(''),
		email: new FormControl(''),
		phoneNumber: new FormControl(''),
		language: new FormControl(),
		marketingConsent: new FormControl(true),
		tenant: new FormGroup({id: new FormControl('')}),
	});

	languageComparisonFunction = (option: LanguageListQueryResponse, value: LanguageListQueryResponse): boolean => {
		return option?.id === value?.id;
	}

	constructor(
		private store: Store,
	) {
		this.profile$ = store.select(ProfileState.state);
	}

	ngOnInit(): void {
		this.store.dispatch([new FetchMyProfile()]);
	}

	onSubmit() {
		this.store.dispatch(new SaveForm());
	}

}
