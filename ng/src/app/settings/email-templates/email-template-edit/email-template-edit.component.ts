import {Component, OnDestroy, OnInit} from '@angular/core';
import {FormControl, FormGroup} from "@angular/forms";
import {Actions, ofActionCompleted, Store} from "@ngxs/store";
import {ActivatedRoute} from "@angular/router";
import {EmailTemplateEditActions} from "./email-template-edit.actions";
import FetchEmailTemplateEdit = EmailTemplateEditActions.FetchEmailTemplateEdit;
import SetEmailTemplateID = EmailTemplateEditActions.SetEmailTemplateID;
import Save = EmailTemplateEditActions.Save;
import {MatDialog} from "@angular/material/dialog";
import {TestEmailTemplateComponent} from "./dialogs/test-email-template/test-email-template.component";
import {debounceTime} from "rxjs/operators";
import {Observable, Subscription} from "rxjs";
import {EmailTemplateMergeType} from "../../../../generated/graphql";
import SetEmailTemplateEdit = EmailTemplateEditActions.SetEmailTemplateEdit;
import {EmailTemplatesListModel, EmailTemplatesListState} from "../email-templates-list/email-templates-list.ngxs";
import {EmailTemplateEditModel, EmailTemplateEditState} from "./email-template-edit.ngxs";

type varStatus = {[key: string]: boolean};

@Component({
	selector: 'app-email-template-edit',
	templateUrl: './email-template-edit.component.html',
	styleUrls: ['./email-template-edit.component.scss']
})
export class EmailTemplateEditComponent implements OnInit, OnDestroy {

	form = new FormGroup({
		name: new FormControl(""),
		subject: new FormControl(""),
		mergeType: new FormControl<EmailTemplateMergeType | null>(null),
		htmlTemplate: new FormControl(""),
	})

	orderLineRowVariables: varStatus = {
		"{{.ProductName}}": false,
		"{{.ProductVariantName}}": false,
		"{{.ProductFirstImageURL}}": false,
		"{{.Quantity}}": false,
		"{{.Price}}": false,
		"{{.Total}}": false,
		"{{.ReturnClaimName}}": false,
		"{{.ReturnClaimDescription}}": false,
	}

	returnLineOrderLineRowVariables: varStatus = {
		"{{.ReturnClaimName}}": false,
		"{{.ReturnClaimDescription}}": false,
	}

	returnOrderLineSection: varStatus = Object.assign(
		{"{{range .OrderLines}}": false},
		this.orderLineRowVariables,
		this.returnLineOrderLineRowVariables,
	{"{{end}}": false},
	);

	orderLineSection: varStatus = Object.assign(
		{"{{range .OrderLines}}": false},
		this.orderLineRowVariables,
		this.returnLineOrderLineRowVariables,
	{"{{end}}": false},
	);

	customerAddressVariables = {
		"{{.CustomerFirstName}}": false,
		"{{.CustomerLastName}}": false,
		"{{.CustomerCompany}}": false,
		"{{.CustomerEmail}}": false,
		"{{.CustomerAddress1}}": false,
		"{{.CustomerAddress2}}": false,
		"{{.CustomerZip}}": false,
		"{{.CustomerCity}}": false,
		"{{.CustomerState}}": false,
		"{{.CustomerCountry}}": false,
	};

	dropPointAddressVariables = {
		"{{.DropPointCompany}}": false,
		"{{.DropPointAddress1}}": false,
		"{{.DropPointAddress2}}": false,
		"{{.DropPointZip}}": false,
		"{{.DropPointCity}}": false,
		"{{.DropPointState}}": false,
		"{{.DropPointCountry}}": false,
	};

	orderPickedVariables: varStatus = Object.assign({
		"{{.TrackingID}}": false,
		"{{.OrderPublicID}}": false,
		"{{.CarrierName}}": false,
		"{{.ExpectedDeliveryDate.Format \"2006 Jan 02\"}}": false,
	}, this.customerAddressVariables,
		this.dropPointAddressVariables);

