/* eslint-disable */
import * as Types from '../generated/graphql';

import { gql } from 'apollo-angular';
import { Injectable } from '@angular/core';
import * as Apollo from 'apollo-angular';
export type FetchUserQueryVariables = Types.Exact<{ [key: string]: never; }>;


export type FetchUserQuery = { user?: { email: string, id: string, tenant: { id: string, name: string } } | null, buildInfo: { Hash: string, Time: string, LimitedSystem: boolean } };

export type SearchQueryVariables = Types.Exact<{
  term: Types.Scalars['String'];
}>;


export type SearchQuery = { search: Array<{ id: string, entity: Types.EntityType, title: string, imagePath?: string | null }> };

export type FetchSelectedWorkstationQueryVariables = Types.Exact<{ [key: string]: never; }>;


export type FetchSelectedWorkstationQuery = { selectedWorkstation?: { limitExceeded: boolean, workstation?: { id: string, name: string, lastPing?: any | null, status: Types.WorkstationStatus } | null, jobs: Array<{ id: string, status: Types.PrintJobStatus, printerMessages?: Array<string> | null, createdAt: any, printer: { name: string, workstation: { name: string } } }> } | null, user?: { pickupDay: Types.UserPickupDay } | null };

export type FetchSelectableWorkstationsQueryVariables = Types.Exact<{ [key: string]: never; }>;


export type FetchSelectableWorkstationsQuery = { workstations: { edges?: Array<{ node?: { id: string, name: string, lastPing?: any | null, status: Types.WorkstationStatus } | null } | null> | null } };

export type SaveSelectedWorkstationMutationVariables = Types.Exact<{
  id: Types.Scalars['ID'];
  pickupDay: Types.UserPickupDay;
}>;


export type SaveSelectedWorkstationMutation = { saveSelectedWorkstation: boolean };

export type SelectedWorkstationFragment = { id: string, name: string, lastPing?: any | null, status: Types.WorkstationStatus };

export const SelectedWorkstationFragmentDoc = gql`
    fragment SelectedWorkstation on Workstation {
  id
  name
  lastPing
  status
}
    `;
export const FetchUserDocument = gql`
    query FetchUser {
  user {
    email
    id
    tenant {
      id
      name
    }
  }
  buildInfo {
    Hash
    Time
    LimitedSystem
  }
}
    `;

  @Injectable({
    providedIn: 'root'
  })
  export class FetchUserGQL extends Apollo.Query<FetchUserQuery, FetchUserQueryVariables> {
    document = FetchUserDocument;
    
    constructor(apollo: Apollo.Apollo) {
      super(apollo);
    }
  }
export const SearchDocument = gql`
    query Search($term: String!) {
  search(term: $term) {
    id
    entity
    title
    imagePath
  }
}
    `;

  @Injectable({
    providedIn: 'root'
  })
  export class SearchGQL extends Apollo.Query<SearchQuery, SearchQueryVariables> {
    document = SearchDocument;
    
    constructor(apollo: Apollo.Apollo) {
      super(apollo);
    }
  }
export const FetchSelectedWorkstationDocument = gql`
    query FetchSelectedWorkstation {
  selectedWorkstation {
    workstation {
      ...SelectedWorkstation
    }
    limitExceeded
    jobs {
      id
      status
      printerMessages
      createdAt
      printer {
        name
        workstation {
          name
        }
      }
    }
  }
  user {
    pickupDay
  }
}
    ${SelectedWorkstationFragmentDoc}`;

  @Injectable({
    providedIn: 'root'
  })
  export class FetchSelectedWorkstationGQL extends Apollo.Query<FetchSelectedWorkstationQuery, FetchSelectedWorkstationQueryVariables> {
    document = FetchSelectedWorkstationDocument;
    
    constructor(apollo: Apollo.Apollo) {
      super(apollo);
    }
  }
export const FetchSelectableWorkstationsDocument = gql`
    query FetchSelectableWorkstations {
  workstations {
    edges {
      node {
        ...SelectedWorkstation
      }
    }
  }
}
    ${SelectedWorkstationFragmentDoc}`;

  @Injectable({
    providedIn: 'root'
  })
  export class FetchSelectableWorkstationsGQL extends Apollo.Query<FetchSelectableWorkstationsQuery, FetchSelectableWorkstationsQueryVariables> {
    document = FetchSelectableWorkstationsDocument;
    
    constructor(apollo: Apollo.Apollo) {
      super(apollo);
    }
  }
export const SaveSelectedWorkstationDocument = gql`
    mutation SaveSelectedWorkstation($id: ID!, $pickupDay: UserPickupDay!) {
  saveSelectedWorkstation(id: $id, pickupDay: $pickupDay)
}
    `;

  @Injectable({
    providedIn: 'root'
  })
  export class SaveSelectedWorkstationGQL extends Apollo.Mutation<SaveSelectedWorkstationMutation, SaveSelectedWorkstationMutationVariables> {
    document = SaveSelectedWorkstationDocument;
    
    constructor(apollo: Apollo.Apollo) {
      super(apollo);
    }
  }