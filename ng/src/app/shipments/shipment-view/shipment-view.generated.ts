/* eslint-disable */
import * as Types from '../../../generated/graphql';

import { gql } from 'apollo-angular';
import { Injectable } from '@angular/core';
import * as Apollo from 'apollo-angular';
export type FetchShipmentQueryVariables = Types.Exact<{
  id: Types.Scalars['ID'];
}>;


export type FetchShipmentQuery = { shipment: { id: string, shipmentPublicID: string, createdAt: any, status: Types.ShipmentStatus, shipmentUSPS?: { postage?: number | null, scheduledDeliveryDate?: any | null, trackingNumber?: string | null } | null, shipmentEasyPost?: { rate?: number | null, estDeliveryDate?: any | null, trackingNumber?: string | null } | null, shipmentParcel?: Array<{ id: string, cancelSyncedAt?: any | null, fulfillmentSyncedAt?: any | null, status: Types.ShipmentParcelStatus, itemID?: string | null, colli?: { deliveryOption?: { carrier: { carrierBrand: { label: string } } } | null, order: { id: string, orderPublicID: string } } | null }> | null, shipmentPallet?: Array<{ id: string, status: Types.ShipmentPalletStatus, barcode: string, labelPdf?: string | null, pallet?: { consolidation: { deliveryOption?: { carrier: { carrierBrand: { label: string } } } | null } } | null }> | null, shipmentPostNord?: { bookingID: string, shipmentReferenceNo: string } | null } };

export type ShipmentViewCancelShipmentMutationVariables = Types.Exact<{
  shipmentID: Types.Scalars['ID'];
}>;


export type ShipmentViewCancelShipmentMutation = { cancelShipment: { id: string } };

export type DebugUpdateLabelIDsMutationVariables = Types.Exact<{
  parcelID: Types.Scalars['ID'];
  itemID: Types.Scalars['String'];
}>;


export type DebugUpdateLabelIDsMutation = { debugUpdateLabelIDs: boolean };

export type CancelFulfillmentSyncMutationVariables = Types.Exact<{
  shipmentParcelID: Types.Scalars['ID'];
}>;


export type CancelFulfillmentSyncMutation = { cancelFulfillmentSync: boolean };

export type CancelCancelSyncMutationVariables = Types.Exact<{
  shipmentParcelID: Types.Scalars['ID'];
}>;


export type CancelCancelSyncMutation = { cancelCancelSync: boolean };

export const FetchShipmentDocument = gql`
    query FetchShipment($id: ID!) {
  shipment(id: $id) {
    id
    shipmentPublicID
    createdAt
    shipmentUSPS {
      postage
      scheduledDeliveryDate
      trackingNumber
    }
    shipmentEasyPost {
      rate
      estDeliveryDate
      trackingNumber
    }
    shipmentParcel {
      id
      cancelSyncedAt
      fulfillmentSyncedAt
      status
      colli {
        deliveryOption {
          carrier {
            carrierBrand {
              label
            }
          }
        }
        order {
          id
          orderPublicID
        }
      }
      itemID
    }
    shipmentPallet {
      id
      status
      barcode
      labelPdf
      pallet {
        consolidation {
          deliveryOption {
            carrier {
              carrierBrand {
                label
              }
            }
          }
        }
      }
    }
    shipmentPostNord {
      bookingID
      shipmentReferenceNo
    }
    status
  }
}
    `;

  @Injectable({
    providedIn: 'root'
  })
  export class FetchShipmentGQL extends Apollo.Query<FetchShipmentQuery, FetchShipmentQueryVariables> {
    document = FetchShipmentDocument;
    
    constructor(apollo: Apollo.Apollo) {
      super(apollo);
    }
  }
export const ShipmentViewCancelShipmentDocument = gql`
    mutation ShipmentViewCancelShipment($shipmentID: ID!) {
  cancelShipment(shipmentID: $shipmentID) {
    id
  }
}
    `;

  @Injectable({
    providedIn: 'root'
  })
  export class ShipmentViewCancelShipmentGQL extends Apollo.Mutation<ShipmentViewCancelShipmentMutation, ShipmentViewCancelShipmentMutationVariables> {
    document = ShipmentViewCancelShipmentDocument;
    
    constructor(apollo: Apollo.Apollo) {
      super(apollo);
    }
  }
export const DebugUpdateLabelIDsDocument = gql`
    mutation DebugUpdateLabelIDs($parcelID: ID!, $itemID: String!) {
  debugUpdateLabelIDs(parcelID: $parcelID, itemID: $itemID)
}
    `;

  @Injectable({
    providedIn: 'root'
  })
  export class DebugUpdateLabelIDsGQL extends Apollo.Mutation<DebugUpdateLabelIDsMutation, DebugUpdateLabelIDsMutationVariables> {
    document = DebugUpdateLabelIDsDocument;
    
    constructor(apollo: Apollo.Apollo) {
      super(apollo);
    }
  }
export const CancelFulfillmentSyncDocument = gql`
    mutation cancelFulfillmentSync($shipmentParcelID: ID!) {
  cancelFulfillmentSync(shipmentParcelID: $shipmentParcelID)
}
    `;

  @Injectable({
    providedIn: 'root'
  })
  export class CancelFulfillmentSyncGQL extends Apollo.Mutation<CancelFulfillmentSyncMutation, CancelFulfillmentSyncMutationVariables> {
    document = CancelFulfillmentSyncDocument;
    
    constructor(apollo: Apollo.Apollo) {
      super(apollo);
    }
  }
export const CancelCancelSyncDocument = gql`
    mutation cancelCancelSync($shipmentParcelID: ID!) {
  cancelCancelSync(shipmentParcelID: $shipmentParcelID)
}
    `;

  @Injectable({
    providedIn: 'root'
  })
  export class CancelCancelSyncGQL extends Apollo.Mutation<CancelCancelSyncMutation, CancelCancelSyncMutationVariables> {
    document = CancelCancelSyncDocument;
    
    constructor(apollo: Apollo.Apollo) {
      super(apollo);
    }
  }