import {
	AfterViewInit,
	Component,
	Input,
	OnDestroy,
	OnInit, Renderer2,
	ViewChild,
	ViewContainerRef,
	ViewEncapsulation
} from '@angular/core';
import {Actions, ofActionDispatched, Store} from '@ngxs/store';
import {ActivatedRoute} from "@angular/router";
import {ReturnPortalFrameActions} from "./return-portal-frame.actions";
import SetReturnPortalID = ReturnPortalFrameActions.SetReturnPortalID;
import {Observable, Subscription, timer} from "rxjs";
import {ItemReturn, ReturnPortalFrameModel, ReturnPortalFrameState} from "./return-portal-frame.ngxs";
import SetSelectedItem = ReturnPortalFrameActions.SetSelectedItem;
import SetSelectedItemReason = ReturnPortalFrameActions.SetSelectedItemReason;
import DecrementQuantity = ReturnPortalFrameActions.DecrementQuantity;
import IncrementQuantity = ReturnPortalFrameActions.IncrementQuantity;
import {CreateReturnColliOutput, Item} from "./return-portal-frame.service";
import SelectDeliveryOption = ReturnPortalFrameActions.SelectDeliveryOption;
import {MatDialog, MatDialogRef} from "@angular/material/dialog";
import {ReturnFrameErrorDialogComponent} from "./dialogs/return-frame-error-dialog/return-frame-error-dialog.component";
import {FormControl} from "@angular/forms";
import {debounceTime} from "rxjs/operators";
import {StoreMapSelectorComponent} from "../store-map-selector/store-map-selector.component";
import SetBaseURL = ReturnPortalFrameActions.SetBaseURL;
import ShowErrorDialog = ReturnPortalFrameActions.ShowErrorDialog;
import UpdateComment = ReturnPortalFrameActions.UpdateComment;
import SetOrderInfo = ReturnPortalFrameActions.SetOrderInfo;
import FetchReturnPortalFrame = ReturnPortalFrameActions.FetchReturnPortalFrame;
import CreateOrder = ReturnPortalFrameActions.CreateOrder;
import SubmitDeliveryOptions = ReturnPortalFrameActions.SubmitDeliveryOptions;
import StopLoading = ReturnPortalFrameActions.StopLoading;
import LoadingRundown = ReturnPortalFrameActions.LoadingRundown;

@Component({
	encapsulation: ViewEncapsulation.ShadowDom,
	selector: 'app-return-portal-frame',
	templateUrl: './return-portal-frame.component.html',
	styleUrls: ['./return-portal-frame.component.scss']
})
export class ReturnPortalFrameComponent implements OnInit, OnDestroy, AfterViewInit {

	@Input()
	get portalID(): string {
		return this._portalID;
	}
	set portalID(value: string) {
		this._portalID = value;
		console.error(value)
		this.store.dispatch(new SetReturnPortalID(this._portalID));
	}

	@Input()
	get url(): string {
		return this._url;
	}
	set url(value: string) {
		this._url = value;
		const validURL = this.isValidURL(this._url);
		if (validURL !== true) {
			console.error("DELIVRIO error: the URL (" + this._url + ") provided is invalid")
		}
		this.store.dispatch(new SetBaseURL(this._url));
	}
	private _url = "";
	private _portalID = "";

	@Input()
	set page1Title(value: string | null) {
		if (!!value) {
			this._page1Title = value;
		}
	}
	get page1Title(): string | null {
		return this._page1Title;
	}
	private _page1Title: string | null = "Return order";

	@Input()
	set page2Title(value: string | null) {
		if (!!value) {
			this._page2Title = value;
		}
	}
	get page2Title(): string | null {
		return this._page2Title;
	}
	private _page2Title: string | null = "Which items would you like to return?";

	@Input()
	set page3Title(value: string | null) {
		if (!!value) {
			this._page3Title = value;
		}
	}
	get page3Title(): string | null {
		return this._page3Title;
	}
	private _page3Title: string | null = "How would you like to return it?";

	@Input()
	set page4Title(value: string | null) {
		if (!!value) {
			this._page4Title = value;
		}
	}
	get page4Title(): string | null {
		return this._page4Title;
	}
	private _page4Title: string | null = "Return submitted";

	@Input()
	set page1Help(value: string | null) {
		if (!!value) {
			this._page1Help = value;
		}
	}
	get page1Help(): string | null {
		return this._page1Help;
	}
	private _page1Help: string | null = `
		To return your order please add your order number
		and the email address used when creating the order.
		This information can be found on your order
		confirmation email.`;

	@Input()
	set page4Help(value: string | null) {
		if (!!value) {
			this._page4Help = value;
		}
	}
	get page4Help(): string | null {
		return this._page4Help;
	}
	private _page4Help: string | null = `Please check your email for further instructions.`;

	@Input()
	set page2HelpSelected(value: string | null) {
		if (!!value) {
			this._page2HelpSelected = value;
		}
	}
	get page2HelpSelected(): string | null {
		return this._page2HelpSelected;
	}
	private _page2HelpSelected: string | null = `You've selected`;

	@Input()
	set page2HelpSelectedItem(value: string | null) {
		if (!!value) {
			this._page2HelpSelectedItem = value;
		}
	}
	get page2HelpSelectedItem(): string | null {
		return this._page2HelpSelectedItem;
	}
	private _page2HelpSelectedItem: string | null = `Item`;

	@Input()
	set page2HelpSelectedItems(value: string | null) {
		if (!!value) {
			this._page2HelpSelectedItems = value;
		}
	}
	get page2HelpSelectedItems(): string | null {
		return this._page2HelpSelectedItems;
	}
	private _page2HelpSelectedItems: string | null = `Items`;

