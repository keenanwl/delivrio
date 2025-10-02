/* eslint-disable */
import * as Types from '../../../../../generated/graphql';

import { gql } from 'apollo-angular';
import { BaseDeliveryOptionFragmentDoc, CarrierServiceItemFragmentDoc } from '../edit-common.generated';
import { Injectable } from '@angular/core';
import * as Apollo from 'apollo-angular';
export type FetchDeliveryOptionEditEasyPostQueryVariables = Types.Exact<{
  id: Types.Scalars['ID'];
}>;


export type FetchDeliveryOptionEditEasyPostQuery = { deliveryOptionEasyPost?: { deliveryOption: { clickOptionDisplayCount?: number | null, name: string, description?: string | null, clickCollect?: boolean | null, overrideReturnAddress?: boolean | null, overrideSenderAddress?: boolean | null, hideDeliveryOption?: boolean | null, deliveryEstimateFrom?: number | null, deliveryEstimateTo?: number | null, webshipperIntegration: boolean, webshipperID?: number | null, shipmondoIntegration: boolean, shipmondoDeliveryOption?: string | null, customsEnabled: boolean, customsSigner?: string | null, hideIfCompanyEmpty: boolean, carrierService: { id: string }, defaultPackaging?: { id: string, name: string } | null }, carrierAddServEasyPost?: Array<{ id: string }> | null } | null, carrierServices: { edges?: Array<{ node?: { id: string, label: string, return: boolean } | null } | null> | null } };

export type FetchAdditionalServiceEasyPostQueryVariables = Types.Exact<{
  carrierServiceID: Types.Scalars['ID'];
}>;


export type FetchAdditionalServiceEasyPostQuery = { carrierAdditionalServiceEasyPosts: { edges?: Array<{ node?: { id: string, label: string } | null } | null> | null } };

export type SaveDeliveryOptionEditEasyPostMutationVariables = Types.Exact<{
  id: Types.Scalars['ID'];
  input: Types.UpdateDeliveryOptionEasyPostInput;
  inputDeliveryOption: Types.UpdateDeliveryOptionInput;
}>;


export type SaveDeliveryOptionEditEasyPostMutation = { updateDeliveryOptionEasyPost: { id: string } };

export const FetchDeliveryOptionEditEasyPostDocument = gql`
    query FetchDeliveryOptionEditEasyPost($id: ID!) {
  deliveryOptionEasyPost(id: $id) {
    deliveryOption {
      carrierService {
        id
      }
      clickOptionDisplayCount
    }
    carrierAddServEasyPost {
      id
    }
    deliveryOption {
      ...BaseDeliveryOption
    }
  }
  carrierServices(where: {hasCarrierBrandWith: {internalID: easy_post}}) {
    edges {
      node {
        ...CarrierServiceItem
      }
    }
  }
}
    ${BaseDeliveryOptionFragmentDoc}
${CarrierServiceItemFragmentDoc}`;

  @Injectable({
    providedIn: 'root'
  })
  export class FetchDeliveryOptionEditEasyPostGQL extends Apollo.Query<FetchDeliveryOptionEditEasyPostQuery, FetchDeliveryOptionEditEasyPostQueryVariables> {
    document = FetchDeliveryOptionEditEasyPostDocument;
    
    constructor(apollo: Apollo.Apollo) {
      super(apollo);
    }
  }
export const FetchAdditionalServiceEasyPostDocument = gql`
    query FetchAdditionalServiceEasyPost($carrierServiceID: ID!) {
  carrierAdditionalServiceEasyPosts(
    where: {hasCarrierServiceEasyPostWith: {hasCarrierServiceWith: {id: $carrierServiceID}}}
  ) {
    edges {
      node {
        id
        label
      }
    }
  }
}
    `;

  @Injectable({
    providedIn: 'root'
  })
  export class FetchAdditionalServiceEasyPostGQL extends Apollo.Query<FetchAdditionalServiceEasyPostQuery, FetchAdditionalServiceEasyPostQueryVariables> {
    document = FetchAdditionalServiceEasyPostDocument;
    
    constructor(apollo: Apollo.Apollo) {
      super(apollo);
    }
  }
export const SaveDeliveryOptionEditEasyPostDocument = gql`
    mutation SaveDeliveryOptionEditEasyPost($id: ID!, $input: UpdateDeliveryOptionEasyPostInput!, $inputDeliveryOption: UpdateDeliveryOptionInput!) {
  updateDeliveryOptionEasyPost(
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
  export class SaveDeliveryOptionEditEasyPostGQL extends Apollo.Mutation<SaveDeliveryOptionEditEasyPostMutation, SaveDeliveryOptionEditEasyPostMutationVariables> {
    document = SaveDeliveryOptionEditEasyPostDocument;
    
    constructor(apollo: Apollo.Apollo) {
      super(apollo);
    }
  }