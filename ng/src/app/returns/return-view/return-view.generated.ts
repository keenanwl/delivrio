/* eslint-disable */
import * as Types from '../../../generated/graphql';

import { gql } from 'apollo-angular';
import { OrderLinesFragmentDoc, AddressInfoFragmentDoc } from '../../orders/order-edit/order-edit.generated';
import { TimelineViewerFragmentDoc } from '../../shared/timeline-viewer/timeline-viewer.generated';
import { Injectable } from '@angular/core';
import * as Apollo from 'apollo-angular';
export type FetchReturnCollisViewQueryVariables = Types.Exact<{
  orderID: Types.Scalars['ID'];
}>;


export type FetchReturnCollisViewQuery = { returnColli: { order: { orderPublicID: string }, collis: Array<{ colli: { id: string, labelPdf?: string | null, status: Types.ReturnColliStatus, deliveryOption?: { id: string, name: string, carrier: { carrierBrand: { label: string } } } | null, returnOrderLine?: Array<{ orderLine: { id: string, units: number, unitPrice: number, discountAllocationAmount: number, currency: { display: string }, productVariant: { id: string, dimensionLength?: number | null, dimensionWidth?: number | null, dimensionHeight?: number | null, weightG?: number | null, description?: string | null, archived: boolean, productImage?: Array<{ url: string }> | null, product: { id: string, externalID?: string | null, title: string, status: Types.ProductStatus } } } }> | null, sender: { id: string, firstName: string, lastName: string, phoneNumber: string, email: string, addressOne: string, addressTwo: string, zip: string, city: string, state?: string | null, company?: string | null, country: { id: string, alpha2: string, label: string } }, recipient: { id: string, firstName: string, lastName: string, phoneNumber: string, email: string, addressOne: string, addressTwo: string, zip: string, city: string, state?: string | null, company?: string | null, country: { id: string, alpha2: string, label: string } } }, timeline: Array<{ id: string, createdAt: any, user?: { id: string, name?: string | null } | null, orderHistory?: Array<{ id: string, type: Types.OrderHistoryType, description: string, order: { id: string, orderPublicID: string } }> | null, shipmentHistory?: Array<{ id: string, type: Types.ShipmentHistoryType }> | null, returnColliHistory?: Array<{ id: string, description: string, type: Types.ReturnColliHistoryType }> | null }> }> } };

export type ReturnColliEditInfoFragment = { order: { orderPublicID: string }, collis: Array<{ colli: { id: string, labelPdf?: string | null, status: Types.ReturnColliStatus, deliveryOption?: { id: string, name: string, carrier: { carrierBrand: { label: string } } } | null, returnOrderLine?: Array<{ orderLine: { id: string, units: number, unitPrice: number, discountAllocationAmount: number, currency: { display: string }, productVariant: { id: string, dimensionLength?: number | null, dimensionWidth?: number | null, dimensionHeight?: number | null, weightG?: number | null, description?: string | null, archived: boolean, productImage?: Array<{ url: string }> | null, product: { id: string, externalID?: string | null, title: string, status: Types.ProductStatus } } } }> | null, sender: { id: string, firstName: string, lastName: string, phoneNumber: string, email: string, addressOne: string, addressTwo: string, zip: string, city: string, state?: string | null, company?: string | null, country: { id: string, alpha2: string, label: string } }, recipient: { id: string, firstName: string, lastName: string, phoneNumber: string, email: string, addressOne: string, addressTwo: string, zip: string, city: string, state?: string | null, company?: string | null, country: { id: string, alpha2: string, label: string } } }, timeline: Array<{ id: string, createdAt: any, user?: { id: string, name?: string | null } | null, orderHistory?: Array<{ id: string, type: Types.OrderHistoryType, description: string, order: { id: string, orderPublicID: string } }> | null, shipmentHistory?: Array<{ id: string, type: Types.ShipmentHistoryType }> | null, returnColliHistory?: Array<{ id: string, description: string, type: Types.ReturnColliHistoryType }> | null }> }> };

