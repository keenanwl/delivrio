/* eslint-disable */
import * as Types from '../../generated/graphql';

import { gql } from 'apollo-angular';
import { Injectable } from '@angular/core';
import * as Apollo from 'apollo-angular';
export type FetchOrdersQueryVariables = Types.Exact<{
  where?: Types.InputMaybe<Types.OrderWhereInput>;
  first?: Types.InputMaybe<Types.Scalars['Int']>;
  last?: Types.InputMaybe<Types.Scalars['Int']>;
  after?: Types.InputMaybe<Types.Scalars['Cursor']>;
  before?: Types.InputMaybe<Types.Scalars['Cursor']>;
  orderBy?: Types.InputMaybe<Types.OrderOrder>;
}>;


export type FetchOrdersQuery = { orders: { totalCount: number, edges?: Array<{ cursor: any, node?: { id: string, createdAt: any, orderPublicID: string, status: Types.OrderStatus, colli?: Array<{ id: string, orderLines?: Array<{ unitPrice: number, units: number, currency: { display: string } }> | null, recipient: { id: string, firstName: string, lastName: string, country: { id: string, label: string, alpha2: string } }, deliveryOption?: { name: string, carrier: { carrierBrand: { label: string } } } | null, shipmentParcel?: { status: Types.ShipmentParcelStatus } | null }> | null, connection: { id: string, name: string, connectionBrand: { label: string } } } | null } | null> | null, pageInfo: { hasNextPage: boolean, hasPreviousPage: boolean, startCursor?: any | null, endCursor?: any | null } }, connections: { edges?: Array<{ node?: { id: string, name: string } | null } | null> | null }, locations: { edges?: Array<{ node?: { id: string, name: string } | null } | null> | null } };

export type CreateEmptyOrderMutationVariables = Types.Exact<{
  input: Types.CreateOrderInput;
}>;


export type CreateEmptyOrderMutation = { createEmptyOrder?: { id: string, colli?: Array<{ id: string }> | null } | null };

export type FetchCountriesQueryVariables = Types.Exact<{ [key: string]: never; }>;


export type FetchCountriesQuery = { countries: { edges?: Array<{ node?: { id: string, label: string } | null } | null> | null } };

export type BulkUpdatePackagingMutationVariables = Types.Exact<{
  orderIDs: Array<Types.Scalars['ID']> | Types.Scalars['ID'];
  packagingID?: Types.InputMaybe<Types.Scalars['ID']>;
}>;


export type BulkUpdatePackagingMutation = { bulkUpdatePackaging: { success: boolean, msg: string } };

export type BulkFetchPackingSlipsByOrderQueryVariables = Types.Exact<{
  orderIDs: Array<Types.Scalars['ID']> | Types.Scalars['ID'];
}>;


export type BulkFetchPackingSlipsByOrderQuery = { packingSlipsByOrder: { allPackingSlips: string, packingSlips: Array<string> } };

export const FetchOrdersDocument = gql`
    query FetchOrders($where: OrderWhereInput, $first: Int, $last: Int, $after: Cursor, $before: Cursor, $orderBy: OrderOrder) {
  orders(
    where: $where
    first: $first
    last: $last
    after: $after
    before: $before
    orderBy: $orderBy
  ) {
    edges {
      node {
        id
        createdAt
        colli {
          id
          orderLines {
            unitPrice
            units
            currency {
              display
            }
          }
          recipient {
            id
            firstName
            lastName
            country {
              id
              label
              alpha2
            }
          }
          deliveryOption {
            name
            carrier {
              carrierBrand {
                label
              }
            }
          }
          shipmentParcel {
            status
          }
        }
        orderPublicID
        status
        connection {
          id
          name
          connectionBrand {
            label
          }
        }
      }
      cursor
    }
    pageInfo {
      hasNextPage
      hasPreviousPage
      startCursor
      endCursor
    }
    totalCount
  }
  connections {
    edges {
      node {
        id
        name
      }
    }
  }
  locations(where: {hasLocationTagsWith: {internalID: "sender"}}) {
    edges {
      node {
        id
        name
      }
    }
  }
}
    `;

  @Injectable({
    providedIn: 'root'
  })
  export class FetchOrdersGQL extends Apollo.Query<FetchOrdersQuery, FetchOrdersQueryVariables> {
    document = FetchOrdersDocument;
    
    constructor(apollo: Apollo.Apollo) {
      super(apollo);
    }
  }
export const CreateEmptyOrderDocument = gql`
    mutation CreateEmptyOrder($input: CreateOrderInput!) {
  createEmptyOrder(input: $input) {
    id
    colli {
      id
    }
  }
}
    `;

  @Injectable({
    providedIn: 'root'
  })
  export class CreateEmptyOrderGQL extends Apollo.Mutation<CreateEmptyOrderMutation, CreateEmptyOrderMutationVariables> {
    document = CreateEmptyOrderDocument;
    
    constructor(apollo: Apollo.Apollo) {
      super(apollo);
    }
  }
export const FetchCountriesDocument = gql`
    query FetchCountries {
  countries {
    edges {
      node {
        id
        label
      }
    }
  }
}
    `;

  @Injectable({
    providedIn: 'root'
  })
  export class FetchCountriesGQL extends Apollo.Query<FetchCountriesQuery, FetchCountriesQueryVariables> {
    document = FetchCountriesDocument;
    
    constructor(apollo: Apollo.Apollo) {
      super(apollo);
    }
  }
export const BulkUpdatePackagingDocument = gql`
    mutation BulkUpdatePackaging($orderIDs: [ID!]!, $packagingID: ID) {
  bulkUpdatePackaging(orderIDs: $orderIDs, packagingID: $packagingID) {
    success
    msg
  }
}
    `;

  @Injectable({
    providedIn: 'root'
  })
  export class BulkUpdatePackagingGQL extends Apollo.Mutation<BulkUpdatePackagingMutation, BulkUpdatePackagingMutationVariables> {
    document = BulkUpdatePackagingDocument;
    
    constructor(apollo: Apollo.Apollo) {
      super(apollo);
    }
  }
export const BulkFetchPackingSlipsByOrderDocument = gql`
    query BulkFetchPackingSlipsByOrder($orderIDs: [ID!]!) {
  packingSlipsByOrder(orderIDs: $orderIDs) {
    allPackingSlips
    packingSlips
  }
}
    `;

  @Injectable({
    providedIn: 'root'
  })
  export class BulkFetchPackingSlipsByOrderGQL extends Apollo.Query<BulkFetchPackingSlipsByOrderQuery, BulkFetchPackingSlipsByOrderQueryVariables> {
    document = BulkFetchPackingSlipsByOrderDocument;
    
    constructor(apollo: Apollo.Apollo) {
      super(apollo);
    }
  }