	returnBaseVariables: varStatus = Object.assign({
		"{{.ReturnFirstName}}": false,
		"{{.ReturnLastName}}": false,
		"{{.ReturnCompany}}": false,
		"{{.ReturnEmail}}": false,
		"{{.ReturnAddress1}}": false,
		"{{.ReturnAddress2}}": false,
		"{{.ReturnZip}}": false,
		"{{.ReturnCity}}": false,
		"{{.ReturnState}}": false,
		"{{.ReturnCountry}}": false,
		"{{.OrderPublicID}}": false,
		"{{.LabelDownloadURL}}": false,
		"{{.LabelPNGURL}}": false,
		"{{.ReturnMethodName}}": false,
		"{{.TrackingID}}": false,
	}, this.customerAddressVariables);

	returnLabelVariables: varStatus = Object.assign(
		this.returnBaseVariables,
		{
			"{{.LabelDownloadURL}}": false,
			"{{.LabelURL}}": false,
		},
	)

	returnQRCodeVariables: varStatus = Object.assign(
		this.returnBaseVariables,
		{
			"{{.QRCodeDownloadURL}}": false,
			"{{.QRCodeURL}}": false,
		},
	)

	singleVariables = this.returnBaseVariables;
	orderLineVariables = this.orderLineSection;

	subscriptions$: Subscription[] = [];

	mergeType = EmailTemplateMergeType

	state$: Observable<EmailTemplateEditModel>;

	constructor(
		private store: Store,
		private route: ActivatedRoute,
		private dialog: MatDialog,
		private actions$: Actions,
	) {
		this.state$ = store.select(EmailTemplateEditState.get);
	}

	ngOnInit() {
		// Yeah, it's a hack to do it this way
		this.form.controls.mergeType.disable();

		this.subscriptions$.push(this.route.queryParams
			.subscribe((params) => {
				this.store.dispatch([
					new SetEmailTemplateID(!!params.id ? params.id : ''),
					new FetchEmailTemplateEdit(),
				]);
			}));

		this.subscriptions$.push(this.actions$
			.pipe(ofActionCompleted(SetEmailTemplateEdit))
			.subscribe(() => this.checkHTMLChanges()));

		this.subscriptions$.push(this.form.controls.mergeType.valueChanges
			.subscribe(() => {
				this.checkHTMLChanges();
			}));

		this.subscriptions$.push(this.form.controls.htmlTemplate.valueChanges
			.pipe(debounceTime(500))
			.subscribe(() => {
				this.checkHTMLChanges();
			}));
	}

	checkHTMLChanges() {

		const html = this.form.controls.htmlTemplate.value?.
			replace(/{{\s/g, "{{").
			replace(/\s}}/g, "}}").
			replace(/\.Format\s"(.*)"}}/g, ".Format \"2006 Jan 02\"}}");

		this.singleVariables = this.returnBaseVariables;
		this.orderLineVariables = this.orderLineSection;
		switch (this.form.controls.mergeType.value) {
			case EmailTemplateMergeType.OrderPicked:
			case EmailTemplateMergeType.OrderConfirmation:
				this.singleVariables = this.orderPickedVariables;
				this.orderLineVariables = this.orderLineSection;
				break;
			case EmailTemplateMergeType.ReturnColliReceived:
				this.singleVariables = this.returnBaseVariables;
				this.orderLineVariables = this.returnOrderLineSection;
				break;
			case EmailTemplateMergeType.ReturnColliAccepted:
				this.singleVariables = this.returnBaseVariables;
				this.orderLineVariables = this.returnOrderLineSection;
				break;
			case EmailTemplateMergeType.ReturnColliQr:
				this.singleVariables = this.returnQRCodeVariables;
				this.orderLineVariables = this.returnOrderLineSection;
				break;
			case EmailTemplateMergeType.ReturnColliLabel:
				this.singleVariables = this.returnLabelVariables;
				this.orderLineVariables = this.returnOrderLineSection;
				break;
		}

		for (let mergeVariablesKey in this.singleVariables) {
			this.singleVariables[mergeVariablesKey] = !!html?.includes(mergeVariablesKey);
		}

		for (let orderLineKey in this.orderLineVariables) {
			this.orderLineVariables[orderLineKey] = !!html?.includes(orderLineKey);
		}

	}

	openTestDialog() {
		this.dialog.open(TestEmailTemplateComponent);
	}

	save() {
		this.store.dispatch(new Save());
	}

	ngOnDestroy(): void {
		this.subscriptions$.forEach((s) => s.unsubscribe());
	}

}
