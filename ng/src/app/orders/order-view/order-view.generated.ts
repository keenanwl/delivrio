/* eslint-disable */
import * as Types from '../../../generated/graphql';

import { gql } from 'apollo-angular';
import { TimelineViewerFragmentDoc } from '../../shared/timeline-viewer/timeline-viewer.generated';
import { AddressInfoFragmentDoc, OrderLinesFragmentDoc } from '../order-edit/order-edit.generated';
import { Injectable } from '@angular/core';
import * as Apollo from 'apollo-angular';
export type FetchOrderViewQueryVariables = Types.Exact<{
  id: Types.Scalars['ID'];
}>;


export type FetchOrderViewQuery = { connections: { edges?: Array<{ node?: { id: string, name: string } | null } | null> | null }, order?: { orderPublicID: string, commentExternal?: string | null, commentInternal?: string | null, status: Types.OrderStatus, connection: { id: string, name: string }, colli?: Array<{ id: string, status: Types.ColliStatus, recipient: { id: string, firstName: string, lastName: string, phoneNumber: string, email: string, addressOne: string, addressTwo: string, zip: string, city: string, state?: string | null, company?: string | null, country: { id: string, alpha2: string, label: string } }, sender: { id: string, firstName: string, lastName: string, phoneNumber: string, email: string, addressOne: string, addressTwo: string, zip: string, city: string, state?: string | null, company?: string | null, country: { id: string, alpha2: string, label: string } }, orderLines?: Array<{ id: string, units: number, unitPrice: number, discountAllocationAmount: number, currency: { display: string }, productVariant: { id: string, dimensionLength?: number | null, dimensionWidth?: number | null, dimensionHeight?: number | null, weightG?: number | null, description?: string | null, archived: boolean, productImage?: Array<{ url: string }> | null, product: { id: string, externalID?: string | null, title: string, status: Types.ProductStatus } } }> | null, deliveryOption?: { id: string, clickCollect?: boolean | null, name: string, carrier: { carrierBrand: { label: string } } } | null, packaging?: { name: string, lengthCm: number, widthCm: number, heightCm: number } | null, clickCollectLocation?: { name: string, address: { addressOne: string, addressTwo: string, zip: string, city: string, state?: string | null, country: { alpha2: string } } } | null, parcelShop?: { name: string, address: { addressOne: string, addressTwo?: string | null, zip: string, city: string, state?: string | null, country: { alpha2: string } } } | null }> | null } | null, orderShipments?: { mayShipRemaining: boolean, shipmentStatuses: Array<{ shipmentID?: string | null, colliID: string, ccSignatures: Array<string> }> } | null, orderTimeline: Array<{ id: string, createdAt: any, user?: { id: string, name?: string | null } | null, orderHistory?: Array<{ id: string, type: Types.OrderHistoryType, description: string, order: { id: string, orderPublicID: string } }> | null, shipmentHistory?: Array<{ id: string, type: Types.ShipmentHistoryType }> | null, returnColliHistory?: Array<{ id: string, description: string, type: Types.ReturnColliHistoryType }> | null } | null> };

export type CancelShipmentByColliIDsMutationVariables = Types.Exact<{
  colliIDs: Array<Types.Scalars['ID']> | Types.Scalars['ID'];
}>;


export type CancelShipmentByColliIDsMutation = { cancelShipmentByColliIDs: boolean };

export type FetchLabelsQueryVariables = Types.Exact<{
  colliIDs: Array<Types.Scalars['ID']> | Types.Scalars['ID'];
}>;


export type FetchLabelsQuery = { shipmentLabels: { allLabels: string, labelsPDF: Array<string> } };

export type FetchPackingSlipsQueryVariables = Types.Exact<{
  colliIDs: Array<Types.Scalars['ID']> | Types.Scalars['ID'];
}>;


export type FetchPackingSlipsQuery = { packingSlips: { allPackingSlips: string, packingSlips: Array<string> } };

