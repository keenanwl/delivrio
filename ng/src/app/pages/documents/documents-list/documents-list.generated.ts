/* eslint-disable */
import * as Types from '../../../../generated/graphql';

import { gql } from 'apollo-angular';
import { Injectable } from '@angular/core';
import * as Apollo from 'apollo-angular';
export type FetchDocumentsQueryVariables = Types.Exact<{ [key: string]: never; }>;


export type FetchDocumentsQuery = { documents: { edges?: Array<{ node?: { id: string, name: string, mergeType: Types.DocumentMergeType, createdAt: any, carrierBrand?: { label: string } | null } | null } | null> | null } };

export type CreateDocumentMutationVariables = Types.Exact<{
  name: Types.Scalars['String'];
  mergeType: Types.DocumentMergeType;
}>;


export type CreateDocumentMutation = { createDocument: string };

export const FetchDocumentsDocument = gql`
    query FetchDocuments {
  documents {
    edges {
      node {
        id
        name
        carrierBrand {
          label
        }
        mergeType
        createdAt
      }
    }
  }
}
    `;

  @Injectable({
    providedIn: 'root'
  })
  export class FetchDocumentsGQL extends Apollo.Query<FetchDocumentsQuery, FetchDocumentsQueryVariables> {
    document = FetchDocumentsDocument;
    
    constructor(apollo: Apollo.Apollo) {
      super(apollo);
    }
  }
export const CreateDocumentDocument = gql`
    mutation CreateDocument($name: String!, $mergeType: DocumentMergeType!) {
  createDocument(name: $name, mergeType: $mergeType)
}
    `;

  @Injectable({
    providedIn: 'root'
  })
  export class CreateDocumentGQL extends Apollo.Mutation<CreateDocumentMutation, CreateDocumentMutationVariables> {
    document = CreateDocumentDocument;
    
    constructor(apollo: Apollo.Apollo) {
      super(apollo);
    }
  }