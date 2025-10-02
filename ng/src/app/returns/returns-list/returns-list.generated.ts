/* eslint-disable */
import * as Types from '../../../generated/graphql';

import { gql } from 'apollo-angular';
import { Injectable } from '@angular/core';
import * as Apollo from 'apollo-angular';
export type FetchReturnsListQueryVariables = Types.Exact<{ [key: string]: never; }>;


export type FetchReturnsListQuery = { returnCollis: { edges?: Array<{ node?: { id: string, status: Types.ReturnColliStatus, createdAt: any, order: { id: string, orderPublicID: string }, recipient: { firstName: string, lastName: string, country: { alpha2: string } }, sender: { country: { alpha2: string } } } | null } | null> | null } };

export type FetchReturnDeliveryOptionsQueryVariables = Types.Exact<{
  returnColliIDs: Array<Types.Scalars['ID']> | Types.Scalars['ID'];
}>;


export type FetchReturnDeliveryOptionsQuery = { returnDeliveryOptions: Array<Array<{ deliveryOptionID: string, name: string, description: string, warning?: string | null } | null> | null> };

export type CreateReturnOrderMutationVariables = Types.Exact<{
  orderID: Types.Scalars['ID'];
  orderLines: Array<Types.MutateReturnItems> | Types.MutateReturnItems;
  portalID: Types.Scalars['ID'];
}>;


export type CreateReturnOrderMutation = { createReturnOrder: Array<string> };

export type MarkReturnDeletedMutationVariables = Types.Exact<{
  returnColliID: Types.Scalars['ID'];
}>;


export type MarkReturnDeletedMutation = { markColliDeleted: boolean };

export type ReturnAddDeliveryOptionQueryVariables = Types.Exact<{
  deliveryOptions: Array<Types.MutateReturnDeliveryOption> | Types.MutateReturnDeliveryOption;
}>;


export type ReturnAddDeliveryOptionQuery = { addReturnDeliveryOption: string };

export type SearchOrdersQueryVariables = Types.Exact<{
  term: Types.Scalars['String'];
}>;


export type SearchOrdersQuery = { search: Array<{ id: string, entity: Types.EntityType, title: string, imagePath?: string | null }> };

export type SearchOrderLinesQueryVariables = Types.Exact<{
  order: Types.Scalars['ID'];
}>;


export type SearchOrderLinesQuery = { order?: { colli?: Array<{ id: string, orderLines?: Array<{ id: string, units: number, unitPrice: number, currency: { display: string }, productVariant: { description?: string | null, productImage?: Array<{ url: string }> | null, product: { title: string } } }> | null }> | null } | null, returnClaimsByOrder: Array<{ id: string, name: string, description: string }> };

export const FetchReturnsListDocument = gql`
    query FetchReturnsList {
  returnCollis(where: {not: {status: Deleted}}) {
    edges {
      node {
        id
        order {
          id
          orderPublicID
        }
        status
        createdAt
        recipient {
          firstName
          lastName
          country {
            alpha2
          }
        }
        sender {
          country {
            alpha2
          }
        }
      }
    }
  }
}
    `;

  @Injectable({
    providedIn: 'root'
  })
  export class FetchReturnsListGQL extends Apollo.Query<FetchReturnsListQuery, FetchReturnsListQueryVariables> {
    document = FetchReturnsListDocument;
    
    constructor(apollo: Apollo.Apollo) {
      super(apollo);
    }
  }
export const FetchReturnDeliveryOptionsDocument = gql`
    query FetchReturnDeliveryOptions($returnColliIDs: [ID!]!) {
  returnDeliveryOptions(returnColliIDs: $returnColliIDs) {
    deliveryOptionID
    name
    description
    warning
  }
}
    `;

  @Injectable({
    providedIn: 'root'
  })
  export class FetchReturnDeliveryOptionsGQL extends Apollo.Query<FetchReturnDeliveryOptionsQuery, FetchReturnDeliveryOptionsQueryVariables> {
    document = FetchReturnDeliveryOptionsDocument;
    
    constructor(apollo: Apollo.Apollo) {
      super(apollo);
    }
  }
export const CreateReturnOrderDocument = gql`
    mutation CreateReturnOrder($orderID: ID!, $orderLines: [MutateReturnItems!]!, $portalID: ID!) {
  createReturnOrder(
    orderID: $orderID
    orderLines: $orderLines
    portalID: $portalID
  )
}
    `;

  @Injectable({
    providedIn: 'root'
  })
  export class CreateReturnOrderGQL extends Apollo.Mutation<CreateReturnOrderMutation, CreateReturnOrderMutationVariables> {
    document = CreateReturnOrderDocument;
    
    constructor(apollo: Apollo.Apollo) {
      super(apollo);
    }
  }
export const MarkReturnDeletedDocument = gql`
    mutation MarkReturnDeleted($returnColliID: ID!) {
  markColliDeleted(returnColliID: $returnColliID)
}
    `;

  @Injectable({
    providedIn: 'root'
  })
  export class MarkReturnDeletedGQL extends Apollo.Mutation<MarkReturnDeletedMutation, MarkReturnDeletedMutationVariables> {
    document = MarkReturnDeletedDocument;
    
    constructor(apollo: Apollo.Apollo) {
      super(apollo);
    }
  }
export const ReturnAddDeliveryOptionDocument = gql`
    query ReturnAddDeliveryOption($deliveryOptions: [MutateReturnDeliveryOption!]!) {
  addReturnDeliveryOption(deliveryOptions: $deliveryOptions)
}
    `;

  @Injectable({
    providedIn: 'root'
  })
  export class ReturnAddDeliveryOptionGQL extends Apollo.Query<ReturnAddDeliveryOptionQuery, ReturnAddDeliveryOptionQueryVariables> {
    document = ReturnAddDeliveryOptionDocument;
    
    constructor(apollo: Apollo.Apollo) {
      super(apollo);
    }
  }
export const SearchOrdersDocument = gql`
    query SearchOrders($term: String!) {
  search(term: $term, filter: [ORDER]) {
    id
    entity
    title
    imagePath
  }
}
    `;

  @Injectable({
    providedIn: 'root'
  })
  export class SearchOrdersGQL extends Apollo.Query<SearchOrdersQuery, SearchOrdersQueryVariables> {
    document = SearchOrdersDocument;
    
    constructor(apollo: Apollo.Apollo) {
      super(apollo);
    }
  }
export const SearchOrderLinesDocument = gql`
    query SearchOrderLines($order: ID!) {
  order(id: $order) {
    colli {
      id
      orderLines {
        id
        units
        unitPrice
        currency {
          display
        }
        productVariant {
          description
          productImage {
            url
          }
          product {
            title
          }
        }
      }
    }
  }
  returnClaimsByOrder(orderID: $order) {
    id
    name
    description
  }
}
    `;

  @Injectable({
    providedIn: 'root'
  })
  export class SearchOrderLinesGQL extends Apollo.Query<SearchOrderLinesQuery, SearchOrderLinesQueryVariables> {
    document = SearchOrderLinesDocument;
    
    constructor(apollo: Apollo.Apollo) {
      super(apollo);
    }
  }