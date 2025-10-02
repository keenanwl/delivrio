/* eslint-disable */
import * as Types from '../../../../generated/graphql';

import { gql } from 'apollo-angular';
import { Injectable } from '@angular/core';
import * as Apollo from 'apollo-angular';
export type FetchNotificationsQueryVariables = Types.Exact<{ [key: string]: never; }>;


export type FetchNotificationsQuery = { notifications: { edges?: Array<{ node?: { id: string, name: string, active: boolean, connection: { id: string, name: string }, emailTemplate: { id: string, name: string, mergeType: Types.EmailTemplateMergeType } } | null } | null> | null }, connections: { edges?: Array<{ node?: { id: string, name: string } | null } | null> | null }, emailTemplates: { edges?: Array<{ node?: { id: string, name: string, mergeType: Types.EmailTemplateMergeType } | null } | null> | null } };

export type NotificationsListItemFragment = { id: string, name: string, active: boolean, connection: { id: string, name: string }, emailTemplate: { id: string, name: string, mergeType: Types.EmailTemplateMergeType } };

export type CreateNotificationMutationVariables = Types.Exact<{
  name: Types.Scalars['String'];
  connectionID: Types.Scalars['ID'];
  emailTemplateID: Types.Scalars['ID'];
}>;


export type CreateNotificationMutation = { createNotification: Array<{ id: string, name: string, active: boolean, connection: { id: string, name: string }, emailTemplate: { id: string, name: string, mergeType: Types.EmailTemplateMergeType } }> };

export type ToggleNotificationMutationVariables = Types.Exact<{
  notificationID: Types.Scalars['ID'];
  checked: Types.Scalars['Boolean'];
}>;


export type ToggleNotificationMutation = { toggleNotification: boolean };

export const NotificationsListItemFragmentDoc = gql`
    fragment NotificationsListItem on Notification {
  id
  name
  active
  connection {
    id
    name
  }
  emailTemplate {
    id
    name
    mergeType
  }
}
    `;
export const FetchNotificationsDocument = gql`
    query FetchNotifications {
  notifications {
    edges {
      node {
        ...NotificationsListItem
      }
    }
  }
  connections {
    edges {
      node {
        id
        name
      }
    }
  }
  emailTemplates {
    edges {
      node {
        id
        name
        mergeType
      }
    }
  }
}
    ${NotificationsListItemFragmentDoc}`;

  @Injectable({
    providedIn: 'root'
  })
  export class FetchNotificationsGQL extends Apollo.Query<FetchNotificationsQuery, FetchNotificationsQueryVariables> {
    document = FetchNotificationsDocument;
    
    constructor(apollo: Apollo.Apollo) {
      super(apollo);
    }
  }
export const CreateNotificationDocument = gql`
    mutation CreateNotification($name: String!, $connectionID: ID!, $emailTemplateID: ID!) {
  createNotification(
    name: $name
    connectionID: $connectionID
    emailTemplateID: $emailTemplateID
  ) {
    ...NotificationsListItem
  }
}
    ${NotificationsListItemFragmentDoc}`;

  @Injectable({
    providedIn: 'root'
  })
  export class CreateNotificationGQL extends Apollo.Mutation<CreateNotificationMutation, CreateNotificationMutationVariables> {
    document = CreateNotificationDocument;
    
    constructor(apollo: Apollo.Apollo) {
      super(apollo);
    }
  }
export const ToggleNotificationDocument = gql`
    mutation ToggleNotification($notificationID: ID!, $checked: Boolean!) {
  toggleNotification(checked: $checked, notificationID: $notificationID)
}
    `;

  @Injectable({
    providedIn: 'root'
  })
  export class ToggleNotificationGQL extends Apollo.Mutation<ToggleNotificationMutation, ToggleNotificationMutationVariables> {
    document = ToggleNotificationDocument;
    
    constructor(apollo: Apollo.Apollo) {
      super(apollo);
    }
  }