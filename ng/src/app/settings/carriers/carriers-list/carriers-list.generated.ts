/* eslint-disable */
import * as Types from '../../../../generated/graphql';

import { gql } from 'apollo-angular';
import { Injectable } from '@angular/core';
import * as Apollo from 'apollo-angular';
export type ListCarriersQueryVariables = Types.Exact<{ [key: string]: never; }>;


export type ListCarriersQuery = { carriers: { edges?: Array<{ node?: { id: string, name: string, carrierBrand: { id: string, internalID: Types.CarrierBrandInternalId, label: string, logoURL?: string | null, backgroundColor?: string | null, textColor?: string | null } } | null } | null> | null } };

export type CarrierBrandsQueryVariables = Types.Exact<{ [key: string]: never; }>;


export type CarrierBrandsQuery = { carrierBrands: { edges?: Array<{ node?: { id: string, label: string, internalID: Types.CarrierBrandInternalId, logoURL?: string | null } | null } | null> | null } };

export type CreateCarrierAgreementMutationVariables = Types.Exact<{
  name: Types.Scalars['String'];
  carrierBrand: Types.Scalars['ID'];
}>;


export type CreateCarrierAgreementMutation = { createCarrierAgreement: { id: string, carrier: Types.CarrierBrandInternalId } };

export const ListCarriersDocument = gql`
    query ListCarriers {
  carriers {
    edges {
      node {
        id
        name
        carrierBrand {
          id
          internalID
          label
          logoURL
          backgroundColor
          textColor
        }
      }
    }
  }
}
    `;

  @Injectable({
    providedIn: 'root'
  })
  export class ListCarriersGQL extends Apollo.Query<ListCarriersQuery, ListCarriersQueryVariables> {
    document = ListCarriersDocument;
    
    constructor(apollo: Apollo.Apollo) {
      super(apollo);
    }
  }
export const CarrierBrandsDocument = gql`
    query CarrierBrands {
  carrierBrands {
    edges {
      node {
        id
        label
        internalID
        logoURL
      }
    }
  }
}
    `;

  @Injectable({
    providedIn: 'root'
  })
  export class CarrierBrandsGQL extends Apollo.Query<CarrierBrandsQuery, CarrierBrandsQueryVariables> {
    document = CarrierBrandsDocument;
    
    constructor(apollo: Apollo.Apollo) {
      super(apollo);
    }
  }
export const CreateCarrierAgreementDocument = gql`
    mutation CreateCarrierAgreement($name: String!, $carrierBrand: ID!) {
  createCarrierAgreement(name: $name, carrierBrand: $carrierBrand) {
    id
    carrier
  }
}
    `;

  @Injectable({
    providedIn: 'root'
  })
  export class CreateCarrierAgreementGQL extends Apollo.Mutation<CreateCarrierAgreementMutation, CreateCarrierAgreementMutationVariables> {
    document = CreateCarrierAgreementDocument;
    
    constructor(apollo: Apollo.Apollo) {
      super(apollo);
    }
  }