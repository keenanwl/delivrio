/* eslint-disable */
import * as Types from '../../../../generated/graphql';

import { gql } from 'apollo-angular';
import { Injectable } from '@angular/core';
import * as Apollo from 'apollo-angular';
export type FetchUserSeatsQueryVariables = Types.Exact<{ [key: string]: never; }>;


export type FetchUserSeatsQuery = { users: { edges?: Array<{ node?: { id: string, email: string, name?: string | null, surname?: string | null, createdAt?: any | null, isAccountOwner: boolean, seatGroup?: { id: string, name: string } | null } | null } | null> | null } };

export const FetchUserSeatsDocument = gql`
    query FetchUserSeats {
  users(where: {and: {isGlobalAdmin: false, emailNEQ: "k.linsly@gmail.com"}}) {
    edges {
      node {
        id
        email
        name
        surname
        createdAt
        isAccountOwner
        seatGroup {
          id
          name
        }
      }
    }
  }
}
    `;

  @Injectable({
    providedIn: 'root'
  })
  export class FetchUserSeatsGQL extends Apollo.Query<FetchUserSeatsQuery, FetchUserSeatsQueryVariables> {
    document = FetchUserSeatsDocument;
    
    constructor(apollo: Apollo.Apollo) {
      super(apollo);
    }
  }