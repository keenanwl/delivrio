/* eslint-disable */
import * as Types from '../../../../generated/graphql';

import { gql } from 'apollo-angular';
import { Injectable } from '@angular/core';
import * as Apollo from 'apollo-angular';
export type FetchEmailTemapltesQueryVariables = Types.Exact<{ [key: string]: never; }>;


export type FetchEmailTemapltesQuery = { emailTemplates: { edges?: Array<{ node?: { id: string, name: string, subject: string, mergeType: Types.EmailTemplateMergeType } | null } | null> | null } };

export type CreateEmailTemplateMutationVariables = Types.Exact<{
  name: Types.Scalars['String'];
  merge: Types.EmailTemplateMergeType;
}>;


export type CreateEmailTemplateMutation = { createEmailTemplates: string };

export const FetchEmailTemapltesDocument = gql`
    query FetchEmailTemapltes {
  emailTemplates {
    edges {
      node {
        id
        name
        subject
        mergeType
      }
    }
  }
}
    `;

  @Injectable({
    providedIn: 'root'
  })
  export class FetchEmailTemapltesGQL extends Apollo.Query<FetchEmailTemapltesQuery, FetchEmailTemapltesQueryVariables> {
    document = FetchEmailTemapltesDocument;
    
    constructor(apollo: Apollo.Apollo) {
      super(apollo);
    }
  }
export const CreateEmailTemplateDocument = gql`
    mutation CreateEmailTemplate($name: String!, $merge: EmailTemplateMergeType!) {
  createEmailTemplates(name: $name, merge: $merge)
}
    `;

  @Injectable({
    providedIn: 'root'
  })
  export class CreateEmailTemplateGQL extends Apollo.Mutation<CreateEmailTemplateMutation, CreateEmailTemplateMutationVariables> {
    document = CreateEmailTemplateDocument;
    
    constructor(apollo: Apollo.Apollo) {
      super(apollo);
    }
  }