/* eslint-disable */
import * as Types from '../../../../../generated/graphql';

import { gql } from 'apollo-angular';
import { Injectable } from '@angular/core';
import * as Apollo from 'apollo-angular';
export type DeliveryOptionsSearchCountriesQueryVariables = Types.Exact<{
  term: Types.Scalars['String'];
}>;


export type DeliveryOptionsSearchCountriesQuery = { countries: { edges?: Array<{ node?: { id: string, label: string, alpha2: string, region: Types.CountryRegion } | null } | null> | null } };

export type FetchDeliveryOptionRulesQueryVariables = Types.Exact<{
  id: Types.Scalars['ID'];
}>;


export type FetchDeliveryOptionRulesQuery = { deliveryRules: { edges?: Array<{ node?: { id: string, name: string, price: number, currency?: { id: string, display: string } | null, country?: Array<{ id: string, label: string, alpha2: string, region: Types.CountryRegion }> | null, deliveryRuleConstraintGroup?: Array<{ id: string, deliveryRuleConstraints?: Array<{ comparison: Types.DeliveryRuleConstraintComparison, selectedValue: { dayOfWeek?: Array<string> | null, text?: string | null, numeric?: number | null, numericRange?: Array<number> | null, timeOfDay?: Array<string> | null, ids?: Array<string> | null, values?: Array<string> | null } }> | null }> | null } | null } | null> | null }, currencies: { edges?: Array<{ node?: { id: string, display: string } | null } | null> | null } };

export type DeliveryOptionRuleFragFragment = { id: string, name: string, price: number, currency?: { id: string, display: string } | null, country?: Array<{ id: string, label: string, alpha2: string, region: Types.CountryRegion }> | null, deliveryRuleConstraintGroup?: Array<{ id: string, deliveryRuleConstraints?: Array<{ comparison: Types.DeliveryRuleConstraintComparison, selectedValue: { dayOfWeek?: Array<string> | null, text?: string | null, numeric?: number | null, numericRange?: Array<number> | null, timeOfDay?: Array<string> | null, ids?: Array<string> | null, values?: Array<string> | null } }> | null }> | null };

export type FetchRuleConstraintsQueryVariables = Types.Exact<{
  id: Types.Scalars['ID'];
}>;


export type FetchRuleConstraintsQuery = { constraintGroup?: { id: string, constraintLogic: Types.DeliveryRuleConstraintGroupConstraintLogic } | null, constraints?: Array<{ constraint?: { comparison: Types.DeliveryRuleConstraintComparison, propertyType: Types.DeliveryRuleConstraintPropertyType, selectedValue: { numeric?: number | null, numericRange?: Array<number> | null, text?: string | null, timeOfDay?: Array<string> | null, dayOfWeek?: Array<string> | null, ids?: Array<string> | null, values?: Array<string> | null } } | null, tags?: Array<{ id: string, name: string }> | null }> | null };

export type ReplaceRuleCountriesMutationVariables = Types.Exact<{
  ruleID: Types.Scalars['ID'];
  countries: Array<Types.Scalars['ID']> | Types.Scalars['ID'];
}>;


export type ReplaceRuleCountriesMutation = { replaceDeliveryRuleCountries: { country?: Array<{ id: string, label: string, alpha2: string, region: Types.CountryRegion }> | null } };

export type SearchProductTagsQueryVariables = Types.Exact<{
  term?: Types.InputMaybe<Types.Scalars['String']>;
}>;


export type SearchProductTagsQuery = { productTags: { edges?: Array<{ node?: { id: string, name: string, createdAt?: any | null } | null } | null> | null } };

export type CreateConstraintGroupMutationVariables = Types.Exact<{
  input: Types.CreateDeliveryRuleConstraintGroupInput;
}>;


export type CreateConstraintGroupMutation = { createDeliveryRuleConstraintGroup: { id: string } };

export type CreateConstraintGroupConstraintsMutationVariables = Types.Exact<{
  ruleID: Types.Scalars['ID'];
  logicType: Types.DeliveryRuleConstraintGroupConstraintLogic;
  input: Array<Types.InputMaybe<Types.CreateDeliveryRuleConstraintInput>> | Types.InputMaybe<Types.CreateDeliveryRuleConstraintInput>;
}>;