export type UpdateReturnColliStatusMutationVariables = Types.Exact<{
  returnColliID: Types.Scalars['ID'];
  status: Types.ReturnColliStatus;
}>;


export type UpdateReturnColliStatusMutation = { updateReturnColliStatus: { order: { orderPublicID: string }, collis: Array<{ colli: { id: string, labelPdf?: string | null, status: Types.ReturnColliStatus, deliveryOption?: { id: string, name: string, carrier: { carrierBrand: { label: string } } } | null, returnOrderLine?: Array<{ orderLine: { id: string, units: number, unitPrice: number, discountAllocationAmount: number, currency: { display: string }, productVariant: { id: string, dimensionLength?: number | null, dimensionWidth?: number | null, dimensionHeight?: number | null, weightG?: number | null, description?: string | null, archived: boolean, productImage?: Array<{ url: string }> | null, product: { id: string, externalID?: string | null, title: string, status: Types.ProductStatus } } } }> | null, sender: { id: string, firstName: string, lastName: string, phoneNumber: string, email: string, addressOne: string, addressTwo: string, zip: string, city: string, state?: string | null, company?: string | null, country: { id: string, alpha2: string, label: string } }, recipient: { id: string, firstName: string, lastName: string, phoneNumber: string, email: string, addressOne: string, addressTwo: string, zip: string, city: string, state?: string | null, company?: string | null, country: { id: string, alpha2: string, label: string } } }, timeline: Array<{ id: string, createdAt: any, user?: { id: string, name?: string | null } | null, orderHistory?: Array<{ id: string, type: Types.OrderHistoryType, description: string, order: { id: string, orderPublicID: string } }> | null, shipmentHistory?: Array<{ id: string, type: Types.ShipmentHistoryType }> | null, returnColliHistory?: Array<{ id: string, description: string, type: Types.ReturnColliHistoryType }> | null }> }> } };

export const ReturnColliEditInfoFragmentDoc = gql`
    fragment ReturnColliEditInfo on ReturnColliEdit {
  order {
    orderPublicID
  }
  collis {
    colli {
      id
      deliveryOption {
        id
        name
        carrier {
          carrierBrand {
            label
          }
        }
      }
      labelPdf
      status
      returnOrderLine {
        orderLine {
          ...OrderLines
        }
      }
      sender: sender {
        ...AddressInfo
      }
      recipient: recipient {
        ...AddressInfo
      }
    }
    timeline {
      ...TimelineViewer
    }
  }
}
    ${OrderLinesFragmentDoc}
${AddressInfoFragmentDoc}
${TimelineViewerFragmentDoc}`;
export const FetchReturnCollisViewDocument = gql`
    query FetchReturnCollisView($orderID: ID!) {
  returnColli(orderID: $orderID) {
    ...ReturnColliEditInfo
  }
}
    ${ReturnColliEditInfoFragmentDoc}`;

  @Injectable({
    providedIn: 'root'
  })
  export class FetchReturnCollisViewGQL extends Apollo.Query<FetchReturnCollisViewQuery, FetchReturnCollisViewQueryVariables> {
    document = FetchReturnCollisViewDocument;
    
    constructor(apollo: Apollo.Apollo) {
      super(apollo);
    }
  }
export const UpdateReturnColliStatusDocument = gql`
    mutation UpdateReturnColliStatus($returnColliID: ID!, $status: ReturnColliStatus!) {
  updateReturnColliStatus(returnColliID: $returnColliID, status: $status) {
    ...ReturnColliEditInfo
  }
}
    ${ReturnColliEditInfoFragmentDoc}`;

  @Injectable({
    providedIn: 'root'
  })
  export class UpdateReturnColliStatusGQL extends Apollo.Mutation<UpdateReturnColliStatusMutation, UpdateReturnColliStatusMutationVariables> {
    document = UpdateReturnColliStatusDocument;
    
    constructor(apollo: Apollo.Apollo) {
      super(apollo);
    }
  }