/* eslint-disable */
import * as Types from '../../../../../generated/graphql';

import { gql } from 'apollo-angular';
import { BaseDeliveryOptionFragmentDoc, CarrierServiceItemFragmentDoc } from '../edit-common.generated';
import { Injectable } from '@angular/core';
import * as Apollo from 'apollo-angular';
export type FetchDeliveryOptionEditPostNordQueryVariables = Types.Exact<{
  id: Types.Scalars['ID'];
}>;


export type FetchDeliveryOptionEditPostNordQuery = { carrierAdditionalServicePostNords: { edges?: Array<{ node?: { internalID: string } | null } | null> | null }, clickCollectLocations?: { deliveryOption: { clickCollectLocation?: Array<{ id: string, name: string, address: { addressOne: string, addressTwo: string, zip: string, city: string, company?: string | null, country: { id: string, label: string, alpha2: string } } }> | null } } | null, selectedEmailTemplates?: { deliveryOption: { emailClickCollectAtStore?: { id: string } | null } } | null, deliveryOptionPostNord?: { deliveryOption: { name: string, description?: string | null, clickCollect?: boolean | null, overrideReturnAddress?: boolean | null, overrideSenderAddress?: boolean | null, hideDeliveryOption?: boolean | null, clickOptionDisplayCount?: number | null, deliveryEstimateFrom?: number | null, deliveryEstimateTo?: number | null, webshipperIntegration: boolean, webshipperID?: number | null, shipmondoIntegration: boolean, shipmondoDeliveryOption?: string | null, customsEnabled: boolean, customsSigner?: string | null, hideIfCompanyEmpty: boolean, carrierService: { id: string }, defaultPackaging?: { id: string, name: string } | null } } | null, locations: { edges?: Array<{ node?: { id: string, name: string, address: { addressOne: string, addressTwo: string, zip: string, city: string, company?: string | null, country: { id: string, label: string, alpha2: string } } } | null } | null> | null }, carrierServices: { edges?: Array<{ node?: { id: string, label: string, return: boolean } | null } | null> | null }, emailTemplates: { edges?: Array<{ node?: { id: string, name: string, mergeType: Types.EmailTemplateMergeType } | null } | null> | null } };

export type AvailableAdditionalServicesPostNordQueryVariables = Types.Exact<{
  carrierServiceID: Types.Scalars['ID'];
}>;


export type AvailableAdditionalServicesPostNordQuery = { availableAdditionalServicesPostNord: Array<string> };

export type SaveDeliveryOptionEditPostNordMutationVariables = Types.Exact<{
  id: Types.Scalars['ID'];
  input: Types.UpdateDeliveryOptionPostNordInput;
  inputDeliveryOption: Types.UpdateDeliveryOptionInput;
  inputAdditionalServices: Array<Types.Scalars['String']> | Types.Scalars['String'];
}>;


export type SaveDeliveryOptionEditPostNordMutation = { updateDeliveryOptionPostNord: { id: string } };

export const FetchDeliveryOptionEditPostNordDocument = gql`
    query FetchDeliveryOptionEditPostNord($id: ID!) {
  carrierAdditionalServicePostNords(
    where: {hasDeliveryOptionPostNordWith: {hasDeliveryOptionWith: {id: $id}}}
  ) {
    edges {
      node {
        internalID
      }
    }
  }
  clickCollectLocations: deliveryOptionPostNord(deliveryOptionID: $id) {
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
  selectedEmailTemplates: deliveryOptionPostNord(deliveryOptionID: $id) {
    deliveryOption {
      emailClickCollectAtStore {
        id
      }
    }
  }
  deliveryOptionPostNord(deliveryOptionID: $id) {
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
  carrierServices(where: {hasCarrierBrandWith: {internalID: post_nord}}) {
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
    ${BaseDeliveryOptionFragmentDoc}
${CarrierServiceItemFragmentDoc}`;

  @Injectable({
    providedIn: 'root'
  })
  export class FetchDeliveryOptionEditPostNordGQL extends Apollo.Query<FetchDeliveryOptionEditPostNordQuery, FetchDeliveryOptionEditPostNordQueryVariables> {
    document = FetchDeliveryOptionEditPostNordDocument;
    
    constructor(apollo: Apollo.Apollo) {
      super(apollo);
    }
  }
export const AvailableAdditionalServicesPostNordDocument = gql`
    query AvailableAdditionalServicesPostNord($carrierServiceID: ID!) {
  availableAdditionalServicesPostNord(carrierServiceID: $carrierServiceID)
}
    `;

  @Injectable({
    providedIn: 'root'
  })
  export class AvailableAdditionalServicesPostNordGQL extends Apollo.Query<AvailableAdditionalServicesPostNordQuery, AvailableAdditionalServicesPostNordQueryVariables> {
    document = AvailableAdditionalServicesPostNordDocument;
    
    constructor(apollo: Apollo.Apollo) {
      super(apollo);
    }
  }
export const SaveDeliveryOptionEditPostNordDocument = gql`
    mutation SaveDeliveryOptionEditPostNord($id: ID!, $input: UpdateDeliveryOptionPostNordInput!, $inputDeliveryOption: UpdateDeliveryOptionInput!, $inputAdditionalServices: [String!]!) {
  updateDeliveryOptionPostNord(
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
  export class SaveDeliveryOptionEditPostNordGQL extends Apollo.Mutation<SaveDeliveryOptionEditPostNordMutation, SaveDeliveryOptionEditPostNordMutationVariables> {
    document = SaveDeliveryOptionEditPostNordDocument;
    
    constructor(apollo: Apollo.Apollo) {
      super(apollo);
    }
  }