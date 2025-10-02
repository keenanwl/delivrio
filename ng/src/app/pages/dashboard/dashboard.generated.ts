/* eslint-disable */
import * as Types from '../../../generated/graphql';

import { gql } from 'apollo-angular';
import { Injectable } from '@angular/core';
import * as Apollo from 'apollo-angular';
export type FetchDashboardTilesQueryVariables = Types.Exact<{ [key: string]: never; }>;


export type FetchDashboardTilesQuery = { trailingProductUpdates: Array<number>, dashboardTiles: Array<{ title: string, value: string }>, hypothesisTestResultsDashboard: Array<{ id: string, Name: string, ControlSuccess: number, ControlFailure: number, TestSuccess: number, TestFailure: number, SignificantlyDifferent: boolean, ControlWin: number, TestWin: number }>, rateRequests: { requests: Array<{ date: string, optionCount: number, req?: string | null }>, requestsError: Array<{ date: string, optionCount: number, req?: string | null, error?: string | null }> } };

export const FetchDashboardTilesDocument = gql`
    query FetchDashboardTiles {
  dashboardTiles {
    title
    value
  }
  hypothesisTestResultsDashboard {
    id
    Name
    ControlSuccess
    ControlFailure
    TestSuccess
    TestFailure
    SignificantlyDifferent
    ControlWin
    TestWin
  }
  trailingProductUpdates
  rateRequests {
    requests {
      date
      optionCount
      req
    }
    requestsError {
      date
      optionCount
      req
      error
    }
  }
}
    `;

  @Injectable({
    providedIn: 'root'
  })
  export class FetchDashboardTilesGQL extends Apollo.Query<FetchDashboardTilesQuery, FetchDashboardTilesQueryVariables> {
    document = FetchDashboardTilesDocument;
    
    constructor(apollo: Apollo.Apollo) {
      super(apollo);
    }
  }