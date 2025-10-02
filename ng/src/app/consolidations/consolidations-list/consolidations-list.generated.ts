/* eslint-disable */
import * as Types from '../../../generated/graphql';

import { gql } from 'apollo-angular';
import { Injectable } from '@angular/core';
import * as Apollo from 'apollo-angular';
export type FetchConsolidationsQueryVariables = Types.Exact<{ [key: string]: never; }>;


export type FetchConsolidationsQuery = { consolidations: { edges?: Array<{ node?: { id: string, createdAt?: any | null, publicID: string, description?: string | null, status: Types.ConsolidationStatus, pallets?: Array<{ id: string }> | null, orders?: Array<{ id: string }> | null } | null } | null> | null } };

export type AddConsolidationMutationVariables = Types.Exact<{
  publicID: Types.Scalars['String'];
  description: Types.Scalars['String'];
}>;


export type AddConsolidationMutation = { createConsolidation: { id: string } };

export const FetchConsolidationsDocument = gql`
    query FetchConsolidations {
  consolidations {
    edges {
      node {
        id
        createdAt
        publicID
        description
        status
        pallets {
          id
        }
        orders {
          id
        }
      }
    }
  }
}
    `;

  @Injectable({
    providedIn: 'root'
  })
  export class FetchConsolidationsGQL extends Apollo.Query<FetchConsolidationsQuery, FetchConsolidationsQueryVariables> {
    document = FetchConsolidationsDocument;
    
    constructor(apollo: Apollo.Apollo) {
      super(apollo);
    }
  }
export const AddConsolidationDocument = gql`
    mutation AddConsolidation($publicID: String!, $description: String!) {
  createConsolidation(publicID: $publicID, description: $description) {
    id
  }
}
    `;

  @Injectable({
    providedIn: 'root'
  })
  export class AddConsolidationGQL extends Apollo.Mutation<AddConsolidationMutation, AddConsolidationMutationVariables> {
    document = AddConsolidationDocument;
    
    constructor(apollo: Apollo.Apollo) {
      super(apollo);
    }
  }