export type CreateConstraintGroupConstraintsMutation = { createDeliveryRuleConstraintGroupConstraints?: Array<{ constraint?: { id: string } | null }> | null };

export type ReplaceConstraintGroupConstraintsMutationVariables = Types.Exact<{
  deliveryGroupId: Types.Scalars['ID'];
  logicType: Types.DeliveryRuleConstraintGroupConstraintLogic;
  input: Array<Types.InputMaybe<Types.CreateDeliveryRuleConstraintInput>> | Types.InputMaybe<Types.CreateDeliveryRuleConstraintInput>;
}>;


export type ReplaceConstraintGroupConstraintsMutation = { replaceDeliveryRuleConstraintGroupConstraints?: Array<{ constraint?: { id: string } | null }> | null };

export type DeleteRuleMutationVariables = Types.Exact<{
  ruleID: Types.Scalars['ID'];
}>;


export type DeleteRuleMutation = { deleteDeliveryRule: Array<{ id: string, name: string, price: number, currency?: { id: string, display: string } | null, country?: Array<{ id: string, label: string, alpha2: string, region: Types.CountryRegion }> | null, deliveryRuleConstraintGroup?: Array<{ id: string, deliveryRuleConstraints?: Array<{ comparison: Types.DeliveryRuleConstraintComparison, selectedValue: { dayOfWeek?: Array<string> | null, text?: string | null, numeric?: number | null, numericRange?: Array<number> | null, timeOfDay?: Array<string> | null, ids?: Array<string> | null, values?: Array<string> | null } }> | null }> | null } | null> };

export type DeleteConstraintGroupConstraintsMutationVariables = Types.Exact<{
  groupID: Types.Scalars['ID'];
}>;


export type DeleteConstraintGroupConstraintsMutation = { deleteDeliveryRuleConstraintGroupConstraints: Array<{ id: string, name: string, price: number, currency?: { id: string, display: string } | null, country?: Array<{ id: string, label: string, alpha2: string, region: Types.CountryRegion }> | null, deliveryRuleConstraintGroup?: Array<{ id: string, deliveryRuleConstraints?: Array<{ comparison: Types.DeliveryRuleConstraintComparison, selectedValue: { dayOfWeek?: Array<string> | null, text?: string | null, numeric?: number | null, numericRange?: Array<number> | null, timeOfDay?: Array<string> | null, ids?: Array<string> | null, values?: Array<string> | null } }> | null }> | null } | null> };

export type CreateDeliveryRuleMutationVariables = Types.Exact<{
  input: Types.CreateDeliveryRuleInput;
}>;


export type CreateDeliveryRuleMutation = { createDeliveryRule: { id: string } };

export type UpdateDeliveryRulePriceMutationVariables = Types.Exact<{
  deliveryRuleID: Types.Scalars['ID'];
  val: Types.UpdateDeliveryRuleInput;
}>;


export type UpdateDeliveryRulePriceMutation = { updateDeliveryRule: { id: string, name: string, price: number, currency?: { id: string, display: string } | null, country?: Array<{ id: string, label: string, alpha2: string, region: Types.CountryRegion }> | null, deliveryRuleConstraintGroup?: Array<{ id: string, deliveryRuleConstraints?: Array<{ comparison: Types.DeliveryRuleConstraintComparison, selectedValue: { dayOfWeek?: Array<string> | null, text?: string | null, numeric?: number | null, numericRange?: Array<number> | null, timeOfDay?: Array<string> | null, ids?: Array<string> | null, values?: Array<string> | null } }> | null }> | null } };

export const DeliveryOptionRuleFragFragmentDoc = gql`
    fragment DeliveryOptionRuleFrag on DeliveryRule {
  id
  name
  price
  currency {
    id
    display
  }
  country {
    id
    label
    alpha2
    region
  }
  deliveryRuleConstraintGroup {
    id
    deliveryRuleConstraints {
      selectedValue {
        dayOfWeek
        text
        numeric
        numericRange
        timeOfDay
        ids
        values
      }
      comparison
    }
  }
}
    `;
