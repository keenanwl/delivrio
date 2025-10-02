/* eslint-disable */
import * as Types from '../../../generated/graphql';

import { gql } from 'apollo-angular';
import { Injectable } from '@angular/core';
import * as Apollo from 'apollo-angular';
export type CompanyInfoQueryVariables = Types.Exact<{
  id: Types.Scalars['ID'];
}>;


export type CompanyInfoQuery = { tenant?: { name: string, invoiceReference?: string | null, defaultLanguage: { id: string, label: string }, adminContact?: { id: string, name: string, surname: string, phoneNumber: string, email: string } | null, billingContact?: { id: string, name: string, surname: string, phoneNumber: string, email: string } | null, companyAddress?: { id: string, company?: string | null, addressOne: string, addressTwo: string, city: string, state?: string | null, zip: string, country: { id: string, label: string, alpha2: string } } | null } | null, languages: { edges?: Array<{ node?: { id: string, label: string } | null } | null> | null }, countries: { edges?: Array<{ node?: { id: string, label: string, alpha2: string } | null } | null> | null } };

export type UpdateCompanyInfoMutationVariables = Types.Exact<{
  input: Types.UpdateTenantInput;
  defaultLanguage: Types.Scalars['ID'];
  adminContact: Types.CreateContactInput;
  billingContact: Types.CreateContactInput;
  address: Types.CreateAddressInput;
}>;


export type UpdateCompanyInfoMutation = { updateCompanyInfo?: { id: string } | null };

export type CompanyInfoSearchCountriesQueryVariables = Types.Exact<{
  term: Types.Scalars['String'];
}>;


export type CompanyInfoSearchCountriesQuery = { countries: { edges?: Array<{ node?: { id: string, label: string, alpha2: string } | null } | null> | null } };

export const CompanyInfoDocument = gql`
    query CompanyInfo($id: ID!) {
  tenant(id: $id) {
    name
    invoiceReference
    defaultLanguage {
      id
      label
    }
    adminContact {
      id
      name
      surname
      phoneNumber
      email
    }
    billingContact {
      id
      name
      surname
      phoneNumber
      email
    }
    companyAddress {
      id
      company
      addressOne
      addressTwo
      city
      state
      zip
      country {
        id
        label
        alpha2
      }
    }
  }
  languages {
    edges {
      node {
        id
        label
      }
    }
  }
  countries {
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
  export class CompanyInfoGQL extends Apollo.Query<CompanyInfoQuery, CompanyInfoQueryVariables> {
    document = CompanyInfoDocument;
    
    constructor(apollo: Apollo.Apollo) {
      super(apollo);
    }
  }
export const UpdateCompanyInfoDocument = gql`
    mutation UpdateCompanyInfo($input: UpdateTenantInput!, $defaultLanguage: ID!, $adminContact: CreateContactInput!, $billingContact: CreateContactInput!, $address: CreateAddressInput!) {
  updateCompanyInfo(
    input: $input
    defaultLanguage: $defaultLanguage
    adminContact: $adminContact
    billingContact: $billingContact
    address: $address
  ) {
    id
  }
}
    `;

  @Injectable({
    providedIn: 'root'
  })
  export class UpdateCompanyInfoGQL extends Apollo.Mutation<UpdateCompanyInfoMutation, UpdateCompanyInfoMutationVariables> {
    document = UpdateCompanyInfoDocument;
    
    constructor(apollo: Apollo.Apollo) {
      super(apollo);
    }
  }
export const CompanyInfoSearchCountriesDocument = gql`
    query CompanyInfoSearchCountries($term: String!) {
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
  export class CompanyInfoSearchCountriesGQL extends Apollo.Query<CompanyInfoSearchCountriesQuery, CompanyInfoSearchCountriesQueryVariables> {
    document = CompanyInfoSearchCountriesDocument;
    
    constructor(apollo: Apollo.Apollo) {
      super(apollo);
    }
  }