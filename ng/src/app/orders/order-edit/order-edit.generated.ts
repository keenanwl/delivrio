/* eslint-disable */
import * as Types from '../../../generated/graphql';

import { gql } from 'apollo-angular';
import { Injectable } from '@angular/core';
import * as Apollo from 'apollo-angular';
export type FetchColliQueryVariables = Types.Exact<{
  id: Types.Scalars['ID'];
}>;


export type FetchColliQuery = { colli?: { order: { connection: { id: string } }, deliveryOption?: { id: string, clickCollect?: boolean | null } | null, recipient: { id: string, firstName: string, lastName: string, phoneNumber: string, email: string, addressOne: string, addressTwo: string, zip: string, city: string, state?: string | null, company?: string | null, country: { id: string, alpha2: string, label: string } }, sender: { id: string, firstName: string, lastName: string, phoneNumber: string, email: string, addressOne: string, addressTwo: string, zip: string, city: string, state?: string | null, company?: string | null, country: { id: string, alpha2: string, label: string } }, packaging?: { id: string, name: string, lengthCm: number, widthCm: number, heightCm: number } | null, orderLines?: Array<{ id: string, units: number, unitPrice: number, discountAllocationAmount: number, currency: { display: string }, productVariant: { id: string, dimensionLength?: number | null, dimensionWidth?: number | null, dimensionHeight?: number | null, weightG?: number | null, description?: string | null, archived: boolean, productImage?: Array<{ url: string }> | null, product: { id: string, externalID?: string | null, title: string, status: Types.ProductStatus } } }> | null } | null, connections: { edges?: Array<{ node?: { id: string, name: string } | null } | null> | null }, deliveryPoint?: { id: string, name: string, address: { addressOne: string, addressTwo?: string | null, zip: string, city: string, state?: string | null, country: { alpha2: string } } } | null, clickCollectLocation?: { id: string, name: string, address: { id: string, firstName: string, lastName: string, phoneNumber: string, email: string, addressOne: string, addressTwo: string, zip: string, city: string, state?: string | null, company?: string | null, country: { id: string, alpha2: string, label: string } } } | null };

export type DeliveryPointListItemFragment = { id: string, name: string, address: { addressOne: string, addressTwo?: string | null, zip: string, city: string, state?: string | null, country: { alpha2: string } } };

export type OrdersSearchCountriesQueryVariables = Types.Exact<{
  term: Types.Scalars['String'];
}>;


export type OrdersSearchCountriesQuery = { countries: { edges?: Array<{ node?: { id: string, label: string, alpha2: string, region: Types.CountryRegion } | null } | null> | null } };

export type OrderLinesFragment = { id: string, units: number, unitPrice: number, discountAllocationAmount: number, currency: { display: string }, productVariant: { id: string, dimensionLength?: number | null, dimensionWidth?: number | null, dimensionHeight?: number | null, weightG?: number | null, description?: string | null, archived: boolean, productImage?: Array<{ url: string }> | null, product: { id: string, externalID?: string | null, title: string, status: Types.ProductStatus } } };

export type AddressInfoFragment = { id: string, firstName: string, lastName: string, phoneNumber: string, email: string, addressOne: string, addressTwo: string, zip: string, city: string, state?: string | null, company?: string | null, country: { id: string, alpha2: string, label: string } };

export type FetchDeliveryOptionsQueryVariables = Types.Exact<{
  orderInfo: Types.DeliveryOptionSeedInput;
}>;


export type FetchDeliveryOptionsQuery = { deliveryOptionsList: Array<{ deliveryOptionID: string, name: string, status: Types.DeliveryOptionBrandNameStatus, price?: string | null, warning?: string | null, requiresDeliveryPoint: boolean, deliveryPoint: boolean, clickAndCollect: boolean, currency?: { id: string, display: string } | null } | null> };

export type FetchAvailableDeliveryPointsQueryVariables = Types.Exact<{
  deliveryOptionID?: Types.InputMaybe<Types.Scalars['ID']>;
  lookupAddress?: Types.InputMaybe<Types.CreateAddressInput>;
}>;


