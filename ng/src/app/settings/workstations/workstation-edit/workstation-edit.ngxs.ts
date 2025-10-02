import {Injectable} from "@angular/core";
import {Action, Selector, State, StateContext, Store} from "@ngxs/store";
import {WorkstationEditActions} from "./workstation-edit.actions";
import {formErrors} from "../../../account/company-info/company-info.ngxs";
import {ArchiveWorkstationGQL, FetchWorkstationGQL, UpdateWorkstationGQL} from "./workstation-edit.generated";
import {produce} from "immer";
import {AppActions} from "../../../app.actions";
import {UpdatePrinterWithIdInput, UpdateWorkstationInput, WorkstationStatus} from "../../../../generated/graphql";
import {Paths} from "../../../app-routing.module";
import SetWorkstationEdit = WorkstationEditActions.SetWorkstationEdit;
import WorkstationEditResponse = WorkstationEditActions.WorkstationEditResponse;
import ShowGlobalSnackbar = AppActions.ShowGlobalSnackbar;
import AppChangeRoute = AppActions.AppChangeRoute;

export interface WorkstationEditModel {
	workstationEditForm: {
		model: WorkstationEditResponse | undefined;
		dirty: boolean;
		status: string;
		errors: formErrors;
	},
	workstationID: string;
}

const defaultState: WorkstationEditModel = {
	workstationEditForm: {
		model: undefined,
		dirty: false,
		status: '',
		errors: {}
	},
	workstationID: '',
};

@Injectable()
@State<WorkstationEditModel>({
	name: 'workstationEdit',
	defaults: defaultState,
})
export class WorkstationEditState {

	constructor(
		private fetchWorkstation: FetchWorkstationGQL,
		private updateWorkstation: UpdateWorkstationGQL,
		private archive: ArchiveWorkstationGQL,
		private store: Store,
	) {}

	@Selector()
	static get(state: WorkstationEditModel) {
		return state;
	}

	@Action(WorkstationEditActions.FetchWorkstationEdit)
	FetchMyWorkstationEdit(ctx: StateContext<WorkstationEditModel>, action: WorkstationEditActions.FetchWorkstationEdit) {
		const id = ctx.getState().workstationID;
		return this.fetchWorkstation.fetch({id})
			.subscribe({next: (r) => {
				const workstation = r.data.workstation;
				if (!!workstation) {
					ctx.dispatch(new SetWorkstationEdit(workstation));
				}
			}});
	}

	@Action(WorkstationEditActions.SetWorkstationID)
	SetWorkstationID(ctx: StateContext<WorkstationEditModel>, action: WorkstationEditActions.SetWorkstationID) {
		ctx.patchState({workstationID: action.payload});
	}

	@Action(WorkstationEditActions.SetWorkstationEdit)
	SetWorkstationEdit(ctx: StateContext<WorkstationEditModel>, action: WorkstationEditActions.SetWorkstationEdit) {
		const state = produce(ctx.getState(), st => {
			st.workstationEditForm.model = action.payload;
		});
		ctx.setState(state);
	}

	@Action(WorkstationEditActions.Save)
	Save(ctx: StateContext<WorkstationEditModel>, action: WorkstationEditActions.Save) {
		const state = ctx.getState();

		const printers: UpdatePrinterWithIdInput[] = (state.workstationEditForm.model?.printer || [])
			.map((p) => {
				return {
					id: p.id,
					updatePrinters: Object.assign({}, p, {id: undefined, lastPing: undefined}),
				}
			});

		const next: UpdateWorkstationInput = Object.assign({},
			state.workstationEditForm.model,
			{
				lastPing: undefined,
				printer: undefined,
				updatePrinters: printers,
			}
		)

		return this.updateWorkstation.mutate({id: state.workstationID, input: next})
			.subscribe((r) => {
				this.store.dispatch([
					new ShowGlobalSnackbar("Workstation saved successfully"),
					new AppChangeRoute({path: Paths.SETTINGS_WORKSTATIONS, queryParams: {}}),
				]);
			});

	}

	@Action(WorkstationEditActions.Disable)
	Disable(ctx: StateContext<WorkstationEditModel>, action: WorkstationEditActions.Disable) {
		const state = ctx.getState();
		return this.archive.mutate({id: state.workstationID})
			.subscribe((r) => {
				this.store.dispatch([
					new ShowGlobalSnackbar("Workstation archived successfully"),
					new AppChangeRoute({path: Paths.SETTINGS_WORKSTATIONS, queryParams: {}}),
				]);
			});
	}

	@Action(WorkstationEditActions.Reset)
	Reset(ctx: StateContext<WorkstationEditModel>, action: WorkstationEditActions.Reset) {
		ctx.setState(defaultState);
	}

}
