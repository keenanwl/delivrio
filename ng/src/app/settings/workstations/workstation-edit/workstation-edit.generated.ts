/* eslint-disable */
import * as Types from '../../../../generated/graphql';

import { gql } from 'apollo-angular';
import { Injectable } from '@angular/core';
import * as Apollo from 'apollo-angular';
export type FetchWorkstationQueryVariables = Types.Exact<{
  id: Types.Scalars['ID'];
}>;


export type FetchWorkstationQuery = { workstation: { status: Types.WorkstationStatus, name: string, autoPrintReceiver: boolean, lastPing?: any | null, printer?: Array<{ id: string, name: string, labelPdf: boolean, labelZpl: boolean, labelPng: boolean, useShell: boolean, printSize: Types.PrinterPrintSize, document: boolean, rotate180: boolean, lastPing: any }> | null } };

export type UpdateWorkstationMutationVariables = Types.Exact<{
  id: Types.Scalars['ID'];
  input: Types.UpdateWorkstationInput;
}>;


export type UpdateWorkstationMutation = { updateWorkstation?: { id: string } | null };

export type ArchiveWorkstationMutationVariables = Types.Exact<{
  id: Types.Scalars['ID'];
}>;


export type ArchiveWorkstationMutation = { archiveWorkstation: boolean };

export const FetchWorkstationDocument = gql`
    query FetchWorkstation($id: ID!) {
  workstation(id: $id) {
    status
    name
    autoPrintReceiver
    lastPing
    printer {
      id
      name
      labelPdf
      labelZpl
      labelPng
      useShell
      printSize
      document
      rotate180
      lastPing
    }
  }
}
    `;

  @Injectable({
    providedIn: 'root'
  })
  export class FetchWorkstationGQL extends Apollo.Query<FetchWorkstationQuery, FetchWorkstationQueryVariables> {
    document = FetchWorkstationDocument;
    
    constructor(apollo: Apollo.Apollo) {
      super(apollo);
    }
  }
export const UpdateWorkstationDocument = gql`
    mutation UpdateWorkstation($id: ID!, $input: UpdateWorkstationInput!) {
  updateWorkstation(id: $id, input: $input) {
    id
  }
}
    `;

  @Injectable({
    providedIn: 'root'
  })
  export class UpdateWorkstationGQL extends Apollo.Mutation<UpdateWorkstationMutation, UpdateWorkstationMutationVariables> {
    document = UpdateWorkstationDocument;
    
    constructor(apollo: Apollo.Apollo) {
      super(apollo);
    }
  }
export const ArchiveWorkstationDocument = gql`
    mutation ArchiveWorkstation($id: ID!) {
  archiveWorkstation(id: $id)
}
    `;

  @Injectable({
    providedIn: 'root'
  })
  export class ArchiveWorkstationGQL extends Apollo.Mutation<ArchiveWorkstationMutation, ArchiveWorkstationMutationVariables> {
    document = ArchiveWorkstationDocument;
    
    constructor(apollo: Apollo.Apollo) {
      super(apollo);
    }
  }