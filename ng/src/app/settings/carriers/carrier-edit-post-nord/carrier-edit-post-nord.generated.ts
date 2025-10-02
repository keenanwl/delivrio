/* eslint-disable */
import * as Types from '../../../../generated/graphql';

import { gql } from 'apollo-angular';
import { Injectable } from '@angular/core';
import * as Apollo from 'apollo-angular';
export type FetchCarrierPostNordQueryVariables = Types.Exact<{
  carrierID: Types.Scalars['ID'];
}>;


export type FetchCarrierPostNordQuery = { carrierPostNord?: { customerNumber: string, carrier: { name: string } } | null };

export type UpdateCarrierPostNordMutationVariables = Types.Exact<{
  id: Types.Scalars['ID'];
  name: Types.Scalars['String'];
  input: Types.UpdateCarrierPostNordInput;
}>;


export type UpdateCarrierPostNordMutation = { updateCarrierAgreementPostNord: { id: string } };

export const FetchCarrierPostNordDocument = gql`
    query FetchCarrierPostNord($carrierID: ID!) {
  carrierPostNord(id: $carrierID) {
    customerNumber
    carrier {
      name
    }
  }
}
    `;

  @Injectable({
    providedIn: 'root'
  })
  export class FetchCarrierPostNordGQL extends Apollo.Query<FetchCarrierPostNordQuery, FetchCarrierPostNordQueryVariables> {
    document = FetchCarrierPostNordDocument;
    
    constructor(apollo: Apollo.Apollo) {
      super(apollo);
    }
  }
export const UpdateCarrierPostNordDocument = gql`
    mutation UpdateCarrierPostNord($id: ID!, $name: String!, $input: UpdateCarrierPostNordInput!) {
  updateCarrierAgreementPostNord(id: $id, name: $name, input: $input) {
    id
  }
}
    `;

  @Injectable({
    providedIn: 'root'
  })
  export class UpdateCarrierPostNordGQL extends Apollo.Mutation<UpdateCarrierPostNordMutation, UpdateCarrierPostNordMutationVariables> {
    document = UpdateCarrierPostNordDocument;
    
    constructor(apollo: Apollo.Apollo) {
      super(apollo);
    }
  }