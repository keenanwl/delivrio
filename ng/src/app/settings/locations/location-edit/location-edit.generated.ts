/* eslint-disable */
import * as Types from '../../../../generated/graphql';

import { gql } from 'apollo-angular';
import { Injectable } from '@angular/core';
import * as Apollo from 'apollo-angular';
export type FetchLocationQueryVariables = Types.Exact<{
  id: Types.Scalars['ID'];
}>;


export type FetchLocationQuery = { location?: { name: string, address: { email: string, firstName: string, lastName: string, phoneNumber: string, vatNumber?: string | null, addressOne: string, addressTwo: string, zip: string, city: string, state?: string | null, company?: string | null, country: { id: string, label: string, alpha2: string } }, locationTags: Array<{ id: string, label: string }> } | null, locationTags: { edges?: Array<{ node?: { id: string, label: string } | null } | null> | null } };

export type LocationSearchCountriesQueryVariables = Types.Exact<{
  term: Types.Scalars['String'];
}>;


export type LocationSearchCountriesQuery = { countries: { edges?: Array<{ node?: { id: string, label: string, alpha2: string, region: Types.CountryRegion } | null } | null> | null } };

export type UpdateLocationMutationVariables = Types.Exact<{
  id: Types.Scalars['ID'];
  input: Types.UpdateLocationInput;
  inputAddress: Types.UpdateAddressInput;
}>;


export type UpdateLocationMutation = { updateLocation?: { id: string } | null };

export const FetchLocationDocument = gql`
    query FetchLocation($id: ID!) {
  location(id: $id) {
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
      state
      country {
        id
        label
        alpha2
      }
      company
    }
    locationTags {
      id
      label
    }
  }
  locationTags {
    edges {
      node {
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
  export class FetchLocationGQL extends Apollo.Query<FetchLocationQuery, FetchLocationQueryVariables> {
    document = FetchLocationDocument;
    
    constructor(apollo: Apollo.Apollo) {
      super(apollo);
    }
  }
export const LocationSearchCountriesDocument = gql`
    query LocationSearchCountries($term: String!) {
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
  export class LocationSearchCountriesGQL extends Apollo.Query<LocationSearchCountriesQuery, LocationSearchCountriesQueryVariables> {
    document = LocationSearchCountriesDocument;
    
    constructor(apollo: Apollo.Apollo) {
      super(apollo);
    }
  }
export const UpdateLocationDocument = gql`
    mutation UpdateLocation($id: ID!, $input: UpdateLocationInput!, $inputAddress: UpdateAddressInput!) {
  updateLocation(id: $id, input: $input, inputAddress: $inputAddress) {
    id
  }
}
    `;

  @Injectable({
    providedIn: 'root'
  })
  export class UpdateLocationGQL extends Apollo.Mutation<UpdateLocationMutation, UpdateLocationMutationVariables> {
    document = UpdateLocationDocument;
    
    constructor(apollo: Apollo.Apollo) {
      super(apollo);
    }
  }