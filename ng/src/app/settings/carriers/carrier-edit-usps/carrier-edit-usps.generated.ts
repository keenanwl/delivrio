/* eslint-disable */
import * as Types from '../../../../generated/graphql';

import { gql } from 'apollo-angular';
import { Injectable } from '@angular/core';
import * as Apollo from 'apollo-angular';
export type FetchCarrierUspsQueryVariables = Types.Exact<{
  carrierID: Types.Scalars['ID'];
}>;


export type FetchCarrierUspsQuery = { carrierUSPS?: { isTestAPI: boolean, epsAccountNumber?: string | null, mid?: string | null, manifestMid?: string | null, consumerKey?: string | null, consumerSecret?: string | null, crid?: string | null, carrier: { name: string } } | null };

export type UpdateCarrierUspsMutationVariables = Types.Exact<{
  id: Types.Scalars['ID'];
  name: Types.Scalars['String'];
  input: Types.UpdateCarrierUspsInput;
}>;


export type UpdateCarrierUspsMutation = { updateCarrierAgreementUSPS: { id: string } };

export const FetchCarrierUspsDocument = gql`
    query FetchCarrierUSPS($carrierID: ID!) {
  carrierUSPS(id: $carrierID) {
    carrier {
      name
    }
    isTestAPI
    epsAccountNumber
    mid
    manifestMid
    consumerKey
    consumerSecret
    crid
  }
}
    `;

  @Injectable({
    providedIn: 'root'
  })
  export class FetchCarrierUspsGQL extends Apollo.Query<FetchCarrierUspsQuery, FetchCarrierUspsQueryVariables> {
    document = FetchCarrierUspsDocument;
    
    constructor(apollo: Apollo.Apollo) {
      super(apollo);
    }
  }
export const UpdateCarrierUspsDocument = gql`
    mutation UpdateCarrierUSPS($id: ID!, $name: String!, $input: UpdateCarrierUSPSInput!) {
  updateCarrierAgreementUSPS(id: $id, name: $name, input: $input) {
    id
  }
}
    `;

  @Injectable({
    providedIn: 'root'
  })
  export class UpdateCarrierUspsGQL extends Apollo.Mutation<UpdateCarrierUspsMutation, UpdateCarrierUspsMutationVariables> {
    document = UpdateCarrierUspsDocument;
    
    constructor(apollo: Apollo.Apollo) {
      super(apollo);
    }
  }