export type OrderViewFragment = { orderPublicID: string, commentExternal?: string | null, commentInternal?: string | null, status: Types.OrderStatus, connection: { id: string, name: string }, colli?: Array<{ id: string, status: Types.ColliStatus, recipient: { id: string, firstName: string, lastName: string, phoneNumber: string, email: string, addressOne: string, addressTwo: string, zip: string, city: string, state?: string | null, company?: string | null, country: { id: string, alpha2: string, label: string } }, sender: { id: string, firstName: string, lastName: string, phoneNumber: string, email: string, addressOne: string, addressTwo: string, zip: string, city: string, state?: string | null, company?: string | null, country: { id: string, alpha2: string, label: string } }, orderLines?: Array<{ id: string, units: number, unitPrice: number, discountAllocationAmount: number, currency: { display: string }, productVariant: { id: string, dimensionLength?: number | null, dimensionWidth?: number | null, dimensionHeight?: number | null, weightG?: number | null, description?: string | null, archived: boolean, productImage?: Array<{ url: string }> | null, product: { id: string, externalID?: string | null, title: string, status: Types.ProductStatus } } }> | null, deliveryOption?: { id: string, clickCollect?: boolean | null, name: string, carrier: { carrierBrand: { label: string } } } | null, packaging?: { name: string, lengthCm: number, widthCm: number, heightCm: number } | null, clickCollectLocation?: { name: string, address: { addressOne: string, addressTwo: string, zip: string, city: string, state?: string | null, country: { alpha2: string } } } | null, parcelShop?: { name: string, address: { addressOne: string, addressTwo?: string | null, zip: string, city: string, state?: string | null, country: { alpha2: string } } } | null }> | null };

export type ColliViewInfoFragment = { id: string, status: Types.ColliStatus, recipient: { id: string, firstName: string, lastName: string, phoneNumber: string, email: string, addressOne: string, addressTwo: string, zip: string, city: string, state?: string | null, company?: string | null, country: { id: string, alpha2: string, label: string } }, sender: { id: string, firstName: string, lastName: string, phoneNumber: string, email: string, addressOne: string, addressTwo: string, zip: string, city: string, state?: string | null, company?: string | null, country: { id: string, alpha2: string, label: string } }, orderLines?: Array<{ id: string, units: number, unitPrice: number, discountAllocationAmount: number, currency: { display: string }, productVariant: { id: string, dimensionLength?: number | null, dimensionWidth?: number | null, dimensionHeight?: number | null, weightG?: number | null, description?: string | null, archived: boolean, productImage?: Array<{ url: string }> | null, product: { id: string, externalID?: string | null, title: string, status: Types.ProductStatus } } }> | null, deliveryOption?: { id: string, clickCollect?: boolean | null, name: string, carrier: { carrierBrand: { label: string } } } | null, packaging?: { name: string, lengthCm: number, widthCm: number, heightCm: number } | null, clickCollectLocation?: { name: string, address: { addressOne: string, addressTwo: string, zip: string, city: string, state?: string | null, country: { alpha2: string } } } | null, parcelShop?: { name: string, address: { addressOne: string, addressTwo?: string | null, zip: string, city: string, state?: string | null, country: { alpha2: string } } } | null };

export type UpdateOrderMutationVariables = Types.Exact<{
  id: Types.Scalars['ID'];
  input: Types.UpdateOrderInput;
}>;


