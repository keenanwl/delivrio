/* eslint-disable */
import * as Types from '../../../generated/graphql';

import { gql } from 'apollo-angular';
import { Injectable } from '@angular/core';
import * as Apollo from 'apollo-angular';
export type ProfileQueryVariables = Types.Exact<{ [key: string]: never; }>;


export type ProfileQuery = { user?: { id: string, email: string, name?: string | null, surname?: string | null, phoneNumber?: string | null, marketingConsent?: boolean | null, language?: { id: string, label: string } | null, tenant: { id: string, name: string } } | null, availableTenants: Array<{ id: string, name: string }>, languages: { edges?: Array<{ node?: { id: string, label: string } | null } | null> | null } };

export type UpdateUserMutationVariables = Types.Exact<{
  input: Types.UpdateUserInput;
  newTenantID?: Types.InputMaybe<Types.Scalars['ID']>;
}>;


export type UpdateUserMutation = { updateUser?: { id: string } | null };

export const ProfileDocument = gql`
    query Profile {
  user {
    id
    email
    name
    surname
    phoneNumber
    marketingConsent
    language {
      id
      label
    }
    tenant {
      id
      name
    }
  }
  availableTenants {
    id
    name
  }
  languages {
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
  export class ProfileGQL extends Apollo.Query<ProfileQuery, ProfileQueryVariables> {
    document = ProfileDocument;
    
    constructor(apollo: Apollo.Apollo) {
      super(apollo);
    }
  }
export const UpdateUserDocument = gql`
    mutation UpdateUser($input: UpdateUserInput!, $newTenantID: ID) {
  updateUser(input: $input, newTenantID: $newTenantID) {
    id
  }
}
    `;

  @Injectable({
    providedIn: 'root'
  })
  export class UpdateUserGQL extends Apollo.Mutation<UpdateUserMutation, UpdateUserMutationVariables> {
    document = UpdateUserDocument;
    
    constructor(apollo: Apollo.Apollo) {
      super(apollo);
    }
  }