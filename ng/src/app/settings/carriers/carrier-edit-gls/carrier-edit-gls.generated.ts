/* eslint-disable */
import * as Types from '../../../../generated/graphql';

import { gql } from 'apollo-angular';
import { Injectable } from '@angular/core';
import * as Apollo from 'apollo-angular';
export type FetchCarrierEditGlsQueryVariables = Types.Exact<{
  id: Types.Scalars['ID'];
}>;


export type FetchCarrierEditGlsQuery = { carrier?: { id: string, name: string, carrierGLS?: { contactID?: string | null, glsUsername?: string | null, glsPassword?: string | null, glsCountryCode?: string | null, customerID?: string | null, syncShipmentCancellation?: boolean | null, printErrorOnLabel?: boolean | null } | null } | null };

export type UpdateCarrierAgreementGlsMutationVariables = Types.Exact<{
  id: Types.Scalars['ID'];
  name: Types.Scalars['String'];
  input: Types.UpdateCarrierGlsInput;
}>;


export type UpdateCarrierAgreementGlsMutation = { updateCarrierAgreementGLS: { id: string } };

export const FetchCarrierEditGlsDocument = gql`
    query FetchCarrierEditGLS($id: ID!) {
  carrier(id: $id) {
    id
    name
    carrierGLS {
      contactID
      glsUsername
      glsPassword
      glsCountryCode
      customerID
      syncShipmentCancellation
      printErrorOnLabel
    }
  }
}
    `;

  @Injectable({
    providedIn: 'root'
  })
  export class FetchCarrierEditGlsGQL extends Apollo.Query<FetchCarrierEditGlsQuery, FetchCarrierEditGlsQueryVariables> {
    document = FetchCarrierEditGlsDocument;
    
    constructor(apollo: Apollo.Apollo) {
      super(apollo);
    }
  }
export const UpdateCarrierAgreementGlsDocument = gql`
    mutation UpdateCarrierAgreementGLS($id: ID!, $name: String!, $input: UpdateCarrierGLSInput!) {
  updateCarrierAgreementGLS(id: $id, name: $name, input: $input) {
    id
  }
}
    `;

  @Injectable({
    providedIn: 'root'
  })
  export class UpdateCarrierAgreementGlsGQL extends Apollo.Mutation<UpdateCarrierAgreementGlsMutation, UpdateCarrierAgreementGlsMutationVariables> {
    document = UpdateCarrierAgreementGlsDocument;
    
    constructor(apollo: Apollo.Apollo) {
      super(apollo);
    }
  }