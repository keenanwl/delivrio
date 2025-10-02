import {Injectable} from "@angular/core";
import {Action, Selector, State, StateContext, Store} from "@ngxs/store";
import {UserSignupOptions} from "../../shared/models/user_signup_options";
import {Register2Actions} from "./register2.actions";
import SaveRegistration2 = Register2Actions.SaveRegistration2;
import {AppActions} from "../../app.actions";
import AppChangeRoute = AppActions.AppChangeRoute;
import SetInvalidMessage = Register2Actions.SetInvalidMessage;
import {CreateSignupOptionsInput} from "../../../generated/graphql";
import {GraphQLError} from "graphql/index";
import {AppState} from "../../app.ngxs";
import {ReplaceSignupOptionsGQL} from "./register2.generated";

export interface Register2Model {
	register2EditForm: {
		model: CreateSignupOptionsInput | undefined;
		dirty: boolean;
		status: string;
		errors: readonly GraphQLError[];
	},
	invalid_message: string;
	users_id: string;
	options: UserSignupOptions;
}

const defaultState: Register2Model = {
	register2EditForm: {
		model: undefined,
		dirty: false,
		status: '',
		errors: [],
	},
	invalid_message: "",
	users_id: "",
	options: {
		better_delivery_options: false,
		click_collect: false,
		custom_docs: false,
		easy_returns: false,
		improve_pick_pack: false,
		num_shipments: 100,
		reduced_costs: true,
		shipping_label: false,
	}
};

@Injectable()
@State<Register2Model>({
	name: 'register2',
	defaults: defaultState,
})
export class Register2State {

	constructor(
	    private replaceSignupOptions: ReplaceSignupOptionsGQL,
	    private store: Store,
	) { }

	@Selector()
	static state(state: Register2Model) {
		return state;
	}

	@Action(SaveRegistration2)
	SaveRegistration2(ctx: StateContext<Register2Model>, action: SaveRegistration2) {
		const myID = this.store.selectSnapshot(AppState.get).my_ids.my_pulid;
		this.replaceSignupOptions.mutate({userID: myID, input: action.payload}).subscribe(
		    () => {
		        ctx.dispatch(new AppChangeRoute({path: `/register/3`, queryParams: {}}))
		    }, (e) => ctx.dispatch(new SetInvalidMessage(e)));
	}

}
