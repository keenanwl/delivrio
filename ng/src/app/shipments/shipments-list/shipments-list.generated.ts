/* eslint-disable */
import * as Types from '../../../generated/graphql';

import { gql } from 'apollo-angular';
import { Injectable } from '@angular/core';
import * as Apollo from 'apollo-angular';
export type FetchShipmentsQueryVariables = Types.Exact<{
  where?: Types.InputMaybe<Types.ShipmentWhereInput>;
  first?: Types.InputMaybe<Types.Scalars['Int']>;
  last?: Types.InputMaybe<Types.Scalars['Int']>;
  after?: Types.InputMaybe<Types.Scalars['Cursor']>;
  before?: Types.InputMaybe<Types.Scalars['Cursor']>;
}>;


export type FetchShipmentsQuery = { shipments: { totalCount: number, edges?: Array<{ cursor: any, node?: { id: string, shipmentPublicID: string, createdAt: any, status: Types.ShipmentStatus, shipmentParcel?: Array<{ id: string, fulfillmentSyncedAt?: any | null, cancelSyncedAt?: any | null }> | null, shipmentPallet?: Array<{ id: string }> | null } | null } | null> | null, pageInfo: { hasNextPage: boolean, hasPreviousPage: boolean, startCursor?: any | null, endCursor?: any | null } }, emailTemplates: { edges?: Array<{ node?: { id: string, name: string } | null } | null> | null } };

export type ShipmentsSearchCcLocationsQueryVariables = Types.Exact<{
  lookup?: Types.InputMaybe<Types.Scalars['String']>;
}>;


export type ShipmentsSearchCcLocationsQuery = { locations: { edges?: Array<{ node?: { id: string, name: string } | null } | null> | null } };

export type ShipmentSendOverviewEmailQueryVariables = Types.Exact<{
  email: Types.Scalars['String'];
  templateID: Types.Scalars['ID'];
  where: Types.ShipmentWhereInput;
}>;


export type ShipmentSendOverviewEmailQuery = { sendOverviewEmail: boolean };

export const FetchShipmentsDocument = gql`
    query FetchShipments($where: ShipmentWhereInput, $first: Int, $last: Int, $after: Cursor, $before: Cursor) {
  shipments(
    where: $where
    first: $first
    last: $last
    after: $after
    before: $before
    orderBy: {direction: DESC, field: CREATED_AT}
  ) {
    edges {
      node {
        id
        shipmentPublicID
        createdAt
        status
        shipmentParcel {
          id
          fulfillmentSyncedAt
          cancelSyncedAt
        }
        shipmentPallet {
          id
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
  emailTemplates(where: {mergeType: return_colli_label}) {
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
  export class FetchShipmentsGQL extends Apollo.Query<FetchShipmentsQuery, FetchShipmentsQueryVariables> {
    document = FetchShipmentsDocument;
    
    constructor(apollo: Apollo.Apollo) {
      super(apollo);
    }
  }
export const ShipmentsSearchCcLocationsDocument = gql`
    query ShipmentsSearchCCLocations($lookup: String) {
  locations(
    where: {and: {hasLocationTagsWith: {internalIDEqualFold: "click_and_collect"}, nameContainsFold: $lookup}}
  ) {
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
  export class ShipmentsSearchCcLocationsGQL extends Apollo.Query<ShipmentsSearchCcLocationsQuery, ShipmentsSearchCcLocationsQueryVariables> {
    document = ShipmentsSearchCcLocationsDocument;
    
    constructor(apollo: Apollo.Apollo) {
      super(apollo);
    }
  }
export const ShipmentSendOverviewEmailDocument = gql`
    query ShipmentSendOverviewEmail($email: String!, $templateID: ID!, $where: ShipmentWhereInput!) {
  sendOverviewEmail(to: $email, emailTpl: $templateID, where: $where)
}
    `;

  @Injectable({
    providedIn: 'root'
  })
  export class ShipmentSendOverviewEmailGQL extends Apollo.Query<ShipmentSendOverviewEmailQuery, ShipmentSendOverviewEmailQueryVariables> {
    document = ShipmentSendOverviewEmailDocument;
    
    constructor(apollo: Apollo.Apollo) {
      super(apollo);
    }
  }