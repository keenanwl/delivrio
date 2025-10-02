/* eslint-disable */
import * as Types from '../../../generated/graphql';

import { gql } from 'apollo-angular';
import { Injectable } from '@angular/core';
import * as Apollo from 'apollo-angular';
export type FetchMyApiTokensQueryVariables = Types.Exact<{ [key: string]: never; }>;


export type FetchMyApiTokensQuery = { myAPITokens: Array<{ id: string, name: string, createdAt?: any | null, lastUsed?: any | null }> };

export type ApiTokenFragmentFragment = { id: string, name: string, createdAt?: any | null, lastUsed?: any | null };

export type CreateApiTokenMutationVariables = Types.Exact<{
  name: Types.Scalars['String'];
}>;


export type CreateApiTokenMutation = { createAPIToken: { id: string, token: string } };

export type DeleteTokenMutationVariables = Types.Exact<{
  ID: Types.Scalars['ID'];
}>;


export type DeleteTokenMutation = { deleteAPIToken: Array<{ id: string, name: string, createdAt?: any | null, lastUsed?: any | null }> };

export const ApiTokenFragmentFragmentDoc = gql`
    fragment APITokenFragment on APIToken {
  id
  name
  createdAt
  lastUsed
}
    `;
export const FetchMyApiTokensDocument = gql`
    query FetchMyAPITokens {
  myAPITokens {
    ...APITokenFragment
  }
}
    ${ApiTokenFragmentFragmentDoc}`;

  @Injectable({
    providedIn: 'root'
  })
  export class FetchMyApiTokensGQL extends Apollo.Query<FetchMyApiTokensQuery, FetchMyApiTokensQueryVariables> {
    document = FetchMyApiTokensDocument;
    
    constructor(apollo: Apollo.Apollo) {
      super(apollo);
    }
  }
export const CreateApiTokenDocument = gql`
    mutation CreateAPIToken($name: String!) {
  createAPIToken(name: $name) {
    id
    token
  }
}
    `;

  @Injectable({
    providedIn: 'root'
  })
  export class CreateApiTokenGQL extends Apollo.Mutation<CreateApiTokenMutation, CreateApiTokenMutationVariables> {
    document = CreateApiTokenDocument;
    
    constructor(apollo: Apollo.Apollo) {
      super(apollo);
    }
  }
export const DeleteTokenDocument = gql`
    mutation DeleteToken($ID: ID!) {
  deleteAPIToken(id: $ID) {
    ...APITokenFragment
  }
}
    ${ApiTokenFragmentFragmentDoc}`;

  @Injectable({
    providedIn: 'root'
  })
  export class DeleteTokenGQL extends Apollo.Mutation<DeleteTokenMutation, DeleteTokenMutationVariables> {
    document = DeleteTokenDocument;
    
    constructor(apollo: Apollo.Apollo) {
      super(apollo);
    }
  }