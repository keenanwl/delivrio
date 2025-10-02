import {Injectable} from "@angular/core";
import { HttpClient, HttpParams } from "@angular/common/http";
import {Observable} from "rxjs";

export interface ReturnPortalViewResponse {
	packages: Package[];
	order_date: Date;
	order_id: string;
	return_reasons: ReturnReason[];
}

export interface Package {
	items: Item[];
}

export interface Item {
	order_line_id: string;
	name: string;
	variant_name: string;
	quantity: number;
	image_url: string;
}

export interface ReturnReason {
	id: string;
	name: string;
	description: string;
}

export interface CreateReturnOrderItemInput {
	claim_id: string;
	order_line_id: string;
	units: number;
}

export interface CreateReturnOrderInput {
	portal_id: string;
	order_lines: CreateReturnOrderItemInput[];
	comment: string;
}

export interface CreateReturnColliDeliveryOption {
	delivery_option_id: string;
	logo_url: string;
	name: string;
	description: string;
	formatted_price: string;
}

export interface CreateReturnColliOutput {
	return_colli_id: string;
	selected_delivery_option_id: string;
	available_delivery_options: CreateReturnColliDeliveryOption[];
}

export interface CreateReturnOrderOutput {
	return_collis: CreateReturnColliOutput[];
}

export interface AddReturnOrderDeliveryOptionsColliIDs {
	return_colli_id: string;
	delivery_option_id: string;
}

export interface AddReturnOrderDeliveryOptionsInput {
	delivery_options: AddReturnOrderDeliveryOptionsColliIDs[];
}

export interface AddReturnOrderDeliveryOptionsOutput {
	success: boolean;
}

export interface GenericError {
	message: string;
}

@Injectable()
export class ReturnPortalFrameService {

	public baseURL = "";

	constructor(
		private http: HttpClient,
	) {}

	public getOrderOverview(
		portalID: string,
		email: string,
		orderPublicID: string
	): Observable<ReturnPortalViewResponse> {

		const params = new HttpParams({fromObject: {
			"return-portal-id": portalID,
			"email": email,
			"order-public-id": orderPublicID,
		}});

		return this.http.get<ReturnPortalViewResponse>(this.baseURL + `/api/return-view`, {params: params})
	}

	public createOrder(
		portalID: string,
	    email: string,
	    orderPublicID: string,
		body: CreateReturnOrderInput
	): Observable<CreateReturnOrderOutput> {
		const params = new HttpParams({fromObject: {
			"return-portal-id": portalID,
			"email": email,
			"order-public-id": orderPublicID,
		}});
		return this.http.post<CreateReturnOrderOutput>(this.baseURL + `/api/return-create`, body, {params})
	}

	public submitDeliveryOptions(
		portalID: string,
	    email: string,
	    orderPublicID: string,
		body: AddReturnOrderDeliveryOptionsInput
	): Observable<AddReturnOrderDeliveryOptionsOutput> {
		const params = new HttpParams({fromObject: {
			"return-portal-id": portalID,
			"email": email,
			"order-public-id": orderPublicID,
		}});
		return this.http.post<AddReturnOrderDeliveryOptionsOutput>(this.baseURL + `/api/return-delivery-options`, body, {params})
	}

}
