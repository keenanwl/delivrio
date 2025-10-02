/* eslint-disable */
import * as Types from '../../../../generated/graphql';

import { gql } from 'apollo-angular';
export type CarrierServiceItemFragment = { id: string, label: string, return: boolean };

export type BaseDeliveryOptionFragment = { name: string, description?: string | null, clickCollect?: boolean | null, overrideReturnAddress?: boolean | null, overrideSenderAddress?: boolean | null, hideDeliveryOption?: boolean | null, clickOptionDisplayCount?: number | null, deliveryEstimateFrom?: number | null, deliveryEstimateTo?: number | null, webshipperIntegration: boolean, webshipperID?: number | null, shipmondoIntegration: boolean, shipmondoDeliveryOption?: string | null, customsEnabled: boolean, customsSigner?: string | null, hideIfCompanyEmpty: boolean, carrierService: { id: string }, defaultPackaging?: { id: string, name: string } | null };

export const CarrierServiceItemFragmentDoc = gql`
    fragment CarrierServiceItem on CarrierService {
  id
  label
  return
}
    `;
export const BaseDeliveryOptionFragmentDoc = gql`
    fragment BaseDeliveryOption on DeliveryOption {
  carrierService {
    id
  }
  defaultPackaging {
    id
    name
  }
  name
  description
  clickCollect
  overrideReturnAddress
  overrideSenderAddress
  hideDeliveryOption
  clickOptionDisplayCount
  deliveryEstimateFrom
  deliveryEstimateTo
  webshipperIntegration
  webshipperID
  shipmondoIntegration
  shipmondoDeliveryOption
  customsEnabled
  customsSigner
  hideIfCompanyEmpty
}
    `;