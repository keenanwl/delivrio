/* eslint-disable */
import * as Types from '../../../../generated/graphql';

import { gql } from 'apollo-angular';
import { Injectable } from '@angular/core';
import * as Apollo from 'apollo-angular';
export type FetchHypothesisTestingListQueryVariables = Types.Exact<{ [key: string]: never; }>;


export type FetchHypothesisTestingListQuery = { hypothesisTests: { edges?: Array<{ node?: { id: string, name: string, active: boolean, connection: { name: string } } | null } | null> | null }, connections: { edges?: Array<{ node?: { id: string, name: string, connectionBrand: { label: string } } | null } | null> | null } };

export type CreateHypothesisTestDeliveryOptionMutationVariables = Types.Exact<{
  name: Types.Scalars['String'];
  connectionID: Types.Scalars['ID'];
}>;


export type CreateHypothesisTestDeliveryOptionMutation = { createHypothesisTestDeliveryOption: string };

export const FetchHypothesisTestingListDocument = gql`
    query FetchHypothesisTestingList {
  hypothesisTests {
    edges {
      node {
        id
        name
        active
        connection {
          name
        }
      }
    }
  }
  connections {
    edges {
      node {
        id
        name
        connectionBrand {
          label
        }
      }
    }
  }
}
    `;

  @Injectable({
    providedIn: 'root'
  })
  export class FetchHypothesisTestingListGQL extends Apollo.Query<FetchHypothesisTestingListQuery, FetchHypothesisTestingListQueryVariables> {
    document = FetchHypothesisTestingListDocument;
    
    constructor(apollo: Apollo.Apollo) {
      super(apollo);
    }
  }
export const CreateHypothesisTestDeliveryOptionDocument = gql`
    mutation CreateHypothesisTestDeliveryOption($name: String!, $connectionID: ID!) {
  createHypothesisTestDeliveryOption(name: $name, connectionID: $connectionID)
}
    `;

  @Injectable({
    providedIn: 'root'
  })
  export class CreateHypothesisTestDeliveryOptionGQL extends Apollo.Mutation<CreateHypothesisTestDeliveryOptionMutation, CreateHypothesisTestDeliveryOptionMutationVariables> {
    document = CreateHypothesisTestDeliveryOptionDocument;
    
    constructor(apollo: Apollo.Apollo) {
      super(apollo);
    }
  }