/* eslint-disable */
import * as Types from '../../../generated/graphql';

import { gql } from 'apollo-angular';
import { Injectable } from '@angular/core';
import * as Apollo from 'apollo-angular';
export type FetchProductQueryVariables = Types.Exact<{
  id: Types.Scalars['ID'];
}>;


export type FetchProductQuery = { product?: { title: string, status: Types.ProductStatus, bodyHTML?: string | null, createdAt?: any | null, productTags?: Array<{ id: string, name: string }> | null } | null, productVariants: { edges?: Array<{ node?: { id: string, archived: boolean, description?: string | null, dimensionWidth?: number | null, dimensionLength?: number | null, dimensionHeight?: number | null, weightG?: number | null, eanNumber?: string | null, createdAt?: any | null, updatedAt: any } | null } | null> | null }, productForImage?: { productImage?: Array<{ id: string, externalID?: string | null, url: string, productVariant?: Array<{ id: string }> | null }> | null } | null, productTags: { edges?: Array<{ node?: { id: string, name: string, createdAt?: any | null } | null } | null> | null } };

export type ProductsSearchCountriesQueryVariables = Types.Exact<{
  term: Types.Scalars['String'];
}>;


export type ProductsSearchCountriesQuery = { countries: { edges?: Array<{ node?: { id: string, label: string, alpha2: string } | null } | null> | null } };

export type MustInventoryItemMutationVariables = Types.Exact<{
  productVariantID: Types.Scalars['ID'];
}>;


export type MustInventoryItemMutation = { mustInventory: { id: string, code?: string | null, sku?: string | null, countryOfOrigin?: { id: string, label: string, alpha2: string } | null, countryHarmonizedCode?: Array<{ code: string, country: { id: string, label: string, alpha2: string } }> | null } };

export type ProductVariantInfoFragment = { id: string, archived: boolean, description?: string | null, dimensionWidth?: number | null, dimensionLength?: number | null, dimensionHeight?: number | null, weightG?: number | null, eanNumber?: string | null, createdAt?: any | null, updatedAt: any };

export type InventoryItemFragFragment = { id: string, code?: string | null, sku?: string | null, countryOfOrigin?: { id: string, label: string, alpha2: string } | null, countryHarmonizedCode?: Array<{ code: string, country: { id: string, label: string, alpha2: string } }> | null };

export type ProductImageInfoFragment = { id: string, externalID?: string | null, url: string, productVariant?: Array<{ id: string }> | null };

export type CreateProductMutationVariables = Types.Exact<{
  input: Types.CreateProductInput;
  variants?: Types.InputMaybe<Array<Types.CreateProductVariantInput> | Types.CreateProductVariantInput>;
  images?: Types.InputMaybe<Array<Types.Scalars['String']> | Types.Scalars['String']>;
}>;


export type CreateProductMutation = { createProduct?: { id: string } | null };

export type UpdateProductMutationVariables = Types.Exact<{
  id: Types.Scalars['ID'];
  input: Types.UpdateProductInput;
  variants?: Types.InputMaybe<Array<Types.UpdateProductVariantIdInput> | Types.UpdateProductVariantIdInput>;
  images: Array<Types.ProductVariantImageInput> | Types.ProductVariantImageInput;
}>;


export type UpdateProductMutation = { updateProduct?: { id: string } | null };

export type CreateVariantMutationVariables = Types.Exact<{
  productID: Types.Scalars['ID'];
  input?: Types.InputMaybe<Types.CreateProductVariantInput>;
}>;


export type CreateVariantMutation = { createVariant: { id: string, description?: string | null, dimensionWidth?: number | null, dimensionLength?: number | null, dimensionHeight?: number | null, weightG?: number | null, eanNumber?: string | null, createdAt?: any | null, updatedAt: any } };

export type ProductUploadImageMutationVariables = Types.Exact<{
  productID: Types.Scalars['ID'];
  image: Types.Scalars['String'];
}>;


export type ProductUploadImageMutation = { uploadProductImage: { productImage?: Array<{ id: string, externalID?: string | null, url: string, productVariant?: Array<{ id: string }> | null }> | null } };

export type ProductDeleteImageMutationVariables = Types.Exact<{
  imageID: Types.Scalars['ID'];
}>;


export type ProductDeleteImageMutation = { deleteProductImage: { productImage?: Array<{ id: string, externalID?: string | null, url: string, productVariant?: Array<{ id: string }> | null }> | null } };

export type ArchiveVariantMutationVariables = Types.Exact<{
  variantID: Types.Scalars['ID'];
}>;


