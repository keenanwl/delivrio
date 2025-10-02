/* eslint-disable */
import * as Types from '../../../../../generated/graphql';

import { gql } from 'apollo-angular';
import { CarrierServiceItemFragmentDoc } from '../edit-common.generated';
import { Injectable } from '@angular/core';
import * as Apollo from 'apollo-angular';
export type FetchDeliveryOptionEditBringQueryVariables = Types.Exact<{
  id: Types.Scalars['ID'];
}>;


export type FetchDeliveryOptionEditBringQuery = { deliveryOptionBring?: { deliveryOption: { clickOptionDisplayCount?: number | null, name: string, description?: string | null, clickCollect?: boolean | null, overrideReturnAddress?: boolean | null, overrideSenderAddress?: boolean | null, hideDeliveryOption?: boolean | null, deliveryEstimateFrom?: number | null, deliveryEstimateTo?: number | null, carrierService: { id: string }, deliveryOptionBring?: { electronicCustoms: boolean } | null } } | null, carrierServices: { edges?: Array<{ node?: { id: string, label: string, return: boolean } | null } | null> | null } };

export type SaveDeliveryOptionEditBringMutationVariables = Types.Exact<{
  id: Types.Scalars['ID'];
  input: Types.UpdateDeliveryOptionBringInput;
  inputDeliveryOption: Types.UpdateDeliveryOptionInput;
}>;


export type SaveDeliveryOptionEditBringMutation = { updateDeliveryOptionBring: { id: string } };

export const FetchDeliveryOptionEditBringDocument = gql`
    query FetchDeliveryOptionEditBring($id: ID!) {
  deliveryOptionBring(id: $id) {
    deliveryOption {
      carrierService {
        id
      }
      clickOptionDisplayCount
      name
      description
      clickCollect
      overrideReturnAddress
      overrideSenderAddress
      hideDeliveryOption
      clickOptionDisplayCount
      deliveryEstimateFrom
      deliveryEstimateTo
      deliveryOptionBring {
        electronicCustoms
      }
    }
  }
  carrierServices(where: {hasCarrierBrandWith: {internalID: bring}}) {
    edges {
      node {
        ...CarrierServiceItem
      }
    }
  }
}
    ${CarrierServiceItemFragmentDoc}`;

  @Injectable({
    providedIn: 'root'
  })
  export class FetchDeliveryOptionEditBringGQL extends Apollo.Query<FetchDeliveryOptionEditBringQuery, FetchDeliveryOptionEditBringQueryVariables> {
    document = FetchDeliveryOptionEditBringDocument;
    
    constructor(apollo: Apollo.Apollo) {
      super(apollo);
    }
  }
export const SaveDeliveryOptionEditBringDocument = gql`
    mutation SaveDeliveryOptionEditBring($id: ID!, $input: UpdateDeliveryOptionBringInput!, $inputDeliveryOption: UpdateDeliveryOptionInput!) {
  updateDeliveryOptionBring(
    id: $id
    input: $input
    inputDeliveryOption: $inputDeliveryOption
  ) {
    id
  }
}
    `;

  @Injectable({
    providedIn: 'root'
  })
  export class SaveDeliveryOptionEditBringGQL extends Apollo.Mutation<SaveDeliveryOptionEditBringMutation, SaveDeliveryOptionEditBringMutationVariables> {
    document = SaveDeliveryOptionEditBringDocument;
    
    constructor(apollo: Apollo.Apollo) {
      super(apollo);
    }
  }