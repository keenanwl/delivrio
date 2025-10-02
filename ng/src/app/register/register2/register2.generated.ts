/* eslint-disable */
import * as Types from '../../../generated/graphql';

import { gql } from 'apollo-angular';
import { Injectable } from '@angular/core';
import * as Apollo from 'apollo-angular';
export type ReplaceSignupOptionsMutationVariables = Types.Exact<{
  userID: Types.Scalars['ID'];
  input: Types.CreateSignupOptionsInput;
}>;


export type ReplaceSignupOptionsMutation = { replaceSignupOptions?: { id: string } | null };

export const ReplaceSignupOptionsDocument = gql`
    mutation ReplaceSignupOptions($userID: ID!, $input: CreateSignupOptionsInput!) {
  replaceSignupOptions(userID: $userID, input: $input) {
    id
  }
}
    `;

  @Injectable({
    providedIn: 'root'
  })
  export class ReplaceSignupOptionsGQL extends Apollo.Mutation<ReplaceSignupOptionsMutation, ReplaceSignupOptionsMutationVariables> {
    document = ReplaceSignupOptionsDocument;
    
    constructor(apollo: Apollo.Apollo) {
      super(apollo);
    }
  }