export type ArchiveVariantMutation = { archiveProductVariant: { id: string } };

export type MustInventoryMutationVariables = Types.Exact<{
  productVariantID: Types.Scalars['ID'];
}>;


export type MustInventoryMutation = { mustInventory: { id: string, code?: string | null, sku?: string | null, countryOfOrigin?: { label: string } | null, countryHarmonizedCode?: Array<{ code: string, country: { id: string, label: string } }> | null } };

export type UpdateInventoryItemMutationVariables = Types.Exact<{
  id: Types.Scalars['ID'];
  input: Types.UpdateInventoryItemInput;
}>;


export type UpdateInventoryItemMutation = { updateInventory: { id: string, code?: string | null, sku?: string | null, countryOfOrigin?: { id: string, label: string, alpha2: string } | null, countryHarmonizedCode?: Array<{ code: string, country: { id: string, label: string, alpha2: string } }> | null } };

export type ProductSearchCountriesQueryVariables = Types.Exact<{
  term: Types.Scalars['String'];
}>;


export type ProductSearchCountriesQuery = { countries: { edges?: Array<{ node?: { id: string, label: string, alpha2: string, region: Types.CountryRegion } | null } | null> | null } };

export const ProductVariantInfoFragmentDoc = gql`
    fragment ProductVariantInfo on ProductVariant {
  id
  archived
  description
  dimensionWidth
  dimensionLength
  dimensionHeight
  weightG
  eanNumber
  createdAt
  updatedAt
}
    `;
export const InventoryItemFragFragmentDoc = gql`
    fragment InventoryItemFrag on InventoryItem {
  id
  code
  sku
  countryOfOrigin {
    id
    label
    alpha2
  }
  countryHarmonizedCode {
    code
    country {
      id
      label
      alpha2
    }
  }
}
    `;
export const ProductImageInfoFragmentDoc = gql`
    fragment ProductImageInfo on ProductImage {
  id
  externalID
  url
  productVariant {
    id
  }
}
    `;
export const FetchProductDocument = gql`
    query FetchProduct($id: ID!) {
  product(id: $id) {
    title
    status
    bodyHTML
    createdAt
    productTags {
      id
      name
    }
  }
  productVariants(where: {and: {archived: false, hasProductWith: {id: $id}}}) {
    edges {
      node {
        ...ProductVariantInfo
      }
    }
  }
  productForImage: product(id: $id) {
    productImage {
      ...ProductImageInfo
    }
  }
  productTags {
    edges {
      node {
        id
        name
        createdAt
      }
    }
  }
}
    ${ProductVariantInfoFragmentDoc}
${ProductImageInfoFragmentDoc}`;

  @Injectable({
    providedIn: 'root'
  })
  export class FetchProductGQL extends Apollo.Query<FetchProductQuery, FetchProductQueryVariables> {
    document = FetchProductDocument;
    
    constructor(apollo: Apollo.Apollo) {
      super(apollo);
    }
  }
