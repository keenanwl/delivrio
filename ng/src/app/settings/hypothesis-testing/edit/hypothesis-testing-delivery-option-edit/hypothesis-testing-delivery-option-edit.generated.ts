/* eslint-disable */
import * as Types from '../../../../../generated/graphql';

import { gql } from 'apollo-angular';
import { Injectable } from '@angular/core';
import * as Apollo from 'apollo-angular';
export type FetchHypothesisTestQueryVariables = Types.Exact<{
  id: Types.Scalars['ID'];
}>;


export type FetchHypothesisTestQuery = { hypothesisTest: { id: string, name: string, active: boolean, hypothesisTestDeliveryOption?: { id: string, byIntervalRotation: boolean, byOrder: boolean, randomizeWithinGroupSort: boolean, rotationIntervalHours: number, deliveryOptionGroupOne?: Array<{ id: string, name: string, description?: string | null, hideDeliveryOption?: boolean | null }> | null, deliveryOptionGroupTwo?: Array<{ id: string, name: string, description?: string | null, hideDeliveryOption?: boolean | null }> | null } | null }, unassignedDeliveryOptions: Array<{ id: string, name: string, description?: string | null, hideDeliveryOption?: boolean | null }> };

export type UpdateHypothesisTestDeliveryOptionMutationVariables = Types.Exact<{
  id: Types.Scalars['ID'];
  input: Types.UpdateHypothesisTestInput;
  inputDeliveryOption: Types.UpdateHypothesisTestDeliveryOptionInput;
}>;


export type UpdateHypothesisTestDeliveryOptionMutation = { updateHypothesisTestDeliveryOption: { id: string } };

export const FetchHypothesisTestDocument = gql`
    query FetchHypothesisTest($id: ID!) {
  hypothesisTest(id: $id) {
    id
    name
    active
    hypothesisTestDeliveryOption {
      id
      byIntervalRotation
      byOrder
      randomizeWithinGroupSort
      rotationIntervalHours
      deliveryOptionGroupOne {
        id
        name
        description
        hideDeliveryOption
      }
      deliveryOptionGroupTwo {
        id
        name
        description
        hideDeliveryOption
      }
    }
  }
  unassignedDeliveryOptions(hypothesisTestID: $id) {
    id
    name
    description
    hideDeliveryOption
  }
}
    `;

  @Injectable({
    providedIn: 'root'
  })
  export class FetchHypothesisTestGQL extends Apollo.Query<FetchHypothesisTestQuery, FetchHypothesisTestQueryVariables> {
    document = FetchHypothesisTestDocument;
    
    constructor(apollo: Apollo.Apollo) {
      super(apollo);
    }
  }
export const UpdateHypothesisTestDeliveryOptionDocument = gql`
    mutation UpdateHypothesisTestDeliveryOption($id: ID!, $input: UpdateHypothesisTestInput!, $inputDeliveryOption: UpdateHypothesisTestDeliveryOptionInput!) {
  updateHypothesisTestDeliveryOption(
    id: $id
    input: $input
    inputDeliveryOption: $inputDeliveryOption
  ) {
    id
  }
}
    `;

  @Injectable({
    providedIn: 'root'
  })
  export class UpdateHypothesisTestDeliveryOptionGQL extends Apollo.Mutation<UpdateHypothesisTestDeliveryOptionMutation, UpdateHypothesisTestDeliveryOptionMutationVariables> {
    document = UpdateHypothesisTestDeliveryOptionDocument;
    
    constructor(apollo: Apollo.Apollo) {
      super(apollo);
    }
  }