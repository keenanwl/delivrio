/* eslint-disable */
import * as Types from '../../../../generated/graphql';

import { gql } from 'apollo-angular';
import { Injectable } from '@angular/core';
import * as Apollo from 'apollo-angular';
export type FetchCarrierEditDsvQueryVariables = Types.Exact<{
  id: Types.Scalars['ID'];
}>;


export type FetchCarrierEditDsvQuery = { carrier?: { id: string, name: string, carrierDSV?: { id: string } | null } | null };

export type UpdateCarrierAgreementDsvMutationVariables = Types.Exact<{
  id: Types.Scalars['ID'];
  name: Types.Scalars['String'];
  input: Types.UpdateCarrierDsvInput;
}>;


export type UpdateCarrierAgreementDsvMutation = { updateCarrierAgreementDSV: { id: string } };

export const FetchCarrierEditDsvDocument = gql`
    query FetchCarrierEditDSV($id: ID!) {
  carrier(id: $id) {
    id
    name
    carrierDSV {
      id
    }
  }
}
    `;

  @Injectable({
    providedIn: 'root'
  })
  export class FetchCarrierEditDsvGQL extends Apollo.Query<FetchCarrierEditDsvQuery, FetchCarrierEditDsvQueryVariables> {
    document = FetchCarrierEditDsvDocument;
    
    constructor(apollo: Apollo.Apollo) {
      super(apollo);
    }
  }
export const UpdateCarrierAgreementDsvDocument = gql`
    mutation UpdateCarrierAgreementDSV($id: ID!, $name: String!, $input: UpdateCarrierDSVInput!) {
  updateCarrierAgreementDSV(id: $id, name: $name, input: $input) {
    id
  }
}
    `;

  @Injectable({
    providedIn: 'root'
  })
  export class UpdateCarrierAgreementDsvGQL extends Apollo.Mutation<UpdateCarrierAgreementDsvMutation, UpdateCarrierAgreementDsvMutationVariables> {
    document = UpdateCarrierAgreementDsvDocument;
    
    constructor(apollo: Apollo.Apollo) {
      super(apollo);
    }
  }