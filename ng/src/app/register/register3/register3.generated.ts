/* eslint-disable */
import * as Types from '../../../generated/graphql';

import { gql } from 'apollo-angular';
import { Injectable } from '@angular/core';
import * as Apollo from 'apollo-angular';
export type FetchPlatformsCarriersQueryVariables = Types.Exact<{ [key: string]: never; }>;


export type FetchPlatformsCarriersQuery = { connectOptionCarriers: { edges?: Array<{ node?: { name: string, id: string } | null } | null> | null }, connectOptionPlatforms: { edges?: Array<{ node?: { name: string, id: string } | null } | null> | null } };

export type ReplaceCarriersPlatformsMutationVariables = Types.Exact<{
  userID: Types.Scalars['ID'];
  inputCarriers: Array<Types.Scalars['ID']> | Types.Scalars['ID'];
  inputPlatforms: Array<Types.Scalars['ID']> | Types.Scalars['ID'];
}>;


export type ReplaceCarriersPlatformsMutation = { replaceInterestedCarriersPlatforms?: { id: string } | null };

export const FetchPlatformsCarriersDocument = gql`
    query FetchPlatformsCarriers {
  connectOptionCarriers {
    edges {
      node {
        name
        id
      }
    }
  }
  connectOptionPlatforms {
    edges {
      node {
        name
        id
      }
    }
  }
}
    `;

  @Injectable({
    providedIn: 'root'
  })
  export class FetchPlatformsCarriersGQL extends Apollo.Query<FetchPlatformsCarriersQuery, FetchPlatformsCarriersQueryVariables> {
    document = FetchPlatformsCarriersDocument;
    
    constructor(apollo: Apollo.Apollo) {
      super(apollo);
    }
  }
export const ReplaceCarriersPlatformsDocument = gql`
    mutation ReplaceCarriersPlatforms($userID: ID!, $inputCarriers: [ID!]!, $inputPlatforms: [ID!]!) {
  replaceInterestedCarriersPlatforms(
    userID: $userID
    inputCarriers: $inputCarriers
    inputPlatforms: $inputPlatforms
  ) {
    id
  }
}
    `;

  @Injectable({
    providedIn: 'root'
  })
  export class ReplaceCarriersPlatformsGQL extends Apollo.Mutation<ReplaceCarriersPlatformsMutation, ReplaceCarriersPlatformsMutationVariables> {
    document = ReplaceCarriersPlatformsDocument;
    
    constructor(apollo: Apollo.Apollo) {
      super(apollo);
    }
  }