/* eslint-disable */
import * as Types from '../../../generated/graphql.js';

import { gql } from 'apollo-angular';
import { Injectable } from '@angular/core';
import * as Apollo from 'apollo-angular';
export type FetchReturnCollisQueryVariables = Types.Exact<{
  id: Types.Scalars['ID'];
}>;


export type FetchReturnCollisQuery = { returnColli: { collis: Array<{ companyPackageSender: { id: string } }>, order: { orderPublicID: string } } };

export const FetchReturnCollisDocument = gql`
    query FetchReturnCollis($id: ID!) {
  returnColli(id: $id) {
    collis {
      companyPackageSender {
        id
      }
    }
    order {
      orderPublicID
    }
  }
}
    `;

  @Injectable({
    providedIn: 'root'
  })
  export class FetchReturnCollisGQL extends Apollo.Query<FetchReturnCollisQuery, FetchReturnCollisQueryVariables> {
    document = FetchReturnCollisDocument;
    
    constructor(apollo: Apollo.Apollo) {
      super(apollo);
    }
  }