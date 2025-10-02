/* eslint-disable */
import * as Types from '../../../../generated/graphql';

import { gql } from 'apollo-angular';
import { Injectable } from '@angular/core';
import * as Apollo from 'apollo-angular';
export type ListDeliveryOptionsQueryVariables = Types.Exact<{
  showArchived: Types.Scalars['Boolean'];
}>;


export type ListDeliveryOptionsQuery = { deliveryOptionsFiltered: Array<{ id: string, name: string, carrier: { name: string, carrierBrand: { internalID: Types.CarrierBrandInternalId, labelShort: string, backgroundColor?: string | null, textColor?: string | null, logoURL?: string | null } }, carrierService: { label: string }, connection: { name: string } }> };

export type DeliveryOptionListItemFragment = { id: string, name: string, carrier: { name: string, carrierBrand: { internalID: Types.CarrierBrandInternalId, labelShort: string, backgroundColor?: string | null, textColor?: string | null, logoURL?: string | null } }, carrierService: { label: string }, connection: { name: string } };

export type FetchCarrierAgreementsAndConnectionsQueryVariables = Types.Exact<{ [key: string]: never; }>;


export type FetchCarrierAgreementsAndConnectionsQuery = { carriers: { edges?: Array<{ node?: { id: string, name: string, carrierBrand: { label: string } } | null } | null> | null }, connections: { edges?: Array<{ node?: { id: string, name: string, connectionBrand: { label: string } } | null } | null> | null } };

export type CreateDeliveryOptionMutationVariables = Types.Exact<{
  name: Types.Scalars['String'];
  agreementID: Types.Scalars['ID'];
  connectionID: Types.Scalars['ID'];
}>;


export type CreateDeliveryOptionMutation = { createDeliveryOption: { id: string, carrier: Types.CarrierBrandInternalId } };

export type UpdateDeliveryOptionSortOrderMutationVariables = Types.Exact<{
  nextSortOrder: Array<Types.Scalars['ID']> | Types.Scalars['ID'];
}>;


export type UpdateDeliveryOptionSortOrderMutation = { updateDeliveryOptionSortOrder: Array<{ id: string, name: string, carrier: { name: string, carrierBrand: { internalID: Types.CarrierBrandInternalId, labelShort: string, backgroundColor?: string | null, textColor?: string | null, logoURL?: string | null } }, carrierService: { label: string }, connection: { name: string } }> };

export type DeliveryOptionArchiveMutationVariables = Types.Exact<{
  id: Types.Scalars['ID'];
}>;


export type DeliveryOptionArchiveMutation = { deliveryOptionArchive: boolean };

export const DeliveryOptionListItemFragmentDoc = gql`
    fragment DeliveryOptionListItem on DeliveryOption {
  id
  name
  carrier {
    name
    carrierBrand {
      internalID
      labelShort
      backgroundColor
      textColor
      logoURL
    }
  }
  carrierService {
    label
  }
  connection {
    name
  }
}
    `;
export const ListDeliveryOptionsDocument = gql`
    query ListDeliveryOptions($showArchived: Boolean!) {
  deliveryOptionsFiltered(showArchived: $showArchived) {
    ...DeliveryOptionListItem
  }
}
    ${DeliveryOptionListItemFragmentDoc}`;

  @Injectable({
    providedIn: 'root'
  })
  export class ListDeliveryOptionsGQL extends Apollo.Query<ListDeliveryOptionsQuery, ListDeliveryOptionsQueryVariables> {
    document = ListDeliveryOptionsDocument;
    
    constructor(apollo: Apollo.Apollo) {
      super(apollo);
    }
  }
export const FetchCarrierAgreementsAndConnectionsDocument = gql`
    query FetchCarrierAgreementsAndConnections {
  carriers {
    edges {
      node {
        id
        name
        carrierBrand {
          label
        }
      }
    }
  }
  connections {
    edges {
      node {
        id
        name
        connectionBrand {
          label
        }
      }
    }
  }
}
    `;

  @Injectable({
    providedIn: 'root'
  })
  export class FetchCarrierAgreementsAndConnectionsGQL extends Apollo.Query<FetchCarrierAgreementsAndConnectionsQuery, FetchCarrierAgreementsAndConnectionsQueryVariables> {
    document = FetchCarrierAgreementsAndConnectionsDocument;
    
    constructor(apollo: Apollo.Apollo) {
      super(apollo);
    }
  }
export const CreateDeliveryOptionDocument = gql`
    mutation CreateDeliveryOption($name: String!, $agreementID: ID!, $connectionID: ID!) {
  createDeliveryOption(
    name: $name
    agreementID: $agreementID
    connectionID: $connectionID
  ) {
    id
    carrier
  }
}
    `;

  @Injectable({
    providedIn: 'root'
  })
  export class CreateDeliveryOptionGQL extends Apollo.Mutation<CreateDeliveryOptionMutation, CreateDeliveryOptionMutationVariables> {
    document = CreateDeliveryOptionDocument;
    
    constructor(apollo: Apollo.Apollo) {
      super(apollo);
    }
  }
export const UpdateDeliveryOptionSortOrderDocument = gql`
    mutation UpdateDeliveryOptionSortOrder($nextSortOrder: [ID!]!) {
  updateDeliveryOptionSortOrder(newOrder: $nextSortOrder) {
    ...DeliveryOptionListItem
  }
}
    ${DeliveryOptionListItemFragmentDoc}`;

  @Injectable({
    providedIn: 'root'
  })
  export class UpdateDeliveryOptionSortOrderGQL extends Apollo.Mutation<UpdateDeliveryOptionSortOrderMutation, UpdateDeliveryOptionSortOrderMutationVariables> {
    document = UpdateDeliveryOptionSortOrderDocument;
    
    constructor(apollo: Apollo.Apollo) {
      super(apollo);
    }
  }
export const DeliveryOptionArchiveDocument = gql`
    mutation DeliveryOptionArchive($id: ID!) {
  deliveryOptionArchive(deliveryOptionID: $id)
}
    `;

  @Injectable({
    providedIn: 'root'
  })
  export class DeliveryOptionArchiveGQL extends Apollo.Mutation<DeliveryOptionArchiveMutation, DeliveryOptionArchiveMutationVariables> {
    document = DeliveryOptionArchiveDocument;
    
    constructor(apollo: Apollo.Apollo) {
      super(apollo);
    }
  }