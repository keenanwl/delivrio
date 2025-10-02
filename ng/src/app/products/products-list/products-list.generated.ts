/* eslint-disable */
import * as Types from '../../../generated/graphql';

import { gql } from 'apollo-angular';
import { Injectable } from '@angular/core';
import * as Apollo from 'apollo-angular';
export type FetchProductsQueryVariables = Types.Exact<{
  first?: Types.InputMaybe<Types.Scalars['Int']>;
  last?: Types.InputMaybe<Types.Scalars['Int']>;
  after?: Types.InputMaybe<Types.Scalars['Cursor']>;
  before?: Types.InputMaybe<Types.Scalars['Cursor']>;
}>;


export type FetchProductsQuery = { products: { totalCount: number, edges?: Array<{ node?: { id: string, title: string, status: Types.ProductStatus, bodyHTML?: string | null, createdAt?: any | null, updatedAt: any, productVariant?: Array<{ id: string, dimensionHeight?: number | null, dimensionWidth?: number | null, dimensionLength?: number | null, productImage?: Array<{ url: string }> | null, product: { productImage?: Array<{ url: string }> | null } }> | null, productTags?: Array<{ id: string, name: string }> | null } | null } | null> | null, pageInfo: { hasNextPage: boolean, hasPreviousPage: boolean, startCursor?: any | null, endCursor?: any | null } } };

export type CreateNewProductMutationVariables = Types.Exact<{
  input: Types.CreateProductInput;
}>;


export type CreateNewProductMutation = { createProduct?: { id: string } | null };

export const FetchProductsDocument = gql`
    query FetchProducts($first: Int, $last: Int, $after: Cursor, $before: Cursor) {
  products(first: $first, last: $last, after: $after, before: $before) {
    edges {
      node {
        id
        title
        status
        bodyHTML
        createdAt
        updatedAt
        productVariant {
          id
          dimensionHeight
          dimensionWidth
          dimensionLength
          productImage {
            url
          }
          product {
            productImage {
              url
            }
          }
        }
        productTags {
          id
          name
        }
      }
    }
    pageInfo {
      hasNextPage
      hasPreviousPage
      startCursor
      endCursor
    }
    totalCount
  }
}
    `;

  @Injectable({
    providedIn: 'root'
  })
  export class FetchProductsGQL extends Apollo.Query<FetchProductsQuery, FetchProductsQueryVariables> {
    document = FetchProductsDocument;
    
    constructor(apollo: Apollo.Apollo) {
      super(apollo);
    }
  }
export const CreateNewProductDocument = gql`
    mutation CreateNewProduct($input: CreateProductInput!) {
  createProduct(input: $input) {
    id
  }
}
    `;

  @Injectable({
    providedIn: 'root'
  })
  export class CreateNewProductGQL extends Apollo.Mutation<CreateNewProductMutation, CreateNewProductMutationVariables> {
    document = CreateNewProductDocument;
    
    constructor(apollo: Apollo.Apollo) {
      super(apollo);
    }
  }