export type FetchAvailableDeliveryPointsQuery = { availableDeliveryPoints: Array<{ id: string, name: string, address: { addressOne: string, addressTwo?: string | null, zip: string, city: string, state?: string | null, country: { alpha2: string } } }> };

export type FetchAvailableClickCollectLocationsQueryVariables = Types.Exact<{
  deliveryOptionID: Types.Scalars['ID'];
}>;


export type FetchAvailableClickCollectLocationsQuery = { locations: { edges?: Array<{ node?: { id: string, name: string, address: { id: string, firstName: string, lastName: string, phoneNumber: string, email: string, addressOne: string, addressTwo: string, zip: string, city: string, state?: string | null, company?: string | null, country: { id: string, alpha2: string, label: string } } } | null } | null> | null } };

export type SearchProductsQueryVariables = Types.Exact<{
  term: Types.Scalars['String'];
}>;


export type SearchProductsQuery = { productVariants: { edges?: Array<{ node?: { id: string, dimensionLength?: number | null, dimensionWidth?: number | null, dimensionHeight?: number | null, weightG?: number | null, description?: string | null, archived: boolean, productImage?: Array<{ url: string }> | null, product: { id: string, externalID?: string | null, title: string, status: Types.ProductStatus, productImage?: Array<{ url: string }> | null } } | null } | null> | null } };

export type AddColliToOrderMutationVariables = Types.Exact<{
  orderID: Types.Scalars['ID'];
  input: Types.CreateColliInput;
  deliveryOptionID?: Types.InputMaybe<Types.Scalars['ID']>;
  deliveryPointID?: Types.InputMaybe<Types.Scalars['ID']>;
  ccLocationID?: Types.InputMaybe<Types.Scalars['ID']>;
  packagingID?: Types.InputMaybe<Types.Scalars['ID']>;
  recipientAddress: Types.CreateAddressInput;
  senderAddress: Types.CreateAddressInput;
  products: Array<Types.ProductVariantQuantity> | Types.ProductVariantQuantity;
}>;


export type AddColliToOrderMutation = { createColli: { id: string } };

export type UpdateColliMutationVariables = Types.Exact<{
  id: Types.Scalars['ID'];
  input: Types.UpdateColliInput;
  deliveryOptionID?: Types.InputMaybe<Types.Scalars['ID']>;
  deliveryPointID?: Types.InputMaybe<Types.Scalars['ID']>;
  ccLocationID?: Types.InputMaybe<Types.Scalars['ID']>;
  packagingID?: Types.InputMaybe<Types.Scalars['ID']>;
  recipientAddressID: Types.Scalars['ID'];
  recipientAddress: Types.UpdateAddressInput;
  senderAddressID: Types.Scalars['ID'];
  senderAddress: Types.UpdateAddressInput;
  updateExistingRecipient?: Types.InputMaybe<Types.Scalars['Boolean']>;
  products: Array<Types.ProductVariantQuantity> | Types.ProductVariantQuantity;
  removeProducts: Array<Types.Scalars['ID']> | Types.Scalars['ID'];
}>;


export type UpdateColliMutation = { updateColli?: { id: string } | null };

export type FetchPackagingQueryVariables = Types.Exact<{ [key: string]: never; }>;


export type FetchPackagingQuery = { packagings: { edges?: Array<{ node?: { id: string, name: string, lengthCm: number, heightCm: number, widthCm: number } | null } | null> | null } };

export const DeliveryPointListItemFragmentDoc = gql`
    fragment DeliveryPointListItem on DeliveryPoint {
  id
  name
  address {
    addressOne
    addressTwo
    zip
    city
    state
    country {
      alpha2
    }
  }
}
    `;
export const OrderLinesFragmentDoc = gql`
    fragment OrderLines on OrderLine {
  id
  units
  unitPrice
  discountAllocationAmount
  currency {
    display
  }
  productVariant {
    id
    dimensionLength
    dimensionWidth
    dimensionHeight
    weightG
    description
    archived
    productImage {
      url
    }
    product {
      id
      externalID
      title
      status
    }
  }
}
    `;
