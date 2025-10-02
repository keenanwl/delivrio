/* eslint-disable */
import * as Types from '../../../../generated/graphql';

import { gql } from 'apollo-angular';
import { Injectable } from '@angular/core';
import * as Apollo from 'apollo-angular';
export type FetchCarrierEasyPostQueryVariables = Types.Exact<{
  carrierID: Types.Scalars['ID'];
}>;


export type FetchCarrierEasyPostQuery = { carrier?: { name: string, carrierEasyPost?: { apiKey: string, carrierAccounts: Array<string>, test: boolean } | null } | null };

export type UpdateCarrierEasyPostMutationVariables = Types.Exact<{
  id: Types.Scalars['ID'];
  name: Types.Scalars['String'];
  input: Types.UpdateCarrierEasyPostInput;
}>;


export type UpdateCarrierEasyPostMutation = { updateCarrierAgreementEasyPost: { id: string } };

export const FetchCarrierEasyPostDocument = gql`
    query FetchCarrierEasyPost($carrierID: ID!) {
  carrier(id: $carrierID) {
    name
    carrierEasyPost {
      apiKey
      carrierAccounts
      test
    }
  }
}
    `;

  @Injectable({
    providedIn: 'root'
  })
  export class FetchCarrierEasyPostGQL extends Apollo.Query<FetchCarrierEasyPostQuery, FetchCarrierEasyPostQueryVariables> {
    document = FetchCarrierEasyPostDocument;
    
    constructor(apollo: Apollo.Apollo) {
      super(apollo);
    }
  }
export const UpdateCarrierEasyPostDocument = gql`
    mutation UpdateCarrierEasyPost($id: ID!, $name: String!, $input: UpdateCarrierEasyPostInput!) {
  updateCarrierAgreementEasyPost(id: $id, name: $name, input: $input) {
    id
  }
}
    `;

  @Injectable({
    providedIn: 'root'
  })
  export class UpdateCarrierEasyPostGQL extends Apollo.Mutation<UpdateCarrierEasyPostMutation, UpdateCarrierEasyPostMutationVariables> {
    document = UpdateCarrierEasyPostDocument;
    
    constructor(apollo: Apollo.Apollo) {
      super(apollo);
    }
  }