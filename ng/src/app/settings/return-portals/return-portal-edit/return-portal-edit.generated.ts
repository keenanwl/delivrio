/* eslint-disable */
import * as Types from '../../../../generated/graphql';

import { gql } from 'apollo-angular';
import { Injectable } from '@angular/core';
import * as Apollo from 'apollo-angular';
export type FetchReturnPortalQueryVariables = Types.Exact<{
  id: Types.Scalars['ID'];
}>;


export type FetchReturnPortalQuery = { returnPortal: { name: string, returnOpenHours: number, automaticallyAccept: boolean, returnLocation?: Array<{ id: string }> | null, returnPortalClaim?: Array<{ id: string, name: string, description: string, restockable: boolean }> | null, connection?: { id: string } | null, emailConfirmationLabel?: { id: string } | null, emailConfirmationQrCode?: { id: string } | null, emailReceived?: { id: string } | null, emailAccepted?: { id: string } | null, deliveryOptions?: Array<{ id: string }> | null }, connections: { edges?: Array<{ node?: { id: string, name: string } | null } | null> | null }, emailTemplates: { edges?: Array<{ node?: { id: string, name: string, mergeType: Types.EmailTemplateMergeType } | null } | null> | null }, deliveryOptions: { edges?: Array<{ node?: { id: string, name: string } | null } | null> | null } };

export type UpdateReturnPortalMutationVariables = Types.Exact<{
  id: Types.Scalars['ID'];
  input: Types.UpdateReturnPortalInput;
  inputClaims: Array<Types.MutateReturnPortalClaim> | Types.MutateReturnPortalClaim;
}>;


export type UpdateReturnPortalMutation = { updateReturnPortal: { id: string } };

export const FetchReturnPortalDocument = gql`
    query FetchReturnPortal($id: ID!) {
  returnPortal(id: $id) {
    name
    returnOpenHours
    automaticallyAccept
    returnLocation {
      id
    }
    returnPortalClaim {
      id
      name
      description
      restockable
    }
    connection {
      id
    }
    emailConfirmationLabel {
      id
    }
    emailConfirmationQrCode {
      id
    }
    emailReceived {
      id
    }
    emailAccepted {
      id
    }
    deliveryOptions {
      id
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
  deliveryOptions(where: {hasCarrierServiceWith: {return: true}}) {
    edges {
      node {
        id
        name
      }
    }
  }
}
    `;

  @Injectable({
    providedIn: 'root'
  })
  export class FetchReturnPortalGQL extends Apollo.Query<FetchReturnPortalQuery, FetchReturnPortalQueryVariables> {
    document = FetchReturnPortalDocument;
    
    constructor(apollo: Apollo.Apollo) {
      super(apollo);
    }
  }
export const UpdateReturnPortalDocument = gql`
    mutation UpdateReturnPortal($id: ID!, $input: UpdateReturnPortalInput!, $inputClaims: [MutateReturnPortalClaim!]!) {
  updateReturnPortal(id: $id, input: $input, inputClaims: $inputClaims) {
    id
  }
}
    `;

  @Injectable({
    providedIn: 'root'
  })
  export class UpdateReturnPortalGQL extends Apollo.Mutation<UpdateReturnPortalMutation, UpdateReturnPortalMutationVariables> {
    document = UpdateReturnPortalDocument;
    
    constructor(apollo: Apollo.Apollo) {
      super(apollo);
    }
  }