export const AddressInfoFragmentDoc = gql`
    fragment AddressInfo on Address {
  id
  firstName
  lastName
  phoneNumber
  email
  addressOne
  addressTwo
  zip
  city
  state
  country {
    id
    alpha2
    label
  }
  company
}
    `;
export const FetchColliDocument = gql`
    query FetchColli($id: ID!) {
  colli(id: $id) {
    order {
      connection {
        id
      }
    }
    deliveryOption {
      id
      clickCollect
    }
    recipient {
      ...AddressInfo
    }
    sender {
      ...AddressInfo
    }
    packaging {
      id
      name
      lengthCm
      widthCm
      heightCm
    }
    orderLines {
      ...OrderLines
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
  deliveryPoint(colliID: $id) {
    ...DeliveryPointListItem
  }
  clickCollectLocation(colliID: $id) {
    id
    name
    address {
      ...AddressInfo
    }
  }
}
    ${AddressInfoFragmentDoc}
${OrderLinesFragmentDoc}
${DeliveryPointListItemFragmentDoc}`;

  @Injectable({
    providedIn: 'root'
  })
  export class FetchColliGQL extends Apollo.Query<FetchColliQuery, FetchColliQueryVariables> {
    document = FetchColliDocument;
    
    constructor(apollo: Apollo.Apollo) {
      super(apollo);
    }
  }
export const OrdersSearchCountriesDocument = gql`
    query OrdersSearchCountries($term: String!) {
  countries(
    where: {or: [{labelContainsFold: $term}, {alpha2ContainsFold: $term}]}
  ) {
    edges {
      node {
        id
        label
        alpha2
        region
      }
    }
  }
}
    `;

  @Injectable({
    providedIn: 'root'
  })
  export class OrdersSearchCountriesGQL extends Apollo.Query<OrdersSearchCountriesQuery, OrdersSearchCountriesQueryVariables> {
    document = OrdersSearchCountriesDocument;
    
    constructor(apollo: Apollo.Apollo) {
      super(apollo);
    }
  }
export const FetchDeliveryOptionsDocument = gql`
    query FetchDeliveryOptions($orderInfo: DeliveryOptionSeedInput!) {
  deliveryOptionsList(orderInfo: $orderInfo) {
    deliveryOptionID
    name
    status
    price
    warning
    requiresDeliveryPoint
    deliveryPoint
    clickAndCollect
    currency {
      id
      display
    }
  }
}
    `;

  @Injectable({
    providedIn: 'root'
  })
  export class FetchDeliveryOptionsGQL extends Apollo.Query<FetchDeliveryOptionsQuery, FetchDeliveryOptionsQueryVariables> {
    document = FetchDeliveryOptionsDocument;
    
    constructor(apollo: Apollo.Apollo) {
      super(apollo);
    }
  }
export const FetchAvailableDeliveryPointsDocument = gql`
    query FetchAvailableDeliveryPoints($deliveryOptionID: ID, $lookupAddress: CreateAddressInput) {
  availableDeliveryPoints(
    deliveryOptionID: $deliveryOptionID
    address: $lookupAddress
  ) {
    ...DeliveryPointListItem
  }
}
    ${DeliveryPointListItemFragmentDoc}`;

  @Injectable({
    providedIn: 'root'
  })
  export class FetchAvailableDeliveryPointsGQL extends Apollo.Query<FetchAvailableDeliveryPointsQuery, FetchAvailableDeliveryPointsQueryVariables> {
    document = FetchAvailableDeliveryPointsDocument;
    
    constructor(apollo: Apollo.Apollo) {
      super(apollo);
    }
  }
export const FetchAvailableClickCollectLocationsDocument = gql`
    query FetchAvailableClickCollectLocations($deliveryOptionID: ID!) {
  locations(where: {hasDeliveryOptionWith: {id: $deliveryOptionID}}) {
    edges {
      node {
        id
        name
        address {
          ...AddressInfo
        }
      }
    }
  }
}
    ${AddressInfoFragmentDoc}`;

  @Injectable({
    providedIn: 'root'
  })
  export class FetchAvailableClickCollectLocationsGQL extends Apollo.Query<FetchAvailableClickCollectLocationsQuery, FetchAvailableClickCollectLocationsQueryVariables> {
    document = FetchAvailableClickCollectLocationsDocument;
    
    constructor(apollo: Apollo.Apollo) {
      super(apollo);
    }
  }
