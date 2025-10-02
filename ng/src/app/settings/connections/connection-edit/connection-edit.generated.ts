/* eslint-disable */
import * as Types from '../../../../generated/graphql';

import { gql } from 'apollo-angular';
import { Injectable } from '@angular/core';
import * as Apollo from 'apollo-angular';
export type FetchConnectionQueryVariables = Types.Exact<{
  id: Types.Scalars['ID'];
}>;


export type FetchConnectionQuery = { connection?: { name: string, convertCurrency: boolean, autoPrintParcelSlip: boolean, dispatchAutomatically: boolean, syncOrders: boolean, syncProducts: boolean, fulfillAutomatically: boolean, currency: { id: string, display: string }, packingSlipTemplate?: { id: string } | null, sellerLocation: { id: string }, senderLocation: { id: string }, returnLocation: { id: string }, pickupLocation: { id: string }, connectionShopify?: { storeURL?: string | null, apiKey?: string | null, rateIntegration: boolean, syncFrom?: any | null, filterTags?: Array<string> | null } | null, defaultDeliveryOption?: { id: string, name: string, description?: string | null } | null } | null, locations: { edges?: Array<{ node?: { id: string, name: string, address: { email: string, firstName: string, lastName: string, phoneNumber: string, vatNumber?: string | null, addressOne: string, addressTwo: string, zip: string, city: string, company?: string | null, country: { id: string, label: string, alpha2: string } }, locationTags: Array<{ id: string, internalID: string }> } | null } | null> | null }, deliveryOptions: { edges?: Array<{ node?: { id: string, name: string, description?: string | null } | null } | null> | null }, documents: { edges?: Array<{ node?: { id: string, name: string } | null } | null> | null }, currencies: { edges?: Array<{ node?: { id: string, display: string } | null } | null> | null } };

export type ConnectionDeliveryOptionFragment = { id: string, name: string, description?: string | null };

export type LocationFragment = { id: string, name: string, address: { email: string, firstName: string, lastName: string, phoneNumber: string, vatNumber?: string | null, addressOne: string, addressTwo: string, zip: string, city: string, company?: string | null, country: { id: string, label: string, alpha2: string } }, locationTags: Array<{ id: string, internalID: string }> };

export type ConnectionBrandsQueryVariables = Types.Exact<{ [key: string]: never; }>;


export type ConnectionBrandsQuery = { connectionBrands: { edges?: Array<{ node?: { id: string, label: string, logoURL?: string | null } | null } | null> | null } };

export type CreateConnectionShopifyMutationVariables = Types.Exact<{
  input: Types.CreateConnectionShopifyInput;
  inputConnection: Types.CreateConnectionInput;
}>;


export type CreateConnectionShopifyMutation = { createShopifyConnection?: { id: string } | null };

export type UpdateConnectionShopifyMutationVariables = Types.Exact<{
  id: Types.Scalars['ID'];
  input: Types.UpdateConnectionShopifyInput;
  inputConnection: Types.UpdateConnectionInput;
}>;


export type UpdateConnectionShopifyMutation = { updateShopifyConnection?: { id: string } | null };

export const ConnectionDeliveryOptionFragmentDoc = gql`
    fragment ConnectionDeliveryOption on DeliveryOption {
  id
  name
  description
}
    `;
export const LocationFragmentDoc = gql`
    fragment Location on Location {
  id
  name
  address {
    email
    firstName
    lastName
    phoneNumber
    vatNumber
    addressOne
    addressTwo
    zip
    city
    country {
      id
      label
      alpha2
    }
    company
  }
  locationTags {
    id
    internalID
  }
}
    `;
export const FetchConnectionDocument = gql`
    query FetchConnection($id: ID!) {
  connection(id: $id) {
    name
    convertCurrency
    autoPrintParcelSlip
    currency {
      id
      display
    }
    dispatchAutomatically
    syncOrders
    syncProducts
    fulfillAutomatically
    packingSlipTemplate {
      id
    }
    sellerLocation {
      id
    }
    senderLocation {
      id
    }
    returnLocation {
      id
    }
    pickupLocation {
      id
    }
    connectionShopify {
      storeURL
      apiKey
      rateIntegration
      syncFrom
      filterTags
    }
    defaultDeliveryOption {
      ...ConnectionDeliveryOption
    }
  }
  locations {
    edges {
      node {
        ...Location
      }
    }
  }
  deliveryOptions(where: {hasConnectionWith: {id: $id}}) {
    edges {
      node {
        ...ConnectionDeliveryOption
      }
    }
  }
  documents(where: {mergeType: PackingSlip}) {
    edges {
      node {
        id
        name
      }
    }
  }
  currencies {
    edges {
      node {
        id
        display
      }
    }
  }
}
    ${ConnectionDeliveryOptionFragmentDoc}
${LocationFragmentDoc}`;

  @Injectable({
    providedIn: 'root'
  })
  export class FetchConnectionGQL extends Apollo.Query<FetchConnectionQuery, FetchConnectionQueryVariables> {
    document = FetchConnectionDocument;
    
    constructor(apollo: Apollo.Apollo) {
      super(apollo);
    }
  }
export const ConnectionBrandsDocument = gql`
    query ConnectionBrands {
  connectionBrands {
    edges {
      node {
        id
        label
        logoURL
      }
    }
  }
}
    `;

  @Injectable({
    providedIn: 'root'
  })
  export class ConnectionBrandsGQL extends Apollo.Query<ConnectionBrandsQuery, ConnectionBrandsQueryVariables> {
    document = ConnectionBrandsDocument;
    
    constructor(apollo: Apollo.Apollo) {
      super(apollo);
    }
  }
export const CreateConnectionShopifyDocument = gql`
    mutation CreateConnectionShopify($input: CreateConnectionShopifyInput!, $inputConnection: CreateConnectionInput!) {
  createShopifyConnection(input: $input, inputConnection: $inputConnection) {
    id
  }
}
    `;

  @Injectable({
    providedIn: 'root'
  })
  export class CreateConnectionShopifyGQL extends Apollo.Mutation<CreateConnectionShopifyMutation, CreateConnectionShopifyMutationVariables> {
    document = CreateConnectionShopifyDocument;
    
    constructor(apollo: Apollo.Apollo) {
      super(apollo);
    }
  }
export const UpdateConnectionShopifyDocument = gql`
    mutation UpdateConnectionShopify($id: ID!, $input: UpdateConnectionShopifyInput!, $inputConnection: UpdateConnectionInput!) {
  updateShopifyConnection(
    id: $id
    input: $input
    inputConnection: $inputConnection
  ) {
    id
  }
}
    `;

  @Injectable({
    providedIn: 'root'
  })
  export class UpdateConnectionShopifyGQL extends Apollo.Mutation<UpdateConnectionShopifyMutation, UpdateConnectionShopifyMutationVariables> {
    document = UpdateConnectionShopifyDocument;
    
    constructor(apollo: Apollo.Apollo) {
      super(apollo);
    }
  }