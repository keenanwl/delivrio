/* eslint-disable */
import * as Types from '../../../generated/graphql';

import { gql } from 'apollo-angular';
import { Injectable } from '@angular/core';
import * as Apollo from 'apollo-angular';
export type FetchUserPlanQueryVariables = Types.Exact<{ [key: string]: never; }>;


export type FetchUserPlanQuery = { user?: { tenant: { plan: { id: string, label: string } }, language?: { id: string, label: string } | null } | null };

export type FetchPlanListQueryVariables = Types.Exact<{ [key: string]: never; }>;


export type FetchPlanListQuery = { plans: { edges?: Array<{ node?: { id: string, label: string, priceDkk: number } | null } | null> | null } };

export type UpdatePlanMutationVariables = Types.Exact<{
  planID: Types.Scalars['ID'];
}>;


export type UpdatePlanMutation = { updatePlan?: { id: string } | null };

export const FetchUserPlanDocument = gql`
    query FetchUserPlan {
  user {
    tenant {
      plan {
        id
        label
      }
    }
    language {
      id
      label
    }
  }
}
    `;

  @Injectable({
    providedIn: 'root'
  })
  export class FetchUserPlanGQL extends Apollo.Query<FetchUserPlanQuery, FetchUserPlanQueryVariables> {
    document = FetchUserPlanDocument;
    
    constructor(apollo: Apollo.Apollo) {
      super(apollo);
    }
  }
export const FetchPlanListDocument = gql`
    query FetchPlanList {
  plans {
    edges {
      node {
        id
        label
        priceDkk
      }
    }
  }
}
    `;

  @Injectable({
    providedIn: 'root'
  })
  export class FetchPlanListGQL extends Apollo.Query<FetchPlanListQuery, FetchPlanListQueryVariables> {
    document = FetchPlanListDocument;
    
    constructor(apollo: Apollo.Apollo) {
      super(apollo);
    }
  }
export const UpdatePlanDocument = gql`
    mutation UpdatePlan($planID: ID!) {
  updatePlan(planID: $planID) {
    id
  }
}
    `;

  @Injectable({
    providedIn: 'root'
  })
  export class UpdatePlanGQL extends Apollo.Mutation<UpdatePlanMutation, UpdatePlanMutationVariables> {
    document = UpdatePlanDocument;
    
    constructor(apollo: Apollo.Apollo) {
      super(apollo);
    }
  }