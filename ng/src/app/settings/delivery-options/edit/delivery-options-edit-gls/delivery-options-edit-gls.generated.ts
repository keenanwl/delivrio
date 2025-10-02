/* eslint-disable */
import * as Types from '../../../../../generated/graphql';

import { gql } from 'apollo-angular';
import { BaseDeliveryOptionFragmentDoc, CarrierServiceItemFragmentDoc } from '../edit-common.generated';
import { Injectable } from '@angular/core';
import * as Apollo from 'apollo-angular';
export type FetchDeliveryOptionsGlsEditQueryVariables = Types.Exact<{
  id: Types.Scalars['ID'];
}>;


export type FetchDeliveryOptionsGlsEditQuery = { deliveryOptionGLS?: { deliveryOption: { name: string, description?: string | null, clickCollect?: boolean | null, overrideReturnAddress?: boolean | null, overrideSenderAddress?: boolean | null, hideDeliveryOption?: boolean | null, clickOptionDisplayCount?: number | null, deliveryEstimateFrom?: number | null, deliveryEstimateTo?: number | null, webshipperIntegration: boolean, webshipperID?: number | null, shipmondoIntegration: boolean, shipmondoDeliveryOption?: string | null, customsEnabled: boolean, customsSigner?: string | null, hideIfCompanyEmpty: boolean, carrierService: { id: string }, defaultPackaging?: { id: string, name: string } | null } } | null, carrierServices: { edges?: Array<{ node?: { id: string, label: string, return: boolean } | null } | null> | null }, emailTemplates: { edges?: Array<{ node?: { id: string, name: string, mergeType: Types.EmailTemplateMergeType } | null } | null> | null } };

export type UpdateDeliveryOptionGlsMutationVariables = Types.Exact<{
  id: Types.Scalars['ID'];
  inputDeliveryOption: Types.UpdateDeliveryOptionInput;
}>;


export type UpdateDeliveryOptionGlsMutation = { updateDeliveryOptionGLS: { id: string } };

export const FetchDeliveryOptionsGlsEditDocument = gql`
    query FetchDeliveryOptionsGLSEdit($id: ID!) {
  deliveryOptionGLS(id: $id) {
    deliveryOption {
      ...BaseDeliveryOption
    }
  }
  carrierServices(where: {hasCarrierBrandWith: {internalID: gls}}) {
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
  export class FetchDeliveryOptionsGlsEditGQL extends Apollo.Query<FetchDeliveryOptionsGlsEditQuery, FetchDeliveryOptionsGlsEditQueryVariables> {
    document = FetchDeliveryOptionsGlsEditDocument;
    
    constructor(apollo: Apollo.Apollo) {
      super(apollo);
    }
  }
export const UpdateDeliveryOptionGlsDocument = gql`
    mutation UpdateDeliveryOptionGLS($id: ID!, $inputDeliveryOption: UpdateDeliveryOptionInput!) {
  updateDeliveryOptionGLS(id: $id, inputDeliveryOption: $inputDeliveryOption) {
    id
  }
}
    `;

  @Injectable({
    providedIn: 'root'
  })
  export class UpdateDeliveryOptionGlsGQL extends Apollo.Mutation<UpdateDeliveryOptionGlsMutation, UpdateDeliveryOptionGlsMutationVariables> {
    document = UpdateDeliveryOptionGlsDocument;
    
    constructor(apollo: Apollo.Apollo) {
      super(apollo);
    }
  }