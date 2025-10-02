import {Component, OnDestroy, OnInit} from '@angular/core';
import {AsyncPipe, NgIf} from "@angular/common";
import {DvoCardComponent} from "../../../shared/dvo-card/dvo-card.component";
import {MatError, MatFormField, MatLabel} from "@angular/material/form-field";
import {MatFabButton, MatIconButton} from "@angular/material/button";
import {MatIcon} from "@angular/material/icon";
import {MatInput} from "@angular/material/input";
import {NgxsFormErrorsPluginModule} from "../../../plugins/ngxs-form-errors/ngxs-form-errors.module";
import {NgxsFormPluginModule} from "@ngxs/form-plugin";
import {FormControl, FormGroup, ReactiveFormsModule} from "@angular/forms";
import {Observable, Subscription} from "rxjs";
import {ActivatedRoute} from "@angular/router";
import {Store} from "@ngxs/store";
import {CarrierEditEasyPostActions} from "./carrier-edit-easy-post.actions";
import SetID = CarrierEditEasyPostActions.SetID;
import FetchCarrierEasyPostEdit = CarrierEditEasyPostActions.FetchCarrierEasyPostEdit;
import Clear = CarrierEditEasyPostActions.Clear;
import SaveForm = CarrierEditEasyPostActions.SaveForm;
import {CarrierEditEasyPostModel, CarrierEditEasyPostState} from "./carrier-edit-easy-post.ngxs";
import {MatSlideToggle} from "@angular/material/slide-toggle";
import {ToggleContainerComponent} from "../../../shared/toggle-container/toggle-container.component";

@Component({
  selector: 'app-carrier-edit-easy-post',
  standalone: true,
	imports: [
		AsyncPipe,
		DvoCardComponent,
		MatError,
		MatFabButton,
		MatFormField,
		MatIcon,
		MatInput,
		MatLabel,
		NgIf,
		NgxsFormErrorsPluginModule,
		NgxsFormPluginModule,
		ReactiveFormsModule,
		MatIconButton,
		MatSlideToggle,
		ToggleContainerComponent
	],
  templateUrl: './carrier-edit-easy-post.component.html',
  styleUrl: './carrier-edit-easy-post.component.scss'
})
export class CarrierEditEasyPostComponent implements OnInit, OnDestroy {

	carrierEditEasyPost$: Observable<CarrierEditEasyPostModel>;

	editForm = new FormGroup({
		carrierEasyPost: new FormGroup({
			apiKey: new FormControl('', {nonNullable: true}),
			carrierAccounts: new FormControl<string[]>([], {nonNullable: true}),
			test: new FormControl(true, {nonNullable: true}),
		}),
		name: new FormControl('', {nonNullable: true}),
	});

	subscriptions$: Subscription[] = [];

	constructor(
		private route: ActivatedRoute,
		private store: Store
	) {
		this.carrierEditEasyPost$ = store.select(CarrierEditEasyPostState.get);
	}

	ngOnInit() {
		this.subscriptions$.push(this.route.queryParams
			.subscribe((params) => {
				this.store.dispatch([
					new SetID(!!params.id ? params.id : ''),
					new FetchCarrierEasyPostEdit(),
				]);
			}));
	}

	ngOnDestroy(): void {
		this.subscriptions$.forEach((s) => s.unsubscribe());
		this.store.dispatch(new Clear());
	}

	save() {
		this.store.dispatch(new SaveForm());
	}
}
