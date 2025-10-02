/* eslint-disable */
import * as Types from '../../../../generated/graphql';

import { gql } from 'apollo-angular';
import { Injectable } from '@angular/core';
import * as Apollo from 'apollo-angular';
export type FetchCarrierEditBringQueryVariables = Types.Exact<{
  id: Types.Scalars['ID'];
}>;


export type FetchCarrierEditBringQuery = { carrier?: { id: string, name: string, carrierBring?: { id: string, apiKey?: string | null, customerNumber?: string | null, test: boolean } | null } | null };

export type UpdateCarrierAgreementBringMutationVariables = Types.Exact<{
  id: Types.Scalars['ID'];
  name: Types.Scalars['String'];
  input: Types.UpdateCarrierBringInput;
}>;


export type UpdateCarrierAgreementBringMutation = { updateCarrierAgreementBring: { id: string } };

export const FetchCarrierEditBringDocument = gql`
    query FetchCarrierEditBring($id: ID!) {
  carrier(id: $id) {
    id
    name
    carrierBring {
      id
      apiKey
      customerNumber
      test
    }
  }
}
    `;

  @Injectable({
    providedIn: 'root'
  })
  export class FetchCarrierEditBringGQL extends Apollo.Query<FetchCarrierEditBringQuery, FetchCarrierEditBringQueryVariables> {
    document = FetchCarrierEditBringDocument;
    
    constructor(apollo: Apollo.Apollo) {
      super(apollo);
    }
  }
export const UpdateCarrierAgreementBringDocument = gql`
    mutation UpdateCarrierAgreementBring($id: ID!, $name: String!, $input: UpdateCarrierBringInput!) {
  updateCarrierAgreementBring(id: $id, name: $name, input: $input) {
    id
  }
}
    `;

  @Injectable({
    providedIn: 'root'
  })
  export class UpdateCarrierAgreementBringGQL extends Apollo.Mutation<UpdateCarrierAgreementBringMutation, UpdateCarrierAgreementBringMutationVariables> {
    document = UpdateCarrierAgreementBringDocument;
    
    constructor(apollo: Apollo.Apollo) {
      super(apollo);
    }
  }