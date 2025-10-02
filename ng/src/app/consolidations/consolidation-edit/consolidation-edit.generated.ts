/* eslint-disable */
import * as Types from '../../../generated/graphql';

import { gql } from 'apollo-angular';
import { AddressInfoFragmentDoc, OrderLinesFragmentDoc } from '../../orders/order-edit/order-edit.generated';
import { Injectable } from '@angular/core';
import * as Apollo from 'apollo-angular';
export type FetchConsolidationQueryVariables = Types.Exact<{
  id: Types.Scalars['ID'];
}>;


export type FetchConsolidationQuery = { consolidation: { publicID: string, description?: string | null, status: Types.ConsolidationStatus, sender?: { id: string, firstName: string, lastName: string, phoneNumber: string, email: string, addressOne: string, addressTwo: string, zip: string, city: string, state?: string | null, company?: string | null, country: { id: string, alpha2: string, label: string } } | null, recipient?: { id: string, firstName: string, lastName: string, phoneNumber: string, email: string, addressOne: string, addressTwo: string, zip: string, city: string, state?: string | null, company?: string | null, country: { id: string, alpha2: string, label: string } } | null, deliveryOption?: { id: string } | null, orders?: Array<{ id: string, orderPublicID: string, status: Types.OrderStatus }> | null, pallets?: Array<{ id: string, publicID: string, description: string, packaging?: { id: string } | null, orders?: Array<{ id: string, orderPublicID: string, status: Types.OrderStatus }> | null }> | null }, consolidationShipments: { mayBook: boolean, mayPrebook: boolean, shipment?: { id: string, shipmentPallet?: Array<{ id: string, pallet?: { id: string } | null }> | null } | null }, locationInfo: { sender?: { id: string, firstName: string, lastName: string, phoneNumber: string, email: string, addressOne: string, addressTwo: string, zip: string, city: string, state?: string | null, company?: string | null, country: { id: string, alpha2: string, label: string } } | null, recipient?: { id: string, firstName: string, lastName: string, phoneNumber: string, email: string, addressOne: string, addressTwo: string, zip: string, city: string, state?: string | null, company?: string | null, country: { id: string, alpha2: string, label: string } } | null }, deliveryOptions: { edges?: Array<{ node?: { id: string, name: string, carrierService: { return: boolean } } | null } | null> | null }, packagings: { edges?: Array<{ node?: { id: string, name: string } | null } | null> | null } };

export type ConsolidationOrderFragment = { id: string, orderPublicID: string, status: Types.OrderStatus };

export type ConsolidationSearchOrdersQueryVariables = Types.Exact<{
  term: Types.Scalars['String'];
}>;


export type ConsolidationSearchOrdersQuery = { orders: { edges?: Array<{ node?: { id: string, orderPublicID: string, status: Types.OrderStatus } | null } | null> | null } };

export type UpdateConsolidationMutationVariables = Types.Exact<{
  id: Types.Scalars['ID'];
  input: Types.UpdateConsolidationInput;
  sender?: Types.InputMaybe<Types.CreateAddressInput>;
  recipient?: Types.InputMaybe<Types.CreateAddressInput>;
  inputPallets: Array<Types.CreateOrUpdatePallet> | Types.CreateOrUpdatePallet;
}>;


export type UpdateConsolidationMutation = { updateConsolidation: { id: string } };

export type ConsolidationSearchCountriesQueryVariables = Types.Exact<{
  term: Types.Scalars['String'];
}>;


export type ConsolidationSearchCountriesQuery = { countries: { edges?: Array<{ node?: { id: string, label: string, alpha2: string } | null } | null> | null } };

export type CreateConsolidationShipmentQueryVariables = Types.Exact<{
  consolidationID: Types.Scalars['ID'];
  prebook: Types.Scalars['Boolean'];
}>;


export type CreateConsolidationShipmentQuery = { createConsolidationShipment: { allLabels: string, labelsPDF: Array<string> } };

export const ConsolidationOrderFragmentDoc = gql`
    fragment ConsolidationOrder on Order {
  id
  orderPublicID
  status
}
    `;
