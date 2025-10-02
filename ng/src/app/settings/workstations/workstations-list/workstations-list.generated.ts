/* eslint-disable */
import * as Types from '../../../../generated/graphql';

import { gql } from 'apollo-angular';
import { Injectable } from '@angular/core';
import * as Apollo from 'apollo-angular';
export type FetchWorkstationsQueryVariables = Types.Exact<{
  showArchived: Types.Scalars['Boolean'];
}>;


export type FetchWorkstationsQuery = { filteredWorkstations: Array<{ id: string, name: string, lastPing?: any | null, status: Types.WorkstationStatus, printer?: Array<{ id: string }> | null, selectedUser?: { name?: string | null, surname?: string | null } | null }> };

export type CreateWorkstationMutationVariables = Types.Exact<{
  input: Types.CreateWorkstationInput;
}>;


export type CreateWorkstationMutation = { createWorkstation?: { id: string, registrationToken: string, registrationTokenImg: string } | null };

export const FetchWorkstationsDocument = gql`
    query FetchWorkstations($showArchived: Boolean!) {
  filteredWorkstations(showArchived: $showArchived) {
    id
    name
    lastPing
    status
    printer {
      id
    }
    selectedUser {
      name
      surname
    }
  }
}
    `;

  @Injectable({
    providedIn: 'root'
  })
  export class FetchWorkstationsGQL extends Apollo.Query<FetchWorkstationsQuery, FetchWorkstationsQueryVariables> {
    document = FetchWorkstationsDocument;
    
    constructor(apollo: Apollo.Apollo) {
      super(apollo);
    }
  }
export const CreateWorkstationDocument = gql`
    mutation CreateWorkstation($input: CreateWorkstationInput!) {
  createWorkstation(input: $input) {
    id
    registrationToken
    registrationTokenImg
  }
}
    `;

  @Injectable({
    providedIn: 'root'
  })
  export class CreateWorkstationGQL extends Apollo.Mutation<CreateWorkstationMutation, CreateWorkstationMutationVariables> {
    document = CreateWorkstationDocument;
    
    constructor(apollo: Apollo.Apollo) {
      super(apollo);
    }
  }