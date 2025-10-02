/* eslint-disable */
import * as Types from '../../../../generated/graphql';

import { gql } from 'apollo-angular';
import { Injectable } from '@angular/core';
import * as Apollo from 'apollo-angular';
export type FetchUserSeatQueryVariables = Types.Exact<{
  id: Types.Scalars['ID'];
}>;


export type FetchUserSeatQuery = { user?: { id: string, name?: string | null, surname?: string | null, email: string, seatGroup?: { id: string, name: string } | null } | null, seatGroups: { edges?: Array<{ node?: { id: string, name: string } | null } | null> | null } };

export type CreateUserSeatMutationVariables = Types.Exact<{
  input: Types.CreateUserInput;
}>;


export type CreateUserSeatMutation = { createUserSeat?: { id: string } | null };

export type UpdateUserSeatMutationVariables = Types.Exact<{
  id: Types.Scalars['ID'];
  input: Types.UpdateUserInput;
}>;


export type UpdateUserSeatMutation = { updateUserSeat?: { id: string } | null };

export type UpdatePasswordMutationVariables = Types.Exact<{
  id: Types.Scalars['ID'];
  input: Types.Scalars['String'];
}>;


export type UpdatePasswordMutation = { updateUserPassword?: { id: string } | null };

export const FetchUserSeatDocument = gql`
    query FetchUserSeat($id: ID!) {
  user(id: $id) {
    id
    name
    surname
    email
    seatGroup {
      id
      name
    }
  }
  seatGroups {
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
  export class FetchUserSeatGQL extends Apollo.Query<FetchUserSeatQuery, FetchUserSeatQueryVariables> {
    document = FetchUserSeatDocument;
    
    constructor(apollo: Apollo.Apollo) {
      super(apollo);
    }
  }
export const CreateUserSeatDocument = gql`
    mutation CreateUserSeat($input: CreateUserInput!) {
  createUserSeat(input: $input) {
    id
  }
}
    `;

  @Injectable({
    providedIn: 'root'
  })
  export class CreateUserSeatGQL extends Apollo.Mutation<CreateUserSeatMutation, CreateUserSeatMutationVariables> {
    document = CreateUserSeatDocument;
    
    constructor(apollo: Apollo.Apollo) {
      super(apollo);
    }
  }
export const UpdateUserSeatDocument = gql`
    mutation UpdateUserSeat($id: ID!, $input: UpdateUserInput!) {
  updateUserSeat(id: $id, input: $input) {
    id
  }
}
    `;

  @Injectable({
    providedIn: 'root'
  })
  export class UpdateUserSeatGQL extends Apollo.Mutation<UpdateUserSeatMutation, UpdateUserSeatMutationVariables> {
    document = UpdateUserSeatDocument;
    
    constructor(apollo: Apollo.Apollo) {
      super(apollo);
    }
  }
export const UpdatePasswordDocument = gql`
    mutation UpdatePassword($id: ID!, $input: String!) {
  updateUserPassword(id: $id, input: $input) {
    id
  }
}
    `;

  @Injectable({
    providedIn: 'root'
  })
  export class UpdatePasswordGQL extends Apollo.Mutation<UpdatePasswordMutation, UpdatePasswordMutationVariables> {
    document = UpdatePasswordDocument;
    
    constructor(apollo: Apollo.Apollo) {
      super(apollo);
    }
  }