/* eslint-disable */
import * as Types from '../../../../generated/graphql';

import { gql } from 'apollo-angular';
import { Injectable } from '@angular/core';
import * as Apollo from 'apollo-angular';
export type FetchEmailTemplateQueryVariables = Types.Exact<{
  id: Types.Scalars['ID'];
}>;


export type FetchEmailTemplateQuery = { emailTemplate: { name: string, subject: string, htmlTemplate: string, mergeType: Types.EmailTemplateMergeType } };

export type UpdateEmailTemplateMutationVariables = Types.Exact<{
  id: Types.Scalars['ID'];
  input: Types.UpdateEmailTemplateInput;
}>;


export type UpdateEmailTemplateMutation = { updateEmailTemplate: { id: string } };

export type FireTestEmailQueryVariables = Types.Exact<{
  id: Types.Scalars['ID'];
  toEmail: Types.Scalars['String'];
}>;


export type FireTestEmailQuery = { sendTestEmail: boolean };

export const FetchEmailTemplateDocument = gql`
    query FetchEmailTemplate($id: ID!) {
  emailTemplate(id: $id) {
    name
    subject
    htmlTemplate
    mergeType
  }
}
    `;

  @Injectable({
    providedIn: 'root'
  })
  export class FetchEmailTemplateGQL extends Apollo.Query<FetchEmailTemplateQuery, FetchEmailTemplateQueryVariables> {
    document = FetchEmailTemplateDocument;
    
    constructor(apollo: Apollo.Apollo) {
      super(apollo);
    }
  }
export const UpdateEmailTemplateDocument = gql`
    mutation UpdateEmailTemplate($id: ID!, $input: UpdateEmailTemplateInput!) {
  updateEmailTemplate(id: $id, input: $input) {
    id
  }
}
    `;

  @Injectable({
    providedIn: 'root'
  })
  export class UpdateEmailTemplateGQL extends Apollo.Mutation<UpdateEmailTemplateMutation, UpdateEmailTemplateMutationVariables> {
    document = UpdateEmailTemplateDocument;
    
    constructor(apollo: Apollo.Apollo) {
      super(apollo);
    }
  }
export const FireTestEmailDocument = gql`
    query FireTestEmail($id: ID!, $toEmail: String!) {
  sendTestEmail(id: $id, toEmail: $toEmail)
}
    `;

  @Injectable({
    providedIn: 'root'
  })
  export class FireTestEmailGQL extends Apollo.Query<FireTestEmailQuery, FireTestEmailQueryVariables> {
    document = FireTestEmailDocument;
    
    constructor(apollo: Apollo.Apollo) {
      super(apollo);
    }
  }