	state$: Observable<ReturnPortalFrameModel>;
	subscriptions$: Subscription[] = [];

	trackByIndex = (index: number, row: any) => {
		return JSON.stringify(row);
	}

	dialogRef: MatDialogRef<ReturnFrameErrorDialogComponent> | null = null;

	cmtFrmCtrl = new FormControl();

	address = "Store Torv 15, 8000 Aarhus"

	constructor(
		private store: Store,
		private route: ActivatedRoute,
		private actions$: Actions,
		private dialog: MatDialog,
		private viewContainerRef: ViewContainerRef,
		private renderer: Renderer2
	) {
		this.state$ = store.select(ReturnPortalFrameState.get);
	}

	ngAfterViewInit() {
		const styleElement = this.renderer.createElement('style');
		const styleText = `
		@font-face {
\tfont-family: "Material Symbols Outlined";
\tfont-style: normal;
\tfont-weight: 100 700;
\tfont-display: block;
\tsrc: url('./assets/material-symbols-outlined.woff2') format("woff2");
}
.material-symbols-outlined {
\tfont-family: "Material Symbols Outlined";
\tfont-weight: normal;
\tfont-style: normal;
\tfont-size: 24px;
\tline-height: 1;
\tletter-spacing: normal;
\ttext-transform: none;
\tdisplay: inline-block;
\twhite-space: nowrap;
\tword-wrap: normal;
\tdirection: ltr;
\t-webkit-font-smoothing: antialiased;
\t-moz-osx-font-smoothing: grayscale;
\ttext-rendering: optimizeLegibility;
\tfont-feature-settings: "liga";
}
		`;
		this.renderer.appendChild(document.head, styleElement);
		this.renderer.setProperty(styleElement, 'textContent', styleText);
	}

	showMaps() {
		const ref = this.dialog.open(StoreMapSelectorComponent, {viewContainerRef: this.viewContainerRef})
		ref.componentInstance.addressChanged.subscribe((newAddress) => {
			this.address = newAddress;
		})
	}

	ngOnInit() {
		this.subscriptions$.push(this.route.queryParams
			.subscribe((params) => {
				// Support both URL & component binding
				if (!!params.id && params.id.length > 0) {
					this.store.dispatch([
						new SetReturnPortalID(!!params.id ? params.id : ''),
					]);
				}
			}));

		this.subscriptions$.push(this.actions$.pipe(ofActionDispatched(ShowErrorDialog))
			.subscribe(({payload}) => {
				this.dialogRef = this.dialog.open(ReturnFrameErrorDialogComponent, {
					viewContainerRef: this.viewContainerRef});
				this.dialogRef.componentInstance.title = payload.title;
				this.dialogRef.componentInstance.body = payload.body;
			}));
		this.subscriptions$.push(this.cmtFrmCtrl.valueChanges.pipe(debounceTime(200))
			.subscribe((v) => {
				this.store.dispatch(new UpdateComment(v));
			}));

		this.subscriptions$.push(this.actions$.pipe(ofActionDispatched(LoadingRundown))
			.subscribe(() => {
				timer(250).subscribe(() => this.store.dispatch(new StopLoading()));
			}));
	}

	ngOnDestroy() {
		this.subscriptions$.forEach((s) => s.unsubscribe());
	}

	fetchView(email: string, orderPublicID: string) {
		this.store.dispatch([
			new SetOrderInfo({email, orderPublicID}),
			new FetchReturnPortalFrame(),
		]);
	}

	toggle(event: { item: Item; selected: boolean }) {
		this.store.dispatch(new SetSelectedItem(event))
	}

	totalCount(selectedItems: ItemReturn[]): number {
		let count = 0;
		selectedItems.forEach((i) => {
			if (i.selected) {
				count += i.quantity;
			}
		})
		return count;
	}

	mayContinue(selectedItems: ItemReturn[]): boolean {
		let mayContinue = false;
		selectedItems.some((i) => {
			let selectedQuantity = false;
			if (i.selected && i.quantity > 0) {
				selectedQuantity = true;
			}
			if (selectedQuantity && i.id.length > 0) {
				mayContinue = true;
				return true;
			}
			return false;
		});
		return mayContinue;
	}

	mayContinueDeliveryOptions(returnCollis: CreateReturnColliOutput[]): boolean {
		let foundIncomplete = false;
		returnCollis.forEach((c) => {
			if (c.selected_delivery_option_id.length <= 0) {
				foundIncomplete = true;
			}
		});
		return !foundIncomplete;
	}

	increment(event: {orderLineID: string}) {
		this.store.dispatch(new IncrementQuantity(event));
	}

	decrement(event: {orderLineID: string}) {
		this.store.dispatch(new DecrementQuantity(event));
	}

	reasonChanged(event: { item: Item; reasonID: string }) {
		this.store.dispatch(new SetSelectedItemReason(event));
	}

	createOrder() {
		this.store.dispatch(new CreateOrder());
	}

	selectDeliveryOption(returnColliID: string, deliveryOptionID: string) {
		this.store.dispatch(new SelectDeliveryOption({returnColliID, deliveryOptionID}));
	}

	submitDeliveryOptions() {
		this.store.dispatch(new SubmitDeliveryOptions());
	}

	isValidURL(tryURL: string): boolean | unknown {
		let url: URL;
		try {
			url = new URL(tryURL);
		} catch (e: unknown) {
			return e;
		}
		return url.protocol === "http:" || url.protocol === "https:";
	}

}
