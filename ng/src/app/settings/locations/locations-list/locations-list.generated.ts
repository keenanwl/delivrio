/* eslint-disable */
import * as Types from '../../../../generated/graphql';

import { gql } from 'apollo-angular';
import { Injectable } from '@angular/core';
import * as Apollo from 'apollo-angular';
export type ListLocationsQueryVariables = Types.Exact<{ [key: string]: never; }>;


export type ListLocationsQuery = { locations: { edges?: Array<{ node?: { id: string, name: string, address: { email: string, firstName: string, lastName: string, addressOne: string, addressTwo: string, zip: string, city: string, country: { id: string, label: string, alpha2: string } }, locationTags: Array<{ id: string, label: string }> } | null } | null> | null } };

export type CreateLocationMutationVariables = Types.Exact<{
  input: Types.CreateLocationInput;
  inputAddress: Types.CreateAddressInput;
}>;


export type CreateLocationMutation = { createLocation?: { id: string } | null };

export const ListLocationsDocument = gql`
    query ListLocations {
  locations {
    edges {
      node {
        id
        name
        address {
          email
          firstName
          lastName
          addressOne
          addressTwo
          zip
          city
          country {
            id
            label
            alpha2
          }
        }
        locationTags {
          id
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
  export class ListLocationsGQL extends Apollo.Query<ListLocationsQuery, ListLocationsQueryVariables> {
    document = ListLocationsDocument;
    
    constructor(apollo: Apollo.Apollo) {
      super(apollo);
    }
  }
export const CreateLocationDocument = gql`
    mutation CreateLocation($input: CreateLocationInput!, $inputAddress: CreateAddressInput!) {
  createLocation(input: $input, inputAddress: $inputAddress) {
    id
  }
}
    `;

  @Injectable({
    providedIn: 'root'
  })
  export class CreateLocationGQL extends Apollo.Mutation<CreateLocationMutation, CreateLocationMutationVariables> {
    document = CreateLocationDocument;
    
    constructor(apollo: Apollo.Apollo) {
      super(apollo);
    }
  }