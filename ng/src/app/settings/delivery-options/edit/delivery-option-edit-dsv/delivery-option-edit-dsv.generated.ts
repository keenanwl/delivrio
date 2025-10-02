/* eslint-disable */
import * as Types from '../../../../../generated/graphql';

import { gql } from 'apollo-angular';
import { CarrierServiceItemFragmentDoc } from '../edit-common.generated';
import { Injectable } from '@angular/core';
import * as Apollo from 'apollo-angular';
export type FetchDeliveryOptionEditDsvQueryVariables = Types.Exact<{
  id: Types.Scalars['ID'];
}>;


export type FetchDeliveryOptionEditDsvQuery = { deliveryOptionDSV?: { deliveryOption: { clickOptionDisplayCount?: number | null, name: string, description?: string | null, clickCollect?: boolean | null, overrideReturnAddress?: boolean | null, overrideSenderAddress?: boolean | null, hideDeliveryOption?: boolean | null, deliveryEstimateFrom?: number | null, deliveryEstimateTo?: number | null, carrierService: { id: string } } } | null, carrierServices: { edges?: Array<{ node?: { id: string, label: string, return: boolean } | null } | null> | null } };

export type SaveDeliveryOptionEditDsvMutationVariables = Types.Exact<{
  id: Types.Scalars['ID'];
  inputDeliveryOption: Types.UpdateDeliveryOptionInput;
}>;


export type SaveDeliveryOptionEditDsvMutation = { updateDeliveryOptionDSV: { id: string } };

export const FetchDeliveryOptionEditDsvDocument = gql`
    query FetchDeliveryOptionEditDSV($id: ID!) {
  deliveryOptionDSV(id: $id) {
    deliveryOption {
      carrierService {
        id
      }
      clickOptionDisplayCount
    }
    deliveryOption {
      name
      description
      clickCollect
      overrideReturnAddress
      overrideSenderAddress
      hideDeliveryOption
      clickOptionDisplayCount
      deliveryEstimateFrom
      deliveryEstimateTo
    }
  }
  carrierServices(where: {hasCarrierBrandWith: {internalID: dsv}}) {
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
  export class FetchDeliveryOptionEditDsvGQL extends Apollo.Query<FetchDeliveryOptionEditDsvQuery, FetchDeliveryOptionEditDsvQueryVariables> {
    document = FetchDeliveryOptionEditDsvDocument;
    
    constructor(apollo: Apollo.Apollo) {
      super(apollo);
    }
  }
export const SaveDeliveryOptionEditDsvDocument = gql`
    mutation SaveDeliveryOptionEditDSV($id: ID!, $inputDeliveryOption: UpdateDeliveryOptionInput!) {
  updateDeliveryOptionDSV(id: $id, inputDeliveryOption: $inputDeliveryOption) {
    id
  }
}
    `;

  @Injectable({
    providedIn: 'root'
  })
  export class SaveDeliveryOptionEditDsvGQL extends Apollo.Mutation<SaveDeliveryOptionEditDsvMutation, SaveDeliveryOptionEditDsvMutationVariables> {
    document = SaveDeliveryOptionEditDsvDocument;
    
    constructor(apollo: Apollo.Apollo) {
      super(apollo);
    }
  }