export type UpdateOrderMutation = { updateOrder?: { orderPublicID: string, commentExternal?: string | null, commentInternal?: string | null, status: Types.OrderStatus, connection: { id: string, name: string }, colli?: Array<{ id: string, status: Types.ColliStatus, recipient: { id: string, firstName: string, lastName: string, phoneNumber: string, email: string, addressOne: string, addressTwo: string, zip: string, city: string, state?: string | null, company?: string | null, country: { id: string, alpha2: string, label: string } }, sender: { id: string, firstName: string, lastName: string, phoneNumber: string, email: string, addressOne: string, addressTwo: string, zip: string, city: string, state?: string | null, company?: string | null, country: { id: string, alpha2: string, label: string } }, orderLines?: Array<{ id: string, units: number, unitPrice: number, discountAllocationAmount: number, currency: { display: string }, productVariant: { id: string, dimensionLength?: number | null, dimensionWidth?: number | null, dimensionHeight?: number | null, weightG?: number | null, description?: string | null, archived: boolean, productImage?: Array<{ url: string }> | null, product: { id: string, externalID?: string | null, title: string, status: Types.ProductStatus } } }> | null, deliveryOption?: { id: string, clickCollect?: boolean | null, name: string, carrier: { carrierBrand: { label: string } } } | null, packaging?: { name: string, lengthCm: number, widthCm: number, heightCm: number } | null, clickCollectLocation?: { name: string, address: { addressOne: string, addressTwo: string, zip: string, city: string, state?: string | null, country: { alpha2: string } } } | null, parcelShop?: { name: string, address: { addressOne: string, addressTwo?: string | null, zip: string, city: string, state?: string | null, country: { alpha2: string } } } | null }> | null } | null };

export type MoveOrderLineMutationVariables = Types.Exact<{
  orderLineID: Types.Scalars['ID'];
  colliID: Types.Scalars['ID'];
}>;


export type MoveOrderLineMutation = { moveOrderLine?: Array<{ id: string, status: Types.ColliStatus, recipient: { id: string, firstName: string, lastName: string, phoneNumber: string, email: string, addressOne: string, addressTwo: string, zip: string, city: string, state?: string | null, company?: string | null, country: { id: string, alpha2: string, label: string } }, sender: { id: string, firstName: string, lastName: string, phoneNumber: string, email: string, addressOne: string, addressTwo: string, zip: string, city: string, state?: string | null, company?: string | null, country: { id: string, alpha2: string, label: string } }, orderLines?: Array<{ id: string, units: number, unitPrice: number, discountAllocationAmount: number, currency: { display: string }, productVariant: { id: string, dimensionLength?: number | null, dimensionWidth?: number | null, dimensionHeight?: number | null, weightG?: number | null, description?: string | null, archived: boolean, productImage?: Array<{ url: string }> | null, product: { id: string, externalID?: string | null, title: string, status: Types.ProductStatus } } }> | null, deliveryOption?: { id: string, clickCollect?: boolean | null, name: string, carrier: { carrierBrand: { label: string } } } | null, packaging?: { name: string, lengthCm: number, widthCm: number, heightCm: number } | null, clickCollectLocation?: { name: string, address: { addressOne: string, addressTwo: string, zip: string, city: string, state?: string | null, country: { alpha2: string } } } | null, parcelShop?: { name: string, address: { addressOne: string, addressTwo?: string | null, zip: string, city: string, state?: string | null, country: { alpha2: string } } } | null }> | null };

export type DeletePackageMutationVariables = Types.Exact<{
  colliID: Types.Scalars['ID'];
}>;


export type DeletePackageMutation = { deleteColli: { orderPublicID: string, commentExternal?: string | null, commentInternal?: string | null, status: Types.OrderStatus, connection: { id: string, name: string }, colli?: Array<{ id: string, status: Types.ColliStatus, recipient: { id: string, firstName: string, lastName: string, phoneNumber: string, email: string, addressOne: string, addressTwo: string, zip: string, city: string, state?: string | null, company?: string | null, country: { id: string, alpha2: string, label: string } }, sender: { id: string, firstName: string, lastName: string, phoneNumber: string, email: string, addressOne: string, addressTwo: string, zip: string, city: string, state?: string | null, company?: string | null, country: { id: string, alpha2: string, label: string } }, orderLines?: Array<{ id: string, units: number, unitPrice: number, discountAllocationAmount: number, currency: { display: string }, productVariant: { id: string, dimensionLength?: number | null, dimensionWidth?: number | null, dimensionHeight?: number | null, weightG?: number | null, description?: string | null, archived: boolean, productImage?: Array<{ url: string }> | null, product: { id: string, externalID?: string | null, title: string, status: Types.ProductStatus } } }> | null, deliveryOption?: { id: string, clickCollect?: boolean | null, name: string, carrier: { carrierBrand: { label: string } } } | null, packaging?: { name: string, lengthCm: number, widthCm: number, heightCm: number } | null, clickCollectLocation?: { name: string, address: { addressOne: string, addressTwo: string, zip: string, city: string, state?: string | null, country: { alpha2: string } } } | null, parcelShop?: { name: string, address: { addressOne: string, addressTwo?: string | null, zip: string, city: string, state?: string | null, country: { alpha2: string } } } | null }> | null } };

