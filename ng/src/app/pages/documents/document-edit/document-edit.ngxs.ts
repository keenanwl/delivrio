import {Injectable} from "@angular/core";
import {Action, Selector, State, StateContext} from "@ngxs/store";
import {formErrors} from "../../../account/company-info/company-info.ngxs";
import {produce} from "immer";
import DocumentResponse = DocumentEditActions.DocumentResponse;
import {AppActions} from "../../../app.actions";
import ShowGlobalSnackbar = AppActions.ShowGlobalSnackbar;
import AppChangeRoute = AppActions.AppChangeRoute;
import {Paths} from "../../../app-routing.module";
import {DocumentEditActions} from "./document-edit.actions";
import {DownloadDocumentGQL, FetchDocumentGQL, UpdateDocumentGQL} from "./document-edit.generated";
import CarrierBrandResponse = DocumentEditActions.CarrierBrandResponse;
import SetCarrierBrands = DocumentEditActions.SetCarrierBrands;
import {toNotNullArray} from "../../../functions/not-null-array";
import SetPDF = DocumentEditActions.SetPDF;

export interface DocumentEditModel {
	id: string;
	form: {
		model: DocumentResponse | undefined;
		dirty: boolean;
		status: string;
		errors: formErrors;
	}
	carrierBrands: CarrierBrandResponse[];
	base64PDF: string;
	loading: boolean;
}

const defaultState: DocumentEditModel = {
	id: "",
	form: {
		model: undefined,
		dirty: false,
		status: '',
		errors: {}
	},
	carrierBrands: [],
	base64PDF: '',
	loading: false,
};

@Injectable()
@State<DocumentEditModel>({
	name: 'documentEdit',
	defaults: defaultState,
})
export class DocumentEditState {

	constructor(
		private fetch: FetchDocumentGQL,
		private update: UpdateDocumentGQL,
		private download: DownloadDocumentGQL,
	) {
	}

	@Selector()
	static get(state: DocumentEditModel) {
		return state;
	}

	@Action(DocumentEditActions.FetchDocumentEdit)
	FetchDocumentEdit(ctx: StateContext<DocumentEditModel>, action: DocumentEditActions.FetchDocumentEdit) {
		return this.fetch.fetch({id: ctx.getState().id})
			.subscribe({next: (r) => {

				ctx.dispatch(new SetCarrierBrands(toNotNullArray(r.data.carrierBrands.edges?.map(n => n?.node))));

				const doc = r.data.document;
				if (!!doc) {
					ctx.dispatch(new DocumentEditActions.SetDocumentEdit(doc));
				}
			}});
	}

	@Action(DocumentEditActions.SetDocumentEdit)
	SetDocumentEdit(ctx: StateContext<DocumentEditModel>, action: DocumentEditActions.SetDocumentEdit) {
		const state = produce(ctx.getState(), st => {
			st.form.model = action.payload;
		});
		ctx.setState(state);
	}

	@Action(DocumentEditActions.SetDocumentID)
	SetDocumentID(ctx: StateContext<DocumentEditModel>, action: DocumentEditActions.SetDocumentID) {
		ctx.patchState({id: action.payload})
	}

	@Action(DocumentEditActions.Clear)
	Clear(ctx: StateContext<DocumentEditModel>, action: DocumentEditActions.Clear) {
		ctx.setState(defaultState);
	}

	@Action(DocumentEditActions.SetCarrierBrands)
	SetCarrierBrands(ctx: StateContext<DocumentEditModel>, action: DocumentEditActions.SetCarrierBrands) {
		ctx.patchState({carrierBrands: action.payload});
	}

	@Action(DocumentEditActions.SetDateTimeRange)
	SetDateTimeRange(ctx: StateContext<DocumentEditModel>, action: DocumentEditActions.SetDateTimeRange) {
		const state = produce(ctx.getState(), st => {
			st.form.model!.startAt = action.payload.start;
			st.form.model!.endAt = action.payload.end;
		});
		ctx.setState(state);
	}

	@Action(DocumentEditActions.SetPDF)
	SetPDF(ctx: StateContext<DocumentEditModel>, action: DocumentEditActions.SetPDF) {
		ctx.patchState({base64PDF: action.payload});
	}

	@Action(DocumentEditActions.Download)
	Download(ctx: StateContext<DocumentEditModel>, action: DocumentEditActions.Download) {
		ctx.patchState({loading: true, base64PDF: ""});
		const state = ctx.getState();
		return this.download.fetch({id: state.id})
			.subscribe((r) => {
				ctx.patchState({loading: false});
				ctx.dispatch(new SetPDF(r.data.documentDownload.base64PDF));
			});
	}

	@Action(DocumentEditActions.Save)
	Save(ctx: StateContext<DocumentEditModel>, action: DocumentEditActions.Save) {
		ctx.patchState({loading: true});
		const state = ctx.getState();
		const next = Object.assign({},
			state.form?.model,
			{carrierBrand: undefined},
			{carrierBrandID: state.form.model?.carrierBrand?.id},
		);
		return this.update.mutate({
			id: ctx.getState().id,
			input: next,
		}, {errorPolicy: "all"})
			.subscribe((r) => {
				ctx.patchState({loading: false});
				if (!!r.errors) {
					ctx.dispatch(new ShowGlobalSnackbar(`An error occurred: ${r.errors.map((e) => e.message).join(" ")}`));
				} else {
					ctx.dispatch(new AppChangeRoute({path: Paths.DOCUMENTS, queryParams: {}}));
				}
			});
	}

}
