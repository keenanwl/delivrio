/* eslint-disable */
import * as Types from '../../../generated/graphql';

import { gql } from 'apollo-angular';
import { Injectable } from '@angular/core';
import * as Apollo from 'apollo-angular';
export type HistoryLogsQueryVariables = Types.Exact<{ [key: string]: never; }>;


export type HistoryLogsQuery = { historyLogs: { histories: Array<{ id: string, createdAt: any, origin: Types.ChangeHistoryOrigin, returnColliHistory?: Array<{ description: string }> | null, orderHistory?: Array<{ description: string }> | null, user?: { id: string, name?: string | null, surname?: string | null } | null }>, system_event: Array<{ id: string, eventType: Types.SystemEventsEventType, status: Types.SystemEventsStatus, createdAt?: any | null, description: string, data?: string | null }> } };

export const HistoryLogsDocument = gql`
    query HistoryLogs {
  historyLogs {
    histories {
      id
      createdAt
      origin
      returnColliHistory {
        description
      }
      orderHistory {
        description
      }
      user {
        id
        name
        surname
      }
    }
    system_event {
      id
      eventType
      status
      createdAt
      description
      data
    }
  }
}
    `;

  @Injectable({
    providedIn: 'root'
  })
  export class HistoryLogsGQL extends Apollo.Query<HistoryLogsQuery, HistoryLogsQueryVariables> {
    document = HistoryLogsDocument;
    
    constructor(apollo: Apollo.Apollo) {
      super(apollo);
    }
  }