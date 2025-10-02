/* eslint-disable */
import * as Types from '../../../generated/graphql';

import { gql } from 'apollo-angular';
import { Injectable } from '@angular/core';
import * as Apollo from 'apollo-angular';
export type FetchProductTagsEditQueryVariables = Types.Exact<{ [key: string]: never; }>;


export type FetchProductTagsEditQuery = { productTags: { edges?: Array<{ node?: { id: string, name: string, createdAt?: any | null, products?: Array<{ id: string }> | null } | null } | null> | null } };

export type CreateTagsMutationVariables = Types.Exact<{
  input?: Types.InputMaybe<Array<Types.Scalars['String']> | Types.Scalars['String']>;
}>;


export type CreateTagsMutation = { createProductTags?: Array<{ id: string, name: string, createdAt?: any | null, products?: Array<{ id: string }> | null } | null> | null };

export type DeleteTagMutationVariables = Types.Exact<{
  id: Types.Scalars['ID'];
}>;


export type DeleteTagMutation = { deleteTag?: Array<{ id: string, name: string, createdAt?: any | null, products?: Array<{ id: string }> | null } | null> | null };

export const FetchProductTagsEditDocument = gql`
    query FetchProductTagsEdit {
  productTags {
    edges {
      node {
        id
        name
        createdAt
        products {
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
  export class FetchProductTagsEditGQL extends Apollo.Query<FetchProductTagsEditQuery, FetchProductTagsEditQueryVariables> {
    document = FetchProductTagsEditDocument;
    
    constructor(apollo: Apollo.Apollo) {
      super(apollo);
    }
  }
export const CreateTagsDocument = gql`
    mutation CreateTags($input: [String!]) {
  createProductTags(input: $input) {
    id
    name
    createdAt
    products {
      id
    }
  }
}
    `;

  @Injectable({
    providedIn: 'root'
  })
  export class CreateTagsGQL extends Apollo.Mutation<CreateTagsMutation, CreateTagsMutationVariables> {
    document = CreateTagsDocument;
    
    constructor(apollo: Apollo.Apollo) {
      super(apollo);
    }
  }
export const DeleteTagDocument = gql`
    mutation DeleteTag($id: ID!) {
  deleteTag(id: $id) {
    id
    name
    createdAt
    products {
      id
    }
  }
}
    `;

  @Injectable({
    providedIn: 'root'
  })
  export class DeleteTagGQL extends Apollo.Mutation<DeleteTagMutation, DeleteTagMutationVariables> {
    document = DeleteTagDocument;
    
    constructor(apollo: Apollo.Apollo) {
      super(apollo);
    }
  }