export const FetchConsolidationDocument = gql`
    query FetchConsolidation($id: ID!) {
  consolidation(id: $id) {
    publicID
    description
    status
    sender {
      ...AddressInfo
    }
    recipient {
      ...AddressInfo
    }
    deliveryOption {
      id
    }
    orders {
      ...ConsolidationOrder
    }
    pallets {
      id
      publicID
      description
      packaging {
        id
      }
      orders {
        ...ConsolidationOrder
      }
    }
  }
  consolidationShipments(consolidationID: $id) {
    mayBook
    mayPrebook
    shipment {
      id
      shipmentPallet {
        id
        pallet {
          id
        }
      }
    }
  }
  locationInfo: consolidation(id: $id) {
    sender {
      ...AddressInfo
    }
    recipient {
      ...AddressInfo
    }
  }
  deliveryOptions(where: {hasCarrierServiceWith: {consolidation: true}}) {
    edges {
      node {
        id
        name
        carrierService {
          return
        }
      }
    }
  }
  packagings(where: {hasCarrierBrandWith: {internalID: df}}) {
    edges {
      node {
        id
        name
      }
    }
  }
}
    ${AddressInfoFragmentDoc}
${ConsolidationOrderFragmentDoc}`;

  @Injectable({
    providedIn: 'root'
  })
  export class FetchConsolidationGQL extends Apollo.Query<FetchConsolidationQuery, FetchConsolidationQueryVariables> {
    document = FetchConsolidationDocument;
    
    constructor(apollo: Apollo.Apollo) {
      super(apollo);
    }
  }
export const ConsolidationSearchOrdersDocument = gql`
    query ConsolidationSearchOrders($term: String!) {
  orders(where: {orderPublicIDContainsFold: $term}) {
    edges {
      node {
        ...ConsolidationOrder
      }
    }
  }
}
    ${ConsolidationOrderFragmentDoc}`;

  @Injectable({
    providedIn: 'root'
  })
  export class ConsolidationSearchOrdersGQL extends Apollo.Query<ConsolidationSearchOrdersQuery, ConsolidationSearchOrdersQueryVariables> {
    document = ConsolidationSearchOrdersDocument;
    
    constructor(apollo: Apollo.Apollo) {
      super(apollo);
    }
  }
export const UpdateConsolidationDocument = gql`
    mutation UpdateConsolidation($id: ID!, $input: UpdateConsolidationInput!, $sender: CreateAddressInput, $recipient: CreateAddressInput, $inputPallets: [CreateOrUpdatePallet!]!) {
  updateConsolidation(
    id: $id
    input: $input
    sender: $sender
    recipient: $recipient
    inputPallets: $inputPallets
  ) {
    id
  }
}
    `;

  @Injectable({
    providedIn: 'root'
  })
  export class UpdateConsolidationGQL extends Apollo.Mutation<UpdateConsolidationMutation, UpdateConsolidationMutationVariables> {
    document = UpdateConsolidationDocument;
    
    constructor(apollo: Apollo.Apollo) {
      super(apollo);
    }
  }
export const ConsolidationSearchCountriesDocument = gql`
    query ConsolidationSearchCountries($term: String!) {
  countries(
    where: {or: [{labelContainsFold: $term}, {alpha2ContainsFold: $term}]}
  ) {
    edges {
      node {
        id
        label
        alpha2
      }
    }
  }
}
    `;

  @Injectable({
    providedIn: 'root'
  })
  export class ConsolidationSearchCountriesGQL extends Apollo.Query<ConsolidationSearchCountriesQuery, ConsolidationSearchCountriesQueryVariables> {
    document = ConsolidationSearchCountriesDocument;
    
    constructor(apollo: Apollo.Apollo) {
      super(apollo);
    }
  }
export const CreateConsolidationShipmentDocument = gql`
    query CreateConsolidationShipment($consolidationID: ID!, $prebook: Boolean!) {
  createConsolidationShipment(
    consolidationID: $consolidationID
    prebook: $prebook
  ) {
    allLabels
    labelsPDF
  }
}
    `;

  @Injectable({
    providedIn: 'root'
  })
  export class CreateConsolidationShipmentGQL extends Apollo.Query<CreateConsolidationShipmentQuery, CreateConsolidationShipmentQueryVariables> {
    document = CreateConsolidationShipmentDocument;
    
    constructor(apollo: Apollo.Apollo) {
      super(apollo);
    }
  }