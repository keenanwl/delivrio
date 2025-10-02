/* eslint-disable */
import * as Types from '../../../../generated/graphql';

import { gql } from 'apollo-angular';
import { Injectable } from '@angular/core';
import * as Apollo from 'apollo-angular';
export type FetchSeatGroupQueryVariables = Types.Exact<{
  id: Types.Scalars['ID'];
}>;


export type FetchSeatGroupQuery = { seatGroup?: { id: string, name: string, seatGroupAccessRight?: Array<{ id: string, level: Types.SeatGroupAccessRightLevel, accessRight: { id: string, label: string, internalID: string } }> | null } | null, accessRights: { edges?: Array<{ node?: { id: string, label: string, internalID: string } | null } | null> | null } };

export type CreateSeatGroupMutationVariables = Types.Exact<{
  input: Types.CreateSeatGroupInput;
  accessRights?: Types.InputMaybe<Array<Types.CreateSeatGroupAccessRightInput> | Types.CreateSeatGroupAccessRightInput>;
}>;


export type CreateSeatGroupMutation = { createSeatGroup?: { id: string } | null };

export type ReplaceSeatGroupMutationVariables = Types.Exact<{
  id: Types.Scalars['ID'];
  input: Types.UpdateSeatGroupInput;
  accessRights?: Types.InputMaybe<Array<Types.CreateSeatGroupAccessRightInput> | Types.CreateSeatGroupAccessRightInput>;
}>;


export type ReplaceSeatGroupMutation = { replaceSeatGroup?: { id: string } | null };

export const FetchSeatGroupDocument = gql`
    query FetchSeatGroup($id: ID!) {
  seatGroup(id: $id) {
    id
    name
    seatGroupAccessRight {
      id
      level
      accessRight {
        id
        label
        internalID
      }
    }
  }
  accessRights {
    edges {
      node {
        id
        label
        internalID
      }
    }
  }
}
    `;

  @Injectable({
    providedIn: 'root'
  })
  export class FetchSeatGroupGQL extends Apollo.Query<FetchSeatGroupQuery, FetchSeatGroupQueryVariables> {
    document = FetchSeatGroupDocument;
    
    constructor(apollo: Apollo.Apollo) {
      super(apollo);
    }
  }
export const CreateSeatGroupDocument = gql`
    mutation CreateSeatGroup($input: CreateSeatGroupInput!, $accessRights: [CreateSeatGroupAccessRightInput!]) {
  createSeatGroup(input: $input, accessRights: $accessRights) {
    id
  }
}
    `;

  @Injectable({
    providedIn: 'root'
  })
  export class CreateSeatGroupGQL extends Apollo.Mutation<CreateSeatGroupMutation, CreateSeatGroupMutationVariables> {
    document = CreateSeatGroupDocument;
    
    constructor(apollo: Apollo.Apollo) {
      super(apollo);
    }
  }
export const ReplaceSeatGroupDocument = gql`
    mutation ReplaceSeatGroup($id: ID!, $input: UpdateSeatGroupInput!, $accessRights: [CreateSeatGroupAccessRightInput!]) {
  replaceSeatGroup(id: $id, input: $input, accessRights: $accessRights) {
    id
  }
}
    `;

  @Injectable({
    providedIn: 'root'
  })
  export class ReplaceSeatGroupGQL extends Apollo.Mutation<ReplaceSeatGroupMutation, ReplaceSeatGroupMutationVariables> {
    document = ReplaceSeatGroupDocument;
    
    constructor(apollo: Apollo.Apollo) {
      super(apollo);
    }
  }