export const ProductsSearchCountriesDocument = gql`
    query ProductsSearchCountries($term: String!) {
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
  export class ProductsSearchCountriesGQL extends Apollo.Query<ProductsSearchCountriesQuery, ProductsSearchCountriesQueryVariables> {
    document = ProductsSearchCountriesDocument;
    
    constructor(apollo: Apollo.Apollo) {
      super(apollo);
    }
  }
export const MustInventoryItemDocument = gql`
    mutation MustInventoryItem($productVariantID: ID!) {
  mustInventory(productVariantID: $productVariantID) {
    ...InventoryItemFrag
  }
}
    ${InventoryItemFragFragmentDoc}`;

  @Injectable({
    providedIn: 'root'
  })
  export class MustInventoryItemGQL extends Apollo.Mutation<MustInventoryItemMutation, MustInventoryItemMutationVariables> {
    document = MustInventoryItemDocument;
    
    constructor(apollo: Apollo.Apollo) {
      super(apollo);
    }
  }
export const CreateProductDocument = gql`
    mutation CreateProduct($input: CreateProductInput!, $variants: [CreateProductVariantInput!], $images: [String!]) {
  createProduct(input: $input, variants: $variants, images: $images) {
    id
  }
}
    `;

  @Injectable({
    providedIn: 'root'
  })
  export class CreateProductGQL extends Apollo.Mutation<CreateProductMutation, CreateProductMutationVariables> {
    document = CreateProductDocument;
    
    constructor(apollo: Apollo.Apollo) {
      super(apollo);
    }
  }
export const UpdateProductDocument = gql`
    mutation UpdateProduct($id: ID!, $input: UpdateProductInput!, $variants: [UpdateProductVariantIDInput!], $images: [ProductVariantImageInput!]!) {
  updateProduct(id: $id, input: $input, variants: $variants, images: $images) {
    id
  }
}
    `;

  @Injectable({
    providedIn: 'root'
  })
  export class UpdateProductGQL extends Apollo.Mutation<UpdateProductMutation, UpdateProductMutationVariables> {
    document = UpdateProductDocument;
    
    constructor(apollo: Apollo.Apollo) {
      super(apollo);
    }
  }
export const CreateVariantDocument = gql`
    mutation CreateVariant($productID: ID!, $input: CreateProductVariantInput) {
  createVariant(productID: $productID, input: $input) {
    id
    description
    dimensionWidth
    dimensionLength
    dimensionHeight
    weightG
    eanNumber
    createdAt
    updatedAt
  }
}
    `;

  @Injectable({
    providedIn: 'root'
  })
  export class CreateVariantGQL extends Apollo.Mutation<CreateVariantMutation, CreateVariantMutationVariables> {
    document = CreateVariantDocument;
    
    constructor(apollo: Apollo.Apollo) {
      super(apollo);
    }
  }
export const ProductUploadImageDocument = gql`
    mutation ProductUploadImage($productID: ID!, $image: String!) {
  uploadProductImage(productID: $productID, image: $image) {
    productImage {
      ...ProductImageInfo
    }
  }
}
    ${ProductImageInfoFragmentDoc}`;

  @Injectable({
    providedIn: 'root'
  })
  export class ProductUploadImageGQL extends Apollo.Mutation<ProductUploadImageMutation, ProductUploadImageMutationVariables> {
    document = ProductUploadImageDocument;
    
    constructor(apollo: Apollo.Apollo) {
      super(apollo);
    }
  }
export const ProductDeleteImageDocument = gql`
    mutation ProductDeleteImage($imageID: ID!) {
  deleteProductImage(imageID: $imageID) {
    productImage {
      ...ProductImageInfo
    }
  }
}
    ${ProductImageInfoFragmentDoc}`;

  @Injectable({
    providedIn: 'root'
  })
  export class ProductDeleteImageGQL extends Apollo.Mutation<ProductDeleteImageMutation, ProductDeleteImageMutationVariables> {
    document = ProductDeleteImageDocument;
    
    constructor(apollo: Apollo.Apollo) {
      super(apollo);
    }
  }
export const ArchiveVariantDocument = gql`
    mutation ArchiveVariant($variantID: ID!) {
  archiveProductVariant(variantID: $variantID) {
    id
  }
}
    `;

  @Injectable({
    providedIn: 'root'
  })
  export class ArchiveVariantGQL extends Apollo.Mutation<ArchiveVariantMutation, ArchiveVariantMutationVariables> {
    document = ArchiveVariantDocument;
    
    constructor(apollo: Apollo.Apollo) {
      super(apollo);
    }
  }
export const MustInventoryDocument = gql`
    mutation MustInventory($productVariantID: ID!) {
  mustInventory(productVariantID: $productVariantID) {
    id
    code
    sku
    countryOfOrigin {
      label
    }
    countryHarmonizedCode {
      code
      country {
        id
        label
      }
    }
  }
}
    `;

  @Injectable({
    providedIn: 'root'
  })
  export class MustInventoryGQL extends Apollo.Mutation<MustInventoryMutation, MustInventoryMutationVariables> {
    document = MustInventoryDocument;
    
    constructor(apollo: Apollo.Apollo) {
      super(apollo);
    }
  }
export const UpdateInventoryItemDocument = gql`
    mutation UpdateInventoryItem($id: ID!, $input: UpdateInventoryItemInput!) {
  updateInventory(input: $input, iventoryItemID: $id) {
    ...InventoryItemFrag
  }
}
    ${InventoryItemFragFragmentDoc}`;

  @Injectable({
    providedIn: 'root'
  })
  export class UpdateInventoryItemGQL extends Apollo.Mutation<UpdateInventoryItemMutation, UpdateInventoryItemMutationVariables> {
    document = UpdateInventoryItemDocument;
    
    constructor(apollo: Apollo.Apollo) {
      super(apollo);
    }
  }
export const ProductSearchCountriesDocument = gql`
    query ProductSearchCountries($term: String!) {
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
  export class ProductSearchCountriesGQL extends Apollo.Query<ProductSearchCountriesQuery, ProductSearchCountriesQueryVariables> {
    document = ProductSearchCountriesDocument;
    
    constructor(apollo: Apollo.Apollo) {
      super(apollo);
    }
  }