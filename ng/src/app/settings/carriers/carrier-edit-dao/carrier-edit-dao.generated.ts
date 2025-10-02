/* eslint-disable */
import * as Types from '../../../../generated/graphql';

import { gql } from 'apollo-angular';
import { Injectable } from '@angular/core';
import * as Apollo from 'apollo-angular';
export type FetchCarrierEditDaoQueryVariables = Types.Exact<{
  id: Types.Scalars['ID'];
}>;


export type FetchCarrierEditDaoQuery = { carrier?: { id: string, name: string, carrierDAO?: { id: string, customerID?: string | null, apiKey?: string | null } | null } | null };

export type UpdateCarrierAgreementDaoMutationVariables = Types.Exact<{
  id: Types.Scalars['ID'];
  name: Types.Scalars['String'];
  input: Types.UpdateCarrierDaoInput;
}>;


export type UpdateCarrierAgreementDaoMutation = { updateCarrierAgreementDAO: { id: string } };

export const FetchCarrierEditDaoDocument = gql`
    query FetchCarrierEditDAO($id: ID!) {
  carrier(id: $id) {
    id
    name
    carrierDAO {
      id
      customerID
      apiKey
    }
  }
}
    `;

  @Injectable({
    providedIn: 'root'
  })
  export class FetchCarrierEditDaoGQL extends Apollo.Query<FetchCarrierEditDaoQuery, FetchCarrierEditDaoQueryVariables> {
    document = FetchCarrierEditDaoDocument;
    
    constructor(apollo: Apollo.Apollo) {
      super(apollo);
    }
  }
export const UpdateCarrierAgreementDaoDocument = gql`
    mutation UpdateCarrierAgreementDAO($id: ID!, $name: String!, $input: UpdateCarrierDAOInput!) {
  updateCarrierAgreementDAO(id: $id, name: $name, input: $input) {
    id
  }
}
    `;

  @Injectable({
    providedIn: 'root'
  })
  export class UpdateCarrierAgreementDaoGQL extends Apollo.Mutation<UpdateCarrierAgreementDaoMutation, UpdateCarrierAgreementDaoMutationVariables> {
    document = UpdateCarrierAgreementDaoDocument;
    
    constructor(apollo: Apollo.Apollo) {
      super(apollo);
    }
  }