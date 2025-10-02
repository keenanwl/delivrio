import {
    Maybe,
    Allocation,
    MailingAddress,
    CountryCode,
    HttpRequest,
    HttpRequestMethod,
    FunctionFetchResult,
    FetchInput
} from '../generated/api';

export function fetch(input: FetchInput): FunctionFetchResult {
    return {
        request: buildExternalApiRequest(input),
    };
}

function buildExternalApiRequest(input: FetchInput): HttpRequest {
    // The latitude and longitude parameters are included in the URL for demonstration purposes only. They do not influence the result.
    let url = `https://main.delivrio.io/api/shopify-lookup-pickup-points`;

    return {
        method: HttpRequestMethod.Post,
        url,
        headers: [{
            name: "Accept",
            value: "application/json; charset=utf-8"
        }],
        body: JSON.stringify(input),
        policy: {
            readTimeoutMs: 2_000,
        },
    };

}

function getUniformDeliveryAddress(allocations: Allocation[]): Maybe<MailingAddress> {
    if (allocations.length === 0) {
        return null;
    }

    let deliveryAddress = allocations[0].deliveryAddress;

    for (let i = 1; i < allocations.length; i++) {
        if (!isDeliveryAddressEqual(allocations[i].deliveryAddress, deliveryAddress)) {
            return null;
        }
    }

    return deliveryAddress;
}

function isDeliveryAddressEqual(address1: MailingAddress, address2: MailingAddress): boolean {
    return address1.countryCode === address2.countryCode &&
        address1.longitude === address2.longitude &&
        address1.latitude === address2.latitude;
}
