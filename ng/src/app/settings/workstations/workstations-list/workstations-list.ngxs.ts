import {Injectable} from "@angular/core";
import {Action, Selector, State, StateContext, Store} from "@ngxs/store";
import {WorkstationsListActions} from "./workstations-list.actions";
import WorkstationsResponse = WorkstationsListActions.WorkstationsResponse;
import {CreateWorkstationGQL, FetchWorkstationsGQL} from "./workstations-list.generated";
import SetRegistrationToken = WorkstationsListActions.SetRegistrationToken;
import FetchWorkstations = WorkstationsListActions.FetchWorkstations;

export interface WorkstationsModel {
	workstations: WorkstationsResponse[];
	registrationToken: string;
	registrationTokenImg: string;
	loading: boolean;
	showArchived: boolean;
}

const defaultState: WorkstationsModel = {
	workstations: [],
	registrationToken: '',
	registrationTokenImg: '',
	loading: false,
	showArchived: false,
};

@Injectable()
@State<WorkstationsModel>({
	name: 'workstationsList',
	defaults: defaultState,
})
export class WorkstationsListState {

	constructor(
		private list: FetchWorkstationsGQL,
		private create: CreateWorkstationGQL,
		private store: Store,
	) {
	}

	@Selector()
	static get(state: WorkstationsModel) {
		return state;
	}

	@Action(WorkstationsListActions.FetchWorkstations)
	FetchWorkstations(ctx: StateContext<WorkstationsModel>, action: WorkstationsListActions.FetchWorkstations) {
		ctx.patchState({loading: true});
		const state = ctx.getState();
		return this.list.fetch({showArchived: state.showArchived})
			.subscribe({next: (r) => {
				ctx.patchState({loading: false});
				ctx.dispatch(new WorkstationsListActions.SetWorkstations(r.data.filteredWorkstations));
			}});
	}

	@Action(WorkstationsListActions.SetWorkstations)
	SetWorkstations(ctx: StateContext<WorkstationsModel>, action: WorkstationsListActions.SetWorkstations) {
		ctx.patchState({workstations: action.payload})
	}

	@Action(WorkstationsListActions.SetRegistrationToken)
	SetRegistrationToken(ctx: StateContext<WorkstationsModel>, action: WorkstationsListActions.SetRegistrationToken) {
		ctx.patchState({registrationToken: action.payload.registrationToken, registrationTokenImg: action.payload.registrationTokenImg})
	}

	@Action(WorkstationsListActions.ToggleArchived)
	ToggleArchived(ctx: StateContext<WorkstationsModel>, action: WorkstationsListActions.ToggleArchived) {
		ctx.patchState({showArchived: !ctx.getState().showArchived});
		ctx.dispatch(new FetchWorkstations());
	}

	@Action(WorkstationsListActions.CreateNewWorkstation)
	CreateNewWorkstation(ctx: StateContext<WorkstationsModel>, action: WorkstationsListActions.CreateNewWorkstation) {
		ctx.patchState({loading: true});
		return this.create.mutate({
			input: {name: action.payload.name, deviceType: action.payload.deviceType},
		}).subscribe((r) => {
			ctx.patchState({loading: false});
			const resp = r.data?.createWorkstation;
			if (!!resp) {
				ctx.dispatch(new SetRegistrationToken({registrationToken: resp.registrationToken, registrationTokenImg: resp.registrationTokenImg}))
			}
		});
	}

}