export const SearchProductsDocument = gql`
    query SearchProducts($term: String!) {
  productVariants(
    where: {or: [{descriptionContainsFold: $term}, {hasProductWith: {or: [{titleContainsFold: $term}, {bodyHTMLContainsFold: $term}]}}]}
  ) {
    edges {
      node {
        id
        dimensionLength
        dimensionWidth
        dimensionHeight
        weightG
        description
        archived
        productImage {
          url
        }
        product {
          id
          externalID
          title
          status
          productImage {
            url
          }
        }
      }
    }
  }
}
    `;

  @Injectable({
    providedIn: 'root'
  })
  export class SearchProductsGQL extends Apollo.Query<SearchProductsQuery, SearchProductsQueryVariables> {
    document = SearchProductsDocument;
    
    constructor(apollo: Apollo.Apollo) {
      super(apollo);
    }
  }
export const AddColliToOrderDocument = gql`
    mutation AddColliToOrder($orderID: ID!, $input: CreateColliInput!, $deliveryOptionID: ID, $deliveryPointID: ID, $ccLocationID: ID, $packagingID: ID, $recipientAddress: CreateAddressInput!, $senderAddress: CreateAddressInput!, $products: [ProductVariantQuantity!]!) {
  createColli(
    orderID: $orderID
    input: $input
    deliveryOptionID: $deliveryOptionID
    deliveryPointID: $deliveryPointID
    ccLocationID: $ccLocationID
    packagingID: $packagingID
    recipientAddress: $recipientAddress
    senderAddress: $senderAddress
    products: $products
  ) {
    id
  }
}
    `;

  @Injectable({
    providedIn: 'root'
  })
  export class AddColliToOrderGQL extends Apollo.Mutation<AddColliToOrderMutation, AddColliToOrderMutationVariables> {
    document = AddColliToOrderDocument;
    
    constructor(apollo: Apollo.Apollo) {
      super(apollo);
    }
  }
export const UpdateColliDocument = gql`
    mutation UpdateColli($id: ID!, $input: UpdateColliInput!, $deliveryOptionID: ID, $deliveryPointID: ID, $ccLocationID: ID, $packagingID: ID, $recipientAddressID: ID!, $recipientAddress: UpdateAddressInput!, $senderAddressID: ID!, $senderAddress: UpdateAddressInput!, $updateExistingRecipient: Boolean, $products: [ProductVariantQuantity!]!, $removeProducts: [ID!]!) {
  updateColli(
    id: $id
    input: $input
    deliveryOptionID: $deliveryOptionID
    deliveryPointID: $deliveryPointID
    ccLocationID: $ccLocationID
    packagingID: $packagingID
    recipientAddressID: $recipientAddressID
    recipientAddress: $recipientAddress
    senderAddressID: $senderAddressID
    senderAddress: $senderAddress
    updateExistingRecipient: $updateExistingRecipient
    products: $products
    removeProducts: $removeProducts
  ) {
    id
  }
}
    `;

  @Injectable({
    providedIn: 'root'
  })
  export class UpdateColliGQL extends Apollo.Mutation<UpdateColliMutation, UpdateColliMutationVariables> {
    document = UpdateColliDocument;
    
    constructor(apollo: Apollo.Apollo) {
      super(apollo);
    }
  }
export const FetchPackagingDocument = gql`
    query FetchPackaging {
  packagings {
    edges {
      node {
        id
        name
        lengthCm
        heightCm
        widthCm
      }
    }
  }
}
    `;

  @Injectable({
    providedIn: 'root'
  })
  export class FetchPackagingGQL extends Apollo.Query<FetchPackagingQuery, FetchPackagingQueryVariables> {
    document = FetchPackagingDocument;
    
    constructor(apollo: Apollo.Apollo) {
      super(apollo);
    }
  }