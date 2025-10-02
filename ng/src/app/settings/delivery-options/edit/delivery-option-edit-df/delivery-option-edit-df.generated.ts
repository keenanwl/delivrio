/* eslint-disable */
import * as Types from '../../../../../generated/graphql';

import { gql } from 'apollo-angular';
import { BaseDeliveryOptionFragmentDoc, CarrierServiceItemFragmentDoc } from '../edit-common.generated';
import { Injectable } from '@angular/core';
import * as Apollo from 'apollo-angular';
export type FetchDeliveryOptionEditDfQueryVariables = Types.Exact<{
  id: Types.Scalars['ID'];
}>;


export type FetchDeliveryOptionEditDfQuery = { deliveryOptionDF?: { deliveryOption: { name: string, description?: string | null, clickCollect?: boolean | null, overrideReturnAddress?: boolean | null, overrideSenderAddress?: boolean | null, hideDeliveryOption?: boolean | null, clickOptionDisplayCount?: number | null, deliveryEstimateFrom?: number | null, deliveryEstimateTo?: number | null, webshipperIntegration: boolean, webshipperID?: number | null, shipmondoIntegration: boolean, shipmondoDeliveryOption?: string | null, customsEnabled: boolean, customsSigner?: string | null, hideIfCompanyEmpty: boolean, carrierService: { id: string }, defaultPackaging?: { id: string, name: string } | null } } | null, carrierServices: { edges?: Array<{ node?: { id: string, label: string, return: boolean } | null } | null> | null } };

export type SaveDeliveryOptionEditDfMutationVariables = Types.Exact<{
  id: Types.Scalars['ID'];
  inputDeliveryOption: Types.UpdateDeliveryOptionInput;
}>;


export type SaveDeliveryOptionEditDfMutation = { updateDeliveryOptionDF: { id: string } };

export const FetchDeliveryOptionEditDfDocument = gql`
    query FetchDeliveryOptionEditDF($id: ID!) {
  deliveryOptionDF(id: $id) {
    deliveryOption {
      ...BaseDeliveryOption
    }
  }
  carrierServices(where: {hasCarrierBrandWith: {internalID: df}}) {
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
  export class FetchDeliveryOptionEditDfGQL extends Apollo.Query<FetchDeliveryOptionEditDfQuery, FetchDeliveryOptionEditDfQueryVariables> {
    document = FetchDeliveryOptionEditDfDocument;
    
    constructor(apollo: Apollo.Apollo) {
      super(apollo);
    }
  }
export const SaveDeliveryOptionEditDfDocument = gql`
    mutation SaveDeliveryOptionEditDF($id: ID!, $inputDeliveryOption: UpdateDeliveryOptionInput!) {
  updateDeliveryOptionDF(id: $id, inputDeliveryOption: $inputDeliveryOption) {
    id
  }
}
    `;

  @Injectable({
    providedIn: 'root'
  })
  export class SaveDeliveryOptionEditDfGQL extends Apollo.Mutation<SaveDeliveryOptionEditDfMutation, SaveDeliveryOptionEditDfMutationVariables> {
    document = SaveDeliveryOptionEditDfDocument;
    
    constructor(apollo: Apollo.Apollo) {
      super(apollo);
    }
  }