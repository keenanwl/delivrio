import {Injectable} from "@angular/core";
import {Action, Selector, State, StateContext} from "@ngxs/store";
import DocumentsResponse = DocumentsListActions.DocumentsResponse;
import {CreateDocumentGQL, FetchDocumentsGQL} from "./documents-list.generated";
import {DocumentsListActions} from "./documents-list.actions";
import {AppActions} from "../../../app.actions";
import ShowGlobalSnackbar = AppActions.ShowGlobalSnackbar;
import AppChangeRoute = AppActions.AppChangeRoute;
import {Paths} from "../../../app-routing.module";
import {toNotNullArray} from "../../../functions/not-null-array";
import SetDocumentsList = DocumentsListActions.SetDocumentsList;

export interface DocumentsListModel {
	documentsList: DocumentsResponse[];
	loading: boolean;
}

const defaultState: DocumentsListModel = {
	documentsList: [],
	loading: false,
};

@Injectable()
@State<DocumentsListModel>({
	name: 'documentsList',
	defaults: defaultState,
})
export class DocumentsListState {

	constructor(
		private list: FetchDocumentsGQL,
		private create: CreateDocumentGQL,
	) {
	}

	@Selector()
	static get(state: DocumentsListModel) {
		return state;
	}

	@Action(DocumentsListActions.FetchDocumentsList)
	FetchMyDocumentsList(ctx: StateContext<DocumentsListModel>, action: DocumentsListActions.FetchDocumentsList) {
		ctx.patchState({loading: true});
		return this.list.fetch({}, {errorPolicy: "all"})
			.subscribe({
				next: (r) => {
					ctx.patchState({loading: false});
					const documents = toNotNullArray(r.data.documents.edges?.map((l) => l?.node));
					ctx.dispatch(new SetDocumentsList(documents));
				},
				error: () => {
					ctx.patchState({loading: false});
				},
			});
	}

	@Action(DocumentsListActions.SetDocumentsList)
	SetMyDocumentsList(ctx: StateContext<DocumentsListModel>, action: DocumentsListActions.SetDocumentsList) {
		ctx.patchState({documentsList: action.payload})
	}

	@Action(DocumentsListActions.Clear)
	Clear(ctx: StateContext<DocumentsListModel>, action: DocumentsListActions.Clear) {
		ctx.setState(defaultState);
	}

	@Action(DocumentsListActions.Create)
	Create(ctx: StateContext<DocumentsListModel>, action: DocumentsListActions.Create) {
		ctx.patchState({loading: true});
		return this.create.mutate({name: action.payload.name, mergeType: action.payload.mergeType}, {errorPolicy: "all"})
			.subscribe({
				next: (r) => {
					ctx.patchState({loading: false});
					if (!!r.errors) {
						ctx.dispatch(new ShowGlobalSnackbar("An error occurred"));
					} else {
						ctx.dispatch(new AppChangeRoute({path: Paths.DOCUMENT_EDIT, queryParams: {id: r.data?.createDocument}}));
					}
				},
			});
	}

}
