/* eslint-disable */
import * as Types from '../../../../../generated/graphql';

import { gql } from 'apollo-angular';
import { BaseDeliveryOptionFragmentDoc, CarrierServiceItemFragmentDoc } from '../edit-common.generated';
import { Injectable } from '@angular/core';
import * as Apollo from 'apollo-angular';
export type FetchDeliveryOptionEditUspsQueryVariables = Types.Exact<{
  id: Types.Scalars['ID'];
}>;


export type FetchDeliveryOptionEditUspsQuery = { carrierAdditionalServiceUspSs: { edges?: Array<{ node?: { id: string, internalID: Types.CarrierAdditionalServiceUspsInternalId, label: string, commonlyUsed: boolean } | null } | null> | null }, clickCollectLocations?: { deliveryOption: { clickCollectLocation?: Array<{ id: string, name: string, address: { addressOne: string, addressTwo: string, zip: string, city: string, company?: string | null, country: { id: string, label: string, alpha2: string } } }> | null } } | null, selectedEmailTemplates?: { deliveryOption: { emailClickCollectAtStore?: { id: string } | null } } | null, deliveryOptionUSPS?: { deliveryOption: { name: string, description?: string | null, clickCollect?: boolean | null, overrideReturnAddress?: boolean | null, overrideSenderAddress?: boolean | null, hideDeliveryOption?: boolean | null, clickOptionDisplayCount?: number | null, deliveryEstimateFrom?: number | null, deliveryEstimateTo?: number | null, webshipperIntegration: boolean, webshipperID?: number | null, shipmondoIntegration: boolean, shipmondoDeliveryOption?: string | null, customsEnabled: boolean, customsSigner?: string | null, hideIfCompanyEmpty: boolean, carrierService: { id: string }, defaultPackaging?: { id: string, name: string } | null } } | null, locations: { edges?: Array<{ node?: { id: string, name: string, address: { addressOne: string, addressTwo: string, zip: string, city: string, company?: string | null, country: { id: string, label: string, alpha2: string } } } | null } | null> | null }, carrierServices: { edges?: Array<{ node?: { id: string, label: string, return: boolean } | null } | null> | null }, emailTemplates: { edges?: Array<{ node?: { id: string, name: string, mergeType: Types.EmailTemplateMergeType } | null } | null> | null } };

export type UspsAdditionalServicesFragment = { id: string, internalID: Types.CarrierAdditionalServiceUspsInternalId, label: string, commonlyUsed: boolean };

export type AvailableAdditionalServicesUspsQueryVariables = Types.Exact<{
  carrierServiceID: Types.Scalars['ID'];
}>;


export type AvailableAdditionalServicesUspsQuery = { availableAdditionalServicesUSPS: Array<{ id: string, internalID: Types.CarrierAdditionalServiceUspsInternalId, label: string, commonlyUsed: boolean }> };

export type SaveDeliveryOptionEditUspsMutationVariables = Types.Exact<{
  id: Types.Scalars['ID'];
  input: Types.UpdateDeliveryOptionUspsInput;
  inputDeliveryOption: Types.UpdateDeliveryOptionInput;
  inputAdditionalServices: Array<Types.Scalars['ID']> | Types.Scalars['ID'];
}>;


export type SaveDeliveryOptionEditUspsMutation = { updateDeliveryOptionUSPS: { id: string } };

export const UspsAdditionalServicesFragmentDoc = gql`
    fragment USPSAdditionalServices on CarrierAdditionalServiceUSPS {
  id
  internalID
  label
  commonlyUsed
}
    `;
export const FetchDeliveryOptionEditUspsDocument = gql`
    query FetchDeliveryOptionEditUSPS($id: ID!) {
  carrierAdditionalServiceUspSs(
    where: {hasDeliveryOptionUSPSWith: {hasDeliveryOptionWith: {id: $id}}}
  ) {
    edges {
      node {
        ...USPSAdditionalServices
      }
    }
  }
  clickCollectLocations: deliveryOptionUSPS(deliveryOptionID: $id) {
    deliveryOption {
      clickCollectLocation {
        id
        name
        address {
          addressOne
          addressTwo
          zip
          city
          country {
            id
            label
            alpha2
          }
          company
        }
      }
    }
  }
  selectedEmailTemplates: deliveryOptionUSPS(deliveryOptionID: $id) {
    deliveryOption {
      emailClickCollectAtStore {
        id
      }
    }
  }
  deliveryOptionUSPS(deliveryOptionID: $id) {
    deliveryOption {
      ...BaseDeliveryOption
    }
  }
  locations(
    where: {hasLocationTagsWith: {internalIDContainsFold: "click_and_collect"}}
  ) {
    edges {
      node {
        id
        name
        address {
          addressOne
          addressTwo
          zip
          city
          country {
            id
            label
            alpha2
          }
          company
        }
      }
    }
  }
  carrierServices(where: {hasCarrierBrandWith: {internalID: usps}}) {
    edges {
      node {
        ...CarrierServiceItem
      }
    }
  }
  emailTemplates {
    edges {
      node {
        id
        name
        mergeType
      }
    }
  }
}
    ${UspsAdditionalServicesFragmentDoc}
${BaseDeliveryOptionFragmentDoc}
${CarrierServiceItemFragmentDoc}`;

  @Injectable({
    providedIn: 'root'
  })
  export class FetchDeliveryOptionEditUspsGQL extends Apollo.Query<FetchDeliveryOptionEditUspsQuery, FetchDeliveryOptionEditUspsQueryVariables> {
    document = FetchDeliveryOptionEditUspsDocument;
    
    constructor(apollo: Apollo.Apollo) {
      super(apollo);
    }
  }
export const AvailableAdditionalServicesUspsDocument = gql`
    query AvailableAdditionalServicesUSPS($carrierServiceID: ID!) {
  availableAdditionalServicesUSPS(carrierServiceID: $carrierServiceID) {
    ...USPSAdditionalServices
  }
}
    ${UspsAdditionalServicesFragmentDoc}`;

  @Injectable({
    providedIn: 'root'
  })
  export class AvailableAdditionalServicesUspsGQL extends Apollo.Query<AvailableAdditionalServicesUspsQuery, AvailableAdditionalServicesUspsQueryVariables> {
    document = AvailableAdditionalServicesUspsDocument;
    
    constructor(apollo: Apollo.Apollo) {
      super(apollo);
    }
  }
export const SaveDeliveryOptionEditUspsDocument = gql`
    mutation SaveDeliveryOptionEditUSPS($id: ID!, $input: UpdateDeliveryOptionUSPSInput!, $inputDeliveryOption: UpdateDeliveryOptionInput!, $inputAdditionalServices: [ID!]!) {
  updateDeliveryOptionUSPS(
    id: $id
    input: $input
    inputDeliveryOption: $inputDeliveryOption
    inputAdditionalServices: $inputAdditionalServices
  ) {
    id
  }
}
    `;

  @Injectable({
    providedIn: 'root'
  })
  export class SaveDeliveryOptionEditUspsGQL extends Apollo.Mutation<SaveDeliveryOptionEditUspsMutation, SaveDeliveryOptionEditUspsMutationVariables> {
    document = SaveDeliveryOptionEditUspsDocument;
    
    constructor(apollo: Apollo.Apollo) {
      super(apollo);
    }
  }