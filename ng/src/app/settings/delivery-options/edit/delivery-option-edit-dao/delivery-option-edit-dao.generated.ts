/* eslint-disable */
import * as Types from '../../../../../generated/graphql';

import { gql } from 'apollo-angular';
import { CarrierServiceItemFragmentDoc } from '../edit-common.generated';
import { Injectable } from '@angular/core';
import * as Apollo from 'apollo-angular';
export type FetchDeliveryOptionEditDaoQueryVariables = Types.Exact<{
  id: Types.Scalars['ID'];
}>;


export type FetchDeliveryOptionEditDaoQuery = { deliveryOptionDAO?: { deliveryOption: { clickOptionDisplayCount?: number | null, name: string, description?: string | null, clickCollect?: boolean | null, overrideReturnAddress?: boolean | null, overrideSenderAddress?: boolean | null, hideDeliveryOption?: boolean | null, deliveryEstimateFrom?: number | null, deliveryEstimateTo?: number | null, carrierService: { id: string } } } | null, carrierServices: { edges?: Array<{ node?: { id: string, label: string, return: boolean } | null } | null> | null } };

export type SaveDeliveryOptionEditDaoMutationVariables = Types.Exact<{
  id: Types.Scalars['ID'];
  inputDeliveryOption: Types.UpdateDeliveryOptionInput;
}>;


export type SaveDeliveryOptionEditDaoMutation = { updateDeliveryOptionDAO: { id: string } };

export const FetchDeliveryOptionEditDaoDocument = gql`
    query FetchDeliveryOptionEditDAO($id: ID!) {
  deliveryOptionDAO(id: $id) {
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
  carrierServices(where: {hasCarrierBrandWith: {internalID: dao}}) {
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
  export class FetchDeliveryOptionEditDaoGQL extends Apollo.Query<FetchDeliveryOptionEditDaoQuery, FetchDeliveryOptionEditDaoQueryVariables> {
    document = FetchDeliveryOptionEditDaoDocument;
    
    constructor(apollo: Apollo.Apollo) {
      super(apollo);
    }
  }
export const SaveDeliveryOptionEditDaoDocument = gql`
    mutation SaveDeliveryOptionEditDAO($id: ID!, $inputDeliveryOption: UpdateDeliveryOptionInput!) {
  updateDeliveryOptionDAO(id: $id, inputDeliveryOption: $inputDeliveryOption) {
    id
  }
}
    `;

  @Injectable({
    providedIn: 'root'
  })
  export class SaveDeliveryOptionEditDaoGQL extends Apollo.Mutation<SaveDeliveryOptionEditDaoMutation, SaveDeliveryOptionEditDaoMutationVariables> {
    document = SaveDeliveryOptionEditDaoDocument;
    
    constructor(apollo: Apollo.Apollo) {
      super(apollo);
    }
  }