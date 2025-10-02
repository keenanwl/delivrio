/* eslint-disable */
import * as Types from '../../../../generated/graphql';

import { gql } from 'apollo-angular';
import { Injectable } from '@angular/core';
import * as Apollo from 'apollo-angular';
export type FetchDocumentQueryVariables = Types.Exact<{
  id: Types.Scalars['ID'];
}>;


export type FetchDocumentQuery = { document: { name: string, endAt: any, startAt: any, htmlTemplate?: string | null, htmlFooter?: string | null, htmlHeader?: string | null, paperSize: Types.DocumentPaperSize, mergeType: Types.DocumentMergeType, carrierBrand?: { id: string } | null }, carrierBrands: { edges?: Array<{ node?: { id: string, label: string } | null } | null> | null } };

export type UpdateDocumentMutationVariables = Types.Exact<{
  id: Types.Scalars['ID'];
  input: Types.UpdateDocumentInput;
}>;


export type UpdateDocumentMutation = { updateDocument: { id: string } };

export type DownloadDocumentQueryVariables = Types.Exact<{
  id: Types.Scalars['ID'];
}>;


export type DownloadDocumentQuery = { documentDownload: { base64PDF: string } };

export const FetchDocumentDocument = gql`
    query FetchDocument($id: ID!) {
  document(id: $id) {
    name
    endAt
    startAt
    htmlTemplate
    htmlFooter
    htmlHeader
    carrierBrand {
      id
    }
    paperSize
    mergeType
  }
  carrierBrands {
    edges {
      node {
        id
        label
      }
    }
  }
}
    `;

  @Injectable({
    providedIn: 'root'
  })
  export class FetchDocumentGQL extends Apollo.Query<FetchDocumentQuery, FetchDocumentQueryVariables> {
    document = FetchDocumentDocument;
    
    constructor(apollo: Apollo.Apollo) {
      super(apollo);
    }
  }
export const UpdateDocumentDocument = gql`
    mutation UpdateDocument($id: ID!, $input: UpdateDocumentInput!) {
  updateDocument(id: $id, input: $input) {
    id
  }
}
    `;

  @Injectable({
    providedIn: 'root'
  })
  export class UpdateDocumentGQL extends Apollo.Mutation<UpdateDocumentMutation, UpdateDocumentMutationVariables> {
    document = UpdateDocumentDocument;
    
    constructor(apollo: Apollo.Apollo) {
      super(apollo);
    }
  }
export const DownloadDocumentDocument = gql`
    query DownloadDocument($id: ID!) {
  documentDownload(id: $id) {
    base64PDF
  }
}
    `;

  @Injectable({
    providedIn: 'root'
  })
  export class DownloadDocumentGQL extends Apollo.Query<DownloadDocumentQuery, DownloadDocumentQueryVariables> {
    document = DownloadDocumentDocument;
    
    constructor(apollo: Apollo.Apollo) {
      super(apollo);
    }
  }