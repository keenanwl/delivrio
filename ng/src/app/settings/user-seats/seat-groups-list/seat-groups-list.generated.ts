/* eslint-disable */
import * as Types from '../../../../generated/graphql';

import { gql } from 'apollo-angular';
import { Injectable } from '@angular/core';
import * as Apollo from 'apollo-angular';
export type FetchSeatGroupsQueryVariables = Types.Exact<{ [key: string]: never; }>;


export type FetchSeatGroupsQuery = { seatGroups: { edges?: Array<{ node?: { id: string, name: string, createdAt?: any | null, user?: Array<{ id: string, email: string }> | null, seatGroupAccessRight?: Array<{ id: string, level: Types.SeatGroupAccessRightLevel, accessRight: { id: string, label: string } }> | null } | null } | null> | null } };

export const FetchSeatGroupsDocument = gql`
    query FetchSeatGroups {
  seatGroups {
    edges {
      node {
        id
        name
        createdAt
        user {
          id
          email
        }
        seatGroupAccessRight {
          id
          level
          accessRight {
            id
            label
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
  export class FetchSeatGroupsGQL extends Apollo.Query<FetchSeatGroupsQuery, FetchSeatGroupsQueryVariables> {
    document = FetchSeatGroupsDocument;
    
    constructor(apollo: Apollo.Apollo) {
      super(apollo);
    }
  }