export type DuplicateColliMutationVariables = Types.Exact<{
  fromColliID: Types.Scalars['ID'];
}>;


export type DuplicateColliMutation = { duplicateColli: { orderPublicID: string, commentExternal?: string | null, commentInternal?: string | null, status: Types.OrderStatus, connection: { id: string, name: string }, colli?: Array<{ id: string, status: Types.ColliStatus, recipient: { id: string, firstName: string, lastName: string, phoneNumber: string, email: string, addressOne: string, addressTwo: string, zip: string, city: string, state?: string | null, company?: string | null, country: { id: string, alpha2: string, label: string } }, sender: { id: string, firstName: string, lastName: string, phoneNumber: string, email: string, addressOne: string, addressTwo: string, zip: string, city: string, state?: string | null, company?: string | null, country: { id: string, alpha2: string, label: string } }, orderLines?: Array<{ id: string, units: number, unitPrice: number, discountAllocationAmount: number, currency: { display: string }, productVariant: { id: string, dimensionLength?: number | null, dimensionWidth?: number | null, dimensionHeight?: number | null, weightG?: number | null, description?: string | null, archived: boolean, productImage?: Array<{ url: string }> | null, product: { id: string, externalID?: string | null, title: string, status: Types.ProductStatus } } }> | null, deliveryOption?: { id: string, clickCollect?: boolean | null, name: string, carrier: { carrierBrand: { label: string } } } | null, packaging?: { name: string, lengthCm: number, widthCm: number, heightCm: number } | null, clickCollectLocation?: { name: string, address: { addressOne: string, addressTwo: string, zip: string, city: string, state?: string | null, country: { alpha2: string } } } | null, parcelShop?: { name: string, address: { addressOne: string, addressTwo?: string | null, zip: string, city: string, state?: string | null, country: { alpha2: string } } } | null }> | null } };

export type CreateShipmentsQueryVariables = Types.Exact<{
  orderID: Types.Scalars['ID'];
  parcelIDs: Array<Types.Scalars['ID']> | Types.Scalars['ID'];
}>;


export type CreateShipmentsQuery = { createShipments: { labelsPDF: Array<string>, allLabels: string } };

export type CreatePackingSlipPrintJobsQueryVariables = Types.Exact<{
  colliIDs: Array<Types.Scalars['ID']> | Types.Scalars['ID'];
}>;


export type CreatePackingSlipPrintJobsQuery = { createPackingListPrintJob: boolean };

export type CreateLabelPrintJobsQueryVariables = Types.Exact<{
  colliIDs: Array<Types.Scalars['ID']> | Types.Scalars['ID'];
}>;


export type CreateLabelPrintJobsQuery = { createLabelsPrintJob: boolean };

export type PackingSlipsClearCacheQueryVariables = Types.Exact<{
  orderIDs: Array<Types.Scalars['ID']> | Types.Scalars['ID'];
}>;


export type PackingSlipsClearCacheQuery = { packingSlipsClearCache: boolean };

