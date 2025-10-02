/* eslint-disable */
import * as Types from '../../../../generated/graphql';

import { gql } from 'apollo-angular';
import { Injectable } from '@angular/core';
import * as Apollo from 'apollo-angular';
export type ListConnectionsQueryVariables = Types.Exact<{ [key: string]: never; }>;


export type ListConnectionsQuery = { connections: { edges?: Array<{ node?: { id: string, name: string, syncOrders: boolean, syncProducts: boolean, connectionBrand: { id: string, label: string, logoURL?: string | null } } | null } | null> | null } };

export const ListConnectionsDocument = gql`
    query ListConnections {
  connections {
    edges {
      node {
        id
        name
        syncOrders
        syncProducts
        connectionBrand {
          id
          label
          logoURL
        }
      }
    }
  }
}
    `;

  @Injectable({
    providedIn: 'root'
  })
  export class ListConnectionsGQL extends Apollo.Query<ListConnectionsQuery, ListConnectionsQueryVariables> {
    document = ListConnectionsDocument;
    
    constructor(apollo: Apollo.Apollo) {
      super(apollo);
    }
  }