/* eslint-disable */
import * as Types from '../../../generated/graphql';

import { gql } from 'apollo-angular';
export type TimelineViewerFragment = { id: string, createdAt: any, user?: { id: string, name?: string | null } | null, orderHistory?: Array<{ id: string, type: Types.OrderHistoryType, description: string, order: { id: string, orderPublicID: string } }> | null, shipmentHistory?: Array<{ id: string, type: Types.ShipmentHistoryType }> | null, returnColliHistory?: Array<{ id: string, description: string, type: Types.ReturnColliHistoryType }> | null };

export const TimelineViewerFragmentDoc = gql`
    fragment TimelineViewer on ChangeHistory {
  id
  createdAt
  user {
    id
    name
  }
  orderHistory {
    id
    type
    description
    order {
      id
      orderPublicID
    }
  }
  shipmentHistory {
    id
    type
  }
  returnColliHistory {
    id
    description
    type
  }
}
    `;