export const ColliViewInfoFragmentDoc = gql`
    fragment ColliViewInfo on Colli {
  id
  recipient {
    ...AddressInfo
  }
  sender {
    ...AddressInfo
  }
  status
  orderLines {
    ...OrderLines
  }
  deliveryOption {
    id
    clickCollect
    name
    carrier {
      carrierBrand {
        label
      }
    }
  }
  packaging {
    name
    lengthCm
    widthCm
    heightCm
  }
  clickCollectLocation {
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
  parcelShop {
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
}
    ${AddressInfoFragmentDoc}
${OrderLinesFragmentDoc}`;
export const OrderViewFragmentDoc = gql`
    fragment OrderView on Order {
  orderPublicID
  commentExternal
  commentInternal
  connection {
    id
    name
  }
  status
  colli {
    ...ColliViewInfo
  }
}
    ${ColliViewInfoFragmentDoc}`;
export const FetchOrderViewDocument = gql`
    query FetchOrderView($id: ID!) {
  connections {
    edges {
      node {
        id
        name
      }
    }
  }
  order(id: $id) {
    ...OrderView
  }
  orderShipments(orderID: $id) {
    mayShipRemaining
    shipmentStatuses {
      shipmentID
      colliID
      ccSignatures
    }
  }
  orderTimeline(orderID: $id) {
    ...TimelineViewer
  }
}
    ${OrderViewFragmentDoc}
${TimelineViewerFragmentDoc}`;

  @Injectable({
    providedIn: 'root'
  })
  export class FetchOrderViewGQL extends Apollo.Query<FetchOrderViewQuery, FetchOrderViewQueryVariables> {
    document = FetchOrderViewDocument;
    
    constructor(apollo: Apollo.Apollo) {
      super(apollo);
    }
  }
export const CancelShipmentByColliIDsDocument = gql`
    mutation CancelShipmentByColliIDs($colliIDs: [ID!]!) {
  cancelShipmentByColliIDs(colliIDs: $colliIDs)
}
    `;

  @Injectable({
    providedIn: 'root'
  })
  export class CancelShipmentByColliIDsGQL extends Apollo.Mutation<CancelShipmentByColliIDsMutation, CancelShipmentByColliIDsMutationVariables> {
    document = CancelShipmentByColliIDsDocument;
    
    constructor(apollo: Apollo.Apollo) {
      super(apollo);
    }
  }
export const FetchLabelsDocument = gql`
    query FetchLabels($colliIDs: [ID!]!) {
  shipmentLabels(colliIDs: $colliIDs) {
    allLabels
    labelsPDF
  }
}
    `;

  @Injectable({
    providedIn: 'root'
  })
  export class FetchLabelsGQL extends Apollo.Query<FetchLabelsQuery, FetchLabelsQueryVariables> {
    document = FetchLabelsDocument;
    
    constructor(apollo: Apollo.Apollo) {
      super(apollo);
    }
  }
export const FetchPackingSlipsDocument = gql`
    query FetchPackingSlips($colliIDs: [ID!]!) {
  packingSlips(colliIDs: $colliIDs) {
    allPackingSlips
    packingSlips
  }
}
    `;

  @Injectable({
    providedIn: 'root'
  })
  export class FetchPackingSlipsGQL extends Apollo.Query<FetchPackingSlipsQuery, FetchPackingSlipsQueryVariables> {
    document = FetchPackingSlipsDocument;
    
    constructor(apollo: Apollo.Apollo) {
      super(apollo);
    }
  }
export const UpdateOrderDocument = gql`
    mutation UpdateOrder($id: ID!, $input: UpdateOrderInput!) {
  updateOrder(id: $id, input: $input) {
    ...OrderView
  }
}
    ${OrderViewFragmentDoc}`;

  @Injectable({
    providedIn: 'root'
  })
  export class UpdateOrderGQL extends Apollo.Mutation<UpdateOrderMutation, UpdateOrderMutationVariables> {
    document = UpdateOrderDocument;
    
    constructor(apollo: Apollo.Apollo) {
      super(apollo);
    }
  }
