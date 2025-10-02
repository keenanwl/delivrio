/* eslint-disable */
import * as Types from '../../../generated/graphql';

import { gql } from 'apollo-angular';
import { Injectable } from '@angular/core';
import * as Apollo from 'apollo-angular';
export type FetchPackagingQueryVariables = Types.Exact<{ [key: string]: never; }>;


export type FetchPackagingQuery = { packagings: { edges?: Array<{ node?: { id: string, name: string, lengthCm: number, heightCm: number, widthCm: number } | null } | null> | null } };

export const FetchPackagingDocument = gql`
    query FetchPackaging {
  packagings {
    edges {
      node {
        id
        name
        lengthCm
        heightCm
        widthCm
      }
    }
  }
}
    `;

  @Injectable({
    providedIn: 'root'
  })
  export class FetchPackagingGQL extends Apollo.Query<FetchPackagingQuery, FetchPackagingQueryVariables> {
    document = FetchPackagingDocument;
    
    constructor(apollo: Apollo.Apollo) {
      super(apollo);
    }
  }