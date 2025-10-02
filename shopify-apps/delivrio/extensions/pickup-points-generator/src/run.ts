import {
    BusinessHours,
    FunctionRunResult,
    Operation,
    PickupAddress,
    PickupPointDeliveryOption,
    Provider,
    RunInput
} from "../generated/api";

export function run(input: RunInput): FunctionRunResult {
    console.error(input)
    const { fetchResult } = input;
    const status = fetchResult?.status;
    const body = fetchResult?.body;

    let operations: Operation[] = [];

    if (status === 200 && body) {
        const g = JSON.parse(body);
        operations = buildPickupPointDeliveryOptionOperations(g);
    }

    return { operations };

}

function buildPickupPointDeliveryOptionOperations(externalApiDeliveryPoints: PickupPointDeliveryOption[]): Operation[] {
    return externalApiDeliveryPoints
        .map(externalApiDeliveryPoint => {
            const next = externalApiDeliveryPoint;
            next.cost = undefined;
            return { add: next }
        });
}

function buildPickupPointDeliveryOption(externalApiDeliveryPoint: any): PickupPointDeliveryOption {

    return {
        cost: 300_00,
        pickupPoint: {
            externalId: externalApiDeliveryPoint.pointId,
            name: externalApiDeliveryPoint.pointName,
            provider: buildProvider(),
            address: buildAddress(externalApiDeliveryPoint),
            businessHours: buildBusinessHours(externalApiDeliveryPoint),
        },
    };
}

function buildProvider(): Provider {
    return {
        name: "DELIVRIO",
        logoUrl: "https://cdn.shopify.com/s/files/1/0577/7778/2844/files/GLS_Icon_blue_DELIVRIO.webp?v=1712829523",
    };
}

function buildAddress(externalApiDeliveryPoint: any): PickupAddress {
    let location = externalApiDeliveryPoint.location;
    let addressComponents = location.addressComponents;
    let geometry = location.geometry.location;

    return {
        address1: `${addressComponents.streetNumber} ${addressComponents.route}`,
        address2: null,
        city: addressComponents.locality,
        country: addressComponents.country,
        countryCode: addressComponents.countryCode,
        latitude: geometry.lat,
        longitude: geometry.lng,
        phone: null,
        province: addressComponents.administrativeAreaLevel1,
        provinceCode: null,
        zip: addressComponents.postalCode,
    };
}

// Transforms the opening hours of a delivery point into a vector of `BusinessHours` objects.
// Each day's opening hours are represented using a `BusinessHours` object as follows:
// "Monday: 9:00 AM – 5:00 PM" is transformed to {day: "MONDAY", periods: [{opening_time: "09:00:00", closing_time: "17:00:00"}]}
// "Tuesday: Closed" is transformed to {day: "TUESDAY", periods: []}
function buildBusinessHours(externalApiDeliveryPoint: any): BusinessHours[] {
    return externalApiDeliveryPoint.openingHours.weekdayText
        .map((dayOpeningHours: string) => {
            let dayOpeningHoursParts = dayOpeningHours.split(": ");
            let dayName = dayOpeningHoursParts[0].toUpperCase();
            if (dayOpeningHoursParts[1] === "Closed") {
                return { day: dayName, periods: [] };
            } else {
                let openingClosingTimes = dayOpeningHoursParts[1].split(" – ");
                return {
                    day: dayName,
                    periods: [{
                        openingTime: formatTime(openingClosingTimes[0]),
                        closingTime: formatTime(openingClosingTimes[1]),
                    }],
                };
            }
        });
}

// Converts a time string from 12-hour to 24-hour format.
// Example: "9:00 AM" => "09:00:00", "5:00 PM" => "17:00:00"
function formatTime(time: string): string {
    let timeParts = time.split(' ');
    let hourMin = timeParts[0].split(':');
    let hour = parseInt(hourMin[0]);
    let min = hourMin[1];
    let period = timeParts[1];

    let hourIn24Format = period === 'AM' ? (hour === 12 ? 0 : hour) : (hour === 12 ? hour : hour + 12);

    return `${hourIn24Format.toString().padStart(2, '0')}:${min}:00`;
}
