/* eslint-disable */
import * as Types from '../../../../generated/graphql';

import { gql } from 'apollo-angular';
import { Injectable } from '@angular/core';
import * as Apollo from 'apollo-angular';
export type FetchPackagingListQueryVariables = Types.Exact<{ [key: string]: never; }>;


export type FetchPackagingListQuery = { packagingFiltered: Array<{ id: string, name: string, lengthCm: number, heightCm: number, widthCm: number, carrierBrand?: { id: string, internalID: Types.CarrierBrandInternalId, label: string } | null }>, carrierBrands: { edges?: Array<{ node?: { id: string, internalID: Types.CarrierBrandInternalId, label: string } | null } | null> | null }, packagingUSPSRateIndicators: { edges?: Array<{ node?: { id: string, name: string } | null } | null> | null }, packagingUSPSProcessingCategories: { edges?: Array<{ node?: { id: string, name: string } | null } | null> | null } };

export type PackagingListItemFragment = { id: string, name: string, lengthCm: number, heightCm: number, widthCm: number, carrierBrand?: { id: string, internalID: Types.CarrierBrandInternalId, label: string } | null };

export type CreatePackagingMutationVariables = Types.Exact<{
  input: Types.CreatePackagingInput;
  inputPackagingUSPS?: Types.InputMaybe<Types.CreatePackagingUspsInput>;
  inputPackagingDF?: Types.InputMaybe<Types.CreatePackagingDfInput>;
}>;


export type CreatePackagingMutation = { createPackaging: Array<{ id: string, name: string, lengthCm: number, heightCm: number, widthCm: number, carrierBrand?: { id: string, internalID: Types.CarrierBrandInternalId, label: string } | null }> };

export type UpdatePackagingMutationVariables = Types.Exact<{
  input: Types.UpdatePackagingInput;
}>;


export type UpdatePackagingMutation = { updatePackaging: Array<{ id: string, name: string, lengthCm: number, heightCm: number, widthCm: number, carrierBrand?: { id: string, internalID: Types.CarrierBrandInternalId, label: string } | null }> };

export type ArchivePackagingMutationVariables = Types.Exact<{
  id: Types.Scalars['ID'];
}>;


export type ArchivePackagingMutation = { archivePackaging: boolean };

export const PackagingListItemFragmentDoc = gql`
    fragment PackagingListItem on Packaging {
  id
  name
  lengthCm
  heightCm
  widthCm
  carrierBrand {
    id
    internalID
    label
  }
}
    `;
export const FetchPackagingListDocument = gql`
    query FetchPackagingList {
  packagingFiltered(showArchived: false) {
    ...PackagingListItem
  }
  carrierBrands {
    edges {
      node {
        id
        internalID
        label
      }
    }
  }
  packagingUSPSRateIndicators {
    edges {
      node {
        id
        name
      }
    }
  }
  packagingUSPSProcessingCategories {
    edges {
      node {
        id
        name
      }
    }
  }
}
    ${PackagingListItemFragmentDoc}`;

  @Injectable({
    providedIn: 'root'
  })
  export class FetchPackagingListGQL extends Apollo.Query<FetchPackagingListQuery, FetchPackagingListQueryVariables> {
    document = FetchPackagingListDocument;
    
    constructor(apollo: Apollo.Apollo) {
      super(apollo);
    }
  }
export const CreatePackagingDocument = gql`
    mutation CreatePackaging($input: CreatePackagingInput!, $inputPackagingUSPS: CreatePackagingUSPSInput, $inputPackagingDF: CreatePackagingDFInput) {
  createPackaging(
    input: $input
    inputPackagingUSPS: $inputPackagingUSPS
    inputPackagingDF: $inputPackagingDF
  ) {
    ...PackagingListItem
  }
}
    ${PackagingListItemFragmentDoc}`;

  @Injectable({
    providedIn: 'root'
  })
  export class CreatePackagingGQL extends Apollo.Mutation<CreatePackagingMutation, CreatePackagingMutationVariables> {
    document = CreatePackagingDocument;
    
    constructor(apollo: Apollo.Apollo) {
      super(apollo);
    }
  }
export const UpdatePackagingDocument = gql`
    mutation UpdatePackaging($input: UpdatePackagingInput!) {
  updatePackaging(input: $input) {
    ...PackagingListItem
  }
}
    ${PackagingListItemFragmentDoc}`;

  @Injectable({
    providedIn: 'root'
  })
  export class UpdatePackagingGQL extends Apollo.Mutation<UpdatePackagingMutation, UpdatePackagingMutationVariables> {
    document = UpdatePackagingDocument;
    
    constructor(apollo: Apollo.Apollo) {
      super(apollo);
    }
  }
export const ArchivePackagingDocument = gql`
    mutation ArchivePackaging($id: ID!) {
  archivePackaging(id: $id)
}
    `;

  @Injectable({
    providedIn: 'root'
  })
  export class ArchivePackagingGQL extends Apollo.Mutation<ArchivePackagingMutation, ArchivePackagingMutationVariables> {
    document = ArchivePackagingDocument;
    
    constructor(apollo: Apollo.Apollo) {
      super(apollo);
    }
  }