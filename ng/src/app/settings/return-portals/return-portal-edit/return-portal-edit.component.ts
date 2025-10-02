import {Component, OnDestroy, OnInit} from '@angular/core';
import {Observable, Subscription} from "rxjs";
import {ReturnPortalEditModel, ReturnPortalEditState} from "./return-portal-edit.ngxs";
import {Actions, ofActionCompleted, Store} from "@ngxs/store";
import {ActivatedRoute} from "@angular/router";
import {ReturnPortalEditActions} from "./return-portal-edit.actions";
import FetchReturnPortalEdit = ReturnPortalEditActions.FetchReturnPortalEdit;
import SetReturnPortalEdit = ReturnPortalEditActions.SetReturnPortalEdit;
import ClaimResponse = ReturnPortalEditActions.ClaimResponse;
import {FormArray, FormControl, FormGroup} from "@angular/forms";
import SetReturnPortalID = ReturnPortalEditActions.SetReturnPortalID;
import AddClaim = ReturnPortalEditActions.AddClaim;
import DeleteClaim = ReturnPortalEditActions.DeleteClaim;
import Save = ReturnPortalEditActions.Save;
import {Paths} from "../../../app-routing.module";
import {MatChipListboxChange} from "@angular/material/chips";
import SetSelectedDeliveryOptions = ReturnPortalEditActions.SetSelectedDeliveryOptions;
import {EmailTemplateMergeType} from "../../../../generated/graphql";

@Component({
	selector: 'app-return-portal-edit',
	templateUrl: './return-portal-edit.component.html',
	styleUrls: ['./return-portal-edit.component.scss']
})
export class ReturnPortalEditComponent implements OnInit, OnDestroy {

	returnPortalEdit$: Observable<ReturnPortalEditModel>;

	editForm = new FormGroup({
		name: new FormControl(''),
		returnOpenHours: new FormControl(24 * 30),
		automaticallyAccept: new FormControl(false),
		returnLocation: new FormGroup({
			id: new FormControl(''),
		}),
		deliveryOptions: new FormArray<ReturnType<typeof this.newDeliveryOption>>([]),
		emailConfirmationLabel: new FormGroup({
			id: new FormControl(null),
		}),
		emailConfirmationQrCode: new FormGroup({
			id: new FormControl(null),
		}),
		emailReceived: new FormGroup({
			id: new FormControl(null),
		}),
		emailAccepted: new FormGroup({
			id: new FormControl(null),
		}),
		connection: new FormGroup({
			id: new FormControl(''),
			name: new FormControl(''),
		}),
		returnPortalClaim: new FormArray<ReturnType<typeof this.newClaim>>([]),
	});

	returnMethodPath = Paths.SETTINGS_DELIVERY_OPTIONS;
	emailTemplatePath = Paths.SETTINGS_EMAIL_TEMPLATES;

	iframe = `<iframe style="width: 100%;
\t\t\theight: 100%;
\t\t\tborder: none;" src="HOST/return-portal-viewer?id=ID"></iframe>`

	webComponent = `
	<script src="HOST/static/return-portal-assets/polyfills.js" type="module"></script>
	<script src="HOST/static/return-portal-assets/main.js" type="module"></script>

	<delivrio-return-portal
		style="--default-font: 'Courier New'; --primary-color: #4fb500;"
		page1-title="Return order"
		page2-title="Which items?"
		page3-title="How would you like to return it?"
		page4-title="Return submitted"
		page1-help="Some nice help text..."
		portalid="ID"
		url="HOST">
	</delivrio-return-portal>
	`

	viewerPath = Paths.RETURN_PORTAL_VIEWER;
	subscriptions$: Subscription[] = [];

	mergeTypeReturnColliLabel = EmailTemplateMergeType.ReturnColliLabel
	mergeTypeReturnColliQRCode = EmailTemplateMergeType.ReturnColliQr
	mergeTypeReturnColliReceived = EmailTemplateMergeType.ReturnColliReceived
	mergeTypeReturnColliAccepted = EmailTemplateMergeType.ReturnColliAccepted

	constructor(
		private store: Store,
		private route: ActivatedRoute,
		private actions$: Actions,
	) {
		this.returnPortalEdit$ = store.select(ReturnPortalEditState.get);
	}

	ngOnInit() {
		this.subscriptions$.push(
			this.route.queryParams
				.subscribe((params) => {
					this.store.dispatch([
						new SetReturnPortalID(!!params.id ? params.id : ''),
						new FetchReturnPortalEdit(),
					]);
				}));

		this.subscriptions$.push(
			this.actions$.pipe(ofActionCompleted(SetReturnPortalEdit, AddClaim, DeleteClaim))
				.subscribe((r) => {
					this.editForm.controls.returnPortalClaim?.clear();
					const portal = this.store.selectSnapshot(ReturnPortalEditState.get).returnPortalEditForm.model
					portal?.returnPortalClaim?.forEach((p) => {
						if (!!p) {
							this.editForm.controls.returnPortalClaim.push(this.newClaim(p));
						}
					});

					portal?.deliveryOptions?.forEach((p) => {
						if (!!p) {
							this.editForm.controls.deliveryOptions.push(this.newDeliveryOption(p.id));
						}
					});
				}));

	}

	ngOnDestroy() {
		this.subscriptions$.forEach((s) => s.unsubscribe());
	}

	save() {
		this.store.dispatch(new Save());
	}

	newClaim(claim: ClaimResponse) {
		return new FormGroup({
			id: new FormControl(claim.id, {nonNullable: true}),
			name: new FormControl(claim.name, {nonNullable: true}),
			description: new FormControl(claim.description, {nonNullable: true}),
			restockable: new FormControl(claim.restockable, {nonNullable: true}),
		});
	}

	newDeliveryOption(id: string) {
		return new FormGroup({
			id: new FormControl(id),
		});
	}

	addClaim() {
		this.store.dispatch(new AddClaim());
	}

	deleteClaimRow(index: number) {
		this.store.dispatch(new DeleteClaim(index));
	}

	replaceTokens(id: string, base: string): string {
		return base
			.replace(new RegExp('ID', "g"), id)
			.replace(
				new RegExp('HOST', "g"),
				`${window.location.protocol}//${window.location.host}`
			)
	}

	selectDeliveryOptions(change: MatChipListboxChange) {
		this.store.dispatch(new SetSelectedDeliveryOptions(change.value));
	}

}
