/* eslint-disable */
import * as Types from '../../../../generated/graphql';

import { gql } from 'apollo-angular';
import { Injectable } from '@angular/core';
import * as Apollo from 'apollo-angular';
export type FetchCarrierEditDfQueryVariables = Types.Exact<{
  id: Types.Scalars['ID'];
}>;


export type FetchCarrierEditDfQuery = { carrier?: { id: string, name: string, carrierDF?: { id: string, customerID: string, agreementNumber: string } | null } | null };

export type UpdateCarrierAgreementDfMutationVariables = Types.Exact<{
  id: Types.Scalars['ID'];
  name: Types.Scalars['String'];
  input: Types.UpdateCarrierDfInput;
}>;


export type UpdateCarrierAgreementDfMutation = { updateCarrierAgreementDF: { id: string } };

export const FetchCarrierEditDfDocument = gql`
    query FetchCarrierEditDF($id: ID!) {
  carrier(id: $id) {
    id
    name
    carrierDF {
      id
      customerID
      agreementNumber
    }
  }
}
    `;

  @Injectable({
    providedIn: 'root'
  })
  export class FetchCarrierEditDfGQL extends Apollo.Query<FetchCarrierEditDfQuery, FetchCarrierEditDfQueryVariables> {
    document = FetchCarrierEditDfDocument;
    
    constructor(apollo: Apollo.Apollo) {
      super(apollo);
    }
  }
export const UpdateCarrierAgreementDfDocument = gql`
    mutation UpdateCarrierAgreementDF($id: ID!, $name: String!, $input: UpdateCarrierDFInput!) {
  updateCarrierAgreementDF(id: $id, name: $name, input: $input) {
    id
  }
}
    `;

  @Injectable({
    providedIn: 'root'
  })
  export class UpdateCarrierAgreementDfGQL extends Apollo.Mutation<UpdateCarrierAgreementDfMutation, UpdateCarrierAgreementDfMutationVariables> {
    document = UpdateCarrierAgreementDfDocument;
    
    constructor(apollo: Apollo.Apollo) {
      super(apollo);
    }
  }