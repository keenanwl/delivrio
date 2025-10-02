/* eslint-disable */
import * as Types from '../../../../generated/graphql';

import { gql } from 'apollo-angular';
import { Injectable } from '@angular/core';
import * as Apollo from 'apollo-angular';
export type FetchReturnPortalsQueryVariables = Types.Exact<{ [key: string]: never; }>;


export type FetchReturnPortalsQuery = { returnPortals: { edges?: Array<{ node?: { id: string, name: string, returnOpenHours: number, automaticallyAccept: boolean, returnLocation?: Array<{ name: string }> | null, connection?: { name: string } | null } | null } | null> | null }, connections: { edges?: Array<{ node?: { id: string, name: string } | null } | null> | null } };

export type CreateReturnPortalMutationVariables = Types.Exact<{
  name: Types.Scalars['String'];
  connection: Types.Scalars['ID'];
}>;


export type CreateReturnPortalMutation = { createReturnPortal: string };

export const FetchReturnPortalsDocument = gql`
    query FetchReturnPortals {
  returnPortals {
    edges {
      node {
        id
        name
        returnOpenHours
        automaticallyAccept
        returnLocation {
          name
        }
        connection {
          name
        }
      }
    }
  }
  connections {
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
  export class FetchReturnPortalsGQL extends Apollo.Query<FetchReturnPortalsQuery, FetchReturnPortalsQueryVariables> {
    document = FetchReturnPortalsDocument;
    
    constructor(apollo: Apollo.Apollo) {
      super(apollo);
    }
  }
export const CreateReturnPortalDocument = gql`
    mutation CreateReturnPortal($name: String!, $connection: ID!) {
  createReturnPortal(name: $name, connection: $connection)
}
    `;

  @Injectable({
    providedIn: 'root'
  })
  export class CreateReturnPortalGQL extends Apollo.Mutation<CreateReturnPortalMutation, CreateReturnPortalMutationVariables> {
    document = CreateReturnPortalDocument;
    
    constructor(apollo: Apollo.Apollo) {
      super(apollo);
    }
  }