export const DeliveryOptionsSearchCountriesDocument = gql`
    query DeliveryOptionsSearchCountries($term: String!) {
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
  export class DeliveryOptionsSearchCountriesGQL extends Apollo.Query<DeliveryOptionsSearchCountriesQuery, DeliveryOptionsSearchCountriesQueryVariables> {
    document = DeliveryOptionsSearchCountriesDocument;
    
    constructor(apollo: Apollo.Apollo) {
      super(apollo);
    }
  }
export const FetchDeliveryOptionRulesDocument = gql`
    query FetchDeliveryOptionRules($id: ID!) {
  deliveryRules(where: {hasDeliveryOptionWith: {id: $id}}) {
    edges {
      node {
        ...DeliveryOptionRuleFrag
      }
    }
  }
  currencies {
    edges {
      node {
        id
        display
      }
    }
  }
}
    ${DeliveryOptionRuleFragFragmentDoc}`;

  @Injectable({
    providedIn: 'root'
  })
  export class FetchDeliveryOptionRulesGQL extends Apollo.Query<FetchDeliveryOptionRulesQuery, FetchDeliveryOptionRulesQueryVariables> {
    document = FetchDeliveryOptionRulesDocument;
    
    constructor(apollo: Apollo.Apollo) {
      super(apollo);
    }
  }
export const FetchRuleConstraintsDocument = gql`
    query FetchRuleConstraints($id: ID!) {
  constraintGroup(id: $id) {
    id
    constraintLogic
  }
  constraints(groupID: $id) {
    constraint {
      comparison
      propertyType
      selectedValue {
        numeric
        numericRange
        text
        timeOfDay
        dayOfWeek
        ids
        values
      }
    }
    tags {
      id
      name
    }
  }
}
    `;

  @Injectable({
    providedIn: 'root'
  })
  export class FetchRuleConstraintsGQL extends Apollo.Query<FetchRuleConstraintsQuery, FetchRuleConstraintsQueryVariables> {
    document = FetchRuleConstraintsDocument;
    
    constructor(apollo: Apollo.Apollo) {
      super(apollo);
    }
  }
export const ReplaceRuleCountriesDocument = gql`
    mutation ReplaceRuleCountries($ruleID: ID!, $countries: [ID!]!) {
  replaceDeliveryRuleCountries(ruleID: $ruleID, countries: $countries) {
    country {
      id
      label
      alpha2
      region
    }
  }
}
    `;

  @Injectable({
    providedIn: 'root'
  })
  export class ReplaceRuleCountriesGQL extends Apollo.Mutation<ReplaceRuleCountriesMutation, ReplaceRuleCountriesMutationVariables> {
    document = ReplaceRuleCountriesDocument;
    
    constructor(apollo: Apollo.Apollo) {
      super(apollo);
    }
  }
export const SearchProductTagsDocument = gql`
    query SearchProductTags($term: String) {
  productTags(where: {nameContainsFold: $term}) {
    edges {
      node {
        id
        name
        createdAt
      }
    }
  }
}
    `;

  @Injectable({
    providedIn: 'root'
  })
  export class SearchProductTagsGQL extends Apollo.Query<SearchProductTagsQuery, SearchProductTagsQueryVariables> {
    document = SearchProductTagsDocument;
    
    constructor(apollo: Apollo.Apollo) {
      super(apollo);
    }
  }
export const CreateConstraintGroupDocument = gql`
    mutation CreateConstraintGroup($input: CreateDeliveryRuleConstraintGroupInput!) {
  createDeliveryRuleConstraintGroup(input: $input) {
    id
  }
}
    `;

  @Injectable({
    providedIn: 'root'
  })
  export class CreateConstraintGroupGQL extends Apollo.Mutation<CreateConstraintGroupMutation, CreateConstraintGroupMutationVariables> {
    document = CreateConstraintGroupDocument;
    
    constructor(apollo: Apollo.Apollo) {
      super(apollo);
    }
  }
export const CreateConstraintGroupConstraintsDocument = gql`
    mutation CreateConstraintGroupConstraints($ruleID: ID!, $logicType: DeliveryRuleConstraintGroupConstraintLogic!, $input: [CreateDeliveryRuleConstraintInput]!) {
  createDeliveryRuleConstraintGroupConstraints(
    deliveryRuleId: $ruleID
    logicType: $logicType
    input: $input
  ) {
    constraint {
      id
    }
  }
}
    `;

  @Injectable({
    providedIn: 'root'
  })
  export class CreateConstraintGroupConstraintsGQL extends Apollo.Mutation<CreateConstraintGroupConstraintsMutation, CreateConstraintGroupConstraintsMutationVariables> {
    document = CreateConstraintGroupConstraintsDocument;
    
    constructor(apollo: Apollo.Apollo) {
      super(apollo);
    }
  }
export const ReplaceConstraintGroupConstraintsDocument = gql`
    mutation ReplaceConstraintGroupConstraints($deliveryGroupId: ID!, $logicType: DeliveryRuleConstraintGroupConstraintLogic!, $input: [CreateDeliveryRuleConstraintInput]!) {
  replaceDeliveryRuleConstraintGroupConstraints(
    deliveryGroupId: $deliveryGroupId
    logicType: $logicType
    input: $input
  ) {
    constraint {
      id
    }
  }
}
    `;

  @Injectable({
    providedIn: 'root'
  })
  export class ReplaceConstraintGroupConstraintsGQL extends Apollo.Mutation<ReplaceConstraintGroupConstraintsMutation, ReplaceConstraintGroupConstraintsMutationVariables> {
    document = ReplaceConstraintGroupConstraintsDocument;
    
    constructor(apollo: Apollo.Apollo) {
      super(apollo);
    }
  }
export const DeleteRuleDocument = gql`
    mutation DeleteRule($ruleID: ID!) {
  deleteDeliveryRule(deliveryRuleID: $ruleID) {
    ...DeliveryOptionRuleFrag
  }
}
    ${DeliveryOptionRuleFragFragmentDoc}`;

  @Injectable({
    providedIn: 'root'
  })
  export class DeleteRuleGQL extends Apollo.Mutation<DeleteRuleMutation, DeleteRuleMutationVariables> {
    document = DeleteRuleDocument;
    
    constructor(apollo: Apollo.Apollo) {
      super(apollo);
    }
  }
export const DeleteConstraintGroupConstraintsDocument = gql`
    mutation DeleteConstraintGroupConstraints($groupID: ID!) {
  deleteDeliveryRuleConstraintGroupConstraints(deliveryGroupId: $groupID) {
    ...DeliveryOptionRuleFrag
  }
}
    ${DeliveryOptionRuleFragFragmentDoc}`;

  @Injectable({
    providedIn: 'root'
  })
  export class DeleteConstraintGroupConstraintsGQL extends Apollo.Mutation<DeleteConstraintGroupConstraintsMutation, DeleteConstraintGroupConstraintsMutationVariables> {
    document = DeleteConstraintGroupConstraintsDocument;
    
    constructor(apollo: Apollo.Apollo) {
      super(apollo);
    }
  }
export const CreateDeliveryRuleDocument = gql`
    mutation createDeliveryRule($input: CreateDeliveryRuleInput!) {
  createDeliveryRule(input: $input) {
    id
  }
}
    `;

  @Injectable({
    providedIn: 'root'
  })
  export class CreateDeliveryRuleGQL extends Apollo.Mutation<CreateDeliveryRuleMutation, CreateDeliveryRuleMutationVariables> {
    document = CreateDeliveryRuleDocument;
    
    constructor(apollo: Apollo.Apollo) {
      super(apollo);
    }
  }
export const UpdateDeliveryRulePriceDocument = gql`
    mutation UpdateDeliveryRulePrice($deliveryRuleID: ID!, $val: UpdateDeliveryRuleInput!) {
  updateDeliveryRule(deliveryRuleID: $deliveryRuleID, val: $val) {
    ...DeliveryOptionRuleFrag
  }
}
    ${DeliveryOptionRuleFragFragmentDoc}`;

  @Injectable({
    providedIn: 'root'
  })
  export class UpdateDeliveryRulePriceGQL extends Apollo.Mutation<UpdateDeliveryRulePriceMutation, UpdateDeliveryRulePriceMutationVariables> {
    document = UpdateDeliveryRulePriceDocument;
    
    constructor(apollo: Apollo.Apollo) {
      super(apollo);
    }
  }