export const MoveOrderLineDocument = gql`
    mutation MoveOrderLine($orderLineID: ID!, $colliID: ID!) {
  moveOrderLine(colliID: $colliID, orderLineID: $orderLineID) {
    ...ColliViewInfo
  }
}
    ${ColliViewInfoFragmentDoc}`;

  @Injectable({
    providedIn: 'root'
  })
  export class MoveOrderLineGQL extends Apollo.Mutation<MoveOrderLineMutation, MoveOrderLineMutationVariables> {
    document = MoveOrderLineDocument;
    
    constructor(apollo: Apollo.Apollo) {
      super(apollo);
    }
  }
export const DeletePackageDocument = gql`
    mutation DeletePackage($colliID: ID!) {
  deleteColli(colliID: $colliID) {
    ...OrderView
  }
}
    ${OrderViewFragmentDoc}`;

  @Injectable({
    providedIn: 'root'
  })
  export class DeletePackageGQL extends Apollo.Mutation<DeletePackageMutation, DeletePackageMutationVariables> {
    document = DeletePackageDocument;
    
    constructor(apollo: Apollo.Apollo) {
      super(apollo);
    }
  }
export const DuplicateColliDocument = gql`
    mutation DuplicateColli($fromColliID: ID!) {
  duplicateColli(fromColliID: $fromColliID) {
    ...OrderView
  }
}
    ${OrderViewFragmentDoc}`;

  @Injectable({
    providedIn: 'root'
  })
  export class DuplicateColliGQL extends Apollo.Mutation<DuplicateColliMutation, DuplicateColliMutationVariables> {
    document = DuplicateColliDocument;
    
    constructor(apollo: Apollo.Apollo) {
      super(apollo);
    }
  }
export const CreateShipmentsDocument = gql`
    query CreateShipments($orderID: ID!, $parcelIDs: [ID!]!) {
  createShipments(orderID: $orderID, packageIDs: $parcelIDs) {
    labelsPDF
    allLabels
  }
}
    `;

  @Injectable({
    providedIn: 'root'
  })
  export class CreateShipmentsGQL extends Apollo.Query<CreateShipmentsQuery, CreateShipmentsQueryVariables> {
    document = CreateShipmentsDocument;
    
    constructor(apollo: Apollo.Apollo) {
      super(apollo);
    }
  }
export const CreatePackingSlipPrintJobsDocument = gql`
    query CreatePackingSlipPrintJobs($colliIDs: [ID!]!) {
  createPackingListPrintJob(colliIDs: $colliIDs)
}
    `;

  @Injectable({
    providedIn: 'root'
  })
  export class CreatePackingSlipPrintJobsGQL extends Apollo.Query<CreatePackingSlipPrintJobsQuery, CreatePackingSlipPrintJobsQueryVariables> {
    document = CreatePackingSlipPrintJobsDocument;
    
    constructor(apollo: Apollo.Apollo) {
      super(apollo);
    }
  }
export const CreateLabelPrintJobsDocument = gql`
    query CreateLabelPrintJobs($colliIDs: [ID!]!) {
  createLabelsPrintJob(colliIDs: $colliIDs)
}
    `;

  @Injectable({
    providedIn: 'root'
  })
  export class CreateLabelPrintJobsGQL extends Apollo.Query<CreateLabelPrintJobsQuery, CreateLabelPrintJobsQueryVariables> {
    document = CreateLabelPrintJobsDocument;
    
    constructor(apollo: Apollo.Apollo) {
      super(apollo);
    }
  }
export const PackingSlipsClearCacheDocument = gql`
    query PackingSlipsClearCache($orderIDs: [ID!]!) {
  packingSlipsClearCache(orderIDs: $orderIDs)
}
    `;

  @Injectable({
    providedIn: 'root'
  })
  export class PackingSlipsClearCacheGQL extends Apollo.Query<PackingSlipsClearCacheQuery, PackingSlipsClearCacheQueryVariables> {
    document = PackingSlipsClearCacheDocument;
    
    constructor(apollo: Apollo.Apollo) {
      super(apollo);
    }
  }